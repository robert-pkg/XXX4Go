package conf

import (
	"errors"
	"flag"
	"io/ioutil"

	"github.com/robert-pkg/micro-go/cache/redis"
	"github.com/robert-pkg/micro-go/db/mysql"

	"github.com/robert-pkg/micro-go/log"
	jaeger_trace "github.com/robert-pkg/micro-go/trace/jaeger-trace"
	"gopkg.in/yaml.v2"
)

// Conf global variable.
var (
	confPath string
	Conf     = &Config{}
)

// Config struct of conf.
type Config struct {
	Log log.LogConfig `yaml:"log"`

	TraceConfig jaeger_trace.Config `yaml:"jaeger"`

	// 登录数据库
	MysqlConfig *mysql.Config `yaml:"db"`
	RedisConfig *redis.Config `yaml:"redis"`
}

func (c *Config) loadFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init int config
func Init() error {

	if confPath != "" {
		return Conf.loadFromFile(confPath)
	}

	registerCodes()

	return errors.New("暂未实现配置中心")

}
