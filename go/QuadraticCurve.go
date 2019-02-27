// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type QuadraticCurve struct {
	_tab flatbuffers.Table
}

func GetRootAsQuadraticCurve(buf []byte, offset flatbuffers.UOffsetT) *QuadraticCurve {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &QuadraticCurve{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *QuadraticCurve) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *QuadraticCurve) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *QuadraticCurve) X1() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *QuadraticCurve) MutateX1(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *QuadraticCurve) Y1() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *QuadraticCurve) MutateY1(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *QuadraticCurve) X() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *QuadraticCurve) MutateX(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *QuadraticCurve) Y() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *QuadraticCurve) MutateY(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func QuadraticCurveStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func QuadraticCurveAddX1(builder *flatbuffers.Builder, x1 uint32) {
	builder.PrependUint32Slot(0, x1, 0)
}
func QuadraticCurveAddY1(builder *flatbuffers.Builder, y1 uint32) {
	builder.PrependUint32Slot(1, y1, 0)
}
func QuadraticCurveAddX(builder *flatbuffers.Builder, x uint32) {
	builder.PrependUint32Slot(2, x, 0)
}
func QuadraticCurveAddY(builder *flatbuffers.Builder, y uint32) {
	builder.PrependUint32Slot(3, y, 0)
}
func QuadraticCurveEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}