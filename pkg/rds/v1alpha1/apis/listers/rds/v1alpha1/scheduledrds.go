package v1alpha1

import (
	v1alpha1 "github.com/eumel8/echoperator/pkg/rds/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScheduledRdsLister helps list ScheduledRdss.
// All objects returned here must be treated as read-only.
type ScheduledRdsLister interface {
	// List lists all ScheduledRdss in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledRds, err error)
	// ScheduledRdss returns an object that can list and get ScheduledRdss.
	ScheduledRdss(namespace string) ScheduledRdsNamespaceLister
	ScheduledRdsListerExpansion
}

// scheduledRdsLister implements the ScheduledRdsLister interface.
type scheduledRdsLister struct {
	indexer cache.Indexer
}

// NewScheduledRdsLister returns a new ScheduledRdsLister.
func NewScheduledRdsLister(indexer cache.Indexer) ScheduledRdsLister {
	return &scheduledRdsLister{indexer: indexer}
}

// List lists all ScheduledRdss in the indexer.
func (s *scheduledRdsLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledRds, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledRds))
	})
	return ret, err
}

// ScheduledRdss returns an object that can list and get ScheduledRdss.
func (s *scheduledRdsLister) ScheduledRdss(namespace string) ScheduledRdsNamespaceLister {
	return scheduledRdsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScheduledRdsNamespaceLister helps list and get ScheduledRdss.
// All objects returned here must be treated as read-only.
type ScheduledRdsNamespaceLister interface {
	// List lists all ScheduledRdss in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledRds, err error)
	// Get retrieves the ScheduledRds from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScheduledRds, error)
	ScheduledRdsNamespaceListerExpansion
}

// scheduledRdsNamespaceLister implements the ScheduledRdsNamespaceLister
// interface.
type scheduledRdsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScheduledRdss in the indexer for a given namespace.
func (s scheduledRdsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledRds, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledRds))
	})
	return ret, err
}

// Get retrieves the ScheduledRds from the indexer for a given namespace and name.
func (s scheduledRdsNamespaceLister) Get(name string) (*v1alpha1.ScheduledRds, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scheduledecho"), name)
	}
	return obj.(*v1alpha1.ScheduledRds), nil
}
