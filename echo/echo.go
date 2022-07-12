package echoredoc

import (
	"github.com/labstack/echo/v4"
	"github.com/simakmak/go-redoc"
)

func New(doc redoc.Redoc) echo.MiddlewareFunc {
	handle := doc.Handler()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			handle(ctx.Response(), ctx.Request())
			return nil
		}
	}
}
