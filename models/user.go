package models

type User struct {
	Id uint `gorm:"primary_key"`
	Name string
	Pass string
	CreatedAt uint
	UpdatedAt uint
}

func (User) TableName() string {
	return "user"
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.First(&user, "name = ?", username).Error
	return &user, err
}
