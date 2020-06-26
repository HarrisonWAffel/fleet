/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	gitopscattleiov1 "github.com/rancher/gitjobs/pkg/apis/gitops.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGitJobs implements GitJobInterface
type FakeGitJobs struct {
	Fake *FakeGitopsV1
	ns   string
}

var gitjobsResource = schema.GroupVersionResource{Group: "gitops.cattle.io", Version: "v1", Resource: "gitjobs"}

var gitjobsKind = schema.GroupVersionKind{Group: "gitops.cattle.io", Version: "v1", Kind: "GitJob"}

// Get takes name of the gitJob, and returns the corresponding gitJob object, and an error if there is any.
func (c *FakeGitJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *gitopscattleiov1.GitJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gitjobsResource, c.ns, name), &gitopscattleiov1.GitJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitopscattleiov1.GitJob), err
}

// List takes label and field selectors, and returns the list of GitJobs that match those selectors.
func (c *FakeGitJobs) List(ctx context.Context, opts v1.ListOptions) (result *gitopscattleiov1.GitJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gitjobsResource, gitjobsKind, c.ns, opts), &gitopscattleiov1.GitJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gitopscattleiov1.GitJobList{ListMeta: obj.(*gitopscattleiov1.GitJobList).ListMeta}
	for _, item := range obj.(*gitopscattleiov1.GitJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gitJobs.
func (c *FakeGitJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gitjobsResource, c.ns, opts))

}

// Create takes the representation of a gitJob and creates it.  Returns the server's representation of the gitJob, and an error, if there is any.
func (c *FakeGitJobs) Create(ctx context.Context, gitJob *gitopscattleiov1.GitJob, opts v1.CreateOptions) (result *gitopscattleiov1.GitJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gitjobsResource, c.ns, gitJob), &gitopscattleiov1.GitJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitopscattleiov1.GitJob), err
}

// Update takes the representation of a gitJob and updates it. Returns the server's representation of the gitJob, and an error, if there is any.
func (c *FakeGitJobs) Update(ctx context.Context, gitJob *gitopscattleiov1.GitJob, opts v1.UpdateOptions) (result *gitopscattleiov1.GitJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gitjobsResource, c.ns, gitJob), &gitopscattleiov1.GitJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitopscattleiov1.GitJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGitJobs) UpdateStatus(ctx context.Context, gitJob *gitopscattleiov1.GitJob, opts v1.UpdateOptions) (*gitopscattleiov1.GitJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gitjobsResource, "status", c.ns, gitJob), &gitopscattleiov1.GitJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitopscattleiov1.GitJob), err
}

// Delete takes name of the gitJob and deletes it. Returns an error if one occurs.
func (c *FakeGitJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gitjobsResource, c.ns, name), &gitopscattleiov1.GitJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGitJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gitjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &gitopscattleiov1.GitJobList{})
	return err
}

// Patch applies the patch and returns the patched gitJob.
func (c *FakeGitJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *gitopscattleiov1.GitJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gitjobsResource, c.ns, name, pt, data, subresources...), &gitopscattleiov1.GitJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitopscattleiov1.GitJob), err
}