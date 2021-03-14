package model

type Profile struct {
	ID        string `gorm:"column:id"`
	Age       uint64 `gorm:"column:age"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	UserID    string `gorm:"column:userId"`
}

func (Profile) TableName() string {
	return "profile"
}
