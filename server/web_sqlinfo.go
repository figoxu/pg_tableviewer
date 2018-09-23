package server

import (
	"github.com/gin-gonic/gin"
	"github.com/figoxu/pg_tableviewer/server/pg"
	"net/http"
)

func h_tableInfo_all(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	dbid := env.ph.Int("dbid")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.db)
	pgDbInfo := pgDbInfoDao.GetByKey(dbid)
	infoes := pg.GetAllTableInfoes(pgDbInfo.ConStr())
	c.JSON(http.StatusOK, infoes)
}

func h_columnInfo_by_tablename(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	dbid := env.ph.Int("dbid")
	tablename := env.fh.String("tablename")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.db)
	pgDbInfo := pgDbInfoDao.GetByKey(dbid)
	infoes := pg.GetAllColumnInfoesByTableName(tablename, pgDbInfo.ConStr())
	c.JSON(http.StatusOK, infoes)
}
