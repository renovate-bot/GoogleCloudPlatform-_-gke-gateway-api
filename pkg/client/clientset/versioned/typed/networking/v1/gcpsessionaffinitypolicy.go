/*
* Copyright 2024 Google LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     https://www.apache.org/licenses/LICENSE-2.0
*
*     Unless required by applicable law or agreed to in writing, software
*     distributed under the License is distributed on an "AS IS" BASIS,
*     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*     See the License for the specific language governing permissions and
*     limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/GoogleCloudPlatform/gke-gateway-api/apis/networking/v1"
	scheme "github.com/GoogleCloudPlatform/gke-gateway-api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GCPSessionAffinityPoliciesGetter has a method to return a GCPSessionAffinityPolicyInterface.
// A group's client should implement this interface.
type GCPSessionAffinityPoliciesGetter interface {
	GCPSessionAffinityPolicies(namespace string) GCPSessionAffinityPolicyInterface
}

// GCPSessionAffinityPolicyInterface has methods to work with GCPSessionAffinityPolicy resources.
type GCPSessionAffinityPolicyInterface interface {
	Create(ctx context.Context, gCPSessionAffinityPolicy *v1.GCPSessionAffinityPolicy, opts metav1.CreateOptions) (*v1.GCPSessionAffinityPolicy, error)
	Update(ctx context.Context, gCPSessionAffinityPolicy *v1.GCPSessionAffinityPolicy, opts metav1.UpdateOptions) (*v1.GCPSessionAffinityPolicy, error)
	UpdateStatus(ctx context.Context, gCPSessionAffinityPolicy *v1.GCPSessionAffinityPolicy, opts metav1.UpdateOptions) (*v1.GCPSessionAffinityPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GCPSessionAffinityPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GCPSessionAffinityPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GCPSessionAffinityPolicy, err error)
	GCPSessionAffinityPolicyExpansion
}

// gCPSessionAffinityPolicies implements GCPSessionAffinityPolicyInterface
type gCPSessionAffinityPolicies struct {
	client rest.Interface
	ns     string
}

// newGCPSessionAffinityPolicies returns a GCPSessionAffinityPolicies
func newGCPSessionAffinityPolicies(c *NetworkingV1Client, namespace string) *gCPSessionAffinityPolicies {
	return &gCPSessionAffinityPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gCPSessionAffinityPolicy, and returns the corresponding gCPSessionAffinityPolicy object, and an error if there is any.
func (c *gCPSessionAffinityPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GCPSessionAffinityPolicy, err error) {
	result = &v1.GCPSessionAffinityPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GCPSessionAffinityPolicies that match those selectors.
func (c *gCPSessionAffinityPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GCPSessionAffinityPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GCPSessionAffinityPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gCPSessionAffinityPolicies.
func (c *gCPSessionAffinityPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gCPSessionAffinityPolicy and creates it.  Returns the server's representation of the gCPSessionAffinityPolicy, and an error, if there is any.
func (c *gCPSessionAffinityPolicies) Create(ctx context.Context, gCPSessionAffinityPolicy *v1.GCPSessionAffinityPolicy, opts metav1.CreateOptions) (result *v1.GCPSessionAffinityPolicy, err error) {
	result = &v1.GCPSessionAffinityPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gCPSessionAffinityPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gCPSessionAffinityPolicy and updates it. Returns the server's representation of the gCPSessionAffinityPolicy, and an error, if there is any.
func (c *gCPSessionAffinityPolicies) Update(ctx context.Context, gCPSessionAffinityPolicy *v1.GCPSessionAffinityPolicy, opts metav1.UpdateOptions) (result *v1.GCPSessionAffinityPolicy, err error) {
	result = &v1.GCPSessionAffinityPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		Name(gCPSessionAffinityPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gCPSessionAffinityPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gCPSessionAffinityPolicies) UpdateStatus(ctx context.Context, gCPSessionAffinityPolicy *v1.GCPSessionAffinityPolicy, opts metav1.UpdateOptions) (result *v1.GCPSessionAffinityPolicy, err error) {
	result = &v1.GCPSessionAffinityPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		Name(gCPSessionAffinityPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gCPSessionAffinityPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gCPSessionAffinityPolicy and deletes it. Returns an error if one occurs.
func (c *gCPSessionAffinityPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gCPSessionAffinityPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gCPSessionAffinityPolicy.
func (c *gCPSessionAffinityPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GCPSessionAffinityPolicy, err error) {
	result = &v1.GCPSessionAffinityPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gcpsessionaffinitypolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}