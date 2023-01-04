package model

type Inventory struct {
	BaseModel
	Goods int32 `gorm:"type:int;index"`
	Stocks int32 `gorm:"type:int32"`
	Version int32 `gorm:"type:int32"`
}
