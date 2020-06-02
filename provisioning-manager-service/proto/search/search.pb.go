// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/provisioning-manager-service/proto/search/search.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_service_provisioningmanager_search is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/provisioning-manager-service/proto/search/search.proto

It has these top-level messages:
	Request
	Result
	Response
*/
package com_HailoOSS_service_provisioningmanager_search

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	ServiceName      *string `protobuf:"bytes,1,opt,name=serviceName" json:"serviceName,omitempty"`
	MachineClass     *string `protobuf:"bytes,2,opt,name=machineClass" json:"machineClass,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Request) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

type Result struct {
	ServiceName      *string `protobuf:"bytes,1,req,name=serviceName" json:"serviceName,omitempty"`
	ServiceVersion   *uint64 `protobuf:"varint,2,req,name=serviceVersion" json:"serviceVersion,omitempty"`
	MachineClass     *string `protobuf:"bytes,3,req,name=machineClass" json:"machineClass,omitempty"`
	NoFileSoftLimit  *uint64 `protobuf:"varint,4,opt,name=noFileSoftLimit" json:"noFileSoftLimit,omitempty"`
	NoFileHardLimit  *uint64 `protobuf:"varint,5,opt,name=noFileHardLimit" json:"noFileHardLimit,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}

func (m *Result) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Result) GetServiceVersion() uint64 {
	if m != nil && m.ServiceVersion != nil {
		return *m.ServiceVersion
	}
	return 0
}

func (m *Result) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

func (m *Result) GetNoFileSoftLimit() uint64 {
	if m != nil && m.NoFileSoftLimit != nil {
		return *m.NoFileSoftLimit
	}
	return 0
}

func (m *Result) GetNoFileHardLimit() uint64 {
	if m != nil && m.NoFileHardLimit != nil {
		return *m.NoFileHardLimit
	}
	return 0
}

type Response struct {
	Results          []*Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
}
