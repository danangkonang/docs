# adb
```bash
adb devices

adb shell

# list all apk
pm list packages -s

# uninstall apk
pm uninstall -k --user 0 <package name>

```
```bash
# mode fastboot
sudo adb reboot-bootloader
```
# fastboot
```bash
fastboot --help

sudo fastboot devices
```

## Flash android
```bash
sudo adb reboot-bootloader
sudo sh flash_all.sh
```
## Root
```sh
download firmware
extract and find file boot.img
open magisk and install boot.img
sudo adb reboot-bootloader
fastboot flash boot.img
fastboot reboot
```