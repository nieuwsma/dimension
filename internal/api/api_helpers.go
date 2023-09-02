package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nieuwsma/dimension/pkg/presentation"
	"net/http"
)

// WriteJsonWithHeaders - writes JSON to the open http connection along with headers
func WriteJsonWithHeaders(c *gin.Context, pb presentation.APIPayload) {
	if pb.IsError {
		c.Header("Content-Type", "application/problem+json")
		c.AbortWithStatusJSON(pb.StatusCode, pb.Error)
	} else if pb.Obj != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(pb.StatusCode, pb.Obj)
	} else {
		c.Status(pb.StatusCode)
	}
}

func BuildErrorPayload(err error) (pb presentation.APIPayload) {
	var se *presentation.StorageFailed
	var ie *presentation.InternalServerError
	var nf *presentation.NotFound
	var ua *presentation.Forbidden

	if errors.As(err, &se) {
		pb = presentation.BuildErrorPassback(http.StatusInternalServerError, err)
		return
	} else if errors.As(err, &nf) {
		pb = presentation.BuildErrorPassback(http.StatusNotFound, err)
		return
	} else if errors.As(err, &ua) {
		pb = presentation.BuildErrorPassback(http.StatusForbidden, err)
		return
	} else if errors.As(err, &ie) {
		pb = presentation.BuildErrorPassback(http.StatusInternalServerError, err)
		return
	}

	pb = presentation.BuildErrorPassback(http.StatusBadRequest, err)
	return

}
