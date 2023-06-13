sudo rm -vf images/*
sudo rm -rf config/includes.chroot/home/user/go/
echo
echo
cat version
echo
echo
echo -n "started "; date +%I:%M%p
./build.sh $@
echo -n "ended "; date +%I:%M%p

