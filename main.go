package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/st3v/runtime-configurator/bosh"
	"github.com/st3v/runtime-configurator/configurator"
	"github.com/st3v/runtime-configurator/deployer"
)

var (
	dirTarget         string
	dirUser           string
	dirPassword       string
	dirCACert         string
	boshLogLevel      string
	uaaTarget         string
	uaaClientID       string
	uaaClientSecret   string
	uaaCACert         string
	runtimeConfigName string
	addManifest       string
	removeManifest    string
	deploy            bool
	deployDryRun      bool
	deploySkip        deployments
)

type deployments []string

func (d *deployments) String() string {
	return strings.Join(*d, ", ")
}

func (d *deployments) Set(v string) error {
	for _, dep := range strings.Split(v, ",") {
		*d = append(*d, dep)
	}
	return nil
}

func init() {
	log.SetFlags(0)

	flag.StringVar(
		&dirTarget,
		"director.url",
		getEnvString("DIRECTOR_URL", ""),
		"BOSH director URL [DIRECTOR_URL]",
	)

	flag.StringVar(
		&dirUser,
		"director.user",
		getEnvString("DIRECTOR_USER", ""),
		"BOSH director user [DIRECTOR_USER]",
	)

	flag.StringVar(
		&dirPassword,
		"director.password",
		getEnvString("DIRECTOR_PASSWORD", ""),
		"BOSH director password [DIRECTOR_PASSWORD]",
	)

	flag.StringVar(
		&dirCACert,
		"director.ca-cert",
		getEnvString("DIRECTOR_CA_CERT", ""),
		"BOSH director CA cert path [DIRECTOR_CA_CERT]",
	)

	flag.StringVar(
		&uaaTarget,
		"uaa.url",
		getEnvString("UAA_URL", ""),
		"UAA URL [UAA_URL]",
	)

	flag.StringVar(
		&uaaClientID,
		"uaa.client-id",
		getEnvString("UAA_CLIENT_ID", "bosh_cli"),
		"UAA client id [UAA_CLIENT_ID]",
	)

	flag.StringVar(
		&uaaClientSecret,
		"uaa.client-secret",
		getEnvString("UAA_CLIENT_SECRET", ""),
		"UAA client secret [UAA_CLIENT_SECRET]",
	)

	flag.StringVar(
		&uaaCACert,
		"uaa.ca-cert",
		getEnvString("UAA_CA_CERT", ""),
		"UAA CA cert path [UAA_CA_CERT]",
	)

	flag.StringVar(
		&runtimeConfigName,
		"runtime-config.name",
		getEnvString("RUNTIME_CONFIG_NAME", ""),
		"BOSH runtime-config name [RUNTIME_CONFIG_NAME]",
	)

	flag.StringVar(
		&addManifest,
		"runtime-config.add",
		getEnvString("RUNTIME_CONFIG_ADD", ""),
		"Path to BOSH runtime-config manifest to add [RUNTIME_CONFIG_ADD]",
	)

	flag.StringVar(
		&removeManifest,
		"runtime-config.remove",
		getEnvString("RUNTIME_CONFIG_REMOVE", ""),
		"Path to BOSH runtime-config manifest to remove [RUNTIME_CONFIG_REMOVE]",
	)

	flag.BoolVar(
		&deploy,
		"runtime-config.deploy",
		getEnvBool("RUNTIME_CONFIG_DEPLOY", false),
		"Update existing deployments [RUNTIME_CONFIG_DEPLOY]",
	)

	flag.BoolVar(
		&deployDryRun,
		"runtime-config.deploy.dry-run",
		getEnvBool("RUNTIME_CONFIG_DEPLOY_DRY_RUN", false),
		"Deploy dry-run for existing deployments [RUNTIME_CONFIG_DEPLOY_DRY_RUN]",
	)

	flag.Var(
		&deploySkip,
		"runtime-config.deploy.skip",
		"Comma-separated list of deployment names that should be skipped during deploy operation [RUNTIME_CONFIG_DEPLOY_SKIP]",
	)

	flag.StringVar(
		&boshLogLevel,
		"bosh.log-level",
		getEnvString("BOSH_LOG_LEVEL", "info"),
		"BOSH log level (debug, info, warn, error, none) [BOSH_LOG_LEVEL]",
	)
}

func main() {
	flag.Parse()

	if len(deploySkip) == 0 {
		flag.Set("runtime-config.deploy.skip", getEnvString("RUNTIME_CONFIG_DEPLOY_SKIP", ""))
	}

	if addManifest == "" && removeManifest == "" {
		flag.Usage()
		log.Fatalln("Missing runtime-config manifest to add or remove")
	}

	if dirTarget == "" {
		flag.Usage()
		log.Fatalln("Missing BOSH director URL")
	}

	dirURL, err := url.Parse(dirTarget)
	if err != nil {
		log.Fatalf("Error parsing BOSH dirctor URL: %v\n", err)
	}

	if uaaTarget == "" {
		uaaTarget = fmt.Sprintf("https://%s:8443", dirURL.Hostname())
	}

	if uaaCACert == "" {
		uaaCACert = dirCACert
	}

	uaa, err := bosh.NewUAA(uaaTarget, uaaClientID, uaaClientSecret, uaaCACert)
	if err != nil {
		log.Fatalf("Error instantiating UAA client: %v\n", err)
	}

	logger, err := bosh.NewLogger(boshLogLevel)
	if err != nil {
		log.Fatalf("Error instantiating BOSH logger: %v\n", err)
	}

	director, err := bosh.NewDirector(dirTarget, dirUser, dirPassword, dirCACert, uaa, logger)
	if err != nil {
		log.Fatalf("Error instantiating BOSH client: %v\n", err)
	}

	cfgr := configurator.New(director, logger)

	if addManifest != "" {
		if err := cfgr.Add(runtimeConfigName, addManifest); err != nil {
			log.Fatalf("Error adding runtime config: %v", err)
		}
	}

	if removeManifest != "" {
		if err := cfgr.Remove(runtimeConfigName, removeManifest); err != nil {
			log.Fatalf("Error adding runtime config: %v", err)
		}
	}

	if deploy || deployDryRun {
		dep := deployer.New(director, deployDryRun, logger)
		if err := dep.DeployAllBut(deploySkip); err != nil {
			log.Fatalf("Error deploying deployments: %v", err)
		}
	}
}

func getEnvString(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}

func getEnvBool(key string, def bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}

	return b
}
