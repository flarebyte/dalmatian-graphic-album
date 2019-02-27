// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TileIntableionItem struct {
	_tab flatbuffers.Table
}

func GetRootAsTileIntableionItem(buf []byte, offset flatbuffers.UOffsetT) *TileIntableionItem {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TileIntableionItem{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TileIntableionItem) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TileIntableionItem) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TileIntableionItem) ItemType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TileIntableionItem) MutateItemType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *TileIntableionItem) Item(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func TileIntableionItemStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func TileIntableionItemAddItemType(builder *flatbuffers.Builder, itemType byte) {
	builder.PrependByteSlot(0, itemType, 0)
}
func TileIntableionItemAddItem(builder *flatbuffers.Builder, item flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(item), 0)
}
func TileIntableionItemEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}