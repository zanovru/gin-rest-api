package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/zanovru/gin-rest-api/pkg"
	"io"
)

const ctxKeyUuid string = "uuid"

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func logRequests(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)

	uuid := pkg.NewUuid()
	c.Set(ctxKeyUuid, uuid)

	log.Infof("uuid=%s, method=%v, url=%v ,body=%s", uuid, c.Request.Method, c.Request.URL, body)
	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()

	statusCode := c.Writer.Status()
	elapsed_ms := c.GetInt64("elapsed")
	if statusCode >= 400 {
		log.Errorf("uuid=%s, status=%d, url=%v ,response=%s, elapsed_ms=%d",
			uuid, statusCode, c.Request.URL, w.body.String(), elapsed_ms)
	} else {
		log.Infof("uuid=%s, status=%d, url=%v ,response=%s, elapsed_ms=%d",
			uuid, statusCode, c.Request.URL, w.body.String(), elapsed_ms)
	}
}
