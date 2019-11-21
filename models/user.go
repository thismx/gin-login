package models

type User struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100);unique_index"`
	Pass      string
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
