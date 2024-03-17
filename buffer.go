package rkcodec

//#include <rockchip/mpp_buffer.h>
import "C"

type MppBufferMode = C.int

const (
	MppBufferInternal = MppBufferMode(C.MPP_BUFFER_INTERNAL)
	MppBufferExternal = MppBufferMode(C.MPP_BUFFER_EXTERNAL)
	MppBufferModeButt = MppBufferMode(C.MPP_BUFFER_MODE_BUTT)
)

type MppBufferType = C.int

const (
	MppBufferTypeNormal  = MppBufferType(C.MPP_BUFFER_TYPE_NORMAL)
	MppBufferTypeIon     = MppBufferType(C.MPP_BUFFER_TYPE_ION)
	MppBufferTypeExtDma  = MppBufferType(C.MPP_BUFFER_TYPE_EXT_DMA)
	MppBufferTypeDrm     = MppBufferType(C.MPP_BUFFER_TYPE_DRM)
	MppBufferTypeDmaHeap = MppBufferType(C.MPP_BUFFER_TYPE_DMA_HEAP)
	MppBufferTypeButt    = MppBufferType(C.MPP_BUFFER_TYPE_BUTT)
)

const MppBufferTypeMask = C.MPP_BUFFER_TYPE_MASK

const MppBufferFlagsMask = C.MPP_BUFFER_FLAGS_MASK            // ROCKCHIP_BO_MASK << 16
const MppBufferFlagsContig = C.MPP_BUFFER_FLAGS_CONTIG        // ROCKCHIP_BO_CONTIG << 16
const MppBufferFlagsCachable = C.MPP_BUFFER_FLAGS_CACHABLE    // ROCKCHIP_BO_CACHABLE << 16
const MppBufferFlagsWC = C.MPP_BUFFER_FLAGS_WC                // ROCKCHIP_BO_WC << 16
const MppBufferFlagsSecure = C.MPP_BUFFER_FLAGS_SECURE        // ROCKCHIP_BO_SECURE << 16
const MppBufferFlagsAllocKmap = C.MPP_BUFFER_FLAGS_ALLOC_KMAP // ROCKCHIP_BO_ALLOC_KMAP << 16
const MppBufferFlagsDMA32 = C.MPP_BUFFER_FLAGS_DMA32          // ROCKCHIP_BO_DMA32 << 16

type MppBufferInfo struct {
	c *C.struct_MppBufferInfo
}

// const BufferGroupSizeDefault = BUFFER_GROUP_SIZE_DEFAULT // TODO: not found

// TODO: all function port
