package models

import "time"

// Report represents the full report table (add other fields if they exist in your DB)
type Report struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"` // Optional: Good for sorting
}

// ReportListItem is a simplified struct for your specific API response
type ReportListItem struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}
