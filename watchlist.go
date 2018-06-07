package tvtime

import "fmt"

// DisplayWatchlist displays the upcoming episodes
func DisplayWatchlist() error {
	watchlist, err := GetWatchlist()
	if err != nil {
		return err
	}
	for _, episode := range watchlist {
		fmt.Printf("%s S%02dE%02d - %s on %s at %s\n", episode.Show.Name, episode.SeasonNumber, episode.Number, episode.Name, episode.AirDate, episode.AirTime)
	}
	return nil
}
