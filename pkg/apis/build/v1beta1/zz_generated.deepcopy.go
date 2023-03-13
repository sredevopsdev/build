//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Build) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildList) DeepCopyInto(out *BuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Build, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildList.
func (in *BuildList) DeepCopy() *BuildList {
	if in == nil {
		return nil
	}
	out := new(BuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRef) DeepCopyInto(out *BuildRef) {
	*out = *in
	if in.APIVersion != nil {
		in, out := &in.APIVersion, &out.APIVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRef.
func (in *BuildRef) DeepCopy() *BuildRef {
	if in == nil {
		return nil
	}
	out := new(BuildRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRetention) DeepCopyInto(out *BuildRetention) {
	*out = *in
	if in.FailedLimit != nil {
		in, out := &in.FailedLimit, &out.FailedLimit
		*out = new(uint)
		**out = **in
	}
	if in.SucceededLimit != nil {
		in, out := &in.SucceededLimit, &out.SucceededLimit
		*out = new(uint)
		**out = **in
	}
	if in.TTLAfterFailed != nil {
		in, out := &in.TTLAfterFailed, &out.TTLAfterFailed
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TTLAfterSucceeded != nil {
		in, out := &in.TTLAfterSucceeded, &out.TTLAfterSucceeded
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRetention.
func (in *BuildRetention) DeepCopy() *BuildRetention {
	if in == nil {
		return nil
	}
	out := new(BuildRetention)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRun) DeepCopyInto(out *BuildRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRun.
func (in *BuildRun) DeepCopy() *BuildRun {
	if in == nil {
		return nil
	}
	out := new(BuildRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRunList) DeepCopyInto(out *BuildRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BuildRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRunList.
func (in *BuildRunList) DeepCopy() *BuildRunList {
	if in == nil {
		return nil
	}
	out := new(BuildRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRunRetention) DeepCopyInto(out *BuildRunRetention) {
	*out = *in
	if in.TTLAfterFailed != nil {
		in, out := &in.TTLAfterFailed, &out.TTLAfterFailed
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TTLAfterSucceeded != nil {
		in, out := &in.TTLAfterSucceeded, &out.TTLAfterSucceeded
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRunRetention.
func (in *BuildRunRetention) DeepCopy() *BuildRunRetention {
	if in == nil {
		return nil
	}
	out := new(BuildRunRetention)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRunSpec) DeepCopyInto(out *BuildRunSpec) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(ReferencedBuild)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccount)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ParamValues != nil {
		in, out := &in.ParamValues, &out.ParamValues
		*out = make([]ParamValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(Image)
		(*in).DeepCopyInto(*out)
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BuildRunRequestedState)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(BuildRunRetention)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]BuildVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRunSpec.
func (in *BuildRunSpec) DeepCopy() *BuildRunSpec {
	if in == nil {
		return nil
	}
	out := new(BuildRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRunStatus) DeepCopyInto(out *BuildRunStatus) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SourceResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(Output)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TaskRunName != nil {
		in, out := &in.TaskRunName, &out.TaskRunName
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.BuildSpec != nil {
		in, out := &in.BuildSpec, &out.BuildSpec
		*out = new(BuildSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.FailedAt != nil {
		in, out := &in.FailedAt, &out.FailedAt
		*out = new(FailedAt)
		**out = **in
	}
	if in.FailureDetails != nil {
		in, out := &in.FailureDetails, &out.FailureDetails
		*out = new(FailureDetails)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRunStatus.
func (in *BuildRunStatus) DeepCopy() *BuildRunStatus {
	if in == nil {
		return nil
	}
	out := new(BuildRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildSpec) DeepCopyInto(out *BuildSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(Trigger)
		(*in).DeepCopyInto(*out)
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.ParamValues != nil {
		in, out := &in.ParamValues, &out.ParamValues
		*out = make([]ParamValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Output.DeepCopyInto(&out.Output)
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(BuildRetention)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]BuildVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildSpec.
func (in *BuildSpec) DeepCopy() *BuildSpec {
	if in == nil {
		return nil
	}
	out := new(BuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStatus) DeepCopyInto(out *BuildStatus) {
	*out = *in
	if in.Registered != nil {
		in, out := &in.Registered, &out.Registered
		*out = new(corev1.ConditionStatus)
		**out = **in
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(BuildReason)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStatus.
func (in *BuildStatus) DeepCopy() *BuildStatus {
	if in == nil {
		return nil
	}
	out := new(BuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStrategy) DeepCopyInto(out *BuildStrategy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStrategy.
func (in *BuildStrategy) DeepCopy() *BuildStrategy {
	if in == nil {
		return nil
	}
	out := new(BuildStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildStrategy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStrategyList) DeepCopyInto(out *BuildStrategyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BuildStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStrategyList.
func (in *BuildStrategyList) DeepCopy() *BuildStrategyList {
	if in == nil {
		return nil
	}
	out := new(BuildStrategyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildStrategyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStrategySpec) DeepCopyInto(out *BuildStrategySpec) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]BuildStrategyVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStrategySpec.
func (in *BuildStrategySpec) DeepCopy() *BuildStrategySpec {
	if in == nil {
		return nil
	}
	out := new(BuildStrategySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStrategyStatus) DeepCopyInto(out *BuildStrategyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStrategyStatus.
func (in *BuildStrategyStatus) DeepCopy() *BuildStrategyStatus {
	if in == nil {
		return nil
	}
	out := new(BuildStrategyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStrategyVolume) DeepCopyInto(out *BuildStrategyVolume) {
	*out = *in
	if in.Overridable != nil {
		in, out := &in.Overridable, &out.Overridable
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStrategyVolume.
func (in *BuildStrategyVolume) DeepCopy() *BuildStrategyVolume {
	if in == nil {
		return nil
	}
	out := new(BuildStrategyVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildVolume) DeepCopyInto(out *BuildVolume) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildVolume.
func (in *BuildVolume) DeepCopy() *BuildVolume {
	if in == nil {
		return nil
	}
	out := new(BuildVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleSourceResult) DeepCopyInto(out *BundleSourceResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleSourceResult.
func (in *BundleSourceResult) DeepCopy() *BundleSourceResult {
	if in == nil {
		return nil
	}
	out := new(BundleSourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneSecret) DeepCopyInto(out *CloneSecret) {
	*out = *in
	if in.CloneSecret != nil {
		in, out := &in.CloneSecret, &out.CloneSecret
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneSecret.
func (in *CloneSecret) DeepCopy() *CloneSecret {
	if in == nil {
		return nil
	}
	out := new(CloneSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBuildStrategy) DeepCopyInto(out *ClusterBuildStrategy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBuildStrategy.
func (in *ClusterBuildStrategy) DeepCopy() *ClusterBuildStrategy {
	if in == nil {
		return nil
	}
	out := new(ClusterBuildStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBuildStrategy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBuildStrategyList) DeepCopyInto(out *ClusterBuildStrategyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBuildStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBuildStrategyList.
func (in *ClusterBuildStrategyList) DeepCopy() *ClusterBuildStrategyList {
	if in == nil {
		return nil
	}
	out := new(ClusterBuildStrategyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBuildStrategyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailedAt) DeepCopyInto(out *FailedAt) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailedAt.
func (in *FailedAt) DeepCopy() *FailedAt {
	if in == nil {
		return nil
	}
	out := new(FailedAt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailureDetails) DeepCopyInto(out *FailureDetails) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(FailedAt)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailureDetails.
func (in *FailureDetails) DeepCopy() *FailureDetails {
	if in == nil {
		return nil
	}
	out := new(FailureDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Git) DeepCopyInto(out *Git) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(string)
		**out = **in
	}
	in.Credentials.DeepCopyInto(&out.Credentials)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Git.
func (in *Git) DeepCopy() *Git {
	if in == nil {
		return nil
	}
	out := new(Git)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSourceResult) DeepCopyInto(out *GitSourceResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSourceResult.
func (in *GitSourceResult) DeepCopy() *GitSourceResult {
	if in == nil {
		return nil
	}
	out := new(GitSourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	if in.PushSecret != nil {
		in, out := &in.PushSecret, &out.PushSecret
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Local) DeepCopyInto(out *Local) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Local.
func (in *Local) DeepCopy() *Local {
	if in == nil {
		return nil
	}
	out := new(Local)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIArtifact) DeepCopyInto(out *OCIArtifact) {
	*out = *in
	if in.Prune != nil {
		in, out := &in.Prune, &out.Prune
		*out = new(PruneOption)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIArtifact.
func (in *OCIArtifact) DeepCopy() *OCIArtifact {
	if in == nil {
		return nil
	}
	out := new(OCIArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectKeyRef) DeepCopyInto(out *ObjectKeyRef) {
	*out = *in
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectKeyRef.
func (in *ObjectKeyRef) DeepCopy() *ObjectKeyRef {
	if in == nil {
		return nil
	}
	out := new(ObjectKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Output) DeepCopyInto(out *Output) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Output.
func (in *Output) DeepCopy() *Output {
	if in == nil {
		return nil
	}
	out := new(Output)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamValue) DeepCopyInto(out *ParamValue) {
	*out = *in
	if in.SingleValue != nil {
		in, out := &in.SingleValue, &out.SingleValue
		*out = new(SingleValue)
		(*in).DeepCopyInto(*out)
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]SingleValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamValue.
func (in *ParamValue) DeepCopy() *ParamValue {
	if in == nil {
		return nil
	}
	out := new(ParamValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(string)
		**out = **in
	}
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferencedBuild) DeepCopyInto(out *ReferencedBuild) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(BuildSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferencedBuild.
func (in *ReferencedBuild) DeepCopy() *ReferencedBuild {
	if in == nil {
		return nil
	}
	out := new(ReferencedBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleValue) DeepCopyInto(out *SingleValue) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ConfigMapValue != nil {
		in, out := &in.ConfigMapValue, &out.ConfigMapValue
		*out = new(ObjectKeyRef)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretValue != nil {
		in, out := &in.SecretValue, &out.SecretValue
		*out = new(ObjectKeyRef)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleValue.
func (in *SingleValue) DeepCopy() *SingleValue {
	if in == nil {
		return nil
	}
	out := new(SingleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	if in.ContextDir != nil {
		in, out := &in.ContextDir, &out.ContextDir
		*out = new(string)
		**out = **in
	}
	if in.OCIArtifact != nil {
		in, out := &in.OCIArtifact, &out.OCIArtifact
		*out = new(OCIArtifact)
		(*in).DeepCopyInto(*out)
	}
	if in.GitSource != nil {
		in, out := &in.GitSource, &out.GitSource
		*out = new(Git)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceResult) DeepCopyInto(out *SourceResult) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitSourceResult)
		**out = **in
	}
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(BundleSourceResult)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceResult.
func (in *SourceResult) DeepCopy() *SourceResult {
	if in == nil {
		return nil
	}
	out := new(SourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]corev1.ContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
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
func (in *Strategy) DeepCopyInto(out *Strategy) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(BuildStrategyKind)
		**out = **in
	}
	if in.APIVersion != nil {
		in, out := &in.APIVersion, &out.APIVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Strategy.
func (in *Strategy) DeepCopy() *Strategy {
	if in == nil {
		return nil
	}
	out := new(Strategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.When != nil {
		in, out := &in.When, &out.When
		*out = make([]TriggerWhen, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerWhen) DeepCopyInto(out *TriggerWhen) {
	*out = *in
	if in.GitHub != nil {
		in, out := &in.GitHub, &out.GitHub
		*out = new(WhenGitHub)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(WhenImage)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectRef != nil {
		in, out := &in.ObjectRef, &out.ObjectRef
		*out = new(WhenObjectRef)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerWhen.
func (in *TriggerWhen) DeepCopy() *TriggerWhen {
	if in == nil {
		return nil
	}
	out := new(TriggerWhen)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhenGitHub) DeepCopyInto(out *WhenGitHub) {
	*out = *in
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]GitHubEventName, len(*in))
		copy(*out, *in)
	}
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhenGitHub.
func (in *WhenGitHub) DeepCopy() *WhenGitHub {
	if in == nil {
		return nil
	}
	out := new(WhenGitHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhenImage) DeepCopyInto(out *WhenImage) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhenImage.
func (in *WhenImage) DeepCopy() *WhenImage {
	if in == nil {
		return nil
	}
	out := new(WhenImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhenObjectRef) DeepCopyInto(out *WhenObjectRef) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhenObjectRef.
func (in *WhenObjectRef) DeepCopy() *WhenObjectRef {
	if in == nil {
		return nil
	}
	out := new(WhenObjectRef)
	in.DeepCopyInto(out)
	return out
}
