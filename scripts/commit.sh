#!/bin/sh
#
# Commits to github.com
#
set -e

msg="$@"
if [[ -z "$msg" ]]; then
	msg="update $(date)"
fi
git add .
git commit -m "$msg"
open /Applications/GitHub\ Desktop.app