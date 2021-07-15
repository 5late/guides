# Setup MatterMost

## Tested on

- Arch (Server)

## End Goal

- Have a functional MatterMost server which can be connected to from the browser

## Notice: I will be following [this](https://wiki.archlinux.org/title/Mattermost) tutorial for most of this set up.

## Start

- Install the mattermost package (``sudo pacman -S mattermost``)
- [Setup the Postgres database](https://wiki.archlinux.org/title/PostgreSQL#Installation)
    - If you run into a permission issue for ``/var/lib/postgres/data``, use [this guide](./mattermost/FIX-POSTGRES-PERMISSIONS.md)
- I will be configuring MatterMost with Unix Socket
- Make sure you setup the ``/etc/webapps/mattermost/config.json`` file to have SITE URL set to whatever domain you are using
- Set your port to the one you have configured to expose on your server
- Visit the site url in your browser, followed by the port you set. (Ex: ``localhost:8065``)
- If you cannot change your profile picture, or cannot send images, read [this guide](./mattermost/FIX-MATTERMOST-PERMS.md)

## You should now have MatterMost installed

- ~5late