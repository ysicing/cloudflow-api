// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/ysicing/cloudflow-api/apps/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWebs implements WebInterface
type FakeWebs struct {
	Fake *FakeAppsV1beta1
	ns   string
}

var websResource = schema.GroupVersionResource{Group: "apps.ysicing.me", Version: "v1beta1", Resource: "webs"}

var websKind = schema.GroupVersionKind{Group: "apps.ysicing.me", Version: "v1beta1", Kind: "Web"}

// Get takes name of the web, and returns the corresponding web object, and an error if there is any.
func (c *FakeWebs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Web, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(websResource, c.ns, name), &v1beta1.Web{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Web), err
}

// List takes label and field selectors, and returns the list of Webs that match those selectors.
func (c *FakeWebs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.WebList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(websResource, websKind, c.ns, opts), &v1beta1.WebList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.WebList{ListMeta: obj.(*v1beta1.WebList).ListMeta}
	for _, item := range obj.(*v1beta1.WebList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested webs.
func (c *FakeWebs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(websResource, c.ns, opts))

}

// Create takes the representation of a web and creates it.  Returns the server's representation of the web, and an error, if there is any.
func (c *FakeWebs) Create(ctx context.Context, web *v1beta1.Web, opts v1.CreateOptions) (result *v1beta1.Web, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(websResource, c.ns, web), &v1beta1.Web{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Web), err
}

// Update takes the representation of a web and updates it. Returns the server's representation of the web, and an error, if there is any.
func (c *FakeWebs) Update(ctx context.Context, web *v1beta1.Web, opts v1.UpdateOptions) (result *v1beta1.Web, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(websResource, c.ns, web), &v1beta1.Web{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Web), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWebs) UpdateStatus(ctx context.Context, web *v1beta1.Web, opts v1.UpdateOptions) (*v1beta1.Web, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(websResource, "status", c.ns, web), &v1beta1.Web{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Web), err
}

// Delete takes name of the web and deletes it. Returns an error if one occurs.
func (c *FakeWebs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(websResource, c.ns, name, opts), &v1beta1.Web{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWebs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(websResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.WebList{})
	return err
}

// Patch applies the patch and returns the patched web.
func (c *FakeWebs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Web, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(websResource, c.ns, name, pt, data, subresources...), &v1beta1.Web{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Web), err
}
