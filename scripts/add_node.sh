cd /nfs
if [ -d $1 ]
then
	echo "Already registered"
	exit 1
fi
btrfs subvol snap reference $1
cd $1
rm etc/ssh/ssh_host_*
ssh-keygen -t dsa -N "" -f etc/ssh/ssh_host_dsa_key
ssh-keygen -t rsa -N "" -f etc/ssh/ssh_host_rsa_key
ssh-keygen -t ecdsa -N "" -f etc/ssh/ssh_host_ecdsa_key
echo "$1" > etc/hostname
echo "127.0.1.1 $1" >> etc/hosts
echo "/nfs/$1 *(rw,async,no_subtree_check,no_root_squash)" >> /etc/exports
exportfs -ra
