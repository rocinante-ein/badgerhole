package main

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/rocinante-ein/badgerhole/internal/interfaces/controllers"
)

func init() {

	// viper initialize section
	// default server setting
	viper.SetDefault("ServerPort", 8080)
	viper.SetDefault("ServerName", "Apache")
	viper.SetDefault("LogOutputAdapterType", "file")
	viper.SetDefault("LogOutputAdapterConnection", "/opt/badgerhole/logs/badgerhole_%s_%s.json")

	// badger hole web contents directorey
	viper.SetDefault("TemplateDir", "/opt/badgerhole/web/template/")
	viper.SetDefault("StaticWebDir", "/opt/badgerhole/web/static/")

	// pflag config
	pflag.String("ConfigDir", "/opt/badgerhole/configs/", "ConfigDir is badgerhole config directory.")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// viper config load setting
	viper.SetConfigType("yaml")

	// load yaml config section
	viper.AddConfigPath(viper.GetString("ConfigDir"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func main() {

	// http routing section
	http.Handle("/static/", http.FileServer(http.Dir(viper.GetString("StaticWebDir"))))
	http.Handle("/favicon.ico", http.FileServer(
		http.Dir(path.Join(viper.GetString("StaticWebDir"), "favicon.ico"))))
	http.HandleFunc("/", controllers.NewLoggingController().HandleFunc)

	// port normalize section
	if viper.GetInt("ServerPort") <= 1024 || viper.GetInt("ServerPort") > 65535 {
		panic(fmt.Errorf("ServerPort is out of range error. :%d", viper.GetInt("ServerPort")))
	}

	// run badger hole server
	http.ListenAndServe(":"+strconv.Itoa(viper.GetInt("ServerPort")), nil)
}
