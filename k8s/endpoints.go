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

package k8s

import (
	"log"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/pkg/watch"
)

func createEndpointsWatcher(clientset *kubernetes.Clientset) (watch.Interface, error) {
	watcher, err := clientset.Core().Endpoints("").Watch(api.ListOptions{})
	if err != nil {
		return nil, err
	}
	return watcher, nil
}

func manageEndpointsEvent(eventType watch.EventType, endpoints *v1.Endpoints) {
	switch eventType {
	case watch.Added:
		log.Printf("[INFO] Add endpoint: %s\n", endpoints.Name)
	case watch.Deleted:
		log.Printf("[INFO] Deleted endpoint: %s\n", endpoints.Name)
	case watch.Modified:
		log.Printf("[INFO] Modified endpoint: %s\n", endpoints.Name)

	}
}
