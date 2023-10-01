package browser

import (
	ua "github.com/mileusna/useragent"
)

type DeviceInfo struct {
	VersionNo   string
	OSVersionNo string
	URL         string
	UserAgent   string
	Name        string
	Version     string
	OS          string
	OSVersion   string
	Device      string
	Mobile      bool
	Tablet      bool
	Desktop     bool
	Bot         bool
}

func GetInfo(userAgent string) DeviceInfo {
	deviceInfo := ua.Parse(userAgent)
	return DeviceInfo{
		VersionNo:   deviceInfo.VersionNoFull(),
		OSVersionNo: deviceInfo.OSVersionNoFull(),
		URL:         deviceInfo.URL,
		UserAgent:   deviceInfo.String,
		Name:        deviceInfo.Name,
		Version:     deviceInfo.Version,
		OS:          deviceInfo.OS,
		OSVersion:   deviceInfo.OSVersion,
		Device:      deviceInfo.Device,
		Mobile:      deviceInfo.Mobile,
		Tablet:      deviceInfo.Tablet,
		Desktop:     deviceInfo.Desktop,
		Bot:         deviceInfo.Bot,
	}
}
