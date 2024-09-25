/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kyverno/kyverno/api/kyverno/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackgroundScanReportLister helps list BackgroundScanReports.
// All objects returned here must be treated as read-only.
type BackgroundScanReportLister interface {
	// List lists all BackgroundScanReports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.BackgroundScanReport, err error)
	// BackgroundScanReports returns an object that can list and get BackgroundScanReports.
	BackgroundScanReports(namespace string) BackgroundScanReportNamespaceLister
	BackgroundScanReportListerExpansion
}

// backgroundScanReportLister implements the BackgroundScanReportLister interface.
type backgroundScanReportLister struct {
	indexer cache.Indexer
}

// NewBackgroundScanReportLister returns a new BackgroundScanReportLister.
func NewBackgroundScanReportLister(indexer cache.Indexer) BackgroundScanReportLister {
	return &backgroundScanReportLister{indexer: indexer}
}

// List lists all BackgroundScanReports in the indexer.
func (s *backgroundScanReportLister) List(selector labels.Selector) (ret []*v1alpha2.BackgroundScanReport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.BackgroundScanReport))
	})
	return ret, err
}

// BackgroundScanReports returns an object that can list and get BackgroundScanReports.
func (s *backgroundScanReportLister) BackgroundScanReports(namespace string) BackgroundScanReportNamespaceLister {
	return backgroundScanReportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackgroundScanReportNamespaceLister helps list and get BackgroundScanReports.
// All objects returned here must be treated as read-only.
type BackgroundScanReportNamespaceLister interface {
	// List lists all BackgroundScanReports in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.BackgroundScanReport, err error)
	// Get retrieves the BackgroundScanReport from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.BackgroundScanReport, error)
	BackgroundScanReportNamespaceListerExpansion
}

// backgroundScanReportNamespaceLister implements the BackgroundScanReportNamespaceLister
// interface.
type backgroundScanReportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackgroundScanReports in the indexer for a given namespace.
func (s backgroundScanReportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.BackgroundScanReport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.BackgroundScanReport))
	})
	return ret, err
}

// Get retrieves the BackgroundScanReport from the indexer for a given namespace and name.
func (s backgroundScanReportNamespaceLister) Get(name string) (*v1alpha2.BackgroundScanReport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("backgroundscanreport"), name)
	}
	return obj.(*v1alpha2.BackgroundScanReport), nil
}
