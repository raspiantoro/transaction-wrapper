package model

type User struct {
	ID       string `gorm:"primaryKey"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Profile  Profile
}

func (User) TableName() string {
	return "user"
}
