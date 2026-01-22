#!/bin/bash
# shellcheck disable=SC2006,SC2086,SC2116

# IMPORTANT: This file is not meant to be executed. It only serves as a test file for semgrep

# ruleid: curl-eval
eval ${x}

# ruleid: curl-eval
eval "$(curl -s "file://$p/script.sh")"
