// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/lds.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import envoy_api_v2_core1 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import envoy_api_v2_listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
import _ "github.com/gogo/googleapis/google/api"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Listener_DrainType int32

const (
	// Drain in response to calling /healthcheck/fail admin endpoint (along with the health check
	// filter), listener removal/modification, and hot restart.
	Listener_DEFAULT Listener_DrainType = 0
	// Drain in response to listener removal/modification and hot restart. This setting does not
	// include /healthcheck/fail. This setting may be desirable if Envoy is hosting both ingress
	// and egress listeners.
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}
var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}
func (Listener_DrainType) EnumDescriptor() ([]byte, []int) { return fileDescriptorLds, []int{0, 0} }

type Listener struct {
	// The unique name by which this listener is known. If no name is provided,
	// Envoy will allocate an internal UUID for the listener. If the listener is to be dynamically
	// updated or removed via :ref:`LDS <config_listeners_lds>` a unique name must be provided.
	// By default, the maximum length of a listener's name is limited to 60 characters. This limit can
	// be increased by setting the :option:`--max-obj-name-len` command line argument to the desired
	// value.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The address that the listener should listen on. In general, the address must be unique, though
	// that is governed by the bind rules of the OS. E.g., multiple listeners can listen on port 0 on
	// Linux as the actual port will be allocated by the OS.
	Address envoy_api_v2_core.Address `protobuf:"bytes,2,opt,name=address" json:"address"`
	// A list of filter chains to consider for this listener. The
	// :ref:`FilterChain <envoy_api_msg_listener.FilterChain>` with the most specific
	// :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` criteria is used on a
	// connection.
	//
	// .. attention::
	//
	//   In the current version, multiple filter chains are supported **only** so that SNI can be
	//   configured. See the :ref:`FAQ entry <faq_how_to_setup_sni>` on how to configure SNI for more
	//   information. When multiple filter chains are configured, each filter chain must have an
	//   **identical** set of :ref:`filters <envoy_api_field_listener.FilterChain.filters>`. If the
	//   filters differ, the configuration will fail to load. In the future, this limitation will be
	//   relaxed such that different filters can be used depending on which filter chain matches
	//   (based on SNI or some other parameter).
	FilterChains []envoy_api_v2_listener.FilterChain `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains" json:"filter_chains"`
	// If a connection is redirected using *iptables*, the port on which the proxy
	// receives it might be different from the original destination address. When this flag is set to
	// true, the listener hands off redirected connections to the listener associated with the
	// original destination address. If there is no listener associated with the original destination
	// address, the connection is handled by the listener that receives it. Defaults to false.
	//
	// .. attention::
	//
	//   This field is deprecated. Use :ref:`an original_dst <config_listener_filters_original_dst>`
	//   :ref:`listener filter <envoy_api_field_Listener.listener_filters>` instead.
	//
	//   Note that hand off to another listener is *NOT* performed without this flag. Once
	//   :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` is implemented this flag
	//   will be removed, as filter chain matching can be used to select a filter chain based on the
	//   restored destination address.
	UseOriginalDst *google_protobuf.BoolValue `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst" json:"use_original_dst,omitempty"`
	// Soft limit on size of the listener’s new connection read and write buffers.
	// If unspecified, an implementation defined default is applied (1MiB).
	PerConnectionBufferLimitBytes *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes" json:"per_connection_buffer_limit_bytes,omitempty"`
	// Listener metadata.
	Metadata *envoy_api_v2_core1.Metadata `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	// [#not-implemented-hide:]
	DeprecatedV1 *Listener_DeprecatedV1 `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
	// The type of draining to perform at a listener-wide level.
	DrainType Listener_DrainType `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
	// Listener filters have the opportunity to manipulate and augment the connection metadata that
	// is used in connection filter chain matching, for example. These filters are run before any in
	// :ref:`filter_chains <envoy_api_field_Listener.filter_chains>`. Order matters as the
	// filters are processed sequentially right after a socket has been accepted by the listener, and
	// before a connection is created.
	ListenerFilters []envoy_api_v2_listener.ListenerFilter `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters" json:"listener_filters"`
	// Whether the listener should be set as a transparent socket.
	// When this flag is set to true, connections can be redirected to the listener using an
	// *iptables* *TPROXY* target, in which case the original source and destination addresses and
	// ports are preserved on accepted connections. This flag should be used in combination with
	// :ref:`an original_dst <config_listener_filters_original_dst>` :ref:`listener filter
	// <envoy_api_field_Listener.listener_filters>` to mark the connections' local addresses as
	// "restored." This can be used to hand off each redirected connection to another listener
	// associated with the connection's destination address. Direct connections to the socket without
	// using *TPROXY* cannot be distinguished from connections redirected using *TPROXY* and are
	// therefore treated as if they were redirected.
	// When this flag is set to false, the listener's socket is explicitly reset as non-transparent.
	// Setting this flag requires Envoy to run with the *CAP_NET_ADMIN* capability.
	// When this flag is not set (default), the socket is not modified, i.e. the transparent option
	// is neither set nor reset.
	Transparent *google_protobuf.BoolValue `protobuf:"bytes,10,opt,name=transparent" json:"transparent,omitempty"`
	// [#not-implemented-hide:] Whether the listener should set the IP_FREEBIND socket option. When
	// this flag is set to true, listeners can be bound to an IP address that is not configured on
	// the system running Envoy.
	// When this flag is set to false, the option IP_FREEBIND is disabled on the socket.
	// When this flag is not set (default), the socket is not modified, i.e. the option is neither
	// enabled nor disabled.
	Freebind *google_protobuf.BoolValue `protobuf:"bytes,11,opt,name=freebind" json:"freebind,omitempty"`
}

func (m *Listener) Reset()                    { *m = Listener{} }
func (m *Listener) String() string            { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()               {}
func (*Listener) Descriptor() ([]byte, []int) { return fileDescriptorLds, []int{0} }

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() envoy_api_v2_core.Address {
	if m != nil {
		return m.Address
	}
	return envoy_api_v2_core.Address{}
}

func (m *Listener) GetFilterChains() []envoy_api_v2_listener.FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

func (m *Listener) GetUseOriginalDst() *google_protobuf.BoolValue {
	if m != nil {
		return m.UseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *google_protobuf.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *envoy_api_v2_core1.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilters() []envoy_api_v2_listener.ListenerFilter {
	if m != nil {
		return m.ListenerFilters
	}
	return nil
}

func (m *Listener) GetTransparent() *google_protobuf.BoolValue {
	if m != nil {
		return m.Transparent
	}
	return nil
}

func (m *Listener) GetFreebind() *google_protobuf.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

// [#not-implemented-hide:]
type Listener_DeprecatedV1 struct {
	// Whether the listener should bind to the port. A listener that doesn’t
	// bind can only receive connections redirected from other listeners that
	// set use_original_dst parameter to true. Default is true.
	//
	// [V2-API-DIFF] This is deprecated in v2, all Listeners will bind to their
	// port. An additional filter chain must be created for every original
	// destination port this listener may redirect to in v2, with the original
	// port specified in the FilterChainMatch destination_port field.
	BindToPort *google_protobuf.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort" json:"bind_to_port,omitempty"`
}

func (m *Listener_DeprecatedV1) Reset()                    { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()               {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptorLds, []int{0, 0} }

func (m *Listener_DeprecatedV1) GetBindToPort() *google_protobuf.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

func init() {
	proto.RegisterType((*Listener)(nil), "envoy.api.v2.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v2.Listener.DeprecatedV1")
	proto.RegisterEnum("envoy.api.v2.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
}
func (this *Listener) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Listener)
	if !ok {
		that2, ok := that.(Listener)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Address.Equal(&that1.Address) {
		return false
	}
	if len(this.FilterChains) != len(that1.FilterChains) {
		return false
	}
	for i := range this.FilterChains {
		if !this.FilterChains[i].Equal(&that1.FilterChains[i]) {
			return false
		}
	}
	if !this.UseOriginalDst.Equal(that1.UseOriginalDst) {
		return false
	}
	if !this.PerConnectionBufferLimitBytes.Equal(that1.PerConnectionBufferLimitBytes) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.DeprecatedV1.Equal(that1.DeprecatedV1) {
		return false
	}
	if this.DrainType != that1.DrainType {
		return false
	}
	if len(this.ListenerFilters) != len(that1.ListenerFilters) {
		return false
	}
	for i := range this.ListenerFilters {
		if !this.ListenerFilters[i].Equal(&that1.ListenerFilters[i]) {
			return false
		}
	}
	if !this.Transparent.Equal(that1.Transparent) {
		return false
	}
	if !this.Freebind.Equal(that1.Freebind) {
		return false
	}
	return true
}
func (this *Listener_DeprecatedV1) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Listener_DeprecatedV1)
	if !ok {
		that2, ok := that.(Listener_DeprecatedV1)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.BindToPort.Equal(that1.BindToPort) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ListenerDiscoveryService service

type ListenerDiscoveryServiceClient interface {
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.ListenerDiscoveryService/FetchListeners", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ListenerDiscoveryService service

type ListenerDiscoveryServiceServer interface {
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/lds.proto",
}

func (m *Listener) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listener) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLds(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintLds(dAtA, i, uint64(m.Address.Size()))
	n1, err := m.Address.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.FilterChains) > 0 {
		for _, msg := range m.FilterChains {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintLds(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.UseOriginalDst != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.UseOriginalDst.Size()))
		n2, err := m.UseOriginalDst.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.PerConnectionBufferLimitBytes != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.PerConnectionBufferLimitBytes.Size()))
		n3, err := m.PerConnectionBufferLimitBytes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Metadata != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.Metadata.Size()))
		n4, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.DeprecatedV1 != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.DeprecatedV1.Size()))
		n5, err := m.DeprecatedV1.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.DrainType != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.DrainType))
	}
	if len(m.ListenerFilters) > 0 {
		for _, msg := range m.ListenerFilters {
			dAtA[i] = 0x4a
			i++
			i = encodeVarintLds(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Transparent != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.Transparent.Size()))
		n6, err := m.Transparent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Freebind != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.Freebind.Size()))
		n7, err := m.Freebind.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *Listener_DeprecatedV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listener_DeprecatedV1) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BindToPort != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLds(dAtA, i, uint64(m.BindToPort.Size()))
		n8, err := m.BindToPort.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func encodeVarintLds(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Listener) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLds(uint64(l))
	}
	l = m.Address.Size()
	n += 1 + l + sovLds(uint64(l))
	if len(m.FilterChains) > 0 {
		for _, e := range m.FilterChains {
			l = e.Size()
			n += 1 + l + sovLds(uint64(l))
		}
	}
	if m.UseOriginalDst != nil {
		l = m.UseOriginalDst.Size()
		n += 1 + l + sovLds(uint64(l))
	}
	if m.PerConnectionBufferLimitBytes != nil {
		l = m.PerConnectionBufferLimitBytes.Size()
		n += 1 + l + sovLds(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovLds(uint64(l))
	}
	if m.DeprecatedV1 != nil {
		l = m.DeprecatedV1.Size()
		n += 1 + l + sovLds(uint64(l))
	}
	if m.DrainType != 0 {
		n += 1 + sovLds(uint64(m.DrainType))
	}
	if len(m.ListenerFilters) > 0 {
		for _, e := range m.ListenerFilters {
			l = e.Size()
			n += 1 + l + sovLds(uint64(l))
		}
	}
	if m.Transparent != nil {
		l = m.Transparent.Size()
		n += 1 + l + sovLds(uint64(l))
	}
	if m.Freebind != nil {
		l = m.Freebind.Size()
		n += 1 + l + sovLds(uint64(l))
	}
	return n
}

func (m *Listener_DeprecatedV1) Size() (n int) {
	var l int
	_ = l
	if m.BindToPort != nil {
		l = m.BindToPort.Size()
		n += 1 + l + sovLds(uint64(l))
	}
	return n
}

func sovLds(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLds(x uint64) (n int) {
	return sovLds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Listener) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLds
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
			return fmt.Errorf("proto: Listener: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Listener: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterChains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilterChains = append(m.FilterChains, envoy_api_v2_listener.FilterChain{})
			if err := m.FilterChains[len(m.FilterChains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseOriginalDst", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UseOriginalDst == nil {
				m.UseOriginalDst = &google_protobuf.BoolValue{}
			}
			if err := m.UseOriginalDst.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerConnectionBufferLimitBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PerConnectionBufferLimitBytes == nil {
				m.PerConnectionBufferLimitBytes = &google_protobuf.UInt32Value{}
			}
			if err := m.PerConnectionBufferLimitBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &envoy_api_v2_core1.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedV1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeprecatedV1 == nil {
				m.DeprecatedV1 = &Listener_DeprecatedV1{}
			}
			if err := m.DeprecatedV1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrainType", wireType)
			}
			m.DrainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DrainType |= (Listener_DrainType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenerFilters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenerFilters = append(m.ListenerFilters, envoy_api_v2_listener.ListenerFilter{})
			if err := m.ListenerFilters[len(m.ListenerFilters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transparent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Transparent == nil {
				m.Transparent = &google_protobuf.BoolValue{}
			}
			if err := m.Transparent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Freebind", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Freebind == nil {
				m.Freebind = &google_protobuf.BoolValue{}
			}
			if err := m.Freebind.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLds
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
func (m *Listener_DeprecatedV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLds
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
			return fmt.Errorf("proto: DeprecatedV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeprecatedV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BindToPort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLds
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
				return ErrInvalidLengthLds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BindToPort == nil {
				m.BindToPort = &google_protobuf.BoolValue{}
			}
			if err := m.BindToPort.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLds
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
func skipLds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLds
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
					return 0, ErrIntOverflowLds
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
					return 0, ErrIntOverflowLds
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
				return 0, ErrInvalidLengthLds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLds
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
				next, err := skipLds(dAtA[start:])
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
	ErrInvalidLengthLds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLds   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/api/v2/lds.proto", fileDescriptorLds) }

var fileDescriptorLds = []byte{
	// 716 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6b, 0x13, 0x4d,
	0x18, 0xc7, 0x3b, 0xe9, 0xaf, 0x64, 0x92, 0xa6, 0x61, 0xde, 0x97, 0xb7, 0x4b, 0xde, 0x9a, 0xc6,
	0xa8, 0x10, 0x3d, 0x6c, 0x6c, 0x0a, 0x0a, 0xa5, 0x20, 0x4d, 0x63, 0x68, 0x21, 0xb5, 0xb2, 0xfd,
	0xa1, 0x3d, 0x2d, 0x93, 0xec, 0x93, 0x74, 0x60, 0x33, 0xb3, 0xce, 0x4c, 0x22, 0xb9, 0x7a, 0x12,
	0x8f, 0xfa, 0x4f, 0xf8, 0x37, 0x78, 0xf2, 0xd8, 0xa3, 0xe0, 0x5d, 0x24, 0x78, 0x11, 0xff, 0x03,
	0x4f, 0xb2, 0x9b, 0xdd, 0x35, 0xa1, 0xad, 0xbd, 0x78, 0x7b, 0xf6, 0x79, 0xbe, 0xcf, 0x67, 0x86,
	0xef, 0x97, 0x1d, 0xfc, 0x1f, 0xf0, 0x81, 0x18, 0x56, 0xa8, 0xc7, 0x2a, 0x83, 0x6a, 0xc5, 0x75,
	0x94, 0xe9, 0x49, 0xa1, 0x05, 0xc9, 0x04, 0x7d, 0x93, 0x7a, 0xcc, 0x1c, 0x54, 0xf3, 0x6b, 0x53,
	0xaa, 0xb6, 0x90, 0x50, 0xa1, 0x8e, 0x23, 0x41, 0x85, 0xf2, 0xfc, 0xea, 0x45, 0x41, 0x8b, 0x2a,
	0xb8, 0x74, 0xea, 0x30, 0xd5, 0x16, 0x03, 0x90, 0xc3, 0x70, 0x7a, 0x7b, 0xfa, 0x0a, 0x4c, 0x69,
	0xe0, 0x20, 0xe3, 0x22, 0x62, 0x74, 0x85, 0xe8, 0xba, 0x10, 0xc8, 0x28, 0xe7, 0x42, 0x53, 0xcd,
	0x04, 0x8f, 0xce, 0x2f, 0x84, 0xd3, 0xe0, 0xab, 0xd5, 0xef, 0x54, 0x5e, 0x4a, 0xea, 0x79, 0x20,
	0xa3, 0xf9, 0xca, 0x80, 0xba, 0xcc, 0xa1, 0x1a, 0x2a, 0x51, 0x11, 0x0e, 0xfe, 0xed, 0x8a, 0xae,
	0x08, 0xca, 0x8a, 0x5f, 0x8d, 0xbb, 0xa5, 0x9f, 0x0b, 0x38, 0xd9, 0x0c, 0xcf, 0x27, 0x04, 0xcf,
	0x71, 0xda, 0x03, 0x03, 0x15, 0x51, 0x39, 0x65, 0x05, 0x35, 0xa9, 0xe3, 0xc5, 0xd0, 0x00, 0x23,
	0x51, 0x44, 0xe5, 0x74, 0x35, 0x6f, 0x4e, 0x1a, 0x66, 0xfa, 0x0e, 0x98, 0xdb, 0x63, 0x45, 0x2d,
	0x7b, 0xfe, 0x65, 0x6d, 0xe6, 0xc3, 0xf7, 0x8f, 0xb3, 0xf3, 0x6f, 0x50, 0x22, 0x87, 0xac, 0x68,
	0x95, 0x3c, 0xc3, 0x4b, 0x1d, 0xe6, 0x6a, 0x90, 0x76, 0xfb, 0x8c, 0x32, 0xae, 0x8c, 0xd9, 0xe2,
	0x6c, 0x39, 0x5d, 0x2d, 0x4d, 0xb3, 0x62, 0x23, 0x1a, 0x81, 0x76, 0xc7, 0x97, 0x4e, 0x30, 0xdf,
	0xa2, 0x44, 0x12, 0x59, 0x99, 0xce, 0xef, 0xa1, 0x22, 0xbb, 0x38, 0xd7, 0x57, 0x60, 0x0b, 0xc9,
	0xba, 0x8c, 0x53, 0xd7, 0x76, 0x94, 0x36, 0xe6, 0xc2, 0x7b, 0x8e, 0x9d, 0x32, 0x23, 0xa7, 0xcc,
	0x9a, 0x10, 0xee, 0x09, 0x75, 0xfb, 0x50, 0x4b, 0x18, 0xc8, 0xca, 0xf6, 0x15, 0x1c, 0x84, 0x6b,
	0x75, 0xa5, 0x49, 0x07, 0xdf, 0xf4, 0xfc, 0xfb, 0x09, 0xce, 0xa1, 0xed, 0x3b, 0x6e, 0xb7, 0xfa,
	0x9d, 0x0e, 0x48, 0xdb, 0x65, 0x3d, 0xa6, 0xed, 0xd6, 0x50, 0x83, 0x32, 0xe6, 0x03, 0xf4, 0xea,
	0x05, 0xf4, 0xf1, 0x1e, 0xd7, 0x1b, 0xd5, 0x00, 0x6e, 0xdd, 0xf0, 0x40, 0xee, 0xc4, 0x94, 0x5a,
	0x00, 0x69, 0xfa, 0x8c, 0x9a, 0x8f, 0x20, 0x0f, 0x71, 0xb2, 0x07, 0x9a, 0x3a, 0x54, 0x53, 0x63,
	0x21, 0xc0, 0xfd, 0x7f, 0x89, 0xa3, 0xfb, 0xa1, 0xc4, 0x8a, 0xc5, 0x64, 0x17, 0x2f, 0x39, 0xe0,
	0x49, 0x68, 0x53, 0x0d, 0x8e, 0x3d, 0x58, 0x37, 0x16, 0x83, 0xed, 0x5b, 0xd3, 0xdb, 0x51, 0x98,
	0x66, 0x3d, 0xd6, 0x9e, 0xac, 0x5b, 0x19, 0x67, 0xe2, 0x8b, 0x3c, 0xc2, 0xd8, 0x91, 0x94, 0x71,
	0x5b, 0x0f, 0x3d, 0x30, 0x92, 0x45, 0x54, 0xce, 0x56, 0x8b, 0x57, 0x61, 0x7c, 0xe1, 0xd1, 0xd0,
	0x03, 0x2b, 0xe5, 0x44, 0x25, 0x39, 0xc1, 0xb9, 0x28, 0x2b, 0x7b, 0x1c, 0x87, 0x32, 0x52, 0x41,
	0xa2, 0x77, 0xae, 0x48, 0x34, 0xe2, 0x8d, 0x93, 0xad, 0xcd, 0xf9, 0xa1, 0x5a, 0xcb, 0xee, 0x54,
	0x57, 0x91, 0x2d, 0x9c, 0xd6, 0x92, 0x72, 0xe5, 0x51, 0x09, 0x5c, 0x1b, 0xf8, 0xba, 0x20, 0xad,
	0x49, 0x39, 0x79, 0x80, 0x93, 0x1d, 0x09, 0xd0, 0x62, 0xdc, 0x31, 0xd2, 0xd7, 0xae, 0xc6, 0xda,
	0x7c, 0x13, 0x67, 0x26, 0xcd, 0x22, 0x5b, 0x38, 0xe3, 0xf7, 0x6d, 0x2d, 0x6c, 0x4f, 0x48, 0x1d,
	0xfc, 0x0e, 0x7f, 0x66, 0x61, 0x5f, 0x7f, 0x24, 0x9e, 0x0a, 0xa9, 0x4b, 0x77, 0x71, 0x2a, 0xf6,
	0x8c, 0xa4, 0xf1, 0x62, 0xfd, 0x71, 0x63, 0xfb, 0xb8, 0x79, 0x94, 0x9b, 0x21, 0xcb, 0x38, 0xbd,
	0x7f, 0x50, 0xdf, 0x6b, 0x9c, 0xda, 0x07, 0x4f, 0x9a, 0xa7, 0x39, 0x54, 0xfd, 0x81, 0xb0, 0x11,
	0x19, 0x53, 0x8f, 0xde, 0x8a, 0x43, 0x90, 0x03, 0xd6, 0x06, 0xf2, 0x1c, 0x2f, 0x1f, 0x6a, 0x09,
	0xb4, 0x17, 0x29, 0x14, 0x29, 0x4c, 0x9b, 0x1b, 0xaf, 0x58, 0xf0, 0xa2, 0x0f, 0x4a, 0xe7, 0xd7,
	0xae, 0x9c, 0x2b, 0x4f, 0x70, 0x05, 0xa5, 0x99, 0x32, 0xba, 0x8f, 0x48, 0x1f, 0x67, 0x1b, 0xa0,
	0xdb, 0x67, 0x7f, 0x11, 0x5c, 0x7a, 0xf5, 0xf9, 0xdb, 0xbb, 0xc4, 0x6a, 0x69, 0x65, 0xea, 0xd9,
	0xdb, 0x8c, 0xf2, 0x55, 0x9b, 0xe8, 0x5e, 0xed, 0x9f, 0xf7, 0xa3, 0x02, 0x3a, 0x1f, 0x15, 0xd0,
	0xa7, 0x51, 0x01, 0x7d, 0x1d, 0x15, 0xd0, 0x6b, 0x84, 0x5a, 0x0b, 0x81, 0x9b, 0x1b, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0x51, 0x8d, 0x4f, 0x9e, 0x05, 0x00, 0x00,
}
