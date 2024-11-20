package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/sarathkumar17/quiz/utils"
)

type Config struct {
	Filename string
	Timeout  time.Duration
	Help     bool
}

func iniTializeFlags() []string {
	f := reflect.TypeOf(Config{})
	flags := make([]string, f.NumField())
	for i := 0; i < f.NumField(); i++ {
		flag := strings.ToLower(f.Field(i).Name)
		flags[i] = flag
	}
	return flags

}

var Flags = iniTializeFlags()

// GetConfig parses the command line flags and returns a Config object.
// It handles the flags: -filename, -timeout, and -help.
// If the -help flag is present, it returns a Config object with Help set to true.
// If there is an error parsing the flags, it returns an error.
func GetConfig() (config Config, err error) {
	config = Config{}
	err = nil

	if len(os.Args) == 1 {
		return
	}

	if os.Args[1] == "-help" {
		config.Help = true
		return
	}
	flagArgs := os.Args[1:]
	argLen := len(flagArgs)
	for i, arg := range flagArgs {
		if checkNullFlagValue(arg, i, argLen) {
			err := fmt.Errorf("invalid flag value: %v", arg)
			return config, err
		}
		if i%2 != 0 {
			continue
		}
		switch arg {
		case "-filename":
			config.Filename = flagArgs[i+1]
		case "-timeout":
			timeout, err_ := utils.ParseTimeDuration(string(flagArgs[i+1]))
			if err_ != nil {
				return config, err_
			}
			config.Timeout = timeout
		case "-help":
			help := true
			config.Help = help
		default:
			err = errors.New("invalid flag")
		}
	}
	return

}

// checkNullFlagValue takes an argument from the command line and
// checks if it is a flag (i.e. a parameter that is passed to a
// command to modify its behavior). If the argument is a flag, it
// checks if the flag is the last argument in the list of arguments
// passed to the command. If it is, it returns true. Otherwise it
// returns false.
func checkNullFlagValue(arg string, position int, argLength int) bool {
	argString := strings.TrimPrefix(arg, "-")
	isFlag := slices.Contains(Flags, argString)
	if isFlag {
		if position+1 == argLength {
			return true
		}
	}
	return false
}
