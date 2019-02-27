// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Shape struct {
	_tab flatbuffers.Table
}

func GetRootAsShape(buf []byte, offset flatbuffers.UOffsetT) *Shape {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Shape{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Shape) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Shape) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Shape) Metadata(obj *Metadata) *Metadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Metadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Shape) Drawing(obj *DrawItem, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Shape) DrawingLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Shape) EntityId() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Shape) MutateEntityId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func (rcv *Shape) IllustrationId() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Shape) MutateIllustrationId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func ShapeStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ShapeAddMetadata(builder *flatbuffers.Builder, metadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(metadata), 0)
}
func ShapeAddDrawing(builder *flatbuffers.Builder, drawing flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(drawing), 0)
}
func ShapeStartDrawingVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ShapeAddEntityId(builder *flatbuffers.Builder, entityId uint16) {
	builder.PrependUint16Slot(2, entityId, 0)
}
func ShapeAddIllustrationId(builder *flatbuffers.Builder, illustrationId uint16) {
	builder.PrependUint16Slot(3, illustrationId, 0)
}
func ShapeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
