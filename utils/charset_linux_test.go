package utils

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func TestConvertToUTF8(t *testing.T) {
	r := "qscamel 是一个用于在不同的端点 (Endpoint) 中高效迁移数据的工具。"

	s, err := simplifiedchinese.GB18030.NewEncoder().String(r)
	assert.NoError(t, err)

	// Check result.Charset.
	x, err := charsetDetector.DetectBest([]byte(s))
	t.Log(x.Charset)
	assert.Equal(t, "GB18030", x.Charset)

	// Check convert result.
	ans, err := ConvertToUTF8(s)
	t.Log(ans)

	assert.NoError(t, err)
	assert.Equal(t, r, ans)
}
