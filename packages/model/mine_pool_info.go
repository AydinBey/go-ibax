package model

import (
	"github.com/shopspring/decimal"
)

//MinePoolInfo example
type MinePoolInfo struct {
	Id               int64           `gorm:"not null" ` //
	Poolid           int64           `gorm:"not null"`  //
	LogoId           int64           `gorm:"not null" ` //logo
	Name             string          `gorm:"not null" ` //
	SettlementType   int64           `gorm:"not null" ` //  1 pps   2  pplns
	SettlementRate   float64         `gorm:"not null" ` //
	SettlementAmount decimal.Decimal `gorm:"not null `  //
	SettlementCycle  int64           `gorm:"not null" ` //
	Status           int64           `gorm:"not null" ` //
	HomeUrl          string          `gorm:"null" `     //
	Date_created     int64           `gorm:"not null" ` //
	var pools []MinePoolInfo
	err := GetDB(dbt).Table(m.TableName()).Find(&pools).Error
	if err != nil {
		return pools, err
	}
	return pools, err
}