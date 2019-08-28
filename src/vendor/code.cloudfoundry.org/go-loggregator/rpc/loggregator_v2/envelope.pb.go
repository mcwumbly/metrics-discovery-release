// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envelope.proto

package loggregator_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Log_Type int32

const (
	Log_OUT Log_Type = 0
	Log_ERR Log_Type = 1
)

var Log_Type_name = map[int32]string{
	0: "OUT",
	1: "ERR",
}
var Log_Type_value = map[string]int32{
	"OUT": 0,
	"ERR": 1,
}

func (x Log_Type) String() string {
	return proto.EnumName(Log_Type_name, int32(x))
}
func (Log_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{3, 0}
}

type Envelope struct {
	Timestamp      int64             `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SourceId       string            `protobuf:"bytes,2,opt,name=source_id,proto3" json:"source_id,omitempty"`
	InstanceId     string            `protobuf:"bytes,8,opt,name=instance_id,proto3" json:"instance_id,omitempty"`
	DeprecatedTags map[string]*Value `protobuf:"bytes,3,rep,name=deprecated_tags,proto3" json:"deprecated_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags           map[string]string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to Message:
	//	*Envelope_Log
	//	*Envelope_Counter
	//	*Envelope_Gauge
	//	*Envelope_Timer
	//	*Envelope_Event
	Message              isEnvelope_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{0}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Envelope) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Envelope) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Envelope) GetDeprecatedTags() map[string]*Value {
	if m != nil {
		return m.DeprecatedTags
	}
	return nil
}

func (m *Envelope) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type isEnvelope_Message interface {
	isEnvelope_Message()
}

type Envelope_Log struct {
	Log *Log `protobuf:"bytes,4,opt,name=log,proto3,oneof"`
}

type Envelope_Counter struct {
	Counter *Counter `protobuf:"bytes,5,opt,name=counter,proto3,oneof"`
}

type Envelope_Gauge struct {
	Gauge *Gauge `protobuf:"bytes,6,opt,name=gauge,proto3,oneof"`
}

type Envelope_Timer struct {
	Timer *Timer `protobuf:"bytes,7,opt,name=timer,proto3,oneof"`
}

type Envelope_Event struct {
	Event *Event `protobuf:"bytes,10,opt,name=event,proto3,oneof"`
}

func (*Envelope_Log) isEnvelope_Message() {}

func (*Envelope_Counter) isEnvelope_Message() {}

func (*Envelope_Gauge) isEnvelope_Message() {}

func (*Envelope_Timer) isEnvelope_Message() {}

func (*Envelope_Event) isEnvelope_Message() {}

func (m *Envelope) GetMessage() isEnvelope_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetLog() *Log {
	if x, ok := m.GetMessage().(*Envelope_Log); ok {
		return x.Log
	}
	return nil
}

func (m *Envelope) GetCounter() *Counter {
	if x, ok := m.GetMessage().(*Envelope_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *Envelope) GetGauge() *Gauge {
	if x, ok := m.GetMessage().(*Envelope_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (m *Envelope) GetTimer() *Timer {
	if x, ok := m.GetMessage().(*Envelope_Timer); ok {
		return x.Timer
	}
	return nil
}

func (m *Envelope) GetEvent() *Event {
	if x, ok := m.GetMessage().(*Envelope_Event); ok {
		return x.Event
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Envelope) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Envelope_OneofMarshaler, _Envelope_OneofUnmarshaler, _Envelope_OneofSizer, []interface{}{
		(*Envelope_Log)(nil),
		(*Envelope_Counter)(nil),
		(*Envelope_Gauge)(nil),
		(*Envelope_Timer)(nil),
		(*Envelope_Event)(nil),
	}
}

func _Envelope_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case *Envelope_Counter:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Counter); err != nil {
			return err
		}
	case *Envelope_Gauge:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gauge); err != nil {
			return err
		}
	case *Envelope_Timer:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Timer); err != nil {
			return err
		}
	case *Envelope_Event:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Event); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Envelope.Message has unexpected type %T", x)
	}
	return nil
}

func _Envelope_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Envelope)
	switch tag {
	case 4: // message.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Log)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Log{msg}
		return true, err
	case 5: // message.counter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Counter)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Counter{msg}
		return true, err
	case 6: // message.gauge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Gauge)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Gauge{msg}
		return true, err
	case 7: // message.timer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Timer)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Timer{msg}
		return true, err
	case 10: // message.event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Event{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Envelope_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		s := proto.Size(x.Log)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Counter:
		s := proto.Size(x.Counter)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Gauge:
		s := proto.Size(x.Gauge)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Timer:
		s := proto.Size(x.Timer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Event:
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

type EnvelopeBatch struct {
	Batch                []*Envelope `protobuf:"bytes,1,rep,name=batch,proto3" json:"batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EnvelopeBatch) Reset()         { *m = EnvelopeBatch{} }
func (m *EnvelopeBatch) String() string { return proto.CompactTextString(m) }
func (*EnvelopeBatch) ProtoMessage()    {}
func (*EnvelopeBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{1}
}
func (m *EnvelopeBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvelopeBatch.Unmarshal(m, b)
}
func (m *EnvelopeBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvelopeBatch.Marshal(b, m, deterministic)
}
func (dst *EnvelopeBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvelopeBatch.Merge(dst, src)
}
func (m *EnvelopeBatch) XXX_Size() int {
	return xxx_messageInfo_EnvelopeBatch.Size(m)
}
func (m *EnvelopeBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvelopeBatch.DiscardUnknown(m)
}

var xxx_messageInfo_EnvelopeBatch proto.InternalMessageInfo

func (m *EnvelopeBatch) GetBatch() []*Envelope {
	if m != nil {
		return m.Batch
	}
	return nil
}

type Value struct {
	// Types that are valid to be assigned to Data:
	//	*Value_Text
	//	*Value_Integer
	//	*Value_Decimal
	Data                 isValue_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{2}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Data interface {
	isValue_Data()
}

type Value_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type Value_Integer struct {
	Integer int64 `protobuf:"varint,2,opt,name=integer,proto3,oneof"`
}

type Value_Decimal struct {
	Decimal float64 `protobuf:"fixed64,3,opt,name=decimal,proto3,oneof"`
}

func (*Value_Text) isValue_Data() {}

func (*Value_Integer) isValue_Data() {}

func (*Value_Decimal) isValue_Data() {}

func (m *Value) GetData() isValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Value) GetText() string {
	if x, ok := m.GetData().(*Value_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Value) GetInteger() int64 {
	if x, ok := m.GetData().(*Value_Integer); ok {
		return x.Integer
	}
	return 0
}

func (m *Value) GetDecimal() float64 {
	if x, ok := m.GetData().(*Value_Decimal); ok {
		return x.Decimal
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_Text)(nil),
		(*Value_Integer)(nil),
		(*Value_Decimal)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *Value_Integer:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Integer))
	case *Value_Decimal:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Decimal))
	case nil:
	default:
		return fmt.Errorf("Value.Data has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // data.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &Value_Text{x}
		return true, err
	case 2: // data.integer
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Value_Integer{int64(x)}
		return true, err
	case 3: // data.decimal
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Data = &Value_Decimal{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *Value_Integer:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Integer))
	case *Value_Decimal:
		n += 1 // tag and wire
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Log struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type                 Log_Type `protobuf:"varint,2,opt,name=type,proto3,enum=loggregator.v2.Log_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{3}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (dst *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(dst, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Log) GetType() Log_Type {
	if m != nil {
		return m.Type
	}
	return Log_OUT
}

type Counter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Delta                uint64   `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
	Total                uint64   `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}
func (*Counter) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{4}
}
func (m *Counter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Counter.Unmarshal(m, b)
}
func (m *Counter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Counter.Marshal(b, m, deterministic)
}
func (dst *Counter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counter.Merge(dst, src)
}
func (m *Counter) XXX_Size() int {
	return xxx_messageInfo_Counter.Size(m)
}
func (m *Counter) XXX_DiscardUnknown() {
	xxx_messageInfo_Counter.DiscardUnknown(m)
}

var xxx_messageInfo_Counter proto.InternalMessageInfo

func (m *Counter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Counter) GetDelta() uint64 {
	if m != nil {
		return m.Delta
	}
	return 0
}

func (m *Counter) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type Gauge struct {
	Metrics              map[string]*GaugeValue `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{5}
}
func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gauge.Unmarshal(m, b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
}
func (dst *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(dst, src)
}
func (m *Gauge) XXX_Size() int {
	return xxx_messageInfo_Gauge.Size(m)
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetMetrics() map[string]*GaugeValue {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type GaugeValue struct {
	Unit                 string   `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GaugeValue) Reset()         { *m = GaugeValue{} }
func (m *GaugeValue) String() string { return proto.CompactTextString(m) }
func (*GaugeValue) ProtoMessage()    {}
func (*GaugeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{6}
}
func (m *GaugeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaugeValue.Unmarshal(m, b)
}
func (m *GaugeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaugeValue.Marshal(b, m, deterministic)
}
func (dst *GaugeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaugeValue.Merge(dst, src)
}
func (m *GaugeValue) XXX_Size() int {
	return xxx_messageInfo_GaugeValue.Size(m)
}
func (m *GaugeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_GaugeValue.DiscardUnknown(m)
}

var xxx_messageInfo_GaugeValue proto.InternalMessageInfo

func (m *GaugeValue) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *GaugeValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Timer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	Stop                 int64    `protobuf:"varint,3,opt,name=stop,proto3" json:"stop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timer) Reset()         { *m = Timer{} }
func (m *Timer) String() string { return proto.CompactTextString(m) }
func (*Timer) ProtoMessage()    {}
func (*Timer) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{7}
}
func (m *Timer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timer.Unmarshal(m, b)
}
func (m *Timer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timer.Marshal(b, m, deterministic)
}
func (dst *Timer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timer.Merge(dst, src)
}
func (m *Timer) XXX_Size() int {
	return xxx_messageInfo_Timer.Size(m)
}
func (m *Timer) XXX_DiscardUnknown() {
	xxx_messageInfo_Timer.DiscardUnknown(m)
}

var xxx_messageInfo_Timer proto.InternalMessageInfo

func (m *Timer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Timer) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Timer) GetStop() int64 {
	if m != nil {
		return m.Stop
	}
	return 0
}

type Event struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_envelope_1843aa18364a6e12, []int{8}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Envelope)(nil), "loggregator.v2.Envelope")
	proto.RegisterMapType((map[string]*Value)(nil), "loggregator.v2.Envelope.DeprecatedTagsEntry")
	proto.RegisterMapType((map[string]string)(nil), "loggregator.v2.Envelope.TagsEntry")
	proto.RegisterType((*EnvelopeBatch)(nil), "loggregator.v2.EnvelopeBatch")
	proto.RegisterType((*Value)(nil), "loggregator.v2.Value")
	proto.RegisterType((*Log)(nil), "loggregator.v2.Log")
	proto.RegisterType((*Counter)(nil), "loggregator.v2.Counter")
	proto.RegisterType((*Gauge)(nil), "loggregator.v2.Gauge")
	proto.RegisterMapType((map[string]*GaugeValue)(nil), "loggregator.v2.Gauge.MetricsEntry")
	proto.RegisterType((*GaugeValue)(nil), "loggregator.v2.GaugeValue")
	proto.RegisterType((*Timer)(nil), "loggregator.v2.Timer")
	proto.RegisterType((*Event)(nil), "loggregator.v2.Event")
	proto.RegisterEnum("loggregator.v2.Log_Type", Log_Type_name, Log_Type_value)
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor_envelope_1843aa18364a6e12) }

var fileDescriptor_envelope_1843aa18364a6e12 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xd1, 0x6b, 0xd4, 0x4e,
	0x10, 0xc7, 0x2f, 0x4d, 0xd2, 0x34, 0xd3, 0xfe, 0xfa, 0x2b, 0xdb, 0x8a, 0xe1, 0x10, 0x3c, 0xf2,
	0xe2, 0x81, 0x35, 0x68, 0x0b, 0x55, 0x44, 0x10, 0x4e, 0x0f, 0x4f, 0x38, 0x11, 0x96, 0xb3, 0xf8,
	0x22, 0x65, 0x9b, 0xac, 0x6b, 0x30, 0x97, 0x0d, 0xc9, 0xe6, 0x30, 0x8f, 0xfe, 0x21, 0xfe, 0xaf,
	0xb2, 0xb3, 0x89, 0x77, 0x3d, 0x73, 0xbe, 0xcd, 0xcc, 0xf7, 0x33, 0x73, 0xb3, 0x33, 0x93, 0x83,
	0x63, 0x9e, 0xaf, 0x78, 0x26, 0x0b, 0x1e, 0x15, 0xa5, 0x54, 0x92, 0x1c, 0x67, 0x52, 0x88, 0x92,
	0x0b, 0xa6, 0x64, 0x19, 0xad, 0x2e, 0xc2, 0x9f, 0x2e, 0x1c, 0x4c, 0x5b, 0x84, 0x3c, 0x00, 0x5f,
	0xa5, 0x4b, 0x5e, 0x29, 0xb6, 0x2c, 0x02, 0x6b, 0x64, 0x8d, 0x6d, 0xba, 0x0e, 0x68, 0xb5, 0x92,
	0x75, 0x19, 0xf3, 0x9b, 0x34, 0x09, 0xf6, 0x46, 0xd6, 0xd8, 0xa7, 0xeb, 0x00, 0x19, 0xc1, 0x61,
	0x9a, 0x57, 0x8a, 0xe5, 0x46, 0x3f, 0x40, 0x7d, 0x33, 0x44, 0xae, 0xe1, 0xff, 0x84, 0x17, 0x25,
	0x8f, 0x99, 0xe2, 0xc9, 0x8d, 0x62, 0xa2, 0x0a, 0xec, 0x91, 0x3d, 0x3e, 0xbc, 0x38, 0x8f, 0xee,
	0x36, 0x15, 0x75, 0x0d, 0x45, 0x6f, 0xff, 0xf0, 0x0b, 0x26, 0xaa, 0x69, 0xae, 0xca, 0x86, 0x6e,
	0x17, 0x21, 0x57, 0xe0, 0x60, 0x31, 0x1f, 0x8b, 0x85, 0x3b, 0x8b, 0xad, 0x4b, 0x20, 0x4f, 0x1e,
	0x81, 0x9d, 0x49, 0x11, 0x38, 0x23, 0x6b, 0x7c, 0x78, 0x71, 0xba, 0x9d, 0x36, 0x97, 0x62, 0x36,
	0xa0, 0x9a, 0x20, 0x97, 0xe0, 0xc5, 0xb2, 0xce, 0x15, 0x2f, 0x03, 0x17, 0xe1, 0xfb, 0xdb, 0xf0,
	0x1b, 0x23, 0xcf, 0x06, 0xb4, 0x23, 0xc9, 0x13, 0x70, 0x05, 0xab, 0x05, 0x0f, 0xf6, 0x31, 0xe5,
	0xde, 0x76, 0xca, 0x3b, 0x2d, 0xce, 0x06, 0xd4, 0x50, 0x1a, 0xd7, 0x93, 0x2e, 0x03, 0xaf, 0x1f,
	0x5f, 0x68, 0x51, 0xe3, 0x48, 0x69, 0x9c, 0xaf, 0x78, 0xae, 0x02, 0xe8, 0xc7, 0xa7, 0x5a, 0xd4,
	0x38, 0x52, 0xc3, 0xcf, 0x70, 0xda, 0x33, 0x4a, 0x72, 0x02, 0xf6, 0x77, 0xde, 0xe0, 0xa6, 0x7d,
	0xaa, 0x4d, 0xf2, 0x18, 0xdc, 0x15, 0xcb, 0x6a, 0x8e, 0xfb, 0xed, 0xa9, 0x7b, 0xad, 0x45, 0x6a,
	0x98, 0x97, 0x7b, 0x2f, 0xac, 0xe1, 0x73, 0xf0, 0xff, 0x55, 0xef, 0x6c, 0xb3, 0x9e, 0xbf, 0x91,
	0x38, 0xf1, 0xc1, 0x5b, 0xf2, 0xaa, 0x62, 0x82, 0x87, 0xaf, 0xe1, 0xbf, 0x6e, 0x49, 0x13, 0xa6,
	0xe2, 0x6f, 0x24, 0x02, 0xf7, 0x56, 0x1b, 0x81, 0x85, 0x2b, 0x0d, 0x76, 0xad, 0x94, 0x1a, 0x2c,
	0xfc, 0x02, 0x2e, 0x36, 0x46, 0xce, 0xc0, 0x51, 0xfc, 0x87, 0x32, 0x1d, 0xcc, 0x06, 0x14, 0x3d,
	0x32, 0x04, 0x2f, 0xcd, 0x15, 0x17, 0xbc, 0xc4, 0x36, 0x6c, 0xbd, 0xa6, 0x36, 0xa0, 0xb5, 0x84,
	0xc7, 0xe9, 0x92, 0x65, 0x81, 0x3d, 0xb2, 0xc6, 0x96, 0xd6, 0xda, 0xc0, 0x64, 0x1f, 0x9c, 0x84,
	0x29, 0x16, 0x0a, 0xb0, 0xe7, 0x52, 0x90, 0x00, 0xbc, 0x82, 0x35, 0x99, 0x64, 0x09, 0xd6, 0x3f,
	0xa2, 0x9d, 0x4b, 0xce, 0xc1, 0x51, 0x4d, 0x61, 0x1e, 0x79, 0xfc, 0x77, 0xbb, 0x73, 0x29, 0xa2,
	0x45, 0x53, 0x70, 0x8a, 0x54, 0x18, 0x80, 0xa3, 0x3d, 0xe2, 0x81, 0xfd, 0xf1, 0xd3, 0xe2, 0x64,
	0xa0, 0x8d, 0x29, 0xa5, 0x27, 0x56, 0xf8, 0x1e, 0xbc, 0xf6, 0x92, 0x08, 0x01, 0x27, 0x67, 0x4b,
	0xde, 0xce, 0x12, 0x6d, 0x3d, 0xcc, 0x84, 0x67, 0x8a, 0xe1, 0xef, 0x38, 0xd4, 0x38, 0x3a, 0xaa,
	0xa4, 0x6a, 0xfb, 0x77, 0xa8, 0x71, 0xc2, 0x5f, 0x16, 0xb8, 0x78, 0x62, 0xe4, 0x95, 0x1e, 0xb4,
	0x2a, 0xd3, 0xb8, 0x6a, 0xc7, 0x19, 0xf6, 0x9e, 0x62, 0xf4, 0xc1, 0x40, 0xe6, 0x0b, 0xe9, 0x52,
	0x86, 0xd7, 0x70, 0xb4, 0x29, 0xf4, 0xac, 0xf8, 0xe9, 0xdd, 0x93, 0x19, 0xf6, 0x56, 0xdf, 0xbe,
	0x9b, 0xf0, 0x0a, 0x60, 0x2d, 0xe8, 0xd7, 0xd6, 0x79, 0xaa, 0xba, 0xd7, 0x6a, 0xfb, 0xee, 0xe9,
	0x58, 0x6d, 0x6e, 0x38, 0x05, 0x17, 0x3f, 0x85, 0x5d, 0x03, 0xaa, 0x14, 0x2b, 0x95, 0x59, 0x33,
	0x35, 0x8e, 0x26, 0x2b, 0x25, 0x0b, 0x9c, 0x8f, 0x4d, 0xd1, 0x0e, 0x9f, 0x81, 0x8b, 0x9f, 0x08,
	0x4e, 0x2f, 0x55, 0x59, 0x57, 0xc7, 0x38, 0x3a, 0xe5, 0x56, 0x26, 0x4d, 0x7b, 0xb5, 0x68, 0x4f,
	0xae, 0xe0, 0xa1, 0x2c, 0x45, 0x14, 0x67, 0xb2, 0x4e, 0xbe, 0xca, 0x3a, 0x4f, 0xca, 0x66, 0xeb,
	0xa9, 0x93, 0xd3, 0xf9, 0xda, 0xef, 0x6e, 0xf4, 0x76, 0x1f, 0xff, 0x78, 0x2f, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x0a, 0x63, 0xfb, 0xc3, 0x8a, 0x05, 0x00, 0x00,
}
