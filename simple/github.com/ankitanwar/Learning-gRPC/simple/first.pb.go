// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: first.proto

package simple

import (
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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age              int32  `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	FirstName        string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName         string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	IsProfileVerifid bool   `protobuf:"varint,5,opt,name=is_profile_verifid,json=isProfileVerifid,proto3" json:"is_profile_verifid,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetIsProfileVerifid() bool {
	if x != nil {
		return x.IsProfileVerifid
	}
	return false
}

var File_first_proto protoreflect.FileDescriptor

var file_first_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x64, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x6b, 0x69, 0x74, 0x61,
	0x6e, 0x77, 0x61, 0x72, 0x2f, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x52,
	0x50, 0x43, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_first_proto_rawDescOnce sync.Once
	file_first_proto_rawDescData = file_first_proto_rawDesc
)

func file_first_proto_rawDescGZIP() []byte {
	file_first_proto_rawDescOnce.Do(func() {
		file_first_proto_rawDescData = protoimpl.X.CompressGZIP(file_first_proto_rawDescData)
	})
	return file_first_proto_rawDescData
}

var file_first_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_first_proto_goTypes = []interface{}{
	(*Person)(nil), // 0: right.Person
}
var file_first_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_first_proto_init() }
func file_first_proto_init() {
	if File_first_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_first_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
			RawDescriptor: file_first_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_first_proto_goTypes,
		DependencyIndexes: file_first_proto_depIdxs,
		MessageInfos:      file_first_proto_msgTypes,
	}.Build()
	File_first_proto = out.File
	file_first_proto_rawDesc = nil
	file_first_proto_goTypes = nil
	file_first_proto_depIdxs = nil
}
