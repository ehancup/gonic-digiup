package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	"base-gin/exception"
	"errors"
	"strconv"

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
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id",h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
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

func (h *PublisherHandler) getList(c *gin.Context) {
	// var req dto.Filter
	// if err := c.ShouldBindQuery(&req); err != nil {
	// 	c.JSON(h.hr.BindingError(err))
	// 	return
	// }

	data, err := h.service.GetList()
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(exception.ErrDataNotFound.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.PublisherDetailResp]{
		Success: true,
		Message: "Daftar anggota",
		Data:    data,
	})
}

func (h *PublisherHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(exception.ErrDataNotFound.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.PublisherDetailResp]{
		Success: true,
		Message: "Detail anggota",
		Data:    data,
	})
}

// update godoc
//
//	@Summary Update a publisher's detail
//	@Description Update a publisher's detail.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Publisher's ID"
//	@Param detail body dto.PublisherUpdateReq true "Publishers's detail"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers/{id} [put]
func (h *PublisherHandler) update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	var req dto.PublisherUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}
	req.ID = uint(id)

	err = h.service.Update(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDateParsing):
			c.JSON(http.StatusBadRequest, h.hr.ErrorResponse(err.Error()))
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil disimpan",
	})
}

func (h *PublisherHandler) delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDateParsing):
			c.JSON(http.StatusBadRequest, h.hr.ErrorResponse(err.Error()))
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil dihapus",
	})
}