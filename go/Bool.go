// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Bool struct {
	_tab flatbuffers.Table
}

func GetRootAsBool(buf []byte, offset flatbuffers.UOffsetT) *Bool {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Bool{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Bool) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Bool) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Bool) Value() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Bool) MutateValue(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func BoolStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BoolAddValue(builder *flatbuffers.Builder, value byte) {
	builder.PrependByteSlot(0, value, 0)
}
func BoolEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
