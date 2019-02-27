// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MoveTo struct {
	_tab flatbuffers.Table
}

func GetRootAsMoveTo(buf []byte, offset flatbuffers.UOffsetT) *MoveTo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MoveTo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MoveTo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MoveTo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MoveTo) X() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MoveTo) MutateX(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *MoveTo) Y() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MoveTo) MutateY(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func MoveToStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MoveToAddX(builder *flatbuffers.Builder, x uint32) {
	builder.PrependUint32Slot(0, x, 0)
}
func MoveToAddY(builder *flatbuffers.Builder, y uint32) {
	builder.PrependUint32Slot(1, y, 0)
}
func MoveToEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}