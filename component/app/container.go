package app

import (
	kconfig "github.com/go-kratos/kratos/v2/config"
	klog "github.com/go-kratos/kratos/v2/log"
)

type Container interface {
	Setup(kconfig.Config, klog.Logger) error
}
