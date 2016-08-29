package etc

import "gopkg.in/yaml.v2"
import gb "github.com/gtlservice/gtlgateway/base"
import "github.com/gtlservice/gtlappserver/api"
import "github.com/gtlservice/gtlappserver/base"
import "github.com/gtlservice/gutils/logger"
import "github.com/gtlservice/gzkwrapper"

import (
	"io/ioutil"
	"os"
)

type Parameters map[string]string
type Cors Parameters

type Configuration struct {
	Version   string `yaml:"version,omitempty"`
	Fork      bool   `yaml:"fork,omitempty"`
	PidFile   string `yaml:"pidfile,omitempty"`
	ZkWrapper struct {
		Hosts     string `yaml:"hosts,omitempty"`
		Root      string `yaml:"root,omitempty"`
		Device    string `yaml:"device,omitempty"`
		Location  string `yaml:"location,omitempty"`
		OS        string `yaml:"os,omitempty"`
		Platform  string `yaml:"platform,omitempty"`
		Pulse     string `yaml:"pulse,omitempty"`
		Threshold int    `yaml:"threshold,omitempty"`
	} `yaml:"zkwrapper,omitempty"`
	Service struct {
		Name string `yaml:"name,omitempty"`
		Host string `yaml:"host,omitempty"`
	} `yaml:"service,omitempty"`
	ApiServer struct {
		Bind string `yaml:"bind,omitempty"`
		Cors Cors   `yaml:"cors,omitempty"`
	} `yaml:"api,omitempty"`
	Logger struct {
		LogFile  string `yaml:"logfile,omitempty"`
		LogLevel string `yaml:"loglevel,omitempty"`
		LogSize  int64  `yaml:"logsize,omitempty"`
	} `yaml:"logger,omitempty"`
}

var configuration *Configuration

func NewConfiguration(file string) (*Configuration, error) {

	fp, err := os.OpenFile(file, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	defer fp.Close()
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	c := makeDefault()
	if err := yaml.Unmarshal([]byte(data), c); err != nil {
		return nil, err
	}
	configuration = c
	return configuration, nil
}

func makeDefault() *Configuration {

	return &Configuration{
		Version: "default",
		Fork:    false,
		PidFile: "./gtlappserver.pid",
		ZkWrapper: struct {
			Hosts     string `yaml:"hosts,omitempty"`
			Root      string `yaml:"root,omitempty"`
			Device    string `yaml:"device,omitempty"`
			Location  string `yaml:"location,omitempty"`
			OS        string `yaml:"os,omitempty"`
			Platform  string `yaml:"platform,omitempty"`
			Pulse     string `yaml:"pulse,omitempty"`
			Threshold int    `yaml:"threshold,omitempty"`
		}{
			Hosts:     "127.0.0.1:2818",
			Root:      "/gtlservice",
			Device:    "",
			Location:  "center",
			OS:        "",
			Platform:  "",
			Pulse:     "10s",
			Threshold: 3,
		},
		Service: struct {
			Name string `yaml:"name,omitempty"`
			Host string `yaml:"host,omitempty"`
		}{
			Name: "service1",
			Host: "",
		},
		ApiServer: struct {
			Bind string `yaml:"bind,omitempty"`
			Cors Cors   `yaml:"cors,omitempty"`
		}{
			Bind: ":8982",
			Cors: map[string]string{
				"origin":  "*",
				"methods": "GET",
			},
		},
		Logger: struct {
			LogFile  string `yaml:"logfile,omitempty"`
			LogLevel string `yaml:"loglevel,omitempty"`
			LogSize  int64  `yaml:"logsize,omitempty"`
		}{
			LogFile:  "logs/jobworker.log",
			LogLevel: "debug",
			LogSize:  2097152,
		},
	}
}

func GetConfiguration() *Configuration {

	return configuration
}

func (c *Configuration) GetVersion() string {

	if c != nil {
		return c.Version
	}
	return "default"
}

func (c *Configuration) GetFork() bool {

	if c != nil {
		return c.Fork
	}
	return false
}

func (c *Configuration) GetPidFile() string {

	if c != nil {
		return c.PidFile
	}
	return ""
}

func (c *Configuration) GetZkWrapper() *gzkwrapper.WorkerArgs {

	if c != nil {
		return &gzkwrapper.WorkerArgs{
			Hosts:     c.ZkWrapper.Hosts,
			Root:      c.ZkWrapper.Root,
			Device:    c.ZkWrapper.Device,
			Location:  c.ZkWrapper.Location,
			OS:        c.ZkWrapper.OS,
			Platform:  c.ZkWrapper.Platform,
			Pulse:     c.ZkWrapper.Pulse,
			Threshold: c.ZkWrapper.Threshold,
		}
	}
	return nil
}

func (c *Configuration) GetService() *base.ServiceArgs {

	if c != nil {
		return &base.ServiceArgs{
			Service: gb.Service{
				Name: c.Service.Name,
				Host: c.Service.Host,
			},
		}
	}
	return nil
}

func (c *Configuration) GetApiServer() *api.ApiServerArgs {

	if c != nil {
		return &api.ApiServerArgs{
			Bind: c.ApiServer.Bind,
			Cors: c.ApiServer.Cors,
		}
	}
	return nil
}

func (c *Configuration) GetLogger() *logger.Args {

	if c != nil {
		return &logger.Args{
			FileName: c.Logger.LogFile,
			Level:    c.Logger.LogLevel,
			MaxSize:  c.Logger.LogSize,
		}
	}
	return nil
}
