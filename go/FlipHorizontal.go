// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlipHorizontal struct {
	_tab flatbuffers.Table
}

func GetRootAsFlipHorizontal(buf []byte, offset flatbuffers.UOffsetT) *FlipHorizontal {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FlipHorizontal{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FlipHorizontal) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FlipHorizontal) Table() flatbuffers.Table {
	return rcv._tab
}

func FlipHorizontalStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func FlipHorizontalEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
