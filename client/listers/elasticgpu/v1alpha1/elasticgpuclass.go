// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "elasticgpu.io/elastic-gpu/apis/elasticgpu/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ElasticGPUClassLister helps list ElasticGPUClasses.
// All objects returned here must be treated as read-only.
type ElasticGPUClassLister interface {
	// List lists all ElasticGPUClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticGPUClass, err error)
	// Get retrieves the ElasticGPUClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ElasticGPUClass, error)
	ElasticGPUClassListerExpansion
}

// elasticGPUClassLister implements the ElasticGPUClassLister interface.
type elasticGPUClassLister struct {
	indexer cache.Indexer
}

// NewElasticGPUClassLister returns a new ElasticGPUClassLister.
func NewElasticGPUClassLister(indexer cache.Indexer) ElasticGPUClassLister {
	return &elasticGPUClassLister{indexer: indexer}
}

// List lists all ElasticGPUClasses in the indexer.
func (s *elasticGPUClassLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticGPUClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticGPUClass))
	})
	return ret, err
}

// Get retrieves the ElasticGPUClass from the index for a given name.
func (s *elasticGPUClassLister) Get(name string) (*v1alpha1.ElasticGPUClass, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elasticgpuclass"), name)
	}
	return obj.(*v1alpha1.ElasticGPUClass), nil
}