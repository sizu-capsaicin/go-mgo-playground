package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlPath = "tcp(127.0.0.1:3306)/"
	userName  = "go"
)

func main() {
	var (
		dbName = flag.String("db", "adb", "Database name")
		passwd = flag.String("pw", "xxxx", "Enter the password")
	)
	flag.Parse()

	path := userName + ":" + *passwd + "@" + mysqlPath

	// MySQL への接続
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// データベースが既に作成済みであれば一度 drop する
	query := "drop database if exists " + *dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	// データベースの作成
	query = "create database " + *dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	// データベースのスイッチ
	query = "use " + *dbName
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	// テーブルが既に作成済みであれば一度 drop する
	query = "drop temporary table if exists `companies`, `safeties`, `lines`, `trains`"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	// テーブルの作成
	query = "create table `companies`("
	query += "id int auto_increment not null primary key, "
	query += "name varchar(50), "
	query += "location varchar(50), "
	query += "url varchar(50)"
	query += ") charset=utf8"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	query = "create table `safeties`("
	query += "id int auto_increment not null primary key, "
	query += "name varchar(50) "
	query += ") charset=utf8"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	query = "create table `lines`("
	query += "id int auto_increment not null primary key, "
	query += "name varchar(50), "
	query += "rail_width int, "
	query += "rail_range int, "
	query += "stations int, "
	query += "safety_id int not null, "
	query += "operator_id int not null, "
	query += "constraint fk_safety_id "
	query += "foreign key (safety_id) "
	query += "references `safeties` (id) "
	query += "on delete restrict on update restrict, "
	query += "constraint fk_operator_id "
	query += "foreign key (operator_id) "
	query += "references `companies` (id) "
	query += "on delete restrict on update restrict"
	query += ") charset=utf8"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	query = "create table `trains`("
	query += "id int auto_increment not null primary key, "
	query += "name varchar(50), "
	query += "line_id int not null, "
	query += "operator_id int not null, "
	query += "max_speed int, "
	query += "acceleration int, "
	query += "safety_id int not null, "
	query += "constraint fk_operator_id "
	query += "foreign key (operator_id) "
	query += "references `companies` (id) "
	query += "on delete restrict on update restrict, "
	query += "constraint fk_line_id "
	query += "foreign key (line_id) "
	query += "references `lines` (id) "
	query += "on delete restrict on update restrict, "
	query += "constraint fk_safety_id "
	query += "foreign key (safety_id) "
	query += "references `safeties` (id) "
	query += "on delete restrict on update restrict"
	query += ") charset=utf8"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}
