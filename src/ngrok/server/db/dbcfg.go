package db

import (
	"database/sql"
	"fmt"
	"ngrok/log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var Db, _ = sql.Open("mysql", "root:@/ngrokdb")

var StaticProxy = make(map[string]string)

func InitStaticProxy() {
	rows, err := Db.Query("select host,proto,dstHost from staticProxy")
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	defer rows.Close()

	var host string
	var proto string
	var dstHost string
	for rows.Next() {
		err := rows.Scan(&host, &proto, &dstHost)
		if err != nil {
			log.Error(err.Error())
		}
		if dstHost != "" {
			StaticProxy[fmt.Sprintf("%s://%s", proto, strings.ToLower(host))] = dstHost
		}
	}
}

func UpdateStaticProxy(host string) {
	hostL := strings.ToLower(host)
	delete(StaticProxy, "http://"+hostL)
	delete(StaticProxy, "https://"+hostL)

	rows, err := Db.Query("select proto,dstHost from staticProxy where host=?", host)
	if err != nil {
		log.Error(fmt.Sprintf("UpdateStaticProxy faild:%s  ==> %v", host, err.Error()))
		return
	}
	defer rows.Close()

	var proto string
	var dstHost string
	for rows.Next() {
		err := rows.Scan(&proto, &dstHost)
		if err != nil {
			log.Error(err.Error())
		}
		if dstHost != "" {
			StaticProxy[fmt.Sprintf("%s://%s", proto, hostL)] = dstHost
		}
	}
}
