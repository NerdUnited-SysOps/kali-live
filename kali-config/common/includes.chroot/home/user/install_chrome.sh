#!/usr/bin/zsh

## Chrome may be needed for compatibility reasons: 
## https://www.geekersdigest.com/how-to-install-google-chrome-from-apt-repository/

wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | sudo apt-key add -
echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" | sudo tee /etc/apt/sources.list.d/chrome.list
sudo apt-get update
## sudo apt-cache search chrome
sudo apt-get install google-chrome-stable --yes

