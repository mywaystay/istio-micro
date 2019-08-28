package cinit

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
)

var Pg *sqlx.DB

//初始化连接
func pgInit() {
	var err error
	dataSourceName := "postgres://" + Config.Postgres.User + ":" + Config.Postgres.Password + "@" + Config.Postgres.Addr + ":" + strconv.Itoa(Config.Postgres.Port) +
		"/" + Config.Postgres.DbName + "?sslmode=disable"
	Pg, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	Pg.SetMaxIdleConns(Config.Postgres.IdleConn)
	Pg.SetMaxOpenConns(Config.Postgres.MaxConn)
	err = Pg.Ping()
	if err != nil {
		panic(err)
	}
}

//关闭
func pgClose() {
	Pg.Close()
}
