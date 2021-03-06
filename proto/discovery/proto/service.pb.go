// Code generated by protoc-gen-go.
// source: github.com/hailocab/h2/proto/discovery/proto/service.proto
// DO NOT EDIT!

/*
Package com_hailocab_kernel_discovery is a generated protocol buffer package.

It is generated from these files:
	github.com/hailocab/h2/proto/discovery/proto/service.proto

It has these top-level messages:
	Service
*/
package com_hailocab_kernel_discovery

import proto "github.com/hailocab/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Service struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Description      *string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Version          *uint64 `protobuf:"varint,3,req,name=version" json:"version,omitempty"`
	Source           *string `protobuf:"bytes,4,req,name=source" json:"source,omitempty"`
	OwnerEmail       *string `protobuf:"bytes,5,req,name=ownerEmail" json:"ownerEmail,omitempty"`
	OwnerMobile      *string `protobuf:"bytes,6,req,name=ownerMobile" json:"ownerMobile,omitempty"`
	OwnerTeam        *string `protobuf:"bytes,7,opt,name=ownerTeam" json:"ownerTeam,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}

func (m *Service) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Service) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Service) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *Service) GetOwnerEmail() string {
	if m != nil && m.OwnerEmail != nil {
		return *m.OwnerEmail
	}
	return ""
}

func (m *Service) GetOwnerMobile() string {
	if m != nil && m.OwnerMobile != nil {
		return *m.OwnerMobile
	}
	return ""
}

func (m *Service) GetOwnerTeam() string {
	if m != nil && m.OwnerTeam != nil {
		return *m.OwnerTeam
	}
	return ""
}

func init() {
}
