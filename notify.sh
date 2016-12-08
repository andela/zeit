#! /bin/bash
seconds=0
divider=$1

function getHours {
	seconds=$1
	hour=$((seconds / 60))
	if [ $hour -gt 0 ] && [ $hour -ne 1 ]; then
		echo "${hour} hours";
		return 0
	fi
	echo "${hour} hour"
}

function getMinutes {
	seconds=$1
	minutes=$((seconds % 60))
	if [ $minutes -gt 0 ] && [ $minutes -ne 1 ]; then
		echo "${minutes} minutes";
		return 0
	fi
	echo "${minutes} minute"
}

while [[ $seconds -ne 1440 ]]; do # Max of 24 hours before timer quits
	((seconds++))
	quotient=$((seconds % divider))
	hours=$(getHours ${seconds})
	minutes=$(getMinutes ${seconds})
	sleep 60
	if [ $quotient -eq 0 ];
		then
			osascript -e "display notification \"You have Logged about ${hours} and ${minutes}\" with title \"Zeit Info\" subtitle \" Time Logged\" sound name \"Blow\""
	fi
done;
