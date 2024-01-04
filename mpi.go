package rkcodec

//#cgo pkg-config: rockchip_mpp
//##include <rockchip/rk_mpi.h>
import "C"

type MppFrame struct {
	c *C.struct_MppFrame
}

type MppCtx struct {
	c *C.struct_MppCtx
}

type MppApi struct {
	c *C.struct_MppApi
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
