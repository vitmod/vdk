package main

import (
	"github.com/vitmod/vdk/av/pktque"
	"github.com/vitmod/vdk/format"
	"github.com/vitmod/vdk/av/avutil"
	"github.com/vitmod/vdk/format/rtmp"
)

func init() {
	format.RegisterAll()
}

// as same as: ffmpeg -re -i projectindex.flv -c copy -f flv rtmp://localhost:1936/app/publish

func main() {
	file, _ := avutil.Open("projectindex.flv")
	conn, _ := rtmp.Dial("rtmp://localhost:1936/app/publish")
	// conn, _ := avutil.Create("rtmp://localhost:1936/app/publish")

	demuxer := &pktque.FilterDemuxer{Demuxer: file, Filter: &pktque.Walltime{}}
	avutil.CopyFile(conn, demuxer)

	file.Close()
	conn.Close()
}

