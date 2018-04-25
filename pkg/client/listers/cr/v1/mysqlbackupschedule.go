/*
Copyright 2017 The MySQL Operator Authors.

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

package v1

import (
	v1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MySQLBackupScheduleLister helps list MySQLBackupSchedules.
type MySQLBackupScheduleLister interface {
	// List lists all MySQLBackupSchedules in the indexer.
	List(selector labels.Selector) (ret []*v1.MySQLBackupSchedule, err error)
	// MySQLBackupSchedules returns an object that can list and get MySQLBackupSchedules.
	MySQLBackupSchedules(namespace string) MySQLBackupScheduleNamespaceLister
	MySQLBackupScheduleListerExpansion
}

// mySQLBackupScheduleLister implements the MySQLBackupScheduleLister interface.
type mySQLBackupScheduleLister struct {
	indexer cache.Indexer
}

// NewMySQLBackupScheduleLister returns a new MySQLBackupScheduleLister.
func NewMySQLBackupScheduleLister(indexer cache.Indexer) MySQLBackupScheduleLister {
	return &mySQLBackupScheduleLister{indexer: indexer}
}

// List lists all MySQLBackupSchedules in the indexer.
func (s *mySQLBackupScheduleLister) List(selector labels.Selector) (ret []*v1.MySQLBackupSchedule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MySQLBackupSchedule))
	})
	return ret, err
}

// MySQLBackupSchedules returns an object that can list and get MySQLBackupSchedules.
func (s *mySQLBackupScheduleLister) MySQLBackupSchedules(namespace string) MySQLBackupScheduleNamespaceLister {
	return mySQLBackupScheduleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MySQLBackupScheduleNamespaceLister helps list and get MySQLBackupSchedules.
type MySQLBackupScheduleNamespaceLister interface {
	// List lists all MySQLBackupSchedules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MySQLBackupSchedule, err error)
	// Get retrieves the MySQLBackupSchedule from the indexer for a given namespace and name.
	Get(name string) (*v1.MySQLBackupSchedule, error)
	MySQLBackupScheduleNamespaceListerExpansion
}

// mySQLBackupScheduleNamespaceLister implements the MySQLBackupScheduleNamespaceLister
// interface.
type mySQLBackupScheduleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MySQLBackupSchedules in the indexer for a given namespace.
func (s mySQLBackupScheduleNamespaceLister) List(selector labels.Selector) (ret []*v1.MySQLBackupSchedule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MySQLBackupSchedule))
	})
	return ret, err
}

// Get retrieves the MySQLBackupSchedule from the indexer for a given namespace and name.
func (s mySQLBackupScheduleNamespaceLister) Get(name string) (*v1.MySQLBackupSchedule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("mysqlbackupschedule"), name)
	}
	return obj.(*v1.MySQLBackupSchedule), nil
}
