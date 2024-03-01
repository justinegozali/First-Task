package domain

type Role struct {
	RoleId   int    `gorm:"column:role_id;primaryKey;autoIncrement"`
	RoleName string `gorm:"column:role_name"`
}

func (r *Role) TableName() string {
	return "roles"
}
