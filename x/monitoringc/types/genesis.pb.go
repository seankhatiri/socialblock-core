// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monitoringc/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the monitoringc module's genesis state.
type GenesisState struct {
	PortId                        string                         `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	VerifiedClientIDs             []VerifiedClientID             `protobuf:"bytes,2,rep,name=verifiedClientIDs,proto3" json:"verifiedClientIDs"`
	ProviderClientIDs             []ProviderClientID             `protobuf:"bytes,3,rep,name=providerClientIDs,proto3" json:"providerClientIDs"`
	LaunchIDsFromVerifiedClientID []LaunchIDFromVerifiedClientID `protobuf:"bytes,4,rep,name=launchIDsFromVerifiedClientID,proto3" json:"launchIDsFromVerifiedClientID"`
	LaunchIDsFromChannelID        []LaunchIDFromChannelID        `protobuf:"bytes,5,rep,name=launchIDsFromChannelID,proto3" json:"launchIDsFromChannelID"`
	MonitoringHistories           []MonitoringHistory            `protobuf:"bytes,6,rep,name=monitoringHistories,proto3" json:"monitoringHistories"`
	Params                        Params                         `protobuf:"bytes,7,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c269cb89bd5f03, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetVerifiedClientIDs() []VerifiedClientID {
	if m != nil {
		return m.VerifiedClientIDs
	}
	return nil
}

func (m *GenesisState) GetProviderClientIDs() []ProviderClientID {
	if m != nil {
		return m.ProviderClientIDs
	}
	return nil
}

func (m *GenesisState) GetLaunchIDsFromVerifiedClientID() []LaunchIDFromVerifiedClientID {
	if m != nil {
		return m.LaunchIDsFromVerifiedClientID
	}
	return nil
}

func (m *GenesisState) GetLaunchIDsFromChannelID() []LaunchIDFromChannelID {
	if m != nil {
		return m.LaunchIDsFromChannelID
	}
	return nil
}

func (m *GenesisState) GetMonitoringHistories() []MonitoringHistory {
	if m != nil {
		return m.MonitoringHistories
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.monitoringc.GenesisState")
}

func init() { proto.RegisterFile("monitoringc/genesis.proto", fileDescriptor_10c269cb89bd5f03) }

var fileDescriptor_10c269cb89bd5f03 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8f, 0xd2, 0x40,
	0x18, 0x86, 0x5b, 0x81, 0x12, 0x07, 0x2f, 0x56, 0xa3, 0xb5, 0x89, 0x95, 0x10, 0x0f, 0x24, 0x4a,
	0x1b, 0xe1, 0xe2, 0xd1, 0x00, 0x11, 0x9b, 0x68, 0x62, 0x30, 0xf1, 0xe0, 0xa5, 0x96, 0x76, 0x68,
	0x27, 0xa1, 0x33, 0xcd, 0xcc, 0x40, 0xe4, 0xee, 0xd1, 0x83, 0x3f, 0x8b, 0x23, 0x47, 0x4f, 0x66,
	0x03, 0x7f, 0x64, 0xd3, 0x69, 0xd9, 0x76, 0x0b, 0xed, 0xee, 0x6d, 0x08, 0xef, 0xfb, 0x3c, 0x5f,
	0xbe, 0x99, 0x82, 0x17, 0x11, 0xc1, 0x88, 0x13, 0x8a, 0x70, 0xe0, 0x59, 0x01, 0xc4, 0x90, 0x21,
	0x66, 0xc6, 0x94, 0x70, 0xa2, 0xea, 0x1c, 0x62, 0x1f, 0xd2, 0x08, 0x61, 0x6e, 0xb2, 0x18, 0x9b,
	0x85, 0xa4, 0xfe, 0x34, 0x20, 0x01, 0x11, 0x31, 0x2b, 0x39, 0xa5, 0x0d, 0x5d, 0x2b, 0xc2, 0x62,
	0x97, 0xba, 0x51, 0xc6, 0xd2, 0x5f, 0x17, 0xff, 0xd9, 0x40, 0x8a, 0x96, 0x08, 0xfa, 0x8e, 0xb7,
	0x42, 0x10, 0x73, 0x07, 0xf9, 0x97, 0x52, 0x31, 0x25, 0x1b, 0xe4, 0x43, 0x7a, 0x96, 0x1a, 0x15,
	0x53, 0x2b, 0x77, 0x8d, 0xbd, 0xd0, 0x41, 0xbe, 0xb3, 0xa4, 0x24, 0x72, 0x2a, 0xd1, 0x6f, 0x6a,
	0x4a, 0x5e, 0xe8, 0x62, 0x0c, 0x57, 0x15, 0x73, 0xe4, 0x67, 0x27, 0x44, 0x8c, 0x13, 0xba, 0x4d,
	0x53, 0xbd, 0x3f, 0x2d, 0xf0, 0x68, 0x96, 0x6e, 0xec, 0x1b, 0x77, 0x39, 0x54, 0x9f, 0x83, 0x76,
	0x4c, 0x68, 0x22, 0xd5, 0xe4, 0xae, 0xdc, 0x7f, 0x38, 0x57, 0x92, 0x9f, 0xb6, 0xaf, 0xfe, 0x04,
	0x8f, 0x4f, 0x83, 0x4d, 0xc4, 0x5c, 0xf6, 0x94, 0x69, 0x0f, 0xba, 0x8d, 0x7e, 0x67, 0xf8, 0xd6,
	0xac, 0xde, 0xb2, 0xf9, 0xbd, 0x54, 0x1a, 0x37, 0x77, 0xff, 0x5f, 0x49, 0xf3, 0x73, 0x58, 0x62,
	0x38, 0xed, 0x2b, 0x37, 0x34, 0xee, 0x36, 0x7c, 0x2d, 0x95, 0x4e, 0x86, 0x33, 0x98, 0xfa, 0x5b,
	0x06, 0x2f, 0xd3, 0xbd, 0xd9, 0x53, 0xf6, 0x91, 0x92, 0xa8, 0x3c, 0x9c, 0xd6, 0x14, 0xba, 0xf7,
	0x75, 0xba, 0xcf, 0x19, 0xe0, 0x52, 0x3f, 0x53, 0xd7, 0x4b, 0x54, 0x02, 0x9e, 0xdd, 0x0a, 0x4c,
	0xd2, 0xbb, 0xb3, 0xa7, 0x5a, 0x4b, 0xe8, 0xdf, 0xdd, 0x57, 0x7f, 0x53, 0xcc, 0xbc, 0x15, 0x58,
	0x15, 0x82, 0x27, 0x39, 0xe2, 0x93, 0x78, 0x00, 0x08, 0x32, 0x4d, 0x11, 0xb6, 0x41, 0x9d, 0xed,
	0x4b, 0xa9, 0xb6, 0xcd, 0x4c, 0x97, 0x78, 0xea, 0x07, 0xa0, 0xa4, 0x1f, 0x8c, 0xd6, 0xee, 0xca,
	0xfd, 0xce, 0xb0, 0x57, 0x7b, 0x6b, 0x22, 0x99, 0xe1, 0xb2, 0xde, 0x78, 0xb6, 0x3b, 0x18, 0xf2,
	0xfe, 0x60, 0xc8, 0x57, 0x07, 0x43, 0xfe, 0x7b, 0x34, 0xa4, 0xfd, 0xd1, 0x90, 0xfe, 0x1d, 0x0d,
	0xe9, 0xc7, 0x20, 0x40, 0x3c, 0x5c, 0x2f, 0x4c, 0x8f, 0x44, 0x56, 0x4e, 0xb5, 0x58, 0x8c, 0xad,
	0x5f, 0x56, 0xf1, 0xa9, 0xf3, 0x6d, 0x0c, 0xd9, 0x42, 0x11, 0xcf, 0x7b, 0x74, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0xe9, 0x4c, 0x80, 0x1b, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.MonitoringHistories) > 0 {
		for iNdEx := len(m.MonitoringHistories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MonitoringHistories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LaunchIDsFromChannelID) > 0 {
		for iNdEx := len(m.LaunchIDsFromChannelID) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDsFromChannelID[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LaunchIDsFromVerifiedClientID) > 0 {
		for iNdEx := len(m.LaunchIDsFromVerifiedClientID) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDsFromVerifiedClientID[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ProviderClientIDs) > 0 {
		for iNdEx := len(m.ProviderClientIDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProviderClientIDs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VerifiedClientIDs) > 0 {
		for iNdEx := len(m.VerifiedClientIDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerifiedClientIDs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.VerifiedClientIDs) > 0 {
		for _, e := range m.VerifiedClientIDs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProviderClientIDs) > 0 {
		for _, e := range m.ProviderClientIDs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LaunchIDsFromVerifiedClientID) > 0 {
		for _, e := range m.LaunchIDsFromVerifiedClientID {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LaunchIDsFromChannelID) > 0 {
		for _, e := range m.LaunchIDsFromChannelID {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MonitoringHistories) > 0 {
		for _, e := range m.MonitoringHistories {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifiedClientIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifiedClientIDs = append(m.VerifiedClientIDs, VerifiedClientID{})
			if err := m.VerifiedClientIDs[len(m.VerifiedClientIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderClientIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderClientIDs = append(m.ProviderClientIDs, ProviderClientID{})
			if err := m.ProviderClientIDs[len(m.ProviderClientIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDsFromVerifiedClientID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LaunchIDsFromVerifiedClientID = append(m.LaunchIDsFromVerifiedClientID, LaunchIDFromVerifiedClientID{})
			if err := m.LaunchIDsFromVerifiedClientID[len(m.LaunchIDsFromVerifiedClientID)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDsFromChannelID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LaunchIDsFromChannelID = append(m.LaunchIDsFromChannelID, LaunchIDFromChannelID{})
			if err := m.LaunchIDsFromChannelID[len(m.LaunchIDsFromChannelID)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoringHistories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MonitoringHistories = append(m.MonitoringHistories, MonitoringHistory{})
			if err := m.MonitoringHistories[len(m.MonitoringHistories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
