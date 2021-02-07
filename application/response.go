package application

import "github.com/allentom/haruka"

func AbortError(ctx *haruka.Context, err error, status int) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(status)
	ctx.JSON(map[string]interface{}{
		"success": false,
		"err":     err,
	})
}
