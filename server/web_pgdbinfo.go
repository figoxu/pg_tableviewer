package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func h_pgdbinfo_add(c *gin.Context){
	pgDbInfo := PgDbInfo{}
	c.BindJSON(&pgDbInfo)
	userGroupDao := NewPgDbInfoDao(sysEnv.db)
	userGroupDao.Save(&pgDbInfo)
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func h_pgdbinfo_update(c *gin.Context){
	pgDbInfo := PgDbInfo{}
	c.BindJSON(&pgDbInfo)
	userGroupDao := NewPgDbInfoDao(sysEnv.db)
	userGroupDao.Save(&pgDbInfo)
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func h_pgdbinfo_list(c *gin.Context){
	env := c.MustGet("env").(*Env)
	ph := env.ph
	pg, size := ph.Int("pg"), ph.Int("size")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.db)
	total := pgDbInfoDao.Count()
	totalPg := (total + size - 1) / size
	if pg <= 0 {
		pg = 1
	} else if pg > totalPg {
		pg = totalPg
	}
	start := (pg - 1) * size
	pgDbInfoes := pgDbInfoDao.Paging(start, size)

	result := make(map[string]interface{})
	result["total"] = total
	result["totalPg"] = totalPg
	result["curPg"] = pg
	result["data"] = pgDbInfoes
	c.JSON(http.StatusOK, result)
}

func h_pgdbinfo_del(c *gin.Context){
	env := c.MustGet("env").(*Env)
	id := env.ph.Int("id")
	pgDbInfoDao := NewPgDbInfoDao(sysEnv.db)
	pgDbInfoDao.DelById(id)
}
