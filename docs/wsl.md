# Membaca file linux dengan wsl windows

## jalankan ps1 sebagai administrator
```ps1
# list disk
Get-Disk

# Number Friendly Name Serial Number                    HealthStatus         OperationalStatus      Total Size Partition
#                                                                                                              Style
# ------ ------------- -------------                    ------------         -----------------      ---------- ----------
# 1      Lexar E6      AA202412120563                   Healthy              Online                  238.47 GB GPT

# List partisi
Get-Partition -DiskNumber 1


# PartitionNumber  DriveLetter Offset                                        Size Type
# ---------------  ----------- ------                                        ---- ----
# 1                           1048576                                    1.17 GB System
# 2                           1256194048                                 1.95 GB Unknown
# 3                           3353346048                                15.62 GB Unknown
# 4                           20130562048                              219.73 GB Unknown

# mounting disk dari partisi
wsl --mount \\.\PHYSICALDRIVE1 --partition 4

# masuk ke wsl
wsl
```
## Di dalam wls
```bash
cd /mnt/host/wsl/PHYSICALDRIVE1p4/

# unmount via wsl
umount /mnt/host/wsl/PHYSICALDRIVE1p4/
```

## Unmount via ps1
```ps1
wsl --unmount \\.\PHYSICALDRIVE1
```