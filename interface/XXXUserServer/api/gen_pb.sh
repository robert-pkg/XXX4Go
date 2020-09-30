#!/usr/bin/env bash

protoc -I . XXXUserServer.proto --go_out=plugins=grpc:.