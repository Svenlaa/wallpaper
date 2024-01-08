package main

import (
	"database/sql"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", DATABASE_URL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	currentWallpaper := getWallpaper()

	var url string;
	attempts := 0
	for {
		db.QueryRow("SELECT url FROM wallpapers ORDER BY RAND()").Scan(&url)
		if url != currentWallpaper || attempts >= 3 {
			break
		}
	}
	setWallpaper(url)
}

func getWallpaper () string {
	schema := "org.gnome.desktop.background"
	key := "picture-uri"

	res, err := exec.Command("gsettings", "get", schema, key).Output()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(strings.ReplaceAll(string(res), "'", ""));
}

func setWallpaper(url string) {
	schema := "org.gnome.desktop.background"
	key := "picture-uri"

	cmd := exec.Command("gsettings", "set", schema, key, url)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}