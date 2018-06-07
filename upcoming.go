package tvtime

import "fmt"

func DisplayUpcoming() error {
	upcoming, err := GetUpcoming()
	if err != nil {
		return err
	}
	for _, episode := range upcoming {
		fmt.Printf("%s S%02dE%02d - %s on %s at %s\n", episode.Show.Name, episode.SeasonNumber, episode.Number, episode.Name, episode.AirDate, episode.AirTime)
	}
	//	fmt.Println(upcoming)
	return nil
}
