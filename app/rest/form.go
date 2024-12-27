package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"


	// "base-gin/app/sevice"
	"base-gin/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormHandler struct {

	hr *server.Handler
	service *service.FormService
}

func newFormHandler(hr *server.Handler, formService *service.FormService) *FormHandler {
	return &FormHandler{hr: hr, service: formService}
}

func (h *FormHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootForm)
	grp.POST("", h.create)

}

func (h *FormHandler) create(c *gin.Context) {
	var req dto.FormPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return 
	}

	if (len(req.GolonganDarah) > 2 ) {
		c.JSON(http.StatusCreated, dto.SuccessResponse[*any]{
			Success: false,
			Message: "golongan darah invalid",
		})

		return 
	}

	if (req.Gender != "f" && req.Gender != "m") {
		c.JSON(http.StatusCreated, dto.SuccessResponse[*any]{
			Success: false,
			Message: "gender harus diisi dengan 'f' atau 'm'",
		})

		return 
	}

	err := h.service.Create(&req)

	if err != nil {

		h.hr.ErrorInternalServer(c, err)
		return 
	}

	
	c.JSON(http.StatusCreated, dto.SuccessResponse[*any]{
		Success: true,
		Message: "Data form berhasil disimpan",
	})
}


// create godoc
//
//	@Summary Create New Publisher
//	@Description Create new publlisher
//	@Accept json
//	@Produce json
// 	@Security BearerAuth
//	@Param newItem body dto.PublisherCreateReq true "Payload"
//	@Success 201 {object} dto.SuccessResponse[dto.PublisherCreateResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers [post]
