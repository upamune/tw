package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandTweet,
	commandRt,
	commandFav,
	commandDel,
	commandSearch,
	commandTimeline,
	commandDm,
	commandReply,
}

var commandTweet = cli.Command{
	Name:  "tweet",
	Usage: "",
	Description: `
`,
	Action: doTweet,
}

var commandRt = cli.Command{
	Name:  "rt",
	Usage: "",
	Description: `
`,
	Action: doRt,
}

var commandFav = cli.Command{
	Name:  "fav",
	Usage: "",
	Description: `
`,
	Action: doFav,
}

var commandDel = cli.Command{
	Name:  "del",
	Usage: "",
	Description: `
`,
	Action: doDel,
}

var commandSearch = cli.Command{
	Name:  "search",
	Usage: "",
	Description: `
`,
	Action: doSearch,
}

var commandTimeline = cli.Command{
	Name:  "timeline",
	Usage: "",
	Description: `
`,
	Action: doTimeline,
}

var commandDm = cli.Command{
	Name:  "dm",
	Usage: "",
	Description: `
`,
	Action: doDm,
}

var commandReply = cli.Command{
	Name:  "reply",
	Usage: "",
	Description: `
`,
	Action: doReply,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doTweet(c *cli.Context) {
}

func doRt(c *cli.Context) {
}

func doFav(c *cli.Context) {
}

func doDel(c *cli.Context) {
}

func doSearch(c *cli.Context) {
}

func doTimeline(c *cli.Context) {
}

func doDm(c *cli.Context) {
}

func doReply(c *cli.Context) {
}
