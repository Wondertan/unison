// Code generated by capnpc-go. DO NOT EDIT.

package block

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Block capnp.Struct

// Block_TypeID is the unique identifier for the type Block.
const Block_TypeID = 0x92df2bd885bf4d5e

func NewBlock(s *capnp.Segment) (Block, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Block(st), err
}

func NewRootBlock(s *capnp.Segment) (Block, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Block(st), err
}

func ReadRootBlock(msg *capnp.Message) (Block, error) {
	root, err := msg.Root()
	return Block(root.Struct()), err
}

func (s Block) String() string {
	str, _ := text.Marshal(0x92df2bd885bf4d5e, capnp.Struct(s))
	return str
}

func (s Block) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Block) DecodeFromPtr(p capnp.Ptr) Block {
	return Block(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Block) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Block) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Block) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Block) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Block) Round() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Block) SetRound(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Block) Signer() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s Block) HasSigner() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Block) SetSigner(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

func (s Block) Batches() (capnp.DataList, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return capnp.DataList(p.List()), err
}

func (s Block) HasBatches() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Block) SetBatches(v capnp.DataList) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewBatches sets the batches field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Block) NewBatches(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}
func (s Block) Parents() (capnp.DataList, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return capnp.DataList(p.List()), err
}

func (s Block) HasParents() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Block) SetParents(v capnp.DataList) error {
	return capnp.Struct(s).SetPtr(2, v.ToPtr())
}

// NewParents sets the parents field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Block) NewParents(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = capnp.Struct(s).SetPtr(2, l.ToPtr())
	return l, err
}

// Block_List is a list of Block.
type Block_List = capnp.StructList[Block]

// NewBlock creates a new list of Block.
func NewBlock_List(s *capnp.Segment, sz int32) (Block_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return capnp.StructList[Block](l), err
}

// Block_Future is a wrapper for a Block promised by a client call.
type Block_Future struct{ *capnp.Future }

func (f Block_Future) Struct() (Block, error) {
	p, err := f.Future.Ptr()
	return Block(p.Struct()), err
}

type BlockID capnp.Struct

// BlockID_TypeID is the unique identifier for the type BlockID.
const BlockID_TypeID = 0xe8be3e5a06a33e53

func NewBlockID(s *capnp.Segment) (BlockID, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return BlockID(st), err
}

func NewRootBlockID(s *capnp.Segment) (BlockID, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return BlockID(st), err
}

func ReadRootBlockID(msg *capnp.Message) (BlockID, error) {
	root, err := msg.Root()
	return BlockID(root.Struct()), err
}

func (s BlockID) String() string {
	str, _ := text.Marshal(0xe8be3e5a06a33e53, capnp.Struct(s))
	return str
}

func (s BlockID) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (BlockID) DecodeFromPtr(p capnp.Ptr) BlockID {
	return BlockID(capnp.Struct{}.DecodeFromPtr(p))
}

func (s BlockID) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s BlockID) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s BlockID) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s BlockID) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s BlockID) Round() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s BlockID) SetRound(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s BlockID) Signer() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s BlockID) HasSigner() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s BlockID) SetSigner(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

func (s BlockID) Hash() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return []byte(p.Data()), err
}

func (s BlockID) HasHash() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s BlockID) SetHash(v []byte) error {
	return capnp.Struct(s).SetData(1, v)
}

// BlockID_List is a list of BlockID.
type BlockID_List = capnp.StructList[BlockID]

// NewBlockID creates a new list of BlockID.
func NewBlockID_List(s *capnp.Segment, sz int32) (BlockID_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[BlockID](l), err
}

// BlockID_Future is a wrapper for a BlockID promised by a client call.
type BlockID_Future struct{ *capnp.Future }

func (f BlockID_Future) Struct() (BlockID, error) {
	p, err := f.Future.Ptr()
	return BlockID(p.Struct()), err
}

const schema_ebe99359e631e3a9 = "x\xda\x84\xcf\xcdJ:Q\x18\xc7\xf1\xdf\xef\x1c_\x10" +
	"\xfe\xfawP\xaa\x9d\xfb\xa2\xc4\x96.L\xc4 !\xc1" +
	"Cm\x8a\x88\xc6qpD\x99\x19F\x836a\x9bn" +
	"\xa0\xba\x83V\xad\xba\x83Zu\x0f\xed\x82\x82^\x88\x16" +
	"\x05-\\\xc4\xc4\x09zA\x88V\xe7<_\x9e\xc5\xf3" +
	"I_\x94#\x85\xe4\xa4\x80PS\xd1X\xb8Y?\xdf" +
	"\xbf\x9c\xb9:\x80\x9a \xc3\x93\x9b\xc2\xed\xda\xe1\xc3#" +
	"\xa22\x0e\x18\xcf\xd7\xc6\x9b~Gw`\xb8R:\x8e" +
	"\xad\x97\xce\xee\xc77E\x1c\xc8<\xf1%3\xa2\xfe\xbd" +
	"\xf2\x14\xb3a\xcbl\xe7\xfd\xc0\x1b\x08/\xdf\xecyV" +
	"w\xce2}\xd7/Vzq\xcf\xea6H\x95\x96\x11" +
	" B\xc00\xe7\x01\xb5!\xa9\x1cA2K\xdd\xec\"" +
	"\xa0\xb6$UO\xd0\x10\xccR\x00F\xa7\x02\xa8\x96\xa4" +
	"\xda\x134\xa4\xc8R\x02\xc6\xae\x8e;\x92\xeaH0\x17" +
	"x\xdbn\x8b\x09\x08&\xc0\x85~\xa7\xed\xda\x01\x93\x10" +
	"L\x82\xc3\xa69\xb0\x1c\xbb\xcf\x14\xd8\x90\xfc\xc8)p" +
	"\xe8\x9b\x81\xed\x0e\xc6\xf3\xef\x80\x9cgukUM\xf8" +
	"\xf7EX\xd4\x84\xb2\xa4Z\xfe&\xd44\xa1*\xa9\x1a" +
	"?\x08\xf5i@-I\xaa\xd5?\xae\xfd\xef\x98}\xe7" +
	"sx\x0f\x00\x00\xff\xff&V_\x1f"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_ebe99359e631e3a9,
		Nodes: []uint64{
			0x92df2bd885bf4d5e,
			0xe8be3e5a06a33e53,
		},
		Compressed: true,
	})
}