package model

type Stream struct {
	DeviceID              string `json:"DeviceID"`  //设备编号
	ChannelID             string `json:"ChannelID"` //通道编号
	AudioEnable           bool   `json:"AudioEnable"`
	CDN                   string `json:"CDN"`
	CascadeSize           int    `json:"CascadeSize"`
	ChannelCustomName     string `json:"ChannelCustomName"`
	ChannelName           string `json:"ChannelName"`
	ChannelPTZType        int    `json:"ChannelPTZType"`
	Duration              int    `json:"Duration"`
	FLV                   string `json:"FLV"`
	HLS                   string `json:"HLS"`
	InBitRate             int    `json:"InBitRate"`
	InBytes               int    `json:"InBytes"`
	NumOutputs            int    `json:"NumOutputs"`
	Ondemand              bool   `json:"Ondemand"`
	OutBytes              int    `json:"OutBytes"`
	RTMP                  string `json:"RTMP"`
	RTSP                  string `json:"RTSP"`
	RecordStartAt         string `json:"RecordStartAt"`
	RelaySize             int    `json:"RelaySize"`
	SnapURL               string `json:"SnapURL"`
	SourceAudioCodecName  string `json:"SourceAudioCodecName"`
	SourceAudioSampleRate int    `json:"SourceAudioSampleRate"`
	SourceVideoCodecName  string `json:"SourceVideoCodecName"`
	SourceVideoFrameRate  int    `json:"SourceVideoFrameRate"`
	SourceVideoHeight     int    `json:"SourceVideoHeight"`
	SourceVideoWidth      int    `json:"SourceVideoWidth"`
	StartAt               string `json:"StartAt"`
	StreamID              string `json:"StreamID"`
	Transport             string `json:"Transport"`
	WSFLV                 string `json:"WS_FLV"`
}
