cd /nfs
btrfs subvol snap reference $1
cd $1
rm etc/ssh/ssh_host_*
ssh-keygen -t dsa -N "" -f etc/ssh/ssh_host_dsa_key
ssh-keygen -t rsa -N "" -f etc/ssh/ssh_host_rsa_key
ssh-keygen -t ecdsa -N "" -f etc/ssh/ssh_host_ecdsa_key
echo "/nfs/$1 *(rw,async,no_subtree_check,no_root_squash)" >> /etc/exports
exportfs -ra
cd /tftpboot/
cp -r reference $1
sed -i "s,:/nfs/client01,:/nfs/$1," $1/cmdline.txt

