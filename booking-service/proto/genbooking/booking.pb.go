// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: booking.proto

package genbooking

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	HotelId      int32                  `protobuf:"varint,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	Roomtype     string                 `protobuf:"bytes,3,opt,name=roomtype,proto3" json:"roomtype,omitempty"`
	CheckInDate  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	CheckOutDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
	TotalAmount  float32                `protobuf:"fixed32,6,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BookingRequest) GetHotelId() int32 {
	if x != nil {
		return x.HotelId
	}
	return 0
}

func (x *BookingRequest) GetRoomtype() string {
	if x != nil {
		return x.Roomtype
	}
	return ""
}

func (x *BookingRequest) GetCheckInDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CheckInDate
	}
	return nil
}

func (x *BookingRequest) GetCheckOutDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CheckOutDate
	}
	return nil
}

func (x *BookingRequest) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId    int32   `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	UserId       int32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	HotelId      int32   `protobuf:"varint,3,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	RoomType     string  `protobuf:"bytes,4,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	CheckInDate  string  `protobuf:"bytes,5,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	CheckOutDate string  `protobuf:"bytes,6,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
	TotalAmount  float32 `protobuf:"fixed32,7,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Status       string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"` //status yozamannnnn
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *BookingResponse) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *BookingResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BookingResponse) GetHotelId() int32 {
	if x != nil {
		return x.HotelId
	}
	return 0
}

func (x *BookingResponse) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *BookingResponse) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *BookingResponse) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *BookingResponse) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *BookingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BookingIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId int32 `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *BookingIdReq) Reset() {
	*x = BookingIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingIdReq) ProtoMessage() {}

func (x *BookingIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingIdReq.ProtoReflect.Descriptor instead.
func (*BookingIdReq) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *BookingIdReq) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type BookingIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId    int32   `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	UserId       int32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	HotelId      int32   `protobuf:"varint,3,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	RoomType     string  `protobuf:"bytes,4,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	CheckInDate  string  `protobuf:"bytes,5,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	CheckOutDate string  `protobuf:"bytes,6,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
	TotalAmount  float32 `protobuf:"fixed32,7,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Status       string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BookingIdResp) Reset() {
	*x = BookingIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingIdResp) ProtoMessage() {}

func (x *BookingIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingIdResp.ProtoReflect.Descriptor instead.
func (*BookingIdResp) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *BookingIdResp) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *BookingIdResp) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BookingIdResp) GetHotelId() int32 {
	if x != nil {
		return x.HotelId
	}
	return 0
}

func (x *BookingIdResp) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *BookingIdResp) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *BookingIdResp) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *BookingIdResp) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *BookingIdResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateBookIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId    int32                  `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	CheckInDate  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	CheckOutDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
	RoomType     string                 `protobuf:"bytes,4,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
}

func (x *UpdateBookIdReq) Reset() {
	*x = UpdateBookIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookIdReq) ProtoMessage() {}

func (x *UpdateBookIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookIdReq.ProtoReflect.Descriptor instead.
func (*UpdateBookIdReq) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBookIdReq) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *UpdateBookIdReq) GetCheckInDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CheckInDate
	}
	return nil
}

func (x *UpdateBookIdReq) GetCheckOutDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CheckOutDate
	}
	return nil
}

func (x *UpdateBookIdReq) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

type UpdateBookIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId    int32                  `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	UserId       int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	HotelId      int32                  `protobuf:"varint,3,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	RoomType     string                 `protobuf:"bytes,4,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	CheckInDate  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	CheckOutDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
	TotalAmount  float32                `protobuf:"fixed32,7,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Status       string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"` //status yozamannnnn
}

func (x *UpdateBookIdResp) Reset() {
	*x = UpdateBookIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookIdResp) ProtoMessage() {}

func (x *UpdateBookIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookIdResp.ProtoReflect.Descriptor instead.
func (*UpdateBookIdResp) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBookIdResp) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *UpdateBookIdResp) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateBookIdResp) GetHotelId() int32 {
	if x != nil {
		return x.HotelId
	}
	return 0
}

func (x *UpdateBookIdResp) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *UpdateBookIdResp) GetCheckInDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CheckInDate
	}
	return nil
}

func (x *UpdateBookIdResp) GetCheckOutDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CheckOutDate
	}
	return nil
}

func (x *UpdateBookIdResp) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *UpdateBookIdResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteBookIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId int32 `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *DeleteBookIDReq) Reset() {
	*x = DeleteBookIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookIDReq) ProtoMessage() {}

func (x *DeleteBookIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookIDReq.ProtoReflect.Descriptor instead.
func (*DeleteBookIDReq) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookIDReq) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type DeleteBookIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteBookIdResp) Reset() {
	*x = DeleteBookIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookIdResp) ProtoMessage() {}

func (x *DeleteBookIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookIdResp.ProtoReflect.Descriptor instead.
func (*DeleteBookIdResp) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookIdResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x85, 0x02, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x86, 0x02, 0x0a, 0x0f, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x2d, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x22, 0x84, 0x02, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbf, 0x02, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x40,
	0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xdd, 0x01, 0x0a, 0x0e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x0f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x34, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x10, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_booking_proto_goTypes = []any{
	(*BookingRequest)(nil),        // 0: BookingRequest
	(*BookingResponse)(nil),       // 1: BookingResponse
	(*BookingIdReq)(nil),          // 2: BookingIdReq
	(*BookingIdResp)(nil),         // 3: BookingIdResp
	(*UpdateBookIdReq)(nil),       // 4: UpdateBookIdReq
	(*UpdateBookIdResp)(nil),      // 5: UpdateBookIdResp
	(*DeleteBookIDReq)(nil),       // 6: DeleteBookIDReq
	(*DeleteBookIdResp)(nil),      // 7: DeleteBookIdResp
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_booking_proto_depIdxs = []int32{
	8,  // 0: BookingRequest.check_in_date:type_name -> google.protobuf.Timestamp
	8,  // 1: BookingRequest.check_out_date:type_name -> google.protobuf.Timestamp
	8,  // 2: UpdateBookIdReq.check_in_date:type_name -> google.protobuf.Timestamp
	8,  // 3: UpdateBookIdReq.check_out_date:type_name -> google.protobuf.Timestamp
	8,  // 4: UpdateBookIdResp.check_in_date:type_name -> google.protobuf.Timestamp
	8,  // 5: UpdateBookIdResp.check_out_date:type_name -> google.protobuf.Timestamp
	0,  // 6: BookingService.CreateBooking:input_type -> BookingRequest
	2,  // 7: BookingService.GetBooking:input_type -> BookingIdReq
	4,  // 8: BookingService.UpdateBooking:input_type -> UpdateBookIdReq
	6,  // 9: BookingService.DeleteBooking:input_type -> DeleteBookIDReq
	1,  // 10: BookingService.CreateBooking:output_type -> BookingResponse
	3,  // 11: BookingService.GetBooking:output_type -> BookingIdResp
	5,  // 12: BookingService.UpdateBooking:output_type -> UpdateBookIdResp
	7,  // 13: BookingService.DeleteBooking:output_type -> DeleteBookIdResp
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BookingRequest); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BookingResponse); i {
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
		file_booking_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BookingIdReq); i {
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
		file_booking_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BookingIdResp); i {
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
		file_booking_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBookIdReq); i {
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
		file_booking_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBookIdResp); i {
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
		file_booking_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteBookIDReq); i {
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
		file_booking_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteBookIdResp); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
