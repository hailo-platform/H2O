// Code generated by protoc-gen-go.
// source: github.com/hailocab/h2/proto/login/proto/keyvalue.proto
// DO NOT EDIT!

/*
Package com_hailocab_service_login is a generated protocol buffer package.

It is generated from these files:
	github.com/hailocab/h2/proto/login/proto/keyvalue.proto

It has these top-level messages:
	KeyValue
*/
package com_hailocab_service_login

import proto "github.com/hailocab/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type KeyValue struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}

func (m *KeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
}
