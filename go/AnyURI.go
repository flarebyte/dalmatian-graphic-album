// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AnyURI struct {
	_tab flatbuffers.Table
}

func GetRootAsAnyURI(buf []byte, offset flatbuffers.UOffsetT) *AnyURI {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AnyURI{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AnyURI) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AnyURI) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AnyURI) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AnyURI) MediaType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AnyURIStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AnyURIAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func AnyURIAddMediaType(builder *flatbuffers.Builder, mediaType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(mediaType), 0)
}
func AnyURIEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
