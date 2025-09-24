package utils

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	enforcerInitMu       sync.Mutex
)

// GetCasbin 获取casbin实例
func GetCasbin() *casbin.SyncedCachedEnforcer {
	if syncedCachedEnforcer != nil {
		return syncedCachedEnforcer
	}
	enforcerInitMu.Lock()
	defer enforcerInitMu.Unlock()
	if syncedCachedEnforcer != nil {
		return syncedCachedEnforcer
	}
	if global.GVA_DB == nil {
		// DB not ready yet
		return nil
	}
	a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
	if err != nil {
		zap.L().Error("Casbin 适配数据库失败", zap.Error(err))
		return nil
	}
	text := `
    [request_definition]
    r = sub, obj, act
    
    [policy_definition]
    p = sub, obj, act
    
    [role_definition]
    g = _, _
    
    [policy_effect]
    e = some(where (p.eft == allow))
    
    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
    `
	m, err := model.NewModelFromString(text)
	if err != nil {
		zap.L().Error("Casbin 字符串加载模型失败", zap.Error(err))
		return nil
	}
	enforcer, err := casbin.NewSyncedCachedEnforcer(m, a)
	if err != nil {
		zap.L().Error("Casbin Enforcer 创建失败", zap.Error(err))
		return nil
	}
	enforcer.SetExpireTime(60 * 60)
	if err := enforcer.LoadPolicy(); err != nil {
		zap.L().Error("Casbin 加载策略失败", zap.Error(err))
		return nil
	}
	syncedCachedEnforcer = enforcer
	return syncedCachedEnforcer
}
