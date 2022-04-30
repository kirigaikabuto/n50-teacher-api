package common

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
	Admin   UserType = "админ"
	Teacher UserType = "учитель"
	Student UserType = "ученик"
)

var (
	userTypeToString = map[UserType]string{
		Admin:   "админ",
		Teacher: "учитель",
		Student: "ученик",
	}
	stringToUserType = map[string]UserType{
		"админ":   Admin,
		"учитель": Teacher,
		"ученик":  Student,
	}
)

func (u UserType) ToString() string {
	return userTypeToString[u]
}

func ToUserType(s string) UserType {
	return stringToUserType[s]
}

func IsUserTypeExist(s string) bool {
	userTypes := []string{"админ", "учитель", "ученик"}
	for _, v := range userTypes {
		if v == s {
			return true
		}
	}
	return false
}