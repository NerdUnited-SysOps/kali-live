#!/bin/bash
# set-wallpaper.sh - Apply custom wallpaper to all connected monitors at login.
# Runs as XFCE autostart after xfdesktop starts. Handles any monitor name.
WALLPAPER=/usr/share/backgrounds/kali-16x9/blockchain_1.png

# Wait for xfconf daemon to be ready
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

# Force xfdesktop to reload its backdrop config
xfdesktop --reload 2>/dev/null || pkill -USR1 xfdesktop 2>/dev/null

exit 0
