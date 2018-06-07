package main

import (
	"log"
	"os"

	tvtime "github.com/Sh4d1/go-tvtime"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()

	app.Name = "tvtime"
	app.Version = version
	app.Author = "Patrik Cyvoct (patrik@ptrk.io)"
	app.Usage = "tvtime"

	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login to TV Time",
			Action: func(c *cli.Context) error {
				return tvtime.Login()
			},
		},
		{
			Name:  "user",
			Usage: "Display current user",
			Action: func(c *cli.Context) error {
				return tvtime.DisplayUser()
			},
		},
		{
			Name:  "upcoming",
			Usage: "Display upcoming episodes",
			Action: func(c *cli.Context) error {
				return tvtime.DisplayUpcoming()
			},
		},
		{
			Name:  "watchlist",
			Usage: "Display watchlist",
			Action: func(c *cli.Context) error {
				return tvtime.DisplayWatchlist()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
