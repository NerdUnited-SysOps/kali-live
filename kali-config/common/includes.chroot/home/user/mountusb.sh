#!/usr/bin/zsh

usb_device=""

echo;echo
echo "  === BEFORE ==="
lsblk -p
echo
echo "Enter the usb drive and partition (\"sd\" already provided) : "
echo -n "example: /dev/sdb1 is showing above.  You should enter only \"b1\" : "
read usb_device
sudo mount /dev/sd$usb_device /media/usb

echo;echo
echo "  === AFTER ==="
lsblk -p
echo
