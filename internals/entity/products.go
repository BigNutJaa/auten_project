package entity

type Products struct {
	ID        int    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Name      string `gorm:"column:name"`
	Detail    string `gorm:"column:detail"`
	Qty       int32  `gorm:"column:qty"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
	CreatedBy string `gorm:"column:created_by"`
	UpdatedBy string `gorm:"column:updated_by"`
}

func (Products) TableName() string {
	return "products"
}
