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

	"github.com/google/go-github/github"
)

func getReleases() string {
	var buffer bytes.Buffer
	buffer.WriteString("Releases:\n")

	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, _ := client.Repositories.ListByOrg("Zeiot", opt)
	for _, repo := range repos {
		fmt.Printf("[DEBUG] Repo %s: %s\n", *repo.Owner.Login, *repo.ReleasesURL)
		tags, _, _ := client.Repositories.ListTags("zeiot", *repo.Name, nil)
		if len(tags) > 0 {
			buffer.WriteString(fmt.Sprintf("- %s : %s\n", *repo.Name, *tags[0].Name))
		}
		// for _, tag := range tags {
		// 	fmt.Printf("[DEBUG] Tag %s\n", *tag.Name)
		// }
	}
	text := buffer.String()
	log.Printf("[INFO] Github releases: %s", text)
	return text
}

func getRepositories(client *github.Client) []github.Repository {
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, _ := client.Repositories.ListByOrg("Zeiot", opt)
	return repos
}

// func getGithubRelease(client *github.Client, project string) (*string, error) {
// 	log.Printf("[DEBUG] Get last release from Github: %s", project)
// 	latestRelease, _, err := client.Repositories.GetLatestRelease("zeiot", project)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return latestRelease.TagName, nil
// }
