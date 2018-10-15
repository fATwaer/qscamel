// +build !linux

package utils

// ConvertToUTF8 will convert the file name to UTF-8.
func ConvertToUTF8(name string) (string, error) {
	return name, nil
}
