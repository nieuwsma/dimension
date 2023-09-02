package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nieuwsma/dimension/internal/logger"
	"net/http"
)

// WriteJSON - writes JSON to the open http connection
func WriteJSON(c *gin.Context, i interface{}) {
	obj, err := json.Marshal(i)
	if err != nil {
		logger.Log.Error(err)
	}
	c.Writer.Write(obj)
}

// WriteHeaders - writes JSON to the open http connection along with headers
func WriteHeaders(c *gin.Context, pb APIPayload) {
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

func WriteHeadersWithLocation(c *gin.Context, pb APIPayload, location string) {
	c.Header("Location", location)
	WriteHeaders(c, pb)
}

func BuildErrorPayload(err error) (pb APIPayload) {
	var se *StorageFailed
	var ie *InternalServerError
	var nf *NotFound
	var ua *Forbidden

	if errors.As(err, &se) {
		pb = BuildErrorPassback(http.StatusInternalServerError, err)
		return
	} else if errors.As(err, &nf) {
		pb = BuildErrorPassback(http.StatusNotFound, err)
		return
	} else if errors.As(err, &ua) {
		pb = BuildErrorPassback(http.StatusForbidden, err)
		return
	} else if errors.As(err, &ie) {
		pb = BuildErrorPassback(http.StatusInternalServerError, err)
		return
	}

	pb = BuildErrorPassback(http.StatusBadRequest, err)
	return

}
