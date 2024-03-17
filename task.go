package rkcodec

//#include <rockchip/mpp_task.h>
import "C"
import "unsafe"

type MppPortType = C.MppPortType

const (
	MppPortInput  = MppPortType(C.MPP_PORT_INPUT)
	MppPortOutput = MppPortType(C.MPP_PORT_OUTPUT)
	MppPortButt   = MppPortType(C.MPP_PORT_BUTT)
)

type MppTaskWorkMode = C.MppTaskWorkMode

const (
	MppTaskAsync        = MppTaskWorkMode(C.MPP_TASK_ASYNC)
	MppTaskSync         = MppTaskWorkMode(C.MPP_TASK_SYNC)
	MppTaskWorkModeButt = MppTaskWorkMode(C.MPP_TASK_WORK_MODE_BUTT)
)

type MppPollType = C.MppPollType

const (
	MppPollButt     = MppPollType(C.MPP_POLL_BUTT)
	MppPollBlock    = MppPollType(C.MPP_POLL_BLOCK)
	MppPollNonBlock = MppPollType(C.MPP_POLL_NON_BLOCK)
	MppPollMax      = MppPollType(C.MPP_POLL_MAX)
)

const (
	MppTimeoutButt     = C.MPP_TIMEOUT_BUTT
	MppTimeoutBlock    = C.MPP_TIMEOUT_BLOCK
	MppTimeoutNonBlock = C.MPP_TIMEOUT_NON_BLOCK
	MppTimeoutMax      = C.MPP_TIMEOUT_MAX
)

func (task *MppTask) MetaSetS32(key MppMetaKey, val int32) MppRet {
	return MppRet(C.mpp_task_meta_set_s32(*task.c, C.MppMetaKey(key), cS32(val)))
}

func (task *MppTask) MetaSetS64(key MppMetaKey, val int64) MppRet {
	return MppRet(C.mpp_task_meta_set_s64(*task.c, C.MppMetaKey(key), cS64(val)))
}

func (task *MppTask) MetaSetPtr(key MppMetaKey, val unsafe.Pointer) MppRet {
	return MppRet(C.mpp_task_meta_set_ptr(*task.c, C.MppMetaKey(key), val))
}

func (task *MppTask) MetaSetFrame(key MppMetaKey, frame *MppFrame) MppRet {
	return MppRet(C.mpp_task_meta_set_frame(*task.c, C.MppMetaKey(key), *frame.c))
}

func (task *MppTask) MetaSetPacket(key MppMetaKey, packet *MppPacket) MppRet {
	return MppRet(C.mpp_task_meta_set_packet(*task.c, C.MppMetaKey(key), *packet.c))
}

func (task *MppTask) MetaSetBuffer(key MppMetaKey, buffer MppBuffer) MppRet {
	return MppRet(C.mpp_task_meta_set_buffer(*task.c, C.MppMetaKey(key), buffer.c))
}

func (task *MppTask) MetaGetS32(key MppMetaKey, defaultVal int32) (int32, MppRet) {
	var val cS32
	ret := MppRet(C.mpp_task_meta_get_s32(*task.c, C.MppMetaKey(key), &val, cS32(defaultVal)))
	return int32(val), ret
}

func (task *MppTask) MetaGetS64(key MppMetaKey, defaultVal int64) (int64, MppRet) {
	var val cS64
	ret := MppRet(C.mpp_task_meta_get_s64(*task.c, C.MppMetaKey(key), &val, cS64(defaultVal)))
	return int64(val), ret
}

func (task *MppTask) MetaGetPtr(key MppMetaKey, defaultVal unsafe.Pointer) (unsafe.Pointer, MppRet) {
	var val unsafe.Pointer
	ret := MppRet(C.mpp_task_meta_get_ptr(*task.c, C.MppMetaKey(key), &val, defaultVal))
	return val, ret
}

func (task *MppTask) MetaGetFrame(key MppMetaKey) (*MppFrame, MppRet) {
	var frame MppFrame
	ret := MppRet(C.mpp_task_meta_get_frame(*task.c, C.MppMetaKey(key), frame.c))
	return &frame, ret
}

func (task *MppTask) MetaGetPacket(key MppMetaKey) (*MppPacket, MppRet) {
	var packet MppPacket
	ret := MppRet(C.mpp_task_meta_get_packet(*task.c, C.MppMetaKey(key), packet.c))
	return &packet, ret
}

func (task *MppTask) MetaGetBuffer(key MppMetaKey) (MppBuffer, MppRet) {
	var buffer MppBuffer
	ret := MppRet(C.mpp_task_meta_get_buffer(*task.c, C.MppMetaKey(key), &buffer.c))
	return buffer, ret
}
