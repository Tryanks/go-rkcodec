package rkcodec

//#include <rockchip/mpp_meta.h>
//#include <rockchip/rk_type.h>
import "C"
import "unsafe"

type MppMetaType = cU32

const (
	MetaTypeFrame  = MppMetaType(C.TYPE_FRAME)
	MetaTypePacket = MppMetaType(C.TYPE_PACKET)
	MetaTypeBuffer = MppMetaType(C.TYPE_BUFFER)

	MetaTypeS32 = MppMetaType(C.TYPE_S32)
	MetaTypeS64 = MppMetaType(C.TYPE_S64)
	MetaTypePtr = MppMetaType(C.TYPE_PTR)
)

type MppMetaKey = C.MppMetaKey

const (
	// data flow key
	MetaKeyInputFrame   = MppMetaKey(C.KEY_INPUT_FRAME)
	MetaKeyInputPacket  = MppMetaKey(C.KEY_INPUT_PACKET)
	MetaKeyOutputFrame  = MppMetaKey(C.KEY_OUTPUT_FRAME)
	MetaKeyOutputPacket = MppMetaKey(C.KEY_OUTPUT_PACKET)
	// output motion information for motion detection
	MetaKeyMotionInfo    = MppMetaKey(C.KEY_MOTION_INFO)
	MetaKeyHdrInfo       = MppMetaKey(C.KEY_HDR_INFO)
	MetaKeyHdrMetaOffset = MppMetaKey(C.KEY_HDR_META_OFFSET)
	MetaKeyHdrMetaSize   = MppMetaKey(C.KEY_HDR_META_SIZE)

	// flow control key
	MetaKeyInputBlock  = MppMetaKey(C.KEY_INPUT_BLOCK)
	MetaKeyOutputBlock = MppMetaKey(C.KEY_OUTPUT_BLOCK)
	MetaKeyInputIdrReq = MppMetaKey(C.KEY_INPUT_IDR_REQ) // input idr frame request flag
	MetaKeyOutputIntra = MppMetaKey(C.KEY_OUTPUT_INTRA)  // output intra frame indicator

	// mpp_frame / mpp_packet meta data info key
	MetaKeyTemporalId   = MppMetaKey(C.KEY_TEMPORAL_ID)
	MetaKeyLongRefIdx   = MppMetaKey(C.KEY_LONG_REF_IDX)
	MetaKeyEncAverageQp = MppMetaKey(C.KEY_ENC_AVERAGE_QP)
	MetaKeyEncStartQp   = MppMetaKey(C.KEY_ENC_START_QP)
	MetaKeyRoiData      = MppMetaKey(C.KEY_ROI_DATA)
	MetaKeyOsdData      = MppMetaKey(C.KEY_OSD_DATA)
	MetaKeyOsdData2     = MppMetaKey(C.KEY_OSD_DATA2)
	MetaKeyUserData     = MppMetaKey(C.KEY_USER_DATA)
	MetaKeyUserDatas    = MppMetaKey(C.KEY_USER_DATAS)

	// num of inter different size predicted block
	MetaKeyLvl64InterNum = MppMetaKey(C.KEY_LVL64_INTER_NUM)
	MetaKeyLvl32InterNum = MppMetaKey(C.KEY_LVL32_INTER_NUM)
	MetaKeyLvl16InterNum = MppMetaKey(C.KEY_LVL16_INTER_NUM)
	MetaKeyLvl8InterNum  = MppMetaKey(C.KEY_LVL8_INTER_NUM)
	// num of intra different size predicted block
	MetaKeyLvl32IntraNum = MppMetaKey(C.KEY_LVL32_INTRA_NUM)
	MetaKeyLvl16IntraNum = MppMetaKey(C.KEY_LVL16_INTRA_NUM)
	MetaKeyLvl8IntraNum  = MppMetaKey(C.KEY_LVL8_INTRA_NUM)
	MetaKeyLvl4IntraNum  = MppMetaKey(C.KEY_LVL4_INTRA_NUM)
	// output P skip frame indicator
	MetaKeyOutputPskip = MppMetaKey(C.KEY_OUTPUT_PSKIP)
	MetaKeyEncSse      = MppMetaKey(C.KEY_ENC_SSE)

	// For vepu580 roi buffer config mode
	MetaKeyRoiData2 = MppMetaKey(C.KEY_ROI_DATA2)

	// qpmap for rv1109/1126 encoder qpmap config
	MetaKeyQpmap0 = MppMetaKey(C.KEY_QPMAP0)

	// input motion list for smart p rate control
	MetaKeyMvList = MppMetaKey(C.KEY_MV_LIST)

	// frame long-term reference frame operation
	MetaKeyEncMarkLtr = MppMetaKey(C.KEY_ENC_MARK_LTR)
	MetaKeyEncUseLtr  = MppMetaKey(C.KEY_ENC_USE_LTR)

	// MLVEC specified encoder feature
	MetaKeyEncFrameQp      = MppMetaKey(C.KEY_ENC_FRAME_QP)
	MetaKeyEncBaseLayerPid = MppMetaKey(C.KEY_ENC_BASE_LAYER_PID)

	// Thumbnail info for decoder output frame
	MetaKeyDecTbnEn       = MppMetaKey(C.KEY_DEC_TBN_EN)
	MetaKeyDecTbnYOffset  = MppMetaKey(C.KEY_DEC_TBN_Y_OFFSET)
	MetaKeyDecTbnUvOffset = MppMetaKey(C.KEY_DEC_TBN_UV_OFFSET)
)

func (m *MppMeta) GetWithTag(tag string, caller string) MppRet {
	return MppRet(C.mpp_meta_get_with_tag(&m.c, C.CString(tag), C.CString(caller)))
}

func (m *MppMeta) Put() MppRet {
	return MppRet(C.mpp_meta_put(m.c))
}

func (m *MppMeta) Size() int32 {
	return int32(C.mpp_meta_size(m.c))
}

func (m *MppMeta) SetS32(key MppMetaKey, val int32) MppRet {
	return MppRet(C.mpp_meta_set_s32(m.c, key, C.RK_S32(val)))
}

func (m *MppMeta) SetS64(key MppMetaKey, val int64) MppRet {
	return MppRet(C.mpp_meta_set_s64(m.c, key, C.RK_S64(val)))
}

func (m *MppMeta) SetPtr(key MppMetaKey, val unsafe.Pointer) MppRet {
	return MppRet(C.mpp_meta_set_ptr(m.c, key, val))
}

func (m *MppMeta) GetS32(key MppMetaKey) (int32, MppRet) {
	var val C.RK_S32
	ret := MppRet(C.mpp_meta_get_s32(m.c, key, &val))
	return int32(val), ret
}

func (m *MppMeta) GetS64(key MppMetaKey) (int64, MppRet) {
	var val C.RK_S64
	ret := MppRet(C.mpp_meta_get_s64(m.c, key, &val))
	return int64(val), ret
}

func (m *MppMeta) GetPtr(key MppMetaKey) (unsafe.Pointer, MppRet) {
	var val unsafe.Pointer
	ret := MppRet(C.mpp_meta_get_ptr(m.c, key, &val))
	return val, ret
}

func (m *MppMeta) SetFrame(key MppMetaKey, frame MppFrame) MppRet {
	return MppRet(C.mpp_meta_set_frame(m.c, key, *frame.c))
}

func (m *MppMeta) SetPacket(key MppMetaKey, packet MppPacket) MppRet {
	return MppRet(C.mpp_meta_set_packet(m.c, key, *packet.c))
}

func (m *MppMeta) SetBuffer(key MppMetaKey, buffer MppBuffer) MppRet {
	return MppRet(C.mpp_meta_set_buffer(m.c, key, buffer.c))
}

func (m *MppMeta) GetFrame(key MppMetaKey) (MppFrame, MppRet) {
	var frame MppFrame
	ret := MppRet(C.mpp_meta_get_frame(m.c, key, frame.c))
	return frame, ret
}

func (m *MppMeta) GetPacket(key MppMetaKey) (MppPacket, MppRet) {
	var packet MppPacket
	ret := MppRet(C.mpp_meta_get_packet(m.c, key, packet.c))
	return packet, ret
}

func (m *MppMeta) GetBuffer(key MppMetaKey) (MppBuffer, MppRet) {
	var buffer MppBuffer
	ret := MppRet(C.mpp_meta_get_buffer(m.c, key, &buffer.c))
	return buffer, ret
}

func (m *MppMeta) GetS32D(key MppMetaKey, def int32) (int32, MppRet) {
	var val C.RK_S32
	ret := MppRet(C.mpp_meta_get_s32_d(m.c, key, &val, C.RK_S32(def)))
	return int32(val), ret
}

func (m *MppMeta) GetS64D(key MppMetaKey, def int64) (int64, MppRet) {
	var val C.RK_S64
	ret := MppRet(C.mpp_meta_get_s64_d(m.c, key, &val, C.RK_S64(def)))
	return int64(val), ret
}

func (m *MppMeta) GetPtrD(key MppMetaKey, def unsafe.Pointer) (unsafe.Pointer, MppRet) {
	var val unsafe.Pointer
	ret := MppRet(C.mpp_meta_get_ptr_d(m.c, key, &val, def))
	return val, ret
}

func (m *MppMeta) GetFrameD(key MppMetaKey, def MppFrame) (MppFrame, MppRet) {
	var frame MppFrame
	ret := MppRet(C.mpp_meta_get_frame_d(m.c, key, frame.c, *def.c))
	return frame, ret
}

func (m *MppMeta) GetPacketD(key MppMetaKey, def MppPacket) (MppPacket, MppRet) {
	var packet MppPacket
	ret := MppRet(C.mpp_meta_get_packet_d(m.c, key, packet.c, *def.c))
	return packet, ret
}

func (m *MppMeta) GetBufferD(key MppMetaKey, def MppBuffer) (MppBuffer, MppRet) {
	var buffer MppBuffer
	ret := MppRet(C.mpp_meta_get_buffer_d(m.c, key, &buffer.c, def.c))
	return buffer, ret
}
