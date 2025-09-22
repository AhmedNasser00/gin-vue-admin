package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
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
	return ctx, nil
}

func (i *initRating) DataInserted(ctx context.Context) bool {
	return true
}
