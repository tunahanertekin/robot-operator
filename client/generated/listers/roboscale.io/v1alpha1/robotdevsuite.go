/*
Copyright 2022.

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

package v1alpha1

import (
	v1alpha1 "github.com/robolaunch/robot-operator/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RobotDevSuiteLister helps list RobotDevSuites.
// All objects returned here must be treated as read-only.
type RobotDevSuiteLister interface {
	// List lists all RobotDevSuites in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RobotDevSuite, err error)
	// RobotDevSuites returns an object that can list and get RobotDevSuites.
	RobotDevSuites(namespace string) RobotDevSuiteNamespaceLister
	RobotDevSuiteListerExpansion
}

// robotDevSuiteLister implements the RobotDevSuiteLister interface.
type robotDevSuiteLister struct {
	indexer cache.Indexer
}

// NewRobotDevSuiteLister returns a new RobotDevSuiteLister.
func NewRobotDevSuiteLister(indexer cache.Indexer) RobotDevSuiteLister {
	return &robotDevSuiteLister{indexer: indexer}
}

// List lists all RobotDevSuites in the indexer.
func (s *robotDevSuiteLister) List(selector labels.Selector) (ret []*v1alpha1.RobotDevSuite, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RobotDevSuite))
	})
	return ret, err
}

// RobotDevSuites returns an object that can list and get RobotDevSuites.
func (s *robotDevSuiteLister) RobotDevSuites(namespace string) RobotDevSuiteNamespaceLister {
	return robotDevSuiteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RobotDevSuiteNamespaceLister helps list and get RobotDevSuites.
// All objects returned here must be treated as read-only.
type RobotDevSuiteNamespaceLister interface {
	// List lists all RobotDevSuites in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RobotDevSuite, err error)
	// Get retrieves the RobotDevSuite from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RobotDevSuite, error)
	RobotDevSuiteNamespaceListerExpansion
}

// robotDevSuiteNamespaceLister implements the RobotDevSuiteNamespaceLister
// interface.
type robotDevSuiteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RobotDevSuites in the indexer for a given namespace.
func (s robotDevSuiteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RobotDevSuite, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RobotDevSuite))
	})
	return ret, err
}

// Get retrieves the RobotDevSuite from the indexer for a given namespace and name.
func (s robotDevSuiteNamespaceLister) Get(name string) (*v1alpha1.RobotDevSuite, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("robotdevsuite"), name)
	}
	return obj.(*v1alpha1.RobotDevSuite), nil
}
