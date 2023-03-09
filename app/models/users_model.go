package models

type User struct {
	ID       uint   `gorm:"primaryKey" db:"id" json:"id" validate:"required"`
	Email    string `db:"email" json:"email" validate:"required,email,lte=255"`
	Password string `db:"password" json:"password" validate:"required,lte=255"`
	UserRole string `db:"role" json:"role" validate:"required,lte=255"`
}
