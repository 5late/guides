# Fix Broken Kernel Boot Perms

## Tested on

- Arch

## Problem

- You cannot boot into your system due to a broken kernel image

## Reason

- The drive had a small error which broke/overwrote the kernel

## Fix

- Boot into a live usb
- Mount, then chroot into your system
- Run ``pacman -Syu linux``
- Note the kernel version
- Run ``mkinitcpio -P``

#### At this point it may work, test it. If it does not, continue on:

- Run ``grub-install --target=x86_64-efi --efi-directory=/boot/efi --bootloader-id=GRUB``
- Run ``grub-mkconfig -o /boot/grub/grub.cfg``

## It should now work, and you should be able to boot properly

- ~5late