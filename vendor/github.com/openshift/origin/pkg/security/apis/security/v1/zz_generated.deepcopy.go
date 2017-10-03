// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_FSGroupStrategyOptions, InType: reflect.TypeOf(&FSGroupStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_IDRange, InType: reflect.TypeOf(&IDRange{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicyReview, InType: reflect.TypeOf(&PodSecurityPolicyReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicyReviewSpec, InType: reflect.TypeOf(&PodSecurityPolicyReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicyReviewStatus, InType: reflect.TypeOf(&PodSecurityPolicyReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicySelfSubjectReview, InType: reflect.TypeOf(&PodSecurityPolicySelfSubjectReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicySelfSubjectReviewSpec, InType: reflect.TypeOf(&PodSecurityPolicySelfSubjectReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicySubjectReview, InType: reflect.TypeOf(&PodSecurityPolicySubjectReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicySubjectReviewSpec, InType: reflect.TypeOf(&PodSecurityPolicySubjectReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSecurityPolicySubjectReviewStatus, InType: reflect.TypeOf(&PodSecurityPolicySubjectReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RunAsUserStrategyOptions, InType: reflect.TypeOf(&RunAsUserStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SELinuxContextStrategyOptions, InType: reflect.TypeOf(&SELinuxContextStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SecurityContextConstraints, InType: reflect.TypeOf(&SecurityContextConstraints{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SecurityContextConstraintsList, InType: reflect.TypeOf(&SecurityContextConstraintsList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ServiceAccountPodSecurityPolicyReviewStatus, InType: reflect.TypeOf(&ServiceAccountPodSecurityPolicyReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SupplementalGroupsStrategyOptions, InType: reflect.TypeOf(&SupplementalGroupsStrategyOptions{})},
	)
}

func DeepCopy_v1_FSGroupStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FSGroupStrategyOptions)
		out := out.(*FSGroupStrategyOptions)
		*out = *in
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_IDRange(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IDRange)
		out := out.(*IDRange)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicyReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicyReview)
		out := out.(*PodSecurityPolicyReview)
		*out = *in
		if err := DeepCopy_v1_PodSecurityPolicyReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_PodSecurityPolicyReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicyReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicyReviewSpec)
		out := out.(*PodSecurityPolicyReviewSpec)
		*out = *in
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.ServiceAccountNames != nil {
			in, out := &in.ServiceAccountNames, &out.ServiceAccountNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicyReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicyReviewStatus)
		out := out.(*PodSecurityPolicyReviewStatus)
		*out = *in
		if in.AllowedServiceAccounts != nil {
			in, out := &in.AllowedServiceAccounts, &out.AllowedServiceAccounts
			*out = make([]ServiceAccountPodSecurityPolicyReviewStatus, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ServiceAccountPodSecurityPolicyReviewStatus(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicySelfSubjectReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicySelfSubjectReview)
		out := out.(*PodSecurityPolicySelfSubjectReview)
		*out = *in
		if err := DeepCopy_v1_PodSecurityPolicySelfSubjectReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicySelfSubjectReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicySelfSubjectReviewSpec)
		out := out.(*PodSecurityPolicySelfSubjectReviewSpec)
		*out = *in
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicySubjectReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicySubjectReview)
		out := out.(*PodSecurityPolicySubjectReview)
		*out = *in
		if err := DeepCopy_v1_PodSecurityPolicySubjectReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicySubjectReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicySubjectReviewSpec)
		out := out.(*PodSecurityPolicySubjectReviewSpec)
		*out = *in
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_PodSecurityPolicySubjectReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicySubjectReviewStatus)
		out := out.(*PodSecurityPolicySubjectReviewStatus)
		*out = *in
		if in.AllowedBy != nil {
			in, out := &in.AllowedBy, &out.AllowedBy
			*out = new(api_v1.ObjectReference)
			**out = **in
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_RunAsUserStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RunAsUserStrategyOptions)
		out := out.(*RunAsUserStrategyOptions)
		*out = *in
		if in.UID != nil {
			in, out := &in.UID, &out.UID
			*out = new(int64)
			**out = **in
		}
		if in.UIDRangeMin != nil {
			in, out := &in.UIDRangeMin, &out.UIDRangeMin
			*out = new(int64)
			**out = **in
		}
		if in.UIDRangeMax != nil {
			in, out := &in.UIDRangeMax, &out.UIDRangeMax
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_SELinuxContextStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SELinuxContextStrategyOptions)
		out := out.(*SELinuxContextStrategyOptions)
		*out = *in
		if in.SELinuxOptions != nil {
			in, out := &in.SELinuxOptions, &out.SELinuxOptions
			*out = new(api_v1.SELinuxOptions)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_SecurityContextConstraints(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SecurityContextConstraints)
		out := out.(*SecurityContextConstraints)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if in.Priority != nil {
			in, out := &in.Priority, &out.Priority
			*out = new(int32)
			**out = **in
		}
		if in.DefaultAddCapabilities != nil {
			in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
			*out = make([]api_v1.Capability, len(*in))
			copy(*out, *in)
		}
		if in.RequiredDropCapabilities != nil {
			in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
			*out = make([]api_v1.Capability, len(*in))
			copy(*out, *in)
		}
		if in.AllowedCapabilities != nil {
			in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
			*out = make([]api_v1.Capability, len(*in))
			copy(*out, *in)
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make([]FSType, len(*in))
			copy(*out, *in)
		}
		if err := DeepCopy_v1_SELinuxContextStrategyOptions(&in.SELinuxContext, &out.SELinuxContext, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, c); err != nil {
			return err
		}
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.SeccompProfiles != nil {
			in, out := &in.SeccompProfiles, &out.SeccompProfiles
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_SecurityContextConstraintsList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SecurityContextConstraintsList)
		out := out.(*SecurityContextConstraintsList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]SecurityContextConstraints, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_SecurityContextConstraints(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_ServiceAccountPodSecurityPolicyReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceAccountPodSecurityPolicyReviewStatus)
		out := out.(*ServiceAccountPodSecurityPolicyReviewStatus)
		*out = *in
		if err := DeepCopy_v1_PodSecurityPolicySubjectReviewStatus(&in.PodSecurityPolicySubjectReviewStatus, &out.PodSecurityPolicySubjectReviewStatus, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_SupplementalGroupsStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SupplementalGroupsStrategyOptions)
		out := out.(*SupplementalGroupsStrategyOptions)
		*out = *in
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
