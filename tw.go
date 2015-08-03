package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tw"
	app.Version = Version
	app.Usage = ""
	app.Author = "upamune"
	app.Email = "jajkeqos@gmail.com"
	app.Flags = GlobalFlags

	app.Action = commandTweet.Action
	app.Commands = Commands
	app.Run(os.Args)
}
