// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Uint8 struct {
	_tab flatbuffers.Table
}

func GetRootAsUint8(buf []byte, offset flatbuffers.UOffsetT) *Uint8 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Uint8{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Uint8) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Uint8) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Uint8) Value(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Uint8) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Uint8) ValueBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func Uint8Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Uint8AddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func Uint8StartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func Uint8End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
