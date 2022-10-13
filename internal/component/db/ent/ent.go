package ent

import "github.com/cnartlu/area-service/pkg/strings"

// Config 数据库配置
type Config struct {
	// driver 连接驱动
	Driver string `json:"driver,omitempty"`
	// source 驱动dsn连接字符集
	Source string `json:"source,omitempty"`
	// hostname 服务器地址
	Hostname string `json:"hostname,omitempty"`
	// database 数据库名
	Database string `json:"database,omitempty"`
	// username 数据库用户名
	Username string `json:"username,omitempty"`
	// password 数据库密码
	Password string `json:"password,omitempty"`
	// hostport 数据库连接端口
	Hostport int32 `json:"hostport,omitempty"`
	// params 数据库连接参数
	Params map[string]interface{} `json:"params,omitempty"`
	// charset 数据库编码默认采用utf8
	Charset string `json:"charset,omitempty"`
	// prefix 数据库表前缀
	Prefix string `json:"prefix,omitempty"`
	// timeout 超时时间
	Timeout *strings.Duration `json:"timeout,omitempty"`
}
