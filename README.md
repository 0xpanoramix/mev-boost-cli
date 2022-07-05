# mev-boost-cli

Use as a playground to implement a CLI for mev-boost, without distractions.

## How does it work ?

Well it's a CLI.

## Getting started !

### Installation

```shell
# Probably a bash command here
```

### Quickstart

To add new flags to the root command, open `cmd/root.go`.
The file's structure is the following:
```go
package cmd

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// High level initialization (config and flags).	
}

func initConfig() {
	// Viper methods to read the configuration file (any format).
}

const (
	abcdFlag = "ab-cd"
	abcdViperKey = "ab.cd"
	abcdEnv = "BOOST_AB_CD"
)

func abcdFlags(v *viper.Viper, f *pflag.FlagSet) {
	f.Bool(abcdFlag, false, "help for --ab-cd flag")
	err := v.BindPFlag(abcdViperKey, f.Lookup(abcdFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(abcdViperKey, abcdEnv)
	cobra.CheckErr(err)
} 
```

You'll just have to call a new method in the `init()` method and declare the function at the end 
of the file.
Right above the function, add the flag, the viper key (used to search in the YAML or JSON 
configuration files) and the environment representation.

For this last one, prefix with `BOOST_` to avoid collision with already existing variables.

In the method itself, you should only have to update the `f.<TYPE>()` method, in order to update 
the default value for example.

## Author

Made with ❤️ by 🤖 [Luca Georges François](https://github.com/0xpanoramix) 🤖
