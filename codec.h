#ifndef __CODEC_H__
#define __CODEC_H__

#include <rockchip/rk_mpi.h>
#include <stdlib.h>

typedef struct MppCodec {
    MppCtx ctx;
    MppApi *api;
} MppCodec;

static inline MppCodec *mpp_ctx_api_alloc() {
    MppCtx ctx          = NULL;
    MppApi *mpi         = NULL;
    mpp_create(&ctx, &mpi);
    MppCodec *codec;
    codec = (MppCodec *)malloc(sizeof(MppCodec));
    codec->ctx = ctx;
    codec->api = mpi;
    return codec;
}

static inline MPP_RET codec_init(MppCodec *codec, MppCtxType type, MppCodingType coding) {
    return mpp_init(codec->ctx, type, coding);
}

static inline MPP_RET codec_destroy(MppCodec *codec) {
    MPP_RET ret = mpp_destroy(codec->ctx);
    free(codec);
    return ret;
}

static inline RK_U32 codec_size(MppCodec *codec) {
    return codec->api->size;
}

static inline RK_U32 codec_version(MppCodec *codec) {
    return codec->api->version;
}

static inline MPP_RET codec_decode(MppCodec *codec, MppPacket packet, MppFrame *frame) {
    return codec->api->decode(codec->ctx, packet, frame);
}

static inline MPP_RET codec_decode_put_packet(MppCodec *codec, MppPacket packet) {
    return codec->api->decode_put_packet(codec->ctx, packet);
}

static inline MPP_RET codec_decode_get_frame(MppCodec *codec, MppFrame *frame) {
    return codec->api->decode_get_frame(codec->ctx, frame);
}

static inline MPP_RET codec_encode(MppCodec *codec, MppFrame frame, MppPacket *packet) {
    return codec->api->encode(codec->ctx, frame, packet);
}

static inline MPP_RET codec_encode_put_frame(MppCodec *codec, MppFrame frame) {
    return codec->api->encode_put_frame(codec->ctx, frame);
}

static inline MPP_RET codec_encode_get_packet(MppCodec *codec, MppPacket *packet) {
    return codec->api->encode_get_packet(codec->ctx, packet);
}

static inline MPP_RET codec_isp(MppCodec *codec, MppFrame dst, MppFrame src) {
    return codec->api->isp(codec->ctx, dst, src);
}

static inline MPP_RET codec_isp_put_frame(MppCodec *codec, MppFrame frame) {
    return codec->api->isp_put_frame(codec->ctx, frame);
}

static inline MPP_RET codec_isp_get_frame(MppCodec *codec, MppFrame *frame) {
    return codec->api->isp_get_frame(codec->ctx, frame);
}

static inline MPP_RET codec_poll(MppCodec *codec, MppPortType type, MppPollType timeout) {
    return codec->api->poll(codec->ctx, type, timeout);
}

static inline MPP_RET codec_dequeue(MppCodec *codec, MppPortType type, MppTask *task) {
    return codec->api->dequeue(codec->ctx, type, task);
}

static inline MPP_RET codec_enqueue(MppCodec *codec, MppPortType type, MppTask task) {
    return codec->api->enqueue(codec->ctx, type, task);
}

static inline MPP_RET codec_reset(MppCodec *codec) {
    return codec->api->reset(codec->ctx);
}

static inline MPP_RET codec_control(MppCodec *codec, MpiCmd cmd, void* param) {
    return codec->api->control(codec->ctx, cmd, param);
}

#endif /* __CODEC_H__ */