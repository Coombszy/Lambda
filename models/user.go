package model

type (
	User struct {
		Name     string `json:"name" gorm:"primaryKey"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
