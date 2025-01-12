// The MIT License
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: temporal/server/api/persistence/v1/history_tree.proto

package persistence

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// branch column
type HistoryTreeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchInfo *HistoryBranch `protobuf:"bytes,1,opt,name=branch_info,json=branchInfo,proto3" json:"branch_info,omitempty"`
	// For fork operation to prevent race condition of leaking event data when forking branches fail. Also can be used for clean up leaked data.
	ForkTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=fork_time,json=forkTime,proto3" json:"fork_time,omitempty"`
	// For lookup back to workflow during debugging, also background cleanup when fork operation cannot finish self cleanup due to crash.
	Info string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	// Deprecating branch token in favor of branch info.
	//
	// Deprecated: Marked as deprecated in temporal/server/api/persistence/v1/history_tree.proto.
	BranchToken []byte `protobuf:"bytes,4,opt,name=branch_token,json=branchToken,proto3" json:"branch_token,omitempty"`
}

func (x *HistoryTreeInfo) Reset() {
	*x = HistoryTreeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryTreeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryTreeInfo) ProtoMessage() {}

func (x *HistoryTreeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryTreeInfo.ProtoReflect.Descriptor instead.
func (*HistoryTreeInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_history_tree_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryTreeInfo) GetBranchInfo() *HistoryBranch {
	if x != nil {
		return x.BranchInfo
	}
	return nil
}

func (x *HistoryTreeInfo) GetForkTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ForkTime
	}
	return nil
}

func (x *HistoryTreeInfo) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

// Deprecated: Marked as deprecated in temporal/server/api/persistence/v1/history_tree.proto.
func (x *HistoryTreeInfo) GetBranchToken() []byte {
	if x != nil {
		return x.BranchToken
	}
	return nil
}

// For history persistence to serialize/deserialize branch details.
type HistoryBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreeId    string                `protobuf:"bytes,1,opt,name=tree_id,json=treeId,proto3" json:"tree_id,omitempty"`
	BranchId  string                `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Ancestors []*HistoryBranchRange `protobuf:"bytes,3,rep,name=ancestors,proto3" json:"ancestors,omitempty"`
}

func (x *HistoryBranch) Reset() {
	*x = HistoryBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryBranch) ProtoMessage() {}

func (x *HistoryBranch) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryBranch.ProtoReflect.Descriptor instead.
func (*HistoryBranch) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_history_tree_proto_rawDescGZIP(), []int{1}
}

func (x *HistoryBranch) GetTreeId() string {
	if x != nil {
		return x.TreeId
	}
	return ""
}

func (x *HistoryBranch) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *HistoryBranch) GetAncestors() []*HistoryBranchRange {
	if x != nil {
		return x.Ancestors
	}
	return nil
}

// HistoryBranchRange represents a piece of range for a branch.
type HistoryBranchRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BranchId of original branch forked from.
	BranchId string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	// Beginning node for the range, inclusive.
	BeginNodeId int64 `protobuf:"varint,2,opt,name=begin_node_id,json=beginNodeId,proto3" json:"begin_node_id,omitempty"`
	// Ending node for the range, exclusive.
	EndNodeId int64 `protobuf:"varint,3,opt,name=end_node_id,json=endNodeId,proto3" json:"end_node_id,omitempty"`
}

func (x *HistoryBranchRange) Reset() {
	*x = HistoryBranchRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryBranchRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryBranchRange) ProtoMessage() {}

func (x *HistoryBranchRange) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryBranchRange.ProtoReflect.Descriptor instead.
func (*HistoryBranchRange) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_history_tree_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryBranchRange) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *HistoryBranchRange) GetBeginNodeId() int64 {
	if x != nil {
		return x.BeginNodeId
	}
	return 0
}

func (x *HistoryBranchRange) GetEndNodeId() int64 {
	if x != nil {
		return x.EndNodeId
	}
	return 0
}

var File_temporal_server_api_persistence_v1_history_tree_proto protoreflect.FileDescriptor

var file_temporal_server_api_persistence_v1_history_tree_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x72, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a,
	0x0f, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x52, 0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x66, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x25, 0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72,
	0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x65,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x54, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x75, 0x0a, 0x12, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_persistence_v1_history_tree_proto_rawDescOnce sync.Once
	file_temporal_server_api_persistence_v1_history_tree_proto_rawDescData = file_temporal_server_api_persistence_v1_history_tree_proto_rawDesc
)

func file_temporal_server_api_persistence_v1_history_tree_proto_rawDescGZIP() []byte {
	file_temporal_server_api_persistence_v1_history_tree_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_persistence_v1_history_tree_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_persistence_v1_history_tree_proto_rawDescData)
	})
	return file_temporal_server_api_persistence_v1_history_tree_proto_rawDescData
}

var file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporal_server_api_persistence_v1_history_tree_proto_goTypes = []interface{}{
	(*HistoryTreeInfo)(nil),       // 0: temporal.server.api.persistence.v1.HistoryTreeInfo
	(*HistoryBranch)(nil),         // 1: temporal.server.api.persistence.v1.HistoryBranch
	(*HistoryBranchRange)(nil),    // 2: temporal.server.api.persistence.v1.HistoryBranchRange
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_temporal_server_api_persistence_v1_history_tree_proto_depIdxs = []int32{
	1, // 0: temporal.server.api.persistence.v1.HistoryTreeInfo.branch_info:type_name -> temporal.server.api.persistence.v1.HistoryBranch
	3, // 1: temporal.server.api.persistence.v1.HistoryTreeInfo.fork_time:type_name -> google.protobuf.Timestamp
	2, // 2: temporal.server.api.persistence.v1.HistoryBranch.ancestors:type_name -> temporal.server.api.persistence.v1.HistoryBranchRange
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_temporal_server_api_persistence_v1_history_tree_proto_init() }
func file_temporal_server_api_persistence_v1_history_tree_proto_init() {
	if File_temporal_server_api_persistence_v1_history_tree_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryTreeInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryBranch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryBranchRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_server_api_persistence_v1_history_tree_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_persistence_v1_history_tree_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_persistence_v1_history_tree_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_persistence_v1_history_tree_proto_msgTypes,
	}.Build()
	File_temporal_server_api_persistence_v1_history_tree_proto = out.File
	file_temporal_server_api_persistence_v1_history_tree_proto_rawDesc = nil
	file_temporal_server_api_persistence_v1_history_tree_proto_goTypes = nil
	file_temporal_server_api_persistence_v1_history_tree_proto_depIdxs = nil
}
