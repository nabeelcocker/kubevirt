/*
Copyright 2021 The KubeVirt Authors.

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

package v1

import (
	"time"

	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	scheme "kubevirt.io/client-go/generated/prometheus-operator/clientset/versioned/scheme"
)

// AlertmanagersGetter has a method to return a AlertmanagerInterface.
// A group's client should implement this interface.
type AlertmanagersGetter interface {
	Alertmanagers(namespace string) AlertmanagerInterface
}

// AlertmanagerInterface has methods to work with Alertmanager resources.
type AlertmanagerInterface interface {
	Create(*v1.Alertmanager) (*v1.Alertmanager, error)
	Update(*v1.Alertmanager) (*v1.Alertmanager, error)
	UpdateStatus(*v1.Alertmanager) (*v1.Alertmanager, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Alertmanager, error)
	List(opts metav1.ListOptions) (*v1.AlertmanagerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Alertmanager, err error)
	AlertmanagerExpansion
}

// alertmanagers implements AlertmanagerInterface
type alertmanagers struct {
	client rest.Interface
	ns     string
}

// newAlertmanagers returns a Alertmanagers
func newAlertmanagers(c *MonitoringV1Client, namespace string) *alertmanagers {
	return &alertmanagers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *alertmanagers) Get(name string, options metav1.GetOptions) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *alertmanagers) List(opts metav1.ListOptions) (result *v1.AlertmanagerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AlertmanagerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *alertmanagers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *alertmanagers) Create(alertmanager *v1.Alertmanager) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alertmanagers").
		Body(alertmanager).
		Do().
		Into(result)
	return
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *alertmanagers) Update(alertmanager *v1.Alertmanager) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(alertmanager.Name).
		Body(alertmanager).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *alertmanagers) UpdateStatus(alertmanager *v1.Alertmanager) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(alertmanager.Name).
		SubResource("status").
		Body(alertmanager).
		Do().
		Into(result)
	return
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *alertmanagers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *alertmanagers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertmanagers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched alertmanager.
func (c *alertmanagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alertmanagers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
