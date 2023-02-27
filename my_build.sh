sudo rm -vf images/*
echo
echo
cat version
echo
echo
echo -n "started "; date +%I:%M%p
./build.sh
echo -n "ended "; date +%I:%M%p

