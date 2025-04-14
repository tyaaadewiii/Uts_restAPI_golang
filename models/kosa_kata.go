package models

type KosaKata struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Kata string `json:"kata"`
	Arti string `json:"arti"`
}

func (KosaKata) TableName() string {
	return "kosa_kata"
}
