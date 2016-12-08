cd /nfs
if [ ! -d $1 ]
then
        echo "Node does not exist"
        exit 1
fi
btrfs subvol delete $1
cd /tftpboot
rm -rf $1
