package db

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     int
	UserName string
	PassWord string
	DbName   string
}

func (s *Sql) Connect() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.UserName, s.PassWord, s.Host, s.Port, s.DbName)
	s.Db = sqlx.MustConnect("mysql", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("connect sus")
}

func (s *Sql) Close() {
	s.Db.Close()
}
func  NewSqlConfig() *Sql {
	return &Sql{
		Host: os.Getenv("DB_HOST"),
		Port: func() int {
			port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
			return port
		}(),
		UserName: os.Getenv("DB_USER"),
		PassWord: os.Getenv("DB_PASS"),
		DbName:   os.Getenv("DB_DATABASE"),
	}
}
