package cmd

import (
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
		log.Println(viper.GetViper().Get(logLevelViperKey))
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

const (
	genesisForkVersionMainnetFlag     = "gfv-mainnet"
	genesisForkVersionMainnetViperKey = "gfv.mainnet"
	genesisForkVersionMainnetEnv      = "BOOST_GENESIS_FORK_VERSION_MAINNET"

	genesisForkVersionKilnFlag     = "gfv-kiln"
	genesisForkVersionKilnViperKey = "gfv.kiln"
	genesisForkVersionKilnEnv      = "BOOST_GENESIS_FORK_VERSION_KILN"

	genesisForkVersionRopstenFlag     = "gfv-ropsten"
	genesisForkVersionRopstenViperKey = "gfv.ropsten"
	genesisForkVersionRopstenEnv      = "BOOST_GENESIS_FORK_VERSION_ROPSTEN"

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
	f.String(logLevelFlag, "debug", "help for --log-level flag")
	err = v.BindPFlag(logLevelViperKey, f.Lookup(logLevelFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(logLevelViperKey, logLevelEnv)
	cobra.CheckErr(err)
}
