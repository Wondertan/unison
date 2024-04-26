// Code generated by capnpc-go. DO NOT EDIT.

package batchmsg

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Batch capnp.Struct
type Batch_signature Batch

// Batch_TypeID is the unique identifier for the type Batch.
const Batch_TypeID = 0xa4e968f5be59ce03

func NewBatch(s *capnp.Segment) (Batch, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Batch(st), err
}

func NewRootBatch(s *capnp.Segment) (Batch, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Batch(st), err
}

func ReadRootBatch(msg *capnp.Message) (Batch, error) {
	root, err := msg.Root()
	return Batch(root.Struct()), err
}

func (s Batch) String() string {
	str, _ := text.Marshal(0xa4e968f5be59ce03, capnp.Struct(s))
	return str
}

func (s Batch) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Batch) DecodeFromPtr(p capnp.Ptr) Batch {
	return Batch(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Batch) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Batch) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Batch) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Batch) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Batch) Data() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s Batch) HasData() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Batch) SetData(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

func (s Batch) Signature() Batch_signature { return Batch_signature(s) }

func (s Batch_signature) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Batch_signature) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Batch_signature) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Batch_signature) Signer() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return []byte(p.Data()), err
}

func (s Batch_signature) HasSigner() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Batch_signature) SetSigner(v []byte) error {
	return capnp.Struct(s).SetData(1, v)
}

func (s Batch_signature) Signature() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return []byte(p.Data()), err
}

func (s Batch_signature) HasSignature() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Batch_signature) SetSignature(v []byte) error {
	return capnp.Struct(s).SetData(2, v)
}

// Batch_List is a list of Batch.
type Batch_List = capnp.StructList[Batch]

// NewBatch creates a new list of Batch.
func NewBatch_List(s *capnp.Segment, sz int32) (Batch_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return capnp.StructList[Batch](l), err
}

// Batch_Future is a wrapper for a Batch promised by a client call.
type Batch_Future struct{ *capnp.Future }

func (f Batch_Future) Struct() (Batch, error) {
	p, err := f.Future.Ptr()
	return Batch(p.Struct()), err
}
func (p Batch_Future) Signature() Batch_signature_Future { return Batch_signature_Future{p.Future} }

// Batch_signature_Future is a wrapper for a Batch_signature promised by a client call.
type Batch_signature_Future struct{ *capnp.Future }

func (f Batch_signature_Future) Struct() (Batch_signature, error) {
	p, err := f.Future.Ptr()
	return Batch_signature(p.Struct()), err
}

const schema_ae6a3c85d527139f = "x\xda\x12\xe8q`1\xe4\x15gb`\x0a\x94`e" +
	"\xfb\x9f\x9d\xd3\xa5w\xf7\xea\xd6\x0e\x06A\x05\xc6\xff\xcc" +
	"\xe7\"\xf7}\xcdx\xb9\x84A\x94\x99\x9d\x91\x81\xc1\xf0" +
	"c\x10#\x03\xa3\xe0\xdfr\x06$)A)\xc6\xff\xf3" +
	"\x85\xd5\xaf\xb6\xdad\xadc`efg`0\x8ee" +
	"db\x14\xceddg`\x10Ne,g\xd0\xfd\x9f" +
	"\x94X\x90\xa3\x9f\x94X\xc2\x92\x9c\x91[\x9c\x0eb%" +
	"g\xe8%'\x16\xe4\x15X9\x81\xd9\xc5\x99\xe9y\x89" +
	"%\xa5E\x8c\xa9\x81\x1c\xcc,\x02\x8c\"\x8c\x8c\x0c\x0c" +
	"\x82\x9aV\x0c\x0c\x81*\xcc\x8c\x81\x06L\x8c\x82\x8cL" +
	"\"\x8cL\x0c\x0c\x82\xbaA\x0c\x0c\x81:\xcc\x8c\x81\x16" +
	"L\x8c\xf6 }\xa9E\x8c\xbc\x0cL\x8c\xbc\x0c\x8c\xff" +
	"a\xc600\xa6\xc2\xc5`v3\xe3\xb0\x9b!\x80\x91" +
	"\x11d)\x03\x03\x0b\xd8N-\x84\x9d \x80\x14 \xba" +
	"A\x0cL\xfc)\x89%\x89\xd8\xec\x03\x04\x00\x00\xff\xff" +
	"\xac\xc0T\xbb"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_ae6a3c85d527139f,
		Nodes: []uint64{
			0x88b5d5dd2e8a6c6b,
			0xa4e968f5be59ce03,
		},
		Compressed: true,
	})
}