# How to fix FlameShot

## Tested on

- Arch

## Problem

- Flameshot will not launch/errors due to wayland/kde/plasma

## Fix

## Works on ARCH-12 KDE

### Do **NOT** install flameshot with ``pacman``

- Run ``pacman -R flameshot`` if you have it installed already
- Run ``pamac build flameshot-git`` (pulls from official github repo)
- Run ``flameshot gui`` to get the screenshot menu you may be familiar with
- Create a system shortcut to bind ``PRNTSCRN`` to ``/usr/bin/flameshot gui``

**AT THIS POINT IT SHOULD WORK**

- ~5late