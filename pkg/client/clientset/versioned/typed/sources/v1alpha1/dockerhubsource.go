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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/tom24d/eventing-dockerhub/pkg/apis/sources/v1alpha1"
	scheme "github.com/tom24d/eventing-dockerhub/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DockerHubSourcesGetter has a method to return a DockerHubSourceInterface.
// A group's client should implement this interface.
type DockerHubSourcesGetter interface {
	DockerHubSources(namespace string) DockerHubSourceInterface
}

// DockerHubSourceInterface has methods to work with DockerHubSource resources.
type DockerHubSourceInterface interface {
	Create(*v1alpha1.DockerHubSource) (*v1alpha1.DockerHubSource, error)
	Update(*v1alpha1.DockerHubSource) (*v1alpha1.DockerHubSource, error)
	UpdateStatus(*v1alpha1.DockerHubSource) (*v1alpha1.DockerHubSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DockerHubSource, error)
	List(opts v1.ListOptions) (*v1alpha1.DockerHubSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DockerHubSource, err error)
	DockerHubSourceExpansion
}

// dockerHubSources implements DockerHubSourceInterface
type dockerHubSources struct {
	client rest.Interface
	ns     string
}

// newDockerHubSources returns a DockerHubSources
func newDockerHubSources(c *SourcesV1alpha1Client, namespace string) *dockerHubSources {
	return &dockerHubSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dockerHubSource, and returns the corresponding dockerHubSource object, and an error if there is any.
func (c *dockerHubSources) Get(name string, options v1.GetOptions) (result *v1alpha1.DockerHubSource, err error) {
	result = &v1alpha1.DockerHubSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dockerhubsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DockerHubSources that match those selectors.
func (c *dockerHubSources) List(opts v1.ListOptions) (result *v1alpha1.DockerHubSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DockerHubSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dockerhubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dockerHubSources.
func (c *dockerHubSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dockerhubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dockerHubSource and creates it.  Returns the server's representation of the dockerHubSource, and an error, if there is any.
func (c *dockerHubSources) Create(dockerHubSource *v1alpha1.DockerHubSource) (result *v1alpha1.DockerHubSource, err error) {
	result = &v1alpha1.DockerHubSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dockerhubsources").
		Body(dockerHubSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dockerHubSource and updates it. Returns the server's representation of the dockerHubSource, and an error, if there is any.
func (c *dockerHubSources) Update(dockerHubSource *v1alpha1.DockerHubSource) (result *v1alpha1.DockerHubSource, err error) {
	result = &v1alpha1.DockerHubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dockerhubsources").
		Name(dockerHubSource.Name).
		Body(dockerHubSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dockerHubSources) UpdateStatus(dockerHubSource *v1alpha1.DockerHubSource) (result *v1alpha1.DockerHubSource, err error) {
	result = &v1alpha1.DockerHubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dockerhubsources").
		Name(dockerHubSource.Name).
		SubResource("status").
		Body(dockerHubSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the dockerHubSource and deletes it. Returns an error if one occurs.
func (c *dockerHubSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dockerhubsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dockerHubSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dockerhubsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dockerHubSource.
func (c *dockerHubSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DockerHubSource, err error) {
	result = &v1alpha1.DockerHubSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dockerhubsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
