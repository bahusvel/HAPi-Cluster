cd /nfs
find . -maxdepth 1 ! \( -name 'reference' -o -name 'share' -o -path . \) -exec btrfs subvol delete {}  +

cd /tftpboot
find . -maxdepth 1 ! \( -name 'reference' -o -name 'original' -o -name 'bootcode.bin' -o -path . \) -exec rm -rf {}  +

echo '/nfs/share *(rw,sync,no_subtree_check,no_root_squash)' > /etc/exports
exportfs -ra
