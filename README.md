# Kali-Live Build-Scripts

_`live-build` configuration for Kali ISO images._

These are the same [build-scripts](https://gitlab.com/kalilinux/build-scripts) that the [Kali team](https://www.kali.org/) uses to generate the official Kali Linux base images, found here: [kali.org/get-kali/](https://www.kali.org/get-kali/).

_Build your Kali Linux image today!_

- - -

These images can be used to live boot into Kali, from such a USB/CD/DVD/sdCard, as well offers a basic installation. For more customization during setup, see [kali-installer](https://gitlab.com/kalilinux/build-scripts/kali-installer).

- [kali-installer](https://gitlab.com/kalilinux/build-scripts/kali-installer) uses [Simple-CDD](https://wiki.debian.org/Simple-CDD) _(which is a wrapper for [debian-cd](https://wiki.debian.org/debian-cd))_
- [kali-live](https://gitlab.com/kalilinux/build-scripts/kali-live) uses [live-build](https://live-team.pages.debian.net/live-manual/html/live-manual/index.en.html)

- - -

Have a look at [Live Build a Custom Kali ISO](https://www.kali.org/docs/development/live-build-a-custom-kali-iso/) for explanations on how to use this repository.

There are also other [code examples of live-build](https://gitlab.com/kalilinux/recipes/live-build-config-examples), as well as [code examples for pre-seed to automate/unattended installation](https://gitlab.com/kalilinux/recipes/kali-preseed-examples).

- - -

## Help

```console
$ ./build.sh --help
Usage: ./build.sh [<option>...]

  --distribution <arg>
  --proposed-updates
  --arch <arg>
  --verbose
  --debug
  --variant <arg>
  --version <arg>
  --subdir <arg>
  --get-image-path
  --no-clean
  --clean
  --help

More information: https://www.kali.org/docs/development/live-build-a-custom-kali-iso/
$
```

## Using this repo to create a new ceremony.iso
Boot kali-linux-2025.4-qemu-amd64.qcow2 (process has been tested with kali linux 2024 and 2025.4)
more memory and vcpus will make the build faster


```console
git clone https://github.com/NerdUnited-SysOps/kali-live.git
cd kali-live
./build.sh --verbose
```
images are built in kali-live/images directory.



# Writing the iso to a flash drive

## Mac instructions
```sh
diskutil list
diskutil unmountdisk /dev/diskX
sudo dd if=ceremony_v1.1.x.iso of=/dev/rdiskX bs=4M status=progress
diskutil list
diskutil eject /dev/diskX
```
