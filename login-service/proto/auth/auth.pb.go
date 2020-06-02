// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/login-service/proto/auth/auth.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_service_login_auth is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/login-service/proto/auth/auth.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_service_login_auth

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_service_login "github.com/HailoOSS/login-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Mech       *string `protobuf:"bytes,1,req,name=mech" json:"mech,omitempty"`
	DeviceType *string `protobuf:"bytes,2,req,name=deviceType" json:"deviceType,omitempty"`
	Password   *string `protobuf:"bytes,3,req,name=password" json:"password,omitempty"`
	Username   *string `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	// the plan is to get rid of these when we port
	Email       *string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Phone       *string `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
	NewPassword *string `protobuf:"bytes,7,opt,name=newPassword" json:"newPassword,omitempty"`
	ApiToken    *string `protobuf:"bytes,8,opt,name=apiToken" json:"apiToken,omitempty"`
	// application is for h2 login and namespaces our user pool
	Application *string `protobuf:"bytes,9,opt,name=application" json:"application,omitempty"`
	// meta data is optional meta data for h2 logins to attach to the login record, things like IP etc.
	Meta []*com_HailoOSS_service_login.KeyValue `protobuf:"bytes,10,rep,name=meta" json:"meta,omitempty"`
	// If true will authenticate the user with the current session
	NoExpire         *bool   `protobuf:"varint,11,opt,name=noExpire" json:"noExpire,omitempty"`
	OauthToken       *string `protobuf:"bytes,12,opt,name=oauthToken" json:"oauthToken,omitempty"`
	Provider         *string `protobuf:"bytes,13,opt,name=provider" json:"provider,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetMech() string {
	if m != nil && m.Mech != nil {
		return *m.Mech
	}
	return ""
}

func (m *Request) GetDeviceType() string {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *Request) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Request) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *Request) GetPhone() string {
	if m != nil && m.Phone != nil {
		return *m.Phone
	}
	return ""
}

func (m *Request) GetNewPassword() string {
	if m != nil && m.NewPassword != nil {
		return *m.NewPassword
	}
	return ""
}

func (m *Request) GetApiToken() string {
	if m != nil && m.ApiToken != nil {
		return *m.ApiToken
	}
	return ""
}

func (m *Request) GetApplication() string {
	if m != nil && m.Application != nil {
		return *m.Application
	}
	return ""
}

func (m *Request) GetMeta() []*com_HailoOSS_service_login.KeyValue {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Request) GetNoExpire() bool {
	if m != nil && m.NoExpire != nil {
		return *m.NoExpire
	}
	return false
}

func (m *Request) GetOauthToken() string {
	if m != nil && m.OauthToken != nil {
		return *m.OauthToken
	}
	return ""
}

func (m *Request) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

type Response struct {
	SessId           *string `protobuf:"bytes,1,req,name=sessId" json:"sessId,omitempty"`
	Token            *string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetSessId() string {
	if m != nil && m.SessId != nil {
		return *m.SessId
	}
	return ""
}

func (m *Response) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func init() {
}
