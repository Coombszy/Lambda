package model

type (
	User struct {
		Name            string `json:"name" gorm:"primaryKey"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		User_Group      string `json:"user_group"`
		Workspace_Group string `json:"workspace_group"`
	}
)
