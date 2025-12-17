package models

type GlobalSummary struct {
	TotalCritical int `json:"critical"`
	TotalHigh     int `json:"high"`
	TotalMedium   int `json:"medium"`
	TotalLow      int `json:"low"`
	TotalInfo     int `json:"info"`
}
