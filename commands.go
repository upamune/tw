package main

import (
	"fmt"
	"log"
	"net/url"
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
	Usage: "tw rt TWEET_ID",
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
	Usage: "tw del TWEET_ID",
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
	Name:    "timeline",
	Aliases: []string{"tl"},
	Usage:   "tw timeline [NUM]",
	Description: `
`,
	Action: doTimeline,
}

var commandDm = cli.Command{
	Name:  "dm",
	Usage: "tw dm [TEXT...]",
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
	defer api.Close()

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
	api := doOauth()
	defer api.Close()

	for i := 0; i < len(c.Args()); i++ {
		tweetID, _ := strconv.ParseInt(c.Args()[i], 10, 64)
		tweet, err := api.Retweet(tweetID, true)
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Println(tweet.Text)
	}
}

func doFav(c *cli.Context) {
	api := doOauth()
	defer api.Close()

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
	api := doOauth()
	defer api.Close()

	for i := 0; i < len(c.Args()); i++ {
		tweetID, _ := strconv.ParseInt(c.Args()[i], 10, 64)
		tweet, err := api.DeleteTweet(tweetID, true)
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Println("Del:", tweet.Text)
	}
}

func doSearch(c *cli.Context) {
}

func doTimeline(c *cli.Context) {
	api := doOauth()
	defer api.Close()

	var cnt string
	if len(c.Args()) > 0 {
		_, err := strconv.Atoi(c.Args()[0])
		if err != nil {
			panic(err)
		}
		cnt = c.Args()[0]
	}

	v := url.Values{}
	v.Add("count", cnt)
	timeline, err := api.GetHomeTimeline(v)
	if err != nil {
		panic(err)
	}

	for _, tweet := range timeline {
		user := tweet.User.Name
		screenName := tweet.User.ScreenName
		user += "(@" + screenName + ")"

		blue := ansi.ColorCode("blue")
		reset := ansi.ColorCode("reset")

		fmt.Println(blue, user, ":", reset, tweet.Text, tweet.Id)
	}

}

func doDm(c *cli.Context) {
	api := doOauth()
	defer api.Close()

	screenName := c.Args()[0]
	var message string
	for i := 1; i < len(c.Args()); i++ {
		message += c.Args()[i]
		if i == len(c.Args())-1 {
			continue
		}
		message += " "
	}
	if message == "" {
		log.Fatal("DMする文字列を指定してください")
	}
	dm, err := api.PostDMToScreenName(message, screenName)
	if err != nil {
		log.Fatal(err)
	} else {
		red := ansi.ColorCode("red")
		reset := ansi.ColorCode("reset")
		fmt.Println("TO:", red, screenName, reset, dm.Text)
	}

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
