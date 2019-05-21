/*
Copyright 2019 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/openfaas-incubator/openfaas-operator/pkg/apis/openfaas/v1alpha2"
	"github.com/openfaas-incubator/openfaas-operator/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type OpenfaasV1alpha2Interface interface {
	RESTClient() rest.Interface
	FunctionsGetter
}

// OpenfaasV1alpha2Client is used to interact with features provided by the openfaas.com group.
type OpenfaasV1alpha2Client struct {
	restClient rest.Interface
}

func (c *OpenfaasV1alpha2Client) Functions(namespace string) FunctionInterface {
	return newFunctions(c, namespace)
}

// NewForConfig creates a new OpenfaasV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*OpenfaasV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OpenfaasV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new OpenfaasV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OpenfaasV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OpenfaasV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *OpenfaasV1alpha2Client {
	return &OpenfaasV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *OpenfaasV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
