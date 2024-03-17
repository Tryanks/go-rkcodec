package rkcodec

//#include <rockchip/mpp_err.h>
import "C"

type MppRet = C.int

const (
	MppSuccess = MppRet(C.MPP_SUCCESS)
	MppOK      = MppRet(C.MPP_OK)

	MppNOK         = MppRet(C.MPP_NOK)
	MppErrUnknow   = MppRet(C.MPP_ERR_UNKNOW)
	MppErrNullPtr  = MppRet(C.MPP_ERR_NULL_PTR)
	MppErrMalloc   = MppRet(C.MPP_ERR_MALLOC)
	MppErrOpenFile = MppRet(C.MPP_ERR_OPEN_FILE)
	MppErrValue    = MppRet(C.MPP_ERR_VALUE)
	MppErrReadBit  = MppRet(C.MPP_ERR_READ_BIT)
	MppErrTimeout  = MppRet(C.MPP_ERR_TIMEOUT)
	MppErrPerm     = MppRet(C.MPP_ERR_PERM)

	MppErrBase = MppRet(C.MPP_ERR_BASE)

	/* The error in stream processing */
	MppErrListStream    = MppRet(C.MPP_ERR_LIST_STREAM)
	MppErrInit          = MppRet(C.MPP_ERR_INIT)
	MppErrVpuCodecInit  = MppRet(C.MPP_ERR_VPU_CODEC_INIT)
	MppErrStream        = MppRet(C.MPP_ERR_STREAM)
	MppErrFatalThread   = MppRet(C.MPP_ERR_FATAL_THREAD)
	MppErrNoMem         = MppRet(C.MPP_ERR_NOMEM)
	MppErrProtol        = MppRet(C.MPP_ERR_PROTOL)
	MppFailSplitFrame   = MppRet(C.MPP_FAIL_SPLIT_FRAME)
	MppErrVpuHW         = MppRet(C.MPP_ERR_VPUHW)
	MppEosStreamReached = MppRet(C.MPP_EOS_STREAM_REACHED)
	MppErrBufferFull    = MppRet(C.MPP_ERR_BUFFER_FULL)
	MppErrDisplayFull   = MppRet(C.MPP_ERR_DISPLAY_FULL)
)

func (r MppRet) String() string {
	switch MppRet(r) {
	case MppSuccess: // MppOK
		return "MPP_SUCCESS" // "MPP_OK"
	case MppNOK:
		return "MPP_NOK"
	case MppErrUnknow:
		return "MPP_ERR_UNKNOW"
	case MppErrNullPtr:
		return "MPP_ERR_NULL_PTR"
	case MppErrMalloc:
		return "MPP_ERR_MALLOC"
	case MppErrOpenFile:
		return "MPP_ERR_OPEN_FILE"
	case MppErrValue:
		return "MPP_ERR_VALUE"
	case MppErrReadBit:
		return "MPP_ERR_READ_BIT"
	case MppErrTimeout:
		return "MPP_ERR_TIMEOUT"
	case MppErrPerm:
		return "MPP_ERR_PERM"
	case MppErrBase:
		return "MPP_ERR_BASE"
	case MppErrListStream:
		return "MPP_ERR_LIST_STREAM"
	case MppErrInit:
		return "MPP_ERR_INIT"
	case MppErrVpuCodecInit:
		return "MPP_ERR_VPU_CODEC_INIT"
	case MppErrStream:
		return "MPP_ERR_STREAM"
	case MppErrFatalThread:
		return "MPP_ERR_FATAL_THREAD"
	case MppErrNoMem:
		return "MPP_ERR_NOMEM"
	case MppErrProtol:
		return "MPP_ERR_PROTOL"
	case MppFailSplitFrame:
		return "MPP_FAIL_SPLIT_FRAME"
	case MppErrVpuHW:
		return "MPP_ERR_VPUHW"
	case MppEosStreamReached:
		return "MPP_EOS_STREAM_REACHED"
	case MppErrBufferFull:
		return "MPP_ERR_BUFFER_FULL"
	case MppErrDisplayFull:
		return "MPP_ERR_DISPLAY_FULL"
	default:
		return "MPP_ERR_UNKNOWN"
	}
}

func (r MppRet) Error() string {
	return r.String()
}
