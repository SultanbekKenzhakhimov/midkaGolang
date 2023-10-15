package models

import "gorm.io/gorm"

type PowerTool struct {
	gorm.Model
	Name  string
	Power int
}

type Paint struct {
	gorm.Model
	Color string
	Type  string
}

type NailScrew struct {
	gorm.Model
	Type     string
	Size     string
	Material string
}

type PlumbingSupply struct {
	gorm.Model
	Name   string
	Type   string
	Length int
}

type ElectricalFixture struct {
	gorm.Model
	Name    string
	Type    string
	Wattage int
}
