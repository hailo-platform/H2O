// Code generated by protoc-gen-go.
// source: github.com/hailo-platform/H2O/binding-service/proto/setupservice/setupservice.proto
// DO NOT EDIT!

/*
Package com_hailo-platform/H2O_kernel_binding_setupservice is a generated protocol buffer package.

It is generated from these files:
	github.com/hailo-platform/H2O/binding-service/proto/setupservice/setupservice.proto

It has these top-level messages:
	Request
	Response
*/
package com_hailo-platform/H2O_kernel_binding_setupservice

import proto "github.com/hailo-platform/H2O/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Service          *string `protobuf:"bytes,1,req,name=service" json:"service,omitempty"`
	Queue            *string `protobuf:"bytes,2,req,name=queue" json:"queue,omitempty"`
	Azname           *string `protobuf:"bytes,3,req,name=azname" json:"azname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetService() string {
	if m != nil && m.Service != nil {
		return *m.Service
	}
	return ""
}

func (m *Request) GetQueue() string {
	if m != nil && m.Queue != nil {
		return *m.Queue
	}
	return ""
}

func (m *Request) GetAzname() string {
	if m != nil && m.Azname != nil {
		return *m.Azname
	}
	return ""
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
