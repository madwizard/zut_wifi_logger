package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

// initDB creates DB if it doesn't exit
func initDB() {
	database, _ := sql.Open("sqlite3", "./wifidata.db")
	wifidata, _ := database.Prepare("CREATE TABLE IF NOT EXISTS wifidata (id INTEGER PRIMARY KEY, essid TEXT, mac TEXT, freq TEXT, siglvl TEXT, " +
		"qual TEXT, enc TEXT, channel INT, mode TEXT, ieee TEXT, bitrates TEXT, wpa text, tmstmp TEXT, position TEXT)")
	wifidata.Exec()
	gpsdata, _ := database.Prepare("CREATE TABLE IF NOT EXISTS gpsdata (id INTEGER PRIMARY KEY, gpsdata TEXT DEFAULT '', tmstmp TEXT DEFAULT '')")
	gpsdata.Exec()
}

// chekcIfIsInDB must check if the ESSID and MAC pair exists in DB, will discard it for now
func checkIfIsInDB(ESSID string, MAC string) bool {
	database, _ := sql.Open("sqlite3", "./wifidata.db")
	defer database.Close()
	var essid string
	var mac string

	rows, _ := database.Query("SELECT essid, mac FROM wifidata WHERE essid = ? AND mac = ?", ESSID, MAC)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&essid, &mac)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Printf("checkIfIsInDB: Couldn't scan rows")
			} else {
				return false
			}
		}

	}
	return false
}

// writeWiFiDB writes all data from scan to DB
// Checks if ESSID + MAC pair already is in DB, then skips write.
func writeWiFiDB(data []wifiData, timestamp int64) {
	database, _ := sql.Open("sqlite3", "./wifidata.db")
	defer database.Close()

	for _, item := range data {
		if item.MAC != "" {
			exists := checkIfIsInDB(item.ESSID, item.MAC)
			if exists {
				continue
			} else {
				statement, _ := database.Prepare("INSERT INTO wifidata (essid, mac, freq, siglvl, qual, enc, channel, mode, ieee, bitrates, wpa, tmstmp, position) " +
					"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)")
				tm := strconv.FormatInt(timestamp, 10)
				statement.Exec(item.ESSID, item.MAC, item.Freq, item.SigLvl, item.Qual, item.Enc,
					item.Channel, item.Mode, item.IEEE, item.Bitrates, item.WPA, tm, GpsData.Data)
			}
		}
	}
}

func readDB() *[]webdata {
	database, _ := sql.Open("sqlite3", "./wifidata.db")
	defer database.Close()
	rows, _ := database.Query("SELECT tmstmp, essid, mac, freq, siglvl, qual, enc, channel, mode, ieee, bitrates, wpa, " +
		" position wifidata")
	var item webdata
	var retdata []webdata
	{
	}
	for rows.Next() {
		rows.Scan(&item.Timestamp, &item.ESSID, &item.MAC, &item.Freq, &item.SigLvl, &item.Qual, &item.Enc,
			&item.Channel, &item.Mode, &item.IEEE, &item.Bitrates, &item.WPA, &item.GPS)

		retdata = append(retdata, item)
	}

	return &retdata
}
