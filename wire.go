//go:build tools
// +build tools

// following https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

//go:generate kratos proto client internal/config/config.proto

package main

import (
	_ "github.com/google/wire/cmd/wire"
)
