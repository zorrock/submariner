/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	submarineriov1 "github.com/rancher/submariner/pkg/apis/submariner.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEndpoints implements EndpointInterface
type FakeEndpoints struct {
	Fake *FakeSubmarinerV1
	ns   string
}

var endpointsResource = schema.GroupVersionResource{Group: "submariner.io", Version: "v1", Resource: "endpoints"}

var endpointsKind = schema.GroupVersionKind{Group: "submariner.io", Version: "v1", Kind: "Endpoint"}

// Get takes name of the endpoint, and returns the corresponding endpoint object, and an error if there is any.
func (c *FakeEndpoints) Get(name string, options v1.GetOptions) (result *submarineriov1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &submarineriov1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*submarineriov1.Endpoint), err
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *FakeEndpoints) List(opts v1.ListOptions) (result *submarineriov1.EndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &submarineriov1.EndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &submarineriov1.EndpointList{ListMeta: obj.(*submarineriov1.EndpointList).ListMeta}
	for _, item := range obj.(*submarineriov1.EndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *FakeEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))

}

// Create takes the representation of a endpoint and creates it.  Returns the server's representation of the endpoint, and an error, if there is any.
func (c *FakeEndpoints) Create(endpoint *submarineriov1.Endpoint) (result *submarineriov1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoint), &submarineriov1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*submarineriov1.Endpoint), err
}

// Update takes the representation of a endpoint and updates it. Returns the server's representation of the endpoint, and an error, if there is any.
func (c *FakeEndpoints) Update(endpoint *submarineriov1.Endpoint) (result *submarineriov1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoint), &submarineriov1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*submarineriov1.Endpoint), err
}

// Delete takes name of the endpoint and deletes it. Returns an error if one occurs.
func (c *FakeEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(endpointsResource, c.ns, name), &submarineriov1.Endpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &submarineriov1.EndpointList{})
	return err
}

// Patch applies the patch and returns the patched endpoint.
func (c *FakeEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *submarineriov1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, pt, data, subresources...), &submarineriov1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*submarineriov1.Endpoint), err
}