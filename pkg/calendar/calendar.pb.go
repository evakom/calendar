// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package calendar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventAlertBefore_BeforeType int32

const (
	EventAlertBefore_DAYS    EventAlertBefore_BeforeType = 0
	EventAlertBefore_HOURS   EventAlertBefore_BeforeType = 1
	EventAlertBefore_MINUTES EventAlertBefore_BeforeType = 2
)

var EventAlertBefore_BeforeType_name = map[int32]string{
	0: "DAYS",
	1: "HOURS",
	2: "MINUTES",
}

var EventAlertBefore_BeforeType_value = map[string]int32{
	"DAYS":    0,
	"HOURS":   1,
	"MINUTES": 2,
}

func (x EventAlertBefore_BeforeType) String() string {
	return proto.EnumName(EventAlertBefore_BeforeType_name, int32(x))
}

func (EventAlertBefore_BeforeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0, 0, 0}
}

type Event struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	OccursAt             *timestamp.Timestamp `protobuf:"bytes,5,opt,name=occurs_at,json=occursAt,proto3" json:"occurs_at,omitempty"`
	Subject              string               `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string               `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	Duration             uint32               `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	Location             string               `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	User                 *User                `protobuf:"bytes,10,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Event) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Event) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Event) GetOccursAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccursAt
	}
	return nil
}

func (m *Event) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Event) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Event) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Event) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type EventAlertBefore struct {
	Type                 EventAlertBefore_BeforeType `protobuf:"varint,1,opt,name=type,proto3,enum=calendar.EventAlertBefore_BeforeType" json:"type,omitempty"`
	Before               uint32                      `protobuf:"varint,2,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EventAlertBefore) Reset()         { *m = EventAlertBefore{} }
func (m *EventAlertBefore) String() string { return proto.CompactTextString(m) }
func (*EventAlertBefore) ProtoMessage()    {}
func (*EventAlertBefore) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0, 0}
}

func (m *EventAlertBefore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventAlertBefore.Unmarshal(m, b)
}
func (m *EventAlertBefore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventAlertBefore.Marshal(b, m, deterministic)
}
func (m *EventAlertBefore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAlertBefore.Merge(m, src)
}
func (m *EventAlertBefore) XXX_Size() int {
	return xxx_messageInfo_EventAlertBefore.Size(m)
}
func (m *EventAlertBefore) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAlertBefore.DiscardUnknown(m)
}

var xxx_messageInfo_EventAlertBefore proto.InternalMessageInfo

func (m *EventAlertBefore) GetType() EventAlertBefore_BeforeType {
	if m != nil {
		return m.Type
	}
	return EventAlertBefore_DAYS
}

func (m *EventAlertBefore) GetBefore() uint32 {
	if m != nil {
		return m.Before
	}
	return 0
}

type User struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                []string             `protobuf:"bytes,3,rep,name=email,proto3" json:"email,omitempty"`
	Mobile               []string             `protobuf:"bytes,4,rep,name=mobile,proto3" json:"mobile,omitempty"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() []string {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *User) GetMobile() []string {
	if m != nil {
		return m.Mobile
	}
	return nil
}

func (m *User) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func init() {
	proto.RegisterEnum("calendar.EventAlertBefore_BeforeType", EventAlertBefore_BeforeType_name, EventAlertBefore_BeforeType_value)
	proto.RegisterType((*Event)(nil), "calendar.Event")
	proto.RegisterType((*EventAlertBefore)(nil), "calendar.Event.alertBefore")
	proto.RegisterType((*User)(nil), "calendar.User")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x1c, 0xc5, 0xe7, 0xc4, 0x4e, 0xec, 0x7f, 0x88, 0x09, 0x62, 0x0c, 0xe3, 0xcb, 0x86, 0x0f, 0x63,
	0x27, 0x05, 0x32, 0xd8, 0xd8, 0x6e, 0x19, 0x0b, 0xb4, 0x87, 0xb6, 0xe0, 0x24, 0x87, 0x9e, 0x8a,
	0x6c, 0x29, 0xa9, 0x8b, 0x6d, 0x19, 0x59, 0x2e, 0xe4, 0x5b, 0xb4, 0x9f, 0xa0, 0x5f, 0xb5, 0xb2,
	0x64, 0x27, 0x85, 0x1e, 0xd2, 0x93, 0xf5, 0xde, 0xff, 0xfd, 0xfe, 0x08, 0x3d, 0x83, 0x9f, 0x92,
	0x9c, 0x95, 0x94, 0x08, 0x5c, 0x09, 0x2e, 0x39, 0x72, 0x7b, 0x1d, 0x7e, 0xdd, 0x73, 0xbe, 0xcf,
	0xd9, 0x5c, 0xfb, 0x49, 0xb3, 0x9b, 0xcb, 0xac, 0x60, 0xb5, 0x24, 0x45, 0x65, 0xa2, 0xd1, 0x8b,
	0x0d, 0xce, 0xea, 0x91, 0x95, 0x12, 0xf9, 0x30, 0xc8, 0x68, 0x60, 0x7d, 0xb3, 0x7e, 0x4c, 0x63,
	0x75, 0x42, 0x7f, 0x00, 0x52, 0xc1, 0x88, 0x64, 0xf4, 0x8e, 0xc8, 0x60, 0xa0, 0xfc, 0xc9, 0x22,
	0xc4, 0x66, 0x1f, 0xee, 0xf7, 0xe1, 0x4d, 0xbf, 0x2f, 0xf6, 0xba, 0xf4, 0x52, 0xb6, 0x68, 0x53,
	0xd1, 0x1e, 0x1d, 0x9e, 0x47, 0xbb, 0xb4, 0x41, 0x29, 0xcb, 0x59, 0x87, 0xda, 0xe7, 0xd1, 0x2e,
	0xad, 0xd0, 0xdf, 0xe0, 0xf1, 0x34, 0x6d, 0x44, 0xdd, 0x92, 0xce, 0x59, 0xd2, 0x35, 0x61, 0x05,
	0x06, 0x30, 0xae, 0x9b, 0xe4, 0x81, 0xa5, 0x32, 0x18, 0x29, 0xcc, 0x8b, 0x7b, 0x89, 0x10, 0xd8,
	0x09, 0xa7, 0x87, 0x60, 0xac, 0x6d, 0x7d, 0x46, 0x21, 0xb8, 0xb4, 0x11, 0x44, 0x66, 0xbc, 0x0c,
	0x5c, 0xfd, 0x5a, 0x47, 0xdd, 0xce, 0x72, 0x9e, 0x9a, 0x99, 0xa7, 0x99, 0xa3, 0x46, 0x11, 0xd8,
	0x4d, 0xcd, 0x44, 0x00, 0xfa, 0x66, 0x3e, 0x3e, 0x76, 0xb6, 0x55, 0x6e, 0xac, 0x67, 0xe1, 0xb3,
	0x05, 0x13, 0x65, 0x0b, 0xf9, 0x8f, 0xed, 0xb8, 0x60, 0xe8, 0x2f, 0xd8, 0xf2, 0x50, 0x31, 0xdd,
	0x8a, 0xbf, 0xf8, 0x7e, 0x62, 0x74, 0x65, 0xf8, 0x4d, 0x14, 0x9b, 0xcf, 0x46, 0xa5, 0x63, 0xcd,
	0xa0, 0x2f, 0x30, 0x4a, 0xb4, 0xa7, 0xbb, 0x9b, 0xc6, 0x9d, 0x8a, 0x30, 0xc0, 0x29, 0x8b, 0x5c,
	0xb0, 0xff, 0x2f, 0x6f, 0xd7, 0xb3, 0x4f, 0xc8, 0x03, 0xe7, 0xe2, 0x66, 0x1b, 0xaf, 0x67, 0x16,
	0x9a, 0xc0, 0xf8, 0xea, 0xf2, 0x7a, 0xbb, 0x59, 0xad, 0x67, 0x83, 0xe8, 0xc9, 0x02, 0xbb, 0xbd,
	0xe2, 0xbb, 0x1f, 0x44, 0x3d, 0x4e, 0x49, 0x0a, 0xb3, 0x5e, 0x3d, 0x4e, 0x7b, 0x46, 0x9f, 0xc1,
	0x61, 0x05, 0xc9, 0x72, 0x55, 0xfa, 0x50, 0x99, 0x46, 0xb4, 0x57, 0x29, 0x78, 0x92, 0xe5, 0x4c,
	0x15, 0xda, 0xda, 0x9d, 0x42, 0xbf, 0xc0, 0x4d, 0x32, 0x21, 0xef, 0x29, 0x39, 0x7c, 0xa4, 0xb0,
	0x3e, 0x9b, 0x8c, 0xf4, 0xf4, 0xe7, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xb4, 0x45, 0xab,
	0xf8, 0x02, 0x00, 0x00,
}
