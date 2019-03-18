#!/usr/bin/env bash

protoc -I proto/ proto/stream.proto --go_out=plugins=grpc:proto