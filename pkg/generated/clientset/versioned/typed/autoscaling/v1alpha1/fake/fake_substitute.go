// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane/api/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubstitutes implements SubstituteInterface
type FakeSubstitutes struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var substitutesResource = schema.GroupVersionResource{Group: "autoscaling.crane.io", Version: "v1alpha1", Resource: "substitutes"}

var substitutesKind = schema.GroupVersionKind{Group: "autoscaling.crane.io", Version: "v1alpha1", Kind: "Substitute"}

// Get takes name of the substitute, and returns the corresponding substitute object, and an error if there is any.
func (c *FakeSubstitutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Substitute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(substitutesResource, c.ns, name), &v1alpha1.Substitute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Substitute), err
}

// List takes label and field selectors, and returns the list of Substitutes that match those selectors.
func (c *FakeSubstitutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubstituteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(substitutesResource, substitutesKind, c.ns, opts), &v1alpha1.SubstituteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubstituteList{ListMeta: obj.(*v1alpha1.SubstituteList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubstituteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested substitutes.
func (c *FakeSubstitutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(substitutesResource, c.ns, opts))

}

// Create takes the representation of a substitute and creates it.  Returns the server's representation of the substitute, and an error, if there is any.
func (c *FakeSubstitutes) Create(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.CreateOptions) (result *v1alpha1.Substitute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(substitutesResource, c.ns, substitute), &v1alpha1.Substitute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Substitute), err
}

// Update takes the representation of a substitute and updates it. Returns the server's representation of the substitute, and an error, if there is any.
func (c *FakeSubstitutes) Update(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.UpdateOptions) (result *v1alpha1.Substitute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(substitutesResource, c.ns, substitute), &v1alpha1.Substitute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Substitute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubstitutes) UpdateStatus(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.UpdateOptions) (*v1alpha1.Substitute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(substitutesResource, "status", c.ns, substitute), &v1alpha1.Substitute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Substitute), err
}

// Delete takes name of the substitute and deletes it. Returns an error if one occurs.
func (c *FakeSubstitutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(substitutesResource, c.ns, name, opts), &v1alpha1.Substitute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubstitutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(substitutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubstituteList{})
	return err
}

// Patch applies the patch and returns the patched substitute.
func (c *FakeSubstitutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Substitute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(substitutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Substitute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Substitute), err
}
