# Installing LineageOS with microG and Magisk on Guacamole Device

## Introduction

This guide will walk you through installing LineageOS with microG and Magisk on a "guacamole" device, which could be a OnePlus 7 Pro, OnePlus 7, or OnePlus 7T Pro.

## Prerequisites

- Make sure to unlock your device's bootloader before proceeding. **Do not relock the bootloader after installing a custom OS.**

## Steps

### 1. Follow LineageOS Documentation

First, follow the steps from the [official LineageOS documentation](https://wiki.lineageos.org/devices/guacamole) up to the point of downloading the LineageOS image.

### 2. Download the LineageOS with microG Image

Instead of the official LineageOS image, download the version with microG from [here](https://download.lineage.microg.org/guacamole/).

### 3. Install LineageOS with microG

Continue with the installation as described in the official LineageOS documentation, but use the LineageOS image with microG that you just downloaded.

### 4. Install Magisk

In the step where the official documentation suggests installing Google Apps, install Magisk instead. You can find Magisk on [GitHub](https://github.com/topjohnwu/Magisk).

## Troubleshooting

- If anything goes wrong, perform a wipe and start over.
  
## Important Notes

- **Do not relock the bootloader** after installing a custom OS, as it may permanently lock your device.
  
- Relocking the bootloader is generally safe only on Pixel devices, which allow system signature modification.

## Summary

Smartphones have an internal system that holds permanent information about the system provider. If the bootloader is locked and the system provider doesn't match, the phone will lock itself for "data safety."
