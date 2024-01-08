# Wallpaper changer for GNOME

## About

I'm exploring GO and have recently installed Fedora. I like to change my wallpaper regularly. I decided to combine both and build a little CLI that picks a random wallpaper.\
It fetches an image URL from the database, and uses `gsettings` to set the wallpaper.

## Setup

#### Clone REPO
``` bash
git clone https://github.com/Svenlaa/wallpaper
cd wallpaper
cp constants.example constants.go
```

#### ENV and Database

Set your own `DATABASE_URL` in `constants.go` \
Run the sql query in the `migrations/` folder

#### Add your own wallpapers
Execute SQL commands \
`INSERT INTO wallpapers (url, comment) VALUES (?, ?);`

#### Build and execute

``` bash
go build
./wallpaper
```
