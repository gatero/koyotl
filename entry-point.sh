#!/bin/bash

LOCKFILE="./vendor/.lockfile"
MANIFEST="Gopkg.toml"

if [ ! -e Gopkg.toml ]; then
  dep init -v
  cp $MANIFEST $LOCKFILE > /dev/null 2>&1
fi

if [ -d ./vendor ]; then
  if ! diff $MANIFEST $LOCKFILE > /dev/null 2>&1; then
    dep ensure
    cp $MANIFEST $LOCKFILE > /dev/null 2>&1
  fi
fi

printf "%b" "${PURPLE_BOLD}ENTRY POINT RUNING${NO_COLOR}"

refresh run
