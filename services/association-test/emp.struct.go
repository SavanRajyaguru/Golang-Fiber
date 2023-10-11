package associationtest

type SendCardResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"card_name"`
	EmpDetails []Emp  `gorm:"foreignKey:card_id" json:"emp_details,omitempty"`
}

type SendEmpResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"emp_name"`
	CardDetails []Card `gorm:"foreignKey:ID" json:"card_details,omitempty"`
}
