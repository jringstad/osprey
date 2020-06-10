#!/usr/bin/env bash
set -eu -o pipefail

function generate() {
  say -v Alex "$1" -o "$2"
  sox "$2".aiff "$2".wav
  rm "$2".aiff
}

rm -f *.wav
generate "Bootstrapper initializing" initializing
generate "checking uplink availability" uplink-test
generate "failed to detect uplink" uplink-failed
generate "using cellular uplink" uplink-cellular
generate "using wired uplink" uplink-wired
generate "using wifi uplink" uplink-wifi
generate "installing system prerequisites" prerequisites
generate "failed" generic-failure
generate "succeeded" generic-success
generate "downloading platform" downloading-platform
generate "failed downloading platform release" downloading-failure
generate "downloading diagnostics platform" downloading-platform
generate "failed downloading diagnostics platform release" downloading-failure
generate "acquiring initial configuration" configuration
generate "failed to acquire initial configuration" configuration-failure
generate "initializing platform diagnostics" diagnostics-initializing
generate "failed to initialize platform diagnostics" diagnostics-failed
generate "Diagnostics platform online" diagnostics-online
generate "initializing platform" platform-initializing
generate "failed to initialize platform" platform-failed
generate "Platform online" platform-online
generate "Bootstrapper finished successfully" exit