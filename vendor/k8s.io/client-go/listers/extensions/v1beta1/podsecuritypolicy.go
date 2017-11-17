/*
<<<<<<< HEAD
Copyright 2018 The Kubernetes Authors.
=======
Copyright 2017 The Kubernetes Authors.
>>>>>>> Initial dep workover

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

// This file was automatically generated by lister-gen

package v1beta1

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
<<<<<<< HEAD
=======
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
>>>>>>> Initial dep workover
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodSecurityPolicyLister helps list PodSecurityPolicies.
type PodSecurityPolicyLister interface {
	// List lists all PodSecurityPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.PodSecurityPolicy, err error)
	// Get retrieves the PodSecurityPolicy from the index for a given name.
	Get(name string) (*v1beta1.PodSecurityPolicy, error)
	PodSecurityPolicyListerExpansion
}

// podSecurityPolicyLister implements the PodSecurityPolicyLister interface.
type podSecurityPolicyLister struct {
	indexer cache.Indexer
}

// NewPodSecurityPolicyLister returns a new PodSecurityPolicyLister.
func NewPodSecurityPolicyLister(indexer cache.Indexer) PodSecurityPolicyLister {
	return &podSecurityPolicyLister{indexer: indexer}
}

// List lists all PodSecurityPolicies in the indexer.
func (s *podSecurityPolicyLister) List(selector labels.Selector) (ret []*v1beta1.PodSecurityPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PodSecurityPolicy))
	})
	return ret, err
}

// Get retrieves the PodSecurityPolicy from the index for a given name.
func (s *podSecurityPolicyLister) Get(name string) (*v1beta1.PodSecurityPolicy, error) {
<<<<<<< HEAD
	obj, exists, err := s.indexer.GetByKey(name)
=======
	key := &v1beta1.PodSecurityPolicy{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
>>>>>>> Initial dep workover
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("podsecuritypolicy"), name)
	}
	return obj.(*v1beta1.PodSecurityPolicy), nil
}
