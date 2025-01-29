package basic_controller

import (
	basic_service "Online-Test-GO/internal/service/_basic_service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BasicDelete(c *gin.Context) (context.Context, basic_service.Delete, error) {
	ctx := context.Background()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "id must be number!",
			"status":  false,
		})

		return ctx, basic_service.Delete{}, err
	}

	var data basic_service.Delete

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return ctx, basic_service.Delete{}, err
	}

	data.Id = &id

	return ctx, data, nil
}
