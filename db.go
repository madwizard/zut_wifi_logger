package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)
func initDB() {
	database, _ := sql.Open("sqlite3", "./wifidata.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS wifidata (id INTEGER PRIMARY KEY, essid TEXT, mac TEXT)")
	statement.Exec()
}

func writeDB(data []wifiData) {
	database, _ := sql.Open("sqlite3", "./wifidata.db")
	for _, item := range data {
		statement, _ := database.Prepare("INSERT INTO wifidata (essid, mac) VALUES (?,?)")
		log.Printf("Inserting %s and %s into database", item.ESSID, item.MAC)
		statement.Exec(item.ESSID, item.MAC)
	}
}

func readDB() *[]wifiData {
	database, _ := sql.Open("sqlite3", "./wifidata.db")
	rows, _ := database.Query("SELECT essid, mac from wifidata")
	var data wifiData
	var retdata []wifiData
	{
	}
	for rows.Next() {
		rows.Scan(&data.ESSID, &data.MAC)
		retdata = append(retdata, data)
	}

	return &retdata
}
