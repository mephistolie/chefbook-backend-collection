// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0
// source: v1/get-recipes-collections-map.proto

package v1

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

type RecipeCollections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []string `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *RecipeCollections) Reset() {
	*x = RecipeCollections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipes_collections_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeCollections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeCollections) ProtoMessage() {}

func (x *RecipeCollections) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipes_collections_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeCollections.ProtoReflect.Descriptor instead.
func (*RecipeCollections) Descriptor() ([]byte, []int) {
	return file_v1_get_recipes_collections_map_proto_rawDescGZIP(), []int{0}
}

func (x *RecipeCollections) GetCollections() []string {
	if x != nil {
		return x.Collections
	}
	return nil
}

type GetRecipesCollectionsMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	RecipesIds []string `protobuf:"bytes,2,rep,name=recipesIds,proto3" json:"recipesIds,omitempty"`
}

func (x *GetRecipesCollectionsMapRequest) Reset() {
	*x = GetRecipesCollectionsMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipes_collections_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecipesCollectionsMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecipesCollectionsMapRequest) ProtoMessage() {}

func (x *GetRecipesCollectionsMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipes_collections_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecipesCollectionsMapRequest.ProtoReflect.Descriptor instead.
func (*GetRecipesCollectionsMapRequest) Descriptor() ([]byte, []int) {
	return file_v1_get_recipes_collections_map_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecipesCollectionsMapRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetRecipesCollectionsMapRequest) GetRecipesIds() []string {
	if x != nil {
		return x.RecipesIds
	}
	return nil
}

type GetRecipesCollectionsMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeCollections map[string]*RecipeCollections `protobuf:"bytes,1,rep,name=recipeCollections,proto3" json:"recipeCollections,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CollectionsInfo   map[string]*Collection        `protobuf:"bytes,2,rep,name=collectionsInfo,proto3" json:"collectionsInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetRecipesCollectionsMapResponse) Reset() {
	*x = GetRecipesCollectionsMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipes_collections_map_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecipesCollectionsMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecipesCollectionsMapResponse) ProtoMessage() {}

func (x *GetRecipesCollectionsMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipes_collections_map_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecipesCollectionsMapResponse.ProtoReflect.Descriptor instead.
func (*GetRecipesCollectionsMapResponse) Descriptor() ([]byte, []int) {
	return file_v1_get_recipes_collections_map_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecipesCollectionsMapResponse) GetRecipeCollections() map[string]*RecipeCollections {
	if x != nil {
		return x.RecipeCollections
	}
	return nil
}

func (x *GetRecipesCollectionsMapResponse) GetCollectionsInfo() map[string]*Collection {
	if x != nil {
		return x.CollectionsInfo
	}
	return nil
}

var File_v1_get_recipes_collections_map_proto protoreflect.FileDescriptor

var file_v1_get_recipes_collections_map_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73,
	0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x6d, 0x61, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x59, 0x0a, 0x1f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73,
	0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x73, 0x49, 0x64, 0x73, 0x22, 0xa3, 0x03, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x11, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x63, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x5b, 0x0a, 0x16, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x52, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_get_recipes_collections_map_proto_rawDescOnce sync.Once
	file_v1_get_recipes_collections_map_proto_rawDescData = file_v1_get_recipes_collections_map_proto_rawDesc
)

func file_v1_get_recipes_collections_map_proto_rawDescGZIP() []byte {
	file_v1_get_recipes_collections_map_proto_rawDescOnce.Do(func() {
		file_v1_get_recipes_collections_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_get_recipes_collections_map_proto_rawDescData)
	})
	return file_v1_get_recipes_collections_map_proto_rawDescData
}

var file_v1_get_recipes_collections_map_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_get_recipes_collections_map_proto_goTypes = []interface{}{
	(*RecipeCollections)(nil),                // 0: v1.RecipeCollections
	(*GetRecipesCollectionsMapRequest)(nil),  // 1: v1.GetRecipesCollectionsMapRequest
	(*GetRecipesCollectionsMapResponse)(nil), // 2: v1.GetRecipesCollectionsMapResponse
	nil,                                      // 3: v1.GetRecipesCollectionsMapResponse.RecipeCollectionsEntry
	nil,                                      // 4: v1.GetRecipesCollectionsMapResponse.CollectionsInfoEntry
	(*Collection)(nil),                       // 5: v1.Collection
}
var file_v1_get_recipes_collections_map_proto_depIdxs = []int32{
	3, // 0: v1.GetRecipesCollectionsMapResponse.recipeCollections:type_name -> v1.GetRecipesCollectionsMapResponse.RecipeCollectionsEntry
	4, // 1: v1.GetRecipesCollectionsMapResponse.collectionsInfo:type_name -> v1.GetRecipesCollectionsMapResponse.CollectionsInfoEntry
	0, // 2: v1.GetRecipesCollectionsMapResponse.RecipeCollectionsEntry.value:type_name -> v1.RecipeCollections
	5, // 3: v1.GetRecipesCollectionsMapResponse.CollectionsInfoEntry.value:type_name -> v1.Collection
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_get_recipes_collections_map_proto_init() }
func file_v1_get_recipes_collections_map_proto_init() {
	if File_v1_get_recipes_collections_map_proto != nil {
		return
	}
	file_v1_get_collection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_get_recipes_collections_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeCollections); i {
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
		file_v1_get_recipes_collections_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecipesCollectionsMapRequest); i {
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
		file_v1_get_recipes_collections_map_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecipesCollectionsMapResponse); i {
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
			RawDescriptor: file_v1_get_recipes_collections_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_get_recipes_collections_map_proto_goTypes,
		DependencyIndexes: file_v1_get_recipes_collections_map_proto_depIdxs,
		MessageInfos:      file_v1_get_recipes_collections_map_proto_msgTypes,
	}.Build()
	File_v1_get_recipes_collections_map_proto = out.File
	file_v1_get_recipes_collections_map_proto_rawDesc = nil
	file_v1_get_recipes_collections_map_proto_goTypes = nil
	file_v1_get_recipes_collections_map_proto_depIdxs = nil
}
