package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KpiRouter struct{}

func (s *KpiRouter) InitKpiRouter(Router *gin.RouterGroup) {
	kpiRouter := Router.Group("kpi").Use(middleware.OperationRecord())
	kpiRouterWithoutRecord := Router.Group("kpi")
	var kpiApi = v1.ApiGroupApp.SystemApiGroup.KpiApi

	// Category
	{
		kpiRouter.POST("createCategory", kpiApi.CreateCategory)
		kpiRouter.DELETE("deleteCategory", kpiApi.DeleteCategory)
		kpiRouter.DELETE("deleteCategoryByIds", kpiApi.DeleteCategoryByIds)
		kpiRouter.PUT("updateCategory", kpiApi.UpdateCategory)
	}
	{
		kpiRouterWithoutRecord.GET("findCategory", kpiApi.FindCategory)
		kpiRouterWithoutRecord.GET("getCategoryList", kpiApi.GetCategoryList)
	}

	// Item
	{
		kpiRouter.POST("createItem", kpiApi.CreateItem)
		kpiRouter.DELETE("deleteItem", kpiApi.DeleteItem)
		kpiRouter.DELETE("deleteItemByIds", kpiApi.DeleteItemByIds)
		kpiRouter.PUT("updateItem", kpiApi.UpdateItem)
	}
	{
		kpiRouterWithoutRecord.GET("findItem", kpiApi.FindItem)
		kpiRouterWithoutRecord.GET("getItemList", kpiApi.GetItemList)
	}
}
