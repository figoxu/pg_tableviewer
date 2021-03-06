package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type PgDbInfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func (p *PgDbInfo) ConStr() string {
	return fmt.Sprint("user=", p.User, " password=", p.Password, " dbname=", p.Dbname, " host=", p.Host, " port=", p.Port, " sslmode=disable")
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

type CountVo struct {
	Val int
}

func (p *PgDbInfoDao) Count() int {
	countVo := CountVo{}
	p.db.Raw("SELECT COUNT(*) val FROM pg_db_info  ").Scan(&countVo)
	return countVo.Val
}

func (p *PgDbInfoDao) Paging(start, pageSize int) []PgDbInfo {
	pgDbInfos := make([]PgDbInfo, 0)
	p.db.Raw("SELECT * FROM pg_db_info ORDER BY ID LIMIT ? OFFSET ?", pageSize, start).Scan(&pgDbInfos)
	return pgDbInfos
}

func (p *PgDbInfoDao) DelById(id int) {
	p.db.Exec("DELETE FROM pg_db_info WHERE id=? ", id)
}

func (p *PgDbInfoDao) All() []PgDbInfo {
	pgDbInfoes := make([]PgDbInfo, 0)
	p.db.Raw("SELECT * FROM pg_db_info").Scan(&pgDbInfoes)
	return pgDbInfoes
}
