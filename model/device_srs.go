//+build srs
package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Cherisher/GB28181-Server-Srs/config"
)

//DeviceList 查询设备列表
func DeviceList() ([]*Device, error) {
	URL := config.Config().Srs.APIBase + "/api/v1/gb28181?action=sip_query_session"
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	type sessionResp struct {
		Code int `json:"code"`
		Data struct {
			Sessions []struct {
				ID           string `json:"id"`
				DeviceSumnum int    `json:"device_sumnum"`
				Devices      []struct {
					DeviceID     string `json:"device_id"`
					DeviceName   string `json:"device_name"`
					DeviceStatus string `json:"device_status"`
					InviteStatus string `json:"invite_status"`
					InviteTime   int    `json:"invite_time"`
				} `json:"devices"`
			} `json:"sessions"`
		} `json:"data"`
	}

	fmt.Println(string(body))
	var srsRsp sessionResp
	if err := json.Unmarshal(body, &srsRsp); err != nil {
		return nil, err
	}

	devices := make([]*Device, 0, len(srsRsp.Data.Sessions))
	for _, session := range srsRsp.Data.Sessions {
		devices = append(devices, &Device{
			ID:           session.ID,
			Name:         session.ID,
			ChannelCount: len(session.Devices),
			Online:       true,
		})
	}
	x, err := json.Marshal(devices)
	fmt.Println(string(x))
	return devices, nil
}

//DeviceInfo 查询设备
func (d *Device) DeviceInfo() error {
	URL := config.Config().Srs.APIBase + "/api/v1/gb28181?action=sip_query_session&id=" + d.ID
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	fmt.Println(URL)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type sessionResp struct {
		Code int `json:"code"`
		Data struct {
			Sessions []struct {
				ID           string `json:"id"`
				DeviceSumnum int    `json:"device_sumnum"`
				Devices      []struct {
					DeviceID     string `json:"device_id"`
					DeviceName   string `json:"device_name"`
					DeviceStatus string `json:"device_status"`
					InviteStatus string `json:"invite_status"`
					InviteTime   int    `json:"invite_time"`
				} `json:"devices"`
			} `json:"sessions"`
		} `json:"data"`
	}

	fmt.Println(string(body))
	var srsRsp sessionResp
	if err := json.Unmarshal(body, &srsRsp); err != nil {
		return err
	}

	devices := make([]*Device, 0, len(srsRsp.Data.Sessions))
	for _, session := range srsRsp.Data.Sessions {
		devices = append(devices, &Device{
			ID:           session.ID,
			Name:         session.ID,
			ChannelCount: len(session.Devices),
			Online:       true,
		})
	}
	x, err := json.Marshal(devices)
	fmt.Println(string(x))
	return nil
}

func (d *Device) ChannelList() ([]*DeviceChannel, error) {
	URL := config.Config().Srs.APIBase + "/api/v1/gb28181?action=sip_query_session&id=" + d.ID
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	fmt.Println(URL)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type AutoResp struct {
		Code int `json:"code"`
		Data struct {
			Sessions []struct {
				ID           string `json:"id"`
				DeviceSumnum int    `json:"device_sumnum"`
				Devices      []struct {
					DeviceID     string `json:"device_id"`
					DeviceName   string `json:"device_name"`
					DeviceStatus string `json:"device_status"`
					InviteStatus string `json:"invite_status"`
					InviteTime   int    `json:"invite_time"`
				} `json:"devices"`
			} `json:"sessions"`
		} `json:"data"`
	}
	var srsRsp AutoResp
	if err := json.Unmarshal(body, &srsRsp); err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	if len(srsRsp.Data.Sessions) != 1 {
		return nil, errors.New("bad response")
	}
	fmt.Println(len(srsRsp.Data.Sessions[0].Devices))
	channels := make([]*DeviceChannel, 0, len(srsRsp.Data.Sessions[0].Devices))
	for _, channel := range srsRsp.Data.Sessions[0].Devices {
		channels = append(channels, &DeviceChannel{
			ID:           channel.DeviceID,
			DeviceID:     d.ID,
			Name:         channel.DeviceName,
			DeviceOnline: true,
			Status:       channel.DeviceStatus,
		})
	}

	return channels, nil

}
