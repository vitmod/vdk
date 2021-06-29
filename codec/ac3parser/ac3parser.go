package ac3parser

import (
	"fmt"
	"time"

	"github.com/vitmod/vdk/av"
)

type CodecData struct {
	CodecType_		av.CodecType
	SampleRate_		int
	SampleFormat_	av.SampleFormat
	ChannelLayout_	av.ChannelLayout
}

var (
	A52HalfRate = []int{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, }
	
	A52BitRates = []int{
		32, 40, 48, 56, 64, 80, 96, 112, 128, 160,
		192, 224, 256, 320, 384, 448, 512, 576, 640,
	}

	/* Only store frame sizes for 44.1KHz - others are simply multiples
	   of the bitrate */
	A52FrameSizes441 = []int{
		69, 70, 87, 88, 104, 105, 121, 122, 139, 140, 174, 175, 208, 209,
		243, 244, 278, 279, 348, 349, 417, 418, 487, 488, 557, 558, 696, 697,
		835, 836, 975, 976, 1114, 1115, 1253, 1254, 1393, 1394,
	}
)


var spf int
var bitrate int
var samplerate int

func IsValidAC3FrameHeader(header []byte) bool {
	if header[0] != 0x0B || header[1] != 0x77 {
		return false
	}
	
	return true
}

func ParseAC3(frame []byte) (spf_ int, framelen_ int, samplerate_ int) {
	i := frame[4] & 0x3E

	if i > 36 {
		fmt.Printf("AC3: Invalid frmsizecod: %v\n", i)
		return
	}
	
	r := frame[5] >> 3

	rate_shift := A52HalfRate[r]
	base_bitrate := A52BitRates[i>>1]
	channels := 2
	
	switch frame[4] & 0xC0 {
	case 0x00:
		samplerate_ = 48000
		spf_ = base_bitrate * 2 * channels
	case 0x40:
		samplerate_ = 44100
		spf_ = A52FrameSizes441[i] * channels
	case 0x80:
		samplerate_ = 32000
		spf_ = base_bitrate * 3 * channels
	default:
		fmt.Printf("AC3: Invalid samplerate code: 0x%02x\n", frame[4]&0xc0)
	}
	
	spf = spf_
	samplerate = samplerate_
	bitrate = (base_bitrate * 1000) >> rate_shift
	
	framelen_ = spf_
	
	//fmt.Printf("AC3: spf: %v, framelen: %v, samplerate: %v, bitrate: %v\n", spf_, framelen_, samplerate_, bitrate)
	
	return
}

func (self CodecData) Type() av.CodecType {
	return av.AC3
}

func (self CodecData) SampleFormat() av.SampleFormat {
	return av.FLTP
}

func (self CodecData) ChannelLayout() av.ChannelLayout {
	return av.CH_STEREO
}

func (self CodecData) SampleRate() int {
	var samplerate_ = 48000
	if samplerate > 0 {
		samplerate_ = samplerate
	}
	return samplerate_
}

func (self CodecData) PacketDuration(data []byte) (dur time.Duration, err error) {
	var spf_ = 1536
	if spf > 0 {
		spf_ = spf
	}
	dur = time.Duration(spf_) * time.Second / time.Duration(self.SampleRate())
	return
}

func NewCodecDataAC3() (self CodecData) {
	self.CodecType_ = self.Type()
	self.SampleRate_ = self.SampleRate()
	self.SampleFormat_ = self.SampleFormat()
	self.ChannelLayout_ = self.ChannelLayout()
	return
}
