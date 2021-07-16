# How to Fix Broken Grub (EFI)

## Tested on 

- Arch
- Manjaro

## Problem

- Grub won't show up/won't boot

## Reason

- Grub was damaged/wiped for a number of reasons

## Fix

**If you find yourself with an overwritten GRUB, follow the below links and commands**

## SystemRescue

- Create a system rescue live boot USB and boot into it and select the first option (default/normal presets)
- Once in, make sure you mount your broken system to **``/mnt/linux``**! This is very important as the system will **FREEZE** if you mount directly to ``/mnt``
- After mounting your main drive (presumably ``/dev/sda2`` (yours may be different)), mount your boot drive to ``/mnt/linux/boot/efi`` (if it does not exist, create the folder with ``mkdir``)
- After mounting your boot and main drive, mount your home drive if you have one (if you do it would go on ``/mnt/linux/home``)
- FOLLOW THIS LINK INSTRUCTIONS: [IT GOES VERY IN-DEPTH](https://www.system-rescue.org/disk-partitioning/Repairing-a-damaged-Grub/)
    - Make sure you are following the second example where they ``chroot`` into the broken system.

### ***BEFORE YOU RUN GRUB-INSTALL***

- Make sure you run ``mount -t efivarfs efivarfs /sys/firmware/efi/efivars``
- Make sure you set the ``efi-directory`` to ``/boot/efi``

## NOW RUN GRUB INSTALL

**It should be a success, congratulations, your grub is repaired.**

- ~5late