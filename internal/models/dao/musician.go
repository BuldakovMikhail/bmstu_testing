package dao

type Musician struct {
	ID   uint64 `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (Musician) TableName() string {
	return "musicians"
}
