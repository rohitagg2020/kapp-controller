// Code generated by main. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	internalv1alpha1 "carvel.dev/kapp-controller/pkg/client/clientset/versioned/typed/internalpackaging/v1alpha1"
	kappctrlv1alpha1 "carvel.dev/kapp-controller/pkg/client/clientset/versioned/typed/kappctrl/v1alpha1"
	packagingv1alpha1 "carvel.dev/kapp-controller/pkg/client/clientset/versioned/typed/packaging/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface
	KappctrlV1alpha1() kappctrlv1alpha1.KappctrlV1alpha1Interface
	PackagingV1alpha1() packagingv1alpha1.PackagingV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	internalV1alpha1  *internalv1alpha1.InternalV1alpha1Client
	kappctrlV1alpha1  *kappctrlv1alpha1.KappctrlV1alpha1Client
	packagingV1alpha1 *packagingv1alpha1.PackagingV1alpha1Client
}

// InternalV1alpha1 retrieves the InternalV1alpha1Client
func (c *Clientset) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	return c.internalV1alpha1
}

// KappctrlV1alpha1 retrieves the KappctrlV1alpha1Client
func (c *Clientset) KappctrlV1alpha1() kappctrlv1alpha1.KappctrlV1alpha1Interface {
	return c.kappctrlV1alpha1
}

// PackagingV1alpha1 retrieves the PackagingV1alpha1Client
func (c *Clientset) PackagingV1alpha1() packagingv1alpha1.PackagingV1alpha1Interface {
	return c.packagingV1alpha1
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
	cs.internalV1alpha1, err = internalv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kappctrlV1alpha1, err = kappctrlv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.packagingV1alpha1, err = packagingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
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
	cs.internalV1alpha1 = internalv1alpha1.New(c)
	cs.kappctrlV1alpha1 = kappctrlv1alpha1.New(c)
	cs.packagingV1alpha1 = packagingv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
