package db_types

import (
	"context"

	"gorm.io/gorm"
)

type GormOpener interface {
	Open(dsn string, config *gorm.Config) (*gorm.DB, error)
}

type DBProvider interface {
	GetDB(ctx context.Context) (*gorm.DB, error)
}

type FilterFunc func(db *gorm.DB) *gorm.DB

type SortData struct {
	SortField string
	SortOrder string
}

type PaginationData struct {
	Page     int
	PageSize int
}

type Repository[DbModel any] interface {
	GetList(ctx context.Context, filters []FilterFunc, sort SortData, pagination PaginationData) ([]*DbModel, int64, error)
	Get(ctx context.Context, id uint64) (*DbModel, error)
	Create(ctx context.Context, newItem *DbModel) (*DbModel, error)
	Update(ctx context.Context, id uint64, updatedItem *DbModel) (*DbModel, error)
	Delete(ctx context.Context, id uint64) error
}

type HasAssociations[DbModel any] interface {
	UpdateAssociations(db *gorm.DB, item *DbModel, updatedData *DbModel) error
}
