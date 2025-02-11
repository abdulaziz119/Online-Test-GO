package basic_controller

import (
	basic_repo "Online-Test-GO/internal/repository/_basic_repo"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func BasicDelete(c *gin.Context) (context.Context, basic_repo.Delete, error) {
	ctx := context.Background()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "id must be number!",
			"status":  false,
		})

		return ctx, basic_repo.Delete{}, err
	}

	var data basic_repo.Delete

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return ctx, basic_repo.Delete{}, err
	}

	data.Id = &id

	return ctx, data, nil
}
