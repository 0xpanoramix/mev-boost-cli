package cmd

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd represents the base command of mev-boost.
var rootCmd = &cobra.Command{
	Use:   "mev-boost",
	Short: "A middleware used by PoS Ethereum consensus clients to outsource block construction.",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			for _, key := range viper.GetViper().AllKeys() {
				log.Printf("%s: %v", key, viper.GetViper().Get(key))
			}
		*/
		config, err := boostConfigFromViper()
		if err != nil {
			log.WithError(err).Error("could not start mev-boost")
		}
		log.Printf("Configuration: %+v", config)
	},
}

// Execute runs the base command parser.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVar(&cfgFile, "config", "", "help for config flag")

	// Register genesis fork version flags.
	genesisForkVersionFlags(viper.GetViper(), rootCmd.PersistentFlags())
	// Register log flags.
	logFlags(viper.GetViper(), rootCmd.PersistentFlags())
	// Register relay flags.
	relayFlags(viper.GetViper(), rootCmd.PersistentFlags())
	// Register server flags.
	serverFlags(viper.GetViper(), rootCmd.PersistentFlags())
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".mev-boost" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mev-boost")
	}

	viper.SetEnvPrefix("BOOST")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

type boostConfig struct {
	genesisForkVersionHex string
	serverAddr            string
}

func boostConfigFromViper() (*boostConfig, error) {
	config := &boostConfig{}

	// Tries to set the genesis fork version.
	if viper.GetViper().GetBool(genesisForkVersionMainnetViperKey) {
		config.genesisForkVersionHex = genesisForkVersionMainnet
	} else if viper.GetViper().GetBool(genesisForkVersionKilnViperKey) {
		config.genesisForkVersionHex = genesisForkVersionKiln
	} else if viper.GetViper().GetBool(genesisForkVersionRopstenViperKey) {
		config.genesisForkVersionHex = genesisForkVersionRopsten
	} else if viper.GetViper().GetBool(genesisForkVersionSepoliaViperKey) {
		config.genesisForkVersionHex = genesisForkVersionSepolia
	} else if viper.GetViper().GetString(genesisForkVersionCustomViperKey) != "" {
		config.genesisForkVersionHex = viper.GetViper().GetString(genesisForkVersionCustomViperKey)
	} else {
		return nil, errors.New("invalid genesis fork version")
	}

	// Sets the server listening address.
	config.serverAddr = viper.GetViper().GetString(serverAddrViperKey)

	// Verifies the relay's format.
	// parseRelayURLs(viper.GetViper().GetStringSlice(relayURLsViperKey))
	return config, nil
}

const (
	genesisForkVersionMainnet         = "0x00000000"
	genesisForkVersionMainnetFlag     = "gfv-mainnet"
	genesisForkVersionMainnetViperKey = "gfv.mainnet"
	genesisForkVersionMainnetEnv      = "BOOST_GENESIS_FORK_VERSION_MAINNET"

	genesisForkVersionKiln         = "0x70000069"
	genesisForkVersionKilnFlag     = "gfv-kiln"
	genesisForkVersionKilnViperKey = "gfv.kiln"
	genesisForkVersionKilnEnv      = "BOOST_GENESIS_FORK_VERSION_KILN"

	genesisForkVersionRopsten         = "0x80000069"
	genesisForkVersionRopstenFlag     = "gfv-ropsten"
	genesisForkVersionRopstenViperKey = "gfv.ropsten"
	genesisForkVersionRopstenEnv      = "BOOST_GENESIS_FORK_VERSION_ROPSTEN"

	genesisForkVersionSepolia         = "0x90000069"
	genesisForkVersionSepoliaFlag     = "gfv-sepolia"
	genesisForkVersionSepoliaViperKey = "gfv.sepolia"
	genesisForkVersionSepoliaEnv      = "BOOST_GENESIS_FORK_VERSION_SEPOLIA"

	genesisForkVersionCustomFlag     = "gfv-custom"
	genesisForkVersionCustomViperKey = "gfv.custom"
	genesisForkVersionCustomEnv      = "BOOST_GENESIS_FORK_VERSION_CUSTOM"
)

// genesisForkVersionFlags is used to register and configure the genesis fork version.
func genesisForkVersionFlags(v *viper.Viper, f *pflag.FlagSet) {
	// --gfv-mainnet
	f.Bool(genesisForkVersionMainnetFlag, false, "help for --gfv-mainnet flag")
	err := v.BindPFlag(genesisForkVersionMainnetViperKey, f.Lookup(genesisForkVersionMainnetFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionMainnetViperKey, genesisForkVersionMainnetEnv)
	cobra.CheckErr(err)

	// --gfv-kiln
	f.Bool(genesisForkVersionKilnFlag, false, "help for --gfv-kiln flag")
	err = v.BindPFlag(genesisForkVersionKilnViperKey, f.Lookup(genesisForkVersionKilnFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionKilnViperKey, genesisForkVersionKilnEnv)
	cobra.CheckErr(err)

	// --gfv-ropsten
	f.Bool(genesisForkVersionRopstenFlag, false, "help for --gfv-ropsten flag")
	err = v.BindPFlag(genesisForkVersionRopstenViperKey, f.Lookup(genesisForkVersionRopstenFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionRopstenViperKey, genesisForkVersionRopstenEnv)
	cobra.CheckErr(err)

	// --gfv-sepolia
	f.Bool(genesisForkVersionSepoliaFlag, false, "help for --gfv-sepolia flag")
	err = v.BindPFlag(genesisForkVersionSepoliaViperKey, f.Lookup(genesisForkVersionSepoliaFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionSepoliaViperKey, genesisForkVersionSepoliaEnv)
	cobra.CheckErr(err)

	// --gfv-custom
	f.String(genesisForkVersionCustomFlag, "", "help for --gfv-custom flag")
	err = v.BindPFlag(genesisForkVersionCustomViperKey, f.Lookup(genesisForkVersionCustomFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionCustomViperKey, genesisForkVersionCustomEnv)
	cobra.CheckErr(err)
}

const (
	logJSONFlag     = "log-json"
	logJSONViperKey = "log.json"
	logJSONEnv      = "BOOST_LOG_JSON"

	logLevelFlag     = "log-level"
	logLevelViperKey = "log.level"
	logLevelEnv      = "BOOST_LOG_LEVEL"
)

// logFlags is used to register and configure the log parameters.
func logFlags(v *viper.Viper, f *pflag.FlagSet) {
	// --log-json
	f.Bool(logJSONFlag, false, "help for --log-json flag")
	err := v.BindPFlag(logJSONViperKey, f.Lookup(logJSONFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(logJSONViperKey, logJSONEnv)
	cobra.CheckErr(err)

	// --log-level
	f.String(logLevelFlag, "info", "help for --log-level flag")
	err = v.BindPFlag(logLevelViperKey, f.Lookup(logLevelFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(logLevelViperKey, logLevelEnv)
	cobra.CheckErr(err)
}

const (
	relayURLsFlag     = "relay-urls"
	relayURLsViperKey = "relay.urls"
	relayURLsEnv      = "BOOST_RELAY_URLS"

	relayCheckFlag     = "relay-check"
	relayCheckViperKey = "relay.check"
	relayCheckEnv      = "BOOST_RELAY_CHECK"
)

// relayFlags is used to register and configure the relay parameters.
func relayFlags(v *viper.Viper, f *pflag.FlagSet) {
	// --relay-urls
	f.StringArray(relayURLsFlag, []string{}, "help for --relay-urls flag")
	err := v.BindPFlag(relayURLsViperKey, f.Lookup(relayURLsFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(relayURLsViperKey, relayURLsEnv)
	cobra.CheckErr(err)

	// --relay-check
	f.Bool(relayCheckFlag, false, "help for --relay-check flag")
	err = v.BindPFlag(relayCheckViperKey, f.Lookup(relayCheckFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(relayCheckViperKey, relayCheckEnv)
	cobra.CheckErr(err)
}

const (
	serverAddrFlag     = "server-addr"
	serverAddrViperKey = "server.addr"
	serverAddrEnv      = "BOOST_SERVER_ADDR"
)

// serverFlags is used to register and configure the server parameters.
func serverFlags(v *viper.Viper, f *pflag.FlagSet) {
	// --server-addr
	f.String(serverAddrFlag, "localhost:18550", "help for --server-addr flag")
	err := v.BindPFlag(serverAddrViperKey, f.Lookup(serverAddrFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(serverAddrViperKey, serverAddrEnv)
	cobra.CheckErr(err)
}
