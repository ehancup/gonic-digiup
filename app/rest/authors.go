package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	"base-gin/exception"
	"base-gin/server"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorsHandler struct {

	hr *server.Handler
	service *service.AuthorsService
}

func newAuthorsHandler(hr *server.Handler, authorsService *service.AuthorsService) *AuthorsHandler {
	return &AuthorsHandler{hr: hr, service: authorsService}
}

func (h *AuthorsHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootAuthors)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("/:id", h.getByID)
	grp.GET("/", h.getList)
	grp.PUT("/:id",h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id",h.hr.AuthAccess(), h.delete)
}

func (h *AuthorsHandler) create(c *gin.Context) {
	var req dto.AuthorsCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return 
	}

	data, err := h.service.Create(&req)

	if err != nil {

		h.hr.ErrorInternalServer(c, err)
		return 
	}
	c.JSON(http.StatusCreated, dto.SuccessResponse[*dto.AuthorsCreateResp]{
		Success: true,
		Message: "Data author berhasil disimpan",
		Data: data,
	})
}

func (h *AuthorsHandler) getByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AuthorsDetailResp]{
		Success: true,
		Message: "Detail author",
		Data:    data,
	})
}

func (h *AuthorsHandler) getList(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.AuthorsDetailResp]{
		Success: true,
		Message: "Daftar authors",
		Data:    data,
	})
}

func (h *AuthorsHandler) update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	var req dto.AuthorsUpdateReq
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

func (h *AuthorsHandler) delete(c *gin.Context) {
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
		Message: "Author berhasil dihapus",
	})
}