package entities

// "gorm.io/gorm"

type Test struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
