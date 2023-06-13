#!/usr/bin/zsh

if (( $# < 1 )); then
    echo; echo "Exiting. Expected  volumeX"; echo
    exit 1
elif (( $# > 1 )); then
    echo; echo "Too many arguments.  Only volumeX is expected."; echo
    exit 2
fi


find $1/ -type f -exec sha256sum "{}" + | sudo tee /media/usb/ceremony/$1.laptop.sha

