package rkcodec

/*
#include <rockchip/mpp_packet.h>

RK_S32 MppPktSeg_getType(const MppPktSeg seg) {
    return seg.type;
}

void MppPktSeg_setType(MppPktSeg seg, RK_S32 type) {
    seg.type = type;
}
*/
import "C"

func NewMppPacket() (*MppPacket, MppRet) {
	cPacket := C.MppPacket(nil)
	ret := MppRet(C.mpp_packet_new(&cPacket))
	return &MppPacket{c: &cPacket}, ret
}

func (packet *MppPacket) Init(data []byte, size int64) MppRet {
	return MppRet(C.mpp_packet_init(packet.c, C.CBytes(data), C.size_t(size)))
}

func (packet *MppPacket) InitWithBuffer(buffer *MppBuffer) MppRet {
	return MppRet(C.mpp_packet_init_with_buffer(packet.c, buffer.c))
}

func (packet *MppPacket) CopyInit(src *MppPacket) MppRet {
	return MppRet(C.mpp_packet_copy_init(packet.c, *src.c))
}

func (packet *MppPacket) Deinit() MppRet {
	return MppRet(C.mpp_packet_deinit(packet.c))
}

func (packet *MppPacket) GetData() []byte {
	return C.GoBytes(C.mpp_packet_get_data(*packet.c), C.int(C.mpp_packet_get_size(*packet.c)))
}

func (packet *MppPacket) SetData(data []byte) {
	C.mpp_packet_set_data(*packet.c, C.CBytes(data))
	C.mpp_packet_set_size(*packet.c, C.size_t(len(data)))
}

func (packet *MppPacket) SetPts(pts int64) {
	C.mpp_packet_set_pts(*packet.c, C.RK_S64(pts))
}

func (packet *MppPacket) GetPts() int64 {
	return int64(C.mpp_packet_get_pts(*packet.c))
}

func (packet *MppPacket) SetDts(dts int64) {
	C.mpp_packet_set_dts(*packet.c, C.RK_S64(dts))
}

func (packet *MppPacket) GetDts() int64 {
	return int64(C.mpp_packet_get_dts(*packet.c))
}

func (packet *MppPacket) SetFlag(flag uint32) {
	C.mpp_packet_set_flag(*packet.c, C.RK_U32(flag))
}

func (packet *MppPacket) GetFlag() uint32 {
	return uint32(C.mpp_packet_get_flag(*packet.c))
}

func (packet *MppPacket) SetEos() MppRet {
	return MppRet(C.mpp_packet_set_eos(*packet.c))
}

func (packet *MppPacket) ClrEos() MppRet {
	return MppRet(C.mpp_packet_clr_eos(*packet.c))
}

func (packet *MppPacket) GetEos() uint32 {
	return uint32(C.mpp_packet_get_eos(*packet.c))
}

func (packet *MppPacket) SetExtraData() MppRet {
	return MppRet(C.mpp_packet_set_extra_data(*packet.c))
}

func (packet *MppPacket) SetBuffer(buffer *MppBuffer) {
	C.mpp_packet_set_buffer(*packet.c, buffer.c)
}

func (packet *MppPacket) GetBuffer() *MppBuffer {
	return &MppBuffer{c: C.mpp_packet_get_buffer(*packet.c)}
}

func (packet *MppPacket) Read(offset int64, data []byte) MppRet {
	return MppRet(C.mpp_packet_read(*packet.c, C.size_t(offset), C.CBytes(data), C.size_t(len(data))))
}

func (packet *MppPacket) Write(offset int64, data []byte) MppRet {
	return MppRet(C.mpp_packet_write(*packet.c, C.size_t(offset), C.CBytes(data), C.size_t(len(data))))
}

func (packet *MppPacket) HasMeta() int32 {
	return int32(C.mpp_packet_has_meta(*packet.c))
}

func (packet *MppPacket) GetMeta() *MppMeta {
	return &MppMeta{c: C.mpp_packet_get_meta(*packet.c)}
}

func (packet *MppPacket) IsPartition() uint32 {
	return uint32(C.mpp_packet_is_partition(*packet.c))
}

func (packet *MppPacket) IsSoi() uint32 {
	return uint32(C.mpp_packet_is_soi(*packet.c))
}

func (packet *MppPacket) IsEoi() uint32 {
	return uint32(C.mpp_packet_is_eoi(*packet.c))
}

type MppPktSeg struct {
	c C.struct_MppPktSeg_t
}

func (seg *MppPktSeg) GetIndex() int32 {
	return int32(seg.c.index)
}

func (seg *MppPktSeg) SetIndex(index int32) {
	seg.c.index = cS32(index)
}

func (seg *MppPktSeg) GetType() int32 {
	return int32(C.MppPktSeg_getType(seg.c))
}

func (seg *MppPktSeg) SetType(t int32) {
	C.MppPktSeg_setType(seg.c, cS32(t))
}

func (seg *MppPktSeg) GetOffset() uint32 {
	return uint32(seg.c.offset)
}

func (seg *MppPktSeg) SetOffset(offset uint32) {
	seg.c.offset = cU32(offset)
}

func (seg *MppPktSeg) GetLength() uint32 {
	return uint32(seg.c.len)
}

func (seg *MppPktSeg) SetLength(length uint32) {
	seg.c.len = cU32(length)
}

func (seg *MppPktSeg) Next() *MppPktSeg {
	return &MppPktSeg{c: *seg.c.next}
}

// func (packet *MppPacket) GetSengmentNb() uint32 {
// 	return uint32(C.mpp_packet_get_sengment_nb(*packet.c))
// }

// func (packet *MppPacket) GetSengment(index uint32) *MppPktSeg {
// 	return &MppPktSeg{c: C.mpp_packet_get_sengment(*packet.c, C.RK_U32(index))}
// }
