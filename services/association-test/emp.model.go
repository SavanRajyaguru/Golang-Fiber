package associationtest

type Emp struct {
	ID          uint   `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name        string `gorm:"index" validate:"required" json:"name"`
	CardID      uint   `gorm:"foreignKey:ID" json:"-"`
	CardDetails []Card `gorm:"many2many:emp_cards;"`
}

type Card struct {
	ID        uint   `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name      string `json:"name"`
	Employees []Emp  `gorm:"many2many:emp_cards;"`
}

func (company *Card) TableName() string {
	return "card"
}

// Hook demo
//
//	func (m *Emp) BeforeSave(tx *gorm.DB) (err error) {
//		hexData := hex.EncodeToString([]byte(m.Name))
//		m.Name = hexData
//		return nil
//	}
func (emp *Emp) TableName() string {
	return "emp"
}
