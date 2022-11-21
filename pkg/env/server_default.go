//go:build !windows

package env

import "net"

func ParentHttpListener() net.Listener {
	return nil
}

func ParentGrpcListener() net.Listener {
	return nil
}
