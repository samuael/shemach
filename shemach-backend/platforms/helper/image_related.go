package helper

import "strings"

var ImageExtensions = []string{"jpeg", "png", "jpg", "gif", "btmp"}

// JPEGFileName function
func JPEGFileName(filename string) string {
	filenameSlice := strings.Split(filename, ".")
	if len(filenameSlice) > 1 {
		filenames := strings.Join(filenameSlice[:len(filenameSlice)-1], "")
		filenames += ".jpg"
		return filenames
	}
	return filename + ".jpg"
}
