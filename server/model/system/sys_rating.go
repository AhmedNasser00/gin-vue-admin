package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysRating 结构体
type SysRating struct {
	global.GVA_MODEL
	Rating           *int   `json:"rating" form:"rating" gorm:"column:rating;uniqueIndex;comment:评级值;"`
	Description      string `json:"description" form:"description" gorm:"column:description;comment:描述;"`
	PerformanceLevel string `json:"performanceLevel" form:"performanceLevel" gorm:"column:performance_level;comment:绩效水平;"`
	MinPercentage    *int   `json:"minPercentage" form:"minPercentage" gorm:"column:min_percentage;comment:最小百分比;"`
	MaxPercentage    *int   `json:"maxPercentage" form:"maxPercentage" gorm:"column:max_percentage;comment:最大百分比;"`
}

func (SysRating) TableName() string {
	return "sys_ratings"
}
