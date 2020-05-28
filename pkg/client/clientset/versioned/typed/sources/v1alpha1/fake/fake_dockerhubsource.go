/*
Copyright 2020 The Knative Authors

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
	v1alpha1 "github.com/tom24d/eventing-dockerhub/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDockerHubSources implements DockerHubSourceInterface
type FakeDockerHubSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var dockerhubsourcesResource = schema.GroupVersionResource{Group: "sources.knative.dev", Version: "v1alpha1", Resource: "dockerhubsources"}

var dockerhubsourcesKind = schema.GroupVersionKind{Group: "sources.knative.dev", Version: "v1alpha1", Kind: "DockerHubSource"}

// Get takes name of the dockerHubSource, and returns the corresponding dockerHubSource object, and an error if there is any.
func (c *FakeDockerHubSources) Get(name string, options v1.GetOptions) (result *v1alpha1.DockerHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dockerhubsourcesResource, c.ns, name), &v1alpha1.DockerHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerHubSource), err
}

// List takes label and field selectors, and returns the list of DockerHubSources that match those selectors.
func (c *FakeDockerHubSources) List(opts v1.ListOptions) (result *v1alpha1.DockerHubSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dockerhubsourcesResource, dockerhubsourcesKind, c.ns, opts), &v1alpha1.DockerHubSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DockerHubSourceList{ListMeta: obj.(*v1alpha1.DockerHubSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DockerHubSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dockerHubSources.
func (c *FakeDockerHubSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dockerhubsourcesResource, c.ns, opts))

}

// Create takes the representation of a dockerHubSource and creates it.  Returns the server's representation of the dockerHubSource, and an error, if there is any.
func (c *FakeDockerHubSources) Create(dockerHubSource *v1alpha1.DockerHubSource) (result *v1alpha1.DockerHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dockerhubsourcesResource, c.ns, dockerHubSource), &v1alpha1.DockerHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerHubSource), err
}

// Update takes the representation of a dockerHubSource and updates it. Returns the server's representation of the dockerHubSource, and an error, if there is any.
func (c *FakeDockerHubSources) Update(dockerHubSource *v1alpha1.DockerHubSource) (result *v1alpha1.DockerHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dockerhubsourcesResource, c.ns, dockerHubSource), &v1alpha1.DockerHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerHubSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDockerHubSources) UpdateStatus(dockerHubSource *v1alpha1.DockerHubSource) (*v1alpha1.DockerHubSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dockerhubsourcesResource, "status", c.ns, dockerHubSource), &v1alpha1.DockerHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerHubSource), err
}

// Delete takes name of the dockerHubSource and deletes it. Returns an error if one occurs.
func (c *FakeDockerHubSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dockerhubsourcesResource, c.ns, name), &v1alpha1.DockerHubSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDockerHubSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dockerhubsourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DockerHubSourceList{})
	return err
}

// Patch applies the patch and returns the patched dockerHubSource.
func (c *FakeDockerHubSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DockerHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dockerhubsourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DockerHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerHubSource), err
}
