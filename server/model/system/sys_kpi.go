package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysKpiCategory represents a parent KPI (e.g., Time Management)
type SysKpiCategory struct {
	global.GVA_MODEL
	Name        string       `json:"name" form:"name" gorm:"column:name;uniqueIndex;comment:KPI类别名称;"`
	Description string       `json:"description" form:"description" gorm:"column:description;comment:描述;"`
	Weight      *int         `json:"weight" form:"weight" gorm:"column:weight;comment:权重百分比;"`
	Items       []SysKpiItem `json:"items" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (SysKpiCategory) TableName() string {
	return "sys_kpi_categories"
}

// SysKpiItem represents a KPI key under a category (e.g., Planning)
type SysKpiItem struct {
	global.GVA_MODEL
	CategoryID uint   `json:"categoryID" form:"categoryID" gorm:"column:category_id;index;comment:所属类别ID;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:KPI键名称;"`
	Weight     *int   `json:"weight" form:"weight" gorm:"column:weight;comment:权重百分比;"`
}

func (SysKpiItem) TableName() string {
	return "sys_kpi_items"
}
