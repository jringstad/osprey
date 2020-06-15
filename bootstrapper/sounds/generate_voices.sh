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
generate_bootstrapper "Bootstrapper initializing" initializing # used
generate_bootstrapper "Checking for key" key-checking # used
generate_bootstrapper "Failed to find key" key-failure # used
generate_bootstrapper "Key is present" key-success # used
generate_bootstrapper "adding repository" adding-repo # used
generate_bootstrapper "adding repository failed" adding-repo-failed # used
generate_bootstrapper "self-updating bootstrapper" self-update # used
generate_bootstrapper "installing platform" installing-platform # used
generate_bootstrapper "failed to install packages" install-failure # used
generate_bootstrapper "packages have been updated, rebooting" rebooting # used
generate_bootstrapper "nothing needed updating" no-updates # used
generate_bootstrapper "starting platform services" starting-platform-services # used
generate_bootstrapper "bootstrapper finished" finished # used
