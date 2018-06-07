package tvtime

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var userURL = "https://api.tvtime.com/v1/user"
var upcomingURL = "https://api.tvtime.com/v1/agenda"
var watchlistURL = "https://api.tvtime.com/v1/to_watch"

// GetUserResponse stores the get response of the user
type GetUserResponse struct {
	Result  string `json:"result"`
	User    User   `json:"user"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// GetUpcomingResponse stores the get response of the upcoming
type GetUpcomingResponse struct {
	Result   string    `json:"result"`
	Episodes []Episode `json:"episodes"`
	Message  string    `json:"message"`
	Code     int       `json:"code"`
}

// GetWatchlistResponse stores the get response of the watchlist
type GetWatchlistResponse struct {
	Result   string    `json:"result"`
	Episodes []Episode `json:"episodes"`
	Message  string    `json:"message"`
	Code     int       `json:"code"`
}

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	token, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("TVST_ACCESS_TOKEN", token)
	return client.Do(req)
}

// GetUser returns the user
func GetUser() (User, error) {
	resp, err := getRequest(userURL)
	if err != nil {
		return User{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var userResp = new(GetUserResponse)
	err = json.Unmarshal(body, &userResp)
	if err != nil {
		return User{}, err
	}
	return userResp.User, nil
}

// GetUpcoming returns the upcoming episodes
func GetUpcoming() ([]Episode, error) {
	resp, err := getRequest(upcomingURL)
	if err != nil {
		return []Episode{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Episode{}, err
	}

	var upcomingResp = new(GetUpcomingResponse)
	err = json.Unmarshal(body, &upcomingResp)
	if err != nil {
		return []Episode{}, err
	}
	return upcomingResp.Episodes, nil
}

// GetWatchlist return the watchlist
func GetWatchlist() ([]Episode, error) {
	resp, err := getRequest(watchlistURL)
	if err != nil {
		return []Episode{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Episode{}, err
	}

	var watchlistResp = new(GetWatchlistResponse)
	err = json.Unmarshal(body, &watchlistResp)
	if err != nil {
		return []Episode{}, err
	}
	return watchlistResp.Episodes, nil
}
