package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/srahkmli/grpcTest/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductImpl {
	return &ProductImpl{
		db: db,
	}
}

func (repo *ProductImpl) UpsertProduct(product *model.Product) error {
	tx := repo.db.Table("product").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "qty"}),
	}).Create(product)

	return tx.Error
}

func (repo *ProductImpl) FindProduct(name string) model.Product {
	var out model.Product
	if tx := repo.db.Table("product").First(&out, name); tx.Error != nil {
		logrus.Errorf("Failed to find product due to %v", tx.Error)
	}
	return out
}

func (repo *ProductImpl) DeleteProduct(name string) error {
	whereCondition := model.Product{
		Name: name,
	}
	tx := repo.db.Table("product").Delete(&whereCondition)
	return tx.Error
}
