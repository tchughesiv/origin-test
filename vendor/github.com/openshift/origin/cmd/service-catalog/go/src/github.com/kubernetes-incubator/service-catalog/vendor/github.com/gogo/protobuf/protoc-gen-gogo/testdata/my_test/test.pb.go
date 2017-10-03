// Code generated by protoc-gen-gogo.
// source: my_test/test.proto
// DO NOT EDIT!

/*
Package my_test is a generated protocol buffer package.

This package holds interesting messages.

It is generated from these files:
	my_test/test.proto

It has these top-level messages:
	Request
	Reply
	OtherBase
	ReplyExtensions
	OtherReplyExtensions
	OldReply
	Communique
*/
package my_test

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "multi"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type HatType int32

const (
	// deliberately skipping 0
	HatType_FEDORA HatType = 1
	HatType_FEZ    HatType = 2
)

var HatType_name = map[int32]string{
	1: "FEDORA",
	2: "FEZ",
}
var HatType_value = map[string]int32{
	"FEDORA": 1,
	"FEZ":    2,
}

func (x HatType) Enum() *HatType {
	p := new(HatType)
	*p = x
	return p
}
func (x HatType) String() string {
	return proto.EnumName(HatType_name, int32(x))
}
func (x *HatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HatType_value, data, "HatType")
	if err != nil {
		return err
	}
	*x = HatType(value)
	return nil
}
func (HatType) EnumDescriptor() ([]byte, []int) { return fileDescriptorTest, []int{0} }

// This enum represents days of the week.
type Days int32

const (
	Days_MONDAY  Days = 1
	Days_TUESDAY Days = 2
	Days_LUNDI   Days = 1
)

var Days_name = map[int32]string{
	1: "MONDAY",
	2: "TUESDAY",
	// Duplicate value: 1: "LUNDI",
}
var Days_value = map[string]int32{
	"MONDAY":  1,
	"TUESDAY": 2,
	"LUNDI":   1,
}

func (x Days) Enum() *Days {
	p := new(Days)
	*p = x
	return p
}
func (x Days) String() string {
	return proto.EnumName(Days_name, int32(x))
}
func (x *Days) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Days_value, data, "Days")
	if err != nil {
		return err
	}
	*x = Days(value)
	return nil
}
func (Days) EnumDescriptor() ([]byte, []int) { return fileDescriptorTest, []int{1} }

type Request_Color int32

const (
	Request_RED   Request_Color = 0
	Request_GREEN Request_Color = 1
	Request_BLUE  Request_Color = 2
)

var Request_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Request_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Request_Color) Enum() *Request_Color {
	p := new(Request_Color)
	*p = x
	return p
}
func (x Request_Color) String() string {
	return proto.EnumName(Request_Color_name, int32(x))
}
func (x *Request_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Color_value, data, "Request_Color")
	if err != nil {
		return err
	}
	*x = Request_Color(value)
	return nil
}
func (Request_Color) EnumDescriptor() ([]byte, []int) { return fileDescriptorTest, []int{0, 0} }

type Reply_Entry_Game int32

const (
	Reply_Entry_FOOTBALL Reply_Entry_Game = 1
	Reply_Entry_TENNIS   Reply_Entry_Game = 2
)

var Reply_Entry_Game_name = map[int32]string{
	1: "FOOTBALL",
	2: "TENNIS",
}
var Reply_Entry_Game_value = map[string]int32{
	"FOOTBALL": 1,
	"TENNIS":   2,
}

func (x Reply_Entry_Game) Enum() *Reply_Entry_Game {
	p := new(Reply_Entry_Game)
	*p = x
	return p
}
func (x Reply_Entry_Game) String() string {
	return proto.EnumName(Reply_Entry_Game_name, int32(x))
}
func (x *Reply_Entry_Game) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Entry_Game_value, data, "Reply_Entry_Game")
	if err != nil {
		return err
	}
	*x = Reply_Entry_Game(value)
	return nil
}
func (Reply_Entry_Game) EnumDescriptor() ([]byte, []int) { return fileDescriptorTest, []int{1, 0, 0} }

// This is a message that might be sent somewhere.
type Request struct {
	Key []int64 `protobuf:"varint,1,rep,name=key" json:"key,omitempty"`
	//  optional imp.ImportedMessage imported_message = 2;
	Hue *Request_Color `protobuf:"varint,3,opt,name=hue,enum=my.test.Request_Color" json:"hue,omitempty"`
	Hat *HatType       `protobuf:"varint,4,opt,name=hat,enum=my.test.HatType,def=1" json:"hat,omitempty"`
	//  optional imp.ImportedMessage.Owner owner = 6;
	Deadline  *float32           `protobuf:"fixed32,7,opt,name=deadline,def=inf" json:"deadline,omitempty"`
	Somegroup *Request_SomeGroup `protobuf:"group,8,opt,name=SomeGroup,json=somegroup" json:"somegroup,omitempty"`
	// This is a map field. It will generate map[int32]string.
	NameMapping map[int32]string `protobuf:"bytes,14,rep,name=name_mapping,json=nameMapping" json:"name_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// This is a map field whose value type is a message.
	MsgMapping map[int64]*Reply `protobuf:"bytes,15,rep,name=msg_mapping,json=msgMapping" json:"msg_mapping,omitempty" protobuf_key:"zigzag64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Reset_     *int32           `protobuf:"varint,12,opt,name=reset" json:"reset,omitempty"`
	// This field should not conflict with any getters.
	GetKey_          *string `protobuf:"bytes,16,opt,name=get_key,json=getKey" json:"get_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{0} }

const Default_Request_Hat HatType = HatType_FEDORA

var Default_Request_Deadline float32 = float32(math.Inf(1))

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetHue() Request_Color {
	if m != nil && m.Hue != nil {
		return *m.Hue
	}
	return Request_RED
}

func (m *Request) GetHat() HatType {
	if m != nil && m.Hat != nil {
		return *m.Hat
	}
	return Default_Request_Hat
}

func (m *Request) GetDeadline() float32 {
	if m != nil && m.Deadline != nil {
		return *m.Deadline
	}
	return Default_Request_Deadline
}

func (m *Request) GetSomegroup() *Request_SomeGroup {
	if m != nil {
		return m.Somegroup
	}
	return nil
}

func (m *Request) GetNameMapping() map[int32]string {
	if m != nil {
		return m.NameMapping
	}
	return nil
}

func (m *Request) GetMsgMapping() map[int64]*Reply {
	if m != nil {
		return m.MsgMapping
	}
	return nil
}

func (m *Request) GetReset_() int32 {
	if m != nil && m.Reset_ != nil {
		return *m.Reset_
	}
	return 0
}

func (m *Request) GetGetKey_() string {
	if m != nil && m.GetKey_ != nil {
		return *m.GetKey_
	}
	return ""
}

type Request_SomeGroup struct {
	GroupField       *int32 `protobuf:"varint,9,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request_SomeGroup) Reset()                    { *m = Request_SomeGroup{} }
func (m *Request_SomeGroup) String() string            { return proto.CompactTextString(m) }
func (*Request_SomeGroup) ProtoMessage()               {}
func (*Request_SomeGroup) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{0, 0} }

func (m *Request_SomeGroup) GetGroupField() int32 {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return 0
}

type Reply struct {
	Found            []*Reply_Entry            `protobuf:"bytes,1,rep,name=found" json:"found,omitempty"`
	CompactKeys      []int32                   `protobuf:"varint,2,rep,packed,name=compact_keys,json=compactKeys" json:"compact_keys,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{1} }

var extRange_Reply = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Reply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Reply
}
func (m *Reply) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Reply) GetFound() []*Reply_Entry {
	if m != nil {
		return m.Found
	}
	return nil
}

func (m *Reply) GetCompactKeys() []int32 {
	if m != nil {
		return m.CompactKeys
	}
	return nil
}

type Reply_Entry struct {
	KeyThatNeeds_1234Camel_CasIng *int64 `protobuf:"varint,1,req,name=key_that_needs_1234camel_CasIng,json=keyThatNeeds1234camelCasIng" json:"key_that_needs_1234camel_CasIng,omitempty"`
	Value                         *int64 `protobuf:"varint,2,opt,name=value,def=7" json:"value,omitempty"`
	XMyFieldName_2                *int64 `protobuf:"varint,3,opt,name=_my_field_name_2,json=myFieldName2" json:"_my_field_name_2,omitempty"`
	XXX_unrecognized              []byte `json:"-"`
}

func (m *Reply_Entry) Reset()                    { *m = Reply_Entry{} }
func (m *Reply_Entry) String() string            { return proto.CompactTextString(m) }
func (*Reply_Entry) ProtoMessage()               {}
func (*Reply_Entry) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{1, 0} }

const Default_Reply_Entry_Value int64 = 7

func (m *Reply_Entry) GetKeyThatNeeds_1234Camel_CasIng() int64 {
	if m != nil && m.KeyThatNeeds_1234Camel_CasIng != nil {
		return *m.KeyThatNeeds_1234Camel_CasIng
	}
	return 0
}

func (m *Reply_Entry) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_Reply_Entry_Value
}

func (m *Reply_Entry) GetXMyFieldName_2() int64 {
	if m != nil && m.XMyFieldName_2 != nil {
		return *m.XMyFieldName_2
	}
	return 0
}

type OtherBase struct {
	Name             *string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *OtherBase) Reset()                    { *m = OtherBase{} }
func (m *OtherBase) String() string            { return proto.CompactTextString(m) }
func (*OtherBase) ProtoMessage()               {}
func (*OtherBase) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{2} }

var extRange_OtherBase = []proto.ExtensionRange{
	{100, 536870911},
}

func (*OtherBase) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OtherBase
}
func (m *OtherBase) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *OtherBase) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ReplyExtensions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReplyExtensions) Reset()                    { *m = ReplyExtensions{} }
func (m *ReplyExtensions) String() string            { return proto.CompactTextString(m) }
func (*ReplyExtensions) ProtoMessage()               {}
func (*ReplyExtensions) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{3} }

var E_ReplyExtensions_Time = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*float64)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.time",
	Tag:           "fixed64,101,opt,name=time",
}

var E_ReplyExtensions_Carrot = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         105,
	Name:          "my.test.ReplyExtensions.carrot",
	Tag:           "bytes,105,opt,name=carrot",
}

var E_ReplyExtensions_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*OtherBase)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.donut",
	Tag:           "bytes,101,opt,name=donut",
}

type OtherReplyExtensions struct {
	Key              *int32 `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OtherReplyExtensions) Reset()                    { *m = OtherReplyExtensions{} }
func (m *OtherReplyExtensions) String() string            { return proto.CompactTextString(m) }
func (*OtherReplyExtensions) ProtoMessage()               {}
func (*OtherReplyExtensions) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{4} }

func (m *OtherReplyExtensions) GetKey() int32 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

type OldReply struct {
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *OldReply) Reset()                    { *m = OldReply{} }
func (m *OldReply) String() string            { return proto.CompactTextString(m) }
func (*OldReply) ProtoMessage()               {}
func (*OldReply) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{5} }

func (m *OldReply) Marshal() ([]byte, error) {
	return proto.MarshalMessageSet(m.ExtensionMap())
}
func (m *OldReply) Unmarshal(buf []byte) error {
	return proto.UnmarshalMessageSet(buf, m.ExtensionMap())
}
func (m *OldReply) MarshalJSON() ([]byte, error) {
	return proto.MarshalMessageSetJSON(m.XXX_extensions)
}
func (m *OldReply) UnmarshalJSON(buf []byte) error {
	return proto.UnmarshalMessageSetJSON(buf, m.XXX_extensions)
}

// ensure OldReply satisfies proto.Marshaler and proto.Unmarshaler
var _ proto.Marshaler = (*OldReply)(nil)
var _ proto.Unmarshaler = (*OldReply)(nil)

var extRange_OldReply = []proto.ExtensionRange{
	{100, 2147483646},
}

func (*OldReply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OldReply
}
func (m *OldReply) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

type Communique struct {
	MakeMeCry *bool `protobuf:"varint,1,opt,name=make_me_cry,json=makeMeCry" json:"make_me_cry,omitempty"`
	// This is a oneof, called "union".
	//
	// Types that are valid to be assigned to Union:
	//	*Communique_Number
	//	*Communique_Name
	//	*Communique_Data
	//	*Communique_TempC
	//	*Communique_Height
	//	*Communique_Today
	//	*Communique_Maybe
	//	*Communique_Delta_
	//	*Communique_Msg
	//	*Communique_Somegroup
	Union            isCommunique_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Communique) Reset()                    { *m = Communique{} }
func (m *Communique) String() string            { return proto.CompactTextString(m) }
func (*Communique) ProtoMessage()               {}
func (*Communique) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{6} }

type isCommunique_Union interface {
	isCommunique_Union()
}

type Communique_Number struct {
	Number int32 `protobuf:"varint,5,opt,name=number,oneof"`
}
type Communique_Name struct {
	Name string `protobuf:"bytes,6,opt,name=name,oneof"`
}
type Communique_Data struct {
	Data []byte `protobuf:"bytes,7,opt,name=data,oneof"`
}
type Communique_TempC struct {
	TempC float64 `protobuf:"fixed64,8,opt,name=temp_c,json=tempC,oneof"`
}
type Communique_Height struct {
	Height float32 `protobuf:"fixed32,9,opt,name=height,oneof"`
}
type Communique_Today struct {
	Today Days `protobuf:"varint,10,opt,name=today,enum=my.test.Days,oneof"`
}
type Communique_Maybe struct {
	Maybe bool `protobuf:"varint,11,opt,name=maybe,oneof"`
}
type Communique_Delta_ struct {
	Delta int32 `protobuf:"zigzag32,12,opt,name=delta,oneof"`
}
type Communique_Msg struct {
	Msg *Reply `protobuf:"bytes,13,opt,name=msg,oneof"`
}
type Communique_Somegroup struct {
	Somegroup *Communique_SomeGroup `protobuf:"group,14,opt,name=SomeGroup,json=somegroup,oneof"`
}

func (*Communique_Number) isCommunique_Union()    {}
func (*Communique_Name) isCommunique_Union()      {}
func (*Communique_Data) isCommunique_Union()      {}
func (*Communique_TempC) isCommunique_Union()     {}
func (*Communique_Height) isCommunique_Union()    {}
func (*Communique_Today) isCommunique_Union()     {}
func (*Communique_Maybe) isCommunique_Union()     {}
func (*Communique_Delta_) isCommunique_Union()    {}
func (*Communique_Msg) isCommunique_Union()       {}
func (*Communique_Somegroup) isCommunique_Union() {}

func (m *Communique) GetUnion() isCommunique_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Communique) GetMakeMeCry() bool {
	if m != nil && m.MakeMeCry != nil {
		return *m.MakeMeCry
	}
	return false
}

func (m *Communique) GetNumber() int32 {
	if x, ok := m.GetUnion().(*Communique_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Communique) GetName() string {
	if x, ok := m.GetUnion().(*Communique_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Communique) GetData() []byte {
	if x, ok := m.GetUnion().(*Communique_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Communique) GetTempC() float64 {
	if x, ok := m.GetUnion().(*Communique_TempC); ok {
		return x.TempC
	}
	return 0
}

func (m *Communique) GetHeight() float32 {
	if x, ok := m.GetUnion().(*Communique_Height); ok {
		return x.Height
	}
	return 0
}

func (m *Communique) GetToday() Days {
	if x, ok := m.GetUnion().(*Communique_Today); ok {
		return x.Today
	}
	return Days_MONDAY
}

func (m *Communique) GetMaybe() bool {
	if x, ok := m.GetUnion().(*Communique_Maybe); ok {
		return x.Maybe
	}
	return false
}

func (m *Communique) GetDelta() int32 {
	if x, ok := m.GetUnion().(*Communique_Delta_); ok {
		return x.Delta
	}
	return 0
}

func (m *Communique) GetMsg() *Reply {
	if x, ok := m.GetUnion().(*Communique_Msg); ok {
		return x.Msg
	}
	return nil
}

func (m *Communique) GetSomegroup() *Communique_SomeGroup {
	if x, ok := m.GetUnion().(*Communique_Somegroup); ok {
		return x.Somegroup
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Communique) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Communique_OneofMarshaler, _Communique_OneofUnmarshaler, _Communique_OneofSizer, []interface{}{
		(*Communique_Number)(nil),
		(*Communique_Name)(nil),
		(*Communique_Data)(nil),
		(*Communique_TempC)(nil),
		(*Communique_Height)(nil),
		(*Communique_Today)(nil),
		(*Communique_Maybe)(nil),
		(*Communique_Delta_)(nil),
		(*Communique_Msg)(nil),
		(*Communique_Somegroup)(nil),
	}
}

func _Communique_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Communique)
	// union
	switch x := m.Union.(type) {
	case *Communique_Number:
		_ = b.EncodeVarint(5<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Number))
	case *Communique_Name:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Name)
	case *Communique_Data:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Data)
	case *Communique_TempC:
		_ = b.EncodeVarint(8<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.TempC))
	case *Communique_Height:
		_ = b.EncodeVarint(9<<3 | proto.WireFixed32)
		_ = b.EncodeFixed32(uint64(math.Float32bits(x.Height)))
	case *Communique_Today:
		_ = b.EncodeVarint(10<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Today))
	case *Communique_Maybe:
		t := uint64(0)
		if x.Maybe {
			t = 1
		}
		_ = b.EncodeVarint(11<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *Communique_Delta_:
		_ = b.EncodeVarint(12<<3 | proto.WireVarint)
		_ = b.EncodeZigzag32(uint64(x.Delta))
	case *Communique_Msg:
		_ = b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Msg); err != nil {
			return err
		}
	case *Communique_Somegroup:
		_ = b.EncodeVarint(14<<3 | proto.WireStartGroup)
		if err := b.Marshal(x.Somegroup); err != nil {
			return err
		}
		_ = b.EncodeVarint(14<<3 | proto.WireEndGroup)
	case nil:
	default:
		return fmt.Errorf("Communique.Union has unexpected type %T", x)
	}
	return nil
}

func _Communique_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Communique)
	switch tag {
	case 5: // union.number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Number{int32(x)}
		return true, err
	case 6: // union.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &Communique_Name{x}
		return true, err
	case 7: // union.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Union = &Communique_Data{x}
		return true, err
	case 8: // union.temp_c
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Union = &Communique_TempC{math.Float64frombits(x)}
		return true, err
	case 9: // union.height
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Union = &Communique_Height{math.Float32frombits(uint32(x))}
		return true, err
	case 10: // union.today
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Today{Days(x)}
		return true, err
	case 11: // union.maybe
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Maybe{x != 0}
		return true, err
	case 12: // union.delta
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag32()
		m.Union = &Communique_Delta_{int32(x)}
		return true, err
	case 13: // union.msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Reply)
		err := b.DecodeMessage(msg)
		m.Union = &Communique_Msg{msg}
		return true, err
	case 14: // union.somegroup
		if wire != proto.WireStartGroup {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Communique_SomeGroup)
		err := b.DecodeGroup(msg)
		m.Union = &Communique_Somegroup{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Communique_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Communique)
	// union
	switch x := m.Union.(type) {
	case *Communique_Number:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Number))
	case *Communique_Name:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *Communique_Data:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case *Communique_TempC:
		n += proto.SizeVarint(8<<3 | proto.WireFixed64)
		n += 8
	case *Communique_Height:
		n += proto.SizeVarint(9<<3 | proto.WireFixed32)
		n += 4
	case *Communique_Today:
		n += proto.SizeVarint(10<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Today))
	case *Communique_Maybe:
		n += proto.SizeVarint(11<<3 | proto.WireVarint)
		n += 1
	case *Communique_Delta_:
		n += proto.SizeVarint(12<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64((uint32(x.Delta) << 1) ^ uint32((int32(x.Delta) >> 31))))
	case *Communique_Msg:
		s := proto.Size(x.Msg)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Communique_Somegroup:
		n += proto.SizeVarint(14<<3 | proto.WireStartGroup)
		n += proto.Size(x.Somegroup)
		n += proto.SizeVarint(14<<3 | proto.WireEndGroup)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Communique_SomeGroup struct {
	Member           *string `protobuf:"bytes,15,opt,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Communique_SomeGroup) Reset()                    { *m = Communique_SomeGroup{} }
func (m *Communique_SomeGroup) String() string            { return proto.CompactTextString(m) }
func (*Communique_SomeGroup) ProtoMessage()               {}
func (*Communique_SomeGroup) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{6, 0} }

func (m *Communique_SomeGroup) GetMember() string {
	if m != nil && m.Member != nil {
		return *m.Member
	}
	return ""
}

type Communique_Delta struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Communique_Delta) Reset()                    { *m = Communique_Delta{} }
func (m *Communique_Delta) String() string            { return proto.CompactTextString(m) }
func (*Communique_Delta) ProtoMessage()               {}
func (*Communique_Delta) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{6, 1} }

var E_Tag = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*string)(nil),
	Field:         103,
	Name:          "my.test.tag",
	Tag:           "bytes,103,opt,name=tag",
}

var E_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*OtherReplyExtensions)(nil),
	Field:         106,
	Name:          "my.test.donut",
	Tag:           "bytes,106,opt,name=donut",
}

func init() {
	proto.RegisterType((*Request)(nil), "my.test.Request")
	proto.RegisterType((*Request_SomeGroup)(nil), "my.test.Request.SomeGroup")
	proto.RegisterType((*Reply)(nil), "my.test.Reply")
	proto.RegisterType((*Reply_Entry)(nil), "my.test.Reply.Entry")
	proto.RegisterType((*OtherBase)(nil), "my.test.OtherBase")
	proto.RegisterType((*ReplyExtensions)(nil), "my.test.ReplyExtensions")
	proto.RegisterType((*OtherReplyExtensions)(nil), "my.test.OtherReplyExtensions")
	proto.RegisterType((*OldReply)(nil), "my.test.OldReply")
	proto.RegisterType((*Communique)(nil), "my.test.Communique")
	proto.RegisterType((*Communique_SomeGroup)(nil), "my.test.Communique.SomeGroup")
	proto.RegisterType((*Communique_Delta)(nil), "my.test.Communique.Delta")
	proto.RegisterEnum("my.test.HatType", HatType_name, HatType_value)
	proto.RegisterEnum("my.test.Days", Days_name, Days_value)
	proto.RegisterEnum("my.test.Request_Color", Request_Color_name, Request_Color_value)
	proto.RegisterEnum("my.test.Reply_Entry_Game", Reply_Entry_Game_name, Reply_Entry_Game_value)
	proto.RegisterExtension(E_ReplyExtensions_Time)
	proto.RegisterExtension(E_ReplyExtensions_Carrot)
	proto.RegisterExtension(E_ReplyExtensions_Donut)
	proto.RegisterExtension(E_Tag)
	proto.RegisterExtension(E_Donut)
}

var fileDescriptorTest = []byte{
	// 970 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xe3, 0x54,
	0x10, 0x5e, 0xdb, 0x71, 0x7e, 0x26, 0xd9, 0xd6, 0x1c, 0xad, 0x8a, 0x15, 0xb4, 0xbb, 0xc1, 0xb0,
	0xa8, 0x54, 0x28, 0xab, 0x35, 0x48, 0xac, 0x22, 0x81, 0x68, 0x12, 0xb7, 0x5b, 0xd1, 0x24, 0xd2,
	0x69, 0x7a, 0x01, 0x37, 0x96, 0x37, 0x3e, 0x4d, 0xcd, 0xc6, 0x76, 0x88, 0x4f, 0x10, 0xbe, 0xe3,
	0x29, 0xe0, 0x35, 0xb8, 0xe7, 0x85, 0x78, 0x0b, 0x98, 0x33, 0x0e, 0x4e, 0x9a, 0xac, 0xc8, 0x45,
	0xe4, 0x99, 0xf9, 0xe6, 0x3b, 0xf3, 0x0f, 0x2c, 0xce, 0x7d, 0x29, 0x32, 0xf9, 0x52, 0xfd, 0x75,
	0x97, 0xab, 0x54, 0xa6, 0xac, 0x16, 0xe7, 0x5d, 0x25, 0xb6, 0x59, 0xbc, 0x5e, 0xc8, 0xe8, 0x25,
	0xfd, 0xbf, 0x2a, 0x8c, 0xce, 0xdf, 0x15, 0xa8, 0x71, 0xf1, 0xf3, 0x1a, 0xed, 0xcc, 0x02, 0xe3,
	0x9d, 0xc8, 0x6d, 0xad, 0x63, 0x9c, 0x1a, 0x5c, 0x7d, 0xb2, 0x53, 0x30, 0xee, 0xd7, 0xc2, 0x36,
	0x3a, 0xda, 0xe9, 0x91, 0x7b, 0xd2, 0xdd, 0x10, 0x75, 0x37, 0x0e, 0xdd, 0x41, 0xba, 0x48, 0x57,
	0x5c, 0x41, 0xd8, 0x19, 0x22, 0x03, 0x69, 0x57, 0x08, 0x69, 0x95, 0xc8, 0x37, 0x81, 0x9c, 0xe6,
	0x4b, 0xd1, 0xab, 0x5e, 0x78, 0xc3, 0x09, 0x3f, 0xe7, 0x0a, 0xc4, 0x9e, 0x43, 0x3d, 0x14, 0x41,
	0xb8, 0x88, 0x12, 0x61, 0xd7, 0xd0, 0x41, 0xef, 0x19, 0x51, 0x72, 0xc7, 0x4b, 0x25, 0x7b, 0x0d,
	0x8d, 0x2c, 0x8d, 0xc5, 0x7c, 0x95, 0xae, 0x97, 0x76, 0x1d, 0x11, 0xe0, 0xb6, 0x0f, 0x1e, 0xbf,
	0x41, 0xc4, 0xa5, 0x42, 0xf0, 0x2d, 0x98, 0x0d, 0xa1, 0x95, 0x04, 0xb1, 0xf0, 0xe3, 0x60, 0xb9,
	0x8c, 0x92, 0xb9, 0x7d, 0x84, 0xb9, 0x34, 0xdd, 0x8f, 0x0f, 0x9c, 0xc7, 0x08, 0x1a, 0x15, 0x18,
	0x2f, 0x91, 0xab, 0x9c, 0x37, 0x93, 0xad, 0x86, 0x9d, 0x43, 0x33, 0xce, 0xe6, 0x25, 0xc9, 0x31,
	0x91, 0x74, 0x0e, 0x48, 0x46, 0xd9, 0xfc, 0x01, 0x07, 0xc4, 0xa5, 0x82, 0x3d, 0x01, 0x73, 0x25,
	0x32, 0x21, 0xed, 0x16, 0x86, 0x6f, 0xf2, 0x42, 0x60, 0x1f, 0x42, 0x6d, 0x2e, 0xa4, 0xaf, 0xaa,
	0x6c, 0xa1, 0xbe, 0xc1, 0xab, 0x28, 0x7e, 0x2f, 0xf2, 0xf6, 0x17, 0xd0, 0x28, 0xf3, 0xc1, 0xfa,
	0x34, 0x29, 0x1b, 0xff, 0x2e, 0x12, 0x8b, 0xd0, 0x6e, 0x10, 0x03, 0x90, 0xea, 0x42, 0x69, 0xda,
	0xdf, 0x82, 0xb5, 0x9f, 0xc0, 0xb6, 0x79, 0x0a, 0x4c, 0xcd, 0xc3, 0x10, 0x7e, 0x09, 0x16, 0xd8,
	0x3e, 0x9d, 0x9e, 0x2a, 0x84, 0x9e, 0xfe, 0x5a, 0x6b, 0x8f, 0xe0, 0x78, 0x2f, 0xf6, 0x5d, 0x77,
	0x56, 0xb8, 0x7f, 0xba, 0xeb, 0xde, 0x74, 0x8f, 0x76, 0xd2, 0x5f, 0x2e, 0xf2, 0x1d, 0x3a, 0xe7,
	0x05, 0x98, 0x34, 0x09, 0xac, 0x06, 0x06, 0xf7, 0x86, 0xd6, 0x23, 0xd6, 0x00, 0xf3, 0x92, 0x7b,
	0xde, 0xd8, 0xd2, 0x58, 0x1d, 0x2a, 0xfd, 0xeb, 0x5b, 0xcf, 0xd2, 0x9d, 0x3f, 0x74, 0x30, 0xc9,
	0x17, 0x87, 0xc5, 0xbc, 0x4b, 0xd7, 0x49, 0x48, 0xa3, 0xd6, 0x74, 0x9f, 0x3c, 0xa4, 0xee, 0x16,
	0xd5, 0x2c, 0x20, 0xec, 0x05, 0xb4, 0x66, 0x69, 0xbc, 0x0c, 0x66, 0x54, 0xb6, 0x0c, 0xa3, 0x31,
	0x4e, 0xcd, 0xbe, 0x6e, 0x69, 0xbc, 0xb9, 0xd1, 0x63, 0xfd, 0xb2, 0xf6, 0x9f, 0x1a, 0x98, 0x45,
	0x26, 0x43, 0x78, 0x8e, 0x40, 0x5f, 0xe2, 0xa4, 0xf9, 0x89, 0x10, 0x61, 0xe6, 0xbf, 0x72, 0xbf,
	0xfc, 0x6a, 0x86, 0xf5, 0x5a, 0xf8, 0x83, 0x20, 0xbb, 0xc2, 0x86, 0x6a, 0x1d, 0x1d, 0x27, 0xfc,
	0x23, 0x84, 0x4d, 0x11, 0x35, 0x56, 0xa0, 0x12, 0x53, 0x40, 0xb0, 0x53, 0x3b, 0xd9, 0x1b, 0x3d,
	0xed, 0xeb, 0x4d, 0xc2, 0xec, 0x33, 0xb0, 0x7c, 0x5c, 0x32, 0x6a, 0x8d, 0x4f, 0xb3, 0xe6, 0xd2,
	0x7e, 0x18, 0xbc, 0x15, 0xe7, 0xd4, 0x1e, 0xd5, 0x1a, 0xd7, 0xe9, 0x40, 0xe5, 0x12, 0x3f, 0x58,
	0x0b, 0xea, 0x17, 0x93, 0xc9, 0xb4, 0x7f, 0x7e, 0x7d, 0x8d, 0xd5, 0x00, 0xa8, 0x4e, 0xbd, 0xf1,
	0xf8, 0xea, 0xc6, 0xd2, 0xcf, 0xea, 0xf5, 0xd0, 0xfa, 0x0d, 0x7f, 0xba, 0xf3, 0x39, 0x34, 0x26,
	0xf2, 0x5e, 0xac, 0xfa, 0x41, 0x26, 0x18, 0x83, 0x8a, 0xa2, 0xa5, 0x56, 0x34, 0x38, 0x7d, 0xef,
	0x40, 0xff, 0xd2, 0xe0, 0x98, 0xaa, 0xe4, 0xfd, 0x2a, 0x45, 0x92, 0x45, 0x69, 0x92, 0xb9, 0x0e,
	0x54, 0x64, 0x84, 0x4f, 0xed, 0xb5, 0xc8, 0x16, 0xc8, 0xa0, 0x71, 0xb2, 0xb9, 0xdf, 0x41, 0x75,
	0x16, 0xac, 0x70, 0xe5, 0x0f, 0x50, 0x11, 0xb5, 0xd7, 0x7e, 0xa8, 0xdd, 0xb2, 0xf3, 0x8d, 0x9f,
	0xdb, 0x07, 0x33, 0x4c, 0x93, 0xb5, 0x64, 0xac, 0x84, 0x96, 0x41, 0xd3, 0x53, 0xff, 0x47, 0x52,
	0xb8, 0x3a, 0xa7, 0xf0, 0x84, 0x7c, 0xf6, 0xcc, 0x87, 0xc3, 0xeb, 0xd8, 0x50, 0x9f, 0x2c, 0x42,
	0xc2, 0x51, 0xf6, 0xff, 0xe0, 0xaf, 0xd6, 0xd3, 0xeb, 0x9a, 0xf3, 0xbb, 0x01, 0x30, 0x48, 0xe3,
	0x78, 0x9d, 0x44, 0xb8, 0x86, 0xec, 0x19, 0xee, 0x6a, 0xf0, 0x0e, 0x37, 0x5e, 0xf8, 0xb3, 0x55,
	0x41, 0x51, 0xe7, 0x0d, 0xa5, 0x1a, 0x89, 0x01, 0x8e, 0x83, 0x0d, 0xd5, 0x64, 0x1d, 0xbf, 0x15,
	0x2b, 0xdb, 0x54, 0xec, 0x6f, 0x1e, 0xf1, 0x8d, 0x8c, 0xfb, 0x51, 0x14, 0xba, 0xaa, 0x0a, 0x8d,
	0x7a, 0x92, 0x94, 0x36, 0x0c, 0x64, 0x40, 0x87, 0xa9, 0xa5, 0xb4, 0x4a, 0xc2, 0x71, 0xa8, 0x4a,
	0x11, 0x2f, 0xfd, 0x19, 0x9d, 0x23, 0x0d, 0xf5, 0xa6, 0x92, 0x07, 0x8a, 0xfe, 0x5e, 0x44, 0xf3,
	0x7b, 0x49, 0x6b, 0xaa, 0x2b, 0xfa, 0x42, 0xc6, 0xc1, 0x35, 0x65, 0x1a, 0x06, 0xb9, 0x0d, 0x74,
	0x13, 0x1f, 0x97, 0xb5, 0x19, 0x06, 0x79, 0x46, 0x04, 0xca, 0xca, 0x4e, 0xc0, 0x8c, 0x83, 0xfc,
	0xad, 0xb0, 0x9b, 0x2a, 0x72, 0xa5, 0x27, 0x51, 0xe9, 0x43, 0xb1, 0xc0, 0x40, 0xd4, 0x01, 0xf9,
	0x40, 0xe9, 0x49, 0x64, 0x0e, 0x18, 0x78, 0x66, 0xec, 0xc7, 0xef, 0x5b, 0x4a, 0x44, 0x29, 0x23,
	0xfb, 0x66, 0xf7, 0x7e, 0x1e, 0xd1, 0xfd, 0x7c, 0x5a, 0x22, 0xb7, 0xb5, 0xdb, 0x9e, 0x50, 0x74,
	0xdc, 0x7a, 0xb4, 0x3f, 0xd9, 0x3d, 0x46, 0x27, 0x50, 0x8d, 0x05, 0xd5, 0xef, 0xb8, 0xb8, 0x58,
	0x85, 0xd4, 0xae, 0x81, 0x39, 0x54, 0x01, 0xf5, 0xf1, 0x03, 0xe9, 0xd2, 0xe4, 0xec, 0x19, 0xd4,
	0x36, 0xe7, 0x5e, 0x8d, 0x79, 0x71, 0xf0, 0x71, 0xe4, 0xf1, 0x28, 0x5c, 0x78, 0x3f, 0xe2, 0xbc,
	0x77, 0xa1, 0xa2, 0x52, 0x57, 0xc6, 0xd1, 0x64, 0x3c, 0x3c, 0xff, 0x01, 0x8d, 0x4d, 0xa8, 0x4d,
	0x6f, 0xbd, 0x1b, 0x25, 0xe8, 0xea, 0x6a, 0x5c, 0xdf, 0x8e, 0x87, 0x57, 0x96, 0xd6, 0xc6, 0xfd,
	0xee, 0x75, 0xc0, 0x90, 0xc1, 0xfc, 0x60, 0x5e, 0xe7, 0x14, 0x86, 0x32, 0xf5, 0x06, 0xff, 0x8d,
	0xe4, 0x3e, 0xe6, 0x27, 0xaa, 0xce, 0xd3, 0x87, 0x83, 0xfa, 0xfe, 0x99, 0xfc, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0xad, 0x0f, 0xe3, 0x33, 0x07, 0x00, 0x00,
}
