// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Page struct {
	_tab flatbuffers.Table
}

func GetRootAsPage(buf []byte, offset flatbuffers.UOffsetT) *Page {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Page{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Page) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Page) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Page) SupplementId() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Page) MutateSupplementId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func PageStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func PageAddSupplementId(builder *flatbuffers.Builder, supplementId uint16) {
	builder.PrependUint16Slot(0, supplementId, 0)
}
func PageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
