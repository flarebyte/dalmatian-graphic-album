// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Vector2dFloat struct {
	_tab flatbuffers.Table
}

func GetRootAsVector2dFloat(buf []byte, offset flatbuffers.UOffsetT) *Vector2dFloat {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Vector2dFloat{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Vector2dFloat) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vector2dFloat) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Vector2dFloat) Value(obj *Vec2dFloat) *Vec2dFloat {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Vec2dFloat)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func Vector2dFloatStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Vector2dFloatAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func Vector2dFloatEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}