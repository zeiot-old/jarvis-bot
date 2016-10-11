// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/zeiot/jarvis-bot/version"
)

func main() {
	var (
		showVersion = flag.Bool("version", false, "Print version information.")
		debug       = flag.Bool("debug", false, "Debug mode for Telegram.")
		token       = flag.String("token", "", "Bot token.")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("Jarvis Bot. v%s\n", version.Version)
		os.Exit(0)
	}

	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	if *debug {
		bot.Debug = true
	}
	log.Printf("[INFO] %ss is ready.\n", bot.Self.UserName)
	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[DEBUG] From: %s Message: %s", update.Message.From.UserName, update.Message.Text)
		if len(update.Message.Text) > 1 && string(update.Message.Text[0]) == "/" {
			route(bot, update)
		}
	}
}
