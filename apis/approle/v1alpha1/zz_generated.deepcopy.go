//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendLogin) DeepCopyInto(out *AuthBackendLogin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendLogin.
func (in *AuthBackendLogin) DeepCopy() *AuthBackendLogin {
	if in == nil {
		return nil
	}
	out := new(AuthBackendLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendLogin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendLoginList) DeepCopyInto(out *AuthBackendLoginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackendLogin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendLoginList.
func (in *AuthBackendLoginList) DeepCopy() *AuthBackendLoginList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendLoginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendLoginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendLoginObservation) DeepCopyInto(out *AuthBackendLoginObservation) {
	*out = *in
	if in.Accessor != nil {
		in, out := &in.Accessor, &out.Accessor
		*out = new(string)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ClientToken != nil {
		in, out := &in.ClientToken, &out.ClientToken
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LeaseDuration != nil {
		in, out := &in.LeaseDuration, &out.LeaseDuration
		*out = new(float64)
		**out = **in
	}
	if in.LeaseStarted != nil {
		in, out := &in.LeaseStarted, &out.LeaseStarted
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Renewable != nil {
		in, out := &in.Renewable, &out.Renewable
		*out = new(bool)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendLoginObservation.
func (in *AuthBackendLoginObservation) DeepCopy() *AuthBackendLoginObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendLoginObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendLoginParameters) DeepCopyInto(out *AuthBackendLoginParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendLoginParameters.
func (in *AuthBackendLoginParameters) DeepCopy() *AuthBackendLoginParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendLoginParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendLoginSpec) DeepCopyInto(out *AuthBackendLoginSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendLoginSpec.
func (in *AuthBackendLoginSpec) DeepCopy() *AuthBackendLoginSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendLoginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendLoginStatus) DeepCopyInto(out *AuthBackendLoginStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendLoginStatus.
func (in *AuthBackendLoginStatus) DeepCopy() *AuthBackendLoginStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendLoginStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRole) DeepCopyInto(out *AuthBackendRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRole.
func (in *AuthBackendRole) DeepCopy() *AuthBackendRole {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleList) DeepCopyInto(out *AuthBackendRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackendRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleList.
func (in *AuthBackendRoleList) DeepCopy() *AuthBackendRoleList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleObservation) DeepCopyInto(out *AuthBackendRoleObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.BindSecretID != nil {
		in, out := &in.BindSecretID, &out.BindSecretID
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.SecretIDBoundCidrs != nil {
		in, out := &in.SecretIDBoundCidrs, &out.SecretIDBoundCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecretIDNumUses != nil {
		in, out := &in.SecretIDNumUses, &out.SecretIDNumUses
		*out = new(float64)
		**out = **in
	}
	if in.SecretIDTTL != nil {
		in, out := &in.SecretIDTTL, &out.SecretIDTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenBoundCidrs != nil {
		in, out := &in.TokenBoundCidrs, &out.TokenBoundCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenExplicitMaxTTL != nil {
		in, out := &in.TokenExplicitMaxTTL, &out.TokenExplicitMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenMaxTTL != nil {
		in, out := &in.TokenMaxTTL, &out.TokenMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenNoDefaultPolicy != nil {
		in, out := &in.TokenNoDefaultPolicy, &out.TokenNoDefaultPolicy
		*out = new(bool)
		**out = **in
	}
	if in.TokenNumUses != nil {
		in, out := &in.TokenNumUses, &out.TokenNumUses
		*out = new(float64)
		**out = **in
	}
	if in.TokenPeriod != nil {
		in, out := &in.TokenPeriod, &out.TokenPeriod
		*out = new(float64)
		**out = **in
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenTTL != nil {
		in, out := &in.TokenTTL, &out.TokenTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleObservation.
func (in *AuthBackendRoleObservation) DeepCopy() *AuthBackendRoleObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleParameters) DeepCopyInto(out *AuthBackendRoleParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.BindSecretID != nil {
		in, out := &in.BindSecretID, &out.BindSecretID
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.SecretIDBoundCidrs != nil {
		in, out := &in.SecretIDBoundCidrs, &out.SecretIDBoundCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecretIDNumUses != nil {
		in, out := &in.SecretIDNumUses, &out.SecretIDNumUses
		*out = new(float64)
		**out = **in
	}
	if in.SecretIDTTL != nil {
		in, out := &in.SecretIDTTL, &out.SecretIDTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenBoundCidrs != nil {
		in, out := &in.TokenBoundCidrs, &out.TokenBoundCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenExplicitMaxTTL != nil {
		in, out := &in.TokenExplicitMaxTTL, &out.TokenExplicitMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenMaxTTL != nil {
		in, out := &in.TokenMaxTTL, &out.TokenMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenNoDefaultPolicy != nil {
		in, out := &in.TokenNoDefaultPolicy, &out.TokenNoDefaultPolicy
		*out = new(bool)
		**out = **in
	}
	if in.TokenNumUses != nil {
		in, out := &in.TokenNumUses, &out.TokenNumUses
		*out = new(float64)
		**out = **in
	}
	if in.TokenPeriod != nil {
		in, out := &in.TokenPeriod, &out.TokenPeriod
		*out = new(float64)
		**out = **in
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenTTL != nil {
		in, out := &in.TokenTTL, &out.TokenTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleParameters.
func (in *AuthBackendRoleParameters) DeepCopy() *AuthBackendRoleParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSecretID) DeepCopyInto(out *AuthBackendRoleSecretID) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSecretID.
func (in *AuthBackendRoleSecretID) DeepCopy() *AuthBackendRoleSecretID {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSecretID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendRoleSecretID) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSecretIDList) DeepCopyInto(out *AuthBackendRoleSecretIDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackendRoleSecretID, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSecretIDList.
func (in *AuthBackendRoleSecretIDList) DeepCopy() *AuthBackendRoleSecretIDList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSecretIDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendRoleSecretIDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSecretIDObservation) DeepCopyInto(out *AuthBackendRoleSecretIDObservation) {
	*out = *in
	if in.Accessor != nil {
		in, out := &in.Accessor, &out.Accessor
		*out = new(string)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.WithWrappedAccessor != nil {
		in, out := &in.WithWrappedAccessor, &out.WithWrappedAccessor
		*out = new(bool)
		**out = **in
	}
	if in.WrappingAccessor != nil {
		in, out := &in.WrappingAccessor, &out.WrappingAccessor
		*out = new(string)
		**out = **in
	}
	if in.WrappingTTL != nil {
		in, out := &in.WrappingTTL, &out.WrappingTTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSecretIDObservation.
func (in *AuthBackendRoleSecretIDObservation) DeepCopy() *AuthBackendRoleSecretIDObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSecretIDObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSecretIDParameters) DeepCopyInto(out *AuthBackendRoleSecretIDParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.SecretIDSecretRef != nil {
		in, out := &in.SecretIDSecretRef, &out.SecretIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.WithWrappedAccessor != nil {
		in, out := &in.WithWrappedAccessor, &out.WithWrappedAccessor
		*out = new(bool)
		**out = **in
	}
	if in.WrappingTTL != nil {
		in, out := &in.WrappingTTL, &out.WrappingTTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSecretIDParameters.
func (in *AuthBackendRoleSecretIDParameters) DeepCopy() *AuthBackendRoleSecretIDParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSecretIDParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSecretIDSpec) DeepCopyInto(out *AuthBackendRoleSecretIDSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSecretIDSpec.
func (in *AuthBackendRoleSecretIDSpec) DeepCopy() *AuthBackendRoleSecretIDSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSecretIDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSecretIDStatus) DeepCopyInto(out *AuthBackendRoleSecretIDStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSecretIDStatus.
func (in *AuthBackendRoleSecretIDStatus) DeepCopy() *AuthBackendRoleSecretIDStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSecretIDStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSpec) DeepCopyInto(out *AuthBackendRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSpec.
func (in *AuthBackendRoleSpec) DeepCopy() *AuthBackendRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleStatus) DeepCopyInto(out *AuthBackendRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleStatus.
func (in *AuthBackendRoleStatus) DeepCopy() *AuthBackendRoleStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleStatus)
	in.DeepCopyInto(out)
	return out
}