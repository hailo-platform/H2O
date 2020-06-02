// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/binding-service/proto/createrule/createrule.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_binding_createrule is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/binding-service/proto/createrule/createrule.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_binding_createrule

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_kernel_binding "github.com/HailoOSS/binding-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Rule             *com_HailoOSS_kernel_binding.BindingRule `protobuf:"bytes,1,req,name=rule" json:"rule,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetRule() *com_HailoOSS_kernel_binding.BindingRule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type Response struct {
	Ok               *bool  `protobuf:"varint,1,req,name=ok" json:"ok,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func init() {
}
