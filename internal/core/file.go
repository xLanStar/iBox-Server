package core

// type FileType uint8

// const (
// 	Unknown FileType = iota
// 	MusicFile
// 	VideoFile
// 	PictureFile
// 	TextFile
// )

type File struct {
	// Id       uint32   `json:"id"`
	FileName string `json:"name"`
	// FilePath string   `json:"-"`
	// Type FileType `json:"-"`
}

// func GetType(fileName string) FileType {
// 	switch fileName {
// 	case ".mp3", ".flac", ".ogg", ".wav":
// 		return MusicFile
// 	case ".mp4", ".mpeg", ".mov", ".mkv", ".avi", ".flv":
// 		return VideoFile
// 	case ".jpg", ".png", ".bmp":
// 		return PictureFile
// 	case ".txt":
// 		return TextFile
// 	default:
// 		return Unknown
// 	}
// }
