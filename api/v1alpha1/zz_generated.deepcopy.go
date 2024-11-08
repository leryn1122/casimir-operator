//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (cr *CsNodeDaemon) DeepCopyInto(out *CsNodeDaemon) {
	*out = *cr
	out.TypeMeta = cr.TypeMeta
	cr.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = cr.Spec
	out.Status = cr.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CsNodeDaemon.
func (cr *CsNodeDaemon) DeepCopy() *CsNodeDaemon {
	if cr == nil {
		return nil
	}
	out := new(CsNodeDaemon)
	cr.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (cr *CsNodeDaemon) DeepCopyObject() runtime.Object {
	if c := cr.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CsNodeDaemonList) DeepCopyInto(out *CsNodeDaemonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CsNodeDaemon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CsNodeDaemonList.
func (in *CsNodeDaemonList) DeepCopy() *CsNodeDaemonList {
	if in == nil {
		return nil
	}
	out := new(CsNodeDaemonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CsNodeDaemonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CsNodeDaemonSpec) DeepCopyInto(out *CsNodeDaemonSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CsNodeDaemonSpec.
func (in *CsNodeDaemonSpec) DeepCopy() *CsNodeDaemonSpec {
	if in == nil {
		return nil
	}
	out := new(CsNodeDaemonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CsNodeDaemonStatus) DeepCopyInto(out *CsNodeDaemonStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CsNodeDaemonStatus.
func (in *CsNodeDaemonStatus) DeepCopy() *CsNodeDaemonStatus {
	if in == nil {
		return nil
	}
	out := new(CsNodeDaemonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}
