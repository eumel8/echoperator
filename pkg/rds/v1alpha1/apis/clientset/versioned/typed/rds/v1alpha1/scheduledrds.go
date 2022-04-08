package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/eumel8/echoperator/pkg/rds/v1alpha1"
	scheme "github.com/eumel8/echoperator/pkg/rds/v1alpha1/apis/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScheduledRdssGetter has a method to return a ScheduledRdsInterface.
// A group's client should implement this interface.
type ScheduledRdssGetter interface {
	ScheduledRdss(namespace string) ScheduledRdsInterface
}

// ScheduledRdsInterface has methods to work with ScheduledRds resources.
type ScheduledRdsInterface interface {
	Create(ctx context.Context, scheduledRds *v1alpha1.ScheduledRds, opts v1.CreateOptions) (*v1alpha1.ScheduledRds, error)
	Update(ctx context.Context, scheduledRds *v1alpha1.ScheduledRds, opts v1.UpdateOptions) (*v1alpha1.ScheduledRds, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ScheduledRds, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ScheduledRdsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScheduledRds, err error)
	ScheduledRdsExpansion
}

// scheduledRdss implements ScheduledRdsInterface
type scheduledRdss struct {
	client rest.Interface
	ns     string
}

// newScheduledRdss returns a ScheduledRdss
func newScheduledRdss(c *McspsV1alpha1Client, namespace string) *scheduledRdss {
	return &scheduledRdss{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scheduledRds, and returns the corresponding scheduledRds object, and an error if there is any.
func (c *scheduledRdss) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScheduledRds, err error) {
	result = &v1alpha1.ScheduledRds{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledechos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScheduledRdss that match those selectors.
func (c *scheduledRdss) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScheduledRdsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ScheduledRdsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledechos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scheduledRdss.
func (c *scheduledRdss) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scheduledechos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a scheduledRds and creates it.  Returns the server's representation of the scheduledRds, and an error, if there is any.
func (c *scheduledRdss) Create(ctx context.Context, scheduledRds *v1alpha1.ScheduledRds, opts v1.CreateOptions) (result *v1alpha1.ScheduledRds, err error) {
	result = &v1alpha1.ScheduledRds{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scheduledechos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduledRds).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a scheduledRds and updates it. Returns the server's representation of the scheduledRds, and an error, if there is any.
func (c *scheduledRdss) Update(ctx context.Context, scheduledRds *v1alpha1.ScheduledRds, opts v1.UpdateOptions) (result *v1alpha1.ScheduledRds, err error) {
	result = &v1alpha1.ScheduledRds{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledechos").
		Name(scheduledRds.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduledRds).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the scheduledRds and deletes it. Returns an error if one occurs.
func (c *scheduledRdss) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledechos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scheduledRdss) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledechos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched scheduledRds.
func (c *scheduledRdss) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScheduledRds, err error) {
	result = &v1alpha1.ScheduledRds{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scheduledechos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
