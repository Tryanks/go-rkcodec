package rkcodec

//#include <rockchip/rk_type.h>
import "C"

type cU8 = C.RK_U8
type cU16 = C.RK_U16
type cU32 = C.RK_U32
type cULONG = C.RK_ULONG
type cU64 = C.RK_U64

type cS8 = C.RK_S8
type cS16 = C.RK_S16
type cS32 = C.RK_S32
type cLONG = C.RK_LONG
type cS64 = C.RK_S64

// type MppCtxType = C.enum_MppCtxType
type MppCtxType = C.int

const (
	MppCtxDec  = MppCtxType(C.MPP_CTX_DEC)  // decoder
	MppCtxEnc  = MppCtxType(C.MPP_CTX_ENC)  // encoder
	MppCtxIsp  = MppCtxType(C.MPP_CTX_ISP)  // isp
	MppCtxButt = MppCtxType(C.MPP_CTX_BUTT) // undefined
)

// type MppCodingType = C.enum_MppCodingType
type MppCodingType = C.int

const (
	MppCodingUnused            = MppCodingType(C.MPP_VIDEO_CodingUnused)     // Value when coding is N/A
	MppCodingAutoDetect        = MppCodingType(C.MPP_VIDEO_CodingAutoDetect) // Autodetection of coding type
	MppCodingMPEG2             = MppCodingType(C.MPP_VIDEO_CodingMPEG2)      // AKA: H.262
	MppCodingH263              = MppCodingType(C.MPP_VIDEO_CodingH263)       // H.263
	MppCodingMPEG4             = MppCodingType(C.MPP_VIDEO_CodingMPEG4)      // MPEG-4
	MppCodingWMV               = MppCodingType(C.MPP_VIDEO_CodingWMV)        // Windows Media Video (WMV1,WMV2,WMV3)
	MppCodingRV                = MppCodingType(C.MPP_VIDEO_CodingRV)         // all versions of Real Video
	MppCodingAVC               = MppCodingType(C.MPP_VIDEO_CodingAVC)        // H.264/AVC
	MppCodingMJPEG             = MppCodingType(C.MPP_VIDEO_CodingMJPEG)      // Motion JPEG
	MppCodingVP8               = MppCodingType(C.MPP_VIDEO_CodingVP8)        // VP8
	MppCodingVP9               = MppCodingType(C.MPP_VIDEO_CodingVP9)        // VP9
	MppCodingVC1               = MppCodingType(C.MPP_VIDEO_CodingVC1)        // Windows Media Video (WMV1,WMV2,WMV3)
	MppCodingFLV1              = MppCodingType(C.MPP_VIDEO_CodingFLV1)       // Sorenson H.263
	MppCodingDIVX3             = MppCodingType(C.MPP_VIDEO_CodingDIVX3)      // DIVX3
	MppCodingVP6               = MppCodingType(C.MPP_VIDEO_CodingVP6)
	MppCodingHEVC              = MppCodingType(C.MPP_VIDEO_CodingHEVC)              // H.265/HEVC
	MppCodingAVSPLUS           = MppCodingType(C.MPP_VIDEO_CodingAVSPLUS)           // AVS+
	MppCodingAVS               = MppCodingType(C.MPP_VIDEO_CodingAVS)               // AVS profile=0x20
	MppCodingAVS2              = MppCodingType(C.MPP_VIDEO_CodingAVS2)              // AVS2
	MppCodingAV1               = MppCodingType(C.MPP_VIDEO_CodingAV1)               // av1
	MppCodingKhronosExtensions = MppCodingType(C.MPP_VIDEO_CodingKhronosExtensions) // Reserved region for introducing Khronos Standard Extensions
	MppCodingVendorStartUnused = MppCodingType(C.MPP_VIDEO_CodingVendorStartUnused) // Reserved region for introducing Vendor Extensions
	MppCodingMax               = MppCodingType(C.MPP_VIDEO_CodingMax)
)

type MppCtx struct {
	c *C.MppCtx
}
type MppParam = *C.MppParam

type MppFrame struct {
	c *C.MppFrame
}
type MppPacket struct {
	c *C.MppPacket
}

type MppBuffer struct {
	c C.MppBuffer
}
type MppBufferGroup struct {
	c *C.MppBufferGroup
}

type MppTask struct {
	c *C.MppTask
}
type MppMeta struct {
	c C.MppMeta
}
