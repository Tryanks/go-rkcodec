package rkcodec

//#include <rockchip/rk_hdr_meta_com.h>
import "C"

type HdrCodecType = C.HdrCodecType

const (
	HdrAVS2     = HdrCodecType(C.HDR_AVS2)
	HdrHEVC     = HdrCodecType(C.HDR_HEVC)
	HdrH264     = HdrCodecType(C.HDR_H264)
	HdrAV1      = HdrCodecType(C.HDR_AV1)
	HdrCodecBut = HdrCodecType(C.HDR_CODEC_BUT)
)

type HdrFormat = C.HdrFormat

const (
	HdrNone = HdrFormat(C.HDR_NONE)
	Hdr10   = HdrFormat(C.HDR10)
	Hlg     = HdrFormat(C.HLG)
	// RESERVED3 = 3, //reserved for more future static hdr format
	// RESERVED4 = 4, //reserved for more future static hdr format
	Hdrvivid = HdrFormat(C.HDRVIVID)
	// RESERVED6 = 6, //reserved for hdr vivid
	// RESERVED7 = 7, //reserved for hdr vivid
	Hdr10plus = HdrFormat(C.HDR10PLUS)
	// RESERVED9 = 9, //reserved for hdr10+
	// RESERVED10 = 10,//reserved for hdr10+
	Dolby = HdrFormat(C.DOLBY)
	// RESERVED12 = 12, //reserved for other dynamic hdr format
	// RESERVED13 = 13, //reserved for  other dynamic hdr format
	HdrFormatMax = HdrFormat(C.HDR_FORMAT_MAX)
)

type HdrPayloadFormat = C.HdrPayloadFormat

const (
	HdrPayloadStatic    = HdrPayloadFormat(C.STATIC)
	HdrPayloadDynamic   = HdrPayloadFormat(C.DYNAMIC)
	HdrPayloadFormatMax = HdrPayloadFormat(C.HDR_PAYLOAD_FORMAT_MAX)
)

type HdrStaticMeta struct {
	c C.HdrStaticMeta
}

func (m *HdrStaticMeta) GetColorSpace() uint32 {
	return uint32(m.c.color_space)
}

func (m *HdrStaticMeta) SetColorSpace(v uint32) {
	m.c.color_space = cU32(v)
}

func (m *HdrStaticMeta) GetColorPrimaries() uint32 {
	return uint32(m.c.color_primaries)
}

func (m *HdrStaticMeta) SetColorPrimaries(v uint32) {
	m.c.color_primaries = cU32(v)
}

func (m *HdrStaticMeta) GetColorTrc() uint32 {
	return uint32(m.c.color_trc)
}

func (m *HdrStaticMeta) SetColorTrc(v uint32) {
	m.c.color_trc = cU32(v)
}

func (m *HdrStaticMeta) GetRedX() uint32 {
	return uint32(m.c.red_x)
}

func (m *HdrStaticMeta) SetRedX(v uint32) {
	m.c.red_x = cU32(v)
}

func (m *HdrStaticMeta) GetRedY() uint32 {
	return uint32(m.c.red_y)
}

func (m *HdrStaticMeta) SetRedY(v uint32) {
	m.c.red_y = cU32(v)
}

func (m *HdrStaticMeta) GetGreenX() uint32 {
	return uint32(m.c.green_x)
}

func (m *HdrStaticMeta) SetGreenX(v uint32) {
	m.c.green_x = cU32(v)
}

func (m *HdrStaticMeta) GetGreenY() uint32 {
	return uint32(m.c.green_y)
}

func (m *HdrStaticMeta) SetGreenY(v uint32) {
	m.c.green_y = cU32(v)
}

func (m *HdrStaticMeta) GetBlueX() uint32 {
	return uint32(m.c.blue_x)
}

func (m *HdrStaticMeta) SetBlueX(v uint32) {
	m.c.blue_x = cU32(v)
}

func (m *HdrStaticMeta) GetBlueY() uint32 {
	return uint32(m.c.blue_y)
}

func (m *HdrStaticMeta) SetBlueY(v uint32) {
	m.c.blue_y = cU32(v)
}

func (m *HdrStaticMeta) GetWhitePointX() uint32 {
	return uint32(m.c.white_point_x)
}

func (m *HdrStaticMeta) SetWhitePointX(v uint32) {
	m.c.white_point_x = cU32(v)
}

func (m *HdrStaticMeta) GetWhitePointY() uint32 {
	return uint32(m.c.white_point_y)
}

func (m *HdrStaticMeta) SetWhitePointY(v uint32) {
	m.c.white_point_y = cU32(v)
}

func (m *HdrStaticMeta) GetMinLuminance() uint32 {
	return uint32(m.c.min_luminance)
}

func (m *HdrStaticMeta) SetMinLuminance(v uint32) {
	m.c.min_luminance = cU32(v)
}

func (m *HdrStaticMeta) GetMaxLuminance() uint32 {
	return uint32(m.c.max_luminance)
}

func (m *HdrStaticMeta) SetMaxLuminance(v uint32) {
	m.c.max_luminance = cU32(v)
}

func (m *HdrStaticMeta) GetMaxCll() uint32 {
	return uint32(m.c.max_cll)
}

func (m *HdrStaticMeta) SetMaxCll(v uint32) {
	m.c.max_cll = cU32(v)
}

func (m *HdrStaticMeta) GetMaxFall() uint32 {
	return uint32(m.c.max_fall)
}

func (m *HdrStaticMeta) SetMaxFall(v uint32) {
	m.c.max_fall = cU32(v)
}

type RkMetaHdrHeader struct {
	c C.RkMetaHdrHeader
}

func (m *RkMetaHdrHeader) GetMagic() uint16 {
	return uint16(m.c.magic)
}

func (m *RkMetaHdrHeader) SetMagic(v uint16) {
	m.c.magic = cU16(v)
}

func (m *RkMetaHdrHeader) GetSize() uint16 {
	return uint16(m.c.size)
}

func (m *RkMetaHdrHeader) SetSize(v uint16) {
	m.c.size = cU16(v)
}

func (m *RkMetaHdrHeader) GetMessageTotal() uint16 {
	return uint16(m.c.message_total)
}

func (m *RkMetaHdrHeader) SetMessageTotal(v uint16) {
	m.c.message_total = cU16(v)
}

func (m *RkMetaHdrHeader) GetMessageIndex() uint16 {
	return uint16(m.c.message_index)
}

func (m *RkMetaHdrHeader) SetMessageIndex(v uint16) {
	m.c.message_index = cU16(v)
}

func (m *RkMetaHdrHeader) GetVersion() uint16 {
	return uint16(m.c.version)
}

func (m *RkMetaHdrHeader) SetVersion(v uint16) {
	m.c.version = cU16(v)
}

func (m *RkMetaHdrHeader) GetHdrFormat() uint16 {
	return uint16(m.c.hdr_format)
}

func (m *RkMetaHdrHeader) SetHdrFormat(v uint16) {
	m.c.hdr_format = cU16(v)
}

func (m *RkMetaHdrHeader) GetHdrPayloadType() uint16 {
	return uint16(m.c.hdr_payload_type)
}

func (m *RkMetaHdrHeader) SetHdrPayloadType(v uint16) {
	m.c.hdr_payload_type = cU16(v)
}

func (m *RkMetaHdrHeader) GetVideoFormat() uint16 {
	return uint16(m.c.video_format)
}

func (m *RkMetaHdrHeader) SetVideoFormat(v uint16) {
	m.c.video_format = cU16(v)
}

func (f *MppFrame) FillHdrMeta(t HdrCodecType) {
	C.fill_hdr_meta_to_frame(*f.c, C.HdrCodecType(t))
}
