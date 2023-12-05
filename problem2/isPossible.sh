#!/bin/bash

input_file="games.txt"
sum=0
total=0
red=12
g=13
b=14

while read -r line; do
	lineNum=$(echo $line | grep -o -E '[0-9]+' | head -n 1)
	((total+=lineNum))
	rolls=$(echo $line | grep -o -E '[0-9]+_g|[0-9]+_b|[0-9]+_r')
	#echo "rolls for $line:          $rolls"
	for roll in $rolls; do 
		color="${roll: -1}"
		rollNum=$(echo $roll | grep -o -E '[0-9]+' | head -n 1)
		if [ "$rollNum" -gt 14 ] || { [ "$color" == "g" ] && [ "$rollNum" -gt 13 ]; } || { [ "$color" == "r" ] && [ "$rollNum" -gt 12 ]; }; then
    			echo "inside $lineNum roll $roll color $color rollNum $rollNum"
			((sum+=$lineNum))
			break
		fi

	done 
	#echo $lineNum
done < "$input_file"
ans=$(($total-$sum))
echo $ans
