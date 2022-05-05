package common

import "net"

var (
	imageTypes = []string{
		"image/gif",
		"image/jpeg",
		"image/png",
		"image/tiff",
		"image/vnd.microsoft.icon",
		"image/x-icon",
		"image/vnd.djvu",
		"image/svg+xml",
	}
	videoTypes = []string{
		"video/mpeg",
		"video/mp4",
		"video/webm",
	}
)

func IsImage(contentType string) bool {
	for _, v := range imageTypes {
		if v == contentType {
			return true
		}
	}
	return false
}

func IsVideo(contentType string) bool {
	for _, v := range videoTypes {
		if v == contentType {
			return true
		}
	}
	return false
}

func IsAvailableResource(currentType string, neededTypes []string) bool {
	for _, v := range neededTypes {
		if v == currentType {
			return true
		}
	}
	return false
}

type UserType string

var (
	Admin   UserType = "admin"
	Teacher UserType = "teacher"
	Student UserType = "student"
)

var (
	userTypeToString = map[UserType]string{
		Admin:   "admin",
		Teacher: "teacher",
		Student: "student",
	}
	stringToUserType = map[string]UserType{
		"admin":   Admin,
		"teacher": Teacher,
		"student": Student,
	}
)

func (u UserType) ToString() string {
	return userTypeToString[u]
}

func ToUserType(s string) UserType {
	return stringToUserType[s]
}

func IsUserTypeExist(s string) bool {
	userTypes := []string{"admin", "teacher", "student"}
	for _, v := range userTypes {
		if v == s {
			return true
		}
	}
	return false
}

type Resp struct {
	Objects interface{} `json:"objects"`
}

func GetOutboundIP() net.IP {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
