// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Vec3d struct {
	_tab flatbuffers.Table
}

func GetRootAsVec3d(buf []byte, offset flatbuffers.UOffsetT) *Vec3d {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Vec3d{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Vec3d) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec3d) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Vec3d) X() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Vec3d) MutateX(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Vec3d) Y() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Vec3d) MutateY(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Vec3d) Z() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Vec3d) MutateZ(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func Vec3dStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func Vec3dAddX(builder *flatbuffers.Builder, x uint32) {
	builder.PrependUint32Slot(0, x, 0)
}
func Vec3dAddY(builder *flatbuffers.Builder, y uint32) {
	builder.PrependUint32Slot(1, y, 0)
}
func Vec3dAddZ(builder *flatbuffers.Builder, z uint32) {
	builder.PrependUint32Slot(2, z, 0)
}
func Vec3dEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}