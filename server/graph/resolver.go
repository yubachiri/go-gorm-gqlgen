package graph

import "github.com/jinzhu/gorm"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver this is resolver
type Resolver struct {
	DB *gorm.DB
}
