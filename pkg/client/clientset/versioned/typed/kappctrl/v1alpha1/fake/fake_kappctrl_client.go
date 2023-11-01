// Code generated by main. DO NOT EDIT.

package fake

import (
	v1alpha1 "carvel.dev/kapp-controller/pkg/client/clientset/versioned/typed/kappctrl/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKappctrlV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKappctrlV1alpha1) Apps(namespace string) v1alpha1.AppInterface {
	return &FakeApps{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKappctrlV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
