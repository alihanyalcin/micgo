#!/bin/bash

DIR=$PWD
CMD=../cmd

# Kill all project-* stuff
function cleanup {
	pkill project
}

services

trap cleanup EXIT

while : ; do sleep 1 ; done