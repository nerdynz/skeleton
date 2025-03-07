// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: image.proto

package image

import (
	base "github.com/nerdynz/skeleton/rpc/base"
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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUlid      string  `protobuf:"bytes,1,opt,name=image_ulid,json=imageUlid,proto3" json:"image_ulid,omitempty" db:"image_ulid"`                  // @gotags: db:"image_ulid"
	Image          string  `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty" db:"image"`                                           // @gotags: db:"image"
	OriginalHeight float64 `protobuf:"fixed64,3,opt,name=original_height,json=originalHeight,proto3" json:"original_height,omitempty" db:"original_height"` // @gotags: db:"original_height"
	OriginalWidth  float64 `protobuf:"fixed64,4,opt,name=original_width,json=originalWidth,proto3" json:"original_width,omitempty" db:"original_width"`    // @gotags: db:"original_width"
	Top            float64 `protobuf:"fixed64,5,opt,name=top,proto3" json:"top,omitempty" db:"top"`                                             // @gotags: db:"top"
	Left           float64 `protobuf:"fixed64,6,opt,name=left,proto3" json:"left,omitempty" db:"left"`                                           // @gotags: db:"left"
	Scale          float64 `protobuf:"fixed64,7,opt,name=scale,proto3" json:"scale,omitempty" db:"scale"`                                         // @gotags: db:"scale"
	CropHeight     float64 `protobuf:"fixed64,8,opt,name=crop_height,json=cropHeight,proto3" json:"crop_height,omitempty" db:"crop_height"`             // @gotags: db:"crop_height"
	CropWidth      float64 `protobuf:"fixed64,9,opt,name=crop_width,json=cropWidth,proto3" json:"crop_width,omitempty" db:"crop_width"`                // @gotags: db:"crop_width"
	SiteUlid       string  `protobuf:"bytes,10,opt,name=site_ulid,json=siteUlid,proto3" json:"site_ulid,omitempty" db:"site_ulid"`                    // @gotags: db:"site_ulid"
	DateCreated    string  `protobuf:"bytes,11,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty" db:"date_created"`           // @gotags: db:"date_created"
	DateModified   string  `protobuf:"bytes,12,opt,name=date_modified,json=dateModified,proto3" json:"date_modified,omitempty" db:"date_modified"`        // @gotags: db:"date_modified"
	IsProcessed    bool    `protobuf:"varint,13,opt,name=is_processed,json=isProcessed,proto3" json:"is_processed,omitempty" db:"is_processed"`          // @gotags: db:"is_processed"
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetImageUlid() string {
	if x != nil {
		return x.ImageUlid
	}
	return ""
}

func (x *Image) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Image) GetOriginalHeight() float64 {
	if x != nil {
		return x.OriginalHeight
	}
	return 0
}

func (x *Image) GetOriginalWidth() float64 {
	if x != nil {
		return x.OriginalWidth
	}
	return 0
}

func (x *Image) GetTop() float64 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *Image) GetLeft() float64 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *Image) GetScale() float64 {
	if x != nil {
		return x.Scale
	}
	return 0
}

func (x *Image) GetCropHeight() float64 {
	if x != nil {
		return x.CropHeight
	}
	return 0
}

func (x *Image) GetCropWidth() float64 {
	if x != nil {
		return x.CropWidth
	}
	return 0
}

func (x *Image) GetSiteUlid() string {
	if x != nil {
		return x.SiteUlid
	}
	return ""
}

func (x *Image) GetDateCreated() string {
	if x != nil {
		return x.DateCreated
	}
	return ""
}

func (x *Image) GetDateModified() string {
	if x != nil {
		return x.DateModified
	}
	return ""
}

func (x *Image) GetIsProcessed() bool {
	if x != nil {
		return x.IsProcessed
	}
	return false
}

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{1}
}

func (x *Images) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type ImagesPaged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PagedInfo *base.PagedInfo `protobuf:"bytes,1,opt,name=pagedInfo,proto3" json:"pagedInfo,omitempty"`
	Records   []*Image        `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ImagesPaged) Reset() {
	*x = ImagesPaged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagesPaged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagesPaged) ProtoMessage() {}

func (x *ImagesPaged) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagesPaged.ProtoReflect.Descriptor instead.
func (*ImagesPaged) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{2}
}

func (x *ImagesPaged) GetPagedInfo() *base.PagedInfo {
	if x != nil {
		return x.PagedInfo
	}
	return nil
}

func (x *ImagesPaged) GetRecords() []*Image {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_image_proto protoreflect.FileDescriptor

var file_image_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x6c, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x74, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x6f, 0x70, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x6f, 0x70, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x6f, 0x70, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x72, 0x6f, 0x70, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x6c, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x69, 0x74, 0x65, 0x55, 0x6c, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x31, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x6b, 0x0a, 0x0b, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x6b,
	0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x32, 0xd8, 0x01, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x50, 0x43, 0x12, 0x2e, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x1a, 0x0f, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x0f, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x1a, 0x0f, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x13, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f,
	0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x64, 0x12, 0x32, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x73,
	0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x1a, 0x11,
	0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x65, 0x72, 0x64, 0x79, 0x6e, 0x7a, 0x2f, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData = file_image_proto_rawDesc
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_proto_rawDescData)
	})
	return file_image_proto_rawDescData
}

var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_image_proto_goTypes = []interface{}{
	(*Image)(nil),          // 0: skeleton.Image
	(*Images)(nil),         // 1: skeleton.Images
	(*ImagesPaged)(nil),    // 2: skeleton.ImagesPaged
	(*base.PagedInfo)(nil), // 3: skeleton.PagedInfo
	(*base.Lookup)(nil),    // 4: skeleton.Lookup
	(*base.Deleted)(nil),   // 5: skeleton.Deleted
}
var file_image_proto_depIdxs = []int32{
	0, // 0: skeleton.Images.images:type_name -> skeleton.Image
	3, // 1: skeleton.ImagesPaged.pagedInfo:type_name -> skeleton.PagedInfo
	0, // 2: skeleton.ImagesPaged.records:type_name -> skeleton.Image
	4, // 3: skeleton.ImageRPC.LoadImage:input_type -> skeleton.Lookup
	0, // 4: skeleton.ImageRPC.SaveImage:input_type -> skeleton.Image
	3, // 5: skeleton.ImageRPC.PagedImages:input_type -> skeleton.PagedInfo
	4, // 6: skeleton.ImageRPC.DeleteImage:input_type -> skeleton.Lookup
	0, // 7: skeleton.ImageRPC.LoadImage:output_type -> skeleton.Image
	0, // 8: skeleton.ImageRPC.SaveImage:output_type -> skeleton.Image
	2, // 9: skeleton.ImageRPC.PagedImages:output_type -> skeleton.ImagesPaged
	5, // 10: skeleton.ImageRPC.DeleteImage:output_type -> skeleton.Deleted
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
		file_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagesPaged); i {
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
			RawDescriptor: file_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_rawDesc = nil
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}
