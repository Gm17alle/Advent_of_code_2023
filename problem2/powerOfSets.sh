#!/bin/bash

input_file="games.txt"
sum=0

while read -r line; do
	rolls=$(echo $line | grep -o -E '[0-9]+_g|[0-9]+_b|[0-9]+_r')
	maxR=1
	maxB=1
	maxG=1	
	for roll in $rolls; do
		color="${roll: -1}"
		 rollNum=$(echo $roll | grep -o -E '[0-9]+' | head -n 1)
		
		 if [ "$color" == "g" ] && [ "$rollNum" -gt "$maxG" ]; then
			 maxG=$rollNum
		 fi

		 if [ "$color" == "r" ] && [ "$rollNum" -gt "$maxR" ]; then
			 maxR=$rollNum
		 fi

		 if [ "$color" == "b" ] && [ "$rollNum" -gt "$maxB" ]; then
			 maxB=$rollNum
		 fi
	done
	 ((sum+=$maxG*$maxB*$maxR))
done < "$input_file"
ans=$sum
echo $ans
