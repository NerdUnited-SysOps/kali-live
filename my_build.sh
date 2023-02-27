#!/usr/bin/zsh
sudo rm -vf images/*
echo
echo
cat version
echo
echo
echo -n "started "; date +%I:%M%p
./build.sh $1 $2
echo -n "ended "; date +%I:%M%p
