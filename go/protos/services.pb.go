// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/services.proto

package modelzoo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PayloadType int32

const (
	PayloadType_IMAGE PayloadType = 0
	PayloadType_TEXT  PayloadType = 1
	PayloadType_TABLE PayloadType = 2
)

var PayloadType_name = map[int32]string{
	0: "IMAGE",
	1: "TEXT",
	2: "TABLE",
}

var PayloadType_value = map[string]int32{
	"IMAGE": 0,
	"TEXT":  1,
	"TABLE": 2,
}

func (x PayloadType) String() string {
	return proto.EnumName(PayloadType_name, int32(x))
}

func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{0}
}

// Image is repsented by the dataurl format
//https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URIs
type Image struct {
	Metadata             map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ImageDataUrl         string            `protobuf:"bytes,2,opt,name=image_data_url,json=imageDataUrl,proto3" json:"image_data_url,omitempty"`
	ModelName            string            `protobuf:"bytes,3,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	AccessToken          string            `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Image) GetImageDataUrl() string {
	if m != nil {
		return m.ImageDataUrl
	}
	return ""
}

func (m *Image) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *Image) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

// Text is a list of string
type Text struct {
	Metadata             map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Texts                []string          `protobuf:"bytes,2,rep,name=texts,proto3" json:"texts,omitempty"`
	ModelName            string            `protobuf:"bytes,3,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	AccessToken          string            `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{1}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Text) GetTexts() []string {
	if m != nil {
		return m.Texts
	}
	return nil
}

func (m *Text) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *Text) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

// Table is more complicated.
//It is inspired from pandas orient="index"
//>>> df.to_json(orient='index')
//'{"row 1":{"col 1":"a","col 2":"b"},"row 2":{"col 1":"c","col 2":"d"}}'
type Table struct {
	Metadata             map[string]string     `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ModelName            string                `protobuf:"bytes,2,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	AccessToken          string                `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Table                map[string]*Table_Row `protobuf:"bytes,4,rep,name=table,proto3" json:"table,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ColumnNames          []string              `protobuf:"bytes,5,rep,name=column_names,json=columnNames,proto3" json:"column_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{2}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Table) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *Table) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Table) GetTable() map[string]*Table_Row {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *Table) GetColumnNames() []string {
	if m != nil {
		return m.ColumnNames
	}
	return nil
}

type Table_Row struct {
	ColumnToValue        map[string]string `protobuf:"bytes,1,rep,name=column_to_value,json=columnToValue,proto3" json:"column_to_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Table_Row) Reset()         { *m = Table_Row{} }
func (m *Table_Row) String() string { return proto.CompactTextString(m) }
func (*Table_Row) ProtoMessage()    {}
func (*Table_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{2, 1}
}

func (m *Table_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table_Row.Unmarshal(m, b)
}
func (m *Table_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table_Row.Marshal(b, m, deterministic)
}
func (m *Table_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table_Row.Merge(m, src)
}
func (m *Table_Row) XXX_Size() int {
	return xxx_messageInfo_Table_Row.Size(m)
}
func (m *Table_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Table_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Table_Row proto.InternalMessageInfo

func (m *Table_Row) GetColumnToValue() map[string]string {
	if m != nil {
		return m.ColumnToValue
	}
	return nil
}

// Web
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type KVPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVPair) Reset()         { *m = KVPair{} }
func (m *KVPair) String() string { return proto.CompactTextString(m) }
func (*KVPair) ProtoMessage()    {}
func (*KVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{4}
}

func (m *KVPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVPair.Unmarshal(m, b)
}
func (m *KVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVPair.Marshal(b, m, deterministic)
}
func (m *KVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVPair.Merge(m, src)
}
func (m *KVPair) XXX_Size() int {
	return xxx_messageInfo_KVPair.Size(m)
}
func (m *KVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KVPair.DiscardUnknown(m)
}

var xxx_messageInfo_KVPair proto.InternalMessageInfo

func (m *KVPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Model struct {
	ModelName            string    `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	Metadata             []*KVPair `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{5}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *Model) GetMetadata() []*KVPair {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type User struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{6}
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

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RateLimitToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitToken) Reset()         { *m = RateLimitToken{} }
func (m *RateLimitToken) String() string { return proto.CompactTextString(m) }
func (*RateLimitToken) ProtoMessage()    {}
func (*RateLimitToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{7}
}

func (m *RateLimitToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitToken.Unmarshal(m, b)
}
func (m *RateLimitToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitToken.Marshal(b, m, deterministic)
}
func (m *RateLimitToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitToken.Merge(m, src)
}
func (m *RateLimitToken) XXX_Size() int {
	return xxx_messageInfo_RateLimitToken.Size(m)
}
func (m *RateLimitToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitToken.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitToken proto.InternalMessageInfo

func (m *RateLimitToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ListModelsResponse struct {
	Models               []*Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListModelsResponse) Reset()         { *m = ListModelsResponse{} }
func (m *ListModelsResponse) String() string { return proto.CompactTextString(m) }
func (*ListModelsResponse) ProtoMessage()    {}
func (*ListModelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{8}
}

func (m *ListModelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListModelsResponse.Unmarshal(m, b)
}
func (m *ListModelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListModelsResponse.Marshal(b, m, deterministic)
}
func (m *ListModelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListModelsResponse.Merge(m, src)
}
func (m *ListModelsResponse) XXX_Size() int {
	return xxx_messageInfo_ListModelsResponse.Size(m)
}
func (m *ListModelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListModelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListModelsResponse proto.InternalMessageInfo

func (m *ListModelsResponse) GetModels() []*Model {
	if m != nil {
		return m.Models
	}
	return nil
}

// Downloader
type ImageDownloadRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageDownloadRequest) Reset()         { *m = ImageDownloadRequest{} }
func (m *ImageDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*ImageDownloadRequest) ProtoMessage()    {}
func (*ImageDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{9}
}

func (m *ImageDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageDownloadRequest.Unmarshal(m, b)
}
func (m *ImageDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageDownloadRequest.Marshal(b, m, deterministic)
}
func (m *ImageDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageDownloadRequest.Merge(m, src)
}
func (m *ImageDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_ImageDownloadRequest.Size(m)
}
func (m *ImageDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImageDownloadRequest proto.InternalMessageInfo

func (m *ImageDownloadRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ImageDownloadResponse struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageDownloadResponse) Reset()         { *m = ImageDownloadResponse{} }
func (m *ImageDownloadResponse) String() string { return proto.CompactTextString(m) }
func (*ImageDownloadResponse) ProtoMessage()    {}
func (*ImageDownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{10}
}

func (m *ImageDownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageDownloadResponse.Unmarshal(m, b)
}
func (m *ImageDownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageDownloadResponse.Marshal(b, m, deterministic)
}
func (m *ImageDownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageDownloadResponse.Merge(m, src)
}
func (m *ImageDownloadResponse) XXX_Size() int {
	return xxx_messageInfo_ImageDownloadResponse.Size(m)
}
func (m *ImageDownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageDownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageDownloadResponse proto.InternalMessageInfo

func (m *ImageDownloadResponse) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type Payload struct {
	Type PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=modelzoo.PayloadType" json:"type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Payload_Image
	//	*Payload_Text
	//	*Payload_Table
	Payload              isPayload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_edd031a16b9fa047, []int{11}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetType() PayloadType {
	if m != nil {
		return m.Type
	}
	return PayloadType_IMAGE
}

type isPayload_Payload interface {
	isPayload_Payload()
}

type Payload_Image struct {
	Image *Image `protobuf:"bytes,2,opt,name=image,proto3,oneof"`
}

type Payload_Text struct {
	Text *Text `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type Payload_Table struct {
	Table *Table `protobuf:"bytes,4,opt,name=table,proto3,oneof"`
}

func (*Payload_Image) isPayload_Payload() {}

func (*Payload_Text) isPayload_Payload() {}

func (*Payload_Table) isPayload_Payload() {}

func (m *Payload) GetPayload() isPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Payload) GetImage() *Image {
	if x, ok := m.GetPayload().(*Payload_Image); ok {
		return x.Image
	}
	return nil
}

func (m *Payload) GetText() *Text {
	if x, ok := m.GetPayload().(*Payload_Text); ok {
		return x.Text
	}
	return nil
}

func (m *Payload) GetTable() *Table {
	if x, ok := m.GetPayload().(*Payload_Table); ok {
		return x.Table
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Payload_Image)(nil),
		(*Payload_Text)(nil),
		(*Payload_Table)(nil),
	}
}

func init() {
	proto.RegisterEnum("modelzoo.PayloadType", PayloadType_name, PayloadType_value)
	proto.RegisterType((*Image)(nil), "modelzoo.Image")
	proto.RegisterMapType((map[string]string)(nil), "modelzoo.Image.MetadataEntry")
	proto.RegisterType((*Text)(nil), "modelzoo.Text")
	proto.RegisterMapType((map[string]string)(nil), "modelzoo.Text.MetadataEntry")
	proto.RegisterType((*Table)(nil), "modelzoo.Table")
	proto.RegisterMapType((map[string]string)(nil), "modelzoo.Table.MetadataEntry")
	proto.RegisterMapType((map[string]*Table_Row)(nil), "modelzoo.Table.TableEntry")
	proto.RegisterType((*Table_Row)(nil), "modelzoo.Table.Row")
	proto.RegisterMapType((map[string]string)(nil), "modelzoo.Table.Row.ColumnToValueEntry")
	proto.RegisterType((*Empty)(nil), "modelzoo.Empty")
	proto.RegisterType((*KVPair)(nil), "modelzoo.KVPair")
	proto.RegisterType((*Model)(nil), "modelzoo.Model")
	proto.RegisterType((*User)(nil), "modelzoo.User")
	proto.RegisterType((*RateLimitToken)(nil), "modelzoo.RateLimitToken")
	proto.RegisterType((*ListModelsResponse)(nil), "modelzoo.ListModelsResponse")
	proto.RegisterType((*ImageDownloadRequest)(nil), "modelzoo.ImageDownloadRequest")
	proto.RegisterType((*ImageDownloadResponse)(nil), "modelzoo.ImageDownloadResponse")
	proto.RegisterType((*Payload)(nil), "modelzoo.Payload")
}

func init() { proto.RegisterFile("protos/services.proto", fileDescriptor_edd031a16b9fa047) }

var fileDescriptor_edd031a16b9fa047 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdd, 0x4e, 0xe3, 0x46,
	0x14, 0xb6, 0xe3, 0x18, 0x92, 0x13, 0x48, 0xd2, 0x69, 0x90, 0x2c, 0xab, 0xb4, 0xd4, 0x42, 0xfc,
	0x54, 0x25, 0xd0, 0x70, 0x51, 0xda, 0x0a, 0xa9, 0xfc, 0x44, 0x80, 0x4a, 0x10, 0x72, 0x0d, 0xea,
	0x5d, 0x34, 0x24, 0xd3, 0x2a, 0xc2, 0x3f, 0xa9, 0x67, 0x42, 0x48, 0x5f, 0xa4, 0x8f, 0xd2, 0x27,
	0xe9, 0x33, 0xec, 0xf5, 0x3e, 0xc1, 0xae, 0xe6, 0x8c, 0x83, 0x1d, 0x07, 0xc4, 0xae, 0xf6, 0x62,
	0x6f, 0xa2, 0x9c, 0x33, 0xdf, 0x7c, 0xe7, 0x3b, 0x67, 0xbe, 0x19, 0xc3, 0xca, 0x30, 0x8e, 0x44,
	0xc4, 0x77, 0x39, 0x8b, 0x1f, 0x06, 0x3d, 0xc6, 0x9b, 0x18, 0x93, 0x52, 0x10, 0xf5, 0x99, 0xff,
	0x4f, 0x14, 0x39, 0x6f, 0x74, 0x30, 0x2f, 0x02, 0xfa, 0x17, 0x23, 0x3f, 0x41, 0x29, 0x60, 0x82,
	0xf6, 0xa9, 0xa0, 0x96, 0xbe, 0x66, 0x6c, 0x55, 0x5a, 0xab, 0xcd, 0x29, 0xac, 0x89, 0x90, 0x66,
	0x27, 0x59, 0x6f, 0x87, 0x22, 0x9e, 0xb8, 0x4f, 0x70, 0xb2, 0x0e, 0xd5, 0x81, 0x04, 0x74, 0x65,
	0xd4, 0x1d, 0xc5, 0xbe, 0x55, 0x58, 0xd3, 0xb7, 0xca, 0xee, 0x12, 0x66, 0x4f, 0xa9, 0xa0, 0x37,
	0xb1, 0x4f, 0x56, 0x01, 0x90, 0xaf, 0x1b, 0xd2, 0x80, 0x59, 0x06, 0x22, 0xca, 0x98, 0xb9, 0xa2,
	0x01, 0x23, 0xdf, 0xc2, 0x12, 0xed, 0xf5, 0x18, 0xe7, 0x5d, 0x11, 0xdd, 0xb3, 0xd0, 0x2a, 0x22,
	0xa0, 0xa2, 0x72, 0x9e, 0x4c, 0xd9, 0xbf, 0xc0, 0xf2, 0x8c, 0x04, 0x52, 0x07, 0xe3, 0x9e, 0x4d,
	0x2c, 0x1d, 0xa1, 0xf2, 0x2f, 0x69, 0x80, 0xf9, 0x40, 0xfd, 0x11, 0x4b, 0x14, 0xa8, 0xe0, 0xe7,
	0xc2, 0x81, 0xee, 0xfc, 0xaf, 0x43, 0xd1, 0x63, 0x8f, 0x82, 0x1c, 0xcc, 0x35, 0xfa, 0x55, 0xda,
	0xa8, 0x44, 0xbc, 0xd8, 0x67, 0x03, 0x4c, 0xc1, 0x1e, 0x05, 0xb7, 0x0a, 0x6b, 0x86, 0x24, 0xc7,
	0xe0, 0x73, 0xf7, 0xf5, 0xce, 0x00, 0xd3, 0xa3, 0x77, 0xfe, 0x2b, 0x27, 0x88, 0x90, 0x17, 0x3b,
	0x9b, 0xed, 0xa1, 0xf0, 0x5a, 0x0f, 0xc6, 0x5c, 0x0f, 0x64, 0x0f, 0x4c, 0x21, 0x4b, 0x58, 0x45,
	0xac, 0x6c, 0xe7, 0x2b, 0xe3, 0xaf, 0x2a, 0xab, 0x80, 0x92, 0xb4, 0x17, 0xf9, 0xa3, 0x20, 0xc4,
	0xa2, 0xdc, 0x32, 0x71, 0xa8, 0x15, 0x95, 0x93, 0x65, 0xf9, 0x27, 0x0d, 0xc6, 0xfe, 0x57, 0x07,
	0xc3, 0x8d, 0xc6, 0xe4, 0x0a, 0x6a, 0x49, 0x1d, 0x11, 0x75, 0x15, 0x56, 0x4d, 0x67, 0x23, 0xaf,
	0xd1, 0x8d, 0xc6, 0xcd, 0x13, 0x84, 0x7a, 0xd1, 0xad, 0x04, 0x2a, 0xbd, 0xcb, 0xbd, 0x6c, 0xce,
	0xfe, 0x15, 0xc8, 0x3c, 0xe8, 0xa3, 0x94, 0x75, 0x00, 0xd2, 0x71, 0x3c, 0xb3, 0x73, 0x3b, 0xbb,
	0xb3, 0xd2, 0xfa, 0xf2, 0x19, 0x9d, 0x59, 0x07, 0x2c, 0x82, 0xd9, 0x0e, 0x86, 0x62, 0xe2, 0xec,
	0xc1, 0xc2, 0x6f, 0xb7, 0xd7, 0x74, 0x10, 0x7f, 0xa8, 0x1a, 0xc7, 0x03, 0xb3, 0x23, 0xb9, 0x73,
	0x06, 0xd0, 0xf3, 0x06, 0xf8, 0x3e, 0x63, 0x2d, 0x03, 0x87, 0x57, 0x4f, 0x45, 0xa9, 0x9a, 0xa9,
	0x9b, 0x9c, 0x03, 0x28, 0xde, 0x70, 0x16, 0xcb, 0x9a, 0x2c, 0xa0, 0x03, 0x3f, 0xe1, 0x53, 0x01,
	0xb1, 0xa1, 0x34, 0xa4, 0x9c, 0x8f, 0xa3, 0xb8, 0x9f, 0x88, 0x79, 0x8a, 0x9d, 0x0d, 0xa8, 0xba,
	0x54, 0xb0, 0xcb, 0x41, 0x30, 0x10, 0xca, 0x57, 0xf2, 0xce, 0xa1, 0xe7, 0x12, 0x0e, 0x0c, 0x9c,
	0x43, 0x20, 0x97, 0x03, 0x2e, 0x50, 0x3b, 0x77, 0x19, 0x1f, 0x46, 0x21, 0x67, 0x64, 0x13, 0x16,
	0x50, 0x14, 0x4f, 0x0e, 0xb8, 0x96, 0x6a, 0x44, 0xa4, 0x9b, 0x2c, 0x3b, 0x5b, 0xd0, 0xc0, 0x17,
	0xed, 0x34, 0x1a, 0x87, 0x7e, 0x44, 0xfb, 0x2e, 0xfb, 0x7b, 0xc4, 0xb8, 0x90, 0x63, 0x93, 0xaf,
	0x57, 0x32, 0xb6, 0x51, 0xec, 0x3b, 0x3b, 0xb0, 0x92, 0x43, 0x26, 0xb5, 0x1a, 0x60, 0xe2, 0xeb,
	0x36, 0xd5, 0x85, 0x81, 0xf3, 0x9f, 0x0e, 0x8b, 0xd7, 0x74, 0x22, 0x91, 0x64, 0x1b, 0x8a, 0x62,
	0x32, 0x54, 0x80, 0x6a, 0x6b, 0x25, 0xd5, 0x92, 0x00, 0xbc, 0xc9, 0x90, 0xb9, 0x08, 0x21, 0x9b,
	0x53, 0x32, 0x75, 0xe0, 0xb5, 0xdc, 0xc3, 0x7b, 0xae, 0x25, 0xfc, 0x64, 0x1d, 0x8a, 0xf2, 0xd1,
	0xc1, 0x0b, 0x58, 0x69, 0x55, 0x67, 0xdf, 0xad, 0x73, 0xcd, 0xc5, 0x55, 0x49, 0x37, 0xbd, 0x8b,
	0x39, 0x3a, 0xf4, 0x8f, 0xa4, 0xc3, 0xf5, 0xe3, 0x32, 0x2c, 0x0e, 0x95, 0x98, 0xef, 0x76, 0xa0,
	0x92, 0xd1, 0x45, 0xca, 0x60, 0x5e, 0x74, 0x8e, 0xce, 0xda, 0x75, 0x8d, 0x94, 0xa0, 0xe8, 0xb5,
	0xff, 0xf0, 0xea, 0xba, 0x4c, 0x7a, 0x47, 0xc7, 0x97, 0xed, 0x7a, 0xa1, 0xf5, 0xb6, 0x00, 0xb5,
	0x4e, 0xc2, 0xfa, 0xbb, 0xfa, 0xb8, 0x90, 0x7d, 0x28, 0x5f, 0x84, 0x7f, 0xb2, 0x98, 0x85, 0x3d,
	0x46, 0xbe, 0x98, 0xeb, 0xd7, 0x9e, 0x4f, 0x39, 0x1a, 0xe9, 0x40, 0xe9, 0x8c, 0x09, 0xf5, 0x09,
	0xfa, 0x3a, 0xd7, 0x77, 0xee, 0x78, 0xec, 0x6f, 0x5e, 0x5c, 0x57, 0x87, 0xe2, 0x68, 0xe4, 0x47,
	0xa4, 0x53, 0xd6, 0xc9, 0xf4, 0x8d, 0xf7, 0xc3, 0xb6, 0xd2, 0xc4, 0xac, 0xcb, 0x1c, 0x8d, 0x1c,
	0x02, 0xa4, 0x8e, 0x9a, 0xdf, 0x9a, 0xf9, 0x44, 0xcc, 0x1b, 0xcf, 0xd1, 0xc8, 0x2e, 0xc0, 0x49,
	0xcc, 0xa8, 0x60, 0x68, 0xfc, 0xcc, 0xc1, 0xc8, 0xd8, 0xce, 0xd3, 0x39, 0x1a, 0xf9, 0x01, 0x2a,
	0x6a, 0x83, 0xba, 0x7f, 0x79, 0xab, 0x3e, 0xb3, 0xe5, 0x6e, 0x01, 0x3f, 0xde, 0xfb, 0xef, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x07, 0xfe, 0xdd, 0xd5, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModelzooServiceClient is the client API for ModelzooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelzooServiceClient interface {
	// Inference
	Inference(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
	// Website utils
	GetImage(ctx context.Context, in *ImageDownloadRequest, opts ...grpc.CallOption) (*ImageDownloadResponse, error)
	// Rate limiting
	GetToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RateLimitToken, error)
	// Database
	ListModels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListModelsResponse, error)
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	CreateModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*Empty, error)
}

type modelzooServiceClient struct {
	cc *grpc.ClientConn
}

func NewModelzooServiceClient(cc *grpc.ClientConn) ModelzooServiceClient {
	return &modelzooServiceClient{cc}
}

func (c *modelzooServiceClient) Inference(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := c.cc.Invoke(ctx, "/modelzoo.ModelzooService/Inference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelzooServiceClient) GetImage(ctx context.Context, in *ImageDownloadRequest, opts ...grpc.CallOption) (*ImageDownloadResponse, error) {
	out := new(ImageDownloadResponse)
	err := c.cc.Invoke(ctx, "/modelzoo.ModelzooService/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelzooServiceClient) GetToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RateLimitToken, error) {
	out := new(RateLimitToken)
	err := c.cc.Invoke(ctx, "/modelzoo.ModelzooService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelzooServiceClient) ListModels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	out := new(ListModelsResponse)
	err := c.cc.Invoke(ctx, "/modelzoo.ModelzooService/ListModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelzooServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/modelzoo.ModelzooService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelzooServiceClient) CreateModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/modelzoo.ModelzooService/CreateModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelzooServiceServer is the server API for ModelzooService service.
type ModelzooServiceServer interface {
	// Inference
	Inference(context.Context, *Payload) (*Payload, error)
	// Website utils
	GetImage(context.Context, *ImageDownloadRequest) (*ImageDownloadResponse, error)
	// Rate limiting
	GetToken(context.Context, *Empty) (*RateLimitToken, error)
	// Database
	ListModels(context.Context, *Empty) (*ListModelsResponse, error)
	CreateUser(context.Context, *User) (*Empty, error)
	CreateModel(context.Context, *Model) (*Empty, error)
}

// UnimplementedModelzooServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModelzooServiceServer struct {
}

func (*UnimplementedModelzooServiceServer) Inference(ctx context.Context, req *Payload) (*Payload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inference not implemented")
}
func (*UnimplementedModelzooServiceServer) GetImage(ctx context.Context, req *ImageDownloadRequest) (*ImageDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (*UnimplementedModelzooServiceServer) GetToken(ctx context.Context, req *Empty) (*RateLimitToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedModelzooServiceServer) ListModels(ctx context.Context, req *Empty) (*ListModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (*UnimplementedModelzooServiceServer) CreateUser(ctx context.Context, req *User) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedModelzooServiceServer) CreateModel(ctx context.Context, req *Model) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModel not implemented")
}

func RegisterModelzooServiceServer(s *grpc.Server, srv ModelzooServiceServer) {
	s.RegisterService(&_ModelzooService_serviceDesc, srv)
}

func _ModelzooService_Inference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelzooServiceServer).Inference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelzoo.ModelzooService/Inference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelzooServiceServer).Inference(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelzooService_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelzooServiceServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelzoo.ModelzooService/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelzooServiceServer).GetImage(ctx, req.(*ImageDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelzooService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelzooServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelzoo.ModelzooService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelzooServiceServer).GetToken(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelzooService_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelzooServiceServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelzoo.ModelzooService/ListModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelzooServiceServer).ListModels(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelzooService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelzooServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelzoo.ModelzooService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelzooServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelzooService_CreateModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Model)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelzooServiceServer).CreateModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelzoo.ModelzooService/CreateModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelzooServiceServer).CreateModel(ctx, req.(*Model))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelzooService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modelzoo.ModelzooService",
	HandlerType: (*ModelzooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inference",
			Handler:    _ModelzooService_Inference_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _ModelzooService_GetImage_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _ModelzooService_GetToken_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _ModelzooService_ListModels_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _ModelzooService_CreateUser_Handler,
		},
		{
			MethodName: "CreateModel",
			Handler:    _ModelzooService_CreateModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/services.proto",
}
