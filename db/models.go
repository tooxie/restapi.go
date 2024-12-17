package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uuid.UUID      `gorm:"type:uuid;primaryKey;default:(gen_random_uuid())"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Role struct {
	Id   uuid.UUID `gorm:"type:uuid;primaryKey;default:(gen_random_uuid())"`
	Name string
}

type User struct {
	BaseModel
	Email string `json:"email"`
	Hash  string `json:"-"`
	Roles []Role `gorm:"many2many:user_roles"`
}

type Book struct {
	BaseModel
	Title   string `json:"title"`
	Author  string `json:"author"`
	Slug    string `json:"slug"`
	OwnerId uuid.UUID
	Owner   User `json:"-";gorm:"foreignKey:OwnerId"`
}
