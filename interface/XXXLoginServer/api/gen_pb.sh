#!/usr/bin/env bash

protoc -I . XXXLoginServer.proto --go_out=plugins=grpc:.