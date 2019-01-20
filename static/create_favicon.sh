#!/bin/sh

# based on https://gist.github.com/nateware/900d2d09f4884ac0c073

if [ ! -f input.png ]; then
    echo "You should have an input.png file here"
    exit 1
fi

convert -resize x16 -gravity center -crop 16x16+0+0 -flatten -colors 256 input.png output-16x16.ico
convert -resize x32 -gravity center -crop 32x32+0+0 -flatten -colors 256 input.png output-32x32.ico
convert output-16x16.ico output-32x32.ico favicon.ico
convert -resize x152 input.png apple-touch-icon-152x152.png
convert -resize x120 input.png apple-touch-icon-120x120.png
convert -resize x76  input.png apple-touch-icon-76x76.png
convert -resize x60  input.png apple-touch-icon-60x60.png
