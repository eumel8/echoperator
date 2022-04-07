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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mmontes11/echoperator/pkg/echo/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EchoLister helps list Echos.
// All objects returned here must be treated as read-only.
type EchoLister interface {
	// List lists all Echos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Echo, err error)
	// Echos returns an object that can list and get Echos.
	Echos(namespace string) EchoNamespaceLister
	EchoListerExpansion
}

// echoLister implements the EchoLister interface.
type echoLister struct {
	indexer cache.Indexer
}

// NewEchoLister returns a new EchoLister.
func NewEchoLister(indexer cache.Indexer) EchoLister {
	return &echoLister{indexer: indexer}
}

// List lists all Echos in the indexer.
func (s *echoLister) List(selector labels.Selector) (ret []*v1alpha1.Echo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Echo))
	})
	return ret, err
}

// Echos returns an object that can list and get Echos.
func (s *echoLister) Echos(namespace string) EchoNamespaceLister {
	return echoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EchoNamespaceLister helps list and get Echos.
// All objects returned here must be treated as read-only.
type EchoNamespaceLister interface {
	// List lists all Echos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Echo, err error)
	// Get retrieves the Echo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Echo, error)
	EchoNamespaceListerExpansion
}

// echoNamespaceLister implements the EchoNamespaceLister
// interface.
type echoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Echos in the indexer for a given namespace.
func (s echoNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Echo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Echo))
	})
	return ret, err
}

// Get retrieves the Echo from the indexer for a given namespace and name.
func (s echoNamespaceLister) Get(name string) (*v1alpha1.Echo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("echo"), name)
	}
	return obj.(*v1alpha1.Echo), nil
}
