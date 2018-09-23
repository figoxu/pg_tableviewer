package pg

type ColumnInfo struct {
	Description string
	Relname     string
	Nspname     string
	Attname     string
	Comment     string
}

func GetAllColumnInfoesByTableName(tableName, conStr string) []ColumnInfo {
	columnInfoes := make([]ColumnInfo, 0)
	query := `SELECT
	pd.description,
	tab.relname,
	tab.nspname,
	pa.attname,
	'comment on COLUMN ' || tab.nspname || '.' || tab.relname || '.' || pa.attname || ' is ''' || pd.description || ''';' AS cmd
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
		AND pc.relname =?
	) tab
LEFT OUTER JOIN pg_description pd ON tab.oid = pd.objoid
LEFT OUTER JOIN pg_attribute pa ON pa.attrelid = pd.objoid
ORDER BY
	relname`
	db := initPgByConStr(conStr)
	db.Raw(query, tableName).Scan(&columnInfoes)
	return columnInfoes
}
