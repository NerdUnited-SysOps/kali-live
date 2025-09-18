# Writing the iso to a flash drive

## Mac instructions
```sh
diskutil list
diskutil unmountdisk /dev/diskX
sudo dd if=ceremony_v1.1.x.iso of=/dev/rdiskX bs=4M status=progress
diskutil list
diskutil eject /dev/diskX
```
