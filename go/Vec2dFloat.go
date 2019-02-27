// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Vec2dFloat struct {
	_tab flatbuffers.Table
}

func GetRootAsVec2dFloat(buf []byte, offset flatbuffers.UOffsetT) *Vec2dFloat {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Vec2dFloat{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Vec2dFloat) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec2dFloat) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Vec2dFloat) X() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Vec2dFloat) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *Vec2dFloat) Y() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Vec2dFloat) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func Vec2dFloatStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Vec2dFloatAddX(builder *flatbuffers.Builder, x float32) {
	builder.PrependFloat32Slot(0, x, 0.0)
}
func Vec2dFloatAddY(builder *flatbuffers.Builder, y float32) {
	builder.PrependFloat32Slot(1, y, 0.0)
}
func Vec2dFloatEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
