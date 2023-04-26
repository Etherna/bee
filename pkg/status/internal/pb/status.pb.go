// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: status.proto

package pb

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Get message indicate interest in receiving a node Snapshot.
type Get struct {
}

func (m *Get) Reset()         { *m = Get{} }
func (m *Get) String() string { return proto.CompactTextString(m) }
func (*Get) ProtoMessage()    {}
func (*Get) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}
func (m *Get) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Get) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Get.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Get) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Get.Merge(m, src)
}
func (m *Get) XXX_Size() int {
	return m.Size()
}
func (m *Get) XXX_DiscardUnknown() {
	xxx_messageInfo_Get.DiscardUnknown(m)
}

var xxx_messageInfo_Get proto.InternalMessageInfo

// Snapshot message is a response to the Get message and contains
// the appropriate values that are a snapshot of the current state
// of the running node.
type Snapshot struct {
	ReserveSize      uint64  `protobuf:"varint,1,opt,name=ReserveSize,proto3" json:"ReserveSize,omitempty"`
	PullsyncRate     float64 `protobuf:"fixed64,2,opt,name=PullsyncRate,proto3" json:"PullsyncRate,omitempty"`
	StorageRadius    uint32  `protobuf:"varint,3,opt,name=StorageRadius,proto3" json:"StorageRadius,omitempty"`
	ConnectedPeers   uint64  `protobuf:"varint,4,opt,name=ConnectedPeers,proto3" json:"ConnectedPeers,omitempty"`
	NeighborhoodSize uint64  `protobuf:"varint,5,opt,name=NeighborhoodSize,proto3" json:"NeighborhoodSize,omitempty"`
	BeeMode          string  `protobuf:"bytes,6,opt,name=BeeMode,proto3" json:"BeeMode,omitempty"`
	BatchTotalAmount string  `protobuf:"bytes,7,opt,name=BatchTotalAmount,proto3" json:"BatchTotalAmount,omitempty"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{1}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetReserveSize() uint64 {
	if m != nil {
		return m.ReserveSize
	}
	return 0
}

func (m *Snapshot) GetPullsyncRate() float64 {
	if m != nil {
		return m.PullsyncRate
	}
	return 0
}

func (m *Snapshot) GetStorageRadius() uint32 {
	if m != nil {
		return m.StorageRadius
	}
	return 0
}

func (m *Snapshot) GetConnectedPeers() uint64 {
	if m != nil {
		return m.ConnectedPeers
	}
	return 0
}

func (m *Snapshot) GetNeighborhoodSize() uint64 {
	if m != nil {
		return m.NeighborhoodSize
	}
	return 0
}

func (m *Snapshot) GetBeeMode() string {
	if m != nil {
		return m.BeeMode
	}
	return ""
}

func (m *Snapshot) GetBatchTotalAmount() string {
	if m != nil {
		return m.BatchTotalAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*Get)(nil), "status.Get")
	proto.RegisterType((*Snapshot)(nil), "status.Snapshot")
}

func init() { proto.RegisterFile("status.proto", fileDescriptor_dfe4fce6682daf5b) }

var fileDescriptor_dfe4fce6682daf5b = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd0, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0xe0, 0x6c, 0xda, 0xa6, 0xba, 0xb6, 0x22, 0x7b, 0xda, 0x83, 0x2c, 0x21, 0x88, 0x04,
	0x0f, 0x5e, 0x7c, 0x02, 0xe3, 0xc1, 0x93, 0x52, 0x36, 0x9e, 0xbc, 0x6d, 0x92, 0xa1, 0x09, 0xc4,
	0x4c, 0xc8, 0x4e, 0x04, 0x7d, 0x8a, 0x3e, 0x96, 0xc7, 0x1e, 0x3d, 0x4a, 0xf2, 0x22, 0xe2, 0x8a,
	0x60, 0xdb, 0xe3, 0xff, 0xf1, 0x33, 0x33, 0x0c, 0x5f, 0x58, 0x32, 0xd4, 0xdb, 0xeb, 0xb6, 0x43,
	0x42, 0x11, 0xfc, 0xa6, 0x68, 0xc6, 0x27, 0xf7, 0x40, 0xd1, 0xc6, 0xe7, 0x47, 0x69, 0x63, 0x5a,
	0x5b, 0x22, 0x89, 0x90, 0x9f, 0x68, 0xb0, 0xd0, 0xbd, 0x42, 0x5a, 0xbd, 0x83, 0x64, 0x21, 0x8b,
	0xa7, 0xfa, 0x3f, 0x89, 0x88, 0x2f, 0x56, 0x7d, 0x5d, 0xdb, 0xb7, 0x26, 0xd7, 0x86, 0x40, 0xfa,
	0x21, 0x8b, 0x99, 0xde, 0x31, 0x71, 0xc1, 0x97, 0x29, 0x61, 0x67, 0xd6, 0xa0, 0x4d, 0x51, 0xf5,
	0x56, 0x4e, 0x42, 0x16, 0x2f, 0xf5, 0x2e, 0x8a, 0x4b, 0x7e, 0x7a, 0x87, 0x4d, 0x03, 0x39, 0x41,
	0xb1, 0x02, 0xe8, 0xac, 0x9c, 0xba, 0x75, 0x7b, 0x2a, 0xae, 0xf8, 0xd9, 0x23, 0x54, 0xeb, 0x32,
	0xc3, 0xae, 0x44, 0x2c, 0xdc, 0x61, 0x33, 0xd7, 0x3c, 0x70, 0x21, 0xf9, 0x3c, 0x01, 0x78, 0xc0,
	0x02, 0x64, 0x10, 0xb2, 0xf8, 0x58, 0xff, 0xc5, 0x9f, 0x29, 0x89, 0xa1, 0xbc, 0x7c, 0x42, 0x32,
	0xf5, 0xed, 0x0b, 0xf6, 0x0d, 0xc9, 0xb9, 0xab, 0x1c, 0x78, 0x72, 0xfe, 0x31, 0x28, 0xb6, 0x1d,
	0x14, 0xfb, 0x1a, 0x14, 0xdb, 0x8c, 0xca, 0xdb, 0x8e, 0xca, 0xfb, 0x1c, 0x95, 0xf7, 0xec, 0xb7,
	0x59, 0x16, 0xb8, 0x37, 0xde, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x2a, 0x3a, 0x95, 0x56,
	0x01, 0x00, 0x00,
}

func (m *Get) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Get) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Get) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BatchTotalAmount) > 0 {
		i -= len(m.BatchTotalAmount)
		copy(dAtA[i:], m.BatchTotalAmount)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.BatchTotalAmount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BeeMode) > 0 {
		i -= len(m.BeeMode)
		copy(dAtA[i:], m.BeeMode)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.BeeMode)))
		i--
		dAtA[i] = 0x32
	}
	if m.NeighborhoodSize != 0 {
		i = encodeVarintStatus(dAtA, i, uint64(m.NeighborhoodSize))
		i--
		dAtA[i] = 0x28
	}
	if m.ConnectedPeers != 0 {
		i = encodeVarintStatus(dAtA, i, uint64(m.ConnectedPeers))
		i--
		dAtA[i] = 0x20
	}
	if m.StorageRadius != 0 {
		i = encodeVarintStatus(dAtA, i, uint64(m.StorageRadius))
		i--
		dAtA[i] = 0x18
	}
	if m.PullsyncRate != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.PullsyncRate))))
		i--
		dAtA[i] = 0x11
	}
	if m.ReserveSize != 0 {
		i = encodeVarintStatus(dAtA, i, uint64(m.ReserveSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Get) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReserveSize != 0 {
		n += 1 + sovStatus(uint64(m.ReserveSize))
	}
	if m.PullsyncRate != 0 {
		n += 9
	}
	if m.StorageRadius != 0 {
		n += 1 + sovStatus(uint64(m.StorageRadius))
	}
	if m.ConnectedPeers != 0 {
		n += 1 + sovStatus(uint64(m.ConnectedPeers))
	}
	if m.NeighborhoodSize != 0 {
		n += 1 + sovStatus(uint64(m.NeighborhoodSize))
	}
	l = len(m.BeeMode)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.BatchTotalAmount)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	return n
}

func sovStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Get) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Get: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Get: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStatus
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
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveSize", wireType)
			}
			m.ReserveSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReserveSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field PullsyncRate", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.PullsyncRate = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageRadius", wireType)
			}
			m.StorageRadius = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StorageRadius |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectedPeers", wireType)
			}
			m.ConnectedPeers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectedPeers |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeighborhoodSize", wireType)
			}
			m.NeighborhoodSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NeighborhoodSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeeMode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeeMode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchTotalAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchTotalAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStatus
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
func skipStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStatus
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
			if length < 0 {
				return 0, ErrInvalidLengthStatus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStatus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStatus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStatus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStatus = fmt.Errorf("proto: unexpected end of group")
)