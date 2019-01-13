package resolver

import "github.com/jinzhu/gorm"

type RootResolver struct {
	db *gorm.DB
}

func NewRootResolver(db *gorm.DB) *RootResolver {
	return &RootResolver{db: db}
}
