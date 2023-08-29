#!/bin/sh

msg="$@"
if [[ -z "$msg" ]]; then
    msg="update $(date)"
fi
git add .
git commit -m "$msg"
git push
#TODO: open /Applications/GitHub\ Desktop.app
