// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "carvel.dev/kapp-controller/pkg/apis/packaging/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePackageInstalls implements PackageInstallInterface
type FakePackageInstalls struct {
	Fake *FakePackagingV1alpha1
	ns   string
}

var packageinstallsResource = schema.GroupVersionResource{Group: "packaging.carvel.dev", Version: "v1alpha1", Resource: "packageinstalls"}

var packageinstallsKind = schema.GroupVersionKind{Group: "packaging.carvel.dev", Version: "v1alpha1", Kind: "PackageInstall"}

// Get takes name of the packageInstall, and returns the corresponding packageInstall object, and an error if there is any.
func (c *FakePackageInstalls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PackageInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(packageinstallsResource, c.ns, name), &v1alpha1.PackageInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageInstall), err
}

// List takes label and field selectors, and returns the list of PackageInstalls that match those selectors.
func (c *FakePackageInstalls) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PackageInstallList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(packageinstallsResource, packageinstallsKind, c.ns, opts), &v1alpha1.PackageInstallList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PackageInstallList{ListMeta: obj.(*v1alpha1.PackageInstallList).ListMeta}
	for _, item := range obj.(*v1alpha1.PackageInstallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested packageInstalls.
func (c *FakePackageInstalls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(packageinstallsResource, c.ns, opts))

}

// Create takes the representation of a packageInstall and creates it.  Returns the server's representation of the packageInstall, and an error, if there is any.
func (c *FakePackageInstalls) Create(ctx context.Context, packageInstall *v1alpha1.PackageInstall, opts v1.CreateOptions) (result *v1alpha1.PackageInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(packageinstallsResource, c.ns, packageInstall), &v1alpha1.PackageInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageInstall), err
}

// Update takes the representation of a packageInstall and updates it. Returns the server's representation of the packageInstall, and an error, if there is any.
func (c *FakePackageInstalls) Update(ctx context.Context, packageInstall *v1alpha1.PackageInstall, opts v1.UpdateOptions) (result *v1alpha1.PackageInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(packageinstallsResource, c.ns, packageInstall), &v1alpha1.PackageInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageInstall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePackageInstalls) UpdateStatus(ctx context.Context, packageInstall *v1alpha1.PackageInstall, opts v1.UpdateOptions) (*v1alpha1.PackageInstall, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(packageinstallsResource, "status", c.ns, packageInstall), &v1alpha1.PackageInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageInstall), err
}

// Delete takes name of the packageInstall and deletes it. Returns an error if one occurs.
func (c *FakePackageInstalls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(packageinstallsResource, c.ns, name, opts), &v1alpha1.PackageInstall{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePackageInstalls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(packageinstallsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PackageInstallList{})
	return err
}

// Patch applies the patch and returns the patched packageInstall.
func (c *FakePackageInstalls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PackageInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(packageinstallsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PackageInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageInstall), err
}
