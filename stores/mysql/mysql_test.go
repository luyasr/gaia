package mysql

import (
	"testing"
)

type MysqlCfg struct {
	Mysql *Config
}

func TestNewMysql(t *testing.T) {
	parseTime := true
	// 创建一个模拟的 Config 对象
	c := &Config{
		Username:  "root",
		Password:  "12345678",
		DataBase:  "test",
		Host:      "localhost",
		Port:      3306,
		ParseTime: &parseTime,
	}

	// 调用 NewMysql 函数
	m, err := New(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(m.Client)
}
