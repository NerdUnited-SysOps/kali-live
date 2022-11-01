# Kali Live Build

Create ceremony os build iso



build.sh --verbose

image will be created in images/*.iso

## mac instructions
diskutil list

diskutil unmountdisk /dev/diskX

sudo dd if=kali-linux-rolling-live-amd64.iso of=/dev/rdiskX bs=4M status=progress
