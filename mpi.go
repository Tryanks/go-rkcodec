package rkcodec

//#include <rockchip/rk_mpi.h>
//#include "codec.h"
import "C"
import "unsafe"

type Codec struct {
	c *C.struct_MppCodec
}

func NewMppCodec() *Codec {
	return &Codec{c: C.mpp_ctx_api_alloc()}
}

func (c *Codec) Init(t MppCtxType, coding MppCodingType) MppRet {
	return MppRet(C.codec_init(c.c, C.MppCtxType(t), C.MppCodingType(coding)))
}

func (c *Codec) Destroy() MppRet {
	return MppRet(C.codec_destroy(c.c))
}

func (c *Codec) Size() uint32 {
	return uint32(C.codec_size(c.c))
}

func (c *Codec) Version() uint32 {
	return uint32(C.codec_version(c.c))
}

func (c *Codec) Decode(packet MppPacket, frame *MppFrame) MppRet {
	return MppRet(C.codec_decode(c.c, *packet.c, frame.c))
}

func (c *Codec) DecodePutPacket(packet MppPacket) MppRet {
	return MppRet(C.codec_decode_put_packet(c.c, *packet.c))
}

func (c *Codec) DecodeGetFrame(frame *MppFrame) MppRet {
	return MppRet(C.codec_decode_get_frame(c.c, frame.c))
}

func (c *Codec) Encode(frame MppFrame, packet *MppPacket) MppRet {
	return MppRet(C.codec_encode(c.c, *frame.c, packet.c))
}

func (c *Codec) EncodePutFrame(frame MppFrame) MppRet {
	return MppRet(C.codec_encode_put_frame(c.c, *frame.c))
}

func (c *Codec) EncodeGetPacket(packet *MppPacket) MppRet {
	return MppRet(C.codec_encode_get_packet(c.c, packet.c))
}

func (c *Codec) Isp(dst, src MppFrame) MppRet {
	return MppRet(C.codec_isp(c.c, *dst.c, *src.c))
}

func (c *Codec) IspPutFrame(frame MppFrame) MppRet {
	return MppRet(C.codec_isp_put_frame(c.c, *frame.c))
}

func (c *Codec) IspGetFrame(frame *MppFrame) MppRet {
	return MppRet(C.codec_isp_get_frame(c.c, frame.c))
}

func (c *Codec) Poll(t MppPortType, timeout MppPollType) {
	C.codec_poll(c.c, C.MppPortType(t), C.MppPollType(timeout))
}

func (c *Codec) Dequeue(t MppPortType, task *MppTask) MppRet {
	return MppRet(C.codec_dequeue(c.c, C.MppPortType(t), task.c))
}

func (c *Codec) Enqueue(t MppPortType, task MppTask) MppRet {
	return MppRet(C.codec_enqueue(c.c, C.MppPortType(t), *task.c))
}

func (c *Codec) Reset() MppRet {
	return MppRet(C.codec_reset(c.c))
}

func (c *Codec) Control(cmd MpiCmd, param any) MppRet {
	point := unsafe.Pointer(&param)
	return MppRet(C.codec_control(c.c, C.MpiCmd(cmd), point))
}

// CheckSupportFormat Reutnr 0 for support, -1 for unsupport
func CheckSupportFormat(t MppCtxType, coding MppCodingType) MppRet {
	return MppRet(C.mpp_check_support_format(C.MppCtxType(t), C.MppCodingType(coding)))
}

// ShowSupportFormat No return value, It just print format info to standard output
func ShowSupportFormat() {
	C.mpp_show_support_format()
}

// ShowColorFormat No return value, It just print format info to standard output
func ShowColorFormat() {
	C.mpp_show_color_format()
}
