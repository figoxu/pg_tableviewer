package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type PgDbInfo struct {
	Id       int    `json:id"`
	User     string `json:user"`
	Password string `json:password"`
	Dbname   string `json:dbname"`
	Host     string `json:host"`
	Port     int    `json:"port"`
}

func (p *PgDbInfo) ConStr() string {
	return fmt.Sprint("user=", p.User, " password=", p.Password, " dbname=", p.Dbname, " host=", p.Host, " Port=", p.Port, " sslmode=disable")
}

type PgDbInfoDao struct {
	db *gorm.DB
}

func NewPgDbInfoDao(db *gorm.DB) *PgDbInfoDao {
	return &PgDbInfoDao{
		db: db,
	}
}

func (p *PgDbInfoDao) GetByKey(id int) *PgDbInfo {
	pgDbInfo := &PgDbInfo{}
	p.db.Raw("SELECT * FROM pg_db_info WHERE id=? ", id).Scan(pgDbInfo)
	return pgDbInfo
}

func (p *PgDbInfoDao) Save(dbInfo *PgDbInfo) {
	p.db.Save(dbInfo)
}

func (p *PgDbInfoDao) Count() int {
	return 0
}

func (p *PgDbInfoDao) Paging(start, pageSize int) []PgDbInfo {
	pgDbInfos := make([]PgDbInfo, 0)
	p.db.Raw("SELECT * FROM pg_db_info ORDER BY ID LIMIT ? OFFSET ?", pageSize, start).Scan(&pgDbInfos)
	return pgDbInfos
}

func (p *PgDbInfoDao) DelById(id int) {
	p.db.Exec("DELETE FROM pg_db_info WHERE id=? ", id)
}
