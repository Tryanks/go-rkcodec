package rkcodec

//#include <rockchip/rk_mpi_cmd.h>
import "C"

type MpiCmd = C.MpiCmd

const (
	MppOsalCmdBase = MpiCmd(C.MPP_OSAL_CMD_BASE)
	MppOsalCmdEnd  = MpiCmd(C.MPP_OSAL_CMD_END)

	MppCmdBase               = MpiCmd(C.MPP_CMD_BASE)
	MppEnableDeinterlace     = MpiCmd(C.MPP_ENABLE_DEINTERLACE)
	MppSetInputBlock         = MpiCmd(C.MPP_SET_INPUT_BLOCK)          // deprecated
	MppSetInputBlockTimeout  = MpiCmd(C.MPP_SET_INTPUT_BLOCK_TIMEOUT) // deprecated
	MppSetOutputBlock        = MpiCmd(C.MPP_SET_OUTPUT_BLOCK)         // deprecated
	MppSetOutputBlockTimeout = MpiCmd(C.MPP_SET_OUTPUT_BLOCK_TIMEOUT) // deprecated

	// timeout setup, refer to  MPP_TIMEOUT_XXX
	MppSetInputTimeout  = MpiCmd(C.MPP_SET_INPUT_TIMEOUT)  // parameter type RK_S64
	MppSetOutputTimeout = MpiCmd(C.MPP_SET_OUTPUT_TIMEOUT) // parameter type RK_S64
	MppSetDisableThread = MpiCmd(C.MPP_SET_DISABLE_THREAD) // MPP no thread mode and use external thread to decode

	MppStateCmdBase = MpiCmd(C.MPP_STATE_CMD_BASE)
	MppStart        = MpiCmd(C.MPP_START)
	MppStop         = MpiCmd(C.MPP_STOP)
	MppPause        = MpiCmd(C.MPP_PAUSE)
	MppResume       = MpiCmd(C.MPP_RESUME)

	MppCmdEnd = MpiCmd(C.MPP_CMD_END)

	MppCodecCmdBase      = MpiCmd(C.MPP_CODEC_CMD_BASE)
	MppCodecGetFrameInfo = MpiCmd(C.MPP_CODEC_GET_FRAME_INFO)
	MppCodecCmdEnd       = MpiCmd(C.MPP_CODEC_CMD_END)

	MppDecCmdBase              = MpiCmd(C.MPP_DEC_CMD_BASE)
	MppDecSetFrameInfo         = MpiCmd(C.MPP_DEC_SET_FRAME_INFO)    // vpu api legacy control for buffer slot dimension init
	MppDecSetExtBufGroup       = MpiCmd(C.MPP_DEC_SET_EXT_BUF_GROUP) // IMPORTANT: set external buffer group to mpp decoder
	MppDecSetInfoChangeReady   = MpiCmd(C.MPP_DEC_SET_INFO_CHANGE_READY)
	MppDecSetPresentTimeOrder  = MpiCmd(C.MPP_DEC_SET_PRESENT_TIME_ORDER) // use input time order for output
	MppDecSetParserSplitMode   = MpiCmd(C.MPP_DEC_SET_PARSER_SPLIT_MODE)  // Need to setup before init
	MppDecSetParserFastMode    = MpiCmd(C.MPP_DEC_SET_PARSER_FAST_MODE)   // Need to setup before init
	MppDecGetStreamCount       = MpiCmd(C.MPP_DEC_GET_STREAM_COUNT)
	MppDecGetVpumemUsedCount   = MpiCmd(C.MPP_DEC_GET_VPUMEM_USED_COUNT)
	MppDecSetVc1ExtraData      = MpiCmd(C.MPP_DEC_SET_VC1_EXTRA_DATA)
	MppDecSetOutputFormat      = MpiCmd(C.MPP_DEC_SET_OUTPUT_FORMAT)
	MppDecSetDisableError      = MpiCmd(C.MPP_DEC_SET_DISABLE_ERROR) // When set it will disable sw/hw error (H.264 / H.265)
	MppDecSetImmediateOut      = MpiCmd(C.MPP_DEC_SET_IMMEDIATE_OUT)
	MppDecSetEnableDeinterlace = MpiCmd(C.MPP_DEC_SET_ENABLE_DEINTERLACE) // MPP enable deinterlace by default. Vpuapi can disable it
	MppDecSetEnableFastPlay    = MpiCmd(C.MPP_DEC_SET_ENABLE_FAST_PLAY)   // enable idr output immediately
	MppDecSetDisableThread     = MpiCmd(C.MPP_DEC_SET_DISABLE_THREAD)     // MPP no thread mode and use external thread to decode
	MppDecSetMaxUseBufferSize  = MpiCmd(C.MPP_DEC_SET_MAX_USE_BUFFER_SIZE)
	MppDecSetEnableMVC         = MpiCmd(C.MPP_DEC_SET_ENABLE_MVC) // enable MVC decoding

	MppDecCmdQuery = MpiCmd(C.MPP_DEC_CMD_QUERY) // query decoder runtime information for decode stage
	// query decoder runtime information for decode stage
	MppDecQuery = MpiCmd(C.MPP_DEC_QUERY) // set and get MppDecQueryCfg structure

	CmdDecCmdCfg = MpiCmd(C.CMD_DEC_CMD_CFG)
	MppDecSetCfg = MpiCmd(C.MPP_DEC_SET_CFG) // set MppDecCfg structure
	MppDecGetCfg = MpiCmd(C.MPP_DEC_GET_CFG) // get MppDecCfg structure

	MppDecCmdEnd = MpiCmd(C.MPP_DEC_CMD_END)

	MppEncCmdBase = MpiCmd(C.MPP_ENC_CMD_BASE)
	// basic encoder setup control
	MppEncSetCfg      = MpiCmd(C.MPP_ENC_SET_CFG)       // set MppEncCfg structure
	MppEncGetCfg      = MpiCmd(C.MPP_ENC_GET_CFG)       // get MppEncCfg structure
	MppEncSetPrepCfg  = MpiCmd(C.MPP_ENC_SET_PREP_CFG)  // deprecated set MppEncPrepCfg structure, use MPP_ENC_SET_CFG instead
	MppEncGetPrepCfg  = MpiCmd(C.MPP_ENC_GET_PREP_CFG)  // deprecated get MppEncPrepCfg structure, use MPP_ENC_GET_CFG instead
	MppEncSetRcCfg    = MpiCmd(C.MPP_ENC_SET_RC_CFG)    // deprecated set MppEncRcCfg structure, use MPP_ENC_SET_CFG instead
	MppEncGetRcCfg    = MpiCmd(C.MPP_ENC_GET_RC_CFG)    // deprecated get MppEncRcCfg structure, use MPP_ENC_GET_CFG instead
	MppEncSetCodecCfg = MpiCmd(C.MPP_ENC_SET_CODEC_CFG) // deprecated set MppEncCodecCfg structure, use MPP_ENC_SET_CFG instead
	MppEncGetCodecCfg = MpiCmd(C.MPP_ENC_GET_CODEC_CFG) // deprecated get MppEncCodecCfg structure, use MPP_ENC_GET_CFG instead
	// runtime encoder setup control
	MppEncSetIdrFrame   = MpiCmd(C.MPP_ENC_SET_IDR_FRAME)    // next frame will be encoded as intra frame
	MppEncSetOsdLegacy0 = MpiCmd(C.MPP_ENC_SET_OSD_LEGACY_0) // deprecated
	MppEncSetOsdLegacy1 = MpiCmd(C.MPP_ENC_SET_OSD_LEGACY_1) // deprecated
	MppEncSetOsdLegacy2 = MpiCmd(C.MPP_ENC_SET_OSD_LEGACY_2) // deprecated
	MppEncGetHdrSync    = MpiCmd(C.MPP_ENC_GET_HDR_SYNC)     // get vps / sps / pps which has better sync behavior parameter is MppPacket
	MppEncGetExtraInfo  = MpiCmd(C.MPP_ENC_GET_EXTRA_INFO)   // deprecated
	MppEncSetSeiCfg     = MpiCmd(C.MPP_ENC_SET_SEI_CFG)      // SEI: Supplement Enhancemant Information, parameter is MppSeiMode
	MppEncGetSeiData    = MpiCmd(C.MPP_ENC_GET_SEI_DATA)     // SEI: Supplement Enhancemant Information, parameter is MppPacket
	MppEncPreAllocBuff  = MpiCmd(C.MPP_ENC_PRE_ALLOC_BUFF)   // deprecated
	MppEncSetQpRange    = MpiCmd(C.MPP_ENC_SET_QP_RANGE)     // used for adjusting qp range, the parameter can be 1 or 2
	MppEncSetRoiCfg     = MpiCmd(C.MPP_ENC_SET_ROI_CFG)      // set MppEncROICfg structure
	MppEncSetCtuQp      = MpiCmd(C.MPP_ENC_SET_CTU_QP)       // for H265 Encoder,set CTU's size and QP

	MppEncCmdQuery = MpiCmd(C.MPP_ENC_CMD_QUERY) // query encoder runtime information for encode stage
	// query encoder runtime information for encode stage
	MppEncQuery = MpiCmd(C.MPP_ENC_QUERY) // set and get MppEncQueryCfg structure

	// User define rate control stategy API control
	MppEncCfgRcApi = MpiCmd(C.MPP_ENC_CFG_RC_API)
	// Get RcApiQueryAll structure
	MppEncGetRcApiAll = MpiCmd(C.MPP_ENC_GET_RC_API_ALL)
	// Get RcApiQueryType structure
	MppEncGetRcApiByType = MpiCmd(C.MPP_ENC_GET_RC_API_BY_TYPE)
	// Set RcImplApi structure
	MppEncSetRcApiCfg = MpiCmd(C.MPP_ENC_SET_RC_API_CFG)
	// Get RcApiBrief structure
	MppEncGetRcApiCurrent = MpiCmd(C.MPP_ENC_GET_RC_API_CURRENT)
	// Set RcApiBrief structure
	MppEncSetRcApiCurrent = MpiCmd(C.MPP_ENC_SET_RC_API_CURRENT)

	MppEncCfgMisc       = MpiCmd(C.MPP_ENC_CFG_MISC)
	MppEncSetHeaderMode = MpiCmd(C.MPP_ENC_SET_HEADER_MODE) // set MppEncHeaderMode
	MppEncGetHeaderMode = MpiCmd(C.MPP_ENC_GET_HEADER_MODE) // get MppEncHeaderMode

	MppEncCfgSplit = MpiCmd(C.MPP_ENC_CFG_SPLIT)
	MppEncSetSplit = MpiCmd(C.MPP_ENC_SET_SPLIT) // set MppEncSliceSplit structure
	MppEncGetSplit = MpiCmd(C.MPP_ENC_GET_SPLIT) // get MppEncSliceSplit structure

	MppEncCfgRef    = MpiCmd(C.MPP_ENC_CFG_REF)
	MppEncSetRefCfg = MpiCmd(C.MPP_ENC_SET_REF_CFG) // set MppEncRefCfg structure

	MppEncCfgOsd        = MpiCmd(C.MPP_ENC_CFG_OSD)
	MppEncSetOsdPltCfg  = MpiCmd(C.MPP_ENC_SET_OSD_PLT_CFG)  // set OSD palette, parameter should be pointer to MppEncOSDPltCfg
	MppEncGetOsdPltCfg  = MpiCmd(C.MPP_ENC_GET_OSD_PLT_CFG)  // get OSD palette, parameter should be pointer to MppEncOSDPltCfg
	MppEncSetOsdDataCfg = MpiCmd(C.MPP_ENC_SET_OSD_DATA_CFG) // set OSD data with at most 8 regions, parameter should be pointer to MppEncOSDData

	MppEncCmdEnd = MpiCmd(C.MPP_ENC_CMD_END)

	MppIspCmdBase = MpiCmd(C.MPP_ISP_CMD_BASE)
	MppIspCmdEnd  = MpiCmd(C.MPP_ISP_CMD_END)

	MppHalCmdBase = MpiCmd(C.MPP_HAL_CMD_BASE)
	MppHalCmdEnd  = MpiCmd(C.MPP_HAL_CMD_END)

	MpiCmdButt = MpiCmd(C.MPI_CMD_BUTT)
)
