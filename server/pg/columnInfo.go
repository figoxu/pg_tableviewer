package pg

type ColumnInfo struct {
	Dbid        int       `json:"dbid"`
	Description string    `json:"description"`
	Relname     string    `json:"relname"`
	Nspname     string    `json:"nspname"`
	Attname     string    `json:"attname"`
	Comment     string    `json:"comment"`
	TableInfo   TableInfo `json:"tableinfo"`
}

func (p *ColumnInfo) fixTableInfo(dbid int, tableName, conStr string) {
	k := tableInfoKey(dbid, tableName)
	v := cache.Get(k)
	if v != nil {
		p.TableInfo = v.(TableInfo)
		return
	}
	GetAllTableInfoes(dbid, conStr)
	v = cache.Get(k)
	p.TableInfo = v.(TableInfo)
}

func GetAllColumnInfoes(dbid int, conStr string) []ColumnInfo {
	columnInfoes := make([]ColumnInfo, 0)
	query := `with att as (
	select
		tab.relname,
		tab.nspname,
		pa.attname,
		pa.attnum,
		pa.attrelid
	 from

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
		) tab
	LEFT OUTER JOIN pg_attribute pa ON tab.oid = pa.attrelid
), cmt as (
	select att.*,pd.description from att,pg_description pd where pd.objoid=att.attrelid and pd.objsubid=att.attnum
), emp as (
	select att.*,'' description from att where att.attrelid ||'_'||att.attnum not in ( select objoid||'_'|| objsubid from  pg_description) and att.attnum>0
) , res as (
	select * from emp union select * from cmt
)
select * from res order by relname,attnum
`
	db := initPgByConStr(conStr)
	db.Raw(query).Scan(&columnInfoes)
	infoes := make([]ColumnInfo, 0)
	for _, cloumnInfo := range columnInfoes {
		cloumnInfo.Dbid = dbid
		cloumnInfo.fixTableInfo(dbid, cloumnInfo.Relname, conStr)
		infoes = append(infoes, cloumnInfo)
	}
	return infoes
}
