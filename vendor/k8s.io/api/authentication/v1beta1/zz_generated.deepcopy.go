// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
<<<<<<< HEAD
	runtime "k8s.io/apimachinery/pkg/runtime"
)

=======
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TokenReview).DeepCopyInto(out.(*TokenReview))
			return nil
		}, InType: reflect.TypeOf(&TokenReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TokenReviewSpec).DeepCopyInto(out.(*TokenReviewSpec))
			return nil
		}, InType: reflect.TypeOf(&TokenReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TokenReviewStatus).DeepCopyInto(out.(*TokenReviewStatus))
			return nil
		}, InType: reflect.TypeOf(&TokenReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*UserInfo).DeepCopyInto(out.(*UserInfo))
			return nil
		}, InType: reflect.TypeOf(&UserInfo{})},
	)
}

>>>>>>> Initial dep workover
// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenReview) DeepCopyInto(out *TokenReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenReview.
func (in *TokenReview) DeepCopy() *TokenReview {
	if in == nil {
		return nil
	}
	out := new(TokenReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TokenReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenReviewSpec) DeepCopyInto(out *TokenReviewSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenReviewSpec.
func (in *TokenReviewSpec) DeepCopy() *TokenReviewSpec {
	if in == nil {
		return nil
	}
	out := new(TokenReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenReviewStatus) DeepCopyInto(out *TokenReviewStatus) {
	*out = *in
	in.User.DeepCopyInto(&out.User)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenReviewStatus.
func (in *TokenReviewStatus) DeepCopy() *TokenReviewStatus {
	if in == nil {
		return nil
	}
	out := new(TokenReviewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfo) DeepCopyInto(out *UserInfo) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			(*out)[key] = make(ExtraValue, len(val))
			copy((*out)[key], val)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfo.
func (in *UserInfo) DeepCopy() *UserInfo {
	if in == nil {
		return nil
	}
	out := new(UserInfo)
	in.DeepCopyInto(out)
	return out
}
