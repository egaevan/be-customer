package models

type Customer struct {
	Id        int    `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Address   string `gorm:"column:address"`
}

func (Customer) TableName() string {
	return `Customers`
}
