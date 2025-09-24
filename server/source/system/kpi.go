package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderKpi = initOrderRating + 1

type initKpi struct{}

func init() {
	system.RegisterInit(initOrderKpi, &initKpi{})
}

func (i *initKpi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	if err := db.AutoMigrate(&sysModel.SysKpiCategory{}, &sysModel.SysKpiItem{}); err != nil {
		return ctx, err
	}
	return ctx, nil
}

func (i *initKpi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysKpiCategory{}) && db.Migrator().HasTable(&sysModel.SysKpiItem{})
}

func (i *initKpi) InitializerName() string {
	return "kpi"
}

func (i *initKpi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	// Seed default KPI: Time Management with three items
	weight := 100
	cat := sysModel.SysKpiCategory{Name: "Time Management", Description: "Time management KPI", Weight: &weight}
	if err := db.Create(&cat).Error; err != nil {
		return ctx, errors.Wrap(err, "seed category failed")
	}
	w1, w2, w3 := 30, 30, 40
	items := []sysModel.SysKpiItem{
		{CategoryID: cat.ID, Name: "Planning", Weight: &w1},
		{CategoryID: cat.ID, Name: "Prioritizing Tasks", Weight: &w2},
		{CategoryID: cat.ID, Name: "Setting & Meeting Goals", Weight: &w3},
	}
	if err := db.Create(&items).Error; err != nil {
		return ctx, errors.Wrap(err, "seed items failed")
	}
	return ctx, nil
}

func (i *initKpi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return false
	}
	var cnt int64
	if err := db.Model(&sysModel.SysKpiCategory{}).Count(&cnt).Error; err != nil {
		return false
	}
	return cnt > 0
}
