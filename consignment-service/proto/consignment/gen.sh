#!/usr/bin/env bash

protoc -I. --go_out=plugins=micro:. ./consignment.proto