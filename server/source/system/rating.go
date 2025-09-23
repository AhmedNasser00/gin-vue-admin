package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderRating = initOrderExcelTemplate + 1

type initRating struct{}

// auto run
func init() {
	system.RegisterInit(initOrderRating, &initRating{})
}

func (i *initRating) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysRating{})
}

func (i *initRating) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysRating{})
}

func (i *initRating) InitializerName() string {
	return sysModel.SysRating{}.TableName()
}

func (i *initRating) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// Seed default rating records (1..5)
	r1, r2, r3, r4, r5 := 1, 2, 3, 4, 5
	min1, max1 := 0, 59
	min2, max2 := 60, 69
	min3, max3 := 70, 79
	min4, max4 := 80, 89
	min5, max5 := 90, 100

	entities := []sysModel.SysRating{
		{Rating: &r1, Description: "Unsatisfactory", PerformanceLevel: "Below Minimum Standards", MinPercentage: &min1, MaxPercentage: &max1},
		{Rating: &r2, Description: "Needs Improvement", PerformanceLevel: "Near Minimum Standards", MinPercentage: &min2, MaxPercentage: &max2},
		{Rating: &r3, Description: "Satisfactory", PerformanceLevel: "Meets Minimum Standards", MinPercentage: &min3, MaxPercentage: &max3},
		{Rating: &r4, Description: "Very Good", PerformanceLevel: "Above Minimum Standards", MinPercentage: &min4, MaxPercentage: &max4},
		{Rating: &r5, Description: "Excellent", PerformanceLevel: "Exceeds Standards", MinPercentage: &min5, MaxPercentage: &max5},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysRating{}.TableName()+" seed failed")
	}
	return ctx, nil
}

func (i *initRating) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value(system.ContextKeyDB).(*gorm.DB)
	if !ok {
		return false
	}
	var cnt int64
	if err := db.Model(&sysModel.SysRating{}).Count(&cnt).Error; err != nil {
		return false
	}
	return cnt > 0
}
