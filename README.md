# PLEASE USE joy5 INSTEAD

[joy5](https://github.com/vitmod/joy5)

- High performance Copy-on-write gop cache [code](https://github.com/vitmod/joy5/blob/master/cmd/avtool/pubsub.go)
- Better av.Packet design [code](https://github.com/vitmod/joy5/blob/master/av/av.go)

# JOY4

> Golang audio/video library and streaming server

JOY4 is powerful library written in golang, well-designed interface makes a few lines of code can do a lot of things such as reading, writing, transcoding among variety media formats, or setting up high-performance live streaming server.

# Features 

Well-designed and easy-to-use interfaces:

- Muxer / Demuxer ([doc](https://godoc.org/github.com/vitmod/vdk/av#Demuxer) [example](https://github.com/vitmod/vdk/blob/master/examples/open_probe_file/main.go))
- Audio Decoder ([doc](https://godoc.org/github.com/vitmod/vdk/av#AudioDecoder) [example](https://github.com/vitmod/vdk/blob/master/examples/audio_decode/main.go))
- Transcoding ([doc](https://godoc.org/github.com/vitmod/vdk/av/transcode) [example](https://github.com/vitmod/vdk/blob/master/examples/transcode/main.go))
- Streaming server ([example](https://github.com/vitmod/vdk/blob/master/examples/http_flv_and_rtmp_server/main.go))

Support container formats:

- MP4
- MPEG-TS
- FLV
- AAC (ADTS)

RTSP Client
- High level camera bug tolerance
- Support STAP-A

RTMP Client
- Support publishing to nginx-rtmp-server
- Support playing

RTMP / HTTP-FLV Server 
- Support publishing clients: OBS / ffmpeg / Flash Player (>8)
- Support playing clients: Flash Player 11 / VLC / ffplay / mpv
- High performance


Publisher-subscriber packet buffer queue ([doc](https://godoc.org/github.com/vitmod/vdk/av/pubsub))

- Customize publisher buffer time and subscriber read position


- Multiple channels live streaming ([example](https://github.com/vitmod/vdk/blob/master/examples/rtmp_server_channels/main.go))

Packet filters ([doc](https://godoc.org/github.com/vitmod/vdk/av/pktque))

- Wait first keyframe
- Fix timestamp
- Make A/V sync
- Customize ([example](https://github.com/vitmod/vdk/blob/master/examples/rtmp_server_channels/main.go#L19))

FFMPEG Golang-style binding ([doc](https://godoc.org/github.com/vitmod/vdk/cgo/ffmpeg))
- Audio Encoder / Decoder
- Video Decoder
- Audio Resampler

Support codec and container parsers:

- H264 SPS/PPS/AVCDecoderConfigure parser ([doc](https://godoc.org/github.com/vitmod/vdk/codec/h264parser))
- AAC ADTSHeader/MPEG4AudioConfig parser ([doc](https://godoc.org/github.com/vitmod/vdk/codec/aacparser))
- MP4 Atoms parser ([doc](https://godoc.org/github.com/vitmod/vdk/format/mp4/mp4io))
- FLV AMF0 object parser ([doc](https://godoc.org/github.com/vitmod/vdk/format/flv/flvio))

# Requirements

Go version >= 1.6

ffmpeg version >= 3.0 (optional)

# TODO

HLS / MPEG-DASH Server

ffmpeg.VideoEncoder / ffmpeg.SWScale

# License

MIT
