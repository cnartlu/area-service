//go:build !windows

package env

func ParentHttpListener() net.Listener {
	return nil
}

func ParentGrpcListener() net.Listener {
	return nil
}
