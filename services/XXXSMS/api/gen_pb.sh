#!/usr/bin/env bash

protoc -I . XXXSMS.proto --go_out=plugins=grpc:.