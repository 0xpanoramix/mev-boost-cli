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
		log.Println(viper.GetViper().Get(genesisForkVersionMainnetViperKey))
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

// genesisForkVersionFlags is used to set and configure the genesis fork version.
func genesisForkVersionFlags(v *viper.Viper, f *pflag.FlagSet) {
	// Mainnet
	f.Bool(genesisForkVersionMainnetFlag, false, "help for mainnet flag")
	err := v.BindPFlag(genesisForkVersionMainnetViperKey, f.Lookup(genesisForkVersionMainnetFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionMainnetViperKey, genesisForkVersionMainnetEnv)
	cobra.CheckErr(err)

	// Kiln
	f.Bool(genesisForkVersionKilnFlag, false, "help for kiln flag")
	err = v.BindPFlag(genesisForkVersionKilnViperKey, f.Lookup(genesisForkVersionKilnFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionKilnViperKey, genesisForkVersionKilnEnv)
	cobra.CheckErr(err)

	// Ropsten
	f.Bool(genesisForkVersionRopstenFlag, false, "help for ropsten flag")
	err = v.BindPFlag(genesisForkVersionRopstenViperKey, f.Lookup(genesisForkVersionRopstenFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionRopstenViperKey, genesisForkVersionRopstenEnv)
	cobra.CheckErr(err)

	// Sepolia
	f.Bool(genesisForkVersionSepoliaFlag, false, "help for sepolia flag")
	err = v.BindPFlag(genesisForkVersionSepoliaViperKey, f.Lookup(genesisForkVersionSepoliaFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionSepoliaViperKey, genesisForkVersionSepoliaEnv)
	cobra.CheckErr(err)

	// Custom
	f.String(genesisForkVersionCustomFlag, "", "help for custom flag")
	err = v.BindPFlag(genesisForkVersionCustomViperKey, f.Lookup(genesisForkVersionCustomFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(genesisForkVersionCustomViperKey, genesisForkVersionCustomEnv)
	cobra.CheckErr(err)
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
