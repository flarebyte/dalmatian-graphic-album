// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LocalizedTextId struct {
	_tab flatbuffers.Table
}

func GetRootAsLocalizedTextId(buf []byte, offset flatbuffers.UOffsetT) *LocalizedTextId {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LocalizedTextId{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LocalizedTextId) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LocalizedTextId) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LocalizedTextId) LocalizedId() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LocalizedTextId) MutateLocalizedId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func LocalizedTextIdStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func LocalizedTextIdAddLocalizedId(builder *flatbuffers.Builder, localizedId uint16) {
	builder.PrependUint16Slot(0, localizedId, 0)
}
func LocalizedTextIdEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
