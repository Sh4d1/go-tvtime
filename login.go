package tvtime

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/spin"
)

var tvtimeURL = "https://api.tvtime.com/v1/"

var clientID = "va0D2CEfSPNNlLoYMYYT"
var clientSecret = "RF51gSEZBJAbLXmEUCZ8thJAwJPAyQSafCQCyqOt"

var deviceCodeURL = tvtimeURL + "oauth/device/code"
var deviceCodeURLValues = url.Values{"client_id": {clientID}}

type DeviceCodeResponse struct {
	Result          string `json:"result"`
	Message         string `json:"message"`
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURL string `json:"verification_url"`
	ExpiresIn       int64  `json:"expires_in"`
	Interval        int64  `json:"interval"`
}

var accessTokenURL = tvtimeURL + "oauth/access_token"
var accessTokenURLValues = url.Values{"client_id": {clientID}, "client_secret": {clientSecret}}

type AccessTokenResponse struct {
	Result      string `json:"result"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

func Login() error {

	if HaveSettings() {
		fmt.Printf("You are already logged in, do you wish to logout and login? [y/N] ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if !strings.HasPrefix(input, "y") && !strings.HasPrefix(input, "Y") {
			return nil
		}
	}

	respDeviceCode, err := http.PostForm(deviceCodeURL, deviceCodeURLValues)
	if err != nil {
		return err
	}
	defer respDeviceCode.Body.Close()

	bodyDeviceCode, err := ioutil.ReadAll(respDeviceCode.Body)
	if err != nil {
		return err
	}

	var deviceCodeResp = new(DeviceCodeResponse)
	err = json.Unmarshal(bodyDeviceCode, &deviceCodeResp)
	if err != nil {
		return err
	}

	fmt.Printf("Please connect to %s and enter the following code: %s\nPress enter when its done\n", deviceCodeResp.VerificationURL, deviceCodeResp.UserCode)

	accessTokenURLValues.Add("code", deviceCodeResp.DeviceCode)
	spin := spin.New("%s Waiting...")
	spin.Start()

	var accessTokenResp = new(AccessTokenResponse)
	var counter = 0
	for {
		counter++
		respAccessToken, err := http.PostForm(accessTokenURL, accessTokenURLValues)
		if err != nil {
			return err
		}
		defer respAccessToken.Body.Close()
		bodyAccessToken, err := ioutil.ReadAll(respAccessToken.Body)
		fmt.Println(string(bodyAccessToken))

		if err != nil {
			return err
		}

		err = json.Unmarshal(bodyAccessToken, &accessTokenResp)
		if err != nil {
			return err
		}

		if accessTokenResp.Result == "OK" || counter > 12 {
			spin.Stop()
			break
		}

		time.Sleep(100 * 100 * time.Millisecond)
	}

	if accessTokenResp.Result != "OK" {
		return errors.New("could not login, please try again")
	}

	var settings Settings
	settings.AccessToken = accessTokenResp.AccessToken

	err = WriteSettings(settings)
	if err != nil {
		return err
	}

	fmt.Println("Logged in!")
	return nil
}
