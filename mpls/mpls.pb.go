// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls/mpls.proto

package gnoi_mpls

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/openconfig/gnoi/types"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClearLSPRequest_Mode int32

const (
	ClearLSPRequest_DEFAULT ClearLSPRequest_Mode = 0
	// Reoptimize the LSP using the current bandwidth.
	ClearLSPRequest_NONAGGRESSIVE ClearLSPRequest_Mode = 0
	// Reoptimize the LSP using the current bandwidth. Only use IGP metric in
	// calcuation.
	ClearLSPRequest_AGGRESSIVE ClearLSPRequest_Mode = 1
	// Reset and restart all LSPs that originated from this routing device.
	ClearLSPRequest_RESET ClearLSPRequest_Mode = 2
	// Apply the highest bandwidth collected on a tunnel without waiting for
	// the current application period to end. Only use IGP metric in
	// calcuation.
	ClearLSPRequest_AUTOBW_AGGRESSIVE ClearLSPRequest_Mode = 3
	// Apply the highest bandwidth collected on a tunnel without waiting for
	// the current application period to end.
	ClearLSPRequest_AUTOBW_NONAGGRESSIVE ClearLSPRequest_Mode = 4
)

var ClearLSPRequest_Mode_name = map[int32]string{
	0: "DEFAULT",
	// Duplicate value: 0: "NONAGGRESSIVE",
	1: "AGGRESSIVE",
	2: "RESET",
	3: "AUTOBW_AGGRESSIVE",
	4: "AUTOBW_NONAGGRESSIVE",
}

var ClearLSPRequest_Mode_value = map[string]int32{
	"DEFAULT":              0,
	"NONAGGRESSIVE":        0,
	"AGGRESSIVE":           1,
	"RESET":                2,
	"AUTOBW_AGGRESSIVE":    3,
	"AUTOBW_NONAGGRESSIVE": 4,
}

func (x ClearLSPRequest_Mode) String() string {
	return proto.EnumName(ClearLSPRequest_Mode_name, int32(x))
}

func (ClearLSPRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{0, 0}
}

type MPLSPingRequest_ReplyMode int32

const (
	MPLSPingRequest_IPV4         MPLSPingRequest_ReplyMode = 0
	MPLSPingRequest_ROUTER_ALERT MPLSPingRequest_ReplyMode = 1
)

var MPLSPingRequest_ReplyMode_name = map[int32]string{
	0: "IPV4",
	1: "ROUTER_ALERT",
}

var MPLSPingRequest_ReplyMode_value = map[string]int32{
	"IPV4":         0,
	"ROUTER_ALERT": 1,
}

func (x MPLSPingRequest_ReplyMode) String() string {
	return proto.EnumName(MPLSPingRequest_ReplyMode_name, int32(x))
}

func (MPLSPingRequest_ReplyMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{6, 0}
}

type MPLSPingResponse_EchoResponseCode int32

const (
	// A successful echo response was received.
	MPLSPingResponse_SUCCESS MPLSPingResponse_EchoResponseCode = 0
	// The MPLS ping packet was not sent, for an unknown reason.
	MPLSPingResponse_NOT_SENT MPLSPingResponse_EchoResponseCode = 1
	// The local system timed out waiting for an LSP ping response.
	MPLSPingResponse_TIMEOUT MPLSPingResponse_EchoResponseCode = 2
)

var MPLSPingResponse_EchoResponseCode_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_SENT",
	2: "TIMEOUT",
}

var MPLSPingResponse_EchoResponseCode_value = map[string]int32{
	"SUCCESS":  0,
	"NOT_SENT": 1,
	"TIMEOUT":  2,
}

func (x MPLSPingResponse_EchoResponseCode) String() string {
	return proto.EnumName(MPLSPingResponse_EchoResponseCode_name, int32(x))
}

func (MPLSPingResponse_EchoResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{7, 0}
}

// Request to clear a single tunnel on a target device.
type ClearLSPRequest struct {
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mode                 ClearLSPRequest_Mode `protobuf:"varint,3,opt,name=mode,proto3,enum=gnoi.mpls.ClearLSPRequest_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClearLSPRequest) Reset()         { *m = ClearLSPRequest{} }
func (m *ClearLSPRequest) String() string { return proto.CompactTextString(m) }
func (*ClearLSPRequest) ProtoMessage()    {}
func (*ClearLSPRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{0}
}

func (m *ClearLSPRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearLSPRequest.Unmarshal(m, b)
}
func (m *ClearLSPRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearLSPRequest.Marshal(b, m, deterministic)
}
func (m *ClearLSPRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearLSPRequest.Merge(m, src)
}
func (m *ClearLSPRequest) XXX_Size() int {
	return xxx_messageInfo_ClearLSPRequest.Size(m)
}
func (m *ClearLSPRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearLSPRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClearLSPRequest proto.InternalMessageInfo

func (m *ClearLSPRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClearLSPRequest) GetMode() ClearLSPRequest_Mode {
	if m != nil {
		return m.Mode
	}
	return ClearLSPRequest_DEFAULT
}

type ClearLSPResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearLSPResponse) Reset()         { *m = ClearLSPResponse{} }
func (m *ClearLSPResponse) String() string { return proto.CompactTextString(m) }
func (*ClearLSPResponse) ProtoMessage()    {}
func (*ClearLSPResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{1}
}

func (m *ClearLSPResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearLSPResponse.Unmarshal(m, b)
}
func (m *ClearLSPResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearLSPResponse.Marshal(b, m, deterministic)
}
func (m *ClearLSPResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearLSPResponse.Merge(m, src)
}
func (m *ClearLSPResponse) XXX_Size() int {
	return xxx_messageInfo_ClearLSPResponse.Size(m)
}
func (m *ClearLSPResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearLSPResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClearLSPResponse proto.InternalMessageInfo

// Request to clear a single tunnel counters on a target device.
type ClearLSPCountersRequest struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearLSPCountersRequest) Reset()         { *m = ClearLSPCountersRequest{} }
func (m *ClearLSPCountersRequest) String() string { return proto.CompactTextString(m) }
func (*ClearLSPCountersRequest) ProtoMessage()    {}
func (*ClearLSPCountersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{2}
}

func (m *ClearLSPCountersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearLSPCountersRequest.Unmarshal(m, b)
}
func (m *ClearLSPCountersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearLSPCountersRequest.Marshal(b, m, deterministic)
}
func (m *ClearLSPCountersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearLSPCountersRequest.Merge(m, src)
}
func (m *ClearLSPCountersRequest) XXX_Size() int {
	return xxx_messageInfo_ClearLSPCountersRequest.Size(m)
}
func (m *ClearLSPCountersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearLSPCountersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClearLSPCountersRequest proto.InternalMessageInfo

func (m *ClearLSPCountersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClearLSPCountersResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearLSPCountersResponse) Reset()         { *m = ClearLSPCountersResponse{} }
func (m *ClearLSPCountersResponse) String() string { return proto.CompactTextString(m) }
func (*ClearLSPCountersResponse) ProtoMessage()    {}
func (*ClearLSPCountersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{3}
}

func (m *ClearLSPCountersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearLSPCountersResponse.Unmarshal(m, b)
}
func (m *ClearLSPCountersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearLSPCountersResponse.Marshal(b, m, deterministic)
}
func (m *ClearLSPCountersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearLSPCountersResponse.Merge(m, src)
}
func (m *ClearLSPCountersResponse) XXX_Size() int {
	return xxx_messageInfo_ClearLSPCountersResponse.Size(m)
}
func (m *ClearLSPCountersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearLSPCountersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClearLSPCountersResponse proto.InternalMessageInfo

type MPLSPingPWEDestination struct {
	// The address of the egress LER that the MPLS ping should be sent on when
	// destined to a PWE service.
	Eler string `protobuf:"bytes,1,opt,name=eler,proto3" json:"eler,omitempty"`
	// The virtual circuit ID for the PWE via which the ping should be sent.
	Vcid                 uint32   `protobuf:"varint,2,opt,name=vcid,proto3" json:"vcid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MPLSPingPWEDestination) Reset()         { *m = MPLSPingPWEDestination{} }
func (m *MPLSPingPWEDestination) String() string { return proto.CompactTextString(m) }
func (*MPLSPingPWEDestination) ProtoMessage()    {}
func (*MPLSPingPWEDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{4}
}

func (m *MPLSPingPWEDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MPLSPingPWEDestination.Unmarshal(m, b)
}
func (m *MPLSPingPWEDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MPLSPingPWEDestination.Marshal(b, m, deterministic)
}
func (m *MPLSPingPWEDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MPLSPingPWEDestination.Merge(m, src)
}
func (m *MPLSPingPWEDestination) XXX_Size() int {
	return xxx_messageInfo_MPLSPingPWEDestination.Size(m)
}
func (m *MPLSPingPWEDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_MPLSPingPWEDestination.DiscardUnknown(m)
}

var xxx_messageInfo_MPLSPingPWEDestination proto.InternalMessageInfo

func (m *MPLSPingPWEDestination) GetEler() string {
	if m != nil {
		return m.Eler
	}
	return ""
}

func (m *MPLSPingPWEDestination) GetVcid() uint32 {
	if m != nil {
		return m.Vcid
	}
	return 0
}

// MPLSPingRSVPTEDestination specifies the destination for an MPLS Ping in
// terms of an absolute specification of an RSVP-TE LSP. It can be used to
// identify an individual RSVP-TE session via which a ping should be sent.
type MPLSPingRSVPTEDestination struct {
	// The IPv4 or IPv6 address used by the system initiating (acting as the
	// head-end) of the RSVP-TE LSP.
	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	// The IPv4 or IPv6 address used by the system terminating (acting as the
	// tail-end) of the RSVP-TE LSP.
	Dst string `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	// The extended tunnel ID of the RSVP-TE LSP, expressed as an unsigned, 32b
	// integer.
	ExtendedTunnelId     uint32   `protobuf:"varint,3,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MPLSPingRSVPTEDestination) Reset()         { *m = MPLSPingRSVPTEDestination{} }
func (m *MPLSPingRSVPTEDestination) String() string { return proto.CompactTextString(m) }
func (*MPLSPingRSVPTEDestination) ProtoMessage()    {}
func (*MPLSPingRSVPTEDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{5}
}

func (m *MPLSPingRSVPTEDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MPLSPingRSVPTEDestination.Unmarshal(m, b)
}
func (m *MPLSPingRSVPTEDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MPLSPingRSVPTEDestination.Marshal(b, m, deterministic)
}
func (m *MPLSPingRSVPTEDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MPLSPingRSVPTEDestination.Merge(m, src)
}
func (m *MPLSPingRSVPTEDestination) XXX_Size() int {
	return xxx_messageInfo_MPLSPingRSVPTEDestination.Size(m)
}
func (m *MPLSPingRSVPTEDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_MPLSPingRSVPTEDestination.DiscardUnknown(m)
}

var xxx_messageInfo_MPLSPingRSVPTEDestination proto.InternalMessageInfo

func (m *MPLSPingRSVPTEDestination) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *MPLSPingRSVPTEDestination) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *MPLSPingRSVPTEDestination) GetExtendedTunnelId() uint32 {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return 0
}

// MPLSPingRequest specifies the parameters that should be used as input from
// the client, to a system that is initiating an RFC4379 MPLS ping request.
type MPLSPingRequest struct {
	// One field within the destination field should be specified to determine
	// the destination of the LSP ping.
	//
	// Types that are valid to be assigned to Destination:
	//	*MPLSPingRequest_LdpFec
	//	*MPLSPingRequest_Fec129Pwe
	//	*MPLSPingRequest_RsvpteLspName
	//	*MPLSPingRequest_RsvpteLsp
	Destination isMPLSPingRequest_Destination `protobuf_oneof:"destination"`
	ReplyMode   MPLSPingRequest_ReplyMode     `protobuf:"varint,6,opt,name=reply_mode,json=replyMode,proto3,enum=gnoi.mpls.MPLSPingRequest_ReplyMode" json:"reply_mode,omitempty"`
	// The number of MPLS echo request packets to send.
	Count uint32 `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	// The size (in bytes) of each MPLS echo request packet.
	Size uint32 `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
	// The source IPv4 address that should be used in the request packet.
	SourceAddress string `protobuf:"bytes,9,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// The MPLS TTL that should be set in the packets sent.
	MplsTtl uint32 `protobuf:"varint,10,opt,name=mpls_ttl,json=mplsTtl,proto3" json:"mpls_ttl,omitempty"`
	// The value of the traffic class (TC, formerly known as EXP) bits that
	// should be set in the MPLS ping packets.
	TrafficClass         uint32   `protobuf:"varint,11,opt,name=traffic_class,json=trafficClass,proto3" json:"traffic_class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MPLSPingRequest) Reset()         { *m = MPLSPingRequest{} }
func (m *MPLSPingRequest) String() string { return proto.CompactTextString(m) }
func (*MPLSPingRequest) ProtoMessage()    {}
func (*MPLSPingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{6}
}

func (m *MPLSPingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MPLSPingRequest.Unmarshal(m, b)
}
func (m *MPLSPingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MPLSPingRequest.Marshal(b, m, deterministic)
}
func (m *MPLSPingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MPLSPingRequest.Merge(m, src)
}
func (m *MPLSPingRequest) XXX_Size() int {
	return xxx_messageInfo_MPLSPingRequest.Size(m)
}
func (m *MPLSPingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MPLSPingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MPLSPingRequest proto.InternalMessageInfo

type isMPLSPingRequest_Destination interface {
	isMPLSPingRequest_Destination()
}

type MPLSPingRequest_LdpFec struct {
	LdpFec string `protobuf:"bytes,1,opt,name=ldp_fec,json=ldpFec,proto3,oneof"`
}

type MPLSPingRequest_Fec129Pwe struct {
	Fec129Pwe *MPLSPingPWEDestination `protobuf:"bytes,2,opt,name=fec129_pwe,json=fec129Pwe,proto3,oneof"`
}

type MPLSPingRequest_RsvpteLspName struct {
	RsvpteLspName string `protobuf:"bytes,4,opt,name=rsvpte_lsp_name,json=rsvpteLspName,proto3,oneof"`
}

type MPLSPingRequest_RsvpteLsp struct {
	RsvpteLsp *MPLSPingRSVPTEDestination `protobuf:"bytes,5,opt,name=rsvpte_lsp,json=rsvpteLsp,proto3,oneof"`
}

func (*MPLSPingRequest_LdpFec) isMPLSPingRequest_Destination() {}

func (*MPLSPingRequest_Fec129Pwe) isMPLSPingRequest_Destination() {}

func (*MPLSPingRequest_RsvpteLspName) isMPLSPingRequest_Destination() {}

func (*MPLSPingRequest_RsvpteLsp) isMPLSPingRequest_Destination() {}

func (m *MPLSPingRequest) GetDestination() isMPLSPingRequest_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *MPLSPingRequest) GetLdpFec() string {
	if x, ok := m.GetDestination().(*MPLSPingRequest_LdpFec); ok {
		return x.LdpFec
	}
	return ""
}

func (m *MPLSPingRequest) GetFec129Pwe() *MPLSPingPWEDestination {
	if x, ok := m.GetDestination().(*MPLSPingRequest_Fec129Pwe); ok {
		return x.Fec129Pwe
	}
	return nil
}

func (m *MPLSPingRequest) GetRsvpteLspName() string {
	if x, ok := m.GetDestination().(*MPLSPingRequest_RsvpteLspName); ok {
		return x.RsvpteLspName
	}
	return ""
}

func (m *MPLSPingRequest) GetRsvpteLsp() *MPLSPingRSVPTEDestination {
	if x, ok := m.GetDestination().(*MPLSPingRequest_RsvpteLsp); ok {
		return x.RsvpteLsp
	}
	return nil
}

func (m *MPLSPingRequest) GetReplyMode() MPLSPingRequest_ReplyMode {
	if m != nil {
		return m.ReplyMode
	}
	return MPLSPingRequest_IPV4
}

func (m *MPLSPingRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MPLSPingRequest) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MPLSPingRequest) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MPLSPingRequest) GetMplsTtl() uint32 {
	if m != nil {
		return m.MplsTtl
	}
	return 0
}

func (m *MPLSPingRequest) GetTrafficClass() uint32 {
	if m != nil {
		return m.TrafficClass
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MPLSPingRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MPLSPingRequest_OneofMarshaler, _MPLSPingRequest_OneofUnmarshaler, _MPLSPingRequest_OneofSizer, []interface{}{
		(*MPLSPingRequest_LdpFec)(nil),
		(*MPLSPingRequest_Fec129Pwe)(nil),
		(*MPLSPingRequest_RsvpteLspName)(nil),
		(*MPLSPingRequest_RsvpteLsp)(nil),
	}
}

func _MPLSPingRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MPLSPingRequest)
	// destination
	switch x := m.Destination.(type) {
	case *MPLSPingRequest_LdpFec:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.LdpFec)
	case *MPLSPingRequest_Fec129Pwe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Fec129Pwe); err != nil {
			return err
		}
	case *MPLSPingRequest_RsvpteLspName:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.RsvpteLspName)
	case *MPLSPingRequest_RsvpteLsp:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RsvpteLsp); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MPLSPingRequest.Destination has unexpected type %T", x)
	}
	return nil
}

func _MPLSPingRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MPLSPingRequest)
	switch tag {
	case 1: // destination.ldp_fec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Destination = &MPLSPingRequest_LdpFec{x}
		return true, err
	case 2: // destination.fec129_pwe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MPLSPingPWEDestination)
		err := b.DecodeMessage(msg)
		m.Destination = &MPLSPingRequest_Fec129Pwe{msg}
		return true, err
	case 4: // destination.rsvpte_lsp_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Destination = &MPLSPingRequest_RsvpteLspName{x}
		return true, err
	case 5: // destination.rsvpte_lsp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MPLSPingRSVPTEDestination)
		err := b.DecodeMessage(msg)
		m.Destination = &MPLSPingRequest_RsvpteLsp{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MPLSPingRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MPLSPingRequest)
	// destination
	switch x := m.Destination.(type) {
	case *MPLSPingRequest_LdpFec:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.LdpFec)))
		n += len(x.LdpFec)
	case *MPLSPingRequest_Fec129Pwe:
		s := proto.Size(x.Fec129Pwe)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MPLSPingRequest_RsvpteLspName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.RsvpteLspName)))
		n += len(x.RsvpteLspName)
	case *MPLSPingRequest_RsvpteLsp:
		s := proto.Size(x.RsvpteLsp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// MPLSPingResponse (MPLSPong?) is sent from the target to the client based on
// each MPLS Echo Response packet it receives associated with an input MPLSPing
// RPC.
type MPLSPingResponse struct {
	Seq uint32 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	// The response that was received from the egress LER.
	Response MPLSPingResponse_EchoResponseCode `protobuf:"varint,2,opt,name=response,proto3,enum=gnoi.mpls.MPLSPingResponse_EchoResponseCode" json:"response,omitempty"`
	// The time (in nanoseconds) between transmission of the Echo Response,
	// and the echo reply.
	ResponseTime         uint64   `protobuf:"varint,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MPLSPingResponse) Reset()         { *m = MPLSPingResponse{} }
func (m *MPLSPingResponse) String() string { return proto.CompactTextString(m) }
func (*MPLSPingResponse) ProtoMessage()    {}
func (*MPLSPingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4448d3521f3d4b5, []int{7}
}

func (m *MPLSPingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MPLSPingResponse.Unmarshal(m, b)
}
func (m *MPLSPingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MPLSPingResponse.Marshal(b, m, deterministic)
}
func (m *MPLSPingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MPLSPingResponse.Merge(m, src)
}
func (m *MPLSPingResponse) XXX_Size() int {
	return xxx_messageInfo_MPLSPingResponse.Size(m)
}
func (m *MPLSPingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MPLSPingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MPLSPingResponse proto.InternalMessageInfo

func (m *MPLSPingResponse) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *MPLSPingResponse) GetResponse() MPLSPingResponse_EchoResponseCode {
	if m != nil {
		return m.Response
	}
	return MPLSPingResponse_SUCCESS
}

func (m *MPLSPingResponse) GetResponseTime() uint64 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func init() {
	proto.RegisterEnum("gnoi.mpls.ClearLSPRequest_Mode", ClearLSPRequest_Mode_name, ClearLSPRequest_Mode_value)
	proto.RegisterEnum("gnoi.mpls.MPLSPingRequest_ReplyMode", MPLSPingRequest_ReplyMode_name, MPLSPingRequest_ReplyMode_value)
	proto.RegisterEnum("gnoi.mpls.MPLSPingResponse_EchoResponseCode", MPLSPingResponse_EchoResponseCode_name, MPLSPingResponse_EchoResponseCode_value)
	proto.RegisterType((*ClearLSPRequest)(nil), "gnoi.mpls.ClearLSPRequest")
	proto.RegisterType((*ClearLSPResponse)(nil), "gnoi.mpls.ClearLSPResponse")
	proto.RegisterType((*ClearLSPCountersRequest)(nil), "gnoi.mpls.ClearLSPCountersRequest")
	proto.RegisterType((*ClearLSPCountersResponse)(nil), "gnoi.mpls.ClearLSPCountersResponse")
	proto.RegisterType((*MPLSPingPWEDestination)(nil), "gnoi.mpls.MPLSPingPWEDestination")
	proto.RegisterType((*MPLSPingRSVPTEDestination)(nil), "gnoi.mpls.MPLSPingRSVPTEDestination")
	proto.RegisterType((*MPLSPingRequest)(nil), "gnoi.mpls.MPLSPingRequest")
	proto.RegisterType((*MPLSPingResponse)(nil), "gnoi.mpls.MPLSPingResponse")
}

func init() { proto.RegisterFile("mpls/mpls.proto", fileDescriptor_d4448d3521f3d4b5) }

var fileDescriptor_d4448d3521f3d4b5 = []byte{
	// 778 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4f, 0x8f, 0xda, 0x46,
	0x14, 0xb7, 0xc1, 0x2c, 0xe6, 0x81, 0xc1, 0x19, 0xa5, 0xad, 0x97, 0x1e, 0x9a, 0x3a, 0xad, 0xba,
	0x87, 0xad, 0xd9, 0x90, 0x5e, 0xda, 0x43, 0x55, 0x20, 0xce, 0x2e, 0x12, 0x0b, 0x68, 0x6c, 0x36,
	0xa7, 0xca, 0x22, 0xf6, 0x40, 0x2c, 0x19, 0xdb, 0xf1, 0x0c, 0x9b, 0xa6, 0x5f, 0xb2, 0x52, 0xbe,
	0x48, 0x4f, 0xbd, 0x57, 0x33, 0xb6, 0x61, 0x21, 0x64, 0x2f, 0xd6, 0x9b, 0xdf, 0x7b, 0xef, 0xf7,
	0xfe, 0xfa, 0x41, 0x67, 0x93, 0x46, 0xb4, 0xc7, 0x3f, 0x56, 0x9a, 0x25, 0x2c, 0x41, 0x8d, 0x75,
	0x9c, 0x84, 0x16, 0x07, 0xba, 0x97, 0xeb, 0x90, 0xbd, 0xdb, 0xbe, 0xb5, 0xfc, 0x64, 0xd3, 0x4b,
	0x52, 0x12, 0xfb, 0x49, 0xbc, 0x0a, 0xd7, 0x3d, 0x6e, 0xd0, 0x63, 0x1f, 0x53, 0x42, 0xf3, 0x6f,
	0xee, 0x68, 0x7e, 0x92, 0xa1, 0x33, 0x8a, 0xc8, 0x32, 0x9b, 0x38, 0x73, 0x4c, 0xde, 0x6f, 0x09,
	0x65, 0x08, 0x81, 0x12, 0x2f, 0x37, 0xc4, 0xa8, 0x3c, 0x93, 0x2f, 0x1a, 0x58, 0xc8, 0xe8, 0x25,
	0x28, 0x9b, 0x24, 0x20, 0x46, 0xf5, 0x99, 0x7c, 0xd1, 0xee, 0x7f, 0x67, 0xed, 0xe2, 0x59, 0x47,
	0xde, 0xd6, 0x6d, 0x12, 0x10, 0x2c, 0x8c, 0xcd, 0x7b, 0x50, 0xf8, 0x0b, 0x35, 0xa1, 0xfe, 0xca,
	0x7e, 0x3d, 0x58, 0x4c, 0x5c, 0x5d, 0x42, 0x4f, 0x40, 0x9b, 0xce, 0xa6, 0x83, 0xeb, 0x6b, 0x6c,
	0x3b, 0xce, 0xf8, 0xce, 0xd6, 0x25, 0xd4, 0x06, 0x78, 0xf0, 0x96, 0x51, 0x03, 0x6a, 0xd8, 0x76,
	0x6c, 0x57, 0xaf, 0xa0, 0xaf, 0xe0, 0xc9, 0x60, 0xe1, 0xce, 0x86, 0x6f, 0xbc, 0x07, 0x16, 0x55,
	0x64, 0xc0, 0xd3, 0x02, 0x3e, 0xe4, 0x52, 0xba, 0x15, 0x5d, 0x36, 0x11, 0xe8, 0xfb, 0xac, 0x68,
	0x9a, 0xc4, 0x94, 0x98, 0x3f, 0xc3, 0x37, 0x25, 0x36, 0x4a, 0xb6, 0x31, 0x23, 0x19, 0x7d, 0xa4,
	0x5e, 0xb3, 0x0b, 0xc6, 0xe7, 0xe6, 0x05, 0xd5, 0x1f, 0xf0, 0xf5, 0xed, 0x7c, 0xe2, 0xcc, 0xc3,
	0x78, 0x3d, 0x7f, 0x63, 0xbf, 0x22, 0x94, 0x85, 0xf1, 0x92, 0x85, 0x49, 0xcc, 0x99, 0x48, 0x44,
	0x32, 0x43, 0xce, 0x99, 0xb8, 0xcc, 0xb1, 0x7b, 0x3f, 0x0c, 0x04, 0xbb, 0x86, 0x85, 0x6c, 0x6e,
	0xe0, 0xbc, 0x64, 0xc0, 0xce, 0xdd, 0xdc, 0x3d, 0x20, 0xd1, 0xa1, 0x4a, 0x33, 0xbf, 0xe0, 0xe0,
	0x22, 0x47, 0x02, 0xca, 0x8a, 0xfc, 0xb8, 0x88, 0x2e, 0x01, 0x91, 0xbf, 0x18, 0x89, 0x03, 0x12,
	0x78, 0x6c, 0x1b, 0xc7, 0x24, 0xf2, 0xc2, 0x40, 0x0c, 0x47, 0xc3, 0x7a, 0xa9, 0x71, 0x85, 0x62,
	0x1c, 0x98, 0xff, 0x56, 0xa1, 0xb3, 0x8b, 0x57, 0x14, 0x7d, 0x0e, 0xf5, 0x28, 0x48, 0xbd, 0x15,
	0x29, 0x22, 0xdd, 0x48, 0xf8, 0x2c, 0x0a, 0xd2, 0xd7, 0xc4, 0x47, 0x43, 0x80, 0x15, 0xf1, 0x5f,
	0xf4, 0x7f, 0xf5, 0xd2, 0x0f, 0x79, 0x57, 0x9a, 0xfd, 0xef, 0x1f, 0x4c, 0xfc, 0x74, 0xf1, 0x37,
	0x12, 0x6e, 0xe4, 0x6e, 0xf3, 0x0f, 0x04, 0x5d, 0x40, 0x27, 0xa3, 0xf7, 0x29, 0x23, 0x5e, 0x44,
	0x53, 0x4f, 0xb4, 0x57, 0x29, 0xc2, 0x68, 0xb9, 0x62, 0x42, 0xd3, 0x29, 0xdf, 0x2c, 0x1b, 0x60,
	0x6f, 0x69, 0xd4, 0x44, 0xb4, 0x1f, 0x4e, 0x44, 0xfb, 0xac, 0x51, 0x3c, 0xe0, 0x8e, 0x0a, 0x8d,
	0x00, 0x32, 0x92, 0x46, 0x1f, 0x3d, 0xb1, 0xa6, 0x67, 0x62, 0x4d, 0x4f, 0xd2, 0x14, 0x6b, 0x8a,
	0xb9, 0xb1, 0xd8, 0xd5, 0x46, 0x56, 0x8a, 0xe8, 0x29, 0xd4, 0x7c, 0x3e, 0x6d, 0xa3, 0x2e, 0x3a,
	0x99, 0x3f, 0xf8, 0x04, 0x69, 0xf8, 0x37, 0x31, 0xd4, 0x7c, 0x82, 0x5c, 0x46, 0x3f, 0x42, 0x9b,
	0x26, 0xdb, 0xcc, 0x27, 0xde, 0x32, 0x08, 0x32, 0x42, 0xa9, 0xd1, 0x10, 0xd3, 0xd1, 0x72, 0x74,
	0x90, 0x83, 0xe8, 0x1c, 0x54, 0x1e, 0xdd, 0x63, 0x2c, 0x32, 0x40, 0xb8, 0xd7, 0xf9, 0xdb, 0x65,
	0x11, 0x7a, 0x0e, 0x1a, 0xcb, 0x96, 0xab, 0x55, 0xe8, 0x7b, 0x7e, 0xb4, 0xa4, 0xd4, 0x68, 0x0a,
	0x7d, 0xab, 0x00, 0x47, 0x1c, 0x33, 0x7f, 0x82, 0xc6, 0x2e, 0x51, 0xa4, 0x82, 0x32, 0x9e, 0xdf,
	0xfd, 0xa2, 0x4b, 0x48, 0x87, 0x16, 0x9e, 0x2d, 0x5c, 0x1b, 0x7b, 0x83, 0x89, 0x8d, 0x5d, 0x5d,
	0x1e, 0x6a, 0xd0, 0x0c, 0xf6, 0xad, 0x31, 0xff, 0x91, 0x41, 0xdf, 0x57, 0x9c, 0xef, 0xad, 0x58,
	0x2c, 0xf2, 0x5e, 0x8c, 0x5b, 0xc3, 0x5c, 0x44, 0x37, 0xa0, 0x66, 0x85, 0x56, 0xcc, 0xb9, 0xdd,
	0xbf, 0x3c, 0xd9, 0xb2, 0xdc, 0xc4, 0xb2, 0xfd, 0x77, 0x49, 0xf9, 0x18, 0xf1, 0xd6, 0xed, 0xbc,
	0x79, 0x35, 0xa5, 0xec, 0xb1, 0x70, 0x93, 0x1f, 0x0a, 0x05, 0xb7, 0x4a, 0xd0, 0x0d, 0x37, 0xc4,
	0xfc, 0x0d, 0xf4, 0x63, 0x0a, 0x7e, 0x1b, 0x9c, 0xc5, 0x68, 0x64, 0x3b, 0x8e, 0x2e, 0xa1, 0x16,
	0xa8, 0xd3, 0x99, 0xeb, 0x39, 0xf6, 0xd4, 0xd5, 0x65, 0xae, 0x72, 0xc7, 0xb7, 0xf6, 0x6c, 0xe1,
	0xea, 0x95, 0xfe, 0x7f, 0x32, 0x28, 0x3c, 0x21, 0x64, 0x83, 0x5a, 0xfe, 0x99, 0xa8, 0xfb, 0xe5,
	0x3b, 0xd4, 0xfd, 0xf6, 0xa4, 0xae, 0xf8, 0x85, 0x25, 0xf4, 0xe7, 0xfe, 0x46, 0x94, 0x3f, 0x38,
	0x32, 0x4f, 0xb8, 0x1c, 0x1d, 0x8b, 0xee, 0xf3, 0x47, 0x6d, 0x76, 0xf4, 0xd7, 0xa0, 0x96, 0xed,
	0x3b, 0xc8, 0xf2, 0x68, 0x0d, 0x0f, 0xb2, 0x3c, 0xee, 0xb7, 0x29, 0x5d, 0xc9, 0x43, 0xf5, 0xd3,
	0xef, 0xb5, 0x2b, 0xeb, 0x85, 0x75, 0xf5, 0xf6, 0x4c, 0x5c, 0xec, 0x97, 0xff, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x2b, 0x14, 0x01, 0x1f, 0xfd, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MPLSClient is the client API for MPLS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MPLSClient interface {
	// ClearLSP clears a single tunnel (requests for it's route to be
	// recalculated).
	ClearLSP(ctx context.Context, in *ClearLSPRequest, opts ...grpc.CallOption) (*ClearLSPResponse, error)
	// ClearLSPCounters will clear the MPLS counters for the provided LSP.
	ClearLSPCounters(ctx context.Context, in *ClearLSPCountersRequest, opts ...grpc.CallOption) (*ClearLSPCountersResponse, error)
	// An MPLS ping, specified as per RFC4379.
	MPLSPing(ctx context.Context, in *MPLSPingRequest, opts ...grpc.CallOption) (MPLS_MPLSPingClient, error)
}

type mPLSClient struct {
	cc *grpc.ClientConn
}

func NewMPLSClient(cc *grpc.ClientConn) MPLSClient {
	return &mPLSClient{cc}
}

func (c *mPLSClient) ClearLSP(ctx context.Context, in *ClearLSPRequest, opts ...grpc.CallOption) (*ClearLSPResponse, error) {
	out := new(ClearLSPResponse)
	err := c.cc.Invoke(ctx, "/gnoi.mpls.MPLS/ClearLSP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPLSClient) ClearLSPCounters(ctx context.Context, in *ClearLSPCountersRequest, opts ...grpc.CallOption) (*ClearLSPCountersResponse, error) {
	out := new(ClearLSPCountersResponse)
	err := c.cc.Invoke(ctx, "/gnoi.mpls.MPLS/ClearLSPCounters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPLSClient) MPLSPing(ctx context.Context, in *MPLSPingRequest, opts ...grpc.CallOption) (MPLS_MPLSPingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MPLS_serviceDesc.Streams[0], "/gnoi.mpls.MPLS/MPLSPing", opts...)
	if err != nil {
		return nil, err
	}
	x := &mPLSMPLSPingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MPLS_MPLSPingClient interface {
	Recv() (*MPLSPingResponse, error)
	grpc.ClientStream
}

type mPLSMPLSPingClient struct {
	grpc.ClientStream
}

func (x *mPLSMPLSPingClient) Recv() (*MPLSPingResponse, error) {
	m := new(MPLSPingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MPLSServer is the server API for MPLS service.
type MPLSServer interface {
	// ClearLSP clears a single tunnel (requests for it's route to be
	// recalculated).
	ClearLSP(context.Context, *ClearLSPRequest) (*ClearLSPResponse, error)
	// ClearLSPCounters will clear the MPLS counters for the provided LSP.
	ClearLSPCounters(context.Context, *ClearLSPCountersRequest) (*ClearLSPCountersResponse, error)
	// An MPLS ping, specified as per RFC4379.
	MPLSPing(*MPLSPingRequest, MPLS_MPLSPingServer) error
}

func RegisterMPLSServer(s *grpc.Server, srv MPLSServer) {
	s.RegisterService(&_MPLS_serviceDesc, srv)
}

func _MPLS_ClearLSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearLSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPLSServer).ClearLSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.mpls.MPLS/ClearLSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPLSServer).ClearLSP(ctx, req.(*ClearLSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPLS_ClearLSPCounters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearLSPCountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPLSServer).ClearLSPCounters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.mpls.MPLS/ClearLSPCounters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPLSServer).ClearLSPCounters(ctx, req.(*ClearLSPCountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPLS_MPLSPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MPLSPingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MPLSServer).MPLSPing(m, &mPLSMPLSPingServer{stream})
}

type MPLS_MPLSPingServer interface {
	Send(*MPLSPingResponse) error
	grpc.ServerStream
}

type mPLSMPLSPingServer struct {
	grpc.ServerStream
}

func (x *mPLSMPLSPingServer) Send(m *MPLSPingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MPLS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.mpls.MPLS",
	HandlerType: (*MPLSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClearLSP",
			Handler:    _MPLS_ClearLSP_Handler,
		},
		{
			MethodName: "ClearLSPCounters",
			Handler:    _MPLS_ClearLSPCounters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MPLSPing",
			Handler:       _MPLS_MPLSPing_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mpls/mpls.proto",
}
