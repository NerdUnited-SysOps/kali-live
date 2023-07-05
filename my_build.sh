#!/usr/bin/zsh

iso_version=$1

sudo rm -vf images/*
sudo rm -rf config/includes.chroot/home/user/go/
~/kali-build/prep-kali-build.sh $iso_version

echo
echo
cat version
echo
echo
echo -n "started "; date +%I:%M%p
./build.sh --clean
./build.sh
echo -n "ended "; date +%I:%M%p
