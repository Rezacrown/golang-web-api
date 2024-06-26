package book

import "time"

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    ``
	Author      string    ``
	Description string    ``
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time // Automatically managed by GORM for update time
}
