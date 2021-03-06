// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package customerproto

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

type RequestHeader struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Employee struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	MiddleName           string   `protobuf:"bytes,2,opt,name=middleName,proto3" json:"middleName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Position             string   `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	ContactPhone         string   `protobuf:"bytes,5,opt,name=contactPhone,proto3" json:"contactPhone,omitempty"`
	ContactMobile        string   `protobuf:"bytes,6,opt,name=contactMobile,proto3" json:"contactMobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (m *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(m, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Employee) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *Employee) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Employee) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Employee) GetContactPhone() string {
	if m != nil {
		return m.ContactPhone
	}
	return ""
}

func (m *Employee) GetContactMobile() string {
	if m != nil {
		return m.ContactMobile
	}
	return ""
}

type RequestField struct {
	CustomerType         string      `protobuf:"bytes,1,opt,name=customerType,proto3" json:"customerType,omitempty"`
	Prefix               string      `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	FirstName            string      `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	MiddleName           string      `protobuf:"bytes,4,opt,name=middleName,proto3" json:"middleName,omitempty"`
	LastName             string      `protobuf:"bytes,5,opt,name=lastName,proto3" json:"lastName,omitempty"`
	OrganisationName     string      `protobuf:"bytes,6,opt,name=organisationName,proto3" json:"organisationName,omitempty"`
	ContactEmail         string      `protobuf:"bytes,7,opt,name=contactEmail,proto3" json:"contactEmail,omitempty"`
	ContactMobile        string      `protobuf:"bytes,8,opt,name=contactMobile,proto3" json:"contactMobile,omitempty"`
	ContactPhone         string      `protobuf:"bytes,9,opt,name=contactPhone,proto3" json:"contactPhone,omitempty"`
	AddressStreet        string      `protobuf:"bytes,10,opt,name=addressStreet,proto3" json:"addressStreet,omitempty"`
	AddressCity          string      `protobuf:"bytes,11,opt,name=addressCity,proto3" json:"addressCity,omitempty"`
	AddressPostCode      string      `protobuf:"bytes,12,opt,name=addressPostCode,proto3" json:"addressPostCode,omitempty"`
	AddressState         string      `protobuf:"bytes,13,opt,name=addressState,proto3" json:"addressState,omitempty"`
	AddressCountry       string      `protobuf:"bytes,14,opt,name=addressCountry,proto3" json:"addressCountry,omitempty"`
	BalanceAmount        int64       `protobuf:"varint,15,opt,name=balanceAmount,proto3" json:"balanceAmount,omitempty"`
	EmployeeDetails      []*Employee `protobuf:"bytes,16,rep,name=employeeDetails,proto3" json:"employeeDetails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RequestField) Reset()         { *m = RequestField{} }
func (m *RequestField) String() string { return proto.CompactTextString(m) }
func (*RequestField) ProtoMessage()    {}
func (*RequestField) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *RequestField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestField.Unmarshal(m, b)
}
func (m *RequestField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestField.Marshal(b, m, deterministic)
}
func (m *RequestField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestField.Merge(m, src)
}
func (m *RequestField) XXX_Size() int {
	return xxx_messageInfo_RequestField.Size(m)
}
func (m *RequestField) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestField.DiscardUnknown(m)
}

var xxx_messageInfo_RequestField proto.InternalMessageInfo

func (m *RequestField) GetCustomerType() string {
	if m != nil {
		return m.CustomerType
	}
	return ""
}

func (m *RequestField) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RequestField) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RequestField) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *RequestField) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RequestField) GetOrganisationName() string {
	if m != nil {
		return m.OrganisationName
	}
	return ""
}

func (m *RequestField) GetContactEmail() string {
	if m != nil {
		return m.ContactEmail
	}
	return ""
}

func (m *RequestField) GetContactMobile() string {
	if m != nil {
		return m.ContactMobile
	}
	return ""
}

func (m *RequestField) GetContactPhone() string {
	if m != nil {
		return m.ContactPhone
	}
	return ""
}

func (m *RequestField) GetAddressStreet() string {
	if m != nil {
		return m.AddressStreet
	}
	return ""
}

func (m *RequestField) GetAddressCity() string {
	if m != nil {
		return m.AddressCity
	}
	return ""
}

func (m *RequestField) GetAddressPostCode() string {
	if m != nil {
		return m.AddressPostCode
	}
	return ""
}

func (m *RequestField) GetAddressState() string {
	if m != nil {
		return m.AddressState
	}
	return ""
}

func (m *RequestField) GetAddressCountry() string {
	if m != nil {
		return m.AddressCountry
	}
	return ""
}

func (m *RequestField) GetBalanceAmount() int64 {
	if m != nil {
		return m.BalanceAmount
	}
	return 0
}

func (m *RequestField) GetEmployeeDetails() []*Employee {
	if m != nil {
		return m.EmployeeDetails
	}
	return nil
}

type RawRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RawRequest) Reset()         { *m = RawRequest{} }
func (m *RawRequest) String() string { return proto.CompactTextString(m) }
func (*RawRequest) ProtoMessage()    {}
func (*RawRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RawRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawRequest.Unmarshal(m, b)
}
func (m *RawRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawRequest.Marshal(b, m, deterministic)
}
func (m *RawRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawRequest.Merge(m, src)
}
func (m *RawRequest) XXX_Size() int {
	return xxx_messageInfo_RawRequest.Size(m)
}
func (m *RawRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RawRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RawRequest proto.InternalMessageInfo

func (m *RawRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Field                *RequestField  `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
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

func (m *CreateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type UpdateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Action               string         `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Field                *RequestField  `protobuf:"bytes,4,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UpdateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *UpdateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type DataResponse struct {
	StatusCode           int32              `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *DataResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}
func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}
func (m *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(m, src)
}
func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}
func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DataResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DataResponse) GetData() *DataResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DataResponse_Data struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerType         string      `protobuf:"bytes,2,opt,name=customerType,proto3" json:"customerType,omitempty"`
	Prefix               string      `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	FirstName            string      `protobuf:"bytes,4,opt,name=firstName,proto3" json:"firstName,omitempty"`
	MiddleName           string      `protobuf:"bytes,5,opt,name=middleName,proto3" json:"middleName,omitempty"`
	LastName             string      `protobuf:"bytes,6,opt,name=lastName,proto3" json:"lastName,omitempty"`
	OrganisationName     string      `protobuf:"bytes,7,opt,name=organisationName,proto3" json:"organisationName,omitempty"`
	ContactEmail         string      `protobuf:"bytes,8,opt,name=contactEmail,proto3" json:"contactEmail,omitempty"`
	ContactMobile        string      `protobuf:"bytes,9,opt,name=contactMobile,proto3" json:"contactMobile,omitempty"`
	ContactPhone         string      `protobuf:"bytes,10,opt,name=contactPhone,proto3" json:"contactPhone,omitempty"`
	AddressStreet        string      `protobuf:"bytes,11,opt,name=addressStreet,proto3" json:"addressStreet,omitempty"`
	AddressCity          string      `protobuf:"bytes,12,opt,name=addressCity,proto3" json:"addressCity,omitempty"`
	AddressPostCode      string      `protobuf:"bytes,13,opt,name=addressPostCode,proto3" json:"addressPostCode,omitempty"`
	AddressState         string      `protobuf:"bytes,14,opt,name=addressState,proto3" json:"addressState,omitempty"`
	AddressCountry       string      `protobuf:"bytes,15,opt,name=addressCountry,proto3" json:"addressCountry,omitempty"`
	BalanceAmount        string      `protobuf:"bytes,16,opt,name=balanceAmount,proto3" json:"balanceAmount,omitempty"`
	EmployeeDetails      []*Employee `protobuf:"bytes,17,rep,name=employeeDetails,proto3" json:"employeeDetails,omitempty"`
	CreatedBy            string      `protobuf:"bytes,18,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedDate          string      `protobuf:"bytes,19,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	LastModifiedBy       string      `protobuf:"bytes,20,opt,name=lastModifiedBy,proto3" json:"lastModifiedBy,omitempty"`
	LastModifiedDate     string      `protobuf:"bytes,21,opt,name=lastModifiedDate,proto3" json:"lastModifiedDate,omitempty"`
	RecordState          string      `protobuf:"bytes,22,opt,name=recordState,proto3" json:"recordState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DataResponse_Data) Reset()         { *m = DataResponse_Data{} }
func (m *DataResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DataResponse_Data) ProtoMessage()    {}
func (*DataResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7, 0}
}

func (m *DataResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse_Data.Unmarshal(m, b)
}
func (m *DataResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse_Data.Marshal(b, m, deterministic)
}
func (m *DataResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse_Data.Merge(m, src)
}
func (m *DataResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DataResponse_Data.Size(m)
}
func (m *DataResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse_Data proto.InternalMessageInfo

func (m *DataResponse_Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataResponse_Data) GetCustomerType() string {
	if m != nil {
		return m.CustomerType
	}
	return ""
}

func (m *DataResponse_Data) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *DataResponse_Data) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *DataResponse_Data) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *DataResponse_Data) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *DataResponse_Data) GetOrganisationName() string {
	if m != nil {
		return m.OrganisationName
	}
	return ""
}

func (m *DataResponse_Data) GetContactEmail() string {
	if m != nil {
		return m.ContactEmail
	}
	return ""
}

func (m *DataResponse_Data) GetContactMobile() string {
	if m != nil {
		return m.ContactMobile
	}
	return ""
}

func (m *DataResponse_Data) GetContactPhone() string {
	if m != nil {
		return m.ContactPhone
	}
	return ""
}

func (m *DataResponse_Data) GetAddressStreet() string {
	if m != nil {
		return m.AddressStreet
	}
	return ""
}

func (m *DataResponse_Data) GetAddressCity() string {
	if m != nil {
		return m.AddressCity
	}
	return ""
}

func (m *DataResponse_Data) GetAddressPostCode() string {
	if m != nil {
		return m.AddressPostCode
	}
	return ""
}

func (m *DataResponse_Data) GetAddressState() string {
	if m != nil {
		return m.AddressState
	}
	return ""
}

func (m *DataResponse_Data) GetAddressCountry() string {
	if m != nil {
		return m.AddressCountry
	}
	return ""
}

func (m *DataResponse_Data) GetBalanceAmount() string {
	if m != nil {
		return m.BalanceAmount
	}
	return ""
}

func (m *DataResponse_Data) GetEmployeeDetails() []*Employee {
	if m != nil {
		return m.EmployeeDetails
	}
	return nil
}

func (m *DataResponse_Data) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *DataResponse_Data) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *DataResponse_Data) GetLastModifiedBy() string {
	if m != nil {
		return m.LastModifiedBy
	}
	return ""
}

func (m *DataResponse_Data) GetLastModifiedDate() string {
	if m != nil {
		return m.LastModifiedDate
	}
	return ""
}

func (m *DataResponse_Data) GetRecordState() string {
	if m != nil {
		return m.RecordState
	}
	return ""
}

type RawResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *RawResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RawResponse) Reset()         { *m = RawResponse{} }
func (m *RawResponse) String() string { return proto.CompactTextString(m) }
func (*RawResponse) ProtoMessage()    {}
func (*RawResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *RawResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawResponse.Unmarshal(m, b)
}
func (m *RawResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawResponse.Marshal(b, m, deterministic)
}
func (m *RawResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawResponse.Merge(m, src)
}
func (m *RawResponse) XXX_Size() int {
	return xxx_messageInfo_RawResponse.Size(m)
}
func (m *RawResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RawResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RawResponse proto.InternalMessageInfo

func (m *RawResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *RawResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RawResponse) GetData() *RawResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type RawResponse_Data struct {
	CustomerType         []string `protobuf:"bytes,1,rep,name=customerType,proto3" json:"customerType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawResponse_Data) Reset()         { *m = RawResponse_Data{} }
func (m *RawResponse_Data) String() string { return proto.CompactTextString(m) }
func (*RawResponse_Data) ProtoMessage()    {}
func (*RawResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8, 0}
}

func (m *RawResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawResponse_Data.Unmarshal(m, b)
}
func (m *RawResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawResponse_Data.Marshal(b, m, deterministic)
}
func (m *RawResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawResponse_Data.Merge(m, src)
}
func (m *RawResponse_Data) XXX_Size() int {
	return xxx_messageInfo_RawResponse_Data.Size(m)
}
func (m *RawResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_RawResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_RawResponse_Data proto.InternalMessageInfo

func (m *RawResponse_Data) GetCustomerType() []string {
	if m != nil {
		return m.CustomerType
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "customerproto.RequestHeader")
	proto.RegisterType((*Employee)(nil), "customerproto.Employee")
	proto.RegisterType((*RequestField)(nil), "customerproto.RequestField")
	proto.RegisterType((*RawRequest)(nil), "customerproto.RawRequest")
	proto.RegisterType((*GetRequest)(nil), "customerproto.GetRequest")
	proto.RegisterType((*CreateRequest)(nil), "customerproto.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "customerproto.UpdateRequest")
	proto.RegisterType((*DataResponse)(nil), "customerproto.DataResponse")
	proto.RegisterType((*DataResponse_Data)(nil), "customerproto.DataResponse.Data")
	proto.RegisterType((*RawResponse)(nil), "customerproto.RawResponse")
	proto.RegisterType((*RawResponse_Data)(nil), "customerproto.RawResponse.Data")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc7, 0x05, 0x06, 0x03, 0xc7, 0x18, 0xd2, 0x69, 0x9a, 0xba, 0x24, 0x6a, 0x91, 0xd5, 0x56,
	0x28, 0x17, 0x91, 0x4a, 0x72, 0xd9, 0x5e, 0x24, 0x24, 0x4d, 0x6f, 0x52, 0x45, 0x6e, 0xfb, 0x00,
	0x13, 0x7b, 0x48, 0x46, 0x35, 0x1e, 0xd6, 0x33, 0x6c, 0xc2, 0x63, 0xec, 0x1b, 0xac, 0xb4, 0x7b,
	0xb1, 0xcf, 0xb2, 0xef, 0xb1, 0xef, 0xb1, 0x9a, 0x0f, 0x82, 0x6d, 0x88, 0xd7, 0x2b, 0x2e, 0xcf,
	0x7f, 0x8e, 0xcf, 0x97, 0xce, 0xf9, 0x01, 0xb8, 0x9c, 0xa4, 0xaf, 0x69, 0x48, 0x4e, 0xe6, 0x29,
	0x13, 0x0c, 0xb9, 0xe1, 0x82, 0x0b, 0x36, 0x23, 0xa9, 0x32, 0xfd, 0x5f, 0xc0, 0x0d, 0xc8, 0xab,
	0x05, 0xe1, 0xe2, 0x2f, 0x82, 0x23, 0x92, 0xa2, 0x7d, 0x68, 0x0a, 0xf6, 0x3f, 0x49, 0xbc, 0xda,
	0xb0, 0x36, 0xea, 0x04, 0xda, 0xf0, 0x3f, 0xd6, 0xa0, 0x7d, 0x35, 0x9b, 0xc7, 0x6c, 0x49, 0x08,
	0x3a, 0x82, 0xce, 0x94, 0xa6, 0x5c, 0xfc, 0x8d, 0x67, 0xc4, 0xb8, 0xad, 0x05, 0xf4, 0x23, 0xc0,
	0x8c, 0x46, 0x51, 0x4c, 0xd4, 0x73, 0x5d, 0x3d, 0x67, 0x14, 0x34, 0x80, 0x76, 0x8c, 0xcd, 0xc7,
	0x96, 0x7a, 0x7d, 0xb6, 0xe5, 0xdb, 0x9c, 0x71, 0x2a, 0x28, 0x4b, 0xbc, 0x86, 0x7e, 0x5b, 0xd9,
	0xc8, 0x87, 0x6e, 0xc8, 0x12, 0x81, 0x43, 0x71, 0xfb, 0xc0, 0x12, 0xe2, 0x35, 0xd5, 0x7b, 0x4e,
	0x43, 0x3f, 0x83, 0x6b, 0xec, 0x1b, 0x76, 0x47, 0x63, 0xe2, 0xd9, 0xca, 0x29, 0x2f, 0xfa, 0x9f,
	0x1a, 0xd0, 0x35, 0x4d, 0xff, 0x49, 0x49, 0x1c, 0xa9, 0xd0, 0x66, 0x2a, 0xff, 0x2e, 0xe7, 0xab,
	0x9e, 0x72, 0x1a, 0x3a, 0x00, 0x7b, 0x9e, 0x92, 0x29, 0x7d, 0x32, 0x2d, 0x19, 0x2b, 0x3f, 0x0c,
	0xab, 0x7c, 0x18, 0x8d, 0xd2, 0x61, 0x34, 0x0b, 0xc3, 0x38, 0x86, 0x3d, 0x96, 0xde, 0xe3, 0x84,
	0x72, 0x2c, 0x07, 0xa0, 0x7c, 0x74, 0x3f, 0x1b, 0x7a, 0x66, 0x38, 0x57, 0x33, 0x4c, 0x63, 0xaf,
	0x95, 0x1b, 0x8e, 0xd2, 0x36, 0x87, 0xd3, 0xde, 0x32, 0x9c, 0x8d, 0x31, 0x77, 0xb6, 0x8f, 0x19,
	0x47, 0x51, 0x4a, 0x38, 0xff, 0x47, 0xa4, 0x84, 0x08, 0x0f, 0x74, 0xa4, 0x9c, 0x88, 0x86, 0xe0,
	0x18, 0x61, 0x42, 0xc5, 0xd2, 0x73, 0x94, 0x4f, 0x56, 0x42, 0x23, 0xe8, 0x1b, 0xf3, 0x96, 0x71,
	0x31, 0x61, 0x11, 0xf1, 0xba, 0xca, 0xab, 0x28, 0xcb, 0xaa, 0x9e, 0x83, 0x63, 0x41, 0x3c, 0x57,
	0x57, 0x95, 0xd5, 0xd0, 0xaf, 0xd0, 0x5b, 0x05, 0x67, 0x8b, 0x44, 0xa4, 0x4b, 0xaf, 0xa7, 0xbc,
	0x0a, 0xaa, 0xac, 0xfe, 0x0e, 0xc7, 0x38, 0x09, 0xc9, 0xf9, 0x4c, 0x4a, 0x5e, 0x7f, 0x58, 0x1b,
	0x59, 0x41, 0x5e, 0x44, 0xe7, 0xd0, 0x27, 0x66, 0xe1, 0x2f, 0x89, 0xc0, 0x34, 0xe6, 0xde, 0xde,
	0xd0, 0x1a, 0x39, 0xe3, 0xef, 0x4f, 0x72, 0x17, 0x74, 0xb2, 0x3a, 0x8b, 0xa0, 0xe8, 0xef, 0x5f,
	0x00, 0x04, 0xf8, 0xd1, 0x6c, 0x1a, 0x3a, 0x03, 0xfb, 0x41, 0x9d, 0x98, 0x5a, 0x2f, 0x67, 0x7c,
	0x54, 0x88, 0x93, 0x3b, 0xc3, 0xc0, 0xf8, 0xfa, 0x01, 0xc0, 0x35, 0x11, 0x3b, 0xc5, 0x40, 0x3d,
	0xa8, 0xd3, 0x48, 0xad, 0x6d, 0x33, 0xa8, 0xd3, 0xc8, 0x7f, 0x02, 0x77, 0x92, 0x12, 0x2c, 0xc8,
	0x6e, 0x61, 0x7f, 0x83, 0xe6, 0x54, 0x9e, 0x8f, 0x8a, 0xec, 0x8c, 0x0f, 0xb7, 0x7f, 0xa4, 0x2e,
	0x2c, 0xd0, 0x9e, 0xfe, 0xdb, 0x1a, 0xb8, 0xff, 0xcd, 0xa3, 0x9d, 0x53, 0x17, 0x3a, 0x92, 0xc7,
	0x89, 0x43, 0x45, 0x0d, 0x7d, 0x81, 0xc6, 0x5a, 0x97, 0xd8, 0xa8, 0x5c, 0xe2, 0xfb, 0x16, 0x74,
	0x2f, 0xb1, 0xc0, 0x01, 0xe1, 0x73, 0x96, 0x70, 0x75, 0xc2, 0x5c, 0x60, 0xb1, 0xe0, 0x6a, 0x3f,
	0x6b, 0x2a, 0x67, 0x46, 0x41, 0x1e, 0xb4, 0x66, 0x84, 0x73, 0x7c, 0xbf, 0x82, 0xdd, 0xca, 0x44,
	0x67, 0xd0, 0x88, 0xb0, 0xc0, 0xaa, 0x26, 0x67, 0x3c, 0x2c, 0x24, 0xcf, 0x26, 0xd1, 0x86, 0xf2,
	0x1e, 0xbc, 0xb3, 0xa1, 0x21, 0x4d, 0xd3, 0xa4, 0x66, 0x91, 0x6c, 0xb2, 0x48, 0xa9, 0x7a, 0x29,
	0xa5, 0xac, 0x97, 0x29, 0xd5, 0x28, 0xa7, 0x54, 0xb3, 0x94, 0x52, 0x76, 0x05, 0x4a, 0xb5, 0x2a,
	0x52, 0xaa, 0x5d, 0x85, 0x52, 0x9d, 0x2a, 0x94, 0x82, 0x2a, 0x94, 0x72, 0x2a, 0x50, 0xaa, 0x5b,
	0x89, 0x52, 0x6e, 0x35, 0x4a, 0xf5, 0x2a, 0x51, 0xaa, 0x5f, 0x8d, 0x52, 0x7b, 0xba, 0xfa, 0x2f,
	0x52, 0xea, 0x9b, 0xaf, 0xa3, 0x94, 0x5c, 0x8d, 0x50, 0xd1, 0x20, 0xba, 0x58, 0x7a, 0x48, 0xaf,
	0xc6, 0xb3, 0x20, 0xc7, 0x63, 0x8c, 0x4b, 0xd9, 0xd1, 0xb7, 0x7a, 0x3c, 0x19, 0x49, 0x36, 0x24,
	0x97, 0xe1, 0x86, 0x45, 0x74, 0x4a, 0x55, 0x90, 0x7d, 0xdd, 0x50, 0x5e, 0x95, 0x8b, 0x92, 0x55,
	0x54, 0xb8, 0xef, 0xf4, 0xa2, 0x14, 0x75, 0x99, 0x35, 0x25, 0x21, 0x4b, 0x23, 0x3d, 0xc7, 0x03,
	0x9d, 0x35, 0x23, 0xf9, 0x1f, 0x6a, 0xe0, 0x28, 0xb8, 0xee, 0x7c, 0xa5, 0xa7, 0xb9, 0x2b, 0xfd,
	0xa9, 0x88, 0x88, 0x75, 0x8e, 0xec, 0x91, 0x1e, 0x9b, 0x1b, 0xdd, 0xfc, 0xe7, 0x60, 0x15, 0x6f,
	0x72, 0xfc, 0xa6, 0x0e, 0xed, 0x89, 0x11, 0xd0, 0xef, 0x60, 0x05, 0xf8, 0x11, 0xfd, 0xb0, 0x2d,
	0x8d, 0x82, 0xd1, 0x60, 0xf0, 0x72, 0x05, 0x68, 0x02, 0xb6, 0x26, 0x37, 0x2a, 0x72, 0x32, 0x07,
	0xf4, 0xc1, 0x61, 0x09, 0x6b, 0xd0, 0x1f, 0x60, 0x5d, 0x13, 0xb1, 0x51, 0xc2, 0xfa, 0x67, 0xa6,
	0xfc, 0xf3, 0x09, 0xd8, 0x1a, 0xe1, 0x1b, 0x35, 0xe4, 0xc8, 0x5e, 0x1a, 0xe4, 0xce, 0x56, 0xda,
	0xe9, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xfd, 0xd9, 0x5c, 0x9d, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	Raw(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Raw(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error) {
	out := new(RawResponse)
	err := c.cc.Invoke(ctx, "/customerproto.Customer/Raw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/customerproto.Customer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/customerproto.Customer/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/customerproto.Customer/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	Raw(context.Context, *RawRequest) (*RawResponse, error)
	Create(context.Context, *CreateRequest) (*DataResponse, error)
	Get(context.Context, *GetRequest) (*DataResponse, error)
	Update(context.Context, *UpdateRequest) (*DataResponse, error)
}

// UnimplementedCustomerServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (*UnimplementedCustomerServer) Raw(ctx context.Context, req *RawRequest) (*RawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Raw not implemented")
}
func (*UnimplementedCustomerServer) Create(ctx context.Context, req *CreateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCustomerServer) Get(ctx context.Context, req *GetRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCustomerServer) Update(ctx context.Context, req *UpdateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_Raw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Raw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerproto.Customer/Raw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Raw(ctx, req.(*RawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerproto.Customer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerproto.Customer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerproto.Customer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customerproto.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Raw",
			Handler:    _Customer_Raw_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Customer_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Customer_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Customer_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
