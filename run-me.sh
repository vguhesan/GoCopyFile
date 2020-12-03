#!/usr/bin/env bash
echo "Removing ./target-file.txt - if it exists."
rm ./target-file.txt
echo "Removing previous build - if it exists."
rm ./CopyFile
echo "Building GoLand executable."
go build && chmod 775 CopyFile
echo "Running executable."
./CopyFile 
echo "Examine the contents of target-file.txt"
echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
cat ./target-file.txt
echo "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
echo ""
echo "Cheers from Venkatt Guhesan at MyThinkPond.com"
echo "Visit MyThinkPond.com and consider subscribing to get notified when we post new articles. Thank You!"
echo "~~~~-> https://MyThinkPond.com/ "