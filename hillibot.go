package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {

	// Load Env variables from .dot file
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")
	// Create a new client to slack by giving token
	// Set debug to true while developing
	// Also add a ApplicationToken option to the client
	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))
	// go-slack comes with a SocketMode package that we need to use that accepts a Slack client and outputs a Socket mode client instead
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		// Option to set a custom logger
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	socketClient.Run()
}