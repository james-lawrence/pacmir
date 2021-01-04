## pacmir - torrent based pacman mirror

### benefits
- no need to manage a mirror list.
- no need to figure out fastest/most reliable mirrors (bittorrent more or less does this automatically).
- reduces infrastructure costs for the arch linux community (via bandwidth reductions).
- in theory could support tooling for previous versions of packages

### how it works
pacmir clones and rewrites /etc/pacman.d/mirrorlist prepending itself to the top of the list
thereby becoming the first host pacman will try.

now any requests to download a packages will instead use torrents, falling back to the remaining
mirrors if the torrent cannot be found.

database and signatures requests are upstreamed to the original mirrorlist servers.

### local installation
```bash
systemctl enable --now pacmir.service
```

### hosted mirror installation
#### requirements
- rsync
- systemd

#### configuration
- see https://archlinux.org/mirrors/status/#successful to pick a mirror to sync from.
- create /etc/pacmir/defaults
```environment
REPOSITORY=rsync://example.com/mirror/archlinux/
CACHE_DIRECTORY=/var/cache/pacmir
```

```bash
systemctl enable --now pacmir-mirror-rsync@pool.service
systemctl enable --now pacmir-mirror-rsync@core.service
systemctl enable --now pacmir-mirror-rsync@community.service
systemctl enable --now pacmir-mirror-rsync@extra.service
systemctl enable --now pacmir-mirror-rsync@multilib.service
systemctl enable --now pacmir-mirror.service
```