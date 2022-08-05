package configuration

import (
	"fmt"

	ibmdb "github.com/ibmdb/go_ibm_db"
)

func DbConnect(db2uri db2conn) *ibmdb.DBP {
	pool := ibmdb.Pconnect("PoolSize=20")
	conn := fmt.Sprintf("HOSTNAME=%s;DATABASE=%s;port=%d;PROTOCOL=TCPIP;UID=%s;PWD=%s;AUTHENTICATION=SERVER", db2uri.Host, db2uri.Name, db2uri.Port, db2uri.User, db2uri.Password)
	db := pool.Open(conn, "SetConnMaxLifetime=30")
	return db

}

type db2conn struct {
	Name      string
	Host      string
	Port      int
	User      string
	Password  string
	TableName string
}

func DbSelector(id string) db2conn {

	dbs := map[string]db2conn{
		"2016": {
			Name:      "DBFE2017",
			Host:      "172.19.46.21",
			Port:      53000,
			User:      "ubindownload",
			Password:  "K1ll1nGbyTh3N4w3",
			TableName: "TM_CE_DOCUMENTO",
		},
		"2017": {
			Name:      "DBFE2017",
			Host:      "172.19.46.21",
			Port:      53000,
			User:      "ubindownload",
			Password:  "K1ll1nGbyTh3N4w3",
			TableName: "TM_CE_DOCUMENTO",
		},
		"2018": {
			Name:      "dbfe",
			Host:      "172.19.46.21",
			Port:      53000,
			User:      "ubindownload",
			Password:  "K1ll1nGbyTh3N4w3",
			TableName: "TM_CE_DOCUMENTO",
		},
		"2019": {
			Name:      "DBFEPR",
			Host:      "172.19.46.21",
			Port:      53000,
			User:      "ubindownload",
			Password:  "K1ll1nGbyTh3N4w3",
			TableName: "TM_CE_DOCUMENTO",
		},
		"2020": {
			Name:      "DBFEPR",
			Host:      "172.19.46.21",
			Port:      53000,
			User:      "ubindownload",
			Password:  "K1ll1nGbyTh3N4w3",
			TableName: "TM_CE_DOCUMENTO",
		},
		"2021": {
			Name:      "DBFEPR",
			Host:      "172.19.46.21",
			Port:      53000,
			User:      "ubindownload",
			Password:  "K1ll1nGbyTh3N4w3",
			TableName: "TM_CE_DOCUMENTO_2021",
		},
		"2030": {
			Name:      "DBFE",
			Host:      "172.19.35.121",
			Port:      52000,
			User:      "ubindownload",
			Password:  "K1ll1nGbyTh3N4w3",
			TableName: "TM_CE_DOCUMENTO",
		},
	}
	return dbs[id]

}
