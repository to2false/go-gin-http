package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/to2false/go-gin-http/message"
	"github.com/to2false/go-gin-http/validate"
)

type (
	RequestValidator interface {
		Validate() error
	}
)

func GinHandlerWrap[REQ any, RESP any](
	transformerName string,
	handlerFn func(context.Context, *REQ) (*RESP, error),
) gin.HandlerFunc {
	return func(c *gin.Context) {
		transformer := message.GetTransformer(transformerName)

		if err := transformer.PreProcessRequest(c.Request); err != nil {
			statusCode, data, e := transformer.Err(c.Request.Context(), err)
			if e != nil {
				c.JSON(statusCode, e.Error())
				c.Abort()

				return
			}

			c.Data(statusCode, transformer.ContentType(), data)
			c.Abort()

			return
		}

		var in REQ
		if err := c.ShouldBind(&in); err != nil {
			statusCode, data, e := transformer.Err(c.Request.Context(), err)
			if e != nil {
				c.JSON(statusCode, e.Error())
				c.Abort()

				return
			}

			c.Data(statusCode, transformer.ContentType(), data)
			c.Abort()

			return
		}

		if ginValidate(c, transformer, in) {
			return
		}

		out, err := handlerFn(c.Request.Context(), &in)
		if err != nil {
			statusCode, data, e := transformer.Err(c.Request.Context(), err)
			if e != nil {
				c.JSON(statusCode, e.Error())
				c.Abort()

				return
			}

			c.Data(statusCode, transformer.ContentType(), data)
			c.Abort()

			return
		}

		httpStatusCode, data, err := transformer.Transform(c.Request.Context(), &out)
		if err != nil {
			statusCode, data, e := transformer.Err(c.Request.Context(), err)
			if e != nil {
				c.JSON(statusCode, e.Error())
				c.Abort()

				return
			}

			c.Data(statusCode, transformer.ContentType(), data)
			c.Abort()

			return
		}

		c.Data(httpStatusCode, transformer.ContentType(), data)
		c.Abort()
	}
}

func ginValidate(c *gin.Context, transformer message.Transformer, req any) (interrupted bool) {
	if v, ok := req.(RequestValidator); ok {
		if err := v.Validate(); err != nil {
			statusCode, data, e := transformer.Err(c.Request.Context(), validate.NewValidateError(err.Error()))
			if e != nil {
				c.JSON(statusCode, e.Error())
				c.Abort()

				return true
			}

			c.Data(statusCode, transformer.ContentType(), data)
			c.Abort()

			return true
		}
	}

	return false
}
