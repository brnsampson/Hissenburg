package config

import (
	"strconv"
	"crypto/tls"
)

type staticServer struct {
	Host       string
	IP         string
	Port       int
	TlsConfig  *tls.Config
	TlsEnabled bool
}

func (c staticServer) GetAddr() (string, error) {
	p := strconv.Itoa(c.Port)
	return c.IP + ":" + p, nil
}

func (c staticServer) GetTlsConfig() *tls.Config {
	return c.TlsConfig
}

func (c staticServer) GetTlsEnabled() bool {
	return c.TlsEnabled
}

type ServerConfig struct {
	staticServer
	loader *FullLoader
}

func (c *ServerConfig) Reload() error {
	reloaded, err := c.loader.WithReload()
	if err != nil {
		return err
	}

	appConf, err := reloaded.Finalize()
	if err != nil {
		return err
	}

	c.staticServer = appConf.Server
	return nil
}

type staticFileStore struct {
	Path string
}

type FileStoreConfig struct {
	staticFileStore
	loader *FullLoader
}

func (c *FileStoreConfig) Reload() error {
	reloaded, err := c.loader.WithReload()
	if err != nil {
		return err
	}

	appConf, err := reloaded.Finalize()
	if err != nil {
		return err
	}

	c.staticFileStore = appConf.FileStore
	return nil
}

type staticS3 struct {
	Host       string
	Port       int
	TlsEnabled bool
	TLSConf    *tls.Config
}

type S3Config struct {
	staticS3
	loader  *FullLoader
}

func (c *S3Config) Reload() error {
	reloaded, err := c.loader.WithReload()
	if err != nil {
		return err
	}

	appConf, err := reloaded.Finalize()
	if err != nil {
		return err
	}

	c.staticS3 = appConf.S3
	return nil
}

type staticAppConfig struct {
	Server staticServer
	S3 staticS3
	FileStore staticFileStore
}

type AppConfig struct {
	staticAppConfig
	loader *FullLoader
}

func NewConfig(loader FullLoader) (AppConfig, error) {
	l, err := loader.WithReload()
	if err != nil {
		return AppConfig{}, err
	}

	conf, err := l.Finalize()
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{conf, &loader}, err
}

func (c AppConfig) Server() *ServerConfig {
	return &ServerConfig{c.staticAppConfig.Server, c.loader}
}

func (c AppConfig) S3() *S3Config {
	return &S3Config{c.staticAppConfig.S3, c.loader}
}

func (c AppConfig) FileStore() *FileStoreConfig {
	return &FileStoreConfig{c.staticAppConfig.FileStore, c.loader}
}
