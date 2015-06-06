#!/bin/bash


LIST=`find . -type f -name "*.svg"`

for x in $LIST
do
	NAME=`echo $x | sed -e 's/.svg//'`
	gzip -k $NAME.svg
	mv $NAME.svg.gz $NAME.svgz
	unset NAME
done
