package main

import (
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	dir := getWallpaperDirectory()
	currentWallpaper := getCurrentWallpaper()

	var url string;
	attempts := 0
	for {
		attempts++;
		url = "file://" + pickFile(dir)
		if url != currentWallpaper || attempts >= 3 {
			break
		}
	}
	setWallpaper(url)
}

func pickFile(dir string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		panic("No files found.")
	}
	return filepath.Join(dir, files[rand.Intn(len(files))].Name())
}

func getWallpaperDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	wallpaperDir := filepath.Join(homeDir, "Pictures", "Wallpapers")
	wallpaperDir, err = filepath.Abs(wallpaperDir)
	if err != nil {
		panic(err)
	}

	return wallpaperDir
}

func getCurrentWallpaper () string {
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