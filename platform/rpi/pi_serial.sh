serial=$(cat /proc/cpuinfo | grep "Serial" | awk '{print $3}' | sed 's/^0*//')
echo "My serial is:"$serial
ssh root@192.168.7.48 "bash add_node.sh $serial"
if [ $? -eq 0 ]
then
	echo "Node registration succeeded"
	poweroff
else
	echo "Node registration failure"
fi
