// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog.proto

package catalogv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VendorId             string   `protobuf:"bytes,2,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	Champions            []string `protobuf:"bytes,3,rep,name=champions,proto3" json:"champions,omitempty"`
	BlueEssence          int32    `protobuf:"varint,4,opt,name=blue_essence,json=blueEssence,proto3" json:"blue_essence,omitempty"`
	RiotPoints           int32    `protobuf:"varint,5,opt,name=riot_points,json=riotPoints,proto3" json:"riot_points,omitempty"`
	Solo                 int32    `protobuf:"varint,6,opt,name=solo,proto3" json:"solo,omitempty"`
	Flex                 int32    `protobuf:"varint,7,opt,name=flex,proto3" json:"flex,omitempty"`
	PriceDollars         int32    `protobuf:"varint,8,opt,name=price_dollars,json=priceDollars,proto3" json:"price_dollars,omitempty"`
	PriceCents           int32    `protobuf:"varint,9,opt,name=price_cents,json=priceCents,proto3" json:"price_cents,omitempty"`
	Level                int32    `protobuf:"varint,10,opt,name=level,proto3" json:"level,omitempty"`
	Email                string   `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	EmailPassword        string   `protobuf:"bytes,12,opt,name=email_password,json=emailPassword,proto3" json:"email_password,omitempty"`
	LoginName            string   `protobuf:"bytes,13,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
	LoginPassword        string   `protobuf:"bytes,14,opt,name=login_password,json=loginPassword,proto3" json:"login_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *Item) GetChampions() []string {
	if m != nil {
		return m.Champions
	}
	return nil
}

func (m *Item) GetBlueEssence() int32 {
	if m != nil {
		return m.BlueEssence
	}
	return 0
}

func (m *Item) GetRiotPoints() int32 {
	if m != nil {
		return m.RiotPoints
	}
	return 0
}

func (m *Item) GetSolo() int32 {
	if m != nil {
		return m.Solo
	}
	return 0
}

func (m *Item) GetFlex() int32 {
	if m != nil {
		return m.Flex
	}
	return 0
}

func (m *Item) GetPriceDollars() int32 {
	if m != nil {
		return m.PriceDollars
	}
	return 0
}

func (m *Item) GetPriceCents() int32 {
	if m != nil {
		return m.PriceCents
	}
	return 0
}

func (m *Item) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Item) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Item) GetEmailPassword() string {
	if m != nil {
		return m.EmailPassword
	}
	return ""
}

func (m *Item) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *Item) GetLoginPassword() string {
	if m != nil {
		return m.LoginPassword
	}
	return ""
}

type CreateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Item                 *Item    `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{3}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *RemoveRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{4}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

func (m *RemoveResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *RemoveResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type Specification struct {
	Solo                 int32    `protobuf:"varint,1,opt,name=solo,proto3" json:"solo,omitempty"`
	Flex                 int32    `protobuf:"varint,2,opt,name=flex,proto3" json:"flex,omitempty"`
	PriceDollars         int32    `protobuf:"varint,3,opt,name=price_dollars,json=priceDollars,proto3" json:"price_dollars,omitempty"`
	PriceCents           int32    `protobuf:"varint,4,opt,name=price_cents,json=priceCents,proto3" json:"price_cents,omitempty"`
	ItemId               string   `protobuf:"bytes,5,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	VendorId             string   `protobuf:"bytes,6,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	PageNum              int32    `protobuf:"varint,7,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Item                 *Item    `protobuf:"bytes,8,opt,name=item,proto3" json:"item,omitempty"`
	Api                  string   `protobuf:"bytes,9,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{5}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetSolo() int32 {
	if m != nil {
		return m.Solo
	}
	return 0
}

func (m *Specification) GetFlex() int32 {
	if m != nil {
		return m.Flex
	}
	return 0
}

func (m *Specification) GetPriceDollars() int32 {
	if m != nil {
		return m.PriceDollars
	}
	return 0
}

func (m *Specification) GetPriceCents() int32 {
	if m != nil {
		return m.PriceCents
	}
	return 0
}

func (m *Specification) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *Specification) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *Specification) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *Specification) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Specification) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type Response struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Items                []*Item  `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *Response) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ListRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{7}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ListRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Items                []*Item  `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{8}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ListResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetByIdRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Haha                 string   `protobuf:"bytes,3,opt,name=haha,proto3" json:"haha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIdRequest) Reset()         { *m = GetByIdRequest{} }
func (m *GetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetByIdRequest) ProtoMessage()    {}
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{9}
}

func (m *GetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIdRequest.Unmarshal(m, b)
}
func (m *GetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIdRequest.Merge(m, src)
}
func (m *GetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetByIdRequest.Size(m)
}
func (m *GetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIdRequest proto.InternalMessageInfo

func (m *GetByIdRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetByIdRequest) GetHaha() string {
	if m != nil {
		return m.Haha
	}
	return ""
}

type GetByIdResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Item                 *Item    `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Haha                 string   `protobuf:"bytes,3,opt,name=haha,proto3" json:"haha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIdResponse) Reset()         { *m = GetByIdResponse{} }
func (m *GetByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetByIdResponse) ProtoMessage()    {}
func (*GetByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{10}
}

func (m *GetByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIdResponse.Unmarshal(m, b)
}
func (m *GetByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIdResponse.Marshal(b, m, deterministic)
}
func (m *GetByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIdResponse.Merge(m, src)
}
func (m *GetByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetByIdResponse.Size(m)
}
func (m *GetByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIdResponse proto.InternalMessageInfo

func (m *GetByIdResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetByIdResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *GetByIdResponse) GetHaha() string {
	if m != nil {
		return m.Haha
	}
	return ""
}

func init() {
	proto.RegisterType((*Item)(nil), "catalogv1.Item")
	proto.RegisterType((*CreateRequest)(nil), "catalogv1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "catalogv1.CreateResponse")
	proto.RegisterType((*RemoveRequest)(nil), "catalogv1.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "catalogv1.RemoveResponse")
	proto.RegisterType((*Specification)(nil), "catalogv1.Specification")
	proto.RegisterType((*Response)(nil), "catalogv1.Response")
	proto.RegisterType((*ListRequest)(nil), "catalogv1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "catalogv1.ListResponse")
	proto.RegisterType((*GetByIdRequest)(nil), "catalogv1.GetByIdRequest")
	proto.RegisterType((*GetByIdResponse)(nil), "catalogv1.GetByIdResponse")
}

func init() { proto.RegisterFile("catalog.proto", fileDescriptor_0abbfcf058acdf89) }

var fileDescriptor_0abbfcf058acdf89 = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4f, 0x6f, 0x23, 0xc5,
	0x13, 0x95, 0x67, 0x1c, 0xdb, 0x53, 0xfe, 0x93, 0xdd, 0xce, 0x4f, 0x9b, 0x89, 0x7f, 0x01, 0xcc,
	0x2c, 0x2b, 0xa2, 0x88, 0x78, 0x36, 0x66, 0xb9, 0x44, 0x1c, 0x60, 0x03, 0x59, 0x59, 0x82, 0xd5,
	0x6a, 0x56, 0xe2, 0x00, 0x2b, 0x59, 0xed, 0x99, 0xca, 0xb8, 0x61, 0x66, 0x7a, 0x98, 0x6e, 0x3b,
	0xbb, 0x42, 0x5c, 0x90, 0xf8, 0x02, 0x70, 0x43, 0x7c, 0x1e, 0x0e, 0x5c, 0x39, 0x72, 0xe5, 0x83,
	0xa0, 0xee, 0x1e, 0x1b, 0x4f, 0xec, 0x84, 0x95, 0x38, 0xb9, 0xeb, 0x55, 0xf5, 0xeb, 0xe7, 0xaa,
	0x57, 0x36, 0x74, 0x43, 0x2a, 0x69, 0xc2, 0xe3, 0x61, 0x5e, 0x70, 0xc9, 0x89, 0x53, 0x86, 0x8b,
	0xd3, 0xfe, 0x61, 0xcc, 0x79, 0x9c, 0xa0, 0x4f, 0x73, 0xe6, 0xd3, 0x2c, 0xe3, 0x92, 0x4a, 0xc6,
	0x33, 0x61, 0x0a, 0xfb, 0xef, 0xe9, 0x8f, 0xf0, 0x24, 0xc6, 0xec, 0x44, 0x5c, 0xd1, 0x38, 0xc6,
	0xc2, 0xe7, 0xb9, 0xae, 0xd8, 0xac, 0xf6, 0x7e, 0xb5, 0xa1, 0x3e, 0x96, 0x98, 0x92, 0x1e, 0x58,
	0x2c, 0x72, 0x6b, 0x83, 0xda, 0x91, 0x13, 0x58, 0x2c, 0x22, 0xff, 0x07, 0x67, 0x81, 0x59, 0xc4,
	0x8b, 0x09, 0x8b, 0x5c, 0x4b, 0xc3, 0x2d, 0x03, 0x8c, 0x23, 0x72, 0x08, 0x4e, 0x38, 0xa3, 0x69,
	0xae, 0x88, 0x5c, 0x7b, 0x60, 0x1f, 0x39, 0xc1, 0x3f, 0x00, 0x79, 0x1b, 0x3a, 0xd3, 0x64, 0x8e,
	0x13, 0x14, 0x02, 0xb3, 0x10, 0xdd, 0xfa, 0xa0, 0x76, 0xb4, 0x13, 0xb4, 0x15, 0xf6, 0xa9, 0x81,
	0xc8, 0x5b, 0xd0, 0x2e, 0x18, 0x97, 0x93, 0x9c, 0xb3, 0x4c, 0x0a, 0x77, 0x47, 0x57, 0x80, 0x82,
	0x9e, 0x69, 0x84, 0x10, 0xa8, 0x0b, 0x9e, 0x70, 0xb7, 0xa1, 0x33, 0xfa, 0xac, 0xb0, 0xcb, 0x04,
	0x5f, 0xba, 0x4d, 0x83, 0xa9, 0x33, 0xb9, 0x0f, 0xdd, 0xbc, 0x60, 0x21, 0x4e, 0x22, 0x9e, 0x24,
	0xb4, 0x10, 0x6e, 0x4b, 0x27, 0x3b, 0x1a, 0xfc, 0xc4, 0x60, 0xea, 0x35, 0x53, 0x14, 0xa2, 0x7a,
	0xcd, 0x31, 0xaf, 0x69, 0xe8, 0x5c, 0x21, 0xe4, 0x7f, 0xb0, 0x93, 0xe0, 0x02, 0x13, 0x17, 0x74,
	0xca, 0x04, 0x0a, 0xc5, 0x94, 0xb2, 0xc4, 0x6d, 0xeb, 0xaf, 0x6f, 0x02, 0xf2, 0x00, 0x7a, 0xfa,
	0x30, 0xc9, 0xa9, 0x10, 0x57, 0xbc, 0x88, 0xdc, 0x8e, 0x4e, 0x77, 0x35, 0xfa, 0xac, 0x04, 0xc9,
	0x1b, 0x00, 0x09, 0x8f, 0x59, 0x36, 0xc9, 0x68, 0x8a, 0x6e, 0x57, 0x97, 0x38, 0x1a, 0x79, 0x4a,
	0x53, 0x54, 0x2c, 0x26, 0xbd, 0x62, 0xe9, 0x19, 0x16, 0x8d, 0x2e, 0x59, 0xbc, 0x0b, 0xe8, 0x9e,
	0x17, 0x48, 0x25, 0x06, 0xf8, 0xed, 0x1c, 0x85, 0x24, 0x77, 0xc0, 0xa6, 0x39, 0x2b, 0xe7, 0xa4,
	0x8e, 0xe4, 0x3e, 0xd4, 0x99, 0xc4, 0x54, 0xcf, 0xa8, 0x3d, 0xda, 0x1d, 0xae, 0x7c, 0x32, 0x54,
	0x73, 0x0d, 0x74, 0xd2, 0x1b, 0x41, 0x6f, 0xc9, 0x23, 0x72, 0x9e, 0x09, 0xdc, 0x42, 0x64, 0x1c,
	0xa0, 0x68, 0x6c, 0xe5, 0x00, 0xef, 0x14, 0xba, 0x01, 0xa6, 0x7c, 0x71, 0xcb, 0xdb, 0xd7, 0xaf,
	0x7c, 0x08, 0xbd, 0xe5, 0x95, 0x1b, 0x9f, 0x71, 0xa1, 0x19, 0x61, 0x82, 0x12, 0x97, 0x17, 0x97,
	0xa1, 0xf7, 0xa3, 0x05, 0xdd, 0xe7, 0x39, 0x86, 0xec, 0x92, 0x85, 0xda, 0xa4, 0x2b, 0x17, 0xd4,
	0xb6, 0xb8, 0xc0, 0xba, 0xcd, 0x05, 0xf6, 0xbf, 0xbb, 0xa0, 0xbe, 0xe1, 0x82, 0x7d, 0x68, 0xaa,
	0x66, 0x29, 0xc3, 0xef, 0x68, 0xbd, 0x0d, 0x15, 0x8e, 0xaf, 0xed, 0x42, 0xe3, 0xda, 0x2e, 0x1c,
	0x40, 0x2b, 0xa7, 0x31, 0x4e, 0xb2, 0x79, 0x5a, 0x3a, 0xb3, 0xa9, 0xe2, 0xa7, 0xf3, 0x74, 0x35,
	0x9a, 0xd6, 0x2d, 0xa3, 0x59, 0x76, 0xc8, 0x59, 0x75, 0xc8, 0x3b, 0x87, 0xd6, 0x2d, 0xfd, 0x7b,
	0x00, 0x3b, 0xea, 0x9e, 0x70, 0xad, 0x81, 0xbd, 0x8d, 0xd5, 0x64, 0xbd, 0x31, 0xb4, 0x3f, 0x63,
	0x42, 0xde, 0x3c, 0x3b, 0x02, 0x75, 0xa5, 0xb3, 0x1c, 0x82, 0x3e, 0xeb, 0x3d, 0x60, 0x29, 0x93,
	0xba, 0x7f, 0x76, 0x60, 0x02, 0xef, 0x09, 0x74, 0x0c, 0xd5, 0x7f, 0xd5, 0x74, 0x01, 0xbd, 0x27,
	0x28, 0x1f, 0xbf, 0x1a, 0x47, 0xaf, 0x6d, 0x29, 0x25, 0x73, 0x46, 0x67, 0x54, 0x2b, 0x72, 0x02,
	0x7d, 0xf6, 0x5e, 0xc0, 0xee, 0x8a, 0xe7, 0x46, 0x4d, 0xaf, 0xb3, 0x17, 0xdb, 0xd8, 0x47, 0x7f,
	0xda, 0xd0, 0x3b, 0x37, 0xc5, 0xcf, 0xb1, 0x58, 0xb0, 0x10, 0xc9, 0x97, 0xe0, 0x5c, 0xb0, 0x2c,
	0x52, 0x17, 0x05, 0x71, 0xd7, 0xa8, 0x2a, 0x76, 0xed, 0xef, 0xad, 0x65, 0x96, 0xca, 0xbc, 0x37,
	0x7f, 0xf8, 0xe3, 0xaf, 0x9f, 0x2d, 0xd7, 0xdb, 0xf3, 0x17, 0xa7, 0x7e, 0x99, 0x17, 0xbe, 0x40,
	0x5a, 0x84, 0xb3, 0xb3, 0xda, 0x31, 0xf9, 0x02, 0x1a, 0x66, 0x35, 0x2b, 0xc4, 0x95, 0xad, 0xef,
	0x1f, 0x6c, 0xc9, 0x94, 0xf4, 0xfb, 0x9a, 0xfe, 0xae, 0xd7, 0x59, 0xa7, 0x57, 0xbc, 0x2f, 0x00,
	0xcc, 0x2e, 0xea, 0x9f, 0x77, 0xb7, 0x22, 0x6d, 0x6d, 0xab, 0x2b, 0xdc, 0xd5, 0xe5, 0xf5, 0x0e,
	0x34, 0xf7, 0xde, 0xf1, 0xdd, 0x8a, 0xf4, 0xef, 0x58, 0xf4, 0xbd, 0xea, 0x88, 0xf2, 0x84, 0xe9,
	0xc8, 0xbd, 0x35, 0x8a, 0x35, 0xd3, 0xf5, 0xf7, 0x37, 0xf0, 0x92, 0xf8, 0x50, 0x13, 0xdf, 0xf3,
	0xaa, 0xc4, 0x09, 0x13, 0x52, 0x29, 0xff, 0x0a, 0x9a, 0xe5, 0x78, 0xc9, 0xba, 0xb8, 0xaa, 0x75,
	0xfa, 0xfd, 0x6d, 0xa9, 0xaa, 0x70, 0xb2, 0x29, 0xfc, 0xf1, 0xef, 0xb5, 0x9f, 0x3e, 0xfe, 0xad,
	0x46, 0x10, 0x76, 0xcb, 0x19, 0x0f, 0x84, 0x19, 0xb2, 0xf7, 0x39, 0xec, 0x86, 0x55, 0x88, 0x1c,
	0xcf, 0xa4, 0xcc, 0xc5, 0x99, 0xef, 0xc7, 0x4c, 0xce, 0xe6, 0xd3, 0x61, 0xc8, 0x53, 0x3f, 0xfc,
	0x66, 0x3a, 0xa5, 0x49, 0xe2, 0x8b, 0x74, 0x5e, 0x5c, 0xb2, 0xec, 0xa4, 0xbc, 0xd3, 0xef, 0x4d,
	0x13, 0x3a, 0xfb, 0x68, 0x46, 0xf3, 0xfc, 0xd5, 0x90, 0x17, 0xf1, 0xc8, 0x3e, 0x1d, 0x3e, 0x3c,
	0xb6, 0x6a, 0xd6, 0xe8, 0x0e, 0xcd, 0xf3, 0xa4, 0xb4, 0x88, 0xff, 0xb5, 0xe0, 0xd9, 0xd9, 0x06,
	0x12, 0x7c, 0x00, 0xf6, 0xa3, 0x87, 0x8f, 0xc8, 0x10, 0xde, 0x09, 0x50, 0xce, 0x8b, 0x0c, 0xa3,
	0xc1, 0xd5, 0x0c, 0xb3, 0x41, 0x81, 0x82, 0xcf, 0x8b, 0x10, 0x07, 0x11, 0x47, 0x91, 0xbd, 0x2b,
	0x07, 0xf8, 0x92, 0x09, 0x49, 0x1a, 0x50, 0xff, 0xc5, 0xaa, 0x35, 0xa7, 0x0d, 0xfd, 0x27, 0xfe,
	0xfe, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0x7f, 0xa3, 0x6e, 0x2c, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	FindItems(ctx context.Context, in *Specification, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	RemoveItem(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	ListItems(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
}

type catalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatalogServiceClient(cc *grpc.ClientConn) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) FindItems(ctx context.Context, in *Specification, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/catalogv1.CatalogService/FindItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/catalogv1.CatalogService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) RemoveItem(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/catalogv1.CatalogService/RemoveItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListItems(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/catalogv1.CatalogService/ListItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, "/catalogv1.CatalogService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	FindItems(context.Context, *Specification) (*Response, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	RemoveItem(context.Context, *RemoveRequest) (*RemoveResponse, error)
	ListItems(context.Context, *ListRequest) (*ListResponse, error)
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
}

// UnimplementedCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (*UnimplementedCatalogServiceServer) FindItems(ctx context.Context, req *Specification) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindItems not implemented")
}
func (*UnimplementedCatalogServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCatalogServiceServer) RemoveItem(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (*UnimplementedCatalogServiceServer) ListItems(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (*UnimplementedCatalogServiceServer) GetById(ctx context.Context, req *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_FindItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Specification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).FindItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogv1.CatalogService/FindItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).FindItems(ctx, req.(*Specification))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogv1.CatalogService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogv1.CatalogService/RemoveItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).RemoveItem(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogv1.CatalogService/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListItems(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogv1.CatalogService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "catalogv1.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindItems",
			Handler:    _CatalogService_FindItems_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CatalogService_Create_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _CatalogService_RemoveItem_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _CatalogService_ListItems_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _CatalogService_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}
