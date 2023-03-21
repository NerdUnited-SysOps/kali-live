#!/usr/bin/zsh

# Set aliases
alias ll='ls -la --color=auto'
alias lsblk='lsblk -p'
alias cp='sudo cp "$@"'
alias mountusb='~/mountusb.sh'
alias umountusb='sudo umount /media/usb; lsblk; echo'
alias sha='~/sha.sh $1'
alias chownuser='sudo chown -R user:user ~/.* ~/*'
alias version='cat ~/version'
export genesis=genesis.blockfabric.net

# set a simple prompt
PROMPT=$'%B%F{green}%n@%m%b%F{reset}:%B%F{blue}%~%b%F{reset}%(#.#.$) '

chownuser


if [ "$TERM" = "linux" ]; then
  ## "linux" TERM means it's the first silent autologin to desktop, not a Terminal shell launch
  ## run mkdir and chown commands only once at first bootup, not each time a new shell is launched
  # make a dir for usb mount point
  sudo mkdir /media/usb > /dev/null 2>&1
  # do not modify this setting.  Doing so is counter to the intention and purpose of this ceremony laptop's configuration, which is to boot and remain "air gapped"
  sudo nmcli radio wifi off

else

## These commands should run with every new Terminal shell
# Setup Firefox panel shortcut for block explorer
  mv -v ~/.config/xfce4/panel/launcher-6/blockexplorer.desktop_ ~/.config/xfce4/panel/launcher-6/blockexplorer.desktop > /dev/null 2>&1

  timedatectl set-timezone America/Denver
  timedatectl set-ntp true
  timedatectl status
  echo;echo
  cat ~/version
  echo
  echo "!== DOUBLE CHECK THAT THE TIME IS SYNCED. HAVE FUN !=="
  echo
  echo -n "OS DATE & TIME: "; date '+%A,%B %d, %Y    %I:%M:%S %p'
  echo
fi
