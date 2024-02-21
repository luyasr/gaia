package response

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/errors"
	"net/http"
)

type Response struct {
	Code     int               `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata"`
	Data     any               `json:"data"`
}

func GinJson(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:     200,
		Reason:   "",
		Message:  "success",
		Metadata: nil,
		Data:     data,
	})
}

func GinJsonWithError(c *gin.Context, err error) {
	defer c.Abort()

	e := errors.FromError(err)
	httpCode := int(e.Code)

	if http.StatusText(int(e.Code)) == "" {
		httpCode = http.StatusInternalServerError
	}

	c.JSON(httpCode, Response{
		Code:     int(e.Code),
		Reason:   e.Reason,
		Message:  e.Message,
		Metadata: e.Metadata,
		Data:     nil,
	})
}
