#!/usr/bin/env bash

VOICE=alex

function generate() {
  say -v Alex "$1" -o "$2"
}

generate "Bootstrapper initializing" initializing.aiff
generate "Bootstrapper failed" generic-failure.aiff
generate "Bootstrapper succeeded" generic-success.aiff
generate "Bootstrapper downloading release" downloading.aiff
generate "Bootstrapper failed downloading release" downloading-failure.aiff
generate "Bootstrapper acquiring initial configuration" configuration.aiff
generate "Bootstrapper failed to acquire initial configuration" configuration-failure.aiff