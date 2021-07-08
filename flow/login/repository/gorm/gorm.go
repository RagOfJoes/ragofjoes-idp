package gorm

import (
	"github.com/RagOfJoes/idp/flow/login"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type gormLoginRepository struct {
	DB *gorm.DB
}

func NewGormLoginRepository(d *gorm.DB) login.Repository {
	return &gormLoginRepository{DB: d}
}

func (g *gormLoginRepository) Create(l login.Login) (*login.Login, error) {
	n := l
	if err := g.DB.Create(&n).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

func (g *gormLoginRepository) Get(i string) (*login.Login, error) {
	var n login.Login
	if err := g.DB.First(&n, "id = ?", i).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

func (g *gormLoginRepository) GetByFlowID(i string) (*login.Login, error) {
	var n login.Login
	if err := g.DB.First(&n, "flow_id = ?", i).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

func (g *gormLoginRepository) Update(l login.Login) (*login.Login, error) {
	n := l
	if err := g.DB.Save(&n).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

func (g *gormLoginRepository) Delete(i uuid.UUID) error {
	if err := g.DB.Where("id = ?", i.String()).Delete(login.Login{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
