# Go rkcodec

A cgo package for [rkmpp](https://github.com/rockchip-linux/mpp) media library.

Need **rkmpp** enviorment.

```shell
pkg-config --cflags rockchip_mpp
```

# Usage

## Install
```shell
go get -u github.com/Tryanks/go-rkcodec
```

## Code
```go
package main

import "C"
import "rkcodec"

func main() {
	decoder := rkcodec.NewMppCodec()
	decoder.Control(rkcodec.MppDecSetParserSplitMode, C.int(1))
	decoder.Init(rkcodec.MppCtxDec, rkcodec.MppCodingAVC)

	frame, err := rkcodec.MppFrameInit()
	if err != rkcodec.MppSuccess {
		panic(err)
	}
	defer frame.Deinit()

	packet, err := rkcodec.NewMppPacket()
	if err != rkcodec.MppSuccess {
		panic(err)
	}

	h264NALU := make([]byte, 1024) // H.264 NALU data

	packet.SetData(h264NALU)

	err = decoder.DecodePutPacket(*packet)
	if err != rkcodec.MppSuccess {
		panic(err)
	}

	err = decoder.DecodeGetFrame(frame)
	if err != rkcodec.MppSuccess {
		panic(err)
	}

	// Do something with the frame
}

```

# Dev Reference

[Rockchip_Developer_Guide_MPP](https://github.com/rockchip-linux/mpp/blob/develop/doc/Rockchip_Developer_Guide_MPP_EN.md)
