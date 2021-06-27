package main

import (
	"fmt"
	"strings"
	"github.com/vitmod/vdk/format"
	"github.com/vitmod/vdk/av/avutil"
	"github.com/vitmod/vdk/format/rtmp"
)

func init() {
	format.RegisterAll()
}

func main() {
	server := &rtmp.Server{}

	server.HandlePlay = func(conn *rtmp.Conn) {
		segs := strings.Split(conn.URL.Path, "/")
		url := fmt.Sprintf("%s://%s", segs[1], strings.Join(segs[2:], "/"))
		src, _ := avutil.Open(url)
		avutil.CopyFile(conn, src)
	}

	server.ListenAndServe()

	// ffplay rtmp://localhost/rtsp/192.168.1.1/camera1
	// ffplay rtmp://localhost/rtmp/live.hkstv.hk.lxdns.com/live/hks
}
