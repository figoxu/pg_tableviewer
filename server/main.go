package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/static"
	"net/http"
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
	"github.com/figoxu/gh"
	"github.com/jinzhu/gorm"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"github.com/figoxu/Figo"
)

var sysEnv = SysEnv{}

type SysEnv struct {
	Path_dist string
	Port      string
	Db_path   string
	Db        *gorm.DB
}

func main() {
	log.Println("Hello World")
	initConf()
	fmt.Println(">>>> ", Figo.JsonString(sysEnv))
	initWeb(sysEnv.Port)
}

func initConf() {
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	sysEnv.Path_dist = cfg_core.String("http::dist")
	sysEnv.Port = cfg_core.String("http::Port")
	sysEnv.Db_path = cfg_core.String("db_sqlite::path")

	sqlitedb, err := gorm.Open("sqlite3", sysEnv.Db_path)
	utee.Chk(err)
	sqlitedb.DB().SetConnMaxLifetime(time.Minute * 5)
	sqlitedb.DB().SetMaxIdleConns(0)
	sqlitedb.DB().SetMaxOpenConns(5)
	sqlitedb.SingularTable(true)
	sqlitedb.SetLogger(&Figo.GormLog{})
	sqlitedb.Debug().AutoMigrate(&PgDbInfo{})
	sysEnv.Db = sqlitedb
}

func initWeb(port string) {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(sysEnv.Port, nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	store := cookie.NewStore([]byte("xujh945@qq.com"))
	r.Use(sessions.Sessions("figoxu", store))
	r.Use(static.Serve("/", static.LocalFile(sysEnv.Path_dist, true)))
	api := r.Group("/api", m_gh)
	{
		pg_db_info := api.Group("/pg_db_info")
		{
			pg_db_info.POST("/add", h_pgdbinfo_add)
			pg_db_info.POST("/update", h_pgdbinfo_update)
			pg_db_info.GET("/list/:size/:pg", h_pgdbinfo_list)
			pg_db_info.POST("/del/:id", h_pgdbinfo_del)
			pg_db_info.GET("/all", h_pgdbinfo_all)
			table_info := pg_db_info.Group("/table_info/:dbid")
			{
				table_info.GET("/tables", h_tableInfo_all)
				table_info.GET("/columns", h_columnInfo_all)
				table_info.PUT("/comment/table", h_comment_table)
				table_info.PUT("/comment/column", h_comment_column)
			}
		}
	}
	return r
}

type Env struct {
	fh *gh.FormHelper
	ph *gh.ParamHelper
}

func m_gh(c *gin.Context) {
	c.Set("env", &Env{
		fh: gh.NewFormHelper(c),
		ph: gh.NewParamHelper(c),
	})
	c.Next()
}
