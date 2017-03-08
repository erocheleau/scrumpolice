package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/k0kubun/pp"
	"github.com/nlopes/slack"
	"github.com/scrumpolice/scrumpolice/bot"
)

const header = "                                           _ _\n" +
	" ___  ___ _ __ _   _ _ __ ___  _ __   ___ | (_) ___ ___\n" +
	"/ __|/ __| '__| | | | '_ ` _ \\| '_ \\ / _ \\| | |/ __/ _ \\\n" +
	"\\__ \\ (__| |  | |_| | | | | | | |_) | (_) | | | (_|  __/\n" +
	"|___/\\___|_|   \\__,_|_| |_| |_| .__/ \\___/|_|_|\\___\\___|\n" +
	"                              |_|"

const Version = "0.0.0"

func main() {
	fmt.Println(header)
	fmt.Println("Version", Version)
	fmt.Println("")

	slackBotToken := os.Getenv("SCRUMPOLICE_SLACK_TOKEN")

	if slackBotToken == "" {
		log.Fatalln("slack bot token must be set in SCRUMPOLICE_SLACK_TOKEN")
	}

	configFile := "config.json"
	flag.StringVar(&configFile, "config", configFile, "The configuration file")
	flag.Parse()

	// Injection
	logger := log.New(os.Stderr, "", log.Lshortfile)

	configurationProvider := bot.NewConfigWatcher(configFile)

	configurationProvider.OnChange(func() {
		pp.Println("Configuration File Changed heres the new teams")
		pp.Println(configurationProvider.Config().ToTeams())
	})

	slackAPIClient := slack.New(slackBotToken)

	// Create and run bot
	b := bot.New(slackAPIClient, logger)
	b.Run()
}
