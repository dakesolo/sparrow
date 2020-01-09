package share

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Run(mux *http.ServeMux) {
	config := viper.New()
	config.SetConfigName("server")
	config.AddConfigPath(Path.Config)

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(config.GetString("http.host") + ":" + config.GetString("http.port"))
	err := http.ListenAndServe(config.GetString("http.host")+":"+config.GetString("http.port"), mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
