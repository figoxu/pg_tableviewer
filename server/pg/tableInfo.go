package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/figoxu/Figo"
	"time"
	"github.com/quexer/utee"
)

type TableInfo struct {
	Description string
	Relname     string
	Nspname     string
	Comment     string
}

func GetAllTableInfoes(conStr string) []TableInfo {
	tableInfoes := make([]TableInfo, 0)
	query := `SELECT
	pd.description,
	tab.relname,
	tab.nspname,
	'comment on table ' || tab.nspname || '.' || tab.relname || ' is ''' || pd.description || ''';' AS cmd
FROM
	(
		SELECT
			pc.oid,
			pc.relname,
			pn.nspname
		FROM
			pg_class pc
		LEFT OUTER JOIN pg_namespace pn ON pc.relnamespace = pn.oid
		WHERE
			pc.relnamespace = 2200
		AND pc.relkind IN('r')
		AND pc.oid NOT IN(SELECT inhrelid FROM pg_inherits)
		AND pc.relname NOT LIKE '%peiyb%'
	) tab
LEFT OUTER JOIN pg_description pd ON tab.oid = pd.objoid
AND pd.objsubid = 0
ORDER BY
	relname`
	db := initPgByConStr(conStr)
	db.Raw(query).Scan(&tableInfoes)
	return tableInfoes
}

func initPgByConStr(conStr string) *gorm.DB {
	db, err := gorm.Open("postgres", conStr)
	db.SetLogger(&Figo.GormLog{})
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)
	utee.Chk(err)
	db.LogMode(true)
	return db
}
