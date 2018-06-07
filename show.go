package tvtime

type Show struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	SeenEpisodes  int    `json:"seen_episodes"`
	AiredEpisodes int    `json:"aired_episodes"`
}

type Episode struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Number       int    `json:"number"`
	SeasonNumber int    `json:"season_number"`
	AirDate      string `json:"air_date"`
	AirTime      string `json:"air_time"`
	Network      string `json:"network"`
	IsNew        bool   `json:"is_new"`
	Seen         bool   `json:"seen"`
	Show         Show   `json:"show"`
}
