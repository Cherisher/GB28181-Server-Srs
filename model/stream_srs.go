//+build srs
package model

import "github.com/Cherisher/GB28181-Server-Srs/config"

//Start 获取拉流地址
func (s *Stream) Start() error {
	s.StreamID = "stream:" + s.DeviceID + ":" + s.ChannelID
	s.FLV = config.Config().Srs.FlvBase + "/live/" + s.DeviceID + "@" + s.ChannelID + ".flv"
	s.RTMP = config.Config().Srs.RtmpBase + "/live/" + s.DeviceID + "@" + s.ChannelID
	return nil
}
