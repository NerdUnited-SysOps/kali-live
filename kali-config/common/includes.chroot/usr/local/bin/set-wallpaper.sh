#!/bin/bash
# set-wallpaper.sh - Apply custom wallpaper to all connected monitors at login.
# Handles any monitor name that xrandr reports, including hardware not covered
# by the static xfce4-desktop.xml entries.
WALLPAPER=/usr/share/backgrounds/kali-16x9/default

# Wait for xfconf daemon to be ready (it starts with the session but may lag)
for i in $(seq 1 10); do
    xfconf-query -c xfce4-desktop -l >/dev/null 2>&1 && break
    sleep 1
done

# Get connected monitor names from xrandr
monitors=$(xrandr --listmonitors 2>/dev/null | tail -n +2 | awk "{print \$NF}")

if [ -z "$monitors" ]; then
    exit 0
fi

for monitor in $monitors; do
    base="/backdrop/screen0/monitor${monitor}/workspace0"
    xfconf-query -c xfce4-desktop -p "${base}/image-path"  -s "$WALLPAPER" --create -t string
    xfconf-query -c xfce4-desktop -p "${base}/image-style" -s 5           --create -t int
    xfconf-query -c xfce4-desktop -p "${base}/image-show"  -s true        --create -t bool
done
