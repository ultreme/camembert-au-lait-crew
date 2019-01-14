// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/crew/crew.proto

package crew // import "ultre.me/calcbiz/pkg/crew"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Crew struct {
	Name     string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Website  string        `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
	Members  []*Person     `protobuf:"bytes,3,rep,name=members" json:"members,omitempty"`
	Accounts []*WebAccount `protobuf:"bytes,4,rep,name=accounts" json:"accounts,omitempty"`
	Friends  []*Friend     `protobuf:"bytes,5,rep,name=friends" json:"friends,omitempty"`
}

func (m *Crew) Reset()         { *m = Crew{} }
func (m *Crew) String() string { return proto.CompactTextString(m) }
func (*Crew) ProtoMessage()    {}
func (*Crew) Descriptor() ([]byte, []int) {
	return fileDescriptor_crew_26e212dfbbdddd43, []int{0}
}
func (m *Crew) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Crew) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Crew.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Crew) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crew.Merge(dst, src)
}
func (m *Crew) XXX_Size() int {
	return m.Size()
}
func (m *Crew) XXX_DiscardUnknown() {
	xxx_messageInfo_Crew.DiscardUnknown(m)
}

var xxx_messageInfo_Crew proto.InternalMessageInfo

func (m *Crew) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Crew) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Crew) GetMembers() []*Person {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Crew) GetAccounts() []*WebAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Crew) GetFriends() []*Friend {
	if m != nil {
		return m.Friends
	}
	return nil
}

type Person struct {
	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_crew_26e212dfbbdddd43, []int{1}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Person.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return m.Size()
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type WebAccount struct {
	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Handle   string `protobuf:"bytes,3,opt,name=handle,proto3" json:"handle,omitempty"`
	URL      string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *WebAccount) Reset()         { *m = WebAccount{} }
func (m *WebAccount) String() string { return proto.CompactTextString(m) }
func (*WebAccount) ProtoMessage()    {}
func (*WebAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_crew_26e212dfbbdddd43, []int{2}
}
func (m *WebAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WebAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WebAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *WebAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebAccount.Merge(dst, src)
}
func (m *WebAccount) XXX_Size() int {
	return m.Size()
}
func (m *WebAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_WebAccount.DiscardUnknown(m)
}

var xxx_messageInfo_WebAccount proto.InternalMessageInfo

func (m *WebAccount) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *WebAccount) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *WebAccount) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *WebAccount) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type Link struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	URL  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_crew_26e212dfbbdddd43, []int{3}
}
func (m *Link) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Link.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(dst, src)
}
func (m *Link) XXX_Size() int {
	return m.Size()
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Link) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type Friend struct {
	Key         string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	LogoURL     string  `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Links       []*Link `protobuf:"bytes,5,rep,name=links" json:"links,omitempty"`
	ImageURL    string  `protobuf:"bytes,6,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (m *Friend) Reset()         { *m = Friend{} }
func (m *Friend) String() string { return proto.CompactTextString(m) }
func (*Friend) ProtoMessage()    {}
func (*Friend) Descriptor() ([]byte, []int) {
	return fileDescriptor_crew_26e212dfbbdddd43, []int{4}
}
func (m *Friend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Friend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Friend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Friend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Friend.Merge(dst, src)
}
func (m *Friend) XXX_Size() int {
	return m.Size()
}
func (m *Friend) XXX_DiscardUnknown() {
	xxx_messageInfo_Friend.DiscardUnknown(m)
}

var xxx_messageInfo_Friend proto.InternalMessageInfo

func (m *Friend) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Friend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Friend) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Friend) GetLogoURL() string {
	if m != nil {
		return m.LogoURL
	}
	return ""
}

func (m *Friend) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Friend) GetImageURL() string {
	if m != nil {
		return m.ImageURL
	}
	return ""
}

func init() {
	proto.RegisterType((*Crew)(nil), "calcbiz.crew.Crew")
	proto.RegisterType((*Person)(nil), "calcbiz.crew.Person")
	proto.RegisterType((*WebAccount)(nil), "calcbiz.crew.WebAccount")
	proto.RegisterType((*Link)(nil), "calcbiz.crew.Link")
	proto.RegisterType((*Friend)(nil), "calcbiz.crew.Friend")
}
func (m *Crew) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Crew) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Website) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Website)))
		i += copy(dAtA[i:], m.Website)
	}
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCrew(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Accounts) > 0 {
		for _, msg := range m.Accounts {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCrew(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Friends) > 0 {
		for _, msg := range m.Friends {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintCrew(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Person) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Person) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *WebAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebAccount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Provider) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Provider)))
		i += copy(dAtA[i:], m.Provider)
	}
	if len(m.Handle) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Handle)))
		i += copy(dAtA[i:], m.Handle)
	}
	if len(m.URL) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.URL)))
		i += copy(dAtA[i:], m.URL)
	}
	return i, nil
}

func (m *Link) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Link) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.URL) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.URL)))
		i += copy(dAtA[i:], m.URL)
	}
	return i, nil
}

func (m *Friend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Friend) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.LogoURL) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.LogoURL)))
		i += copy(dAtA[i:], m.LogoURL)
	}
	if len(m.Links) > 0 {
		for _, msg := range m.Links {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintCrew(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ImageURL) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCrew(dAtA, i, uint64(len(m.ImageURL)))
		i += copy(dAtA[i:], m.ImageURL)
	}
	return i, nil
}

func encodeVarintCrew(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Crew) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovCrew(uint64(l))
		}
	}
	if len(m.Accounts) > 0 {
		for _, e := range m.Accounts {
			l = e.Size()
			n += 1 + l + sovCrew(uint64(l))
		}
	}
	if len(m.Friends) > 0 {
		for _, e := range m.Friends {
			l = e.Size()
			n += 1 + l + sovCrew(uint64(l))
		}
	}
	return n
}

func (m *Person) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	return n
}

func (m *WebAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.Handle)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	return n
}

func (m *Link) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	return n
}

func (m *Friend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	l = len(m.LogoURL)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovCrew(uint64(l))
		}
	}
	l = len(m.ImageURL)
	if l > 0 {
		n += 1 + l + sovCrew(uint64(l))
	}
	return n
}

func sovCrew(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCrew(x uint64) (n int) {
	return sovCrew(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Crew) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrew
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
			return fmt.Errorf("proto: Crew: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Crew: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Person{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, &WebAccount{})
			if err := m.Accounts[len(m.Accounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Friends", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Friends = append(m.Friends, &Friend{})
			if err := m.Friends[len(m.Friends)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrew(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrew
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
func (m *Person) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrew
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
			return fmt.Errorf("proto: Person: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Person: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrew(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrew
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
func (m *WebAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrew
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
			return fmt.Errorf("proto: WebAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrew(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrew
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
func (m *Link) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrew
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
			return fmt.Errorf("proto: Link: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Link: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrew(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrew
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
func (m *Friend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrew
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
			return fmt.Errorf("proto: Friend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Friend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogoURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogoURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrew
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
				return ErrInvalidLengthCrew
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrew(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrew
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
func skipCrew(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrew
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
					return 0, ErrIntOverflowCrew
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
					return 0, ErrIntOverflowCrew
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
				return 0, ErrInvalidLengthCrew
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCrew
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
				next, err := skipCrew(dAtA[start:])
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
	ErrInvalidLengthCrew = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrew   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pkg/crew/crew.proto", fileDescriptor_crew_26e212dfbbdddd43) }

var fileDescriptor_crew_26e212dfbbdddd43 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x6e, 0x9a, 0x6e, 0x92, 0x7d, 0xdd, 0x83, 0x8c, 0x22, 0xd9, 0x22, 0x69, 0xe9, 0x41, 0xea,
	0xc1, 0x04, 0x5c, 0xfd, 0x01, 0xae, 0x20, 0x08, 0x3d, 0x48, 0xa0, 0x08, 0x5e, 0x24, 0x49, 0xdf,
	0xce, 0x0e, 0x9d, 0xcc, 0x84, 0x49, 0x62, 0xd1, 0x5f, 0xe1, 0xcf, 0x12, 0x41, 0xd8, 0xa3, 0xa7,
	0x22, 0xe9, 0x1f, 0x91, 0x99, 0x24, 0xed, 0x2e, 0xf4, 0xe0, 0x25, 0xcc, 0xfb, 0xde, 0xf7, 0xcd,
	0xf7, 0xbe, 0xc7, 0x04, 0x1e, 0x17, 0x1b, 0x1a, 0x65, 0x0a, 0xb7, 0xe6, 0x13, 0x16, 0x4a, 0x56,
	0x92, 0x5c, 0x64, 0x09, 0xcf, 0x52, 0xf6, 0x3d, 0xd4, 0xd8, 0xe4, 0x19, 0x95, 0x92, 0x72, 0x8c,
	0x92, 0x82, 0x45, 0x89, 0x10, 0xb2, 0x4a, 0x2a, 0x26, 0x45, 0xd9, 0x72, 0x27, 0x2f, 0x29, 0xab,
	0x6e, 0xeb, 0x34, 0xcc, 0x64, 0x1e, 0x51, 0x49, 0x65, 0x64, 0xe0, 0xb4, 0xbe, 0x31, 0x95, 0x29,
	0xcc, 0xa9, 0xa5, 0xcf, 0x7f, 0x59, 0x30, 0x7a, 0xa7, 0x70, 0x4b, 0x08, 0x8c, 0x44, 0x92, 0xa3,
	0x6f, 0xcd, 0xac, 0xc5, 0x79, 0x6c, 0xce, 0xc4, 0x07, 0x77, 0x8b, 0x69, 0xc9, 0x2a, 0xf4, 0x87,
	0x06, 0xee, 0x4b, 0x12, 0x82, 0x9b, 0x63, 0x9e, 0xa2, 0x2a, 0x7d, 0x7b, 0x66, 0x2f, 0xc6, 0xaf,
	0x9e, 0x84, 0xf7, 0x67, 0x0c, 0x3f, 0xa2, 0x2a, 0xa5, 0x88, 0x7b, 0x12, 0x79, 0x0d, 0x5e, 0x92,
	0x65, 0xb2, 0x16, 0x55, 0xe9, 0x8f, 0x8c, 0xc0, 0x7f, 0x28, 0xf8, 0x84, 0xe9, 0xdb, 0x96, 0x10,
	0x1f, 0x98, 0xda, 0xe5, 0x46, 0x31, 0x14, 0xeb, 0xd2, 0x3f, 0x3b, 0xe5, 0xf2, 0xde, 0x34, 0xe3,
	0x9e, 0x34, 0x0f, 0xc1, 0x69, 0x8d, 0xc9, 0x23, 0xb0, 0x37, 0xf8, 0xad, 0x0b, 0xa3, 0x8f, 0x87,
	0x7c, 0xc3, 0x63, 0xbe, 0x79, 0x0e, 0x70, 0xf4, 0x3d, 0xa1, 0x99, 0x80, 0x57, 0x28, 0xf9, 0x95,
	0xad, 0x51, 0x75, 0xba, 0x43, 0x4d, 0x9e, 0x82, 0x73, 0x9b, 0x88, 0x35, 0x47, 0xdf, 0x36, 0x9d,
	0xae, 0x22, 0x97, 0x60, 0xd7, 0x8a, 0xfb, 0x23, 0x0d, 0x5e, 0xbb, 0xcd, 0x6e, 0x6a, 0xaf, 0xe2,
	0x65, 0xac, 0xb1, 0xf9, 0x1b, 0x18, 0x2d, 0x99, 0xd8, 0x9c, 0x5c, 0x75, 0x27, 0x1b, 0x9e, 0x90,
	0xfd, 0xb6, 0xc0, 0x69, 0x93, 0xfe, 0x5f, 0x2c, 0x32, 0x83, 0xf1, 0x1a, 0xcb, 0x4c, 0xb1, 0x42,
	0x3f, 0x8c, 0x6e, 0xbe, 0xfb, 0x10, 0x79, 0x0e, 0x1e, 0x97, 0x54, 0x7e, 0x39, 0x4e, 0x3a, 0x6e,
	0x76, 0x53, 0x77, 0x29, 0xa9, 0xd4, 0xb6, 0xae, 0x6e, 0xae, 0x14, 0x27, 0x0b, 0x38, 0xe3, 0x4c,
	0x6c, 0xfa, 0xf5, 0x93, 0x87, 0xeb, 0xd7, 0x61, 0xe2, 0x96, 0x40, 0x5e, 0xc0, 0x39, 0xcb, 0x13,
	0x8a, 0xe6, 0x4a, 0xc7, 0x5c, 0x79, 0xd1, 0xec, 0xa6, 0xde, 0x07, 0x0d, 0xea, 0x3b, 0x3d, 0xd3,
	0x5e, 0x29, 0x7e, 0x7d, 0xf5, 0xb3, 0x09, 0xac, 0xbb, 0x26, 0xb0, 0xfe, 0x36, 0x81, 0xf5, 0x63,
	0x1f, 0x0c, 0xee, 0xf6, 0xc1, 0xe0, 0xcf, 0x3e, 0x18, 0x7c, 0xbe, 0xac, 0x79, 0xa5, 0x30, 0xcc,
	0x31, 0xea, 0x7c, 0xa2, 0xfe, 0x6f, 0x48, 0x1d, 0xf3, 0x5c, 0xaf, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x87, 0xcc, 0x09, 0x0f, 0x20, 0x03, 0x00, 0x00,
}
