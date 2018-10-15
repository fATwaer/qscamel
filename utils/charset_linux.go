// +build linux

package utils

import (
	"github.com/gogs/chardet"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	charsetDetector *chardet.Detector

	gb18030Decoder *encoding.Decoder
)

// parseGB18030 will parse GB18030 encoded string to UTF-8 encoded.
func parseGB18030(name string) (string, error) {
	s, err := gb18030Decoder.String(name)
	if err != nil {
		return "", err
	}
	logrus.Infof("Object name is GB18030 encoded, converted to %s.", s)
	return s, nil
}

// ConvertToUTF8 will convert the file name to UTF-8.
func ConvertToUTF8(name string) (string, error) {
	r, err := charsetDetector.DetectBest([]byte(name))
	if err != nil {
		return "", err
	}

	// Currently, we only support convert from gb18030 encoding.
	switch r.Charset {
	case "GB18030":
		return parseGB18030(name)
	default:
		return name, nil
	}
}

func init() {
	charsetDetector = chardet.NewTextDetector()

	gb18030Decoder = simplifiedchinese.GB18030.NewDecoder()
}
