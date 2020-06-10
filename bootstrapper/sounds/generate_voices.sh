#!/usr/bin/env bash
set -eu -o pipefail

function generate() {
  say -v Alex "$1" -o "$2"
  sox "$2".aiff "$2".wav
  rm "$2".aiff
}

rm -f *.wav
generate "Bootstrapper initializing" initializing
generate "Bootstrapper checking uplink availability" uplink-test
generate "Bootstrapper failed to detect uplink" uplink-failed
generate "Bootstrapper using cellular uplink" uplink-cellular
generate "Bootstrapper using wired uplink" uplink-wired
generate "Bootstrapper using wifi uplink" uplink-wifi
generate "Bootstrapper installing system prerequisites" prerequisites
generate "Bootstrapper failed" generic-failure
generate "Bootstrapper succeeded" generic-success
generate "Bootstrapper downloading platform" downloading-platform
generate "Bootstrapper failed downloading platform release" downloading-failure
generate "Bootstrapper downloading diagnostics platform" downloading-platform
generate "Bootstrapper failed downloading diagnostics platform release" downloading-failure
generate "Bootstrapper acquiring initial configuration" configuration
generate "Bootstrapper failed to acquire initial configuration" configuration-failure
generate "Bootstrapper initializing platform diagnostics" diagnostics-initializing
generate "Bootstrapper failed to initialize platform diagnostics" diagnostics-failed
generate "Diagnostics platform online" diagnostics-online
generate "Bootstrapper initializing platform" platform-initializing
generate "Bootstrapper failed to initialize platform" platform-failed
generate "Platform online" platform-online
generate "Bootstrapper finished successfully" exit