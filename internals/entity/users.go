package entity

type Users struct {
	ID        int    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserName  string `gorm:"column:user_name"`
	Password  string `gorm:"column:password"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	RoleCode  string `gorm:"column:role_code"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
	CreatedBy string `gorm:"column:created_by"`
	UpdatedBy string `gorm:"column:updated_by"`
}

func (Users) TableName() string {
	return "users"
}
