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

// RobotsGetter has a method to return a RobotInterface.
// A group's client should implement this interface.
type RobotsGetter interface {
	Robots(namespace string) RobotInterface
}

// RobotInterface has methods to work with Robot resources.
type RobotInterface interface {
	Create(ctx context.Context, robot *v1alpha1.Robot, opts v1.CreateOptions) (*v1alpha1.Robot, error)
	Update(ctx context.Context, robot *v1alpha1.Robot, opts v1.UpdateOptions) (*v1alpha1.Robot, error)
	UpdateStatus(ctx context.Context, robot *v1alpha1.Robot, opts v1.UpdateOptions) (*v1alpha1.Robot, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Robot, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RobotList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Robot, err error)
	RobotExpansion
}

// robots implements RobotInterface
type robots struct {
	client rest.Interface
	ns     string
}

// newRobots returns a Robots
func newRobots(c *RoboscaleV1alpha1Client, namespace string) *robots {
	return &robots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the robot, and returns the corresponding robot object, and an error if there is any.
func (c *robots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Robots that match those selectors.
func (c *robots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RobotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RobotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested robots.
func (c *robots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("robots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a robot and creates it.  Returns the server's representation of the robot, and an error, if there is any.
func (c *robots) Create(ctx context.Context, robot *v1alpha1.Robot, opts v1.CreateOptions) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("robots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(robot).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a robot and updates it. Returns the server's representation of the robot, and an error, if there is any.
func (c *robots) Update(ctx context.Context, robot *v1alpha1.Robot, opts v1.UpdateOptions) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("robots").
		Name(robot.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(robot).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *robots) UpdateStatus(ctx context.Context, robot *v1alpha1.Robot, opts v1.UpdateOptions) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("robots").
		Name(robot.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(robot).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the robot and deletes it. Returns an error if one occurs.
func (c *robots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robots").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *robots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robots").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched robot.
func (c *robots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Robot, err error) {
	result = &v1alpha1.Robot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("robots").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
