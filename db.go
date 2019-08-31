package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"strings"
	"unicode"
)

// initDB creates DB if it doesn't exit
func initDB() {
	database, err := sql.Open("sqlite3", "./wifidata.db")
	if err != nil {
		log.Printf("Error initilizing Database: %v", err)
	}
	wifidata, err := database.Prepare("CREATE TABLE IF NOT EXISTS wifidata (id INTEGER PRIMARY KEY, essid TEXT, mac TEXT, freq TEXT, siglvl TEXT, " +
		"qual TEXT, enc TEXT, channel INT, mode TEXT, ieee TEXT, bitrates TEXT, wpa TEXT, tmstmp TEXT, latitude TEXT, longitude TEXT, read BOOL, UNIQUE(essid, mac))")
	if err != nil {
		log.Printf("Error creating table: %v", err)
	}
	wifidata.Exec()
}

// writeWiFiDB writes all data from scan to DB
func writeWiFiDB(data []wifiData, timestamp int64) {
	database, err := sql.Open("sqlite3", "./wifidata.db")
	if err != nil {
		log.Printf("Can't open database: %v", err)
	}
	defer database.Close()

	for _, item := range data {
		if item.MAC != "" {
				statement, _ := database.Prepare("INSERT OR IGNORE INTO wifidata (essid, mac, freq, siglvl, qual, enc, channel, mode, ieee, bitrates, wpa, tmstmp, latitude, longitude, read) " +
					"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
				tm := strconv.FormatInt(timestamp, 10)
				statement.Exec(stripSpaces(item.ESSID), stripSpaces(item.MAC), item.Freq, item.SigLvl, item.Qual, item.Enc, item.Channel, item.Mode, item.IEEE, item.Bitrates, item.WPA, tm, GPSdata.Latitude, GPSdata.Longitute, GPSdata.GPSRead)
		}
	}
}

func readDB() *[]webdata {
	database, err := sql.Open("sqlite3", "./wifidata.db")
	if err != nil {
		log.Printf("Can't open database: %v", err)
	}
	defer database.Close()
	rows, _ := database.Query("SELECT tmstmp, essid, mac, freq, siglvl, qual, enc, channel, mode, ieee, bitrates, wpa, " +
		" latitude, longitude, read FROM wifidata")
	var item webdata
	var retdata []webdata
	{
	}
	for rows.Next() {
		rows.Scan(&item.Timestamp, &item.ESSID, &item.MAC, &item.Freq, &item.SigLvl, &item.Qual, &item.Enc,
			&item.Channel, &item.Mode, &item.IEEE, &item.Bitrates, &item.WPA, &item.Latitude, &item.Longitude, &item.GPSRead)

		retdata = append(retdata, item)
	}

	return &retdata
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}