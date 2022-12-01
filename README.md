# Kali Live Build

# Prepare environment and install dependencies
sudo apt update
sudo apt install -y git live-build simple-cdd cdebootstrap curl
git clone https://github.com/caseynielson/kali-live.git

Create ceremony os build iso



build.sh --verbose

image will be created in images/*.iso

## mac instructions
diskutil list

diskutil unmountdisk /dev/diskX

sudo dd if=kali-linux-rolling-live-amd64.iso of=/dev/rdiskX bs=4M status=progress
