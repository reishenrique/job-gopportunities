package recruiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List recruiter
// @Description List all recruiters
// @Tags Recruiters
// @Accept json
// @Produce json
// @Success 200 {object} ListRecruitersResponse
// @Failure 500 {object} ErrorResponse
// @Router /recruiters [get]
func ListRecruitersHandler(ctx *gin.Context) {
	recruiters := []schemas.Recruiters{}

	if err := db.Find(&recruiters).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing recruiters")
		return
	}

	sendSuccess(ctx, "list-recruiters", recruiters)
}	