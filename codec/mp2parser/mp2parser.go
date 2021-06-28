package mp2parser

import (
	//"fmt"
	"time"

	"github.com/vitmod/vdk/av"
)

type CodecData struct {
	CodecType_		av.CodecType
	SampleRate_		int
	SampleFormat_	av.SampleFormat
	ChannelLayout_	av.ChannelLayout
}

var srtable = [...]uint32{
	44100, 48000, 32000, 0, // mpeg1
	22050, 24000, 16000, 0, // mpeg2
	11025, 12000, 8000, 0} // mpeg2.5
		
var brtable = [...]uint32{
	0, 32, 64, 96, 128, 160, 192, 224, 256, 288, 320, 352, 384, 416, 448, 0,
	0, 32, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384, 0,
	0, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 0,
	0, 32, 48, 56, 64, 80, 96, 112, 128, 144, 160, 176, 192, 224, 256, 0,
	0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160, 0}

func IsValidFrameHeader(header []byte) bool {
	if (header[0] != 0x0FF) && ((header[1] & 0x0E0) != 0x0E0) {
		return false
	}

	// get and check the mpeg version
	mpegver := (uint16(header[1]) & 0x18) >> 3
	if mpegver == 1 || mpegver > 3 {
		return false
	}

	// get and check mpeg layer
	layer := (header[1] & 0x06) >> 1
	if layer == 0 || layer > 3 {
		return false
	}

	// get and check bitreate index
	brindex := (header[2] & 0x0F0) >> 4
	if brindex > 15 {
		return false
	}

	// get and check the 'sampling_rate_index':
	srindex := (header[2] & 0x0C) >> 2
	if srindex >= 3 {
		return false
	}

	return true
}

func GetSPF(header []byte) int {
	// get and check the mpeg version
	mpegver := byte((header[1] & 0x18) >> 3)

	// get and check mpeg layer
	layer := byte((header[1] & 0x06) >> 1)

	spf := 0
	switch mpegver {
	case 3: // mpeg 1
		if layer == 3 { // layer1
			spf = 384
		} else {
			spf = 1152 // layer2 & layer3
		}
	case 2, 0: // mpeg2 & mpeg2.5
		switch layer {
		case 3: // layer1
			spf = 384
		case 2: // layer2
			spf = 1152
		default:
			spf = 576 // layer 3
		}
	}
	return spf
}

func GetSR(header []byte) int {
	sr := 0
	// get and check the 'sampling_rate_index':
	srindex := (header[2] & 0x0C) >> 2
	if srtable[srindex] == 0 {
		return 0
	}

	// get and check the mpeg version
	mpegver := byte((header[1] & 0x018) >> 3)
	if mpegver == 1 {
		return 0
	}

	if mpegver == 3 {
		// mpeg1
		sr = int(srtable[srindex])
	}
	if mpegver == 2 {
		// mpeg2
		sr = int(srtable[srindex+4])
	}
	if mpegver == 0 {
		// mpeg2.5
		sr = int(srtable[srindex+8])
	}
	return sr
}

func GetFrameSize(header []byte) int {
	var sr, bitrate uint32
	var res int

	// get and check the mpeg version
	mpegver := byte((header[1] & 0x18) >> 3)
	if mpegver == 1 || mpegver > 3 {
		return 0
	}

	// get and check mpeg layer
	layer := byte((header[1] & 0x06) >> 1)
	if layer == 0 || layer > 3 {
		return 0
	}

	brindex := byte((header[2] & 0x0F0) >> 4)

	if mpegver == 3 && layer == 3 {
		// mpeg1, layer1
		bitrate = brtable[brindex]
	}
	if mpegver == 3 && layer == 2 {
		// mpeg1, layer2
		bitrate = brtable[brindex+16]
	}
	if mpegver == 3 && layer == 1 {
		// mpeg1, layer3
		bitrate = brtable[brindex+32]
	}
	if (mpegver == 2 || mpegver == 0) && layer == 3 {
		// mpeg2, 2.5, layer1
		bitrate = brtable[brindex+48]
	}
	if (mpegver == 2 || mpegver == 0) && (layer == 2 || layer == 1) {
		//mpeg2, layer2 or layer3
		bitrate = brtable[brindex+64]
	}
	bitrate = bitrate * 1000
	padding := int(header[2]&0x02) >> 1

	// get and check the 'sampling_rate_index':
	srindex := byte((header[2] & 0x0C) >> 2)
	if srindex >= 3 {
		return 0
	}
	if mpegver == 3 {
		sr = srtable[srindex]
	}
	if mpegver == 2 {
		sr = srtable[srindex+4]
	}
	if mpegver == 0 {
		sr = srtable[srindex+8]
	}

	switch mpegver {
	case 3: // mpeg1
		if layer == 3 { // layer1
			res = (int(12*bitrate/sr) * 4) + (padding * 4)
		}
		if layer == 2 || layer == 1 {
			// layer 2 & 3
			res = int(144*bitrate/sr) + padding
		}

	case 2, 0: //mpeg2, mpeg2.5
		if layer == 3 { // layer1
			res = (int(12*bitrate/sr) * 4) + (padding * 4)
		}
		if layer == 2 { // layer2
			res = int(144*bitrate/sr) + padding
		}
		if layer == 1 { // layer3
			res = int(72*bitrate/sr) + padding
		}
	}
	return res
}

var spf_ int
var samplerate_ int

//func ParseMP2(frame []byte) (spf int, framelen int, samplerate int, err error) {
func ParseMP2(frame []byte) (spf int, framelen int, samplerate int) {	
	spf = GetSPF(frame)
	spf_ = spf
	framelen = GetFrameSize(frame)
	samplerate = GetSR(frame)
	samplerate_ = samplerate
	return
}

func (self CodecData) Type() av.CodecType {
	return av.MPEG2AUDIO
}

func (self CodecData) SampleFormat() av.SampleFormat {
	return av.S16
}

func (self CodecData) ChannelLayout() av.ChannelLayout {
	return av.CH_STEREO
}

func (self CodecData) SampleRate() int {
	var sr int = 48000
	if samplerate_ > 0 {
		sr = samplerate_
	}
	return sr
}

func (self CodecData) PacketDuration(data []byte) (dur time.Duration, err error) {
	var spf int = 1152
	if spf_ > 0 {
		spf = spf_
	}
	dur = time.Duration(spf) * time.Second / time.Duration(self.SampleRate())
	return
}

func NewCodecDataMP2Audio() (self CodecData) {
	var sr int = 48000
	if samplerate_ > 0 {
		sr = samplerate_
	}
	self.CodecType_ = av.MPEG2AUDIO
	self.SampleRate_ = sr
	self.SampleFormat_ = av.S16
	self.ChannelLayout_ = av.CH_STEREO
	return
}
