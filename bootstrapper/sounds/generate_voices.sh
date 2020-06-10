#!/usr/bin/env bash
set -eu -o pipefail

function generate_bootstrapper() {
  say -v Alex "$1" -o "$2"
  sox "$2".aiff "$2".wav
  rm "$2".aiff
}

function generate_diagnostics() {
  say -v Samantha "$1" -o "$2"
  sox "$2".aiff "$2".wav
  rm "$2".aiff
}

rm -f *.wav
generate_bootstrapper "Bootstrapper initializing" initializing
generate_bootstrapper "checking uplink availability" uplink-test
generate_bootstrapper "failed to detect uplink" uplink-failed
generate_bootstrapper "using cellular uplink" uplink-cellular
generate_bootstrapper "using wired uplink" uplink-wired
generate_bootstrapper "using wifi uplink" uplink-wifi
generate_bootstrapper "installing system prerequisites" prerequisites
generate_bootstrapper "failed" generic-failure
generate_bootstrapper "succeeded" generic-success
generate_bootstrapper "downloading platform" downloading-platform
generate_bootstrapper "failed downloading platform release" downloading-failure
generate_bootstrapper "downloading diagnostics platform" downloading-platform
generate_bootstrapper "failed downloading diagnostics platform release" downloading-failure
generate_bootstrapper "acquiring initial configuration" configuration
generate_bootstrapper "failed to acquire initial configuration" configuration-failure
generate_bootstrapper "initializing platform diagnostics" diagnostics-initializing
generate_bootstrapper "failed to initialize platform diagnostics" diagnostics-failed
generate_bootstrapper "initializing platform" platform-initializing
generate_bootstrapper "failed to initialize platform" platform-failed
generate_bootstrapper "Osprey platform online" platform-online
generate_bootstrapper "Bootstrapper finished successfully" exit


generate_diagnostics "Diagnostics platform online" diagnostics-online
generate_diagnostics "Engines armed" engines-armed
generate_diagnostics "Engines disarmed" engines-disarmed