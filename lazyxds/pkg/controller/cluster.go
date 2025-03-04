// Copyright Aeraki Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

// Cluster ...
type Cluster struct {
	name     string
	Client   *kubernetes.Clientset
	Informer informers.SharedInformerFactory
	stopCh   chan struct{}
}

// NewCluster ...
func NewCluster(name string, client *kubernetes.Clientset, stop chan struct{}) *Cluster {
	return &Cluster{
		name:     name,
		Client:   client,
		Informer: informers.NewSharedInformerFactory(client, 0),
		stopCh:   stop,
	}
}
