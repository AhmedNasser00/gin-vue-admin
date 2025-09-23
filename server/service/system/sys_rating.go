package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type RatingService struct{}

var RatingServiceApp = new(RatingService)

// validate performs business validations on a rating entity
func (ratingService *RatingService) validate(r system.SysRating) error {
	// rating must be 1..5
	if r.Rating == nil {
		return errors.New("rating is required")
	}
	if *r.Rating < 1 || *r.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	// percentage range required
	if r.MinPercentage == nil || r.MaxPercentage == nil {
		return errors.New("percentage range is required")
	}
	if *r.MinPercentage < 0 || *r.MinPercentage > 100 || *r.MaxPercentage < 0 || *r.MaxPercentage > 100 {
		return errors.New("percentage must be between 0 and 100")
	}
	if *r.MinPercentage > *r.MaxPercentage {
		return errors.New("minPercentage cannot be greater than maxPercentage")
	}
	return nil
}

// checkOverlap ensures that the range does not overlap with other ratings
func (ratingService *RatingService) checkOverlap(r system.SysRating) error {
	var count int64
	// overlap condition: existing.min <= new.max AND existing.max >= new.min
	if err := global.GVA_DB.Model(&system.SysRating{}).
		Where("min_percentage <= ? AND max_percentage >= ?", *r.MaxPercentage, *r.MinPercentage).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("percentage range overlaps with existing rating")
	}
	return nil
}

// checkOverlapExcludeSelf checks overlap excluding the current record id
func (ratingService *RatingService) checkOverlapExcludeSelf(r system.SysRating) error {
	var count int64
	if err := global.GVA_DB.Model(&system.SysRating{}).
		Where("min_percentage <= ? AND max_percentage >= ? AND id <> ?", *r.MaxPercentage, *r.MinPercentage, r.ID).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("percentage range overlaps with existing rating")
	}
	return nil
}

// CreateRating 创建SysRating记录
func (ratingService *RatingService) CreateRating(rating system.SysRating) (err error) {
	// validate fields
	if err = ratingService.validate(rating); err != nil {
		return err
	}
	// uniqueness check for rating value
	var cnt int64
	if err = global.GVA_DB.Model(&system.SysRating{}).Where("rating = ?", rating.Rating).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("rating already exists")
	}
	// overlap check
	if err = ratingService.checkOverlap(rating); err != nil {
		return err
	}
	// create
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
	// validate fields
	if err = ratingService.validate(rating); err != nil {
		return err
	}
	// uniqueness check for rating value (exclude self)
	var cnt int64
	if err = global.GVA_DB.Model(&system.SysRating{}).Where("rating = ? AND id <> ?", rating.Rating, rating.ID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("rating already exists")
	}
	// overlap check (exclude self)
	if err = ratingService.checkOverlapExcludeSelf(rating); err != nil {
		return err
	}
	// update
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
