// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.1
// source: wso2/discovery/api/Resource.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Resource config model
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path             string              `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Methods          []string            `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
	Summary          string              `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	Description      string              `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	ProductionUrls   []*Endpoint         `protobuf:"bytes,6,rep,name=productionUrls,proto3" json:"productionUrls,omitempty"`
	SandboxUrls      []*Endpoint         `protobuf:"bytes,7,rep,name=sandboxUrls,proto3" json:"sandboxUrls,omitempty"`
	Security         map[string]string   `protobuf:"bytes,8,rep,name=security,proto3" json:"security,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	VendorExtensible map[string]*any.Any `protobuf:"bytes,9,rep,name=vendorExtensible,proto3" json:"vendorExtensible,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Consumes         []string            `protobuf:"bytes,10,rep,name=consumes,proto3" json:"consumes,omitempty"`
	Schemes          []string            `protobuf:"bytes,11,rep,name=schemes,proto3" json:"schemes,omitempty"`
	Tags             []string            `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_Resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_Resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_Resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Resource) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Resource) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Resource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Resource) GetProductionUrls() []*Endpoint {
	if x != nil {
		return x.ProductionUrls
	}
	return nil
}

func (x *Resource) GetSandboxUrls() []*Endpoint {
	if x != nil {
		return x.SandboxUrls
	}
	return nil
}

func (x *Resource) GetSecurity() map[string]string {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Resource) GetVendorExtensible() map[string]*any.Any {
	if x != nil {
		return x.VendorExtensible
	}
	return nil
}

func (x *Resource) GetConsumes() []string {
	if x != nil {
		return x.Consumes
	}
	return nil
}

func (x *Resource) GetSchemes() []string {
	if x != nil {
		return x.Schemes
	}
	return nil
}

func (x *Resource) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_wso2_discovery_api_Resource_proto protoreflect.FileDescriptor

var file_wso2_discovery_api_Resource_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x05, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72,
	0x6c, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x55, 0x72, 0x6c,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x55, 0x72,
	0x6c, 0x73, 0x12, 0x46, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x5e, 0x0a, 0x10, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x59, 0x0a, 0x15, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x70, 0x0a, 0x1e,
	0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0d,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_api_Resource_proto_rawDescOnce sync.Once
	file_wso2_discovery_api_Resource_proto_rawDescData = file_wso2_discovery_api_Resource_proto_rawDesc
)

func file_wso2_discovery_api_Resource_proto_rawDescGZIP() []byte {
	file_wso2_discovery_api_Resource_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_api_Resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_api_Resource_proto_rawDescData)
	})
	return file_wso2_discovery_api_Resource_proto_rawDescData
}

var file_wso2_discovery_api_Resource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wso2_discovery_api_Resource_proto_goTypes = []interface{}{
	(*Resource)(nil), // 0: wso2.discovery.api.Resource
	nil,              // 1: wso2.discovery.api.Resource.SecurityEntry
	nil,              // 2: wso2.discovery.api.Resource.VendorExtensibleEntry
	(*Endpoint)(nil), // 3: wso2.discovery.api.Endpoint
	(*any.Any)(nil),  // 4: google.protobuf.Any
}
var file_wso2_discovery_api_Resource_proto_depIdxs = []int32{
	3, // 0: wso2.discovery.api.Resource.productionUrls:type_name -> wso2.discovery.api.Endpoint
	3, // 1: wso2.discovery.api.Resource.sandboxUrls:type_name -> wso2.discovery.api.Endpoint
	1, // 2: wso2.discovery.api.Resource.security:type_name -> wso2.discovery.api.Resource.SecurityEntry
	2, // 3: wso2.discovery.api.Resource.vendorExtensible:type_name -> wso2.discovery.api.Resource.VendorExtensibleEntry
	4, // 4: wso2.discovery.api.Resource.VendorExtensibleEntry.value:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_wso2_discovery_api_Resource_proto_init() }
func file_wso2_discovery_api_Resource_proto_init() {
	if File_wso2_discovery_api_Resource_proto != nil {
		return
	}
	file_wso2_discovery_api_Endpoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_api_Resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
			RawDescriptor: file_wso2_discovery_api_Resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_api_Resource_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_api_Resource_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_api_Resource_proto_msgTypes,
	}.Build()
	File_wso2_discovery_api_Resource_proto = out.File
	file_wso2_discovery_api_Resource_proto_rawDesc = nil
	file_wso2_discovery_api_Resource_proto_goTypes = nil
	file_wso2_discovery_api_Resource_proto_depIdxs = nil
}