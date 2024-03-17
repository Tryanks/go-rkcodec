package rkcodec

import (
	"testing"
)

func TestType(t *testing.T) {
	t.Log(MppCtxEnc)
	t.Log(MppSuccess)
	t.Log(MppErrUnknow)
	t.Log(MppBufferTypeButt)
	t.Log(MppBufferFlagsDMA32)
	t.Log(MppFramePrimariesNB)
	t.Log(MetaKeyDecTbnUvOffset)
	t.Log(HdrFormatMax)
	t.Log(MpiCmdButt)
	t.Log(MppTimeoutMax)
}
