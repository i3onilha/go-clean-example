package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"go-clean-example/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (h *UserController) GetUserWithOrders(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}
	if offset < 0 {
		offset = 0
	}

	resp, err := h.usecase.GetUserWithOrders(c.Request.Context(), int32(userID), limit, offset)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user with orders"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
