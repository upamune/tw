package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
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
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "ENV_PIPE",
			Name:   "pipe",
			Usage:  "Tweet by stdin",
		},
	},
	Description: `
	`,
	Action: doTweet,
}

var commandRt = cli.Command{
	Name:  "rt",
	Usage: "tw rt TWEET_ID",
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "ENV_PIPE",
			Name:   "pipe",
			Usage:  "Retweet tweet by stdin",
		},
	},
	Description: `
	`,
	Action: doRt,
}

var commandFav = cli.Command{
	Name:  "fav",
	Usage: "tw fav TWEET_ID",
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "ENV_PIPE",
			Name:   "pipe",
			Usage:  "Favorite tweet by stdin",
		},
	},
	Description: `
	`,
	Action: doFav,
}

var commandDel = cli.Command{
	Name:  "del",
	Usage: "tw del TWEET_ID",
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "ENV_PIPE",
			Name:   "pipe",
			Usage:  "Favorite tweet by stdin",
		},
	},
	Description: `
	`,
	Action: doDel,
}

var commandSearch = cli.Command{
	Name:  "search",
	Usage: "tw serach QUERY",
	Description: `
	`,
	Action: doSearch,
}

var commandTimeline = cli.Command{
	Name:    "timeline",
	Aliases: []string{"tl"},
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "ENV_WITH",
			Name:   "with-id",
			Usage:  "Show timeline with Tweet ID",
		},
		cli.BoolFlag{
			Name:  "stream",
			Usage: "Show user stream timeline",
		},
		cli.StringFlag{
			Name:  "user",
			Usage: "Show timeline with screen name",
		},
	},
	Usage: "tw timeline [NUM]",
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
	Usage: "tw reply",
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
	if c.Bool("pipe") {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	} else {
		for i := 0; i < len(c.Args()); i++ {
			text += c.Args()[i]
			if i == len(c.Args())-1 {
				continue
			}
			text += " "
		}
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
	if c.Bool("pipe") {
		var stdin string

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin += scanner.Text()
		}
		tweetID, _ := strconv.ParseInt(stdin, 10, 64)
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		_, err := api.Retweet(tweetID, true)
		if err != nil {
			log.Fatal(err)
		}
		tweet, err := api.GetTweet(tweetID, nil)
		if err != nil {
			log.Fatal(err)
		}
		dispTweet(tweet, "Retweeted: ", "blue")
	} else {

		for i := 0; i < len(c.Args()); i++ {
			tweetID, _ := strconv.ParseInt(c.Args()[i], 10, 64)
			tweet, err := api.Retweet(tweetID, true)
			if err != nil {
				log.Fatal(err)
				break
			}
			dispTweet(tweet, "Retweeted: ", "blue")
		}
	}
}

func doFav(c *cli.Context) {
	api := doOauth()
	defer api.Close()

	if c.Bool("pipe") {
		var stdin string

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin += scanner.Text()
		}
		tweetID, _ := strconv.ParseInt(stdin, 10, 64)
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		_, err := api.Favorite(tweetID)
		if err != nil {
			log.Fatal(err)
		}
		tweet, err := api.GetTweet(tweetID, nil)
		if err != nil {
			log.Fatal(err)
		}
		dispTweet(tweet, "Favorited: ", "blue")
	} else {
		for i := 0; i < len(c.Args()); i++ {
			tweetID, _ := strconv.ParseInt(c.Args()[i], 10, 64)
			_, err := api.Favorite(tweetID)
			if err != nil {
				log.Fatal(err)
				break
			}
			tweet, err := api.GetTweet(tweetID, nil)
			if err != nil {
				log.Fatal(err)
			}
			dispTweet(tweet, "Favorited: ", "blue")
		}
	}
}

func doDel(c *cli.Context) {
	api := doOauth()
	defer api.Close()

	if c.Bool("pipe") {
		var stdin string

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin += scanner.Text()
		}
		tweetID, _ := strconv.ParseInt(stdin, 10, 64)
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		tweet, err := api.DeleteTweet(tweetID, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Deleted: ", tweet.Text)
	} else {

		for i := 0; i < len(c.Args()); i++ {
			tweetID, _ := strconv.ParseInt(c.Args()[i], 10, 64)
			tweet, err := api.DeleteTweet(tweetID, true)
			if err != nil {
				log.Fatal(err)
				break
			}
			fmt.Println("Deleted: ", tweet.Text)
		}
	}
}

func doSearch(c *cli.Context) {
	api := doOauth()
	defer api.Close()
	var query string
	for i := 0; i < len(c.Args()); i++ {
		query += c.Args()[i]
		if i == len(c.Args())-1 {
			continue
		}
		query += " "
	}
	if query == "" {
		log.Fatal("ツイートする文字列を指定してください")
	}
	searchRes, err := api.GetSearch(query, nil)
	if err != nil {
		log.Fatal(err)
	}
	tweets := searchRes.Statuses

	for _, tweet := range tweets {
		dispTweet(tweet, "", "blue")
	}
}

func doTimeline(c *cli.Context) {
	api := doOauth()
	defer api.Close()

	if c.Bool("stream") {
		doStream(api)
	} else {
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

		var timeline []anaconda.Tweet

		if screenName := c.String("user"); len(screenName) > 0 {
			v.Add("screen_name", screenName)
			var err error
			timeline, err = api.GetUserTimeline(v)
			if err != nil {
				panic(err)
			}

		} else {
			timeline, _ = api.GetHomeTimeline(v)
		}

		for _, tweet := range timeline {
			user := tweet.User.Name
			screenName := tweet.User.ScreenName
			user += "(@" + screenName + ")"

			blue := ansi.ColorCode("blue")
			reset := ansi.ColorCode("reset")

			if c.Bool("with-id") {
				fmt.Println(blue, user, ":", reset, tweet.Text, tweet.IdStr)
			} else {
				fmt.Println(blue, user, ":", reset, tweet.Text)
			}
		}
	}

}

func doStream(api *anaconda.TwitterApi) {
	stream := api.UserStream(nil)

	for {
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				user := status.User.Name
				screenName := status.User.ScreenName
				user += "(@" + screenName + ")"

				blue := ansi.ColorCode("blue")
				reset := ansi.ColorCode("reset")

				fmt.Println(blue, user, ":", reset, status.Text)
			default:
			}
		}
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

func dispTweet(tweet anaconda.Tweet, text string, color string) (err error) {

	user := tweet.User.Name
	screenName := tweet.User.ScreenName
	user += "(@" + screenName + ")"

	colorCode := ansi.ColorCode(color)
	reset := ansi.ColorCode("reset")

	fmt.Println(text, colorCode, user, ":", reset, tweet.Text, tweet.Id)

	return
}
