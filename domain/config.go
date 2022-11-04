package domain

import "time"

type DatabaseConfig struct {
	User              string
	Password          string
	Database          string
	Host              string
	Port              string
	TimeZone          string
	MaxConnectionIdle int
	MaxConnectionOpen int
	Schema            string
	Debug             bool
	TimeOut           time.Duration
}
