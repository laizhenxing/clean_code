package persistence

import (
	"database/sql"
	"fmt"
)

const (
	maxOpenConn = 20
	maxIdleConn  = 10
)

type Driver string

func (d *Driver) String() string {
	return string(*d)
}

const (
	DriverMysql    Driver = "mysql"
	DriverMongo    Driver = "mongo"
	DriverPostgres Driver = "postgres"
	DriverSQLite3  Driver = "sqlite3"
)

type Options struct {
	User     string
	Pwd      string
	Host     string
	Port     int
	Database string
}

func GetMysqlDsn(o *Options) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=true&charset=utf8mb4&loc=Local",
		o.User,
		o.Pwd,
		o.Host,
		o.Port,
		o.Database,
	)
}

func GetMongoDsn(o *Options) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?connect=direct",
		o.User,
		o.Pwd,
		o.Host,
		o.Port,
		o.Database,
	)
}

func GetPostgreDsn(o *Options) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&charset=utf8",
		o.User,
		o.Pwd,
		o.Host,
		o.Port,
		o.Database,
	)
}

func GetSQLLiteDsn(file, user, password string) string {
	return fmt.Sprintf(
		"file:%s?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=sha1",
		file,
		user,
		password,
	)
}

// DB 基于dbr的数据库封装
type DB struct {
	conn     *sql.DB
	dsn      string
}

var Db *DB

func NewDB(driver Driver, dsn string) (*DB, error) {
	Db = &DB{
		dsn: dsn,
	}

	conn, err := sql.Open(driver.String(), dsn)
	if err != nil {
		return nil, err
	}

	Db.conn = conn
	// 最大连接数
	Db.conn.SetConnMaxLifetime(maxOpenConn)
	// 最大空闲连接数，空闲连接数应该小于最大连接数，否则被退化为与最大连接数相等
	Db.conn.SetMaxIdleConns(maxIdleConn)
	if err = Db.conn.Ping(); err != nil {
		return nil, err
	}

	return Db, nil
}

func (d *DB) Close() error {
	return d.conn.Close()
}

func (d *DB) GetConn() *sql.DB {
	return d.conn
}
