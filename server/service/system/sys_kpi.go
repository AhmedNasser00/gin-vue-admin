package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type KpiService struct{}

var KpiServiceApp = new(KpiService)

// ---- validation helpers ----
func (s *KpiService) validateCategory(c sysModel.SysKpiCategory) error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Weight == nil {
		return errors.New("weight is required")
	}
	if *c.Weight < 0 || *c.Weight > 100 {
		return errors.New("weight must be between 0 and 100")
	}
	return nil
}

func (s *KpiService) validateItem(i sysModel.SysKpiItem) error {
	if i.CategoryID == 0 {
		return errors.New("categoryID is required")
	}
	if i.Name == "" {
		return errors.New("name is required")
	}
	if i.Weight == nil {
		return errors.New("weight is required")
	}
	if *i.Weight < 0 || *i.Weight > 100 {
		return errors.New("weight must be between 0 and 100")
	}
	return nil
}

// ---- category CRUD ----
func (s *KpiService) CreateCategory(c sysModel.SysKpiCategory) error {
	if err := s.validateCategory(c); err != nil {
		return err
	}
	var cnt int64
	if err := global.GVA_DB.Model(&sysModel.SysKpiCategory{}).Where("name = ?", c.Name).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("category name already exists")
	}
	return global.GVA_DB.Create(&c).Error
}

func (s *KpiService) DeleteCategory(c sysModel.SysKpiCategory) error {
	return global.GVA_DB.Delete(&c).Error
}

func (s *KpiService) UpdateCategory(c sysModel.SysKpiCategory) error {
	if err := s.validateCategory(c); err != nil {
		return err
	}
	var cnt int64
	if err := global.GVA_DB.Model(&sysModel.SysKpiCategory{}).Where("name = ? AND id <> ?", c.Name, c.ID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("category name already exists")
	}
	return global.GVA_DB.Save(&c).Error
}

func (s *KpiService) GetCategory(id uint) (sysModel.SysKpiCategory, error) {
	var c sysModel.SysKpiCategory
	err := global.GVA_DB.First(&c, id).Error
	return c, err
}

func (s *KpiService) GetCategoryList(page, pageSize int) (list []sysModel.SysKpiCategory, total int64, err error) {
	db := global.GVA_DB.Model(&sysModel.SysKpiCategory{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset((page - 1) * pageSize).Order("id desc").Find(&list).Error
	return
}

// ---- item CRUD ----
func (s *KpiService) CreateItem(i sysModel.SysKpiItem) error {
	if err := s.validateItem(i); err != nil {
		return err
	}
	var cnt int64
	if err := global.GVA_DB.Model(&sysModel.SysKpiItem{}).Where("category_id = ? AND name = ?", i.CategoryID, i.Name).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("duplicate item name in category")
	}
	return global.GVA_DB.Create(&i).Error
}

func (s *KpiService) DeleteItem(i sysModel.SysKpiItem) error {
	return global.GVA_DB.Delete(&i).Error
}

func (s *KpiService) UpdateItem(i sysModel.SysKpiItem) error {
	if err := s.validateItem(i); err != nil {
		return err
	}
	var cnt int64
	if err := global.GVA_DB.Model(&sysModel.SysKpiItem{}).Where("category_id = ? AND name = ? AND id <> ?", i.CategoryID, i.Name, i.ID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("duplicate item name in category")
	}
	return global.GVA_DB.Save(&i).Error
}

func (s *KpiService) GetItem(id uint) (sysModel.SysKpiItem, error) {
	var i sysModel.SysKpiItem
	err := global.GVA_DB.First(&i, id).Error
	return i, err
}

func (s *KpiService) GetItemList(categoryID uint, page, pageSize int) (list []sysModel.SysKpiItem, total int64, err error) {
	db := global.GVA_DB.Model(&sysModel.SysKpiItem{})
	if categoryID > 0 {
		db = db.Where("category_id = ?", categoryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset((page - 1) * pageSize).Order("id desc").Find(&list).Error
	return
}
