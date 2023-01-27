package config

import "fmt"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		DBName:   "online_learning_database",
		Password: "password",
	}
}

func BuildDBConfigForTest() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		DBName:   "online_learning_database",
		Password: "password",
	}
}

func (dbConfig *DBConfig) DBUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
