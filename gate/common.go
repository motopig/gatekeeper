package gate

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func PosToQuery(c *gin.Context) (string, error) {
	var keys string
	c.Request.ParseForm()
	params := c.Request.PostForm
	for k, v := range params {
		keys = keys + k + "=" + v[0] + "&"
	}

	if len(keys) > 1 {
		return Substr(keys, 0, len(keys)-1), nil
	} else {
		return "", errors.New("empty post form")
	}

}
