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
)

var sysEnv = SysEnv{}

type SysEnv struct {
	path_dist string
	port      string
	db_path   string
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
	return r
}
