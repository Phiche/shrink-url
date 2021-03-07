package Storage

import (
	"fmt"
	"time"
)

type User struct {
	id         int
	email      string
	name       string
	creation   time.Time
	last_login time.Time
}

type Url struct {
	Hash            string
	Original_url    string
	Name            string
	Creation_date   time.Time
	expiration_date time.Time
}

func CreateUrl(url Url) {
	fmt.Println(" **** Creating new url ****\n", url)
	if err := Session.Query("INSERT INTO url(hash, original_url, creation_date, expiration_date) VALUES(?, ?, ?, ?)",
		url.Hash, url.Original_url, url.Creation_date, url.expiration_date).Exec(); err != nil {
		fmt.Println("Error while inserting Url")
		fmt.Println(err)
	}
}

func GetOriginalUrl(hash string) string {
	m := map[string]interface{}{}
	var originalUrl = ""
	iter := Session.Query("SELECT original_url from url where hash=?", hash).Iter()
	for iter.MapScan(m) {
		originalUrl = m["original_url"].(string)
		m = map[string]interface{}{}
	}

	return originalUrl
}
