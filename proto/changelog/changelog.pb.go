// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/config-service/proto/changelog/changelog.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_service_config_changelog is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/config-service/proto/changelog/changelog.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_service_config_changelog

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_service_config "github.com/HailoOSS/config-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	// specify an ID to filter the change log
	Id *string `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
	// specify a time range to search between
	RangeStart *int64 `protobuf:"varint,1,opt,name=rangeStart" json:"rangeStart,omitempty"`
	RangeEnd   *int64 `protobuf:"varint,2,opt,name=rangeEnd" json:"rangeEnd,omitempty"`
	// paginate
	LastId           *string `protobuf:"bytes,3,opt,name=lastId" json:"lastId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Request) GetRangeStart() int64 {
	if m != nil && m.RangeStart != nil {
		return *m.RangeStart
	}
	return 0
}

func (m *Request) GetRangeEnd() int64 {
	if m != nil && m.RangeEnd != nil {
		return *m.RangeEnd
	}
	return 0
}

func (m *Request) GetLastId() string {
	if m != nil && m.LastId != nil {
		return *m.LastId
	}
	return ""
}

type Response struct {
	Changes          []*com_HailoOSS_service_config.Change `protobuf:"bytes,1,rep,name=changes" json:"changes,omitempty"`
	Last             *string                               `protobuf:"bytes,2,opt,name=last" json:"last,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetChanges() []*com_HailoOSS_service_config.Change {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *Response) GetLast() string {
	if m != nil && m.Last != nil {
		return *m.Last
	}
	return ""
}

func init() {
}
