// Copyright 2023 Red Hat, Inc. and/or its affiliates
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

package kubernetes

import (
	"net/url"

	v1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
)

// TODO: retrieve the cluster domain from the /etc/resolve inside the pod or from the Platform CRD - will be addressed by KOGITO-9198
var defaultClusterDomain = "svc.cluster.local"

// retrieveServiceHost function that based on the service name, namespace and eventually the nodeport, will provide the service URI
func retrieveServiceHost(service *v1.Service) string {
	namespace := service.Namespace
	if len(namespace) == 0 {
		namespace = "default"
	}
	// TODO: Retrieve the cluster domain or use the default one
	return service.Name + "." + namespace + "." + defaultClusterDomain
}

// RetrieveServiceURL function that based on the service name, namespace and eventually the nodeport, will provide the service URI
func RetrieveServiceURL(service *v1.Service) (*apis.URL, error) {
	url := url.URL{
		Scheme: "http",
		Host:   retrieveServiceHost(service),
		Path:   service.Name}
	return apis.ParseURL(url.String())
}
