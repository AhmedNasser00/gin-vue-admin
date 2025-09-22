package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RatingRouter struct{}

func (s *RatingRouter) InitRatingRouter(Router *gin.RouterGroup) {
	ratingRouter := Router.Group("rating").Use(middleware.OperationRecord())
	ratingRouterWithoutRecord := Router.Group("rating")
	var ratingApi = v1.ApiGroupApp.SystemApiGroup.RatingApi
	{
		ratingRouter.POST("createRating", ratingApi.CreateRating)             // 新建SysRating
		ratingRouter.DELETE("deleteRating", ratingApi.DeleteRating)           // 删除SysRating
		ratingRouter.DELETE("deleteRatingByIds", ratingApi.DeleteRatingByIds) // 批量删除SysRating
		ratingRouter.PUT("updateRating", ratingApi.UpdateRating)              // 更新SysRating
	}
	{
		ratingRouterWithoutRecord.GET("findRating", ratingApi.FindRating)       // 根据ID获取SysRating
		ratingRouterWithoutRecord.GET("getRatingList", ratingApi.GetRatingList) // 获取SysRating列表
		ratingRouterWithoutRecord.GET("getRatingListAll", ratingApi.GetRatingListAll)
	}
}
