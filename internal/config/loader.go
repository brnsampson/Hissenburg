package config

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/brnsampson/optional"
	o "github.com/brnsampson/optional/confopt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	"github.com/charmbracelet/log"
)

// Defaults values for ServerLoader, which loads and generates a staticServerConfig
const (
	DEFAULT_HOST        = "localhost"
	DEFAULT_IP          = "127.0.0.1"
	DEFAULT_PORT        = 3000
	DEFAULT_TLS_ENABLED = true
	DEFAULT_TLS_CERT    = "/etc/scammmateo/tls/cert.pem"
	DEFAULT_TLS_KEY     = "/etc/scammmateo/tls/key.pem"
)

type LoadingError struct {
	msg string
}

func NewLoadingError(msg string) LoadingError {
	return LoadingError{msg}
}

func (e LoadingError) Error() string {
	return e.msg
}

// Generic server configuration which can be reloaded on demand.
type ServerLoader struct {
	Host       o.Str        `env:"HISS_HOST"`
	IP         o.Str        `env:"HISS_BIND_IP"`
	Port       o.Int        `env:"HISS_PORT"`
	TlsCert    o.Cert       `env:"HISS_TLS_CERT"`
	TlsKey     o.PrivateKey `env:"HISS_TLS_KEY"`
	TlsEnabled o.Bool       `env:"HISS_TLS_ENABLED"`
}

func SomeServerLoader(loader ServerLoader) optional.Option[ServerLoader] {
	return optional.Some(loader)
}

func NoServerLoader() optional.Option[ServerLoader] {
	return optional.None[ServerLoader]()
}

func NewServerLoader() ServerLoader {
	return ServerLoader {
		Host: o.NoStr(),
		IP: o.NoStr(),
		Port: o.NoInt(),
		TlsCert: o.NoCert(),
		TlsKey: o.NoPrivateKey(),
		TlsEnabled: o.NoBool(),
	}
}

func serverLoaderFromEnv(override optional.Option[ServerLoader]) (loader ServerLoader, err error) {
	loader = NewServerLoader()
	if err = env.Parse(&loader); err != nil {
		log.Error("Failed to load ServerLoader from env variables!")
		return
	}

	log.Debug("Loaded ServerLoader from env variables", "ServerConfig", loader)
	over := optional.GetOrElse(override, NewServerLoader)
	loader = over.Merged(loader)
	return
}

const (
	DEFAULT_S3_HOST            = "localhost"
	DEFAULT_S3_PORT            = 1234
	DEFAULT_S3_ROOT_CA_CERT    = "/etc/scammmateo/db/ca.pem"
	DEFAULT_S3_MTLS_CERT       = "/etc/scammmateo/db/cert.pem"
	DEFAULT_S3_MTLS_KEY        = "/etc/scammmateo/db/key.pem"
	DEFAULT_S3_TLS_ENABLED     = true
	DEFAULT_S3_TLS_SKIP_VERIFY = false
)

type S3Loader struct {
	Host          o.Str        `env:"HISS_S3_HOST"`
	Port          o.Int        `env:"HISS_S3_PORT"`
	RootCACert    o.Cert       `env:"HISS_ROOT_CA"`
	MtlsCert      o.Cert       `env:"HISS_S3_MTLS_CERT"`
	MtlsKey       o.PrivateKey `env:"HISS_S3_MTLS_CERT"`
	TlsEnabled    o.Bool       `env:"HISS_S3_TLS_ENABLED"`
	TlsSkipVerify o.Bool       `env:"HISS_S3_TLS_SKIP_VERIFY"`
}

func SomeS3Loader(loader S3Loader) optional.Option[S3Loader] {
	return optional.Some(loader)
}

func NoS3Loader() optional.Option[S3Loader] {
	return optional.None[S3Loader]()
}

func NewS3Loader() S3Loader {
	return S3Loader {
		Host: o.NoStr(),
		Port: o.NoInt(),
		RootCACert: o.NoCert(),
		MtlsCert: o.NoCert(),
		MtlsKey: o.NoPrivateKey(),
		TlsEnabled: o.NoBool(),
		TlsSkipVerify: o.NoBool(),
	}
}

func dbLoaderFromEnv(override optional.Option[S3Loader]) (loader S3Loader, err error) {
	loader = NewS3Loader()
	if err = env.Parse(&loader); err != nil {
		log.Error("Failed to load S3Loader from env variables!")
		return
	}

	log.Debug("Loaded S3Loader from env variables", "S3Config", loader)

	over := optional.GetOrElse(override, NewS3Loader)
	loader = over.Merged(loader)
	return
}

const (
	DEFAULT_FILESTORE_PATH = "./store"
)

type FileStoreLoader struct {
	Path o.Str   `env:"HISS_FILESTORE_PATH"`
}

func NewFileStoreLoader() FileStoreLoader {
	return FileStoreLoader{
		Path: o.NoStr(),
	}
}

func fileStoreLoaderFromEnv(override optional.Option[FileStoreLoader]) (loader FileStoreLoader, err error) {
	loader = NewFileStoreLoader()
	if err = env.Parse(&loader); err != nil {
		log.Error("Failed to load FileStoreLoader from env variables!")
		return
	}

	log.Debug("Loaded FileStoreLoader from env variables", "FileStoreConfig", loader)

	over := optional.GetOrElse(override, NewFileStoreLoader)
	loader = over.Merged(loader)
	return
}

type FullLoader struct {
	ConfigPath o.Str
	Server ServerLoader
	S3     S3Loader
	FileStore FileStoreLoader
}

func SomeFullLoader(loader FullLoader) optional.Option[FullLoader] {
	return optional.Some(loader)
}

func NoFullLoader() optional.Option[FullLoader] {
	return optional.None[FullLoader]()
}

func NewFullLoader() FullLoader {
	return FullLoader {
		Server: NewServerLoader(),
		S3: NewS3Loader(),
	}
}

func fullLoaderFromEnv() (loader FullLoader, err error) {
	if err = env.Parse(&loader); err != nil {
		log.Error("Failed to load ServerLoader from env variables!")
		return
	}

	loader.Server, err = serverLoaderFromEnv(NoServerLoader())
	if err != nil {
		return
	}

	loader.S3, err = dbLoaderFromEnv(NoS3Loader())
	if err != nil {
		return
	}

	log.Debug("Loaded FullLoader from env variables", "FullConfig", loader)
	return
}

func fullLoaderFromFile(pathStr o.Str) (loader FullLoader, err error) {
	loader = NewFullLoader()

	path, err := pathStr.Get()
	if err != nil {
		log.Debug("config file path was None. Skipping...")
		return loader, nil
	} else {
		loader = loader.WithConfigPath(pathStr)
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		log.Info("Could not get absolute path from path", "path", path)
		return
	}

	if _, err = os.Stat(abs); err != nil {
		log.Debug("Could not stat config file path. Does the file exist? Skipping...", "path", abs)
		return loader, nil
	}

	_, err = toml.DecodeFile(abs, &loader)
	if err != nil {
		return
	}

	log.Info("Loaded Settings config from file", "path", abs, "config", loader)

	return
}

func loadedFullLoader(pathStr o.Str, override optional.Option[FullLoader]) (loader FullLoader, err error) {
	loader = optional.GetOrElse(override, NewFullLoader)

	path := optional.Or(pathStr, loader.ConfigPath)
	file, err := fullLoaderFromFile(path)
	if err != nil {
		return
	}
	loader = loader.Merged(file)

	env, err := fullLoaderFromEnv()
	if err != nil {
		return
	}

	loader = loader.Merged(env)
	return
}

// ServerLoader methods
func (l ServerLoader) AsOption() optional.Option[ServerLoader] {
	return optional.Some(l)
}

func (l ServerLoader) WithHost(with o.Str) ServerLoader {
	l.Host = optional.Or(with, l.Host)
	return l
}

func (l ServerLoader) OrHost(or o.Str) ServerLoader {
	l.Host = optional.Or(l.Host, or)
	return l
}

func (l ServerLoader) WithIP(with o.Str) ServerLoader {
	l.IP = optional.Or(with, l.IP)
	return l
}

func (l ServerLoader) OrIP(or o.Str) ServerLoader {
	l.IP = optional.Or(l.IP, or)
	return l
}

func (l ServerLoader) WithPort(with o.Int) ServerLoader {
	l.Port = optional.Or(with, l.Port)
	return l
}

func (l ServerLoader) OrPort(or o.Int) ServerLoader {
	l.Port = optional.Or(l.Port, or)
	return l
}

func (l ServerLoader) WithTlsCert(with o.Cert) ServerLoader {
	l.TlsCert = optional.Or(with, l.TlsCert)
	return l
}

func (l ServerLoader) OrTlsCert(or o.Cert) ServerLoader {
	l.TlsCert = optional.Or(l.TlsCert, or)
	return l
}

func (l ServerLoader) WithTlsKey(with o.PrivateKey) ServerLoader {
	l.TlsKey = optional.Or(with, l.TlsKey)
	return l
}

func (l ServerLoader) OrTlsKey(or o.PrivateKey) ServerLoader {
	l.TlsKey = optional.Or(l.TlsKey, or)
	return l
}

func (l ServerLoader) WithTlsEnabled(with o.Bool) ServerLoader {
	l.TlsEnabled = optional.Or(with, l.TlsEnabled)
	return l
}

func (l ServerLoader) OrTlsEnabled(or o.Bool) ServerLoader {
	l.TlsEnabled = optional.Or(l.TlsEnabled, or)
	return l
}

func (l ServerLoader) Merged(other ServerLoader) ServerLoader {
	return l.OrHost(other.Host).
		OrIP(other.IP).
		OrPort(other.Port).
		OrTlsCert(other.TlsCert).
		OrTlsKey(other.TlsKey).
		OrTlsEnabled(other.TlsEnabled)
}

func (l ServerLoader) WithEnv() (loader ServerLoader, err error) {
	loader, err = serverLoaderFromEnv(l.AsOption())
	if err != nil {
		return
	}

	return l, err
}

func (l ServerLoader) Finalize() (staticServer, error) {
	host := optional.GetOr(l.Host, DEFAULT_HOST)
	ip := optional.GetOr(l.IP, DEFAULT_IP)
	port := optional.GetOr(l.Port, DEFAULT_PORT)
	tls_enabled := optional.GetOr(l.TlsEnabled, DEFAULT_TLS_ENABLED)

	tls_conf := &tls.Config{
		Certificates: make([]tls.Certificate, 0),

		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	if tls_enabled {
		if l.TlsCert.IsNone() || l.TlsKey.IsNone() {
			log.Error("TLS was enabled, but the cert and key path were not both populated!", "loader", "ServerLoader", "function", "Finalize()",   "config", l)
			return staticServer{}, NewLoadingError("Missing TLS certificate or key for listening server")
		}
		tls_cert, err := l.TlsKey.ReadCert(l.TlsCert)
		if err != nil {
			return staticServer{}, err
		}

		tls_conf.Certificates = append(tls_conf.Certificates, tls_cert)
	}

	return staticServer{host, ip, port, tls_conf, tls_enabled}, nil
}

// S3Loader methods

func (l S3Loader) AsOption() optional.Option[S3Loader] {
	return optional.Some(l)
}

func (l S3Loader) WithHost(with o.Str) S3Loader {
	l.Host = optional.Or(with, l.Host)
	return l
}

func (l S3Loader) OrHost(or o.Str) S3Loader {
	l.Host = optional.Or(l.Host, or)
	return l
}

func (l S3Loader) WithPort(with o.Int) S3Loader {
	l.Port = optional.Or(with, l.Port)
	return l
}

func (l S3Loader) OrPort(or o.Int) S3Loader {
	l.Port = optional.Or(l.Port, or)
	return l
}

func (l S3Loader) WithRootCACert(with o.Cert) S3Loader {
	l.RootCACert = optional.Or(with, l.RootCACert)
	return l
}

func (l S3Loader) OrRootCACert(or o.Cert) S3Loader {
	l.RootCACert = optional.Or(l.RootCACert, or)
	return l
}

func (l S3Loader) WithMtlsKey(with o.PrivateKey) S3Loader {
	l.MtlsKey = optional.Or(with, l.MtlsKey)
	return l
}

func (l S3Loader) OrMtlsKey(or o.PrivateKey) S3Loader {
	l.MtlsKey = optional.Or(l.MtlsKey, or)
	return l
}

func (l S3Loader) WithMtlsCert(with o.Cert) S3Loader {
	l.MtlsCert = optional.Or(with, l.MtlsCert)
	return l
}

func (l S3Loader) OrMtlsCert(or o.Cert) S3Loader {
	l.MtlsCert = optional.Or(l.MtlsCert, or)
	return l
}

func (l S3Loader) WithTlsEnabled(with o.Bool) S3Loader {
	l.TlsEnabled = optional.Or(with, l.TlsEnabled)
	return l
}

func (l S3Loader) OrTlsEnabled(or o.Bool) S3Loader {
	l.TlsEnabled = optional.Or(l.TlsEnabled, or)
	return l
}

func (l S3Loader) WithTlsSkipVerify(with o.Bool) S3Loader {
	l.TlsSkipVerify = optional.Or(with, l.TlsSkipVerify)
	return l
}

func (l S3Loader) OrTlsSkipVerify(or o.Bool) S3Loader {
	l.TlsSkipVerify = optional.Or(l.TlsSkipVerify, or)
	return l
}

func (l S3Loader) Merged(other S3Loader) S3Loader {
	return l.OrHost(other.Host).
		OrPort(other.Port).
		OrRootCACert(other.RootCACert).
		OrMtlsKey(other.MtlsKey).
		OrMtlsCert(other.MtlsCert).
		OrTlsEnabled(other.TlsEnabled).
		OrTlsSkipVerify(other.TlsSkipVerify)
}

func (l S3Loader) WithEnv() (loader S3Loader, err error) {
	loader, err = dbLoaderFromEnv(l.AsOption())
	if err != nil {
		return l, err
	}
	return
}

func (l S3Loader) Finalize() (conf staticS3, err error) {
	host := optional.GetOr(l.Host, DEFAULT_S3_HOST)
	port := optional.GetOr(l.Port, DEFAULT_S3_PORT)
	tlsEnabled := optional.GetOr(l.TlsEnabled, DEFAULT_S3_TLS_ENABLED)
	tlsSkipVerify := optional.GetOr(l.TlsSkipVerify, DEFAULT_S3_TLS_SKIP_VERIFY)

	// Reading the rootCACerts will cause an error if its value is None, but we don't care if TlsSkipVerify is set anyways
	rootCACertPool, err := x509.SystemCertPool()
	if err != nil {
		return
	}

	if !tlsSkipVerify {
		if l.RootCACert.IsNone() {
			log.Error("RootCACert not set, but TlsSkipVerify == false!", "loader", "S3Loader", "function", "Finalize()", "config", l)
			return conf, NewLoadingError("RootCACert must be set if TlsSkipVerify is false")
		}

		rootCACerts, err := l.RootCACert.ReadCerts()
		if err != nil {
			return conf, err
		}

		for _, cert := range rootCACerts {
			rootCACertPool.AddCert(cert)
		}
	}

	tlsConf := &tls.Config{
		Certificates:       make([]tls.Certificate, 0),
		RootCAs:            rootCACertPool,
		InsecureSkipVerify: tlsSkipVerify,

		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	if tlsEnabled {
		if l.MtlsCert.IsNone() || l.MtlsKey.IsNone() {
			log.Error("S3 MTLS was enabled, but the mtls cert and key path were not both populated!", "loader", "S3Loader", "function", "Finalize()", "config", l)
			return conf, NewLoadingError("Missing MTLS certificate or key for S3 connection")
		}

		mtlsCert, err := l.MtlsKey.ReadCert(l.MtlsCert)
		if err != nil {
			return conf, err
		}

		tlsConf.Certificates = append(tlsConf.Certificates, mtlsCert)
	}


	return staticS3{host, port, tlsEnabled, tlsConf}, nil
}

// FileStoreLoader methods
func (l FileStoreLoader) AsOption() optional.Option[FileStoreLoader] {
	return optional.Some(l)
}

func (l FileStoreLoader) WithPath(with o.Str) FileStoreLoader {
	l.Path = optional.Or(with, l.Path)
	return l
}

func (l FileStoreLoader) OrPath(or o.Str) FileStoreLoader {
	l.Path = optional.Or(l.Path, or)
	return l
}

func (l FileStoreLoader) Merged(other FileStoreLoader) FileStoreLoader {
	return l.OrPath(other.Path)
}

func (l FileStoreLoader) WithEnv() (loader FileStoreLoader, err error) {
	loader, err = fileStoreLoaderFromEnv(l.AsOption())
	if err != nil {
		return l, err
	}
	return
}

func (l FileStoreLoader) Finalize() (conf staticFileStore, err error) {
	path := optional.GetOr(l.Path, DEFAULT_FILESTORE_PATH)
	return staticFileStore{ Path: path }, nil
}

// FullLoader methods
func (l FullLoader) AsOption() optional.Option[FullLoader] {
	return optional.Some(l)
}

func (l FullLoader) WithConfigPath(with o.Str) FullLoader {
	l.ConfigPath = optional.Or(with, l.ConfigPath)
	return l
}

func (l FullLoader) OrConfigPath(or o.Str) FullLoader {
	l.ConfigPath = optional.Or(or, l.ConfigPath)
	return l
}

func (l FullLoader) WithServerLoader(with ServerLoader) FullLoader {
	l.Server = with
	return l
}

func (l FullLoader) WithS3Loader(with S3Loader) FullLoader {
	l.S3 = with
	return l
}

func (l FullLoader) WithFileStoreLoader(with FileStoreLoader) FullLoader {
	l.FileStore = with
	return l
}

func (l FullLoader) Merged(other FullLoader) FullLoader {
	l = l.OrConfigPath(other.ConfigPath)
	l.Server = l.Server.Merged(other.Server)
	l.S3 = l.S3.Merged(other.S3)
	return l
}

func (l FullLoader) WithReload() (loader FullLoader, err error) {
	loader, err = loadedFullLoader(l.ConfigPath, l.AsOption())
	if err != nil {
		return l, err
	}
	return
}

func (l FullLoader) Finalize() (conf staticAppConfig, err error) {
	server, err := l.Server.Finalize()
	if err != nil {
		return
	}

	db, err := l.S3.Finalize()
	if err != nil {
		return
	}

	filestore, err := l.FileStore.Finalize()
	if err != nil {
		return
	}

	return staticAppConfig{server, db, filestore}, nil
}
