while read -u 10 host; do
	echo $host
	ssh root@$host "killall cpd; /mnt/share/bin/cpd -c 192.168.7.28:3334 -i $host" &
done 10<hosts.txt

