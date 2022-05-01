package lessons

type Lesson struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	VideoFileUrl    string `json:"video_file_url"`
	DocumentFileUrl string `json:"document_file_url"`
	GroupSubjectId  string `json:"group_subject_id"`
	CreatedDate     string `json:"created_date"`
}

type LessonUpdate struct {
	Id              string  `json:"id"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	VideoFileUrl    *string `json:"video_file_url"`
	DocumentFileUrl *string `json:"document_file_url"`
	GroupSubjectId  *string `json:"group_subject_id"`
}

type File struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	FileUrl     string   `json:"file_url"`
	FileType    FileType `json:"file_type"`
}

type FileUpdate struct {
	Id          string    `json:"id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	FileUrl     *string   `json:"file_url"`
	FileType    *FileType `json:"file_type"`
}

type FileType string

var (
	Video FileType = "video"
	Image FileType = "image"
)

var (
	fileTypeToString = map[FileType]string{
		Video: "video",
		Image: "image",
	}
	stringToFileType = map[string]FileType{
		"video": Video,
		"image": Image,
	}
)

func (c FileType) ToString() string {
	return fileTypeToString[c]
}

func ToFileType(s string) FileType {
	return stringToFileType[s]
}

func IsFileTypeExist(s string) bool {
	fileTypes := []string{"video", "image"}
	for _, v := range fileTypes {
		if v == s {
			return true
		}
	}
	return false
}
