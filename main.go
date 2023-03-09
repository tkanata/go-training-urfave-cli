package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "やるぞやるぞやるぞ",
		Usage: "say やるぞやるぞやるぞ",
		Action: func(c *cli.Context) error {
			fmt.Println("やるぞやるぞやるぞ")
			return nil
		},
	}

	app.Run(os.Args)
}
