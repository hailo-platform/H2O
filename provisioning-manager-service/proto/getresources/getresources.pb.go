// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/provisioning-manager-service/proto/getresources/getresources.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_provisioningmanager_getresources is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/provisioning-manager-service/proto/getresources/getresources.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_provisioningmanager_getresources

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_kernel_provisioningmanager "github.com/HailoOSS/provisioning-manager-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	MachineClass     *string `protobuf:"bytes,1,opt,name=machineClass" json:"machineClass,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

type Response struct {
	Total            *com_HailoOSS_kernel_provisioningmanager.Resource `protobuf:"bytes,1,req,name=total" json:"total,omitempty"`
	Usage            *com_HailoOSS_kernel_provisioningmanager.Resource `protobuf:"bytes,2,req,name=usage" json:"usage,omitempty"`
	XXX_unrecognized []byte                                            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetTotal() *com_HailoOSS_kernel_provisioningmanager.Resource {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *Response) GetUsage() *com_HailoOSS_kernel_provisioningmanager.Resource {
	if m != nil {
		return m.Usage
	}
	return nil
}

func init() {
}
