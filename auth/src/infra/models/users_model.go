package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : efishery/auth
 */

import "time"

type Users struct {
	ID        int64     `gorm:"id"`
	Name      string    `gorm:"name"`
	Phone     string    `gorm:"phone"`
	Role      string    `gorm:"role"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
