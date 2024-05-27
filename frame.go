package rkcodec

//#include <rockchip/mpp_frame.h>
import "C"

// bit definition for mode flag in MppFrame
const MppFrameFlagFrame = C.MPP_FRAME_FLAG_FRAME                     // progressive frame
const MppFrameFlagTopField = C.MPP_FRAME_FLAG_TOP_FIELD              // top field only
const MppFrameFlagBotField = C.MPP_FRAME_FLAG_BOT_FIELD              // bottom field only
const MppFrameFlagPairedField = C.MPP_FRAME_FLAG_PAIRED_FIELD        // paired field
const MppFrameFlagTopFirst = C.MPP_FRAME_FLAG_TOP_FIRST              // paired field with field order of top first
const MppFrameFlagBotFirst = C.MPP_FRAME_FLAG_BOT_FIRST              // paired field with field order of bottom first
const MppFrameFlagDeinterlaced = C.MPP_FRAME_FLAG_DEINTERLACED       // paired field with unknown field order (MBAFF)
const MppFrameFlagFieldOrderMask = C.MPP_FRAME_FLAG_FIELD_ORDER_MASK // field order mask
const MppFrameFlagViewIdMask = C.MPP_FRAME_FLAG_VIEW_ID_MASK         // view id mask
const MppFrameFlagIepDeiMask = C.MPP_FRAME_FLAG_IEP_DEI_MASK         // iep dei mask
const MppFrameFlagIepDeiI2O1 = C.MPP_FRAME_FLAG_IEP_DEI_I2O1         // iep dei i2o1
const MppFrameFlagIepDeiI4O2 = C.MPP_FRAME_FLAG_IEP_DEI_I4O2         // iep dei i4o2
const MppFrameFlagIepDeiI4O1 = C.MPP_FRAME_FLAG_IEP_DEI_I4O1         // iep dei i4o1

// MPEG vs JPEG YUV range
type MppFrameColorRange = C.int

const (
	MppFrameRangeUnspecified = MppFrameColorRange(C.MPP_FRAME_RANGE_UNSPECIFIED)
	MppFrameRangeMpeg        = MppFrameColorRange(C.MPP_FRAME_RANGE_MPEG)
	MppFrameRangeJpeg        = MppFrameColorRange(C.MPP_FRAME_RANGE_JPEG)
	MppFrameRangeNB          = MppFrameColorRange(C.MPP_FRAME_RANGE_NB)
)

type MppFrameChromaDownSampleMode = C.int

const (
	MppFrameChromaDownSampleModeNone    = MppFrameChromaDownSampleMode(C.MPP_FRAME_CHROMA_DOWN_SAMPLE_MODE_NONE)
	MppFrameChromaDownSampleModeAverage = MppFrameChromaDownSampleMode(C.MPP_FRAME_CHORMA_DOWN_SAMPLE_MODE_AVERAGE)
	MppFrameChromaDownSampleModeDiscard = MppFrameChromaDownSampleMode(C.MPP_FRAME_CHORMA_DOWN_SAMPLE_MODE_DISCARD)
)

type MppFrameVideoFormat = C.int

const (
	MppFrameVideoFmtComponent   = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_COMPONEMT)
	MppFrameVideoFmtPal         = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_PAL)
	MppFrameVideoFmtNtsc        = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_NTSC)
	MppFrameVideoFmtSecam       = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_SECAM)
	MppFrameVideoFmtMac         = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_MAC)
	MppFrameVideoFmtUnspecified = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_UNSPECIFIED)
	MppFrameVideoFmtReserved0   = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_RESERVED0)
	MppFrameVideoFmtReserved1   = MppFrameVideoFormat(C.MPP_FRAME_VIDEO_FMT_RESERVED1)
)

// Chromaticity coordinates of the source primaries.
type MppFrameColorPrimaries = C.MppFrameColorPrimaries

const (
	MppFramePrimariesReserved0   = MppFrameColorPrimaries(C.MPP_FRAME_PRI_RESERVED0)
	MppFramePrimariesBt709       = MppFrameColorPrimaries(C.MPP_FRAME_PRI_BT709) // also ITU-R BT1361 / IEC 61966-2-4 / SMPTE RP177 Annex B
	MppFramePrimariesUnspecified = MppFrameColorPrimaries(C.MPP_FRAME_PRI_UNSPECIFIED)
	MppFramePrimariesReserved    = MppFrameColorPrimaries(C.MPP_FRAME_PRI_RESERVED)
	MppFramePrimariesBt470M      = MppFrameColorPrimaries(C.MPP_FRAME_PRI_BT470M) // also FCC Title 47 Code of Federal Regulations 73.682 (a)(20)

	MppFramePrimariesBt470BG    = MppFrameColorPrimaries(C.MPP_FRAME_PRI_BT470BG)      // also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM
	MppFramePrimariesSmpte170M  = MppFrameColorPrimaries(C.MPP_FRAME_PRI_SMPTE170M)    // also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC/SMPTE ST 170 (2004)
	MppFramePrimariesSmpte240M  = MppFrameColorPrimaries(C.MPP_FRAME_PRI_SMPTE240M)    // functionally identical to above/SMPTE ST 240
	MppFramePrimariesFilm       = MppFrameColorPrimaries(C.MPP_FRAME_PRI_FILM)         // colour filters using Illuminant C
	MppFramePrimariesBt2020     = MppFrameColorPrimaries(C.MPP_FRAME_PRI_BT2020)       // ITU-R BT2020 / ITU-R BT.2100-2
	MppFramePrimariesSmpte428_1 = MppFrameColorPrimaries(C.MPP_FRAME_PRI_SMPTEST428_1) // SMPTE ST 428-1 (CIE 1931 XYZ)
	MppFramePrimariesSmpte431   = MppFrameColorPrimaries(C.MPP_FRAME_PRI_SMPTE431)     // SMPTE ST 431-2 (2011) / DCI P3
	MppFramePrimariesSmpte432   = MppFrameColorPrimaries(C.MPP_FRAME_PRI_SMPTE432)     // SMPTE ST 432-1 (2010) / P3 D65 / Display P3
	MppFramePrimariesJedecP22   = MppFrameColorPrimaries(C.MPP_FRAME_PRI_JEDEC_P22)    // JEDEC P22 phosphors
	MppFramePrimariesNB         = MppFrameColorPrimaries(C.MPP_FRAME_PRI_NB)           // Not part of ABI
)

// Color Transfer Characteristic
type MppFrameColorTransferCharacteristic = C.MppFrameColorTransferCharacteristic

const (
	MppFrameTrcReserved0    = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_RESERVED0)
	MppFrameTrcBt709        = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_BT709) // also ITU-R BT1361
	MppFrameTrcUnspecified  = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_UNSPECIFIED)
	MppFrameTrcReserved     = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_RESERVED)
	MppFrameTrcGamma22      = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_GAMMA22)   // also ITU-R BT470M / ITU-R BT1700 625 PAL & SECAM
	MppFrameTrcGamma28      = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_GAMMA28)   // also ITU-R BT470BG
	MppFrameTrcSmpte170M    = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_SMPTE170M) // also ITU-R BT601-6 525 or 625 / ITU-R BT1358 525 or 625 / ITU-R BT1700 NTSC
	MppFrameTrcSmpte240M    = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_SMPTE240M)
	MppFrameTrcLinear       = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_LINEAR)       // "Linear transfer characteristics"
	MppFrameTrcLog          = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_LOG)          // "Logarithmic transfer characteristic (100:1 range)"
	MppFrameTrcLogSqrt      = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_LOG_SQRT)     // "Logarithmic transfer characteristic (100 * Sqrt(10) : 1 range)"
	MppFrameTrcIec61966_2_4 = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_IEC61966_2_4) // IEC 61966-2-4
	MppFrameTrcBt1361Ecg    = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_BT1361_ECG)   // ITU-R BT1361 Extended Colour Gamut
	MppFrameTrcIec61966_2_1 = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_IEC61966_2_1) // IEC 61966-2-1 (sRGB or sYCC)
	MppFrameTrcBt2020_10    = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_BT2020_10)    // ITU-R BT2020 for 10 bit system
	MppFrameTrcBt2020_12    = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_BT2020_12)    // ITU-R BT2020 for 12 bit system
	MppFrameTrcSmpte2084    = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_SMPTEST2084)  // SMPTE ST 2084 for 10-, 12-, 14- and 16-bit systems
	MppFrameTrcSmpte428_1   = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_SMPTEST428_1) // SMPTE ST 428-1
	MppFrameTrcAribStdB67   = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_ARIB_STD_B67) // ARIB STD-B67, known as "Hybrid log-gamma"
	MppFrameTrcNB           = MppFrameColorTransferCharacteristic(C.MPP_FRAME_TRC_NB)           // Not part of ABI
)

// YUV colorspace type
type MppFrameColorSpace = C.MppFrameColorSpace

const (
	MppFrameSpcRGB              = MppFrameColorSpace(C.MPP_FRAME_SPC_RGB)   // order of coefficients is actually GBR, also IEC 61966-2-1 (sRGB)
	MppFrameSpcBt709            = MppFrameColorSpace(C.MPP_FRAME_SPC_BT709) // also ITU-R BT1361 / IEC 61966-2-4 xvYCC709 / SMPTE RP177 Annex B
	MppFrameSpcUnspecified      = MppFrameColorSpace(C.MPP_FRAME_SPC_UNSPECIFIED)
	MppFrameSpcReserved         = MppFrameColorSpace(C.MPP_FRAME_SPC_RESERVED)
	MppFrameSpcFcc              = MppFrameColorSpace(C.MPP_FRAME_SPC_FCC)       // FCC Title 47 Code of Federal Regulations 73.682 (a)(20)
	MppFrameSpcBt470bg          = MppFrameColorSpace(C.MPP_FRAME_SPC_BT470BG)   // also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM / IEC 61966-2-4 xvYCC601
	MppFrameSpcSmpte170m        = MppFrameColorSpace(C.MPP_FRAME_SPC_SMPTE170M) // also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC / functionally identical to above
	MppFrameSpcSmpte240m        = MppFrameColorSpace(C.MPP_FRAME_SPC_SMPTE240M)
	MppFrameSpcYcocg            = MppFrameColorSpace(C.MPP_FRAME_SPC_YCOCG)              // Used by Dirac / VC-2 and H.264 FRext, see ITU-T SG16
	MppFrameSpcBt2020Ncl        = MppFrameColorSpace(C.MPP_FRAME_SPC_BT2020_NCL)         // ITU-R BT2020 non-constant luminance system
	MppFrameSpcBt2020Cl         = MppFrameColorSpace(C.MPP_FRAME_SPC_BT2020_CL)          // ITU-R BT2020 constant luminance system
	MppFrameSpcSmpte2085        = MppFrameColorSpace(C.MPP_FRAME_SPC_SMPTE2085)          // SMPTE 2085, Y'D'zD'x
	MppFrameSpcChromaDerivedNcl = MppFrameColorSpace(C.MPP_FRAME_SPC_CHROMA_DERIVED_NCL) // Chromaticity-derived non-constant luminance system
	MppFrameSpcChromaDerivedCl  = MppFrameColorSpace(C.MPP_FRAME_SPC_CHROMA_DERIVED_CL)  // Chromaticity-derived constant luminance system
	MppFrameSpcIctcp            = MppFrameColorSpace(C.MPP_FRAME_SPC_ICTCP)              // ITU-R BT.2100-0, ICtCp
	MppFrameSpcNB               = MppFrameColorSpace(C.MPP_FRAME_SPC_NB)                 // Not part of ABI
)

// Location of chroma samples
type MppFrameChromaLocation = C.MppFrameChromaLocation

const (
	MppChromaLocUnspecified = MppFrameChromaLocation(C.MPP_CHROMA_LOC_UNSPECIFIED)
	MppChromaLocLeft        = MppFrameChromaLocation(C.MPP_CHROMA_LOC_LEFT)    // mpeg2/4 4:2:0, h264 default for 4:2:0
	MppChromaLocCenter      = MppFrameChromaLocation(C.MPP_CHROMA_LOC_CENTER)  // mpeg1 4:2:0, jpeg 4:2:0, h263 4:2:0
	MppChromaLocTopLeft     = MppFrameChromaLocation(C.MPP_CHROMA_LOC_TOPLEFT) // ITU-R 601, SMPTE 274M 296M S314M(DV 4:1:1), mpeg2 4:2:2
	MppChromaLocTop         = MppFrameChromaLocation(C.MPP_CHROMA_LOC_TOP)
	MppChromaLocBottomLeft  = MppFrameChromaLocation(C.MPP_CHROMA_LOC_BOTTOMLEFT)
	MppChromaLocBottom      = MppFrameChromaLocation(C.MPP_CHROMA_LOC_BOTTOM)
	MppChromaLocNB          = MppFrameChromaLocation(C.MPP_CHROMA_LOC_NB) // Not part of ABI
)

type MppFrameChromaFormat = C.int

const (
	MppChromaFormatUnspecified = MppFrameChromaFormat(C.MPP_CHROMA_UNSPECIFIED)
	MppChromaFormat400         = MppFrameChromaFormat(C.MPP_CHROMA_400)
	MppChromaFormat410         = MppFrameChromaFormat(C.MPP_CHROMA_410)
	MppChromaFormat411         = MppFrameChromaFormat(C.MPP_CHROMA_411)
	MppChromaFormat420         = MppFrameChromaFormat(C.MPP_CHROMA_420)
	MppChromaFormat422         = MppFrameChromaFormat(C.MPP_CHROMA_422)
	MppChromaFormat440         = MppFrameChromaFormat(C.MPP_CHROMA_440)
	MppChromaFormat444         = MppFrameChromaFormat(C.MPP_CHROMA_444)
)

// MppFrameFormat bit flag
const (
	MppFrameFmtMask     = C.MPP_FRAME_FMT_MASK
	MppFrameFmtPropMask = C.MPP_FRAME_FMT_PROP_MASK

	MppFrameFmtColorMask = C.MPP_FRAME_FMT_COLOR_MASK
	MppFrameFmtYuv       = C.MPP_FRAME_FMT_YUV
	MppFrameFmtRgb       = C.MPP_FRAME_FMT_RGB

	MppFrameFbcMask = C.MPP_FRAME_FBC_MASK
	MppFrameFbcNone = C.MPP_FRAME_FBC_NONE

	MppFrameHdrMask = C.MPP_FRAME_HDR_MASK
	MppFrameHdrNone = C.MPP_FRAME_HDR_NONE
	MppFrameHdr     = C.MPP_FRAME_HDR

	MppFrameTileFlag = C.MPP_FRAME_TILE_FLAG
)

// AFBC_V1 is for ISP output.
const MppFrameFbcAfbcV1 = C.MPP_FRAME_FBC_AFBC_V1

// AFBC_V2 is for video decoder output.
const MppFrameFbcAfbcV2 = C.MPP_FRAME_FBC_AFBC_V2

const MppFrameFmtLeMask = C.MPP_FRAME_FMT_LE_MASK

func MppFrameFmtIsYUV(fmt MppFrameFormat) bool {
	return ((fmt & MppFrameFmtColorMask) == MppFrameFmtYuv) && ((fmt & MppFrameFmtMask) < MppFmtYuvButt)
}

func MppFrameFmtIsYUV10Bit(fmt MppFrameFormat) bool {
	return (fmt&MppFrameFmtMask) == MppFmtYuv420sp10Bit || (fmt&MppFrameFmtMask) == MppFmtYuv422sp10Bit
}

func MppFrameFmtIsRGB(fmt MppFrameFormat) bool {
	return ((fmt & MppFrameFmtColorMask) == MppFrameFmtRgb) && ((fmt & MppFrameFmtMask) < MppFmtRgbButt)
}

// For MPP_FRAME_FBC_AFBC_V1 the 16byte aligned stride is used.
func MppFrameFmtIsFbc(fmt MppFrameFormat) bool {
	return (fmt & MppFrameFbcMask) != MppFrameFbcNone
}

func MppFrameFmtIsHdr(fmt MppFrameFormat) bool {
	return (fmt & MppFrameHdrMask) != MppFrameHdrNone
}

func MppFrameFmtIsLe(fmt MppFrameFormat) bool {
	return (fmt & MppFrameFmtLeMask) == MppFrameFmtLeMask
}

func MppFrameFmtIsBe(fmt MppFrameFormat) bool {
	return (fmt & MppFrameFmtLeMask) == 0
}

func MppFrameFmtIsTile(fmt MppFrameFormat) bool {
	return (fmt & MppFrameTileFlag) != 0
}

// mpp color format index definition
type MppFrameFormat = C.MppFrameFormat

const (
	MppFmtYuv420sp = MppFrameFormat(C.MPP_FMT_YUV420SP) // YYYY... UV... (NV12)
	/*
	 * A rockchip specific pixel format, without gap between pixel aganist
	 * the P010_10LE/P010_10BE
	 */
	MppFmtYuv420sp10Bit = MppFrameFormat(C.MPP_FMT_YUV420SP_10BIT)
	MppFmtYuv422sp      = MppFrameFormat(C.MPP_FMT_YUV422SP)       // YYYY... UVUV... (NV16)
	MppFmtYuv422sp10Bit = MppFrameFormat(C.MPP_FMT_YUV422SP_10BIT) ///< Not part of ABI
	MppFmtYuv420p       = MppFrameFormat(C.MPP_FMT_YUV420P)        // YYYY... U...V...  (I420)
	MppFmtYuv420spVu    = MppFrameFormat(C.MPP_FMT_YUV420SP_VU)    // YYYY... VUVUVU... (NV21)
	MppFmtYuv422p       = MppFrameFormat(C.MPP_FMT_YUV422P)        // YYYY... UU...VV...(422P)
	MppFmtYuv422spVu    = MppFrameFormat(C.MPP_FMT_YUV422SP_VU)    // YYYY... VUVUVU... (NV61)
	MppFmtYuv422Yuyv    = MppFrameFormat(C.MPP_FMT_YUV422_YUYV)    // YUYVYUYV... (YUY2)
	MppFmtYuv422Yvyu    = MppFrameFormat(C.MPP_FMT_YUV422_YVYU)    // YVYUYVYU... (YVY2)
	MppFmtYuv422Uyvy    = MppFrameFormat(C.MPP_FMT_YUV422_UYVY)    // UYVYUYVY... (UYVY)
	MppFmtYuv422Vyuy    = MppFrameFormat(C.MPP_FMT_YUV422_VYUY)    // VYUYVYUY... (VYUY)
	MppFmtYuv400        = MppFrameFormat(C.MPP_FMT_YUV400)         // YYYY...
	MppFmtYuv440sp      = MppFrameFormat(C.MPP_FMT_YUV440SP)       // YYYY... UVUV...
	MppFmtYuv411sp      = MppFrameFormat(C.MPP_FMT_YUV411SP)       // YYYY... UV...
	MppFmtYuv444sp      = MppFrameFormat(C.MPP_FMT_YUV444SP)       // YYYY... UVUVUVUV...
	MppFmtYuv444p       = MppFrameFormat(C.MPP_FMT_YUV444P)        // YYYY... UUUU... VVVV...
	MppFmtYuvButt       = MppFrameFormat(C.MPP_FMT_YUV_BUTT)

	MppFmtRgb565    = MppFrameFormat(C.MPP_FMT_RGB565)    // 16-bit RGB
	MppFmtBgr565    = MppFrameFormat(C.MPP_FMT_BGR565)    // 16-bit RGB
	MppFmtRgb555    = MppFrameFormat(C.MPP_FMT_RGB555)    // 15-bit RGB
	MppFmtBgr555    = MppFrameFormat(C.MPP_FMT_BGR555)    // 15-bit RGB
	MppFmtRgb444    = MppFrameFormat(C.MPP_FMT_RGB444)    // 12-bit RGB
	MppFmtBgr444    = MppFrameFormat(C.MPP_FMT_BGR444)    // 12-bit RGB
	MppFmtRgb888    = MppFrameFormat(C.MPP_FMT_RGB888)    // 24-bit RGB
	MppFmtBgr888    = MppFrameFormat(C.MPP_FMT_BGR888)    // 24-bit RGB
	MppFmtRgb101010 = MppFrameFormat(C.MPP_FMT_RGB101010) // 30-bit RGB
	MppFmtBgr101010 = MppFrameFormat(C.MPP_FMT_BGR101010) // 30-bit RGB
	MppFmtArgb8888  = MppFrameFormat(C.MPP_FMT_ARGB8888)  // 32-bit RGB
	MppFmtAbgr8888  = MppFrameFormat(C.MPP_FMT_ABGR8888)  // 32-bit RGB
	MppFmtBgra8888  = MppFrameFormat(C.MPP_FMT_BGRA8888)  // 32-bit RGB
	MppFmtRgba8888  = MppFrameFormat(C.MPP_FMT_RGBA8888)  // 32-bit RGB
	MppFmtRgbButt   = MppFrameFormat(C.MPP_FMT_RGB_BUTT)

	MppFmtButt = MppFrameFormat(C.MPP_FMT_BUTT)
)

type MppFrameRational struct {
	c C.struct_MppFrameRational
}

func (r *MppFrameRational) Num() int32 {
	return int32(r.c.num)
}

func (r *MppFrameRational) Den() int32 {
	return int32(r.c.den)
}

func (r *MppFrameRational) SetNum(num int32) {
	r.c.num = cS32(num)
}

func (r *MppFrameRational) SetDen(den int32) {
	r.c.den = cS32(den)
}

type MppFrameMasteringDisplayMetadata struct {
	c C.struct_MppFrameMasteringDisplayMetadata
}

func (m *MppFrameMasteringDisplayMetadata) DisplayPrimaries() [3][2]uint16 {
	var result [3][2]uint16
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			result[i][j] = uint16(m.c.display_primaries[i][j])
		}
	}
	return result
}

func (m *MppFrameMasteringDisplayMetadata) WhitePoint() [2]uint16 {
	var result [2]uint16
	for i := 0; i < 2; i++ {
		result[i] = uint16(m.c.white_point[i])
	}
	return result
}

func (m *MppFrameMasteringDisplayMetadata) MaxLuminance() uint32 {
	return uint32(m.c.max_luminance)
}

func (m *MppFrameMasteringDisplayMetadata) MinLuminance() uint32 {
	return uint32(m.c.min_luminance)
}

func (m *MppFrameMasteringDisplayMetadata) SetDisplayPrimaries(displayPrimaries [3][2]uint16) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			m.c.display_primaries[i][j] = cU16(displayPrimaries[i][j])
		}
	}
}

func (m *MppFrameMasteringDisplayMetadata) SetWhitePoint(whitePoint [2]uint16) {
	for i := 0; i < 2; i++ {
		m.c.white_point[i] = cU16(whitePoint[i])
	}
}

func (m *MppFrameMasteringDisplayMetadata) SetMaxLuminance(maxLuminance uint32) {
	m.c.max_luminance = cU32(maxLuminance)
}

func (m *MppFrameMasteringDisplayMetadata) SetMinLuminance(minLuminance uint32) {
	m.c.min_luminance = cU32(minLuminance)
}

type MppFrameContentLightMetadata struct {
	c C.struct_MppFrameContentLightMetadata
}

func (m *MppFrameContentLightMetadata) MaxCLL() uint16 {
	return uint16(m.c.MaxCLL)
}

func (m *MppFrameContentLightMetadata) MaxFALL() uint16 {
	return uint16(m.c.MaxFALL)
}

func (m *MppFrameContentLightMetadata) SetMaxCLL(maxCLL uint16) {
	m.c.MaxCLL = cU16(maxCLL)
}

func (m *MppFrameContentLightMetadata) SetMaxFALL(maxFALL uint16) {
	m.c.MaxFALL = cU16(maxFALL)
}

type MppFrameHdrDynamicMeta struct {
	c *C.struct_MppFrameHdrDynamicMeta
}

func (m *MppFrameHdrDynamicMeta) HdrFmt() uint32 {
	return uint32(m.c.hdr_fmt)
}

func (m *MppFrameHdrDynamicMeta) Size() uint32 {
	return uint32(m.c.size)
}

func (m *MppFrameHdrDynamicMeta) SetHdrFmt(hdrFmt uint32) {
	m.c.hdr_fmt = cU32(hdrFmt)
}

func (m *MppFrameHdrDynamicMeta) SetSize(size uint32) {
	m.c.size = cU32(size)
}

type MppFrameError = C.int

const (
	MppFrameErrUnknow     = MppFrameError(C.MPP_FRAME_ERR_UNKNOW)       // General error not specified
	MppFrameErrUnsupport  = MppFrameError(C.MPP_FRAME_ERR_UNSUPPORT)    // Critical error for decoder not support error
	MppFrameErrDecInvalid = MppFrameError(C.MPP_FRAME_ERR_DEC_INVALID)  // Fatal error for decoder can not parse a valid frame for hardware. the pixel data is all invalid.
	MppFrameErrDecHwErr   = MppFrameError(C.MPP_FRAME_ERR_DEC_HW_ERR)   // Normal error for decoder found hardware error on decoding.
	MppFrameErrDecMissRef = MppFrameError(C.MPP_FRAME_ERR_DEC_MISS_REF) // Normal error for decoder found missing reference frame on decoding.
)

func MppFrameInit() (*MppFrame, MppRet) {
	cFrame := C.MppFrame(nil)
	ret := MppRet(C.mpp_frame_init(&cFrame))
	return &MppFrame{c: &cFrame}, ret
}

func (f *MppFrame) Deinit() MppRet {
	return MppRet(C.mpp_frame_deinit(f.c))
}

func (f *MppFrame) GetWidth() uint32 {
	return uint32(C.mpp_frame_get_width(*f.c))
}

func (f *MppFrame) SetWidth(width uint32) {
	C.mpp_frame_set_width(*f.c, cU32(width))
}

func (f *MppFrame) GetHeight() uint32 {
	return uint32(C.mpp_frame_get_height(*f.c))
}

func (f *MppFrame) SetHeight(height uint32) {
	C.mpp_frame_set_height(*f.c, cU32(height))
}

func (f *MppFrame) GetHorStride() uint32 {
	return uint32(C.mpp_frame_get_hor_stride(*f.c))
}

func (f *MppFrame) SetHorStride(horStride uint32) {
	C.mpp_frame_set_hor_stride(*f.c, cU32(horStride))
}

func (f *MppFrame) GetVerStride() uint32 {
	return uint32(C.mpp_frame_get_ver_stride(*f.c))
}

func (f *MppFrame) SetVerStride(verStride uint32) {
	C.mpp_frame_set_ver_stride(*f.c, cU32(verStride))
}

func (f *MppFrame) SetHorStridePixel(horStridePixel uint32) {
	C.mpp_frame_set_hor_stride_pixel(*f.c, cU32(horStridePixel))
}

func (f *MppFrame) GetHorStridePixel() uint32 {
	return uint32(C.mpp_frame_get_hor_stride_pixel(*f.c))
}

func (f *MppFrame) SetFbcHdrStride(fbcHdrStride uint32) {
	C.mpp_frame_set_fbc_hdr_stride(*f.c, cU32(fbcHdrStride))
}

func (f *MppFrame) GetFbcHdrStride() uint32 {
	return uint32(C.mpp_frame_get_fbc_hdr_stride(*f.c))
}

func (f *MppFrame) GetOffsetX() uint32 {
	return uint32(C.mpp_frame_get_offset_x(*f.c))
}

func (f *MppFrame) SetOffsetX(offsetX uint32) {
	C.mpp_frame_set_offset_x(*f.c, cU32(offsetX))
}

func (f *MppFrame) GetOffsetY() uint32 {
	return uint32(C.mpp_frame_get_offset_y(*f.c))
}

func (f *MppFrame) SetOffsetY(offsetY uint32) {
	C.mpp_frame_set_offset_y(*f.c, cU32(offsetY))
}

func (f *MppFrame) GetMode() uint32 {
	return uint32(C.mpp_frame_get_mode(*f.c))
}

func (f *MppFrame) SetMode(mode uint32) {
	C.mpp_frame_set_mode(*f.c, cU32(mode))
}

func (f *MppFrame) GetDiscard() uint32 {
	return uint32(C.mpp_frame_get_discard(*f.c))
}

func (f *MppFrame) SetDiscard(discard uint32) {
	C.mpp_frame_set_discard(*f.c, cU32(discard))
}

func (f *MppFrame) GetViewId() uint32 {
	return uint32(C.mpp_frame_get_viewid(*f.c))
}

func (f *MppFrame) SetViewId(viewId uint32) {
	C.mpp_frame_set_viewid(*f.c, cU32(viewId))
}

func (f *MppFrame) GetPoc() uint32 {
	return uint32(C.mpp_frame_get_poc(*f.c))
}

func (f *MppFrame) SetPoc(poc uint32) {
	C.mpp_frame_set_poc(*f.c, cU32(poc))
}

func (f *MppFrame) GetPts() int64 {
	return int64(C.mpp_frame_get_pts(*f.c))
}

func (f *MppFrame) SetPts(pts int64) {
	C.mpp_frame_set_pts(*f.c, cS64(pts))
}

func (f *MppFrame) GetDts() int64 {
	return int64(C.mpp_frame_get_dts(*f.c))
}

func (f *MppFrame) SetDts(dts int64) {
	C.mpp_frame_set_dts(*f.c, cS64(dts))
}

func (f *MppFrame) GetErrInfo() uint32 {
	return uint32(C.mpp_frame_get_errinfo(*f.c))
}

func (f *MppFrame) SetErrInfo(errInfo uint32) {
	C.mpp_frame_set_errinfo(*f.c, cU32(errInfo))
}

func (f *MppFrame) GetBufSize() int64 {
	return int64(C.mpp_frame_get_buf_size(*f.c))
}

func (f *MppFrame) SetBufSize(bufSize int64) {
	C.mpp_frame_set_buf_size(*f.c, C.size_t(bufSize))
}

func (f *MppFrame) SetThumbnailEn(thumbnailEn uint32) {
	C.mpp_frame_set_thumbnail_en(*f.c, cU32(thumbnailEn))
}

func (f *MppFrame) GetThumbnailEn() uint32 {
	return uint32(C.mpp_frame_get_thumbnail_en(*f.c))
}

// flow control parmeter
func (f *MppFrame) GetEos() uint32 {
	return uint32(C.mpp_frame_get_eos(*f.c))
}

func (f *MppFrame) SetEos(eos uint32) {
	C.mpp_frame_set_eos(*f.c, cU32(eos))
}

func (f *MppFrame) GetInfoChange() uint32 {
	return uint32(C.mpp_frame_get_info_change(*f.c))
}

func (f *MppFrame) SetInfoChange(infoChange uint32) {
	C.mpp_frame_set_info_change(*f.c, cU32(infoChange))
}

// buffer parameter
func (f *MppFrame) GetBuffer() *MppBuffer {
	return &MppBuffer{c: C.mpp_frame_get_buffer(*f.c)}
}

func (f *MppFrame) SetBuffer(buf *MppBuffer) {
	C.mpp_frame_set_buffer(*f.c, buf.c)
}

// meta data parameter
func (f *MppFrame) HasMeta() uint32 {
	return uint32(C.mpp_frame_has_meta(*f.c))
}

func (f *MppFrame) GetMeta() *MppMeta {
	return &MppMeta{c: C.mpp_frame_get_meta(*f.c)}
}

func (f *MppFrame) SetMeta(meta *MppMeta) {
	C.mpp_frame_set_meta(*f.c, meta.c)
}

// color related parameter
func (f *MppFrame) GetColorRange() MppFrameColorRange {
	return MppFrameColorRange(C.mpp_frame_get_color_range(*f.c))
}

func (f *MppFrame) SetColorRange(colorRange MppFrameColorRange) {
	C.mpp_frame_set_color_range(*f.c, C.MppFrameColorRange(colorRange))
}

func (f *MppFrame) GetColorPrimaries() MppFrameColorPrimaries {
	return MppFrameColorPrimaries(C.mpp_frame_get_color_primaries(*f.c))
}

func (f *MppFrame) SetColorPrimaries(colorPrimaries MppFrameColorPrimaries) {
	C.mpp_frame_set_color_primaries(*f.c, colorPrimaries)
}

func (f *MppFrame) GetColorTrc() MppFrameColorTransferCharacteristic {
	return MppFrameColorTransferCharacteristic(C.mpp_frame_get_color_trc(*f.c))
}

func (f *MppFrame) SetColorTrc(colorTrc MppFrameColorTransferCharacteristic) {
	C.mpp_frame_set_color_trc(*f.c, colorTrc)
}

func (f *MppFrame) GetColorSpace() MppFrameColorSpace {
	return MppFrameColorSpace(C.mpp_frame_get_colorspace(*f.c))
}

func (f *MppFrame) SetColorSpace(colorSpace MppFrameColorSpace) {
	C.mpp_frame_set_colorspace(*f.c, colorSpace)
}

func (f *MppFrame) GetChromaLocation() MppFrameChromaLocation {
	return MppFrameChromaLocation(C.mpp_frame_get_chroma_location(*f.c))
}

func (f *MppFrame) SetChromaLocation(chromaLocation MppFrameChromaLocation) {
	C.mpp_frame_set_chroma_location(*f.c, chromaLocation)
}

func (f *MppFrame) GetFmt() MppFrameFormat {
	return MppFrameFormat(C.mpp_frame_get_fmt(*f.c))
}

func (f *MppFrame) SetFmt(fmt MppFrameFormat) {
	C.mpp_frame_set_fmt(*f.c, fmt)
}

func (f *MppFrame) GetSar() *MppFrameRational {
	return &MppFrameRational{c: C.mpp_frame_get_sar(*f.c)}
}

func (f *MppFrame) SetSar(sar *MppFrameRational) {
	C.mpp_frame_set_sar(*f.c, sar.c)
}

func (f *MppFrame) GetMasteringDisplay() *MppFrameMasteringDisplayMetadata {
	return &MppFrameMasteringDisplayMetadata{c: C.mpp_frame_get_mastering_display(*f.c)}
}

func (f *MppFrame) SetMasteringDisplay(masteringDisplay *MppFrameMasteringDisplayMetadata) {
	C.mpp_frame_set_mastering_display(*f.c, masteringDisplay.c)
}

func (f *MppFrame) GetContentLight() *MppFrameContentLightMetadata {
	return &MppFrameContentLightMetadata{c: C.mpp_frame_get_content_light(*f.c)}
}

func (f *MppFrame) SetContentLight(contentLight *MppFrameContentLightMetadata) {
	C.mpp_frame_set_content_light(*f.c, contentLight.c)
}

func (f *MppFrame) GetHdrDynamicMeta() *MppFrameHdrDynamicMeta {
	return &MppFrameHdrDynamicMeta{c: C.mpp_frame_get_hdr_dynamic_meta(*f.c)}
}

func (f *MppFrame) SetHdrDynamicMeta(hdrDynamicMeta *MppFrameHdrDynamicMeta) {
	C.mpp_frame_set_hdr_dynamic_meta(*f.c, hdrDynamicMeta.c)
}
