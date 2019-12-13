// +build !ignore_autogenerated

/*
Copyright 2019 The Kruise Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BroadcastJob) DeepCopyInto(out *BroadcastJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BroadcastJob.
func (in *BroadcastJob) DeepCopy() *BroadcastJob {
	if in == nil {
		return nil
	}
	out := new(BroadcastJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BroadcastJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BroadcastJobList) DeepCopyInto(out *BroadcastJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BroadcastJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BroadcastJobList.
func (in *BroadcastJobList) DeepCopy() *BroadcastJobList {
	if in == nil {
		return nil
	}
	out := new(BroadcastJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BroadcastJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BroadcastJobSpec) DeepCopyInto(out *BroadcastJobSpec) {
	*out = *in
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.Template.DeepCopyInto(&out.Template)
	in.CompletionPolicy.DeepCopyInto(&out.CompletionPolicy)
	out.FailurePolicy = in.FailurePolicy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BroadcastJobSpec.
func (in *BroadcastJobSpec) DeepCopy() *BroadcastJobSpec {
	if in == nil {
		return nil
	}
	out := new(BroadcastJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BroadcastJobStatus) DeepCopyInto(out *BroadcastJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JobCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BroadcastJobStatus.
func (in *BroadcastJobStatus) DeepCopy() *BroadcastJobStatus {
	if in == nil {
		return nil
	}
	out := new(BroadcastJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSet) DeepCopyInto(out *CloneSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSet.
func (in *CloneSet) DeepCopy() *CloneSet {
	if in == nil {
		return nil
	}
	out := new(CloneSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloneSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSetCondition) DeepCopyInto(out *CloneSetCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSetCondition.
func (in *CloneSetCondition) DeepCopy() *CloneSetCondition {
	if in == nil {
		return nil
	}
	out := new(CloneSetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSetInPlaceUpdateStrategy) DeepCopyInto(out *CloneSetInPlaceUpdateStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSetInPlaceUpdateStrategy.
func (in *CloneSetInPlaceUpdateStrategy) DeepCopy() *CloneSetInPlaceUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(CloneSetInPlaceUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSetList) DeepCopyInto(out *CloneSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloneSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSetList.
func (in *CloneSetList) DeepCopy() *CloneSetList {
	if in == nil {
		return nil
	}
	out := new(CloneSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloneSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSetScaleStrategy) DeepCopyInto(out *CloneSetScaleStrategy) {
	*out = *in
	if in.PodsToDelete != nil {
		in, out := &in.PodsToDelete, &out.PodsToDelete
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSetScaleStrategy.
func (in *CloneSetScaleStrategy) DeepCopy() *CloneSetScaleStrategy {
	if in == nil {
		return nil
	}
	out := new(CloneSetScaleStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSetSpec) DeepCopyInto(out *CloneSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ScaleStrategy.DeepCopyInto(&out.ScaleStrategy)
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSetSpec.
func (in *CloneSetSpec) DeepCopy() *CloneSetSpec {
	if in == nil {
		return nil
	}
	out := new(CloneSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSetStatus) DeepCopyInto(out *CloneSetStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CloneSetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSetStatus.
func (in *CloneSetStatus) DeepCopy() *CloneSetStatus {
	if in == nil {
		return nil
	}
	out := new(CloneSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSetUpdateStrategy) DeepCopyInto(out *CloneSetUpdateStrategy) {
	*out = *in
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(int32)
		**out = **in
	}
	if in.PriorityStrategy != nil {
		in, out := &in.PriorityStrategy, &out.PriorityStrategy
		*out = new(UpdatePriorityStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.InPlaceUpdateStrategy != nil {
		in, out := &in.InPlaceUpdateStrategy, &out.InPlaceUpdateStrategy
		*out = new(CloneSetInPlaceUpdateStrategy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSetUpdateStrategy.
func (in *CloneSetUpdateStrategy) DeepCopy() *CloneSetUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(CloneSetUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionPolicy) DeepCopyInto(out *CompletionPolicy) {
	*out = *in
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionPolicy.
func (in *CompletionPolicy) DeepCopy() *CompletionPolicy {
	if in == nil {
		return nil
	}
	out := new(CompletionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailurePolicy) DeepCopyInto(out *FailurePolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailurePolicy.
func (in *FailurePolicy) DeepCopy() *FailurePolicy {
	if in == nil {
		return nil
	}
	out := new(FailurePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InPlaceUpdateContainerStatus) DeepCopyInto(out *InPlaceUpdateContainerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InPlaceUpdateContainerStatus.
func (in *InPlaceUpdateContainerStatus) DeepCopy() *InPlaceUpdateContainerStatus {
	if in == nil {
		return nil
	}
	out := new(InPlaceUpdateContainerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InPlaceUpdateState) DeepCopyInto(out *InPlaceUpdateState) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	if in.LastContainerStatuses != nil {
		in, out := &in.LastContainerStatuses, &out.LastContainerStatuses
		*out = make(map[string]InPlaceUpdateContainerStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InPlaceUpdateState.
func (in *InPlaceUpdateState) DeepCopy() *InPlaceUpdateState {
	if in == nil {
		return nil
	}
	out := new(InPlaceUpdateState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobCondition) DeepCopyInto(out *JobCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobCondition.
func (in *JobCondition) DeepCopy() *JobCondition {
	if in == nil {
		return nil
	}
	out := new(JobCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualUpdate) DeepCopyInto(out *ManualUpdate) {
	*out = *in
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualUpdate.
func (in *ManualUpdate) DeepCopy() *ManualUpdate {
	if in == nil {
		return nil
	}
	out := new(ManualUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateSidecarSet) DeepCopyInto(out *RollingUpdateSidecarSet) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateSidecarSet.
func (in *RollingUpdateSidecarSet) DeepCopy() *RollingUpdateSidecarSet {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateSidecarSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateStatefulSetStrategy) DeepCopyInto(out *RollingUpdateStatefulSetStrategy) {
	*out = *in
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(int32)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.PriorityStrategy != nil {
		in, out := &in.PriorityStrategy, &out.PriorityStrategy
		*out = new(UpdatePriorityStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateStatefulSetStrategy.
func (in *RollingUpdateStatefulSetStrategy) DeepCopy() *RollingUpdateStatefulSetStrategy {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateStatefulSetStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarContainer) DeepCopyInto(out *SidecarContainer) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarContainer.
func (in *SidecarContainer) DeepCopy() *SidecarContainer {
	if in == nil {
		return nil
	}
	out := new(SidecarContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSet) DeepCopyInto(out *SidecarSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSet.
func (in *SidecarSet) DeepCopy() *SidecarSet {
	if in == nil {
		return nil
	}
	out := new(SidecarSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SidecarSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetList) DeepCopyInto(out *SidecarSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SidecarSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetList.
func (in *SidecarSetList) DeepCopy() *SidecarSetList {
	if in == nil {
		return nil
	}
	out := new(SidecarSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SidecarSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetSpec) DeepCopyInto(out *SidecarSetSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]SidecarContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetSpec.
func (in *SidecarSetSpec) DeepCopy() *SidecarSetSpec {
	if in == nil {
		return nil
	}
	out := new(SidecarSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetStatus) DeepCopyInto(out *SidecarSetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetStatus.
func (in *SidecarSetStatus) DeepCopy() *SidecarSetStatus {
	if in == nil {
		return nil
	}
	out := new(SidecarSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetUpdateStrategy) DeepCopyInto(out *SidecarSetUpdateStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateSidecarSet)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetUpdateStrategy.
func (in *SidecarSetUpdateStrategy) DeepCopy() *SidecarSetUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(SidecarSetUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSet) DeepCopyInto(out *StatefulSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSet.
func (in *StatefulSet) DeepCopy() *StatefulSet {
	if in == nil {
		return nil
	}
	out := new(StatefulSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetList) DeepCopyInto(out *StatefulSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatefulSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetList.
func (in *StatefulSetList) DeepCopy() *StatefulSetList {
	if in == nil {
		return nil
	}
	out := new(StatefulSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetSpec) DeepCopyInto(out *StatefulSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetSpec.
func (in *StatefulSetSpec) DeepCopy() *StatefulSetSpec {
	if in == nil {
		return nil
	}
	out := new(StatefulSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetStatus) DeepCopyInto(out *StatefulSetStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]appsv1.StatefulSetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetStatus.
func (in *StatefulSetStatus) DeepCopy() *StatefulSetStatus {
	if in == nil {
		return nil
	}
	out := new(StatefulSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetTemplateSpec) DeepCopyInto(out *StatefulSetTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetTemplateSpec.
func (in *StatefulSetTemplateSpec) DeepCopy() *StatefulSetTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(StatefulSetTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetUpdateStrategy) DeepCopyInto(out *StatefulSetUpdateStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateStatefulSetStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetUpdateStrategy.
func (in *StatefulSetUpdateStrategy) DeepCopy() *StatefulSetUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(StatefulSetUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subset) DeepCopyInto(out *Subset) {
	*out = *in
	in.NodeSelectorTerm.DeepCopyInto(&out.NodeSelectorTerm)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subset.
func (in *Subset) DeepCopy() *Subset {
	if in == nil {
		return nil
	}
	out := new(Subset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubsetTemplate) DeepCopyInto(out *SubsetTemplate) {
	*out = *in
	if in.StatefulSetTemplate != nil {
		in, out := &in.StatefulSetTemplate, &out.StatefulSetTemplate
		*out = new(StatefulSetTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubsetTemplate.
func (in *SubsetTemplate) DeepCopy() *SubsetTemplate {
	if in == nil {
		return nil
	}
	out := new(SubsetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topology) DeepCopyInto(out *Topology) {
	*out = *in
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]Subset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topology.
func (in *Topology) DeepCopy() *Topology {
	if in == nil {
		return nil
	}
	out := new(Topology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitedDeployment) DeepCopyInto(out *UnitedDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitedDeployment.
func (in *UnitedDeployment) DeepCopy() *UnitedDeployment {
	if in == nil {
		return nil
	}
	out := new(UnitedDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnitedDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitedDeploymentCondition) DeepCopyInto(out *UnitedDeploymentCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitedDeploymentCondition.
func (in *UnitedDeploymentCondition) DeepCopy() *UnitedDeploymentCondition {
	if in == nil {
		return nil
	}
	out := new(UnitedDeploymentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitedDeploymentList) DeepCopyInto(out *UnitedDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UnitedDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitedDeploymentList.
func (in *UnitedDeploymentList) DeepCopy() *UnitedDeploymentList {
	if in == nil {
		return nil
	}
	out := new(UnitedDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnitedDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitedDeploymentSpec) DeepCopyInto(out *UnitedDeploymentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	in.Topology.DeepCopyInto(&out.Topology)
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitedDeploymentSpec.
func (in *UnitedDeploymentSpec) DeepCopy() *UnitedDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(UnitedDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitedDeploymentStatus) DeepCopyInto(out *UnitedDeploymentStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.SubsetReplicas != nil {
		in, out := &in.SubsetReplicas, &out.SubsetReplicas
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]UnitedDeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpdateStatus != nil {
		in, out := &in.UpdateStatus, &out.UpdateStatus
		*out = new(UpdateStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitedDeploymentStatus.
func (in *UnitedDeploymentStatus) DeepCopy() *UnitedDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(UnitedDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnitedDeploymentUpdateStrategy) DeepCopyInto(out *UnitedDeploymentUpdateStrategy) {
	*out = *in
	if in.ManualUpdate != nil {
		in, out := &in.ManualUpdate, &out.ManualUpdate
		*out = new(ManualUpdate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnitedDeploymentUpdateStrategy.
func (in *UnitedDeploymentUpdateStrategy) DeepCopy() *UnitedDeploymentUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(UnitedDeploymentUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdatePriorityOrderTerm) DeepCopyInto(out *UpdatePriorityOrderTerm) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdatePriorityOrderTerm.
func (in *UpdatePriorityOrderTerm) DeepCopy() *UpdatePriorityOrderTerm {
	if in == nil {
		return nil
	}
	out := new(UpdatePriorityOrderTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdatePriorityStrategy) DeepCopyInto(out *UpdatePriorityStrategy) {
	*out = *in
	if in.OrderPriority != nil {
		in, out := &in.OrderPriority, &out.OrderPriority
		*out = make([]UpdatePriorityOrderTerm, len(*in))
		copy(*out, *in)
	}
	if in.WeightPriority != nil {
		in, out := &in.WeightPriority, &out.WeightPriority
		*out = make([]UpdatePriorityWeightTerm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdatePriorityStrategy.
func (in *UpdatePriorityStrategy) DeepCopy() *UpdatePriorityStrategy {
	if in == nil {
		return nil
	}
	out := new(UpdatePriorityStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdatePriorityWeightTerm) DeepCopyInto(out *UpdatePriorityWeightTerm) {
	*out = *in
	in.MatchSelector.DeepCopyInto(&out.MatchSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdatePriorityWeightTerm.
func (in *UpdatePriorityWeightTerm) DeepCopy() *UpdatePriorityWeightTerm {
	if in == nil {
		return nil
	}
	out := new(UpdatePriorityWeightTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStatus) DeepCopyInto(out *UpdateStatus) {
	*out = *in
	if in.CurrentPartitions != nil {
		in, out := &in.CurrentPartitions, &out.CurrentPartitions
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStatus.
func (in *UpdateStatus) DeepCopy() *UpdateStatus {
	if in == nil {
		return nil
	}
	out := new(UpdateStatus)
	in.DeepCopyInto(out)
	return out
}
