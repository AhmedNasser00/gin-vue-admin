package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type RatingService struct{}

var RatingServiceApp = new(RatingService)

// CreateRating 创建SysRating记录
func (ratingService *RatingService) CreateRating(rating system.SysRating) (err error) {
	err = global.GVA_DB.Create(&rating).Error
	return err
}

// DeleteRating 删除SysRating记录
func (ratingService *RatingService) DeleteRating(rating system.SysRating) (err error) {
	err = global.GVA_DB.Delete(&rating).Error
	return err
}

// UpdateRating 更新SysRating记录
func (ratingService *RatingService) UpdateRating(rating system.SysRating) (err error) {
	err = global.GVA_DB.Save(&rating).Error
	return err
}

// GetRating 根据id获取SysRating记录
func (ratingService *RatingService) GetRating(id uint) (rating system.SysRating, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&rating).Error
	return
}

// GetRatingInfoList 分页获取SysRating记录
func (ratingService *RatingService) GetRatingInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysRating{})
	var ratings []system.SysRating
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&ratings).Error
	return ratings, total, err
}

// GetRatingInfoListAll 获取所有SysRating记录
func (ratingService *RatingService) GetRatingInfoListAll() (list interface{}, err error) {
	// 创建db
	db := global.GVA_DB.Model(&system.SysRating{})
	var ratings []system.SysRating
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Find(&ratings).Error
	return ratings, err
}

// DeleteRatingByIds 批量删除SysRating记录
func (ratingService *RatingService) DeleteRatingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&[]system.SysRating{}, "id in ?", ids.Ids).Error
		return txErr
	})
	return err
}
