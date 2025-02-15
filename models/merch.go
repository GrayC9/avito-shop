package models

type Merch struct {
	MerchID uint   `gorm:"column:merch_id;primaryKey;autoIncrement"`
	Name    string `gorm:"column:name;type:text;not null"`
	Price   int    `gorm:"column:price;type:int;not null"`
}

func (m Merch) TableName() string {
	return "merch"
}

type MerchUser struct {
	ID      uint `gorm:"column:id;primaryKey;autoIncrement"`
	MerchID uint `gorm:"column:merch_id;type:int;not null"`
	UserID  int  `gorm:"column:user_id;type:int;not null"`
}

func (m MerchUser) TableName() string {
	return "merch_user"
}
