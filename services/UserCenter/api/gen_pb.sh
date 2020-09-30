#!/usr/bin/env bash

protoc -I . UserCenter.proto --go_out=plugins=grpc:.