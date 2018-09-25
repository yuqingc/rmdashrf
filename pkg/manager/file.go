package manager

// FileProperty is a data structure whichdescribes a file
// It is part of the result for list api
type FileProperty struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Mode    string `json:"mode"`
	ModTime string `json:"modTime"`
	IsDir   bool   `json:"isDir"`
}
