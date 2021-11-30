// +build !ignore_autogenerated

// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package resourcebuilder

import (
	"github.com/banzaicloud/operator-tools/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfig) DeepCopyInto(out *ComponentConfig) {
	*out = *in
	in.EnabledComponent.DeepCopyInto(&out.EnabledComponent)
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadMetaOverrides != nil {
		in, out := &in.WorkloadMetaOverrides, &out.WorkloadMetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadOverrides != nil {
		in, out := &in.WorkloadOverrides, &out.WorkloadOverrides
		*out = new(types.PodSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerOverrides != nil {
		in, out := &in.ContainerOverrides, &out.ContainerOverrides
		*out = new(types.ContainerBase)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfig.
func (in *ComponentConfig) DeepCopy() *ComponentConfig {
	if in == nil {
		return nil
	}
	out := new(ComponentConfig)
	in.DeepCopyInto(out)
	return out
}
