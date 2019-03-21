// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/dbnode/generated/proto/namespace/namespace.proto

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
	Package namespace is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/dbnode/generated/proto/namespace/namespace.proto

	It has these top-level messages:
		RetentionOptions
		IndexOptions
		NamespaceOptions
		Registry
*/
package namespace

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type RetentionOptions struct {
	RetentionPeriodNanos                     int64 `protobuf:"varint,1,opt,name=retentionPeriodNanos,proto3" json:"retentionPeriodNanos,omitempty"`
	BlockSizeNanos                           int64 `protobuf:"varint,2,opt,name=blockSizeNanos,proto3" json:"blockSizeNanos,omitempty"`
	BufferFutureNanos                        int64 `protobuf:"varint,3,opt,name=bufferFutureNanos,proto3" json:"bufferFutureNanos,omitempty"`
	BufferPastNanos                          int64 `protobuf:"varint,4,opt,name=bufferPastNanos,proto3" json:"bufferPastNanos,omitempty"`
	BlockDataExpiry                          bool  `protobuf:"varint,5,opt,name=blockDataExpiry,proto3" json:"blockDataExpiry,omitempty"`
	BlockDataExpiryAfterNotAccessPeriodNanos int64 `protobuf:"varint,6,opt,name=blockDataExpiryAfterNotAccessPeriodNanos,proto3" json:"blockDataExpiryAfterNotAccessPeriodNanos,omitempty"`
}

func (m *RetentionOptions) Reset()                    { *m = RetentionOptions{} }
func (m *RetentionOptions) String() string            { return proto.CompactTextString(m) }
func (*RetentionOptions) ProtoMessage()               {}
func (*RetentionOptions) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{0} }

func (m *RetentionOptions) GetRetentionPeriodNanos() int64 {
	if m != nil {
		return m.RetentionPeriodNanos
	}
	return 0
}

func (m *RetentionOptions) GetBlockSizeNanos() int64 {
	if m != nil {
		return m.BlockSizeNanos
	}
	return 0
}

func (m *RetentionOptions) GetBufferFutureNanos() int64 {
	if m != nil {
		return m.BufferFutureNanos
	}
	return 0
}

func (m *RetentionOptions) GetBufferPastNanos() int64 {
	if m != nil {
		return m.BufferPastNanos
	}
	return 0
}

func (m *RetentionOptions) GetBlockDataExpiry() bool {
	if m != nil {
		return m.BlockDataExpiry
	}
	return false
}

func (m *RetentionOptions) GetBlockDataExpiryAfterNotAccessPeriodNanos() int64 {
	if m != nil {
		return m.BlockDataExpiryAfterNotAccessPeriodNanos
	}
	return 0
}

type IndexOptions struct {
	Enabled        bool  `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	BlockSizeNanos int64 `protobuf:"varint,2,opt,name=blockSizeNanos,proto3" json:"blockSizeNanos,omitempty"`
}

func (m *IndexOptions) Reset()                    { *m = IndexOptions{} }
func (m *IndexOptions) String() string            { return proto.CompactTextString(m) }
func (*IndexOptions) ProtoMessage()               {}
func (*IndexOptions) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{1} }

func (m *IndexOptions) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *IndexOptions) GetBlockSizeNanos() int64 {
	if m != nil {
		return m.BlockSizeNanos
	}
	return 0
}

type NamespaceOptions struct {
	BootstrapEnabled  bool              `protobuf:"varint,1,opt,name=bootstrapEnabled,proto3" json:"bootstrapEnabled,omitempty"`
	FlushEnabled      bool              `protobuf:"varint,2,opt,name=flushEnabled,proto3" json:"flushEnabled,omitempty"`
	WritesToCommitLog bool              `protobuf:"varint,3,opt,name=writesToCommitLog,proto3" json:"writesToCommitLog,omitempty"`
	CleanupEnabled    bool              `protobuf:"varint,4,opt,name=cleanupEnabled,proto3" json:"cleanupEnabled,omitempty"`
	RepairEnabled     bool              `protobuf:"varint,5,opt,name=repairEnabled,proto3" json:"repairEnabled,omitempty"`
	RetentionOptions  *RetentionOptions `protobuf:"bytes,6,opt,name=retentionOptions" json:"retentionOptions,omitempty"`
	SnapshotEnabled   bool              `protobuf:"varint,7,opt,name=snapshotEnabled,proto3" json:"snapshotEnabled,omitempty"`
	IndexOptions      *IndexOptions     `protobuf:"bytes,8,opt,name=indexOptions" json:"indexOptions,omitempty"`
	Schema            []byte            `protobuf:"bytes,9,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (m *NamespaceOptions) Reset()                    { *m = NamespaceOptions{} }
func (m *NamespaceOptions) String() string            { return proto.CompactTextString(m) }
func (*NamespaceOptions) ProtoMessage()               {}
func (*NamespaceOptions) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{2} }

func (m *NamespaceOptions) GetBootstrapEnabled() bool {
	if m != nil {
		return m.BootstrapEnabled
	}
	return false
}

func (m *NamespaceOptions) GetFlushEnabled() bool {
	if m != nil {
		return m.FlushEnabled
	}
	return false
}

func (m *NamespaceOptions) GetWritesToCommitLog() bool {
	if m != nil {
		return m.WritesToCommitLog
	}
	return false
}

func (m *NamespaceOptions) GetCleanupEnabled() bool {
	if m != nil {
		return m.CleanupEnabled
	}
	return false
}

func (m *NamespaceOptions) GetRepairEnabled() bool {
	if m != nil {
		return m.RepairEnabled
	}
	return false
}

func (m *NamespaceOptions) GetRetentionOptions() *RetentionOptions {
	if m != nil {
		return m.RetentionOptions
	}
	return nil
}

func (m *NamespaceOptions) GetSnapshotEnabled() bool {
	if m != nil {
		return m.SnapshotEnabled
	}
	return false
}

func (m *NamespaceOptions) GetIndexOptions() *IndexOptions {
	if m != nil {
		return m.IndexOptions
	}
	return nil
}

func (m *NamespaceOptions) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type Registry struct {
	Namespaces map[string]*NamespaceOptions `protobuf:"bytes,1,rep,name=namespaces" json:"namespaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Registry) Reset()                    { *m = Registry{} }
func (m *Registry) String() string            { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()               {}
func (*Registry) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{3} }

func (m *Registry) GetNamespaces() map[string]*NamespaceOptions {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*RetentionOptions)(nil), "namespace.RetentionOptions")
	proto.RegisterType((*IndexOptions)(nil), "namespace.IndexOptions")
	proto.RegisterType((*NamespaceOptions)(nil), "namespace.NamespaceOptions")
	proto.RegisterType((*Registry)(nil), "namespace.Registry")
}
func (m *RetentionOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetentionOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RetentionPeriodNanos != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.RetentionPeriodNanos))
	}
	if m.BlockSizeNanos != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.BlockSizeNanos))
	}
	if m.BufferFutureNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.BufferFutureNanos))
	}
	if m.BufferPastNanos != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.BufferPastNanos))
	}
	if m.BlockDataExpiry {
		dAtA[i] = 0x28
		i++
		if m.BlockDataExpiry {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.BlockDataExpiryAfterNotAccessPeriodNanos != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.BlockDataExpiryAfterNotAccessPeriodNanos))
	}
	return i, nil
}

func (m *IndexOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Enabled {
		dAtA[i] = 0x8
		i++
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.BlockSizeNanos != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.BlockSizeNanos))
	}
	return i, nil
}

func (m *NamespaceOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BootstrapEnabled {
		dAtA[i] = 0x8
		i++
		if m.BootstrapEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.FlushEnabled {
		dAtA[i] = 0x10
		i++
		if m.FlushEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.WritesToCommitLog {
		dAtA[i] = 0x18
		i++
		if m.WritesToCommitLog {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CleanupEnabled {
		dAtA[i] = 0x20
		i++
		if m.CleanupEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RepairEnabled {
		dAtA[i] = 0x28
		i++
		if m.RepairEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RetentionOptions != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.RetentionOptions.Size()))
		n1, err := m.RetentionOptions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.SnapshotEnabled {
		dAtA[i] = 0x38
		i++
		if m.SnapshotEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IndexOptions != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.IndexOptions.Size()))
		n2, err := m.IndexOptions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Schema) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Schema)))
		i += copy(dAtA[i:], m.Schema)
	}
	return i, nil
}

func (m *Registry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Registry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for k, _ := range m.Namespaces {
			dAtA[i] = 0xa
			i++
			v := m.Namespaces[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovNamespace(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovNamespace(uint64(len(k))) + msgSize
			i = encodeVarintNamespace(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintNamespace(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintNamespace(dAtA, i, uint64(v.Size()))
				n3, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n3
			}
		}
	}
	return i, nil
}

func encodeVarintNamespace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RetentionOptions) Size() (n int) {
	var l int
	_ = l
	if m.RetentionPeriodNanos != 0 {
		n += 1 + sovNamespace(uint64(m.RetentionPeriodNanos))
	}
	if m.BlockSizeNanos != 0 {
		n += 1 + sovNamespace(uint64(m.BlockSizeNanos))
	}
	if m.BufferFutureNanos != 0 {
		n += 1 + sovNamespace(uint64(m.BufferFutureNanos))
	}
	if m.BufferPastNanos != 0 {
		n += 1 + sovNamespace(uint64(m.BufferPastNanos))
	}
	if m.BlockDataExpiry {
		n += 2
	}
	if m.BlockDataExpiryAfterNotAccessPeriodNanos != 0 {
		n += 1 + sovNamespace(uint64(m.BlockDataExpiryAfterNotAccessPeriodNanos))
	}
	return n
}

func (m *IndexOptions) Size() (n int) {
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.BlockSizeNanos != 0 {
		n += 1 + sovNamespace(uint64(m.BlockSizeNanos))
	}
	return n
}

func (m *NamespaceOptions) Size() (n int) {
	var l int
	_ = l
	if m.BootstrapEnabled {
		n += 2
	}
	if m.FlushEnabled {
		n += 2
	}
	if m.WritesToCommitLog {
		n += 2
	}
	if m.CleanupEnabled {
		n += 2
	}
	if m.RepairEnabled {
		n += 2
	}
	if m.RetentionOptions != nil {
		l = m.RetentionOptions.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	if m.SnapshotEnabled {
		n += 2
	}
	if m.IndexOptions != nil {
		l = m.IndexOptions.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *Registry) Size() (n int) {
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for k, v := range m.Namespaces {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovNamespace(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovNamespace(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovNamespace(uint64(mapEntrySize))
		}
	}
	return n
}

func sovNamespace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNamespace(x uint64) (n int) {
	return sovNamespace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RetentionOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: RetentionOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetentionOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetentionPeriodNanos", wireType)
			}
			m.RetentionPeriodNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetentionPeriodNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockSizeNanos", wireType)
			}
			m.BlockSizeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockSizeNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferFutureNanos", wireType)
			}
			m.BufferFutureNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BufferFutureNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferPastNanos", wireType)
			}
			m.BufferPastNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BufferPastNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockDataExpiry", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BlockDataExpiry = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockDataExpiryAfterNotAccessPeriodNanos", wireType)
			}
			m.BlockDataExpiryAfterNotAccessPeriodNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockDataExpiryAfterNotAccessPeriodNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *IndexOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: IndexOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockSizeNanos", wireType)
			}
			m.BlockSizeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockSizeNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BootstrapEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BootstrapEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlushEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FlushEnabled = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WritesToCommitLog", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WritesToCommitLog = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CleanupEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CleanupEnabled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepairEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RepairEnabled = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetentionOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RetentionOptions == nil {
				m.RetentionOptions = &RetentionOptions{}
			}
			if err := m.RetentionOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SnapshotEnabled = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IndexOptions == nil {
				m.IndexOptions = &IndexOptions{}
			}
			if err := m.IndexOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = append(m.Schema[:0], dAtA[iNdEx:postIndex]...)
			if m.Schema == nil {
				m.Schema = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *Registry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: Registry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Registry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Namespaces == nil {
				m.Namespaces = make(map[string]*NamespaceOptions)
			}
			var mapkey string
			var mapvalue *NamespaceOptions
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNamespace
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespace
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNamespace
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespace
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthNamespace
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthNamespace
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &NamespaceOptions{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNamespace(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNamespace
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Namespaces[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func skipNamespace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
				return 0, ErrInvalidLengthNamespace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNamespace
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
				next, err := skipNamespace(dAtA[start:])
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
	ErrInvalidLengthNamespace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespace   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/dbnode/generated/proto/namespace/namespace.proto", fileDescriptorNamespace)
}

var fileDescriptorNamespace = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0xd2, 0xa6, 0xc9, 0x34, 0x50, 0xb3, 0x42, 0x10, 0x81, 0x14, 0x45, 0x01, 0x21,
	0x0b, 0xa1, 0x58, 0x24, 0x17, 0x04, 0xa7, 0x52, 0x42, 0x85, 0x84, 0x42, 0xb4, 0x70, 0xea, 0x6d,
	0x6d, 0x4f, 0x12, 0xab, 0xf1, 0xae, 0xb5, 0xbb, 0x86, 0x86, 0xa7, 0xe0, 0x29, 0xb8, 0xf0, 0x22,
	0x1c, 0x38, 0xf0, 0x08, 0x28, 0xbc, 0x08, 0xf2, 0xba, 0x4e, 0xe3, 0x4d, 0x0f, 0xbd, 0x58, 0xbb,
	0xff, 0x7c, 0x3b, 0xe3, 0x9d, 0x7f, 0x6c, 0x38, 0x9d, 0xc7, 0x7a, 0x91, 0x05, 0x83, 0x50, 0x24,
	0x7e, 0x32, 0x8a, 0x02, 0x3f, 0x19, 0xf9, 0x4a, 0x86, 0x7e, 0x14, 0x70, 0x11, 0xa1, 0x3f, 0x47,
	0x8e, 0x92, 0x69, 0x8c, 0xfc, 0x54, 0x0a, 0x2d, 0x7c, 0xce, 0x12, 0x54, 0x29, 0x0b, 0xf1, 0x6a,
	0x35, 0x30, 0x11, 0xd2, 0xda, 0x08, 0xfd, 0xdf, 0x35, 0x70, 0x29, 0x6a, 0xe4, 0x3a, 0x16, 0xfc,
	0x63, 0x9a, 0x3f, 0x15, 0x19, 0xc2, 0x3d, 0x59, 0x6a, 0x53, 0x94, 0xb1, 0x88, 0x26, 0x8c, 0x0b,
	0xd5, 0x71, 0x7a, 0x8e, 0x57, 0xa7, 0xd7, 0xc6, 0xc8, 0x53, 0xb8, 0x13, 0x2c, 0x45, 0x78, 0xfe,
	0x29, 0xfe, 0x86, 0x05, 0x5d, 0x33, 0xb4, 0xa5, 0x92, 0xe7, 0x70, 0x37, 0xc8, 0x66, 0x33, 0x94,
	0xef, 0x32, 0x9d, 0xc9, 0x4b, 0xb4, 0x6e, 0xd0, 0xdd, 0x00, 0xf1, 0xe0, 0xa8, 0x10, 0xa7, 0x4c,
	0xe9, 0x82, 0xdd, 0x33, 0xac, 0x2d, 0x1b, 0x32, 0xaf, 0xf4, 0x96, 0x69, 0x36, 0xbe, 0x48, 0x63,
	0xb9, 0xea, 0xec, 0xf7, 0x1c, 0xaf, 0x49, 0x6d, 0x99, 0x9c, 0x81, 0x67, 0x49, 0xc7, 0x33, 0x8d,
	0x72, 0x22, 0xf4, 0x71, 0x18, 0xa2, 0x52, 0xdb, 0x37, 0x6e, 0x98, 0x62, 0x37, 0xe6, 0xfb, 0x53,
	0x68, 0xbf, 0xe7, 0x11, 0x5e, 0x94, 0x9d, 0xec, 0xc0, 0x01, 0x72, 0x16, 0x2c, 0x31, 0x32, 0xcd,
	0x6b, 0xd2, 0x72, 0x7b, 0xd3, 0x7e, 0xf5, 0x7f, 0xd4, 0xc1, 0x9d, 0x94, 0x76, 0x95, 0x69, 0x9f,
	0x81, 0x1b, 0x08, 0xa1, 0x95, 0x96, 0x2c, 0x1d, 0x57, 0xf2, 0xef, 0xe8, 0xa4, 0x0f, 0xed, 0xd9,
	0x32, 0x53, 0x8b, 0x92, 0xab, 0x19, 0xae, 0xa2, 0xe5, 0xa6, 0x7c, 0x95, 0xb1, 0x46, 0xf5, 0x59,
	0x9c, 0x88, 0x24, 0x89, 0xf5, 0x07, 0x31, 0x37, 0xa6, 0x34, 0xe9, 0x6e, 0x20, 0x7f, 0xf5, 0x70,
	0x89, 0x8c, 0x67, 0x9b, 0xda, 0x7b, 0x06, 0xb5, 0x54, 0xf2, 0x04, 0x6e, 0x4b, 0x4c, 0x59, 0x2c,
	0x4b, 0xac, 0x30, 0xa4, 0x2a, 0x92, 0x53, 0x70, 0xa5, 0x35, 0x80, 0xa6, 0xed, 0x87, 0xc3, 0x47,
	0x83, 0xab, 0xc1, 0xb5, 0x67, 0x94, 0xee, 0x1c, 0xca, 0x27, 0x40, 0x71, 0x96, 0xaa, 0x85, 0xd0,
	0x65, 0xc1, 0x83, 0x62, 0x02, 0x2c, 0x99, 0xbc, 0x86, 0x76, 0xbc, 0xe5, 0x52, 0xa7, 0x69, 0xca,
	0x3d, 0xd8, 0x2a, 0xb7, 0x6d, 0x22, 0xad, 0xc0, 0xe4, 0x3e, 0x34, 0x54, 0xb8, 0xc0, 0x84, 0x75,
	0x5a, 0x3d, 0xc7, 0x6b, 0xd3, 0xcb, 0x5d, 0xff, 0xa7, 0x03, 0x4d, 0x8a, 0xf3, 0x58, 0x69, 0xb9,
	0x22, 0x27, 0x00, 0x9b, 0x64, 0xf9, 0x77, 0x53, 0xf7, 0x0e, 0x87, 0x8f, 0x2b, 0xd7, 0x29, 0xc0,
	0xc1, 0xc6, 0x5a, 0x35, 0xe6, 0x5a, 0xae, 0xe8, 0xd6, 0xb1, 0x87, 0x67, 0x70, 0x64, 0x85, 0x89,
	0x0b, 0xf5, 0x73, 0x5c, 0x19, 0xaf, 0x5b, 0x34, 0x5f, 0x92, 0x17, 0xb0, 0xff, 0x85, 0x2d, 0x33,
	0x34, 0xbe, 0x56, 0x7b, 0x66, 0x8f, 0x0d, 0x2d, 0xc8, 0x57, 0xb5, 0x97, 0xce, 0x1b, 0xf7, 0xd7,
	0xba, 0xeb, 0xfc, 0x59, 0x77, 0x9d, 0xbf, 0xeb, 0xae, 0xf3, 0xfd, 0x5f, 0xf7, 0x56, 0xd0, 0x30,
	0xff, 0x86, 0xd1, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x06, 0xe8, 0x07, 0x66, 0x04, 0x00,
	0x00,
}