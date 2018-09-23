package pg

func Comment(comment,conStr string){
	db := initPgByConStr(conStr)
	db.Exec(comment)
}
