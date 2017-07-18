// Copyright (C) 2016-2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package telegram

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/zeiot/jarvis-bot/k8s"
)

type Client struct {
	Bot *tgbotapi.BotAPI
}

func NewClient(token string, debug bool) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = debug
	log.Printf("[INFO] %s is ready.\n", bot.Self.UserName)
	return &Client{
		Bot: bot,
	}, nil
}

func (client *Client) Run(k8sclient *k8s.Client) error {
	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60
	updates, err := client.Bot.GetUpdatesChan(updConfig)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[DEBUG] From: %s Message: %s", update.Message.From.UserName, update.Message.Text)
		if len(update.Message.Text) > 1 && string(update.Message.Text[0]) == "/" {
			route(client.Bot, update, k8sclient)
		}
	}
	return nil
}
