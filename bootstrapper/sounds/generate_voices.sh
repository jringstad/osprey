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
generate_bootstrapper "Checking for key" key-checking
generate_bootstrapper "Failed to find key" key-failure
generate_bootstrapper "Key is present" key-success
generate_bootstrapper "adding repository" adding-repo
generate_bootstrapper "adding repository failed" adding-repo-failed
generate_bootstrapper "self-updating bootstrapper" self-update
generate_bootstrapper "installing platform" installing-platform
generate_bootstrapper "failed to install packages" install-failure
generate_bootstrapper "packages have been updated, rebooting" rebooting
generate_bootstrapper "nothing needed updating" no-updates
generate_bootstrapper "starting platform services" starting-platform-services
generate_bootstrapper "bootstrapper finished" finished
