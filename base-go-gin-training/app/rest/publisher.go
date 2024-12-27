package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	// "base-gin/app/sevice"
	"base-gin/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublisherHandler struct {

	hr *server.Handler
	service *service.PublisherService
}

func newPublisherHandler(hr *server.Handler, publisherservice *service.PublisherService) *PublisherHandler {
	return &PublisherHandler{hr: hr, service: publisherservice}
}

func (h *PublisherHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootPublisher)
	grp.POST("", h.hr.AuthAccess(), h.create)
}

// create godoc
//
//	@Summary Create New Publisher
//	@Description Create new publlisher
//	@Accept json
//	@Produce json
//	@Param newItem body dto.PublisherCreateReq true "Credential"
//	@Success 201 {object} dto.SuccessResponse[dto.PublisherCreateResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publisher [post]

func (h *PublisherHandler) create(c *gin.Context) {
	var req dto.PublisherCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return 
	}

	data, err := h.service.Create(&req)

	if err != nil {

		h.hr.ErrorInternalServer(c, err)
		return 
	}
	c.JSON(http.StatusCreated, dto.SuccessResponse[*dto.PublisherCreateResp]{
		Success: true,
		Message: "Data penerbit berhasil disimpan",
		Data: data,
	})
}