// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Tile struct {
	_tab flatbuffers.Table
}

func GetRootAsTile(buf []byte, offset flatbuffers.UOffsetT) *Tile {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Tile{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Tile) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Tile) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Tile) SupplementId() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tile) MutateSupplementId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func TileStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TileAddSupplementId(builder *flatbuffers.Builder, supplementId uint16) {
	builder.PrependUint16Slot(0, supplementId, 0)
}
func TileEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}