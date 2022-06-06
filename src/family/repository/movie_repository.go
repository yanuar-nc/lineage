package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/go-boiler-plate/src/family/domain"
	"github.com/yanuar-nc/golang/helper"
	"gorm.io/gorm"
)

// FamilyRepositoryNeo4j struct
type FamilyRepositoryNeo4j struct {
	db *gorm.DB
}

// NewFamilyRepositoryNeo4j function
func NewFamilyRepositoryNeo4j(db *gorm.DB) *FamilyRepositoryNeo4j {
	return &FamilyRepositoryNeo4j{db: db}
}

func (l *FamilyRepositoryNeo4j) Save(ctx context.Context, data *domain.Family) error {
	data.Datetime = time.Now()
	err := l.db.Save(data).Error
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "FamilyRepository", "save")
		return err
	}
	return nil
}
