package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/anaconda"
	"github.com/mitchellh/go-homedir"
	"github.com/mrjones/oauth"
	"github.com/skratchdot/open-golang/open"
)

type AccessToken struct {
	Token  string `toml:"token"`
	Secret string `toml:"secret"`
}

func isFileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func getAccessToken(consumerKey, consumerSecret string) (a AccessToken, err error) {
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
	a.Token = authToken.Token
	a.Secret = authToken.Secret

	err = saveAccessToken(a)

	return
}

func saveAccessToken(a AccessToken) (err error) {
	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	err = encoder.Encode(a)
	if err != nil {
		return err
	}
	home, _ := homedir.Dir()
	saveFilePath := home + "/.tw.toml"
	err = ioutil.WriteFile(saveFilePath, buffer.Bytes(), 0664)
	if err != nil {
		return err
	}

	return err
}

func loadAccessToken() (a AccessToken, err error) {
	home, _ := homedir.Dir()
	filePath := home + "/.tw.toml"
	_, err = toml.DecodeFile(filePath, &a)

	return

}

func doOauth() (twitterApi *anaconda.TwitterApi) {
	consumerKey := "oaFjoL9zSCPdNKIfCfae4iRGf"
	consumerSecret := "R016AC6zTohQtWU0ynUbyjdkNEcJHzTGNclRXmFiDzBtSgZEyj"
	var accessToken AccessToken
	home, _ := homedir.Dir()
	filePath := home + "/.tw.toml"

	if isFileExists(filePath) {
		accessToken, _ = loadAccessToken()
	} else {
		accessToken, _ = getAccessToken(consumerKey, consumerSecret)
	}

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	twitterApi = anaconda.NewTwitterApi(accessToken.Token, accessToken.Secret)

	return
}
