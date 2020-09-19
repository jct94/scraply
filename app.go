// go build -ldflags "-X main.wordLen=6"
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	app        *cli.App
	appVersion = "unknown"
)

func main() {
	app = cli.NewApp()

	app.Name = "scraply"
	app.Version = appVersion
	app.Description = "a very simple but highly scalable scraping platform"
	app.Copyright = "scraply is a software created by Mohamed Al Ashaal <m7medalash3al@gmail.com> under Apache License 2.0"
	app.Usage = app.Description
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

func registerCommand(cmd *cli.Command) {
	app.Commands = append(app.Commands, cmd)
}
