// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "elasticgpu.io/elastic-gpu/apis/elasticgpu/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeElasticGPUs implements ElasticGPUInterface
type FakeElasticGPUs struct {
	Fake *FakeElasticgpuV1alpha1
}

var elasticgpusResource = schema.GroupVersionResource{Group: "elasticgpu.io", Version: "v1alpha1", Resource: "elasticgpus"}

var elasticgpusKind = schema.GroupVersionKind{Group: "elasticgpu.io", Version: "v1alpha1", Kind: "ElasticGPU"}

// Get takes name of the elasticGPU, and returns the corresponding elasticGPU object, and an error if there is any.
func (c *FakeElasticGPUs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ElasticGPU, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(elasticgpusResource, name), &v1alpha1.ElasticGPU{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticGPU), err
}

// List takes label and field selectors, and returns the list of ElasticGPUs that match those selectors.
func (c *FakeElasticGPUs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ElasticGPUList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(elasticgpusResource, elasticgpusKind, opts), &v1alpha1.ElasticGPUList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElasticGPUList{ListMeta: obj.(*v1alpha1.ElasticGPUList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElasticGPUList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticGPUs.
func (c *FakeElasticGPUs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(elasticgpusResource, opts))
}

// Create takes the representation of a elasticGPU and creates it.  Returns the server's representation of the elasticGPU, and an error, if there is any.
func (c *FakeElasticGPUs) Create(ctx context.Context, elasticGPU *v1alpha1.ElasticGPU, opts v1.CreateOptions) (result *v1alpha1.ElasticGPU, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(elasticgpusResource, elasticGPU), &v1alpha1.ElasticGPU{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticGPU), err
}

// Update takes the representation of a elasticGPU and updates it. Returns the server's representation of the elasticGPU, and an error, if there is any.
func (c *FakeElasticGPUs) Update(ctx context.Context, elasticGPU *v1alpha1.ElasticGPU, opts v1.UpdateOptions) (result *v1alpha1.ElasticGPU, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(elasticgpusResource, elasticGPU), &v1alpha1.ElasticGPU{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticGPU), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticGPUs) UpdateStatus(ctx context.Context, elasticGPU *v1alpha1.ElasticGPU, opts v1.UpdateOptions) (*v1alpha1.ElasticGPU, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(elasticgpusResource, "status", elasticGPU), &v1alpha1.ElasticGPU{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticGPU), err
}

// Delete takes name of the elasticGPU and deletes it. Returns an error if one occurs.
func (c *FakeElasticGPUs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(elasticgpusResource, name, opts), &v1alpha1.ElasticGPU{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticGPUs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(elasticgpusResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElasticGPUList{})
	return err
}

// Patch applies the patch and returns the patched elasticGPU.
func (c *FakeElasticGPUs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ElasticGPU, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(elasticgpusResource, name, pt, data, subresources...), &v1alpha1.ElasticGPU{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticGPU), err
}