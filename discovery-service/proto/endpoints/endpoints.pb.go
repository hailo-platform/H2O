// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/discovery-service/proto/endpoints/endpoints.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_discovery_endpoints is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/discovery-service/proto/endpoints/endpoints.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_discovery_endpoints

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Service          *string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
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

type Response struct {
	Endpoints        []*Response_Endpoint `protobuf:"bytes,1,rep,name=endpoints" json:"endpoints,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetEndpoints() []*Response_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type Response_Endpoint struct {
	FqName           *string `protobuf:"bytes,1,req,name=fqName" json:"fqName,omitempty"`
	Version          *uint64 `protobuf:"varint,2,req,name=version" json:"version,omitempty"`
	Mean             *uint32 `protobuf:"varint,3,req,name=mean" json:"mean,omitempty"`
	Upper95          *uint32 `protobuf:"varint,4,req,name=upper95" json:"upper95,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response_Endpoint) Reset()         { *m = Response_Endpoint{} }
func (m *Response_Endpoint) String() string { return proto.CompactTextString(m) }
func (*Response_Endpoint) ProtoMessage()    {}

func (m *Response_Endpoint) GetFqName() string {
	if m != nil && m.FqName != nil {
		return *m.FqName
	}
	return ""
}

func (m *Response_Endpoint) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Response_Endpoint) GetMean() uint32 {
	if m != nil && m.Mean != nil {
		return *m.Mean
	}
	return 0
}

func (m *Response_Endpoint) GetUpper95() uint32 {
	if m != nil && m.Upper95 != nil {
		return *m.Upper95
	}
	return 0
}

func init() {
}