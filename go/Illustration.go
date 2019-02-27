// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Illustration struct {
	_tab flatbuffers.Table
}

func GetRootAsIllustration(buf []byte, offset flatbuffers.UOffsetT) *Illustration {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Illustration{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Illustration) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Illustration) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Illustration) Metadata(obj *Metadata) *Metadata {
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

func (rcv *Illustration) DataType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Illustration) MutateDataType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *Illustration) Data(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func IllustrationStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func IllustrationAddMetadata(builder *flatbuffers.Builder, metadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(metadata), 0)
}
func IllustrationAddDataType(builder *flatbuffers.Builder, dataType byte) {
	builder.PrependByteSlot(1, dataType, 0)
}
func IllustrationAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(data), 0)
}
func IllustrationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
