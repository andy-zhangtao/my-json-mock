package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
)

type RunParams struct {
	Configure Configure `toml:"configure"`
	Mysql     Mysql     `toml:"mysql"`
}
type Configure struct {
	Debug bool `toml:"debug"`
	Port  int  `toml:"port"`
}
type Mysql struct {
	Dsn    string `toml:"dsn"`
	User   string `toml:"user"`
	Passwd string `toml:"passwd"`
	Db     string `toml:"db"`
}

// ParseParams 解析参数,返回RunParams
func ParseParams(data []byte) (rp RunParams, err error) {
	err = toml.Unmarshal(data, &rp)
	if err != nil {
		return rp, err
	}

	return rp, nil
}

func Output(rp RunParams) {
	fmt.Println("================")
	fmt.Printf("debug: %t", rp.Configure.Debug)
	fmt.Printf("port: %d", rp.Configure.Port)
	fmt.Println("================")
}
