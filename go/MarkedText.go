// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MarkedText struct {
	_tab flatbuffers.Table
}

func GetRootAsMarkedText(buf []byte, offset flatbuffers.UOffsetT) *MarkedText {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MarkedText{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MarkedText) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MarkedText) Table() flatbuffers.Table {
	return rcv._tab
}

func MarkedTextStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func MarkedTextEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
