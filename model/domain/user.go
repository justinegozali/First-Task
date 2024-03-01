package domain

type User struct {
	UserId       int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	UserName     string `gorm:"column:user_name"`
	UserEmail    string `gorm:"column:user_email"`
	UserPassword string `gorm:"column:user_password"`
	UserRole     int    `gorm:"column:user_role"`
	Role         Role   `gorm:"foreignKey:role_id;references:user_role"`
}

func (u *User) TableName() string {
	return "users"
}
