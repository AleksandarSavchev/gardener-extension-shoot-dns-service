// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	dnsv1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	apis "github.com/gardener/gardener-extension-shoot-dns-service/pkg/apis"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*DNSEntry)(nil), (*apis.DNSEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSEntry_To_apis_DNSEntry(a.(*DNSEntry), b.(*apis.DNSEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apis.DNSEntry)(nil), (*DNSEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apis_DNSEntry_To_v1alpha1_DNSEntry(a.(*apis.DNSEntry), b.(*DNSEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DNSState)(nil), (*apis.DNSState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSState_To_apis_DNSState(a.(*DNSState), b.(*apis.DNSState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apis.DNSState)(nil), (*DNSState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apis_DNSState_To_v1alpha1_DNSState(a.(*apis.DNSState), b.(*DNSState), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_DNSEntry_To_apis_DNSEntry(in *DNSEntry, out *apis.DNSEntry, s conversion.Scope) error {
	out.Name = in.Name
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.Spec = (*dnsv1alpha1.DNSEntrySpec)(unsafe.Pointer(in.Spec))
	return nil
}

// Convert_v1alpha1_DNSEntry_To_apis_DNSEntry is an autogenerated conversion function.
func Convert_v1alpha1_DNSEntry_To_apis_DNSEntry(in *DNSEntry, out *apis.DNSEntry, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSEntry_To_apis_DNSEntry(in, out, s)
}

func autoConvert_apis_DNSEntry_To_v1alpha1_DNSEntry(in *apis.DNSEntry, out *DNSEntry, s conversion.Scope) error {
	out.Name = in.Name
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.Spec = (*dnsv1alpha1.DNSEntrySpec)(unsafe.Pointer(in.Spec))
	return nil
}

// Convert_apis_DNSEntry_To_v1alpha1_DNSEntry is an autogenerated conversion function.
func Convert_apis_DNSEntry_To_v1alpha1_DNSEntry(in *apis.DNSEntry, out *DNSEntry, s conversion.Scope) error {
	return autoConvert_apis_DNSEntry_To_v1alpha1_DNSEntry(in, out, s)
}

func autoConvert_v1alpha1_DNSState_To_apis_DNSState(in *DNSState, out *apis.DNSState, s conversion.Scope) error {
	out.Entries = *(*[]*apis.DNSEntry)(unsafe.Pointer(&in.Entries))
	return nil
}

// Convert_v1alpha1_DNSState_To_apis_DNSState is an autogenerated conversion function.
func Convert_v1alpha1_DNSState_To_apis_DNSState(in *DNSState, out *apis.DNSState, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSState_To_apis_DNSState(in, out, s)
}

func autoConvert_apis_DNSState_To_v1alpha1_DNSState(in *apis.DNSState, out *DNSState, s conversion.Scope) error {
	out.Entries = *(*[]*DNSEntry)(unsafe.Pointer(&in.Entries))
	return nil
}

// Convert_apis_DNSState_To_v1alpha1_DNSState is an autogenerated conversion function.
func Convert_apis_DNSState_To_v1alpha1_DNSState(in *apis.DNSState, out *DNSState, s conversion.Scope) error {
	return autoConvert_apis_DNSState_To_v1alpha1_DNSState(in, out, s)
}
