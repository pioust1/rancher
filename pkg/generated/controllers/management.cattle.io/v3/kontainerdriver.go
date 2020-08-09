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

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type KontainerDriverHandler func(string, *v3.KontainerDriver) (*v3.KontainerDriver, error)

type KontainerDriverController interface {
	generic.ControllerMeta
	KontainerDriverClient

	OnChange(ctx context.Context, name string, sync KontainerDriverHandler)
	OnRemove(ctx context.Context, name string, sync KontainerDriverHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() KontainerDriverCache
}

type KontainerDriverClient interface {
	Create(*v3.KontainerDriver) (*v3.KontainerDriver, error)
	Update(*v3.KontainerDriver) (*v3.KontainerDriver, error)
	UpdateStatus(*v3.KontainerDriver) (*v3.KontainerDriver, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.KontainerDriver, error)
	List(opts metav1.ListOptions) (*v3.KontainerDriverList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.KontainerDriver, err error)
}

type KontainerDriverCache interface {
	Get(name string) (*v3.KontainerDriver, error)
	List(selector labels.Selector) ([]*v3.KontainerDriver, error)

	AddIndexer(indexName string, indexer KontainerDriverIndexer)
	GetByIndex(indexName, key string) ([]*v3.KontainerDriver, error)
}

type KontainerDriverIndexer func(obj *v3.KontainerDriver) ([]string, error)

type kontainerDriverController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewKontainerDriverController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) KontainerDriverController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &kontainerDriverController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromKontainerDriverHandlerToHandler(sync KontainerDriverHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.KontainerDriver
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.KontainerDriver))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *kontainerDriverController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.KontainerDriver))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateKontainerDriverDeepCopyOnChange(client KontainerDriverClient, obj *v3.KontainerDriver, handler func(obj *v3.KontainerDriver) (*v3.KontainerDriver, error)) (*v3.KontainerDriver, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *kontainerDriverController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *kontainerDriverController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *kontainerDriverController) OnChange(ctx context.Context, name string, sync KontainerDriverHandler) {
	c.AddGenericHandler(ctx, name, FromKontainerDriverHandlerToHandler(sync))
}

func (c *kontainerDriverController) OnRemove(ctx context.Context, name string, sync KontainerDriverHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromKontainerDriverHandlerToHandler(sync)))
}

func (c *kontainerDriverController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *kontainerDriverController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *kontainerDriverController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *kontainerDriverController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *kontainerDriverController) Cache() KontainerDriverCache {
	return &kontainerDriverCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *kontainerDriverController) Create(obj *v3.KontainerDriver) (*v3.KontainerDriver, error) {
	result := &v3.KontainerDriver{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *kontainerDriverController) Update(obj *v3.KontainerDriver) (*v3.KontainerDriver, error) {
	result := &v3.KontainerDriver{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *kontainerDriverController) UpdateStatus(obj *v3.KontainerDriver) (*v3.KontainerDriver, error) {
	result := &v3.KontainerDriver{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *kontainerDriverController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *kontainerDriverController) Get(name string, options metav1.GetOptions) (*v3.KontainerDriver, error) {
	result := &v3.KontainerDriver{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *kontainerDriverController) List(opts metav1.ListOptions) (*v3.KontainerDriverList, error) {
	result := &v3.KontainerDriverList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *kontainerDriverController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *kontainerDriverController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.KontainerDriver, error) {
	result := &v3.KontainerDriver{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type kontainerDriverCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *kontainerDriverCache) Get(name string) (*v3.KontainerDriver, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.KontainerDriver), nil
}

func (c *kontainerDriverCache) List(selector labels.Selector) (ret []*v3.KontainerDriver, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.KontainerDriver))
	})

	return ret, err
}

func (c *kontainerDriverCache) AddIndexer(indexName string, indexer KontainerDriverIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.KontainerDriver))
		},
	}))
}

func (c *kontainerDriverCache) GetByIndex(indexName, key string) (result []*v3.KontainerDriver, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.KontainerDriver, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.KontainerDriver))
	}
	return result, nil
}

type KontainerDriverStatusHandler func(obj *v3.KontainerDriver, status v3.KontainerDriverStatus) (v3.KontainerDriverStatus, error)

type KontainerDriverGeneratingHandler func(obj *v3.KontainerDriver, status v3.KontainerDriverStatus) ([]runtime.Object, v3.KontainerDriverStatus, error)

func RegisterKontainerDriverStatusHandler(ctx context.Context, controller KontainerDriverController, condition condition.Cond, name string, handler KontainerDriverStatusHandler) {
	statusHandler := &kontainerDriverStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromKontainerDriverHandlerToHandler(statusHandler.sync))
}

func RegisterKontainerDriverGeneratingHandler(ctx context.Context, controller KontainerDriverController, apply apply.Apply,
	condition condition.Cond, name string, handler KontainerDriverGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &kontainerDriverGeneratingHandler{
		KontainerDriverGeneratingHandler: handler,
		apply:                            apply,
		name:                             name,
		gvk:                              controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterKontainerDriverStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type kontainerDriverStatusHandler struct {
	client    KontainerDriverClient
	condition condition.Cond
	handler   KontainerDriverStatusHandler
}

func (a *kontainerDriverStatusHandler) sync(key string, obj *v3.KontainerDriver) (*v3.KontainerDriver, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		obj, newErr = a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
	}
	return obj, err
}

type kontainerDriverGeneratingHandler struct {
	KontainerDriverGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *kontainerDriverGeneratingHandler) Remove(key string, obj *v3.KontainerDriver) (*v3.KontainerDriver, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.KontainerDriver{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *kontainerDriverGeneratingHandler) Handle(obj *v3.KontainerDriver, status v3.KontainerDriverStatus) (v3.KontainerDriverStatus, error) {
	objs, newStatus, err := a.KontainerDriverGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
