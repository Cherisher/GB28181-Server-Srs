//+build srs
package model

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Cherisher/GB28181-Server-Srs/config"
)

func (s *Stream) srsStreamAction(action string) error {
	URL := config.Config().Srs.APIBase + "/api/v1/gb28181?action=" + action + "&id=" + s.DeviceID + "&chid=" + s.ChannelID

	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

//Start 获取拉流地址
func (s *Stream) Start() error {
	s.srsStreamAction("sip_invite")
	s.StreamID = "stream:" + s.DeviceID + ":" + s.ChannelID
	s.FLV = config.Config().Srs.FlvBase + "/live/" + s.DeviceID + "@" + s.ChannelID + ".flv"
	s.RTMP = config.Config().Srs.RtmpBase + "/live/" + s.DeviceID + "@" + s.ChannelID
	return nil
}

//Stop 停止BYE 获取拉流地址
func (s *Stream) Stop() error {
	return s.srsStreamAction("sip_bye")
}
