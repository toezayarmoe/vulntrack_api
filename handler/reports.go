package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toezayarmoe/vulntrack_api/config"
	"github.com/toezayarmoe/vulntrack_api/models"
)

func GetReports(ctx *gin.Context) {
	var reports []models.ReportListItem

	// 1. Retrieve user info set by the JWT middleware
	// Note: JWT numeric claims are often float64 by default, so we assert carefully
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	isAdminVal, _ := ctx.Get("is_admin")
	isAdmin, ok := isAdminVal.(bool)
	if !ok {
		isAdmin = false
	}

	// 2. Prepare the database query
	tx := config.DB.Table("reports").Select("reports.id, reports.name")

	// 3. Apply filters based on role
	if !isAdmin {
		// If NOT admin, join with access table to filter results
		// Query: SELECT reports.id, reports.name FROM reports
		//        JOIN user_environment_access ON reports.environment_id = user_environment_access.environment_id
		//        WHERE user_environment_access.user_id = ?
		tx = tx.Joins("JOIN user_environment_access ON reports.environment_id = user_environment_access.environment_id").
			Where("user_environment_access.user_id = ?", userID)
	}

	// 4. Execute the query
	if err := tx.Scan(&reports).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch reports",
		})
		return
	}

	ctx.JSON(http.StatusOK, reports)
}
