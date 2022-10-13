package data

import (
	// 通过MySQL驱动使用Opencensus​
	_ "github.com/go-sql-driver/mysql"
	// 使用pgx驱动PostgreSQL​
	_ "github.com/jackc/pgx/v4/stdlib"
)
