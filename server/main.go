package server

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
)

var sysEnv = SysEnv{}

type SysEnv struct {
	path_dist string
	port      string
	db_path   string
	db        *gorm.DB
}

func main() {
	log.Println("Hello World")
	initConf()
	initWeb(sysEnv.port)
}

func initConf() {
	cfg_core, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	sysEnv.path_dist = cfg_core.String("http::dist")
	sysEnv.port = cfg_core.String("http::port")
	sysEnv.db_path = cfg_core.String("db_sqlite::path")

	sqlitedb, err := gorm.Open("sqlite3", sysEnv.db_path)
	utee.Chk(err)
	sqlitedb.DB().SetConnMaxLifetime(time.Minute * 5)
	sqlitedb.DB().SetMaxIdleConns(0)
	sqlitedb.DB().SetMaxOpenConns(5)
	sqlitedb.SingularTable(true)
	sqlitedb.Debug().AutoMigrate(&PgDbInfo{})
	sysEnv.db = sqlitedb
}

func initWeb(port string) {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(sysEnv.port, nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	store := cookie.NewStore([]byte("xujh945@qq.com"))
	r.Use(sessions.Sessions("figoxu", store))
	r.Use(static.Serve("/", static.LocalFile(sysEnv.path_dist, true)))
	api := r.Group("/api")
	{
		pg_db_info := api.Group("/pg_db_info")
		{
			pg_db_info.POST("/add", h_pgdbinfo_add)
			pg_db_info.POST("/update", h_pgdbinfo_update)
			pg_db_info.GET("/list/:size/:pg", h_pgdbinfo_list)
			pg_db_info.POST("/del/:id", h_pgdbinfo_del)
			table_info := pg_db_info.Group("/table_info/:dbid")
			{
				table_info.GET("/tables",h_tableInfo_all)
				table_info.GET("/columns",h_columnInfo_by_tablename)
				table_info.PUT("/comment",h_comment)
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
