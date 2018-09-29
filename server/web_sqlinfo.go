package main

import (
	"github.com/gin-gonic/gin"
	"github.com/figoxu/pg_tableviewer/server/pg"
	"net/http"
	"fmt"
)

func h_tableInfo_all(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	dbid := env.ph.Int("dbid")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.Db)
	pgDbInfo := pgDbInfoDao.GetByKey(dbid)
	infoes := pg.GetAllTableInfoes(dbid, pgDbInfo.ConStr())
	c.JSON(http.StatusOK, infoes)
}

// curl http://localhost:8080/api/pg_db_info/table_info/11/columns
func h_columnInfo_all(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	dbid := env.ph.Int("dbid")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.Db)
	pgDbInfo := pgDbInfoDao.GetByKey(dbid)
	infoes := pg.GetAllColumnInfoes(dbid, pgDbInfo.ConStr())
	c.JSON(http.StatusOK, infoes)
}

func h_comment_table(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	dbid := env.ph.Int("dbid")
	comment := env.fh.String("comment")
	relname := env.fh.String("relname")
	nspname := env.fh.String("nspname")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.Db)
	pgDbInfo := pgDbInfoDao.GetByKey(dbid)
	sql := fmt.Sprint("comment on table ", nspname, ".", relname, " is '", comment, "'")
	fmt.Println("COMMENT:  >>>  ",sql)
	pg.Comment(sql, pgDbInfo.ConStr())
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func h_comment_column(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	dbid := env.ph.Int("dbid")
	comment := env.fh.String("comment")
	attname := env.fh.String("attname")
	relname := env.fh.String("relname")
	nspname := env.fh.String("nspname")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.Db)
	pgDbInfo := pgDbInfoDao.GetByKey(dbid)
	sql := fmt.Sprint("comment on COLUMN ", nspname, ".", relname, ".", attname, " is '", comment, "'")
	fmt.Println("COMMENT:  >>>  ",sql)
	pg.Comment(sql, pgDbInfo.ConStr())
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
