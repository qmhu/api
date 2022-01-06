// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gocrane/api/analysis/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigSetLister helps list ConfigSets.
// All objects returned here must be treated as read-only.
type ConfigSetLister interface {
	// List lists all ConfigSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigSet, err error)
	// ConfigSets returns an object that can list and get ConfigSets.
	ConfigSets(namespace string) ConfigSetNamespaceLister
	ConfigSetListerExpansion
}

// configSetLister implements the ConfigSetLister interface.
type configSetLister struct {
	indexer cache.Indexer
}

// NewConfigSetLister returns a new ConfigSetLister.
func NewConfigSetLister(indexer cache.Indexer) ConfigSetLister {
	return &configSetLister{indexer: indexer}
}

// List lists all ConfigSets in the indexer.
func (s *configSetLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigSet))
	})
	return ret, err
}

// ConfigSets returns an object that can list and get ConfigSets.
func (s *configSetLister) ConfigSets(namespace string) ConfigSetNamespaceLister {
	return configSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigSetNamespaceLister helps list and get ConfigSets.
// All objects returned here must be treated as read-only.
type ConfigSetNamespaceLister interface {
	// List lists all ConfigSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigSet, err error)
	// Get retrieves the ConfigSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConfigSet, error)
	ConfigSetNamespaceListerExpansion
}

// configSetNamespaceLister implements the ConfigSetNamespaceLister
// interface.
type configSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigSets in the indexer for a given namespace.
func (s configSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigSet))
	})
	return ret, err
}

// Get retrieves the ConfigSet from the indexer for a given namespace and name.
func (s configSetNamespaceLister) Get(name string) (*v1alpha1.ConfigSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configset"), name)
	}
	return obj.(*v1alpha1.ConfigSet), nil
}