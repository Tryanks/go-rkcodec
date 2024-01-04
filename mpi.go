package rkcodec

//#cgo pkg-config: rockchip_mpp
/*
#include <rockchip/rk_mpi.h>

MPP_RET Mpi_decode(MppApi mpi, MppCtx ctx, MppPacket packet, MppFrame *frame) {
    return mpi->decode(ctx, packet, frame);
}

MPP_RET Mpi_decode_put_packet(MppApi mpi, MppCtx ctx, MppPacket packet) {
    return mpi->decode_put_packet(ctx, packet);
}

MPP_RET Mpi_decode_get_frame(MppApi mpi, MppCtx ctx, MppFrame *frame) {
    return mpi->decode_get_frame(ctx, frame);
}

MPP_RET Mpi_encode(MppApi mpi, MppCtx ctx, MppFrame frame, MppPacket *packet) {
    return mpi->encode(ctx, frame, packet);
}

MPP_RET Mpi_encode_put_frame(MppApi mpi, MppCtx ctx, MppFrame frame) {
    return mpi->encode_put_frame(ctx, frame);
}

MPP_RET Mpi_encode_get_packet(MppApi mpi, MppCtx ctx, MppPacket *packet) {
    return mpi->encode_get_packet(ctx, packet);
}

//MPP_RET Mpi_isp(MppApi mpi, MppCtx ctx, MppFrame dst, MppFrame src) {
//    return mpi->isp(ctx, dst, src);
//}
//
//MPP_RET Mpi_isp_put_frame(MppApi mpi, MppCtx ctx, MppFrame frame) {
//    return mpi->isp_put_frame(ctx, frame);
//}
//
//MPP_RET Mpi_isp_get_frame(MppApi mpi, MppCtx ctx, MppFrame *frame) {
//    return mpi->isp_get_frame(ctx, frame);
//}
//
//MPP_RET Mpi_poll(MppApi mpi, MppCtx ctx, MppPortType type, MppPollType timeout) {
//    return mpi->poll(ctx, type, timeout);
//}
//
//MPP_RET Mpi_dequeue(MppApi mpi, MppCtx ctx, MppPortType type, MppTask *task) {
//    return mpi->dequeue(ctx, type, task);
//}
//
//MPP_RET Mpi_enqueue(MppApi mpi, MppCtx ctx, MppPortType type, MppTask task) {
//    return mpi->enqueue(ctx, type, task);
//}

MPP_RET Mpi_reset(MppApi mpi, MppCtx ctx) {
    return mpi->reset(ctx);
}

MPP_RET Mpi_control(MppApi mpi, MppCtx ctx, MpiCmd cmd, MppParam param) {
    return mpi->control(ctx, cmd, param);
}
*/
import "C"

type MppContext struct {
	mpi *C.struct_MppApi
	ctx C.MppCtx
}

func CreateCtxAndApi(ctxType MppCtxType, codingType MppCodingType) *MppContext {
	var mpi *C.struct_MppApi
	var ctx C.MppCtx
	C.mpp_create(&mpi, &ctx, ctxType)
	C.mpp_init(ctx, codingType, 0)
	return &MppContext{mpi: mpi, ctx: ctx}
}

func (c *MppContext) Decode(packet MppPacket, frame *MppFrame) {
	C.Mpi_decode(c.mpi, c.ctx, packet.c, frame.c)
}

func (c *MppContext) DecodePutPacket(packet MppPacket) {
	C.Mpi_decode_put_packet(c.mpi, c.ctx, packet.c)
}

func (c *MppContext) DecodeGetFrame(frame *MppFrame) {
	C.Mpi_decode_get_frame(c.mpi, c.ctx, frame.c)
}

func (c *MppContext) Encode(frame MppFrame, packet *MppPacket) {
	C.Mpi_encode(c.mpi, c.ctx, frame.c, packet.c)
}

func (c *MppContext) EncodePutFrame(frame MppFrame) {
	C.Mpi_encode_put_frame(c.mpi, c.ctx, frame.c)
}

func (c *MppContext) EncodeGetPacket(packet *MppPacket) {
	C.Mpi_encode_get_packet(c.mpi, c.ctx, packet.c)
}

func (c *MppContext) Reset() {
	C.Mpi_reset(c.mpi, c.ctx)
}

func (c *MppContext) Control(cmd MpiCmd, param bool) {
	param2 := C.int(0)
	if param {
		param2 = C.int(1)
	}
	C.Mpi_control(c.mpi, c.ctx, (C.enum_MpiCmd)(cmd), C.MppParam(param2))
}

func (c *MppContext) Destroy() {
	C.mpp_destroy(c.mpi, c.ctx)
}

type MpiCmd C.enum_MpiCmd

const (
	MpiCmdBase               = MpiCmd(C.MPI_CMD_BASE)
	MpiDecSetParserSplitMode = MpiCmd(C.MPI_DEC_SET_PARSER_SPLIT_MODE)
	MppDecSetImmediateOut    = MpiCmd(C.MPP_DEC_SET_IMMEDIATE_OUT)
	MppSetInputBlock         = MpiCmd(C.MPP_SET_INPUT_BLOCK)
	MppSetInputTimeout       = MpiCmd(C.MPP_SET_INPUT_TIMEOUT)
	MppSetOutputBlock        = MpiCmd(C.MPP_SET_OUTPUT_BLOCK)
	MppSetOutputTimeout      = MpiCmd(C.MPP_SET_OUTPUT_TIMEOUT)
)

type MppCtxType C.enum_MppCtxType

const (
	MppCtxTypeDecoder = MppCtxType(C.MPP_CTX_DEC)
	MppCtxTypeEncoder = MppCtxType(C.MPP_CTX_ENC)
)

type MppCodingType C.enum_MppCodingType

const (
	MppCodingTypeAVC   = MppCodingType(C.MPP_VIDEO_CodingAVC)
	MppCodingTypeHEVC  = MppCodingType(C.MPP_VIDEO_CodingHEVC)
	MppCodingTypeVP8   = MppCodingType(C.MPP_VIDEO_CodingVP8)
	MppCodingTypeVP9   = MppCodingType(C.MPP_VIDEO_CodingVP9)
	MppCodingTypeMJPEG = MppCodingType(C.MPP_VIDEO_CodingMJPEG)
)
