while read -u 10 host; do
	echo $host
	ssh root@$host "$1"
done 10<hosts.txt
