package fake

import (
	"context"

	v1alpha1 "github.com/eumel8/echoperator/pkg/rds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScheduledRdss implements ScheduledRdsInterface
type FakeScheduledRdss struct {
	Fake *FakeMcspsV1alpha1
	ns   string
}

var scheduledrdssResource = schema.GroupVersionResource{Group: "mcsps.de", Version: "v1alpha1", Resource: "scheduledrdss"}

var scheduledrdssKind = schema.GroupVersionKind{Group: "mcsps.de", Version: "v1alpha1", Kind: "ScheduledRds"}

// Get takes name of the scheduledRds, and returns the corresponding scheduledRds object, and an error if there is any.
func (c *FakeScheduledRdss) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScheduledRds, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scheduledrdssResource, c.ns, name), &v1alpha1.ScheduledRds{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledRds), err
}

// List takes label and field selectors, and returns the list of ScheduledRdss that match those selectors.
func (c *FakeScheduledRdss) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScheduledRdsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scheduledrdssResource, scheduledrdssKind, c.ns, opts), &v1alpha1.ScheduledRdsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScheduledRdsList{ListMeta: obj.(*v1alpha1.ScheduledRdsList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScheduledRdsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scheduledRdss.
func (c *FakeScheduledRdss) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scheduledrdssResource, c.ns, opts))

}

// Create takes the representation of a scheduledRds and creates it.  Returns the server's representation of the scheduledRds, and an error, if there is any.
func (c *FakeScheduledRdss) Create(ctx context.Context, scheduledRds *v1alpha1.ScheduledRds, opts v1.CreateOptions) (result *v1alpha1.ScheduledRds, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scheduledrdssResource, c.ns, scheduledRds), &v1alpha1.ScheduledRds{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledRds), err
}

// Update takes the representation of a scheduledRds and updates it. Returns the server's representation of the scheduledRds, and an error, if there is any.
func (c *FakeScheduledRdss) Update(ctx context.Context, scheduledRds *v1alpha1.ScheduledRds, opts v1.UpdateOptions) (result *v1alpha1.ScheduledRds, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scheduledrdssResource, c.ns, scheduledRds), &v1alpha1.ScheduledRds{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledRds), err
}

// Delete takes name of the scheduledRds and deletes it. Returns an error if one occurs.
func (c *FakeScheduledRdss) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scheduledrdssResource, c.ns, name), &v1alpha1.ScheduledRds{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScheduledRdss) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scheduledrdssResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScheduledRdsList{})
	return err
}

// Patch applies the patch and returns the patched scheduledRds.
func (c *FakeScheduledRdss) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScheduledRds, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scheduledrdssResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScheduledRds{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledRds), err
}
