package errors

type dbError struct {
	string
}

func (e dbError) Error() string {
	return e.string
}

func ErrDbCreateTableFail() error {
	return dbError{"数据库建表失败"}

}
