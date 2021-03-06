---
title: "Getting Started"
date: 2020-01-26T09:35:23-05:00
draft: false
weight: 3
summary: Quickstart
---

```bash
$ ipsw --help

Download and Parse IPSWs

Usage:
  ipsw [command]

Available Commands:
  device-list     List all iOS devices
  disass          🚧 [WIP] Disassemble ARM binaries at address or symbol
  download        Download and parse IPSW(s) from the internets
  dtree           Parse DeviceTree
  dyld            Parse dyld_shared_cache
  ent             Search IPSW filesystem DMG for MachOs with a given entitlement
  extract         Extract kernelcache, dyld_shared_cache or DeviceTree from IPSW
  help            Help about any command
  iboot           Dump firmwares
  img4            Parse Img4
  info            Display IPSW Info
  kernel          Parse kernelcache
  lipo            Extract single MachO out of a universal/fat MachO
  macho           Parse a MachO file
  ota             Extract file(s) from OTA
  sepfw           Dump MachOs
  shsh            Get shsh blobs from device
  symbolicate     Symbolicate ARM 64-bit crash logs (similar to Apple's symbolicatecrash)
  update          Download an ipsw update if one exists
  version         Print the version number of ipsw

Flags:
  -h, --help      help for ipsw
  -V, --verbose   verbose output

Use "ipsw [command] --help" for more information about a command.
```
