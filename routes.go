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
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/zeiot/jarvis-bot/k8s"
	"github.com/zeiot/jarvis-bot/version"
)

var (
	botName = "Jarvis"

	messages = map[string]string{
		"Help": `You can use the following commands:
        /releases : Display last release of components
	/help     : Display help message
	/version  : Show bot version
	`,
	}
)

func route(bot *tgbotapi.BotAPI, update tgbotapi.Update, k8sclient *k8s.Client) {
	tokens := strings.Fields(update.Message.Text)
	command := tokens[0]
	switch command {
	case "/help":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages["Help"])
		bot.Send(msg)
	case "/help@" + botName:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages["Help"])
		bot.Send(msg)
	case "/releases":
		text := getReleases()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
	case "/k8s-services":
		text := getKubernetesServices(k8sclient)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
	case "/version":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("v%s", version.Version))
		bot.Send(msg)
	}
}

func getKubernetesServices(k8sclient *k8s.Client) string {
	var buffer bytes.Buffer
	buffer.WriteString("Kubernetes Services:\n")
	services, err := k8sclient.GetServices()
	if err != nil {
		log.Printf("[ERROR] Kubernetes services: %s", err.Error())
		buffer.WriteString("Can't retrieve Kubernetes services")
		return buffer.String()
	}
	for _, service := range services.Items {
		buffer.WriteString(fmt.Sprintf("- %s\n", service.Name))
	}
	text := buffer.String()
	log.Printf("[INFO] Kubernetes services: %s", text)
	return text
}
