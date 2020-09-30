package model

import "errors"

func GetDeviceType(deviceTypeName string) (int, error) {
	deviceType := 0
	switch deviceTypeName {
	case "web":
		deviceType = 1
	case "android":
		deviceType = 2
	case "ios":
		deviceType = 3
	default:
		return 0, errors.New("no support device type")
	}

	return deviceType, nil
}
