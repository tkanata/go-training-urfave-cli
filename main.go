package main

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "yaruzoyaruzoyaruzo",
				Value: "empty",
				Usage: "say やるぞやるぞやるぞ",
			},
		},
		Action: func(c *cli.Context) error {
			// fmt.Println("やるぞやるぞやるぞ")
			yaruzo := c.String("yaruzoyaruzoyaruzo")
			if yaruzo == "やるぞやるぞやるぞ" {
				fmt.Println("やるぞやるぞやるぞ")
			} else {
				fmt.Printf("%v\n", yaruzo)
			}
			return nil
		},
	}

	app.Run(os.Args)
}
