package main

import "google.golang.org/protobuf/runtime/protoimpl"

// protobuf 基本类型和默认值在go中的对应关系
// Protobuf 	Go			默认值（Go）
// double		float64		0
// float		float32		0
// int32		int32		0
// int64		int64		0
// uint32		uint32		0
// uint64		uint64		0
// sint32		int32		0
// sint64		int64		0
// fixed32		uint32		0
// fixed64		uint64		0
// sfixed32		int32		0
// sfixed64		int64		0
// bool			bool		FALSE
// string		string		""（空字符串）
// bytes		[]byte		nil（空切片）

// 特殊类型
// Protobuf 	Go 					默认值（Go）
// message		指针（*StructName）	nil（空指针）
// enum	int32	0					（枚举的第一个值）
// repeated		切片（[]T）			nil（空切片）
// map<K, V>	map[K]V				nil（空 map)

// 	1.	Protobuf v3 默认不存储默认值：
// 		如果字段未赋值，Go 解析时不会包含该字段，返回默认值。
// 		例如，int32 默认值是 0，如果未设置该字段，Go 解析时不会在 proto.Marshal() 结果中体现这个字段
// 	2.	处理 optional 类型（Protobuf v3+）：
// 		optional 字段在 Go 代码中会被生成为指针类型（区分 “未设置” 和 “设置了但值为 0”）
//  	如： 	optional int32 value = 1;
//		对应：	Value *int32 `protobuf:"varint,1,opt,name=value"`
//

func main() {
	// message Hello {
	// 	string name = 1;
	// 	int32 age = 2;
	// 	repeated string tags = 3;
	// 	map<string, int32> scores = 4;
	// 	optional bool is_active = 5;
	// }
}

type Hello struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age           int32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Tags          []string               `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Scores        map[string]int32       `protobuf:"bytes,4,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	IsActive      *bool                  `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
