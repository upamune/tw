package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/mgutz/ansi"
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
	Usage: "tw [tweet] TEXT...",
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
	Usage: "tw fav TWEET_ID",
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
	Usage: "tw timeline [NUM]",
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
	api := doOauth()
	var text string
	for i := 0; i < len(c.Args()); i++ {
		text += c.Args()[i]
		if i == len(c.Args())-1 {
			continue
		}
		text += " "
	}
	if text == "" {
		log.Fatal("ツイートする文字列を指定してください")
	}
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Tweeted : ", tweet.Text)
	}
}

func doRt(c *cli.Context) {
}

func doFav(c *cli.Context) {
	api := doOauth()

	for i := 0; i < len(c.Args()); i++ {
		tweetID, _ := strconv.ParseInt(c.Args()[i], 10, 64)
		_, err := api.Favorite(tweetID)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func doDel(c *cli.Context) {
}

func doSearch(c *cli.Context) {
}

func doTimeline(c *cli.Context) {
	api := doOauth()
	defer api.Close()
	timeline, err := api.GetHomeTimeline(nil)
	if err != nil {
		panic(err)
	}

	for _, tweet := range timeline {
		user := tweet.User.Name
		screenName := tweet.User.ScreenName
		user += "(@" + screenName + ")"

		blue := ansi.ColorCode("blue")
		reset := ansi.ColorCode("reset")

		fmt.Println(blue, user, ":", reset, tweet.Text)
	}

}

func doDm(c *cli.Context) {
}

func doReply(c *cli.Context) {
	api := doOauth()
	defer api.Close()
	mentions, err := api.GetMentionsTimeline(nil)
	if err != nil {
		panic(err)
	}

	for _, mention := range mentions {
		red := ansi.ColorCode("red")
		reset := ansi.ColorCode("reset")
		fromUser := mention.User
		from := fromUser.Name + "(" + fromUser.ScreenName + ")"
		fmt.Println(red, from, reset, ":", mention.Text)
	}
}
