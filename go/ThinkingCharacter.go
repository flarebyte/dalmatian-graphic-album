// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ThinkingCharacter struct {
	_tab flatbuffers.Table
}

func GetRootAsThinkingCharacter(buf []byte, offset flatbuffers.UOffsetT) *ThinkingCharacter {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ThinkingCharacter{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ThinkingCharacter) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ThinkingCharacter) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ThinkingCharacter) EntityId() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ThinkingCharacter) MutateEntityId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func ThinkingCharacterStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ThinkingCharacterAddEntityId(builder *flatbuffers.Builder, entityId uint16) {
	builder.PrependUint16Slot(0, entityId, 0)
}
func ThinkingCharacterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
