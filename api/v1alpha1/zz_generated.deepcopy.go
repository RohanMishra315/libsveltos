//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022. projectsveltos.io. All rights reserved.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRequest) DeepCopyInto(out *AccessRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRequest.
func (in *AccessRequest) DeepCopy() *AccessRequest {
	if in == nil {
		return nil
	}
	out := new(AccessRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRequestList) DeepCopyInto(out *AccessRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRequestList.
func (in *AccessRequestList) DeepCopy() *AccessRequestList {
	if in == nil {
		return nil
	}
	out := new(AccessRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRequestSpec) DeepCopyInto(out *AccessRequestSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRequestSpec.
func (in *AccessRequestSpec) DeepCopy() *AccessRequestSpec {
	if in == nil {
		return nil
	}
	out := new(AccessRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRequestStatus) DeepCopyInto(out *AccessRequestStatus) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRequestStatus.
func (in *AccessRequestStatus) DeepCopy() *AccessRequestStatus {
	if in == nil {
		return nil
	}
	out := new(AccessRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Classifier) DeepCopyInto(out *Classifier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Classifier.
func (in *Classifier) DeepCopy() *Classifier {
	if in == nil {
		return nil
	}
	out := new(Classifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Classifier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierLabel) DeepCopyInto(out *ClassifierLabel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierLabel.
func (in *ClassifierLabel) DeepCopy() *ClassifierLabel {
	if in == nil {
		return nil
	}
	out := new(ClassifierLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierList) DeepCopyInto(out *ClassifierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Classifier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierList.
func (in *ClassifierList) DeepCopy() *ClassifierList {
	if in == nil {
		return nil
	}
	out := new(ClassifierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClassifierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReport) DeepCopyInto(out *ClassifierReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReport.
func (in *ClassifierReport) DeepCopy() *ClassifierReport {
	if in == nil {
		return nil
	}
	out := new(ClassifierReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClassifierReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReportList) DeepCopyInto(out *ClassifierReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClassifierReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReportList.
func (in *ClassifierReportList) DeepCopy() *ClassifierReportList {
	if in == nil {
		return nil
	}
	out := new(ClassifierReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClassifierReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReportSpec) DeepCopyInto(out *ClassifierReportSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReportSpec.
func (in *ClassifierReportSpec) DeepCopy() *ClassifierReportSpec {
	if in == nil {
		return nil
	}
	out := new(ClassifierReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierReportStatus) DeepCopyInto(out *ClassifierReportStatus) {
	*out = *in
	if in.Phase != nil {
		in, out := &in.Phase, &out.Phase
		*out = new(ReportPhase)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierReportStatus.
func (in *ClassifierReportStatus) DeepCopy() *ClassifierReportStatus {
	if in == nil {
		return nil
	}
	out := new(ClassifierReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierSpec) DeepCopyInto(out *ClassifierSpec) {
	*out = *in
	if in.DeployedResourceConstraints != nil {
		in, out := &in.DeployedResourceConstraints, &out.DeployedResourceConstraints
		*out = make([]DeployedResourceConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubernetesVersionConstraints != nil {
		in, out := &in.KubernetesVersionConstraints, &out.KubernetesVersionConstraints
		*out = make([]KubernetesVersionConstraint, len(*in))
		copy(*out, *in)
	}
	if in.ClassifierLabels != nil {
		in, out := &in.ClassifierLabels, &out.ClassifierLabels
		*out = make([]ClassifierLabel, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierSpec.
func (in *ClassifierSpec) DeepCopy() *ClassifierSpec {
	if in == nil {
		return nil
	}
	out := new(ClassifierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassifierStatus) DeepCopyInto(out *ClassifierStatus) {
	*out = *in
	if in.MachingClusterStatuses != nil {
		in, out := &in.MachingClusterStatuses, &out.MachingClusterStatuses
		*out = make([]MachingClusterStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = make([]ClusterInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassifierStatus.
func (in *ClassifierStatus) DeepCopy() *ClassifierStatus {
	if in == nil {
		return nil
	}
	out := new(ClassifierStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.ClusterInfo.DeepCopyInto(&out.ClusterInfo)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NotificationSummaries != nil {
		in, out := &in.NotificationSummaries, &out.NotificationSummaries
		*out = make([]NotificationSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealthCheck) DeepCopyInto(out *ClusterHealthCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealthCheck.
func (in *ClusterHealthCheck) DeepCopy() *ClusterHealthCheck {
	if in == nil {
		return nil
	}
	out := new(ClusterHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterHealthCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealthCheckList) DeepCopyInto(out *ClusterHealthCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterHealthCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealthCheckList.
func (in *ClusterHealthCheckList) DeepCopy() *ClusterHealthCheckList {
	if in == nil {
		return nil
	}
	out := new(ClusterHealthCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterHealthCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealthCheckSpec) DeepCopyInto(out *ClusterHealthCheckSpec) {
	*out = *in
	if in.LivenessChecks != nil {
		in, out := &in.LivenessChecks, &out.LivenessChecks
		*out = make([]LivenessCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]Notification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealthCheckSpec.
func (in *ClusterHealthCheckSpec) DeepCopy() *ClusterHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterHealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealthCheckStatus) DeepCopyInto(out *ClusterHealthCheckStatus) {
	*out = *in
	if in.MatchingClusterRefs != nil {
		in, out := &in.MatchingClusterRefs, &out.MatchingClusterRefs
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ClusterConditions != nil {
		in, out := &in.ClusterConditions, &out.ClusterConditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealthCheckStatus.
func (in *ClusterHealthCheckStatus) DeepCopy() *ClusterHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterHealthCheckStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	out.Cluster = in.Cluster
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfiguration) DeepCopyInto(out *ComponentConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfiguration.
func (in *ComponentConfiguration) DeepCopy() *ComponentConfiguration {
	if in == nil {
		return nil
	}
	out := new(ComponentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
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
func (in *DebuggingConfiguration) DeepCopyInto(out *DebuggingConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebuggingConfiguration.
func (in *DebuggingConfiguration) DeepCopy() *DebuggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(DebuggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DebuggingConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebuggingConfigurationList) DeepCopyInto(out *DebuggingConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DebuggingConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebuggingConfigurationList.
func (in *DebuggingConfigurationList) DeepCopy() *DebuggingConfigurationList {
	if in == nil {
		return nil
	}
	out := new(DebuggingConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DebuggingConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebuggingConfigurationSpec) DeepCopyInto(out *DebuggingConfigurationSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ComponentConfiguration, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebuggingConfigurationSpec.
func (in *DebuggingConfigurationSpec) DeepCopy() *DebuggingConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(DebuggingConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployedResourceConstraint) DeepCopyInto(out *DeployedResourceConstraint) {
	*out = *in
	if in.LabelFilters != nil {
		in, out := &in.LabelFilters, &out.LabelFilters
		*out = make([]LabelFilter, len(*in))
		copy(*out, *in)
	}
	if in.FieldFilters != nil {
		in, out := &in.FieldFilters, &out.FieldFilters
		*out = make([]FieldFilter, len(*in))
		copy(*out, *in)
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(int)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployedResourceConstraint.
func (in *DeployedResourceConstraint) DeepCopy() *DeployedResourceConstraint {
	if in == nil {
		return nil
	}
	out := new(DeployedResourceConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldFilter) DeepCopyInto(out *FieldFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldFilter.
func (in *FieldFilter) DeepCopy() *FieldFilter {
	if in == nil {
		return nil
	}
	out := new(FieldFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmResources) DeepCopyInto(out *HelmResources) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]Resource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmResources.
func (in *HelmResources) DeepCopy() *HelmResources {
	if in == nil {
		return nil
	}
	out := new(HelmResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesVersionConstraint) DeepCopyInto(out *KubernetesVersionConstraint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesVersionConstraint.
func (in *KubernetesVersionConstraint) DeepCopy() *KubernetesVersionConstraint {
	if in == nil {
		return nil
	}
	out := new(KubernetesVersionConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelFilter) DeepCopyInto(out *LabelFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelFilter.
func (in *LabelFilter) DeepCopy() *LabelFilter {
	if in == nil {
		return nil
	}
	out := new(LabelFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LivenessCheck) DeepCopyInto(out *LivenessCheck) {
	*out = *in
	if in.LivenessSourceRef != nil {
		in, out := &in.LivenessSourceRef, &out.LivenessSourceRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LivenessCheck.
func (in *LivenessCheck) DeepCopy() *LivenessCheck {
	if in == nil {
		return nil
	}
	out := new(LivenessCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachingClusterStatus) DeepCopyInto(out *MachingClusterStatus) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.ManagedLabels != nil {
		in, out := &in.ManagedLabels, &out.ManagedLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UnManagedLabels != nil {
		in, out := &in.UnManagedLabels, &out.UnManagedLabels
		*out = make([]UnManagedLabel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachingClusterStatus.
func (in *MachingClusterStatus) DeepCopy() *MachingClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MachingClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notification) DeepCopyInto(out *Notification) {
	*out = *in
	if in.NotificationRef != nil {
		in, out := &in.NotificationRef, &out.NotificationRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notification.
func (in *Notification) DeepCopy() *Notification {
	if in == nil {
		return nil
	}
	out := new(Notification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationSummary) DeepCopyInto(out *NotificationSummary) {
	*out = *in
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationSummary.
func (in *NotificationSummary) DeepCopy() *NotificationSummary {
	if in == nil {
		return nil
	}
	out := new(NotificationSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRef) DeepCopyInto(out *PolicyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRef.
func (in *PolicyRef) DeepCopy() *PolicyRef {
	if in == nil {
		return nil
	}
	out := new(PolicyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceHash) DeepCopyInto(out *ResourceHash) {
	*out = *in
	out.Resource = in.Resource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceHash.
func (in *ResourceHash) DeepCopy() *ResourceHash {
	if in == nil {
		return nil
	}
	out := new(ResourceHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSummary) DeepCopyInto(out *ResourceSummary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSummary.
func (in *ResourceSummary) DeepCopy() *ResourceSummary {
	if in == nil {
		return nil
	}
	out := new(ResourceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSummary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSummaryList) DeepCopyInto(out *ResourceSummaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSummaryList.
func (in *ResourceSummaryList) DeepCopy() *ResourceSummaryList {
	if in == nil {
		return nil
	}
	out := new(ResourceSummaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSummaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSummarySpec) DeepCopyInto(out *ResourceSummarySpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]Resource, len(*in))
		copy(*out, *in)
	}
	if in.ChartResources != nil {
		in, out := &in.ChartResources, &out.ChartResources
		*out = make([]HelmResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSummarySpec.
func (in *ResourceSummarySpec) DeepCopy() *ResourceSummarySpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSummarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSummaryStatus) DeepCopyInto(out *ResourceSummaryStatus) {
	*out = *in
	if in.ResourceHashes != nil {
		in, out := &in.ResourceHashes, &out.ResourceHashes
		*out = make([]ResourceHash, len(*in))
		copy(*out, *in)
	}
	if in.HelmResourceHashes != nil {
		in, out := &in.HelmResourceHashes, &out.HelmResourceHashes
		*out = make([]ResourceHash, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSummaryStatus.
func (in *ResourceSummaryStatus) DeepCopy() *ResourceSummaryStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceSummaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleRequest) DeepCopyInto(out *RoleRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleRequest.
func (in *RoleRequest) DeepCopy() *RoleRequest {
	if in == nil {
		return nil
	}
	out := new(RoleRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleRequestList) DeepCopyInto(out *RoleRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleRequestList.
func (in *RoleRequestList) DeepCopy() *RoleRequestList {
	if in == nil {
		return nil
	}
	out := new(RoleRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleRequestSpec) DeepCopyInto(out *RoleRequestSpec) {
	*out = *in
	if in.RoleRefs != nil {
		in, out := &in.RoleRefs, &out.RoleRefs
		*out = make([]PolicyRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleRequestSpec.
func (in *RoleRequestSpec) DeepCopy() *RoleRequestSpec {
	if in == nil {
		return nil
	}
	out := new(RoleRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleRequestStatus) DeepCopyInto(out *RoleRequestStatus) {
	*out = *in
	if in.MatchingClusterRefs != nil {
		in, out := &in.MatchingClusterRefs, &out.MatchingClusterRefs
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = make([]ClusterInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleRequestStatus.
func (in *RoleRequestStatus) DeepCopy() *RoleRequestStatus {
	if in == nil {
		return nil
	}
	out := new(RoleRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SveltosCluster) DeepCopyInto(out *SveltosCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SveltosCluster.
func (in *SveltosCluster) DeepCopy() *SveltosCluster {
	if in == nil {
		return nil
	}
	out := new(SveltosCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SveltosCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SveltosClusterList) DeepCopyInto(out *SveltosClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SveltosCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SveltosClusterList.
func (in *SveltosClusterList) DeepCopy() *SveltosClusterList {
	if in == nil {
		return nil
	}
	out := new(SveltosClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SveltosClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SveltosClusterSpec) DeepCopyInto(out *SveltosClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SveltosClusterSpec.
func (in *SveltosClusterSpec) DeepCopy() *SveltosClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SveltosClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SveltosClusterStatus) DeepCopyInto(out *SveltosClusterStatus) {
	*out = *in
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SveltosClusterStatus.
func (in *SveltosClusterStatus) DeepCopy() *SveltosClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SveltosClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnManagedLabel) DeepCopyInto(out *UnManagedLabel) {
	*out = *in
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnManagedLabel.
func (in *UnManagedLabel) DeepCopy() *UnManagedLabel {
	if in == nil {
		return nil
	}
	out := new(UnManagedLabel)
	in.DeepCopyInto(out)
	return out
}
