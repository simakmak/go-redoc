package ginredoc

import (
	"github.com/gin-gonic/gin"
	"github.com/simakmak/go-redoc"
)

func New(doc redoc.Redoc) gin.HandlerFunc {
	handle := doc.Handler()
	return func(ctx *gin.Context) {
		handle(ctx.Writer, ctx.Request)
	}
}
