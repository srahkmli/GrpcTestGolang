package config

import (
	"errors"
	"fmt"
	"log"
	"micro/config/env"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var (
	// Global config
	confs = Config{}
	lock  = sync.Mutex{}
)

// Config is base of configs we need for project
type Config struct {
	CORS     string      `yaml:"cors" required:"true"`
	Debug    bool        `yaml:"debug"`
	Env      env.EnvMode `yaml:"env" required:"true"`
	Service  Service     `yaml:"service" required:"true"`
	Jaeger   Jaeger      `yaml:"jaeger" required:"true"`
	Etcd     Etcd        `yaml:"etcd" required:"true"`
	Redis    Redis       `yaml:"redis" required:"true"`
	Postgres Database    `yaml:"database" required:"true"`
	Nats     NATS        `yaml:"nats" required:"true"`
	JWT      JWT         `yaml:"jwt" json:"jwt" required:"true"`
	Logstash Logstash    `yaml:"logstash" required:"true"`
}

func Validate(c any) error {
	errmsg := ""
	numFields := reflect.TypeOf(c).NumField()
	for i := 0; i < numFields; i++ {
		fieldType := reflect.TypeOf(c).Field(i)
		tagval, ok := fieldType.Tag.Lookup("required")
		isRequired := ok && tagval == "true"
		if !isRequired {
			continue
		}
		fieldVal := reflect.ValueOf(c).Field(i)
		if fieldVal.Kind() == reflect.Struct {
			if err := Validate(fieldVal.Interface()); err != nil {
				errmsg += fmt.Sprintf("%s > [%v], ", fieldType.Name, err)
			}
		} else {
			if fieldVal.IsZero() {
				errmsg += fmt.Sprintf("%s, ", fieldType.Name)
			}
		}
	}
	if errmsg == "" {
		return nil
	}
	return errors.New(errmsg)
}

func C() *Config {
	return &confs
}

// init configs
func InitConfigs(shutdowner fx.Shutdowner) {
	dir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.AddConfigPath(dir)
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	loadConfigs()
	viper.OnConfigChange(func(in fsnotify.Event) {
		lock.Lock()
		defer lock.Unlock()
		lastUpdate := viper.GetTime("fsnotify")
		if time.Since(lastUpdate) < time.Second {
			return
		}
		viper.Set("fsnotify", time.Now())
		log.Println("config file changed. restarting...")
		shutdowner.Shutdown()
	})
	viper.WatchConfig()
}

func loadConfigs() {
	must(viper.Unmarshal(&confs),
		"could not unmarshal config file")
	must(Validate(confs),
		"some required configurations are missing")
	log.Printf("configs loaded from file successfully \n")
}

func must(err error, logMsg string) {
	if err != nil {
		log.Println(logMsg)
		panic(err)
	}
}
