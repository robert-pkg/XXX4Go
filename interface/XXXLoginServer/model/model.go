package model

func GetDeviceType(deviceTypeName string) int {
	deviceType := 0
	switch deviceTypeName {
	case "web":
		deviceType = 1
	case "android":
		deviceType = 2
	case "ios":
		deviceType = 3
	}

	return deviceType
}
