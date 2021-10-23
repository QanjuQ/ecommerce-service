package config

import (
	"fmt"
)

const (
	sslmode  = "disable"
	timezone = "Asia/Kolkata"
)

type DBConfig struct {
	Name     string
	Host     string
	Username string
	Password string
	Port     int
}

func (db *DBConfig) GetConnectionURI() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s port=%d TimeZone=%s", db.Host, db.Username, db.Name, db.Password, sslmode, db.Port, timezone)
}

