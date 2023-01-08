package entity

type Token struct {
	ID        int    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserName  string `gorm:"column:user_name"`
	Token     string `gorm:"column:token"`
	Status    string `gorm:"column:status"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
	CreatedBy string `gorm:"column:created_by"`
	UpdatedBy string `gorm:"column:updated_by"`
}

func (Token) TableName() string {
	return "token"
}
