package delivery

import (
	"net/http"
	"path/filepath"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"snakealive/m/internal/gateway/media/usecase"
	"snakealive/m/internal/models"
	"snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
)

type MediaDelivery interface {
	UploadFile(ctx *fasthttp.RequestCtx)
}

type mediaDelivery struct {
	manager      usecase.MediaUsecase
	errorAdapter error_adapter.HttpAdapter
}

func (m *mediaDelivery) UploadFile(ctx *fasthttp.RequestCtx) {
	form, err := ctx.FormFile(constants.FileKey)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	file, err := form.Open()
	if err != nil {
		ctx.SetStatusCode(http.StatusConflict)
		return
	}
	defer file.Close()

	filename, err := m.manager.UploadFile(ctx, file, filepath.Ext(form.Filename))
	if err != nil {
		ctx.SetStatusCode(http.StatusTeapot)
		return
	}

	b, _ := easyjson.Marshal(models.UploadFileResponse{Filename: filename})
	ctx.Response.SetBody(b)
	ctx.SetStatusCode(http.StatusOK)
}

func NewMediaDelivery(
	manager usecase.MediaUsecase,
	errorAdapter error_adapter.HttpAdapter,
) MediaDelivery {
	return &mediaDelivery{
		manager:      manager,
		errorAdapter: errorAdapter,
	}
}
