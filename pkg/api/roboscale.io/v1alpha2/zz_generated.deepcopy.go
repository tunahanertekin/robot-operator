//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROS2Workload) DeepCopyInto(out *ROS2Workload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROS2Workload.
func (in *ROS2Workload) DeepCopy() *ROS2Workload {
	if in == nil {
		return nil
	}
	out := new(ROS2Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROS2Workload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROS2WorkloadList) DeepCopyInto(out *ROS2WorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ROS2Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROS2WorkloadList.
func (in *ROS2WorkloadList) DeepCopy() *ROS2WorkloadList {
	if in == nil {
		return nil
	}
	out := new(ROS2WorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROS2WorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROS2WorkloadSpec) DeepCopyInto(out *ROS2WorkloadSpec) {
	*out = *in
	out.DiscoveryServerTemplate = in.DiscoveryServerTemplate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROS2WorkloadSpec.
func (in *ROS2WorkloadSpec) DeepCopy() *ROS2WorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(ROS2WorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROS2WorkloadStatus) DeepCopyInto(out *ROS2WorkloadStatus) {
	*out = *in
	out.DiscoveryServerStatus = in.DiscoveryServerStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROS2WorkloadStatus.
func (in *ROS2WorkloadStatus) DeepCopy() *ROS2WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(ROS2WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}
