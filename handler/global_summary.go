package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toezayarmoe/vulntrack_api/config"
	"github.com/toezayarmoe/vulntrack_api/models"
)

func GlobalSummary(ctx *gin.Context) {

	isAdmin, exist := ctx.Get("is_admin")
	if !exist || isAdmin != true {
		ctx.JSON(403, gin.H{
			"error": "admin access is required",
		})
		return
	}
	var result models.GlobalSummary

	query := `
		SELECT
			SUM(CASE WHEN v.risk_factor = 'Critical' THEN 1 ELSE 0 END) AS total_critical,
			SUM(CASE WHEN v.risk_factor = 'High' THEN 1 ELSE 0 END) AS total_high,
			SUM(CASE WHEN v.risk_factor = 'Medium' THEN 1 ELSE 0 END) AS total_medium,
			SUM(CASE WHEN v.risk_factor = 'Low' THEN 1 ELSE 0 END) AS total_low,
			SUM(CASE WHEN v.risk_factor = 'Info' THEN 1 ELSE 0 END) AS total_info
		FROM asset_severity_summary m
		JOIN vulnerabilities v
		  ON m.current_report_id = v.report_id
		 AND (m.ip = v.ip_address OR m.ip = v.hostname)
	`

	if err := config.DB.Raw(query).Scan(&result).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch global summary",
		})
		return
	}
	fmt.Println(result)
	ctx.JSON(http.StatusOK, result)
}
