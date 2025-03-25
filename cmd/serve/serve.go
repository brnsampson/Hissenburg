package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/brnsampson/Hissenburg/internal/config"
	"github.com/brnsampson/Hissenburg/internal/mux"
	"github.com/brnsampson/Hissenburg/logic/startup"
	"github.com/brnsampson/httprunner"
	co "github.com/brnsampson/optional/confopt"
	"github.com/charmbracelet/log"
	"strings"
)

const (
	DEFAULT_CONFIG_FLAG = "config.toml"
)

// non-value flag vars
var (
	logLevel string
	dev      bool
)

// Override values for dev mode flag. This makes it easy to use for local development but also locks the user off from
// doing something dumb like connecting to the prod DB with their local dev server.
var (
	DEV_LOG_LEVEL       = "debug"
	DEV_HOST            = co.SomeStr("127.0.0.1")
	DEV_IP              = co.SomeStr("127.0.0.1")
	DEV_PORT            = co.SomeInt(8080)
	DEV_TLS_ENABLED     = co.SomeBool(false)
	DEV_TLS_SKIP_VERIFY = co.SomeBool(true)
	DEV_DB_HOST         = co.SomeStr("127.0.0.1")
)

func main() {
	// Configuration, logging, and other infrastructure related config
	configPath := co.SomeStr(DEFAULT_CONFIG_FLAG)

	// HTTP Server config flags
	host := co.NoStr()
	ip := co.NoStr()
	port := co.NoInt()
	tlsCert := co.NoCert()
	tlsKey := co.NoPrivateKey()
	tlsEnabled := co.NoBool()

	// DB connection flags
	dbHost := co.NoStr()
	dbPort := co.NoInt()
	rootCACert := co.NoCert()
	dbMtlsKey := co.NoPrivateKey()
	dbMtlsCert := co.NoCert()
	dbTlsEnabled := co.NoBool()
	dbTlsSkipVerify := co.NoBool()

	// local FileStore flag
	filestorePath := co.NoStr()

	flag.StringVar(&logLevel, "level", "info", "set logging level to selected value {debug, info, warn, error}")
	flag.BoolVar(&dev, "dev", false, "enable development server mode")
	flag.Var(&configPath, "config", "path to config file. Set to `none` to disable loading from co.")

	flag.Var(&host, "host", "hostname for the listening server")
	flag.Var(&ip, "ip", "bind ip for the listening server")
	flag.Var(&port, "port", "bind port for the listening server")
	flag.Var(&tlsKey, "key", "path to server TLS private key")
	flag.Var(&tlsCert, "cert", "path to server TLS certificate")
	flag.Var(&tlsEnabled, "tls_enabled", "Whether to enable or disable TLS")

	flag.Var(&dbHost, "db_host", "hostname for the backend database")
	flag.Var(&dbPort, "db_port", "connection port for the backend database")
	flag.Var(&rootCACert, "root_ca_cert", "path to root CA signing certificate used for self-signed certs. The system CertPool is also used.")
	flag.Var(&dbMtlsKey, "db_mtls_key", "path to MTLS private key for database connection")
	flag.Var(&dbMtlsCert, "db_mtls_cert", "path to MTLS certificate for database connection")
	flag.Var(&dbTlsEnabled, "db_tls_enabled", "Whether to enable or disable TLS when connecting to the database")
	flag.Var(&dbTlsSkipVerify, "db_tls_skip_verify", "Whether to verify the database's certificate against the given root certs")

	flag.Var(&filestorePath, "fs_path", "The root directory for local file storage when using the filestore backend")

	flag.Parse()

	if dev {
		logLevel = DEV_LOG_LEVEL
	}
	s := strings.ToLower(logLevel)
	if strings.HasPrefix(s, "debug") {
		log.SetLevel(log.DebugLevel)
	} else if strings.HasPrefix(s, "info") {
		log.SetLevel(log.InfoLevel)
	} else if strings.HasPrefix(s, "warn") {
		log.SetLevel(log.WarnLevel)
	} else if strings.HasPrefix(s, "error") {
		log.SetLevel(log.ErrorLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	serverLoader := config.
		NewServerLoader().
		WithHost(host).
		WithIP(ip).
		WithPort(port).
		WithTlsKey(tlsKey).
		WithTlsCert(tlsCert).
		WithTlsEnabled(tlsEnabled)

	dbLoader := config.
		NewS3Loader().
		WithHost(dbHost).
		WithPort(dbPort).
		WithRootCACert(rootCACert).
		WithMtlsKey(dbMtlsKey).
		WithMtlsCert(dbMtlsCert).
		WithTlsEnabled(dbTlsEnabled).
		WithTlsSkipVerify(dbTlsSkipVerify)

	filestoreLoader := config.
		NewFileStoreLoader().
		WithPath(filestorePath)

	// Dev mode overrides go here
	if dev {
		serverLoader = serverLoader.WithHost(DEV_HOST).WithIP(DEV_IP).WithPort(DEV_PORT).WithTlsEnabled(DEV_TLS_ENABLED)

		dbLoader = dbLoader.WithHost(DEV_DB_HOST).WithTlsEnabled(DEV_TLS_ENABLED).WithTlsSkipVerify(DEV_TLS_SKIP_VERIFY)
	}

	appLoader := config.
		NewFullLoader().
		WithConfigPath(configPath).
		WithServerLoader(serverLoader).
		WithS3Loader(dbLoader).
		WithFileStoreLoader(filestoreLoader)
	appConfig, err := config.NewConfig(appLoader)
	if err != nil {
		log.Error("Error loading appConfig from appLoader", "error", err, "load state", appLoader)
		panic("Error while loading config!")
	}

	serverConfig := appConfig.Server()
	fsConfig := appConfig.FileStore()

	// We have a static config! Easy, right?!?! ...right?
	log.Info("Full configuration loaded", "config", fmt.Sprintf("%#v", appConfig))

	err = startup.PopulateMissingItems([]string{})
	if err != nil {
		log.Error("Error while populating missing items in DB", "error", err)
		panic("Error runnin PopulateMissingItems!")
	}

	mux, err := mux.New(fsConfig.Path)
	if err != nil {
		os.Exit(1)
	}

	stdlog := log.StandardLog(log.StandardLogOptions{})
	runner := httprunner.NewRunner()
	runner.BlockingRun(mux, serverConfig, stdlog)
}
