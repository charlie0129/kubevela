//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The KubeVela Authors.

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

package v1beta1

import (
	"github.com/kubevela/workflow/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/common"
	core_oam_devv1alpha1 "github.com/oam-dev/kubevela/apis/core.oam.dev/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppPolicy) DeepCopyInto(out *AppPolicy) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppPolicy.
func (in *AppPolicy) DeepCopy() *AppPolicy {
	if in == nil {
		return nil
	}
	out := new(AppPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationRevision) DeepCopyInto(out *ApplicationRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationRevision.
func (in *ApplicationRevision) DeepCopy() *ApplicationRevision {
	if in == nil {
		return nil
	}
	out := new(ApplicationRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationRevisionCompressibleFields) DeepCopyInto(out *ApplicationRevisionCompressibleFields) {
	*out = *in
	in.Application.DeepCopyInto(&out.Application)
	if in.ComponentDefinitions != nil {
		in, out := &in.ComponentDefinitions, &out.ComponentDefinitions
		*out = make(map[string]*ComponentDefinition, len(*in))
		for key, val := range *in {
			var outVal *ComponentDefinition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ComponentDefinition)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.WorkloadDefinitions != nil {
		in, out := &in.WorkloadDefinitions, &out.WorkloadDefinitions
		*out = make(map[string]WorkloadDefinition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.TraitDefinitions != nil {
		in, out := &in.TraitDefinitions, &out.TraitDefinitions
		*out = make(map[string]*TraitDefinition, len(*in))
		for key, val := range *in {
			var outVal *TraitDefinition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(TraitDefinition)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ScopeDefinitions != nil {
		in, out := &in.ScopeDefinitions, &out.ScopeDefinitions
		*out = make(map[string]ScopeDefinition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.PolicyDefinitions != nil {
		in, out := &in.PolicyDefinitions, &out.PolicyDefinitions
		*out = make(map[string]PolicyDefinition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.WorkflowStepDefinitions != nil {
		in, out := &in.WorkflowStepDefinitions, &out.WorkflowStepDefinitions
		*out = make(map[string]*WorkflowStepDefinition, len(*in))
		for key, val := range *in {
			var outVal *WorkflowStepDefinition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(WorkflowStepDefinition)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ScopeGVK != nil {
		in, out := &in.ScopeGVK, &out.ScopeGVK
		*out = make(map[string]v1.GroupVersionKind, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make(map[string]core_oam_devv1alpha1.Policy, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = new(v1alpha1.Workflow)
		(*in).DeepCopyInto(*out)
	}
	if in.ReferredObjects != nil {
		in, out := &in.ReferredObjects, &out.ReferredObjects
		*out = make([]common.ReferredObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationRevisionCompressibleFields.
func (in *ApplicationRevisionCompressibleFields) DeepCopy() *ApplicationRevisionCompressibleFields {
	if in == nil {
		return nil
	}
	out := new(ApplicationRevisionCompressibleFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationRevisionCompression) DeepCopyInto(out *ApplicationRevisionCompression) {
	*out = *in
	out.CompressedText = in.CompressedText
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationRevisionCompression.
func (in *ApplicationRevisionCompression) DeepCopy() *ApplicationRevisionCompression {
	if in == nil {
		return nil
	}
	out := new(ApplicationRevisionCompression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationRevisionList) DeepCopyInto(out *ApplicationRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationRevisionList.
func (in *ApplicationRevisionList) DeepCopy() *ApplicationRevisionList {
	if in == nil {
		return nil
	}
	out := new(ApplicationRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationRevisionSpec) DeepCopyInto(out *ApplicationRevisionSpec) {
	*out = *in
	in.ApplicationRevisionCompressibleFields.DeepCopyInto(&out.ApplicationRevisionCompressibleFields)
	out.Compression = in.Compression
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationRevisionSpec.
func (in *ApplicationRevisionSpec) DeepCopy() *ApplicationRevisionSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationRevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationRevisionStatus) DeepCopyInto(out *ApplicationRevisionStatus) {
	*out = *in
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = new(common.WorkflowStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkflowContext != nil {
		in, out := &in.WorkflowContext, &out.WorkflowContext
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationRevisionStatus.
func (in *ApplicationRevisionStatus) DeepCopy() *ApplicationRevisionStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationRevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]common.ApplicationComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]AppPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = new(Workflow)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinition) DeepCopyInto(out *ComponentDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinition.
func (in *ComponentDefinition) DeepCopy() *ComponentDefinition {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinitionList) DeepCopyInto(out *ComponentDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComponentDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinitionList.
func (in *ComponentDefinitionList) DeepCopy() *ComponentDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinitionSpec) DeepCopyInto(out *ComponentDefinitionSpec) {
	*out = *in
	out.Workload = in.Workload
	if in.ChildResourceKinds != nil {
		in, out := &in.ChildResourceKinds, &out.ChildResourceKinds
		*out = make([]common.ChildResourceKind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(common.Status)
		**out = **in
	}
	if in.Schematic != nil {
		in, out := &in.Schematic, &out.Schematic
		*out = new(common.Schematic)
		(*in).DeepCopyInto(*out)
	}
	if in.Extension != nil {
		in, out := &in.Extension, &out.Extension
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinitionSpec.
func (in *ComponentDefinitionSpec) DeepCopy() *ComponentDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinitionStatus) DeepCopyInto(out *ComponentDefinitionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(common.Revision)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinitionStatus.
func (in *ComponentDefinitionStatus) DeepCopy() *ComponentDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionRevision) DeepCopyInto(out *DefinitionRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionRevision.
func (in *DefinitionRevision) DeepCopy() *DefinitionRevision {
	if in == nil {
		return nil
	}
	out := new(DefinitionRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefinitionRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionRevisionList) DeepCopyInto(out *DefinitionRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefinitionRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionRevisionList.
func (in *DefinitionRevisionList) DeepCopy() *DefinitionRevisionList {
	if in == nil {
		return nil
	}
	out := new(DefinitionRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefinitionRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionRevisionSpec) DeepCopyInto(out *DefinitionRevisionSpec) {
	*out = *in
	in.ComponentDefinition.DeepCopyInto(&out.ComponentDefinition)
	in.TraitDefinition.DeepCopyInto(&out.TraitDefinition)
	in.PolicyDefinition.DeepCopyInto(&out.PolicyDefinition)
	in.WorkflowStepDefinition.DeepCopyInto(&out.WorkflowStepDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionRevisionSpec.
func (in *DefinitionRevisionSpec) DeepCopy() *DefinitionRevisionSpec {
	if in == nil {
		return nil
	}
	out := new(DefinitionRevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedResource) DeepCopyInto(out *ManagedResource) {
	*out = *in
	out.ClusterObjectReference = in.ClusterObjectReference
	out.OAMObjectReference = in.OAMObjectReference
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedResource.
func (in *ManagedResource) DeepCopy() *ManagedResource {
	if in == nil {
		return nil
	}
	out := new(ManagedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinition) DeepCopyInto(out *PolicyDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinition.
func (in *PolicyDefinition) DeepCopy() *PolicyDefinition {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionList) DeepCopyInto(out *PolicyDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionList.
func (in *PolicyDefinitionList) DeepCopy() *PolicyDefinitionList {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionSpec) DeepCopyInto(out *PolicyDefinitionSpec) {
	*out = *in
	out.Reference = in.Reference
	if in.Schematic != nil {
		in, out := &in.Schematic, &out.Schematic
		*out = new(common.Schematic)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionSpec.
func (in *PolicyDefinitionSpec) DeepCopy() *PolicyDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionStatus) DeepCopyInto(out *PolicyDefinitionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(common.Revision)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionStatus.
func (in *PolicyDefinitionStatus) DeepCopy() *PolicyDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTracker) DeepCopyInto(out *ResourceTracker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTracker.
func (in *ResourceTracker) DeepCopy() *ResourceTracker {
	if in == nil {
		return nil
	}
	out := new(ResourceTracker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceTracker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTrackerCompression) DeepCopyInto(out *ResourceTrackerCompression) {
	*out = *in
	out.CompressedText = in.CompressedText
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTrackerCompression.
func (in *ResourceTrackerCompression) DeepCopy() *ResourceTrackerCompression {
	if in == nil {
		return nil
	}
	out := new(ResourceTrackerCompression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTrackerList) DeepCopyInto(out *ResourceTrackerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceTracker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTrackerList.
func (in *ResourceTrackerList) DeepCopy() *ResourceTrackerList {
	if in == nil {
		return nil
	}
	out := new(ResourceTrackerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceTrackerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTrackerSpec) DeepCopyInto(out *ResourceTrackerSpec) {
	*out = *in
	if in.ManagedResources != nil {
		in, out := &in.ManagedResources, &out.ManagedResources
		*out = make([]ManagedResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Compression = in.Compression
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTrackerSpec.
func (in *ResourceTrackerSpec) DeepCopy() *ResourceTrackerSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceTrackerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTrackerStatus) DeepCopyInto(out *ResourceTrackerStatus) {
	*out = *in
	if in.TrackedResources != nil {
		in, out := &in.TrackedResources, &out.TrackedResources
		*out = make([]common.ClusterObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTrackerStatus.
func (in *ResourceTrackerStatus) DeepCopy() *ResourceTrackerStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceTrackerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeDefinition) DeepCopyInto(out *ScopeDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeDefinition.
func (in *ScopeDefinition) DeepCopy() *ScopeDefinition {
	if in == nil {
		return nil
	}
	out := new(ScopeDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScopeDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeDefinitionList) DeepCopyInto(out *ScopeDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScopeDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeDefinitionList.
func (in *ScopeDefinitionList) DeepCopy() *ScopeDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ScopeDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScopeDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeDefinitionSpec) DeepCopyInto(out *ScopeDefinitionSpec) {
	*out = *in
	out.Reference = in.Reference
	if in.Extension != nil {
		in, out := &in.Extension, &out.Extension
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeDefinitionSpec.
func (in *ScopeDefinitionSpec) DeepCopy() *ScopeDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ScopeDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraitDefinition) DeepCopyInto(out *TraitDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraitDefinition.
func (in *TraitDefinition) DeepCopy() *TraitDefinition {
	if in == nil {
		return nil
	}
	out := new(TraitDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TraitDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraitDefinitionList) DeepCopyInto(out *TraitDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TraitDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraitDefinitionList.
func (in *TraitDefinitionList) DeepCopy() *TraitDefinitionList {
	if in == nil {
		return nil
	}
	out := new(TraitDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TraitDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraitDefinitionSpec) DeepCopyInto(out *TraitDefinitionSpec) {
	*out = *in
	out.Reference = in.Reference
	if in.AppliesToWorkloads != nil {
		in, out := &in.AppliesToWorkloads, &out.AppliesToWorkloads
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConflictsWith != nil {
		in, out := &in.ConflictsWith, &out.ConflictsWith
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Schematic != nil {
		in, out := &in.Schematic, &out.Schematic
		*out = new(common.Schematic)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(common.Status)
		**out = **in
	}
	if in.Extension != nil {
		in, out := &in.Extension, &out.Extension
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraitDefinitionSpec.
func (in *TraitDefinitionSpec) DeepCopy() *TraitDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(TraitDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraitDefinitionStatus) DeepCopyInto(out *TraitDefinitionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(common.Revision)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraitDefinitionStatus.
func (in *TraitDefinitionStatus) DeepCopy() *TraitDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(TraitDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workflow) DeepCopyInto(out *Workflow) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(v1alpha1.WorkflowExecuteMode)
		**out = **in
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]v1alpha1.WorkflowStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workflow.
func (in *Workflow) DeepCopy() *Workflow {
	if in == nil {
		return nil
	}
	out := new(Workflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStepDefinition) DeepCopyInto(out *WorkflowStepDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStepDefinition.
func (in *WorkflowStepDefinition) DeepCopy() *WorkflowStepDefinition {
	if in == nil {
		return nil
	}
	out := new(WorkflowStepDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowStepDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStepDefinitionList) DeepCopyInto(out *WorkflowStepDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkflowStepDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStepDefinitionList.
func (in *WorkflowStepDefinitionList) DeepCopy() *WorkflowStepDefinitionList {
	if in == nil {
		return nil
	}
	out := new(WorkflowStepDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowStepDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStepDefinitionSpec) DeepCopyInto(out *WorkflowStepDefinitionSpec) {
	*out = *in
	out.Reference = in.Reference
	if in.Schematic != nil {
		in, out := &in.Schematic, &out.Schematic
		*out = new(common.Schematic)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStepDefinitionSpec.
func (in *WorkflowStepDefinitionSpec) DeepCopy() *WorkflowStepDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowStepDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStepDefinitionStatus) DeepCopyInto(out *WorkflowStepDefinitionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(common.Revision)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStepDefinitionStatus.
func (in *WorkflowStepDefinitionStatus) DeepCopy() *WorkflowStepDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(WorkflowStepDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadDefinition) DeepCopyInto(out *WorkloadDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadDefinition.
func (in *WorkloadDefinition) DeepCopy() *WorkloadDefinition {
	if in == nil {
		return nil
	}
	out := new(WorkloadDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadDefinitionList) DeepCopyInto(out *WorkloadDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadDefinitionList.
func (in *WorkloadDefinitionList) DeepCopy() *WorkloadDefinitionList {
	if in == nil {
		return nil
	}
	out := new(WorkloadDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadDefinitionSpec) DeepCopyInto(out *WorkloadDefinitionSpec) {
	*out = *in
	out.Reference = in.Reference
	if in.ChildResourceKinds != nil {
		in, out := &in.ChildResourceKinds, &out.ChildResourceKinds
		*out = make([]common.ChildResourceKind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(common.Status)
		**out = **in
	}
	if in.Schematic != nil {
		in, out := &in.Schematic, &out.Schematic
		*out = new(common.Schematic)
		(*in).DeepCopyInto(*out)
	}
	if in.Extension != nil {
		in, out := &in.Extension, &out.Extension
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadDefinitionSpec.
func (in *WorkloadDefinitionSpec) DeepCopy() *WorkloadDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadDefinitionStatus) DeepCopyInto(out *WorkloadDefinitionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadDefinitionStatus.
func (in *WorkloadDefinitionStatus) DeepCopy() *WorkloadDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}
