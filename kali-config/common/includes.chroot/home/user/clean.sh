#!/usr/bin/zsh

rm -rfv ~/blockfabric-ceremony/
rm -rfv ~/.aws/
mv -v ~/.config/xfce4/panel/launcher-6/blockexplorer.desktop ~/.config/xfce4/panel/launcher-6/blockexplorer.desktop_
rm -v ~/bootstrap.log
cp ~/.ssh/config.template ~/.ssh/config
sudo chown -R user:user ~/.* ~/*
echo
echo you should now exit
echo
