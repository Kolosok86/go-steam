// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: uifontfile_format.proto

package protobuf

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

type CUIFontFilePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FontFileName     *string `protobuf:"bytes,1,opt,name=font_file_name,json=fontFileName" json:"font_file_name,omitempty"`
	OpentypeFontData []byte  `protobuf:"bytes,2,opt,name=opentype_font_data,json=opentypeFontData" json:"opentype_font_data,omitempty"`
}

func (x *CUIFontFilePB) Reset() {
	*x = CUIFontFilePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uifontfile_format_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUIFontFilePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUIFontFilePB) ProtoMessage() {}

func (x *CUIFontFilePB) ProtoReflect() protoreflect.Message {
	mi := &file_uifontfile_format_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUIFontFilePB.ProtoReflect.Descriptor instead.
func (*CUIFontFilePB) Descriptor() ([]byte, []int) {
	return file_uifontfile_format_proto_rawDescGZIP(), []int{0}
}

func (x *CUIFontFilePB) GetFontFileName() string {
	if x != nil && x.FontFileName != nil {
		return *x.FontFileName
	}
	return ""
}

func (x *CUIFontFilePB) GetOpentypeFontData() []byte {
	if x != nil {
		return x.OpentypeFontData
	}
	return nil
}

type CUIFontFilePackagePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageVersion     *uint32                                        `protobuf:"varint,1,req,name=package_version,json=packageVersion" json:"package_version,omitempty"`
	EncryptedFontFiles []*CUIFontFilePackagePB_CUIEncryptedFontFilePB `protobuf:"bytes,2,rep,name=encrypted_font_files,json=encryptedFontFiles" json:"encrypted_font_files,omitempty"`
}

func (x *CUIFontFilePackagePB) Reset() {
	*x = CUIFontFilePackagePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uifontfile_format_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUIFontFilePackagePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUIFontFilePackagePB) ProtoMessage() {}

func (x *CUIFontFilePackagePB) ProtoReflect() protoreflect.Message {
	mi := &file_uifontfile_format_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUIFontFilePackagePB.ProtoReflect.Descriptor instead.
func (*CUIFontFilePackagePB) Descriptor() ([]byte, []int) {
	return file_uifontfile_format_proto_rawDescGZIP(), []int{1}
}

func (x *CUIFontFilePackagePB) GetPackageVersion() uint32 {
	if x != nil && x.PackageVersion != nil {
		return *x.PackageVersion
	}
	return 0
}

func (x *CUIFontFilePackagePB) GetEncryptedFontFiles() []*CUIFontFilePackagePB_CUIEncryptedFontFilePB {
	if x != nil {
		return x.EncryptedFontFiles
	}
	return nil
}

type CUIFontFilePackagePB_CUIEncryptedFontFilePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedContents []byte `protobuf:"bytes,1,opt,name=encrypted_contents,json=encryptedContents" json:"encrypted_contents,omitempty"`
}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) Reset() {
	*x = CUIFontFilePackagePB_CUIEncryptedFontFilePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uifontfile_format_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUIFontFilePackagePB_CUIEncryptedFontFilePB) ProtoMessage() {}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) ProtoReflect() protoreflect.Message {
	mi := &file_uifontfile_format_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUIFontFilePackagePB_CUIEncryptedFontFilePB.ProtoReflect.Descriptor instead.
func (*CUIFontFilePackagePB_CUIEncryptedFontFilePB) Descriptor() ([]byte, []int) {
	return file_uifontfile_format_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CUIFontFilePackagePB_CUIEncryptedFontFilePB) GetEncryptedContents() []byte {
	if x != nil {
		return x.EncryptedContents
	}
	return nil
}

var File_uifontfile_format_proto protoreflect.FileDescriptor

var file_uifontfile_format_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x69, 0x66, 0x6f, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0d, 0x43, 0x55, 0x49,
	0x46, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x42, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x6f,
	0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x6f, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x46, 0x6f, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0xe8,
	0x01, 0x0a, 0x14, 0x43, 0x55, 0x49, 0x46, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x42, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x5e, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x6f,
	0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x43, 0x55, 0x49, 0x46, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x42, 0x2e, 0x43, 0x55, 0x49, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x46, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x42, 0x52, 0x12, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x1a, 0x47, 0x0a, 0x16, 0x43, 0x55, 0x49, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x46, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x42, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
}

var (
	file_uifontfile_format_proto_rawDescOnce sync.Once
	file_uifontfile_format_proto_rawDescData = file_uifontfile_format_proto_rawDesc
)

func file_uifontfile_format_proto_rawDescGZIP() []byte {
	file_uifontfile_format_proto_rawDescOnce.Do(func() {
		file_uifontfile_format_proto_rawDescData = protoimpl.X.CompressGZIP(file_uifontfile_format_proto_rawDescData)
	})
	return file_uifontfile_format_proto_rawDescData
}

var file_uifontfile_format_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_uifontfile_format_proto_goTypes = []interface{}{
	(*CUIFontFilePB)(nil),                               // 0: CUIFontFilePB
	(*CUIFontFilePackagePB)(nil),                        // 1: CUIFontFilePackagePB
	(*CUIFontFilePackagePB_CUIEncryptedFontFilePB)(nil), // 2: CUIFontFilePackagePB.CUIEncryptedFontFilePB
}
var file_uifontfile_format_proto_depIdxs = []int32{
	2, // 0: CUIFontFilePackagePB.encrypted_font_files:type_name -> CUIFontFilePackagePB.CUIEncryptedFontFilePB
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_uifontfile_format_proto_init() }
func file_uifontfile_format_proto_init() {
	if File_uifontfile_format_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uifontfile_format_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CUIFontFilePB); i {
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
		file_uifontfile_format_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CUIFontFilePackagePB); i {
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
		file_uifontfile_format_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CUIFontFilePackagePB_CUIEncryptedFontFilePB); i {
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
			RawDescriptor: file_uifontfile_format_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uifontfile_format_proto_goTypes,
		DependencyIndexes: file_uifontfile_format_proto_depIdxs,
		MessageInfos:      file_uifontfile_format_proto_msgTypes,
	}.Build()
	File_uifontfile_format_proto = out.File
	file_uifontfile_format_proto_rawDesc = nil
	file_uifontfile_format_proto_goTypes = nil
	file_uifontfile_format_proto_depIdxs = nil
}
