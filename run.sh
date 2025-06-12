#!/bin/bash

# Note: "go build" will ignore *test* files:
# "-o" creates a "bookings" file
# build script but only run if build succeeds
go build -o bookings cmd/web/*.go && ./bookings