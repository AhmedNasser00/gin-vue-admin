package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KpiApi struct{}

// ---- Category APIs ----
func (a *KpiApi) CreateCategory(c *gin.Context) {
	var req sysModel.SysKpiCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kpiService.CreateCategory(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a *KpiApi) DeleteCategory(c *gin.Context) {
	var req sysModel.SysKpiCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kpiService.DeleteCategory(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *KpiApi) DeleteCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	if err := c.ShouldBindJSON(&IDS); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(IDS.Ids) == 0 {
		response.Ok(c)
		return
	}
	if err := global.GVA_DB.Delete(&sysModel.SysKpiCategory{}, IDS.Ids).Error; err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (a *KpiApi) UpdateCategory(c *gin.Context) {
	var req sysModel.SysKpiCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kpiService.UpdateCategory(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (a *KpiApi) FindCategory(c *gin.Context) {
	var req sysModel.SysKpiCategory
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reCategory, err := kpiService.GetCategory(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reCategory": reCategory}, c)
	}
}

func (a *KpiApi) GetCategoryList(c *gin.Context) {
	var pageInfo request.PageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := kpiService.GetCategoryList(int(pageInfo.Page), int(pageInfo.PageSize)); err != nil {
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

// ---- Item APIs ----
func (a *KpiApi) CreateItem(c *gin.Context) {
	var req sysModel.SysKpiItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kpiService.CreateItem(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a *KpiApi) DeleteItem(c *gin.Context) {
	var req sysModel.SysKpiItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kpiService.DeleteItem(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *KpiApi) DeleteItemByIds(c *gin.Context) {
	var IDS request.IdsReq
	if err := c.ShouldBindJSON(&IDS); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(IDS.Ids) == 0 {
		response.Ok(c)
		return
	}
	if err := global.GVA_DB.Delete(&sysModel.SysKpiItem{}, IDS.Ids).Error; err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (a *KpiApi) UpdateItem(c *gin.Context) {
	var req sysModel.SysKpiItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kpiService.UpdateItem(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (a *KpiApi) FindItem(c *gin.Context) {
	var req sysModel.SysKpiItem
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reItem, err := kpiService.GetItem(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reItem": reItem}, c)
	}
}

type KpiItemListReq struct {
	request.PageInfo
	CategoryID uint `json:"categoryID" form:"categoryID"`
}

func (a *KpiApi) GetItemList(c *gin.Context) {
	var req KpiItemListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := kpiService.GetItemList(req.CategoryID, int(req.Page), int(req.PageSize)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}
