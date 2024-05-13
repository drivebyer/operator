// This file is part of MinIO Operator
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	jobv1alpha1 "github.com/minio/operator/pkg/client/clientset/versioned/typed/job.min.io/v1alpha1"
	miniov2 "github.com/minio/operator/pkg/client/clientset/versioned/typed/minio.min.io/v2"
	stsv1alpha1 "github.com/minio/operator/pkg/client/clientset/versioned/typed/sts.min.io/v1alpha1"
	stsv1beta1 "github.com/minio/operator/pkg/client/clientset/versioned/typed/sts.min.io/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	JobV1alpha1() jobv1alpha1.JobV1alpha1Interface
	MinioV2() miniov2.MinioV2Interface
	StsV1alpha1() stsv1alpha1.StsV1alpha1Interface
	StsV1beta1() stsv1beta1.StsV1beta1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	jobV1alpha1 *jobv1alpha1.JobV1alpha1Client
	minioV2     *miniov2.MinioV2Client
	stsV1alpha1 *stsv1alpha1.StsV1alpha1Client
	stsV1beta1  *stsv1beta1.StsV1beta1Client
}

// JobV1alpha1 retrieves the JobV1alpha1Client
func (c *Clientset) JobV1alpha1() jobv1alpha1.JobV1alpha1Interface {
	return c.jobV1alpha1
}

// MinioV2 retrieves the MinioV2Client
func (c *Clientset) MinioV2() miniov2.MinioV2Interface {
	return c.minioV2
}

// StsV1alpha1 retrieves the StsV1alpha1Client
func (c *Clientset) StsV1alpha1() stsv1alpha1.StsV1alpha1Interface {
	return c.stsV1alpha1
}

// StsV1beta1 retrieves the StsV1beta1Client
func (c *Clientset) StsV1beta1() stsv1beta1.StsV1beta1Interface {
	return c.stsV1beta1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.jobV1alpha1, err = jobv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.minioV2, err = miniov2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.stsV1alpha1, err = stsv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.stsV1beta1, err = stsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.jobV1alpha1 = jobv1alpha1.New(c)
	cs.minioV2 = miniov2.New(c)
	cs.stsV1alpha1 = stsv1alpha1.New(c)
	cs.stsV1beta1 = stsv1beta1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
