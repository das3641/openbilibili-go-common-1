// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v0/Smalltv.proto

package v0

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SmalltvStartReq struct {
	// 用户id
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// 房间号
	Roomid int64 `protobuf:"varint,2,opt,name=roomid,proto3" json:"roomid"`
	// 道具id
	GiftId int64 `protobuf:"varint,3,opt,name=gift_id,json=giftId,proto3" json:"gift_id"`
	// 道具个数
	Num int64 `protobuf:"varint,4,opt,name=num,proto3" json:"num"`
	// 业务id
	Tid int64 `protobuf:"varint,5,opt,name=tid,proto3" json:"tid"`
	// 公告样式id
	StyleId int64 `protobuf:"varint,6,opt,name=style_id,json=styleId,proto3" json:"style_id"`
}

func (m *SmalltvStartReq) Reset()         { *m = SmalltvStartReq{} }
func (m *SmalltvStartReq) String() string { return proto.CompactTextString(m) }
func (*SmalltvStartReq) ProtoMessage()    {}
func (*SmalltvStartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_Smalltv_627605adbee0258c, []int{0}
}
func (m *SmalltvStartReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmalltvStartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmalltvStartReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SmalltvStartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmalltvStartReq.Merge(dst, src)
}
func (m *SmalltvStartReq) XXX_Size() int {
	return m.Size()
}
func (m *SmalltvStartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SmalltvStartReq.DiscardUnknown(m)
}

var xxx_messageInfo_SmalltvStartReq proto.InternalMessageInfo

func (m *SmalltvStartReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SmalltvStartReq) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *SmalltvStartReq) GetGiftId() int64 {
	if m != nil {
		return m.GiftId
	}
	return 0
}

func (m *SmalltvStartReq) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *SmalltvStartReq) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *SmalltvStartReq) GetStyleId() int64 {
	if m != nil {
		return m.StyleId
	}
	return 0
}

type SmalltvStartResp struct {
	//
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	//
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data []*SmalltvStartResp_MSG `protobuf:"bytes,3,rep,name=data" json:"data"`
}

func (m *SmalltvStartResp) Reset()         { *m = SmalltvStartResp{} }
func (m *SmalltvStartResp) String() string { return proto.CompactTextString(m) }
func (*SmalltvStartResp) ProtoMessage()    {}
func (*SmalltvStartResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_Smalltv_627605adbee0258c, []int{1}
}
func (m *SmalltvStartResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmalltvStartResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmalltvStartResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SmalltvStartResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmalltvStartResp.Merge(dst, src)
}
func (m *SmalltvStartResp) XXX_Size() int {
	return m.Size()
}
func (m *SmalltvStartResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SmalltvStartResp.DiscardUnknown(m)
}

var xxx_messageInfo_SmalltvStartResp proto.InternalMessageInfo

func (m *SmalltvStartResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SmalltvStartResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SmalltvStartResp) GetData() []*SmalltvStartResp_MSG {
	if m != nil {
		return m.Data
	}
	return nil
}

type SmalltvStartResp_MSG struct {
	//
	Cmd string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd"`
	//
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	MsgText string `protobuf:"bytes,3,opt,name=msg_text,json=msgText,proto3" json:"msg_text"`
	//
	MsgCommon string `protobuf:"bytes,4,opt,name=msg_common,json=msgCommon,proto3" json:"msg_common"`
	//
	MsgSelf string `protobuf:"bytes,5,opt,name=msg_self,json=msgSelf,proto3" json:"msg_self"`
	//
	Rep int64 `protobuf:"varint,6,opt,name=rep,proto3" json:"rep"`
	//
	StyleType int64 `protobuf:"varint,7,opt,name=styleType,proto3" json:"styleType"`
	//
	Url string `protobuf:"bytes,8,opt,name=url,proto3" json:"url"`
	//
	Roomid int64 `protobuf:"varint,9,opt,name=roomid,proto3" json:"roomid"`
	//
	RealRoomid int64 `protobuf:"varint,10,opt,name=real_roomid,json=realRoomid,proto3" json:"real_roomid"`
	//
	Rnd int64 `protobuf:"varint,11,opt,name=rnd,proto3" json:"rnd"`
	//
	BroadcastType int64 `protobuf:"varint,12,opt,name=broadcast_type,json=broadcastType,proto3" json:"broadcast_type"`
}

func (m *SmalltvStartResp_MSG) Reset()         { *m = SmalltvStartResp_MSG{} }
func (m *SmalltvStartResp_MSG) String() string { return proto.CompactTextString(m) }
func (*SmalltvStartResp_MSG) ProtoMessage()    {}
func (*SmalltvStartResp_MSG) Descriptor() ([]byte, []int) {
	return fileDescriptor_Smalltv_627605adbee0258c, []int{1, 0}
}
func (m *SmalltvStartResp_MSG) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmalltvStartResp_MSG) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmalltvStartResp_MSG.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SmalltvStartResp_MSG) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmalltvStartResp_MSG.Merge(dst, src)
}
func (m *SmalltvStartResp_MSG) XXX_Size() int {
	return m.Size()
}
func (m *SmalltvStartResp_MSG) XXX_DiscardUnknown() {
	xxx_messageInfo_SmalltvStartResp_MSG.DiscardUnknown(m)
}

var xxx_messageInfo_SmalltvStartResp_MSG proto.InternalMessageInfo

func (m *SmalltvStartResp_MSG) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *SmalltvStartResp_MSG) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SmalltvStartResp_MSG) GetMsgText() string {
	if m != nil {
		return m.MsgText
	}
	return ""
}

func (m *SmalltvStartResp_MSG) GetMsgCommon() string {
	if m != nil {
		return m.MsgCommon
	}
	return ""
}

func (m *SmalltvStartResp_MSG) GetMsgSelf() string {
	if m != nil {
		return m.MsgSelf
	}
	return ""
}

func (m *SmalltvStartResp_MSG) GetRep() int64 {
	if m != nil {
		return m.Rep
	}
	return 0
}

func (m *SmalltvStartResp_MSG) GetStyleType() int64 {
	if m != nil {
		return m.StyleType
	}
	return 0
}

func (m *SmalltvStartResp_MSG) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SmalltvStartResp_MSG) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *SmalltvStartResp_MSG) GetRealRoomid() int64 {
	if m != nil {
		return m.RealRoomid
	}
	return 0
}

func (m *SmalltvStartResp_MSG) GetRnd() int64 {
	if m != nil {
		return m.Rnd
	}
	return 0
}

func (m *SmalltvStartResp_MSG) GetBroadcastType() int64 {
	if m != nil {
		return m.BroadcastType
	}
	return 0
}

func init() {
	proto.RegisterType((*SmalltvStartReq)(nil), "gift.v0.SmalltvStartReq")
	proto.RegisterType((*SmalltvStartResp)(nil), "gift.v0.SmalltvStartResp")
	proto.RegisterType((*SmalltvStartResp_MSG)(nil), "gift.v0.SmalltvStartResp.MSG")
}
func (m *SmalltvStartReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmalltvStartReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Uid))
	}
	if m.Roomid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Roomid))
	}
	if m.GiftId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.GiftId))
	}
	if m.Num != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Num))
	}
	if m.Tid != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Tid))
	}
	if m.StyleId != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.StyleId))
	}
	return i, nil
}

func (m *SmalltvStartResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmalltvStartResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSmalltv(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SmalltvStartResp_MSG) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmalltvStartResp_MSG) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cmd) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(len(m.Cmd)))
		i += copy(dAtA[i:], m.Cmd)
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.MsgText) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(len(m.MsgText)))
		i += copy(dAtA[i:], m.MsgText)
	}
	if len(m.MsgCommon) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(len(m.MsgCommon)))
		i += copy(dAtA[i:], m.MsgCommon)
	}
	if len(m.MsgSelf) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(len(m.MsgSelf)))
		i += copy(dAtA[i:], m.MsgSelf)
	}
	if m.Rep != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Rep))
	}
	if m.StyleType != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.StyleType))
	}
	if len(m.Url) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if m.Roomid != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Roomid))
	}
	if m.RealRoomid != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.RealRoomid))
	}
	if m.Rnd != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.Rnd))
	}
	if m.BroadcastType != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintSmalltv(dAtA, i, uint64(m.BroadcastType))
	}
	return i, nil
}

func encodeVarintSmalltv(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SmalltvStartReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovSmalltv(uint64(m.Uid))
	}
	if m.Roomid != 0 {
		n += 1 + sovSmalltv(uint64(m.Roomid))
	}
	if m.GiftId != 0 {
		n += 1 + sovSmalltv(uint64(m.GiftId))
	}
	if m.Num != 0 {
		n += 1 + sovSmalltv(uint64(m.Num))
	}
	if m.Tid != 0 {
		n += 1 + sovSmalltv(uint64(m.Tid))
	}
	if m.StyleId != 0 {
		n += 1 + sovSmalltv(uint64(m.StyleId))
	}
	return n
}

func (m *SmalltvStartResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovSmalltv(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovSmalltv(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovSmalltv(uint64(l))
		}
	}
	return n
}

func (m *SmalltvStartResp_MSG) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cmd)
	if l > 0 {
		n += 1 + l + sovSmalltv(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovSmalltv(uint64(l))
	}
	l = len(m.MsgText)
	if l > 0 {
		n += 1 + l + sovSmalltv(uint64(l))
	}
	l = len(m.MsgCommon)
	if l > 0 {
		n += 1 + l + sovSmalltv(uint64(l))
	}
	l = len(m.MsgSelf)
	if l > 0 {
		n += 1 + l + sovSmalltv(uint64(l))
	}
	if m.Rep != 0 {
		n += 1 + sovSmalltv(uint64(m.Rep))
	}
	if m.StyleType != 0 {
		n += 1 + sovSmalltv(uint64(m.StyleType))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovSmalltv(uint64(l))
	}
	if m.Roomid != 0 {
		n += 1 + sovSmalltv(uint64(m.Roomid))
	}
	if m.RealRoomid != 0 {
		n += 1 + sovSmalltv(uint64(m.RealRoomid))
	}
	if m.Rnd != 0 {
		n += 1 + sovSmalltv(uint64(m.Rnd))
	}
	if m.BroadcastType != 0 {
		n += 1 + sovSmalltv(uint64(m.BroadcastType))
	}
	return n
}

func sovSmalltv(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSmalltv(x uint64) (n int) {
	return sovSmalltv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SmalltvStartReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmalltv
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SmalltvStartReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmalltvStartReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roomid", wireType)
			}
			m.Roomid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Roomid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GiftId", wireType)
			}
			m.GiftId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GiftId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tid", wireType)
			}
			m.Tid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StyleId", wireType)
			}
			m.StyleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StyleId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSmalltv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmalltv
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SmalltvStartResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmalltv
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SmalltvStartResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmalltvStartResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &SmalltvStartResp_MSG{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSmalltv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmalltv
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SmalltvStartResp_MSG) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmalltv
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MSG: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cmd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgCommon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgCommon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgSelf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgSelf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rep", wireType)
			}
			m.Rep = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rep |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StyleType", wireType)
			}
			m.StyleType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StyleType |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSmalltv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roomid", wireType)
			}
			m.Roomid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Roomid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RealRoomid", wireType)
			}
			m.RealRoomid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RealRoomid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rnd", wireType)
			}
			m.Rnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rnd |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BroadcastType", wireType)
			}
			m.BroadcastType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BroadcastType |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSmalltv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmalltv
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSmalltv(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSmalltv
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSmalltv
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSmalltv
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSmalltv
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSmalltv(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSmalltv = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSmalltv   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v0/Smalltv.proto", fileDescriptor_Smalltv_627605adbee0258c) }

var fileDescriptor_Smalltv_627605adbee0258c = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x8e, 0x63, 0x82, 0xf1, 0x92, 0x3f, 0xed, 0xc9, 0x41, 0xa9, 0x1d, 0xa1, 0x4a, 0x89, 0x54,
	0xc5, 0xa0, 0xf4, 0x54, 0xb5, 0x27, 0x7a, 0x40, 0x39, 0xe4, 0xb2, 0xe4, 0xd4, 0x0b, 0x32, 0xde,
	0xc5, 0xb5, 0xe4, 0x65, 0x89, 0xbd, 0x46, 0xc9, 0x5b, 0xf4, 0xb1, 0x2a, 0xf5, 0x92, 0x4b, 0xab,
	0x9e, 0xac, 0x0a, 0x6e, 0x7e, 0x8a, 0x6a, 0xc6, 0x26, 0xd0, 0x48, 0xf4, 0xf2, 0x31, 0xf3, 0xcd,
	0xf8, 0x9b, 0x9f, 0x1d, 0xc8, 0xe9, 0xa2, 0xdf, 0x1b, 0xc9, 0x20, 0x49, 0xf4, 0xc2, 0x9f, 0xa7,
	0x4a, 0x2b, 0x6a, 0x45, 0xf1, 0x54, 0xfb, 0x8b, 0x7e, 0xe7, 0x3a, 0x8a, 0xf5, 0xd7, 0x7c, 0xe2,
	0x87, 0x4a, 0xf6, 0x22, 0x15, 0xa9, 0x1e, 0xc6, 0x27, 0xf9, 0x14, 0x3d, 0x74, 0xd0, 0xaa, 0xbe,
	0xeb, 0xfe, 0x34, 0xc8, 0x49, 0xad, 0x34, 0xd2, 0x41, 0xaa, 0x99, 0x78, 0xa0, 0x67, 0xc4, 0xcc,
	0x63, 0xee, 0x18, 0x17, 0xc6, 0x95, 0x39, 0xb0, 0xca, 0xc2, 0x03, 0x97, 0x01, 0xd0, 0x2e, 0x69,
	0xa6, 0x4a, 0xc9, 0x98, 0x3b, 0xfb, 0x18, 0x25, 0x65, 0xe1, 0xd5, 0x0c, 0xab, 0x7f, 0xe9, 0x5b,
	0x82, 0xcd, 0x8c, 0x63, 0xee, 0x98, 0x98, 0xd4, 0x2e, 0x0b, 0x6f, 0x4d, 0xb1, 0x26, 0x18, 0xb7,
	0x1c, 0x8a, 0xcc, 0x72, 0xe9, 0x34, 0x36, 0x45, 0x66, 0xb9, 0x64, 0x00, 0x10, 0xd2, 0x31, 0x77,
	0x0e, 0x36, 0x21, 0x0d, 0xf5, 0x75, 0xcc, 0xe9, 0x25, 0x69, 0x65, 0xfa, 0x29, 0x11, 0x20, 0xde,
	0xc4, 0xf8, 0x61, 0x59, 0x78, 0x2f, 0x1c, 0xb3, 0xd0, 0xba, 0xe5, 0xdd, 0x5f, 0x0d, 0x72, 0xfa,
	0xef, 0x5c, 0xd9, 0x9c, 0x9e, 0x93, 0x46, 0xa8, 0xb8, 0xa8, 0x27, 0x6b, 0x95, 0x85, 0x87, 0x3e,
	0x43, 0x84, 0xb2, 0x32, 0x8b, 0x70, 0x30, 0xbb, 0x2a, 0x2b, 0xb3, 0x88, 0x01, 0xd0, 0x8f, 0xa4,
	0xc1, 0x03, 0x1d, 0x38, 0xe6, 0x85, 0x79, 0xd5, 0xbe, 0x79, 0xe3, 0xd7, 0xcb, 0xf6, 0x5f, 0x57,
	0xf0, 0xef, 0x46, 0xc3, 0x4a, 0x17, 0xd2, 0x19, 0x62, 0xe7, 0x87, 0x49, 0xcc, 0xbb, 0xd1, 0x10,
	0xf4, 0x43, 0x59, 0xad, 0xb5, 0xd6, 0x0f, 0x25, 0x67, 0x00, 0xff, 0x2b, 0x7d, 0x49, 0x5a, 0x32,
	0x8b, 0xc6, 0x5a, 0x3c, 0x6a, 0x5c, 0xa7, 0x5d, 0x4d, 0xbc, 0xe6, 0x98, 0x25, 0xb3, 0xe8, 0x5e,
	0x3c, 0x6a, 0x7a, 0x4d, 0x08, 0x90, 0xa1, 0x92, 0x52, 0xcd, 0x70, 0xaf, 0xf6, 0xe0, 0xb8, 0x2c,
	0xbc, 0x2d, 0x96, 0xd9, 0x32, 0x8b, 0x3e, 0xa3, 0xb9, 0xd6, 0xcd, 0x44, 0x32, 0xc5, 0x4d, 0x6f,
	0xe9, 0x02, 0x87, 0xba, 0x23, 0x91, 0x4c, 0xa1, 0xb7, 0x54, 0xcc, 0xeb, 0x6d, 0x63, 0x6f, 0xa9,
	0x98, 0x33, 0x00, 0xfa, 0x8e, 0xd8, 0xb8, 0xef, 0xfb, 0xa7, 0xb9, 0x70, 0x2c, 0x4c, 0x38, 0x2a,
	0x0b, 0x6f, 0x43, 0xb2, 0x8d, 0x89, 0x57, 0x95, 0x26, 0x4e, 0x6b, 0x33, 0x63, 0x9e, 0x26, 0x0c,
	0x60, 0xeb, 0xaa, 0xec, 0x9d, 0x57, 0xd5, 0x27, 0xed, 0x54, 0x04, 0xc9, 0xb8, 0x4e, 0x24, 0x98,
	0x78, 0x52, 0x16, 0xde, 0x36, 0xcd, 0x08, 0x38, 0xac, 0xfa, 0x02, 0x1a, 0x9f, 0x71, 0xa7, 0xbd,
	0xd5, 0xf8, 0x8c, 0x33, 0x00, 0xfa, 0x81, 0x1c, 0x4f, 0x52, 0x15, 0xf0, 0x30, 0xc8, 0xf4, 0x58,
	0x43, 0xf7, 0x87, 0x98, 0x45, 0xcb, 0xc2, 0x7b, 0x15, 0x61, 0x47, 0x2f, 0x3e, 0x8c, 0x71, 0x33,
	0x24, 0x56, 0xfd, 0xea, 0xf4, 0x13, 0x39, 0xc8, 0xe0, 0xe5, 0xa9, 0xb3, 0xe3, 0x20, 0x1e, 0x3a,
	0x67, 0x3b, 0x4f, 0x65, 0x70, 0xfe, 0x7d, 0xe9, 0x1a, 0xcf, 0x4b, 0xd7, 0xf8, 0xb3, 0x74, 0x8d,
	0x6f, 0x2b, 0x77, 0xef, 0x79, 0xe5, 0xee, 0xfd, 0x5e, 0xb9, 0x7b, 0x5f, 0xf6, 0x17, 0xfd, 0x49,
	0x13, 0xff, 0x9e, 0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x53, 0xf7, 0x74, 0xea, 0x03,
	0x00, 0x00,
}
