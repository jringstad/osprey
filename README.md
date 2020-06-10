# osprey

## bootstrapper
Installs and updates the platform code and diagnostic tools on the device

## groundstation
command-and-control, receives telemetry

## platform
The code running on the platform

## platform-diagnostics
Extra bits running on the platform for diagnostic purposes (like reverse SSH tunnel)


## TODO
- register endpoint (send cpuinfo, meminfo etc, receive device ID, cat /proc/device-tree/model)
- telemetry collection pipeline
- prerequisites (blank image with bootstrapper + voice synth?)