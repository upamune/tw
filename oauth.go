package main

import (
	"fmt"
	"log"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mrjones/oauth"
	"github.com/skratchdot/open-golang/open"
)

func doOauth() (twitterApi *anaconda.TwitterApi) {
	var consumerKey string = "oaFjoL9zSCPdNKIfCfae4iRGf"

	var consumerSecret string = "R016AC6zTohQtWU0ynUbyjdkNEcJHzTGNclRXmFiDzBtSgZEyj"

	c := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	requestToken, url, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("(1) Go to: " + url)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	open.Run(url)
	fmt.Print("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	authToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}
	accessToken := authToken.Token
	accessTokenSecret := authToken.Secret

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	twitterApi = anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	return
}
