package model

//Device 设备字段
type Device struct {
	ID                 string `json:"ID"`
	Name               string `json:"Name"`
	Manufacturer       string `json:"Manufacturer"`
	Type               string `json:"Type"`
	Password           string `json:"Password"`
	MediaTransport     string `json:"MediaTransport"`
	MediaTransportMode string `json:"MediaTransportMode"`
	CommandTransport   string `json:"CommandTransport"`
	RemoteIP           string `json:"RemoteIP"`
	RemotePort         int    `json:"RemotePort"`
	Online             bool   `json:"Online"`
	RecordCenter       bool   `json:"RecordCenter"`
	RecordIndistinct   bool   `json:"RecordIndistinct"`
	CivilCodeFirst     bool   `json:"CivilCodeFirst"`
	ChannelCount       int    `json:"ChannelCount"`
	RecvStreamIP       string `json:"RecvStreamIP"`
	SMSID              string `json:"SMSID"`
	ContactIP          string `json:"ContactIP"`
	Charset            string `json:"Charset"`
	CatalogInterval    int    `json:"CatalogInterval"`
	SubscribeInterval  int    `json:"SubscribeInterval"`
	Longitude          int    `json:"Longitude"`
	Latitude           int    `json:"Latitude"`
	LastRegisterAt     MyTime `json:"LastRegisterAt" gorm:"type:datetime"`
	LastKeepaliveAt    MyTime `json:"LastKeepaliveAt" gorm:"type:datetime"`
	CreatedAt          MyTime `json:"CreatedAt" gorm:"type:datetime"`
	UpdatedAt          MyTime `json:"UpdatedAt" gorm:"type:datetime"`
}

type DeviceChannel struct {
	ID           string `json:"ID"`           //通道编号
	DeviceID     string `json:"DeviceID"`     //设备编号
	DeviceName   string `json:"DeviceName"`   //设备名称
	DeviceOnline bool   `json:"DeviceOnline"` //设备在线状态
	Channel      int    `json:"Channel"`      //通道序号
	Name         string `json:"Name"`         //通道名称

	Address        string `json:"Address"`
	Altitude       int    `json:"Altitude"`
	AudioEnable    bool   `json:"AudioEnable"`
	BatteryLevel   int    `json:"BatteryLevel"`
	Block          string `json:"Block"`
	CivilCode      string `json:"CivilCode"`
	CloudRecord    bool   `json:"CloudRecord"`
	CreatedAt      string `json:"CreatedAt"`
	Custom         bool   `json:"Custom"`
	CustomName     string `json:"CustomName"`
	CustomPTZType  int    `json:"CustomPTZType"`
	CustomParentID string `json:"CustomParentID"`
	Direction      int    `json:"Direction,omitempty"`
	DownloadSpeed  string `json:"DownloadSpeed"`
	Latitude       int    `json:"Latitude"`
	Longitude      int    `json:"Longitude"`
	Manufacturer   string `json:"Manufacturer"`
	Model          string `json:"Model"`
	NumOutputs     int    `json:"NumOutputs"`
	Ondemand       bool   `json:"Ondemand"`
	Owner          string `json:"Owner"`
	PTZType        int    `json:"PTZType"`
	ParentID       string `json:"ParentID"`
	Parental       int    `json:"Parental"`
	Quality        string `json:"Quality"`
	RegisterWay    int    `json:"RegisterWay"`
	Secrecy        int    `json:"Secrecy"`
	Shared         bool   `json:"Shared"`
	SignalLevel    int    `json:"SignalLevel"`
	SnapURL        string `json:"SnapURL,omitempty"` //通道快照链接
	Speed          int    `json:"Speed"`
	Status         string `json:"Status"`
	StreamID       string `json:"StreamID,omitempty"` //直播流ID, 有值表示正在直播
	SubCount       int    `json:"SubCount"`           //子节点数, SubCount > 0 表示该通道为子目录
	UpdatedAt      string `json:"UpdatedAt"`
}
