#!/bin/bash

input_file="trebuchet.txt"
sum=0
count=0

while read -r line; do
	left=$(echo "$line" | grep -o -E '[0-9]|one|two|three|four|five|six|seven|eight|nine' | sed -n '1p')
	right=$(echo "$line" | rev | grep -o -E '[0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin' | sed -n '1p' | rev)
	num="$left$right"
	before=$num
	num=$(echo "$num" | sed 's/one/1/g')
	num=$(echo "$num" | sed 's/two/2/g')
	num=$(echo "$num" | sed 's/three/3/g')
	num=$(echo "$num" | sed 's/four/4/g')
	num=$(echo "$num" | sed 's/five/5/g')
	num=$(echo "$num" | sed 's/six/6/g')
	num=$(echo "$num" | sed 's/seven/7/g')
	num=$(echo "$num" | sed 's/eight/8/g')
	num=$(echo "$num" | sed 's/nine/9/g')
	#echo "the goods: $num"
	((count++))
	#echo "linenum: $count line: $line before: $before , after: $num"
	((sum+=num))
done < "$input_file"
echo "The sum is: '$sum'"
