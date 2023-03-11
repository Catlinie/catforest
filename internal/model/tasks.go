package model

type TaskInfo struct {
	Id          uint32 `json:"id" gorm:"primary_key"`
	Cid         string `json:"cid"`
	Start       uint32 `json:"start"`
	End         uint32 `json:"end"`
	Duration    uint32 `json:"duration"`
	YearsIndex  uint8  `json:"years_index"`
	MonthsIndex uint8  `json:"months_index"`
	DaysIndex   uint8  `json:"days_index"`
	HoursIndex  uint8  `json:"hours_index"`
}

func (model TaskInfo) TableName() string {
	return "tasks"
}
