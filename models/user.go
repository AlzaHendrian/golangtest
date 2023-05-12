package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
}
