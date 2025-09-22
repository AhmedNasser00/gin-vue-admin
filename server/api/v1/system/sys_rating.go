package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RatingApi struct{}

// CreateRating 创建SysRating
func (ratingApi *RatingApi) CreateRating(c *gin.Context) {
	var rating system.SysRating
	err := c.ShouldBindJSON(&rating)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ratingService.CreateRating(rating); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRating 删除SysRating
func (ratingApi *RatingApi) DeleteRating(c *gin.Context) {
	var rating system.SysRating
	err := c.ShouldBindJSON(&rating)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ratingService.DeleteRating(rating); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRatingByIds 批量删除SysRating
func (ratingApi *RatingApi) DeleteRatingByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ratingService.DeleteRatingByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRating 更新SysRating
func (ratingApi *RatingApi) UpdateRating(c *gin.Context) {
	var rating system.SysRating
	err := c.ShouldBindJSON(&rating)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ratingService.UpdateRating(rating); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRating 用id查询SysRating
func (ratingApi *RatingApi) FindRating(c *gin.Context) {
	var rating system.SysRating
	err := c.ShouldBindQuery(&rating)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserrating, err := ratingService.GetRating(rating.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerating": reuserrating}, c)
	}
}

// GetRatingList 分页获取SysRating列表
func (ratingApi *RatingApi) GetRatingList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ratingService.GetRatingInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetRatingListAll 获取所有SysRating列表
func (ratingApi *RatingApi) GetRatingListAll(c *gin.Context) {
	if list, err := ratingService.GetRatingInfoListAll(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
