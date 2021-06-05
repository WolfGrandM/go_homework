package main

import (
	"database/sql"
	"errors"
	"fmt"
)

/*
在数据库操作时,比如dao层中遇到一个sql.errorNoRows时,
	1.是否应该warp这个error抛给上层
	2.为什么
	3.怎么做,用代码实现它
*/

/*
	1.不应该
	2.在该场景下,具有最高可用线性的包只能返回根错误值,否则在会存在多次打印堆栈信息的情况.
	3.在该场景下,故基础库或第三库在会用pkg.error时,应抛出根错误,实现如下
*/



func readDataFromSql(filePath string, db *sql.DB)  *sql.Row {
	err := db.QueryRow( filePath /*sql condition...*/)
	if err !=nil {
		return err
	}
	return nil
}

func main() {
	filepath:="HOME"
	db:=&sql.DB{}
	err:=readDataFromSql(filepath,db)
	if errors.Is(err.Err(),sql.ErrNoRows) {
		fmt.Printf("sql error %s",err.Err())
	}
}