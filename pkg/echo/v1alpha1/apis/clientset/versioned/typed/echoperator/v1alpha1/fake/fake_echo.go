/*
MIT License

Copyright (c) 2021 Martín Montes

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/mmontes11/echoperator/pkg/echo/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEchos implements EchoInterface
type FakeEchos struct {
	Fake *FakeMmontesV1alpha1
	ns   string
}

var echosResource = schema.GroupVersionResource{Group: "mmontes.io", Version: "v1alpha1", Resource: "echos"}

var echosKind = schema.GroupVersionKind{Group: "mmontes.io", Version: "v1alpha1", Kind: "Echo"}

// Get takes name of the echo, and returns the corresponding echo object, and an error if there is any.
func (c *FakeEchos) Get(
	ctx context.Context,
	name string,
	options v1.GetOptions,
) (result *v1alpha1.Echo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(echosResource, c.ns, name), &v1alpha1.Echo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Echo), err
}

// List takes label and field selectors, and returns the list of Echos that match those selectors.
func (c *FakeEchos) List(
	ctx context.Context,
	opts v1.ListOptions,
) (result *v1alpha1.EchoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(echosResource, echosKind, c.ns, opts), &v1alpha1.EchoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EchoList{ListMeta: obj.(*v1alpha1.EchoList).ListMeta}
	for _, item := range obj.(*v1alpha1.EchoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested echos.
func (c *FakeEchos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(echosResource, c.ns, opts))

}

// Create takes the representation of a echo and creates it.  Returns the server's representation of the echo, and an error, if there is any.
func (c *FakeEchos) Create(
	ctx context.Context,
	echo *v1alpha1.Echo,
	opts v1.CreateOptions,
) (result *v1alpha1.Echo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(echosResource, c.ns, echo), &v1alpha1.Echo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Echo), err
}

// Update takes the representation of a echo and updates it. Returns the server's representation of the echo, and an error, if there is any.
func (c *FakeEchos) Update(
	ctx context.Context,
	echo *v1alpha1.Echo,
	opts v1.UpdateOptions,
) (result *v1alpha1.Echo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(echosResource, c.ns, echo), &v1alpha1.Echo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Echo), err
}

// Delete takes name of the echo and deletes it. Returns an error if one occurs.
func (c *FakeEchos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(echosResource, c.ns, name), &v1alpha1.Echo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEchos) DeleteCollection(
	ctx context.Context,
	opts v1.DeleteOptions,
	listOpts v1.ListOptions,
) error {
	action := testing.NewDeleteCollectionAction(echosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EchoList{})
	return err
}

// Patch applies the patch and returns the patched echo.
func (c *FakeEchos) Patch(
	ctx context.Context,
	name string,
	pt types.PatchType,
	data []byte,
	opts v1.PatchOptions,
	subresources ...string,
) (result *v1alpha1.Echo, err error) {
	obj, err := c.Fake.
		Invokes(
			testing.NewPatchSubresourceAction(echosResource, c.ns, name, pt, data, subresources...),
			&v1alpha1.Echo{},
		)

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Echo), err
}