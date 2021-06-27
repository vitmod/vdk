package format

import (
	"github.com/vitmod/vdk/format/mp4"
	"github.com/vitmod/vdk/format/ts"
	"github.com/vitmod/vdk/format/rtmp"
	"github.com/vitmod/vdk/format/rtsp"
	"github.com/vitmod/vdk/format/flv"
	"github.com/vitmod/vdk/format/aac"
	"github.com/vitmod/vdk/av/avutil"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}

