// Code generated by protoc-gen-go. DO NOT EDIT.
// source: egress.proto

package loggregator_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type EgressRequest struct {
	// shard_id instructs Loggregator to shard envelopes between other
	// subscriptions with the same shard_id. Loggregator will do its best to
	// split the load evenly between subscriptions with the same shard_id
	// (unless deterministic_name is set).
	ShardId string `protobuf:"bytes,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	// deterministic_name is used to enable deterministic routing. This implies
	// that gauges and counters are routed based on name. If this is excluded,
	// then they are routed to split load evenly.
	DeterministicName string `protobuf:"bytes,5,opt,name=deterministic_name,json=deterministicName,proto3" json:"deterministic_name,omitempty"`
	// TODO: This can be removed once selector has been around long enough.
	LegacySelector *Selector `protobuf:"bytes,2,opt,name=legacy_selector,json=legacySelector,proto3" json:"legacy_selector,omitempty"`
	// selector is the preferred (over legacy_selector) mechanism to select
	// what envelope types the subscription wants. If there are no selectors
	// given, no data will be sent.
	Selectors []*Selector `protobuf:"bytes,4,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// TODO: This can be removed once the envelope.deprecated_tags is removed.
	UsePreferredTags     bool     `protobuf:"varint,3,opt,name=use_preferred_tags,json=usePreferredTags,proto3" json:"use_preferred_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EgressRequest) Reset()         { *m = EgressRequest{} }
func (m *EgressRequest) String() string { return proto.CompactTextString(m) }
func (*EgressRequest) ProtoMessage()    {}
func (*EgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{0}
}
func (m *EgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EgressRequest.Unmarshal(m, b)
}
func (m *EgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EgressRequest.Marshal(b, m, deterministic)
}
func (dst *EgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EgressRequest.Merge(dst, src)
}
func (m *EgressRequest) XXX_Size() int {
	return xxx_messageInfo_EgressRequest.Size(m)
}
func (m *EgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EgressRequest proto.InternalMessageInfo

func (m *EgressRequest) GetShardId() string {
	if m != nil {
		return m.ShardId
	}
	return ""
}

func (m *EgressRequest) GetDeterministicName() string {
	if m != nil {
		return m.DeterministicName
	}
	return ""
}

func (m *EgressRequest) GetLegacySelector() *Selector {
	if m != nil {
		return m.LegacySelector
	}
	return nil
}

func (m *EgressRequest) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *EgressRequest) GetUsePreferredTags() bool {
	if m != nil {
		return m.UsePreferredTags
	}
	return false
}

type EgressBatchRequest struct {
	// shard_id instructs Loggregator to shard envelopes between other
	// subscriptions with the same shard_id. Loggregator will do its best to
	// split the load evenly between subscriptions with the same shard_id
	// (unless deterministic_name is set).
	ShardId string `protobuf:"bytes,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	// deterministic_name is used to enable deterministic routing. This implies
	// that gauges and counters are routed based on name. If this is excluded,
	// then they are routed to split load evenly.
	DeterministicName string `protobuf:"bytes,5,opt,name=deterministic_name,json=deterministicName,proto3" json:"deterministic_name,omitempty"`
	// TODO: This can be removed once selector has been around long enough.
	LegacySelector *Selector `protobuf:"bytes,2,opt,name=legacy_selector,json=legacySelector,proto3" json:"legacy_selector,omitempty"`
	// selector is the preferred (over legacy_selector) mechanism to select
	// what envelope types the subscription wants. If there are no selectors
	// given, no data will be sent.
	Selectors []*Selector `protobuf:"bytes,4,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// TODO: This can be removed once the envelope.deprecated_tags is removed.
	UsePreferredTags     bool     `protobuf:"varint,3,opt,name=use_preferred_tags,json=usePreferredTags,proto3" json:"use_preferred_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EgressBatchRequest) Reset()         { *m = EgressBatchRequest{} }
func (m *EgressBatchRequest) String() string { return proto.CompactTextString(m) }
func (*EgressBatchRequest) ProtoMessage()    {}
func (*EgressBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{1}
}
func (m *EgressBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EgressBatchRequest.Unmarshal(m, b)
}
func (m *EgressBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EgressBatchRequest.Marshal(b, m, deterministic)
}
func (dst *EgressBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EgressBatchRequest.Merge(dst, src)
}
func (m *EgressBatchRequest) XXX_Size() int {
	return xxx_messageInfo_EgressBatchRequest.Size(m)
}
func (m *EgressBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EgressBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EgressBatchRequest proto.InternalMessageInfo

func (m *EgressBatchRequest) GetShardId() string {
	if m != nil {
		return m.ShardId
	}
	return ""
}

func (m *EgressBatchRequest) GetDeterministicName() string {
	if m != nil {
		return m.DeterministicName
	}
	return ""
}

func (m *EgressBatchRequest) GetLegacySelector() *Selector {
	if m != nil {
		return m.LegacySelector
	}
	return nil
}

func (m *EgressBatchRequest) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *EgressBatchRequest) GetUsePreferredTags() bool {
	if m != nil {
		return m.UsePreferredTags
	}
	return false
}

// Selector instructs Loggregator to only send envelopes that match the given
// criteria.
type Selector struct {
	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	// Types that are valid to be assigned to Message:
	//	*Selector_Log
	//	*Selector_Counter
	//	*Selector_Gauge
	//	*Selector_Timer
	//	*Selector_Event
	Message              isSelector_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{2}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (dst *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(dst, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

type isSelector_Message interface {
	isSelector_Message()
}

type Selector_Log struct {
	Log *LogSelector `protobuf:"bytes,2,opt,name=log,proto3,oneof"`
}

type Selector_Counter struct {
	Counter *CounterSelector `protobuf:"bytes,3,opt,name=counter,proto3,oneof"`
}

type Selector_Gauge struct {
	Gauge *GaugeSelector `protobuf:"bytes,4,opt,name=gauge,proto3,oneof"`
}

type Selector_Timer struct {
	Timer *TimerSelector `protobuf:"bytes,5,opt,name=timer,proto3,oneof"`
}

type Selector_Event struct {
	Event *EventSelector `protobuf:"bytes,6,opt,name=event,proto3,oneof"`
}

func (*Selector_Log) isSelector_Message() {}

func (*Selector_Counter) isSelector_Message() {}

func (*Selector_Gauge) isSelector_Message() {}

func (*Selector_Timer) isSelector_Message() {}

func (*Selector_Event) isSelector_Message() {}

func (m *Selector) GetMessage() isSelector_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Selector) GetLog() *LogSelector {
	if x, ok := m.GetMessage().(*Selector_Log); ok {
		return x.Log
	}
	return nil
}

func (m *Selector) GetCounter() *CounterSelector {
	if x, ok := m.GetMessage().(*Selector_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *Selector) GetGauge() *GaugeSelector {
	if x, ok := m.GetMessage().(*Selector_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (m *Selector) GetTimer() *TimerSelector {
	if x, ok := m.GetMessage().(*Selector_Timer); ok {
		return x.Timer
	}
	return nil
}

func (m *Selector) GetEvent() *EventSelector {
	if x, ok := m.GetMessage().(*Selector_Event); ok {
		return x.Event
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Selector) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Selector_OneofMarshaler, _Selector_OneofUnmarshaler, _Selector_OneofSizer, []interface{}{
		(*Selector_Log)(nil),
		(*Selector_Counter)(nil),
		(*Selector_Gauge)(nil),
		(*Selector_Timer)(nil),
		(*Selector_Event)(nil),
	}
}

func _Selector_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Selector)
	// Message
	switch x := m.Message.(type) {
	case *Selector_Log:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case *Selector_Counter:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Counter); err != nil {
			return err
		}
	case *Selector_Gauge:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gauge); err != nil {
			return err
		}
	case *Selector_Timer:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Timer); err != nil {
			return err
		}
	case *Selector_Event:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Event); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Selector.Message has unexpected type %T", x)
	}
	return nil
}

func _Selector_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Selector)
	switch tag {
	case 2: // Message.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogSelector)
		err := b.DecodeMessage(msg)
		m.Message = &Selector_Log{msg}
		return true, err
	case 3: // Message.counter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CounterSelector)
		err := b.DecodeMessage(msg)
		m.Message = &Selector_Counter{msg}
		return true, err
	case 4: // Message.gauge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GaugeSelector)
		err := b.DecodeMessage(msg)
		m.Message = &Selector_Gauge{msg}
		return true, err
	case 5: // Message.timer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TimerSelector)
		err := b.DecodeMessage(msg)
		m.Message = &Selector_Timer{msg}
		return true, err
	case 6: // Message.event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventSelector)
		err := b.DecodeMessage(msg)
		m.Message = &Selector_Event{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Selector_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Selector)
	// Message
	switch x := m.Message.(type) {
	case *Selector_Log:
		s := proto.Size(x.Log)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Selector_Counter:
		s := proto.Size(x.Counter)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Selector_Gauge:
		s := proto.Size(x.Gauge)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Selector_Timer:
		s := proto.Size(x.Timer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Selector_Event:
		s := proto.Size(x.Event)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// LogSelector instructs Loggregator to egress Log envelopes to the given
// subscription.
type LogSelector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogSelector) Reset()         { *m = LogSelector{} }
func (m *LogSelector) String() string { return proto.CompactTextString(m) }
func (*LogSelector) ProtoMessage()    {}
func (*LogSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{3}
}
func (m *LogSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSelector.Unmarshal(m, b)
}
func (m *LogSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSelector.Marshal(b, m, deterministic)
}
func (dst *LogSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSelector.Merge(dst, src)
}
func (m *LogSelector) XXX_Size() int {
	return xxx_messageInfo_LogSelector.Size(m)
}
func (m *LogSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSelector.DiscardUnknown(m)
}

var xxx_messageInfo_LogSelector proto.InternalMessageInfo

// GaugeSelector instructs Loggregator to egress Gauge envelopes to the
// given subscription.
type GaugeSelector struct {
	// Any egress Gauge envelope must consist of the given names.
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GaugeSelector) Reset()         { *m = GaugeSelector{} }
func (m *GaugeSelector) String() string { return proto.CompactTextString(m) }
func (*GaugeSelector) ProtoMessage()    {}
func (*GaugeSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{4}
}
func (m *GaugeSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaugeSelector.Unmarshal(m, b)
}
func (m *GaugeSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaugeSelector.Marshal(b, m, deterministic)
}
func (dst *GaugeSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaugeSelector.Merge(dst, src)
}
func (m *GaugeSelector) XXX_Size() int {
	return xxx_messageInfo_GaugeSelector.Size(m)
}
func (m *GaugeSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_GaugeSelector.DiscardUnknown(m)
}

var xxx_messageInfo_GaugeSelector proto.InternalMessageInfo

func (m *GaugeSelector) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// CounterSelector instructs Loggregator to egress Counter envelopes to the
// given subscription
type CounterSelector struct {
	// Any egress Counter envelope must have the given name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CounterSelector) Reset()         { *m = CounterSelector{} }
func (m *CounterSelector) String() string { return proto.CompactTextString(m) }
func (*CounterSelector) ProtoMessage()    {}
func (*CounterSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{5}
}
func (m *CounterSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CounterSelector.Unmarshal(m, b)
}
func (m *CounterSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CounterSelector.Marshal(b, m, deterministic)
}
func (dst *CounterSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CounterSelector.Merge(dst, src)
}
func (m *CounterSelector) XXX_Size() int {
	return xxx_messageInfo_CounterSelector.Size(m)
}
func (m *CounterSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_CounterSelector.DiscardUnknown(m)
}

var xxx_messageInfo_CounterSelector proto.InternalMessageInfo

func (m *CounterSelector) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// TimerSelector instructs Loggregator to egress Timer envelopes to the given
// subscription.
type TimerSelector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimerSelector) Reset()         { *m = TimerSelector{} }
func (m *TimerSelector) String() string { return proto.CompactTextString(m) }
func (*TimerSelector) ProtoMessage()    {}
func (*TimerSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{6}
}
func (m *TimerSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimerSelector.Unmarshal(m, b)
}
func (m *TimerSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimerSelector.Marshal(b, m, deterministic)
}
func (dst *TimerSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimerSelector.Merge(dst, src)
}
func (m *TimerSelector) XXX_Size() int {
	return xxx_messageInfo_TimerSelector.Size(m)
}
func (m *TimerSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_TimerSelector.DiscardUnknown(m)
}

var xxx_messageInfo_TimerSelector proto.InternalMessageInfo

// EventSelector instructs Loggregator to egress Event envelopes to the given
// subscription.
type EventSelector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSelector) Reset()         { *m = EventSelector{} }
func (m *EventSelector) String() string { return proto.CompactTextString(m) }
func (*EventSelector) ProtoMessage()    {}
func (*EventSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_egress_fcae6bb65dce0d2e, []int{7}
}
func (m *EventSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSelector.Unmarshal(m, b)
}
func (m *EventSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSelector.Marshal(b, m, deterministic)
}
func (dst *EventSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSelector.Merge(dst, src)
}
func (m *EventSelector) XXX_Size() int {
	return xxx_messageInfo_EventSelector.Size(m)
}
func (m *EventSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSelector.DiscardUnknown(m)
}

var xxx_messageInfo_EventSelector proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EgressRequest)(nil), "loggregator.v2.EgressRequest")
	proto.RegisterType((*EgressBatchRequest)(nil), "loggregator.v2.EgressBatchRequest")
	proto.RegisterType((*Selector)(nil), "loggregator.v2.Selector")
	proto.RegisterType((*LogSelector)(nil), "loggregator.v2.LogSelector")
	proto.RegisterType((*GaugeSelector)(nil), "loggregator.v2.GaugeSelector")
	proto.RegisterType((*CounterSelector)(nil), "loggregator.v2.CounterSelector")
	proto.RegisterType((*TimerSelector)(nil), "loggregator.v2.TimerSelector")
	proto.RegisterType((*EventSelector)(nil), "loggregator.v2.EventSelector")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EgressClient is the client API for Egress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EgressClient interface {
	Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error)
	BatchedReceiver(ctx context.Context, in *EgressBatchRequest, opts ...grpc.CallOption) (Egress_BatchedReceiverClient, error)
}

type egressClient struct {
	cc *grpc.ClientConn
}

func NewEgressClient(cc *grpc.ClientConn) EgressClient {
	return &egressClient{cc}
}

func (c *egressClient) Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Egress_serviceDesc.Streams[0], "/loggregator.v2.Egress/Receiver", opts...)
	if err != nil {
		return nil, err
	}
	x := &egressReceiverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Egress_ReceiverClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type egressReceiverClient struct {
	grpc.ClientStream
}

func (x *egressReceiverClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *egressClient) BatchedReceiver(ctx context.Context, in *EgressBatchRequest, opts ...grpc.CallOption) (Egress_BatchedReceiverClient, error) {
	fmt.Println("before  BatchedReceiver-grpc.NewClientStream")
	stream, err := c.cc.NewStream(ctx, &_Egress_serviceDesc.Streams[1], "/loggregator.v2.Egress/BatchedReceiver", opts...)
	fmt.Println("after  BatchedReceiver-grpc.NewClientStream")
	if err != nil {
		return nil, err
	}
	x := &egressBatchedReceiverClient{stream}
	fmt.Println("before  BatchedReceiver-SendMsg")
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	fmt.Println("after  BatchedReceiver-SendMsg")
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	fmt.Println("after  BatchedReceiver-closesend")
	return x, nil
}

type Egress_BatchedReceiverClient interface {
	Recv() (*EnvelopeBatch, error)
	grpc.ClientStream
}

type egressBatchedReceiverClient struct {
	grpc.ClientStream
}

func (x *egressBatchedReceiverClient) Recv() (*EnvelopeBatch, error) {
	m := new(EnvelopeBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EgressServer is the server API for Egress service.
type EgressServer interface {
	Receiver(*EgressRequest, Egress_ReceiverServer) error
	BatchedReceiver(*EgressBatchRequest, Egress_BatchedReceiverServer) error
}

func RegisterEgressServer(s *grpc.Server, srv EgressServer) {
	s.RegisterService(&_Egress_serviceDesc, srv)
}

func _Egress_Receiver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EgressServer).Receiver(m, &egressReceiverServer{stream})
}

type Egress_ReceiverServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type egressReceiverServer struct {
	grpc.ServerStream
}

func (x *egressReceiverServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func _Egress_BatchedReceiver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EgressBatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EgressServer).BatchedReceiver(m, &egressBatchedReceiverServer{stream})
}

type Egress_BatchedReceiverServer interface {
	Send(*EnvelopeBatch) error
	grpc.ServerStream
}

type egressBatchedReceiverServer struct {
	grpc.ServerStream
}

func (x *egressBatchedReceiverServer) Send(m *EnvelopeBatch) error {
	return x.ServerStream.SendMsg(m)
}

var _Egress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.Egress",
	HandlerType: (*EgressServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receiver",
			Handler:       _Egress_Receiver_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BatchedReceiver",
			Handler:       _Egress_BatchedReceiver_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "egress.proto",
}

func init() { proto.RegisterFile("egress.proto", fileDescriptor_egress_fcae6bb65dce0d2e) }

var fileDescriptor_egress_fcae6bb65dce0d2e = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xeb, 0xfc, 0xb5, 0x27, 0xbf, 0x24, 0xbf, 0xae, 0x38, 0x98, 0x54, 0x55, 0x2d, 0x4b,
	0x95, 0x72, 0x00, 0x83, 0xc2, 0x9f, 0x0b, 0x27, 0x82, 0xa2, 0x52, 0xa9, 0x20, 0x64, 0x7a, 0xe0,
	0x66, 0x2d, 0xf6, 0x74, 0x6b, 0xc9, 0xf1, 0x86, 0xdd, 0x75, 0xa4, 0x5e, 0xb9, 0xf0, 0x30, 0x3c,
	0x07, 0xef, 0x85, 0xbc, 0x1b, 0x37, 0xb6, 0x1b, 0x78, 0x01, 0x6e, 0x9e, 0xfd, 0x7e, 0x3f, 0xb3,
	0xb3, 0xe3, 0xd9, 0x85, 0xff, 0x90, 0x09, 0x94, 0x32, 0xd8, 0x08, 0xae, 0x38, 0x99, 0x64, 0x9c,
	0x31, 0x81, 0x8c, 0x2a, 0x2e, 0x82, 0xed, 0x62, 0x36, 0xc1, 0x7c, 0x8b, 0x19, 0xdf, 0xa0, 0xd1,
	0xfd, 0xef, 0x1d, 0x18, 0xaf, 0x34, 0x10, 0xe2, 0xb7, 0x02, 0xa5, 0x22, 0x8f, 0xc1, 0x96, 0xb7,
	0x54, 0x24, 0x51, 0x9a, 0xb8, 0x96, 0x67, 0xcd, 0x9d, 0x70, 0xa8, 0xe3, 0xcb, 0x84, 0x3c, 0x05,
	0x92, 0xa0, 0x42, 0xb1, 0x4e, 0xf3, 0x54, 0xaa, 0x34, 0x8e, 0x72, 0xba, 0x46, 0xb7, 0xaf, 0x4d,
	0xc7, 0x0d, 0xe5, 0x23, 0x5d, 0x23, 0x79, 0x0b, 0xd3, 0x0c, 0x19, 0x8d, 0xef, 0x22, 0x89, 0x19,
	0xc6, 0x8a, 0x0b, 0xb7, 0xe3, 0x59, 0xf3, 0xd1, 0xc2, 0x0d, 0x9a, 0x55, 0x05, 0x9f, 0x77, 0x7a,
	0x38, 0x31, 0x40, 0x15, 0x93, 0xd7, 0xe0, 0x54, 0xac, 0x74, 0x7b, 0x5e, 0xf7, 0xaf, 0xf0, 0xde,
	0x4a, 0x9e, 0x00, 0x29, 0x24, 0x46, 0x1b, 0x81, 0x37, 0x28, 0x04, 0x26, 0x91, 0xa2, 0x4c, 0xba,
	0x5d, 0xcf, 0x9a, 0xdb, 0xe1, 0xff, 0x85, 0xc4, 0x4f, 0x95, 0x70, 0x4d, 0x99, 0xf4, 0x7f, 0x74,
	0x80, 0x98, 0x26, 0x2c, 0xa9, 0x8a, 0x6f, 0xff, 0xe1, 0x4e, 0xfc, 0xea, 0x80, 0x7d, 0xbf, 0xe5,
	0x09, 0x38, 0x92, 0x17, 0x22, 0xc6, 0x7d, 0x03, 0x6c, 0xb3, 0x70, 0x99, 0x90, 0x67, 0xd0, 0xcd,
	0x38, 0xdb, 0x1d, 0xe3, 0xa4, 0x5d, 0xc9, 0x15, 0x67, 0x55, 0x9a, 0xf7, 0x47, 0x61, 0xe9, 0x24,
	0x6f, 0x60, 0x18, 0xf3, 0x22, 0x57, 0x28, 0xf4, 0xee, 0xa3, 0xc5, 0x59, 0x1b, 0x7a, 0x67, 0xe4,
	0x1a, 0x58, 0x11, 0xe4, 0x15, 0xf4, 0x19, 0x2d, 0x18, 0xba, 0x3d, 0x8d, 0x9e, 0xb6, 0xd1, 0x8b,
	0x52, 0xac, 0x81, 0xc6, 0x5d, 0x62, 0x2a, 0x5d, 0xa3, 0xd0, 0x7f, 0xe6, 0x00, 0x76, 0x5d, 0x8a,
	0x75, 0x4c, 0xbb, 0x4b, 0x0c, 0xb7, 0x98, 0x2b, 0x77, 0x70, 0x18, 0x5b, 0x95, 0x62, 0x1d, 0xd3,
	0xee, 0xa5, 0x03, 0xc3, 0x0f, 0x28, 0x25, 0x65, 0xe8, 0x8f, 0x61, 0x54, 0x6b, 0x81, 0x7f, 0x0e,
	0xe3, 0x46, 0x85, 0xe4, 0x11, 0xf4, 0xcb, 0x89, 0x91, 0xae, 0xe5, 0x75, 0xe7, 0x4e, 0x68, 0x02,
	0xff, 0x1c, 0xa6, 0xad, 0x1e, 0x10, 0x02, 0x3d, 0x3d, 0x5a, 0xa6, 0xfd, 0xfa, 0xdb, 0x9f, 0xc2,
	0xb8, 0x51, 0x78, 0xb9, 0xd0, 0x28, 0x69, 0xf1, 0xd3, 0x82, 0x81, 0x19, 0x68, 0x72, 0x01, 0x76,
	0x88, 0x31, 0xa6, 0x5b, 0x14, 0xe4, 0xe1, 0x41, 0xea, 0x37, 0x7f, 0xf6, 0x60, 0x9e, 0x56, 0xbb,
	0xb7, 0xc2, 0x3f, 0x7a, 0x6e, 0x91, 0x2f, 0x30, 0xd5, 0xb7, 0x03, 0x93, 0xfb, 0x7c, 0xfe, 0xe1,
	0x7c, 0xf5, 0x4b, 0x34, 0x3b, 0xfd, 0x53, 0x52, 0xed, 0x2a, 0x33, 0x2f, 0x5f, 0xc2, 0x19, 0x17,
	0x2c, 0x88, 0x33, 0x5e, 0x24, 0x37, 0xbc, 0xc8, 0x13, 0x71, 0xd7, 0x82, 0x96, 0xc7, 0x57, 0xfb,
	0xd8, 0x6c, 0xf2, 0x75, 0xa0, 0x1f, 0xb0, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x04, 0xbd,
	0xa8, 0xa8, 0xf0, 0x04, 0x00, 0x00,
}
