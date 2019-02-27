// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IsoDateTime struct {
	_tab flatbuffers.Table
}

func GetRootAsIsoDateTime(buf []byte, offset flatbuffers.UOffsetT) *IsoDateTime {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IsoDateTime{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *IsoDateTime) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IsoDateTime) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IsoDateTime) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func IsoDateTimeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func IsoDateTimeAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func IsoDateTimeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
