# Wallpaper changer for GNOME

## About

I'm exploring GO and have recently installed Fedora. I like to change my wallpaper regularly. I decided to combine both and build a little CLI that picks a random wallpaper.\
It fetches an image URL from the database, and uses `gsettings` to set the wallpaper.

## Setup

#### Clone and build
``` bash
git clone https://github.com/Svenlaa/wallpaper
cd wallpaper
go build
```

#### Add your own wallpapers
Add images to the `~/Pictures/Wallpapers` folder. \
You might need to create the folder yourself.
