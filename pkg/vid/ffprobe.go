package vid

import "time"

type FFProbeJSON struct {
	Streams []FFStream `json:"streams"`
	Format  FFFormat   `json:"format"`
}

type FFDisposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
}

type FFTags struct {
	CreationTime     time.Time `json:"creation_time"`
	Language         string    `json:"language"`
	HandlerName      string    `json:"handler_name"`
	MajorBrand       string    `json:"major_brand"`
	MinorVersion     string    `json:"minor_version"`
	CompatibleBrands string    `json:"compatible_brands"`
	Encoder          string    `json:"encoder"`
}

type FFStream struct {
	Index            int           `json:"index"`
	CodecName        string        `json:"codec_name"`
	CodecLongName    string        `json:"codec_long_name"`
	Profile          string        `json:"profile"`
	CodecType        string        `json:"codec_type"`
	CodecTimeBase    string        `json:"codec_time_base"`
	CodecTagString   string        `json:"codec_tag_string"`
	CodecTag         string        `json:"codec_tag"`
	Width            int           `json:"width,omitempty"`
	Height           int           `json:"height,omitempty"`
	CodedWidth       int           `json:"coded_width,omitempty"`
	CodedHeight      int           `json:"coded_height,omitempty"`
	HasBFrames       int           `json:"has_b_frames,omitempty"`
	PixFmt           string        `json:"pix_fmt,omitempty"`
	Level            int           `json:"level,omitempty"`
	ColorRange       string        `json:"color_range,omitempty"`
	ColorSpace       string        `json:"color_space,omitempty"`
	ColorTransfer    string        `json:"color_transfer,omitempty"`
	ColorPrimaries   string        `json:"color_primaries,omitempty"`
	ChromaLocation   string        `json:"chroma_location,omitempty"`
	Refs             int           `json:"refs,omitempty"`
	IsAvc            string        `json:"is_avc,omitempty"`
	NalLengthSize    string        `json:"nal_length_size,omitempty"`
	RFrameRate       string        `json:"r_frame_rate"`
	AvgFrameRate     string        `json:"avg_frame_rate"`
	TimeBase         string        `json:"time_base"`
	StartPts         int           `json:"start_pts"`
	StartTime        string        `json:"start_time"`
	DurationTs       int           `json:"duration_ts"`
	Duration         string        `json:"duration"`
	BitRate          string        `json:"bit_rate"`
	BitsPerRawSample string        `json:"bits_per_raw_sample,omitempty"`
	NbFrames         string        `json:"nb_frames"`
	Disposition      FFDisposition `json:"disposition"`
	Tags             FFTags        `json:"tags"`
	SampleFmt        string        `json:"sample_fmt,omitempty"`
	SampleRate       string        `json:"sample_rate,omitempty"`
	Channels         int           `json:"channels,omitempty"`
	ChannelLayout    string        `json:"channel_layout,omitempty"`
	BitsPerSample    int           `json:"bits_per_sample,omitempty"`
	MaxBitRate       string        `json:"max_bit_rate,omitempty"`
}

type FFFormat struct {
	Filename       string `json:"filename"`
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      string `json:"start_time"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	Tags           FFTags `json:"tags"`
}
