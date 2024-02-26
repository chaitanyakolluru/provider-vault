//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

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
func (in *AuthBackendRoleInitParameters) DeepCopyInto(out *AuthBackendRoleInitParameters) {
	*out = *in
	if in.AllowedCommonNames != nil {
		in, out := &in.AllowedCommonNames, &out.AllowedCommonNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedDNSSans != nil {
		in, out := &in.AllowedDNSSans, &out.AllowedDNSSans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedEmailSans != nil {
		in, out := &in.AllowedEmailSans, &out.AllowedEmailSans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedNames != nil {
		in, out := &in.AllowedNames, &out.AllowedNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrganizationUnits != nil {
		in, out := &in.AllowedOrganizationUnits, &out.AllowedOrganizationUnits
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrganizationalUnits != nil {
		in, out := &in.AllowedOrganizationalUnits, &out.AllowedOrganizationalUnits
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedURISans != nil {
		in, out := &in.AllowedURISans, &out.AllowedURISans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.OcspCACertificates != nil {
		in, out := &in.OcspCACertificates, &out.OcspCACertificates
		*out = new(string)
		**out = **in
	}
	if in.OcspEnabled != nil {
		in, out := &in.OcspEnabled, &out.OcspEnabled
		*out = new(bool)
		**out = **in
	}
	if in.OcspFailOpen != nil {
		in, out := &in.OcspFailOpen, &out.OcspFailOpen
		*out = new(bool)
		**out = **in
	}
	if in.OcspQueryAllServers != nil {
		in, out := &in.OcspQueryAllServers, &out.OcspQueryAllServers
		*out = new(bool)
		**out = **in
	}
	if in.OcspServersOverride != nil {
		in, out := &in.OcspServersOverride, &out.OcspServersOverride
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RequiredExtensions != nil {
		in, out := &in.RequiredExtensions, &out.RequiredExtensions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleInitParameters.
func (in *AuthBackendRoleInitParameters) DeepCopy() *AuthBackendRoleInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleInitParameters)
	in.DeepCopyInto(out)
	return out
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
	if in.AllowedCommonNames != nil {
		in, out := &in.AllowedCommonNames, &out.AllowedCommonNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedDNSSans != nil {
		in, out := &in.AllowedDNSSans, &out.AllowedDNSSans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedEmailSans != nil {
		in, out := &in.AllowedEmailSans, &out.AllowedEmailSans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedNames != nil {
		in, out := &in.AllowedNames, &out.AllowedNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrganizationUnits != nil {
		in, out := &in.AllowedOrganizationUnits, &out.AllowedOrganizationUnits
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrganizationalUnits != nil {
		in, out := &in.AllowedOrganizationalUnits, &out.AllowedOrganizationalUnits
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedURISans != nil {
		in, out := &in.AllowedURISans, &out.AllowedURISans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.OcspCACertificates != nil {
		in, out := &in.OcspCACertificates, &out.OcspCACertificates
		*out = new(string)
		**out = **in
	}
	if in.OcspEnabled != nil {
		in, out := &in.OcspEnabled, &out.OcspEnabled
		*out = new(bool)
		**out = **in
	}
	if in.OcspFailOpen != nil {
		in, out := &in.OcspFailOpen, &out.OcspFailOpen
		*out = new(bool)
		**out = **in
	}
	if in.OcspQueryAllServers != nil {
		in, out := &in.OcspQueryAllServers, &out.OcspQueryAllServers
		*out = new(bool)
		**out = **in
	}
	if in.OcspServersOverride != nil {
		in, out := &in.OcspServersOverride, &out.OcspServersOverride
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RequiredExtensions != nil {
		in, out := &in.RequiredExtensions, &out.RequiredExtensions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.AllowedCommonNames != nil {
		in, out := &in.AllowedCommonNames, &out.AllowedCommonNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedDNSSans != nil {
		in, out := &in.AllowedDNSSans, &out.AllowedDNSSans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedEmailSans != nil {
		in, out := &in.AllowedEmailSans, &out.AllowedEmailSans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedNames != nil {
		in, out := &in.AllowedNames, &out.AllowedNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrganizationUnits != nil {
		in, out := &in.AllowedOrganizationUnits, &out.AllowedOrganizationUnits
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrganizationalUnits != nil {
		in, out := &in.AllowedOrganizationalUnits, &out.AllowedOrganizationalUnits
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedURISans != nil {
		in, out := &in.AllowedURISans, &out.AllowedURISans
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.OcspCACertificates != nil {
		in, out := &in.OcspCACertificates, &out.OcspCACertificates
		*out = new(string)
		**out = **in
	}
	if in.OcspEnabled != nil {
		in, out := &in.OcspEnabled, &out.OcspEnabled
		*out = new(bool)
		**out = **in
	}
	if in.OcspFailOpen != nil {
		in, out := &in.OcspFailOpen, &out.OcspFailOpen
		*out = new(bool)
		**out = **in
	}
	if in.OcspQueryAllServers != nil {
		in, out := &in.OcspQueryAllServers, &out.OcspQueryAllServers
		*out = new(bool)
		**out = **in
	}
	if in.OcspServersOverride != nil {
		in, out := &in.OcspServersOverride, &out.OcspServersOverride
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RequiredExtensions != nil {
		in, out := &in.RequiredExtensions, &out.RequiredExtensions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
func (in *AuthBackendRoleSpec) DeepCopyInto(out *AuthBackendRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
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
