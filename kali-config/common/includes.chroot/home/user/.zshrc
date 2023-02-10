#!/usr/bin/zsh

# Set aliases
alias ll='ls -la --color=auto'
alias count='f(){ find "$1" -type f | wc -l; unset -f f; }; f'
alias cp2usb='f(){ sudo cp -vr "$1" /media/usb/ceremony/; unset -f f; }; f'
alias findf='f(){ find "$1" -type f; unset -f f; }; f'
alias mountusb='sudo fdisk -l; sudo mount /dev/sdb1 /media/usb; sudo mkdir /media/usb/ceremony > /dev/null 2>&1; df -h; echo'
alias umountusb='sudo umount /media/usb; df -h; echo'
alias chownuser='sudo chown -R user:user ~/.* ~/*'
alias version='cat ~/version'
export genesis=genesis.blockfabric.net

# set a simple prompt
PROMPT=$'%B%F{green}%n@%m%b%F{reset}:%B%F{blue}%~%b%F{reset}%(#.#.$) '

chownuser

# do not modify this setting.  Doing so is counter to the intention and purpose of this ceremony laptop's configuration, which is to boot and remain "air gapped"
nmcli radio wifi off

if [ -f system_started ]; then
  mv system_started user_started
# Setup panel shortcut
  mv ~/.config/xfce4/panel/launcher-6/blockexplorer.desktop_ ~/.config/xfce4/panel/launcher-6/blockexplorer.desktop > /dev/null 2>&1

# make a dir for usb mount point
  sudo mkdir /media/usb > /dev/null 2>&1
# time sync and show ISO version, only the first time this boots up
  timedatectl set-timezone America/Denver
  timedatectl set-ntp true
# sleep 1 second to allow for time sync
  sleep 1
  timedatectl status
  echo
  echo "!== CEREMONY LAPTOP INITIALIZATION IS COMPLETE !=="
  echo
else
  touch system_started
fi

cat ~/version
echo
echo "!== DOUBLE CHECK THAT THE TIME IS SYNCED. HAVE FUN !=="
echo
echo -n "OS DATE & TIME: "; date '+%A,%B %d, %Y    %I:%M:%S %p'
echo

