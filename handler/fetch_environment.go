package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toezayarmoe/vulntrack_api/config"
	"github.com/toezayarmoe/vulntrack_api/models"
)

func FetchEnvironment(ctx *gin.Context) {

	isAdmin, exist := ctx.Get("is_admin")
	if !exist || isAdmin != true {
		ctx.JSON(403, gin.H{
			"error": "admin access is required",
		})
		return
	}
	var results []models.Environment

	query := `
	SELECT
	    m.env_id,
	    e.name,
	    SUM(CASE WHEN v.risk_factor = 'Critical' THEN 1 ELSE 0 END) AS critical,
	    SUM(CASE WHEN v.risk_factor = 'High' THEN 1 ELSE 0 END) AS high,
	    SUM(CASE WHEN v.risk_factor = 'Medium' THEN 1 ELSE 0 END) AS medium,
	    SUM(CASE WHEN v.risk_factor = 'Low' THEN 1 ELSE 0 END) AS low,
		SUM(CASE WHEN v.risk_factor = 'Info' THEN 1 ELSE 0 END) AS info
	FROM asset_severity_summary m
	JOIN vulnerabilities v ON m.current_report_id = v.report_id
	    AND (m.ip = v.ip_address OR m.ip = v.hostname)
	JOIN environments e ON e.id = m.env_id
	GROUP BY m.env_id, e.name
	`

	if err := config.DB.Raw(query).Scan(&results).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch global summary",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"total": len(results), "data": results})
}
