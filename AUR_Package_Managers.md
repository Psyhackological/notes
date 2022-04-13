# To install these:

`sudo pacman -S paru`

`sudo pacman -S yay`

# [paru](https://github.com/morganamilo/paru)
Paru is your standard pacman wrapping AUR helper with lots of features and minimal interaction.

[![asciicast](https://asciinema.org/a/sEh1ZpZZUgXUsgqKxuDdhpdEE.svg)](https://asciinema.org/a/sEh1ZpZZUgXUsgqKxuDdhpdEE)

Synchronize and update all packages

`paru`

`paru -Sua`

Interactively search for and install a package

`paru package_name_or_seach_term`

Install a new package from the repos and AUR

`paru -S package_name`

To remove a single package, leaving all of its dependencies installed

`paru -R package_name`

To remove a package and its dependencies which are not required by any other installed package

`paru -Rs package_name`

To remove a package, its dependencies and all the packages that depend on the target package

`paru -Rsc package_name`

Get information about a package

`paru -Si package_name`

# [yay](https://github.com/Jguer/yay)
Yet Another Yogurt - An AUR Helper Written in Go
[![asciicast](https://asciinema.org/a/399431.svg)](https://asciinema.org/a/399431)
[![asciicast](https://asciinema.org/a/399433.svg)](https://asciinema.org/a/399433)

Synchronize and update all packages
`yay`

`yay -Sua`

Interactively search for and install a package

`yay package_name_or_seach_term`

Install a new package from the repos and AUR

`yay -S package_name`

To remove a single package, leaving all of its dependencies installed

`yay -R package_name`

To remove a package and its dependencies which are not required by any other installed package

`yay -Rs package_name`

To remove a package, its dependencies and all the packages that depend on the target package

`yay -Rsc package_name`

Get information about a package

`yay -Si package_name`

# [pacman](https://wiki.archlinux.org/title/Pacman)
Synchronize and update all packages

`sudo pacman -Syu`

Search the package database for a regular expression or keyword

`sudo pacman -Ss search_pattern`

Install a new package from the repos and AUR

`sudo pacman -S package_name`

To remove a single package, leaving all of its dependencies installed

`sudo pacman -R package_name`

To remove a package and its dependencies which are not required by any other installed package

`sudo pacman -Rs package_name`

To remove a package, its dependencies and all the packages that depend on the target package

`sudo pacman -Rsc package_name`

Get information about a package

`pacman -Si package_name`

# Comparison
Synchronize and update all packages

`paru`

`paru -Sua`

`yay`

`yay -Sua`

`sudo pacman -Syu`

Interactively search for and install a package

`paru package_name_or_seach_term`

`yay package_name_or_seach_term`

`sudo pacman -Ss search_pattern`

Install a new package from the repos and AUR

`paru -S package_name`

`yay -S package_name`

`sudo pacman -S package_name`

To remove a single package, leaving all of its dependencies installed

`paru -R package_name`
`yay -R package_name`

`sudo pacman -R package_name`


To remove a package and its dependencies which are not required by any other installed package

`paru -Rs package_name`

`yay -Rs package_name`

`sudo pacman -Rs package_name`

To remove a package, its dependencies and all the packages that depend on the target package

`paru -Rsc package_name`

`yay -Rsc package_name`

`sudo pacman -Rsc package_name`

Get information about a package

`paru -Si package_name`

`yay -Si package_name`

`pacman -Si package_name`

# Conclusion
Flags of the `paru` and `yay` are the same.

# Differences between paru/yay and pacman

| Action | paru | yay | pacman |
| ---- | ---- | --- | ------ |
| Updating system | `paru` | `yay` | `pacman -Syu` |
| Searching for a package | `paru package_name` | `yay package_name` | `pacman -Ss search_term |

## + despite `pacman -Si` everything requires sudo privileges while `paru` and `yay` doesn't.
