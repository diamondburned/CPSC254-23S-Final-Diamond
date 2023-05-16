#!/usr/bin/env bash
mkdir -p final/{src,data,config}
mv main.cpp header.h final/src
mv mainDB.db backupDB.db final/data
mv config.yml final/config

tree final # print the directory structure
