package server

import "github.com/jinzhu/gorm"

type PgDbInfo struct {
	User     string
	Password string
	Dbname   string
	Host     string
	Port     int
}

type PgDbInfoDao struct {
	db *gorm.DB
}

func NewPgDbInfoDao(db *gorm.DB) *PgDbInfoDao {
	return &PgDbInfoDao{
		db: db,
	}
}

func (p *PgDbInfoDao) Save(dbInfo *PgDbInfo){
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


