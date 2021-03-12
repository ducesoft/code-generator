// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	example3io "github.com/ducesoft/code-generator/examples/apiserver/apis/example3.io"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*TestType)(nil), (*example3io.TestType)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TestType_To_example3io_TestType(a.(*TestType), b.(*example3io.TestType), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example3io.TestType)(nil), (*TestType)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example3io_TestType_To_v1_TestType(a.(*example3io.TestType), b.(*TestType), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TestTypeList)(nil), (*example3io.TestTypeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TestTypeList_To_example3io_TestTypeList(a.(*TestTypeList), b.(*example3io.TestTypeList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example3io.TestTypeList)(nil), (*TestTypeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example3io_TestTypeList_To_v1_TestTypeList(a.(*example3io.TestTypeList), b.(*TestTypeList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TestTypeStatus)(nil), (*example3io.TestTypeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TestTypeStatus_To_example3io_TestTypeStatus(a.(*TestTypeStatus), b.(*example3io.TestTypeStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example3io.TestTypeStatus)(nil), (*TestTypeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example3io_TestTypeStatus_To_v1_TestTypeStatus(a.(*example3io.TestTypeStatus), b.(*TestTypeStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_TestType_To_example3io_TestType(in *TestType, out *example3io.TestType, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_TestTypeStatus_To_example3io_TestTypeStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TestType_To_example3io_TestType is an autogenerated conversion function.
func Convert_v1_TestType_To_example3io_TestType(in *TestType, out *example3io.TestType, s conversion.Scope) error {
	return autoConvert_v1_TestType_To_example3io_TestType(in, out, s)
}

func autoConvert_example3io_TestType_To_v1_TestType(in *example3io.TestType, out *TestType, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_example3io_TestTypeStatus_To_v1_TestTypeStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_example3io_TestType_To_v1_TestType is an autogenerated conversion function.
func Convert_example3io_TestType_To_v1_TestType(in *example3io.TestType, out *TestType, s conversion.Scope) error {
	return autoConvert_example3io_TestType_To_v1_TestType(in, out, s)
}

func autoConvert_v1_TestTypeList_To_example3io_TestTypeList(in *TestTypeList, out *example3io.TestTypeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]example3io.TestType)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_TestTypeList_To_example3io_TestTypeList is an autogenerated conversion function.
func Convert_v1_TestTypeList_To_example3io_TestTypeList(in *TestTypeList, out *example3io.TestTypeList, s conversion.Scope) error {
	return autoConvert_v1_TestTypeList_To_example3io_TestTypeList(in, out, s)
}

func autoConvert_example3io_TestTypeList_To_v1_TestTypeList(in *example3io.TestTypeList, out *TestTypeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]TestType)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_example3io_TestTypeList_To_v1_TestTypeList is an autogenerated conversion function.
func Convert_example3io_TestTypeList_To_v1_TestTypeList(in *example3io.TestTypeList, out *TestTypeList, s conversion.Scope) error {
	return autoConvert_example3io_TestTypeList_To_v1_TestTypeList(in, out, s)
}

func autoConvert_v1_TestTypeStatus_To_example3io_TestTypeStatus(in *TestTypeStatus, out *example3io.TestTypeStatus, s conversion.Scope) error {
	out.Blah = in.Blah
	return nil
}

// Convert_v1_TestTypeStatus_To_example3io_TestTypeStatus is an autogenerated conversion function.
func Convert_v1_TestTypeStatus_To_example3io_TestTypeStatus(in *TestTypeStatus, out *example3io.TestTypeStatus, s conversion.Scope) error {
	return autoConvert_v1_TestTypeStatus_To_example3io_TestTypeStatus(in, out, s)
}

func autoConvert_example3io_TestTypeStatus_To_v1_TestTypeStatus(in *example3io.TestTypeStatus, out *TestTypeStatus, s conversion.Scope) error {
	out.Blah = in.Blah
	return nil
}

// Convert_example3io_TestTypeStatus_To_v1_TestTypeStatus is an autogenerated conversion function.
func Convert_example3io_TestTypeStatus_To_v1_TestTypeStatus(in *example3io.TestTypeStatus, out *TestTypeStatus, s conversion.Scope) error {
	return autoConvert_example3io_TestTypeStatus_To_v1_TestTypeStatus(in, out, s)
}
