// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/annotations.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         72295728,
	Name:          "google.api.http",
	Tag:           "bytes,72295728,opt,name=http",
	Filename:      "google/api/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Http)
}

func init() { proto.RegisterFile("google/api/annotations.proto", fileDescriptorAnnotations) }

var fileDescriptorAnnotations = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xc8, 0xea, 0x25, 0x16, 0x64,
	0x4a, 0x89, 0x22, 0xa9, 0xcc, 0x28, 0x29, 0x29, 0x80, 0x28, 0x91, 0x52, 0x80, 0x0a, 0x83, 0x79,
	0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x10, 0x15,
	0x56, 0xde, 0x5c, 0x2c, 0x20, 0xf5, 0x42, 0x72, 0x7a, 0x50, 0xd3, 0x60, 0x4a, 0xf5, 0x7c, 0x53,
	0x4b, 0x32, 0xf2, 0x53, 0xfc, 0x0b, 0xc0, 0x56, 0x4a, 0x6c, 0x38, 0xb5, 0x47, 0x49, 0x81, 0x51,
	0x83, 0xdb, 0x48, 0x44, 0x0f, 0x61, 0xad, 0x9e, 0x47, 0x49, 0x49, 0x41, 0x50, 0x69, 0x4e, 0x6a,
	0x10, 0xd8, 0x10, 0xa7, 0x84, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50,
	0x8e, 0xf1, 0xc7, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c,
	0x00, 0x89, 0x3d, 0x96, 0x63, 0xe4, 0xe2, 0x4b, 0xce, 0xcf, 0x45, 0x32, 0xd0, 0x49, 0xc0, 0x11,
	0xe1, 0xcd, 0x00, 0x90, 0x4b, 0x02, 0x18, 0xa3, 0x98, 0x13, 0x0b, 0x32, 0x17, 0x31, 0xb1, 0xb8,
	0x3b, 0x06, 0x78, 0x26, 0xb1, 0x81, 0x9d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x92, 0x1b,
	0xe6, 0x45, 0x1a, 0x01, 0x00, 0x00,
}
