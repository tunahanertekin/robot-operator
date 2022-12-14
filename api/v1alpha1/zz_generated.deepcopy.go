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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedBuildObject) DeepCopyInto(out *AttachedBuildObject) {
	*out = *in
	out.Reference = in.Reference
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedBuildObject.
func (in *AttachedBuildObject) DeepCopy() *AttachedBuildObject {
	if in == nil {
		return nil
	}
	out := new(AttachedBuildObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedLaunchObject) DeepCopyInto(out *AttachedLaunchObject) {
	*out = *in
	out.Reference = in.Reference
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedLaunchObject.
func (in *AttachedLaunchObject) DeepCopy() *AttachedLaunchObject {
	if in == nil {
		return nil
	}
	out := new(AttachedLaunchObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BridgeDistro) DeepCopyInto(out *BridgeDistro) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BridgeDistro.
func (in *BridgeDistro) DeepCopy() *BridgeDistro {
	if in == nil {
		return nil
	}
	out := new(BridgeDistro)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BridgePodStatus) DeepCopyInto(out *BridgePodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BridgePodStatus.
func (in *BridgePodStatus) DeepCopy() *BridgePodStatus {
	if in == nil {
		return nil
	}
	out := new(BridgePodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BridgeServiceStatus) DeepCopyInto(out *BridgeServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BridgeServiceStatus.
func (in *BridgeServiceStatus) DeepCopy() *BridgeServiceStatus {
	if in == nil {
		return nil
	}
	out := new(BridgeServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildManager) DeepCopyInto(out *BuildManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildManager.
func (in *BuildManager) DeepCopy() *BuildManager {
	if in == nil {
		return nil
	}
	out := new(BuildManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildManagerList) DeepCopyInto(out *BuildManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BuildManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildManagerList.
func (in *BuildManagerList) DeepCopy() *BuildManagerList {
	if in == nil {
		return nil
	}
	out := new(BuildManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildManagerSpec) DeepCopyInto(out *BuildManagerSpec) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildManagerSpec.
func (in *BuildManagerSpec) DeepCopy() *BuildManagerSpec {
	if in == nil {
		return nil
	}
	out := new(BuildManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildManagerStatus) DeepCopyInto(out *BuildManagerStatus) {
	*out = *in
	out.ScriptConfigMapStatus = in.ScriptConfigMapStatus
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]StepStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildManagerStatus.
func (in *BuildManagerStatus) DeepCopy() *BuildManagerStatus {
	if in == nil {
		return nil
	}
	out := new(BuildManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionInfo) DeepCopyInto(out *ConnectionInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionInfo.
func (in *ConnectionInfo) DeepCopy() *ConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(ConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServer) DeepCopyInto(out *DiscoveryServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServer.
func (in *DiscoveryServer) DeepCopy() *DiscoveryServer {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServerConfigMapStatus) DeepCopyInto(out *DiscoveryServerConfigMapStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServerConfigMapStatus.
func (in *DiscoveryServerConfigMapStatus) DeepCopy() *DiscoveryServerConfigMapStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServerConfigMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServerInstanceStatus) DeepCopyInto(out *DiscoveryServerInstanceStatus) {
	*out = *in
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServerInstanceStatus.
func (in *DiscoveryServerInstanceStatus) DeepCopy() *DiscoveryServerInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServerInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServerList) DeepCopyInto(out *DiscoveryServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveryServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServerList.
func (in *DiscoveryServerList) DeepCopy() *DiscoveryServerList {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServerPodStatus) DeepCopyInto(out *DiscoveryServerPodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServerPodStatus.
func (in *DiscoveryServerPodStatus) DeepCopy() *DiscoveryServerPodStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServerPodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServerServiceStatus) DeepCopyInto(out *DiscoveryServerServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServerServiceStatus.
func (in *DiscoveryServerServiceStatus) DeepCopy() *DiscoveryServerServiceStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServerServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServerSpec) DeepCopyInto(out *DiscoveryServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServerSpec.
func (in *DiscoveryServerSpec) DeepCopy() *DiscoveryServerSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServerStatus) DeepCopyInto(out *DiscoveryServerStatus) {
	*out = *in
	out.ServiceStatus = in.ServiceStatus
	out.PodStatus = in.PodStatus
	out.ConfigMapStatus = in.ConfigMapStatus
	out.ConnectionInfo = in.ConnectionInfo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServerStatus.
func (in *DiscoveryServerStatus) DeepCopy() *DiscoveryServerStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Launch) DeepCopyInto(out *Launch) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Prelaunch = in.Prelaunch
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Launch.
func (in *Launch) DeepCopy() *Launch {
	if in == nil {
		return nil
	}
	out := new(Launch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchManager) DeepCopyInto(out *LaunchManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchManager.
func (in *LaunchManager) DeepCopy() *LaunchManager {
	if in == nil {
		return nil
	}
	out := new(LaunchManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchManagerList) DeepCopyInto(out *LaunchManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LaunchManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchManagerList.
func (in *LaunchManagerList) DeepCopy() *LaunchManagerList {
	if in == nil {
		return nil
	}
	out := new(LaunchManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchManagerSpec) DeepCopyInto(out *LaunchManagerSpec) {
	*out = *in
	if in.Launch != nil {
		in, out := &in.Launch, &out.Launch
		*out = make(map[string]Launch, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchManagerSpec.
func (in *LaunchManagerSpec) DeepCopy() *LaunchManagerSpec {
	if in == nil {
		return nil
	}
	out := new(LaunchManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchManagerStatus) DeepCopyInto(out *LaunchManagerStatus) {
	*out = *in
	in.LaunchPodStatus.DeepCopyInto(&out.LaunchPodStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchManagerStatus.
func (in *LaunchManagerStatus) DeepCopy() *LaunchManagerStatus {
	if in == nil {
		return nil
	}
	out := new(LaunchManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchPodStatus) DeepCopyInto(out *LaunchPodStatus) {
	*out = *in
	if in.LaunchStatus != nil {
		in, out := &in.LaunchStatus, &out.LaunchStatus
		*out = make(map[string]LaunchStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchPodStatus.
func (in *LaunchPodStatus) DeepCopy() *LaunchPodStatus {
	if in == nil {
		return nil
	}
	out := new(LaunchPodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchStatus) DeepCopyInto(out *LaunchStatus) {
	*out = *in
	in.ContainerStatus.DeepCopyInto(&out.ContainerStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchStatus.
func (in *LaunchStatus) DeepCopy() *LaunchStatus {
	if in == nil {
		return nil
	}
	out := new(LaunchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoaderJobStatus) DeepCopyInto(out *LoaderJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoaderJobStatus.
func (in *LoaderJobStatus) DeepCopy() *LoaderJobStatus {
	if in == nil {
		return nil
	}
	out := new(LoaderJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prelaunch) DeepCopyInto(out *Prelaunch) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prelaunch.
func (in *Prelaunch) DeepCopy() *Prelaunch {
	if in == nil {
		return nil
	}
	out := new(Prelaunch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSBridge) DeepCopyInto(out *ROSBridge) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSBridge.
func (in *ROSBridge) DeepCopy() *ROSBridge {
	if in == nil {
		return nil
	}
	out := new(ROSBridge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROSBridge) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSBridgeInstanceStatus) DeepCopyInto(out *ROSBridgeInstanceStatus) {
	*out = *in
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSBridgeInstanceStatus.
func (in *ROSBridgeInstanceStatus) DeepCopy() *ROSBridgeInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ROSBridgeInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSBridgeList) DeepCopyInto(out *ROSBridgeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ROSBridge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSBridgeList.
func (in *ROSBridgeList) DeepCopy() *ROSBridgeList {
	if in == nil {
		return nil
	}
	out := new(ROSBridgeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ROSBridgeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSBridgeSpec) DeepCopyInto(out *ROSBridgeSpec) {
	*out = *in
	out.ROS = in.ROS
	out.ROS2 = in.ROS2
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSBridgeSpec.
func (in *ROSBridgeSpec) DeepCopy() *ROSBridgeSpec {
	if in == nil {
		return nil
	}
	out := new(ROSBridgeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ROSBridgeStatus) DeepCopyInto(out *ROSBridgeStatus) {
	*out = *in
	out.PodStatus = in.PodStatus
	out.ServiceStatus = in.ServiceStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ROSBridgeStatus.
func (in *ROSBridgeStatus) DeepCopy() *ROSBridgeStatus {
	if in == nil {
		return nil
	}
	out := new(ROSBridgeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Robot) DeepCopyInto(out *Robot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Robot.
func (in *Robot) DeepCopy() *Robot {
	if in == nil {
		return nil
	}
	out := new(Robot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Robot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotDevSuite) DeepCopyInto(out *RobotDevSuite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotDevSuite.
func (in *RobotDevSuite) DeepCopy() *RobotDevSuite {
	if in == nil {
		return nil
	}
	out := new(RobotDevSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotDevSuite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotDevSuiteList) DeepCopyInto(out *RobotDevSuiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RobotDevSuite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotDevSuiteList.
func (in *RobotDevSuiteList) DeepCopy() *RobotDevSuiteList {
	if in == nil {
		return nil
	}
	out := new(RobotDevSuiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotDevSuiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotDevSuiteSpec) DeepCopyInto(out *RobotDevSuiteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotDevSuiteSpec.
func (in *RobotDevSuiteSpec) DeepCopy() *RobotDevSuiteSpec {
	if in == nil {
		return nil
	}
	out := new(RobotDevSuiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotDevSuiteStatus) DeepCopyInto(out *RobotDevSuiteStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotDevSuiteStatus.
func (in *RobotDevSuiteStatus) DeepCopy() *RobotDevSuiteStatus {
	if in == nil {
		return nil
	}
	out := new(RobotDevSuiteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotList) DeepCopyInto(out *RobotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Robot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotList.
func (in *RobotList) DeepCopy() *RobotList {
	if in == nil {
		return nil
	}
	out := new(RobotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotSpec) DeepCopyInto(out *RobotSpec) {
	*out = *in
	out.Storage = in.Storage
	out.DiscoveryServerTemplate = in.DiscoveryServerTemplate
	out.ROSBridgeTemplate = in.ROSBridgeTemplate
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.RootDNSConfig = in.RootDNSConfig
	out.TLSSecretReference = in.TLSSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotSpec.
func (in *RobotSpec) DeepCopy() *RobotSpec {
	if in == nil {
		return nil
	}
	out := new(RobotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotStatus) DeepCopyInto(out *RobotStatus) {
	*out = *in
	out.VolumeStatuses = in.VolumeStatuses
	out.DiscoveryServerStatus = in.DiscoveryServerStatus
	out.ROSBridgeStatus = in.ROSBridgeStatus
	out.LoaderJobStatus = in.LoaderJobStatus
	in.AttachedBuildObject.DeepCopyInto(&out.AttachedBuildObject)
	if in.AttachedLaunchObjects != nil {
		in, out := &in.AttachedLaunchObjects, &out.AttachedLaunchObjects
		*out = make([]AttachedLaunchObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotStatus.
func (in *RobotStatus) DeepCopy() *RobotStatus {
	if in == nil {
		return nil
	}
	out := new(RobotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDI) DeepCopyInto(out *RobotVDI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDI.
func (in *RobotVDI) DeepCopy() *RobotVDI {
	if in == nil {
		return nil
	}
	out := new(RobotVDI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotVDI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDIIngressStatus) DeepCopyInto(out *RobotVDIIngressStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDIIngressStatus.
func (in *RobotVDIIngressStatus) DeepCopy() *RobotVDIIngressStatus {
	if in == nil {
		return nil
	}
	out := new(RobotVDIIngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDIList) DeepCopyInto(out *RobotVDIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RobotVDI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDIList.
func (in *RobotVDIList) DeepCopy() *RobotVDIList {
	if in == nil {
		return nil
	}
	out := new(RobotVDIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotVDIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDIPVCStatus) DeepCopyInto(out *RobotVDIPVCStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDIPVCStatus.
func (in *RobotVDIPVCStatus) DeepCopy() *RobotVDIPVCStatus {
	if in == nil {
		return nil
	}
	out := new(RobotVDIPVCStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDIPodStatus) DeepCopyInto(out *RobotVDIPodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDIPodStatus.
func (in *RobotVDIPodStatus) DeepCopy() *RobotVDIPodStatus {
	if in == nil {
		return nil
	}
	out := new(RobotVDIPodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDIServiceTCPStatus) DeepCopyInto(out *RobotVDIServiceTCPStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDIServiceTCPStatus.
func (in *RobotVDIServiceTCPStatus) DeepCopy() *RobotVDIServiceTCPStatus {
	if in == nil {
		return nil
	}
	out := new(RobotVDIServiceTCPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDIServiceUDPStatus) DeepCopyInto(out *RobotVDIServiceUDPStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDIServiceUDPStatus.
func (in *RobotVDIServiceUDPStatus) DeepCopy() *RobotVDIServiceUDPStatus {
	if in == nil {
		return nil
	}
	out := new(RobotVDIServiceUDPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDISpec) DeepCopyInto(out *RobotVDISpec) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDISpec.
func (in *RobotVDISpec) DeepCopy() *RobotVDISpec {
	if in == nil {
		return nil
	}
	out := new(RobotVDISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotVDIStatus) DeepCopyInto(out *RobotVDIStatus) {
	*out = *in
	out.PodStatus = in.PodStatus
	out.ServiceTCPStatus = in.ServiceTCPStatus
	out.ServiceUDPStatus = in.ServiceUDPStatus
	out.IngressStatus = in.IngressStatus
	out.PVCStatus = in.PVCStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotVDIStatus.
func (in *RobotVDIStatus) DeepCopy() *RobotVDIStatus {
	if in == nil {
		return nil
	}
	out := new(RobotVDIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootDNSConfig) DeepCopyInto(out *RootDNSConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootDNSConfig.
func (in *RootDNSConfig) DeepCopy() *RootDNSConfig {
	if in == nil {
		return nil
	}
	out := new(RootDNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptConfigMapStatus) DeepCopyInto(out *ScriptConfigMapStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptConfigMapStatus.
func (in *ScriptConfigMapStatus) DeepCopy() *ScriptConfigMapStatus {
	if in == nil {
		return nil
	}
	out := new(ScriptConfigMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepStatus) DeepCopyInto(out *StepStatus) {
	*out = *in
	in.Step.DeepCopyInto(&out.Step)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepStatus.
func (in *StepStatus) DeepCopy() *StepStatus {
	if in == nil {
		return nil
	}
	out := new(StepStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	out.StorageClassConfig = in.StorageClassConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassConfig) DeepCopyInto(out *StorageClassConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassConfig.
func (in *StorageClassConfig) DeepCopy() *StorageClassConfig {
	if in == nil {
		return nil
	}
	out := new(StorageClassConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSecretReference) DeepCopyInto(out *TLSSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSecretReference.
func (in *TLSSecretReference) DeepCopy() *TLSSecretReference {
	if in == nil {
		return nil
	}
	out := new(TLSSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeStatus) DeepCopyInto(out *VolumeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeStatus.
func (in *VolumeStatus) DeepCopy() *VolumeStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeStatuses) DeepCopyInto(out *VolumeStatuses) {
	*out = *in
	out.Var = in.Var
	out.Etc = in.Etc
	out.Usr = in.Usr
	out.Opt = in.Opt
	out.Display = in.Display
	out.Workspace = in.Workspace
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeStatuses.
func (in *VolumeStatuses) DeepCopy() *VolumeStatuses {
	if in == nil {
		return nil
	}
	out := new(VolumeStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make(map[string]Repository, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}
