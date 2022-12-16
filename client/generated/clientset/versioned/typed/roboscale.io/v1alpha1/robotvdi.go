/*
Copyright 2022.

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
	"context"
	"time"

	v1alpha1 "github.com/robolaunch/robot-operator/api/roboscale.io/v1alpha1"
	scheme "github.com/robolaunch/robot-operator/client/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RobotVDIsGetter has a method to return a RobotVDIInterface.
// A group's client should implement this interface.
type RobotVDIsGetter interface {
	RobotVDIs(namespace string) RobotVDIInterface
}

// RobotVDIInterface has methods to work with RobotVDI resources.
type RobotVDIInterface interface {
	Create(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.CreateOptions) (*v1alpha1.RobotVDI, error)
	Update(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.UpdateOptions) (*v1alpha1.RobotVDI, error)
	UpdateStatus(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.UpdateOptions) (*v1alpha1.RobotVDI, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RobotVDI, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RobotVDIList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RobotVDI, err error)
	RobotVDIExpansion
}

// robotVDIs implements RobotVDIInterface
type robotVDIs struct {
	client rest.Interface
	ns     string
}

// newRobotVDIs returns a RobotVDIs
func newRobotVDIs(c *RoboscaleV1alpha1Client, namespace string) *robotVDIs {
	return &robotVDIs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the robotVDI, and returns the corresponding robotVDI object, and an error if there is any.
func (c *robotVDIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RobotVDI, err error) {
	result = &v1alpha1.RobotVDI{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robotvdis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RobotVDIs that match those selectors.
func (c *robotVDIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RobotVDIList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RobotVDIList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robotvdis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested robotVDIs.
func (c *robotVDIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("robotvdis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a robotVDI and creates it.  Returns the server's representation of the robotVDI, and an error, if there is any.
func (c *robotVDIs) Create(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.CreateOptions) (result *v1alpha1.RobotVDI, err error) {
	result = &v1alpha1.RobotVDI{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("robotvdis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(robotVDI).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a robotVDI and updates it. Returns the server's representation of the robotVDI, and an error, if there is any.
func (c *robotVDIs) Update(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.UpdateOptions) (result *v1alpha1.RobotVDI, err error) {
	result = &v1alpha1.RobotVDI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("robotvdis").
		Name(robotVDI.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(robotVDI).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *robotVDIs) UpdateStatus(ctx context.Context, robotVDI *v1alpha1.RobotVDI, opts v1.UpdateOptions) (result *v1alpha1.RobotVDI, err error) {
	result = &v1alpha1.RobotVDI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("robotvdis").
		Name(robotVDI.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(robotVDI).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the robotVDI and deletes it. Returns an error if one occurs.
func (c *robotVDIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robotvdis").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *robotVDIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robotvdis").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched robotVDI.
func (c *robotVDIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RobotVDI, err error) {
	result = &v1alpha1.RobotVDI{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("robotvdis").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}