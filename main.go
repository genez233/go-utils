package main

import (
	"github.com/genez233/go-utils/glog"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TEST
	log := glog.New(&glog.Config{
		ServerName:       "glog",
		Version:          "1.0.0",
		ConsoleLog:       true,
		IsUpload:         true,
		LogUrl:           "http://your.domain.com/api/default/%s/_json",
		OpenobserveToken: "token",
	})

	log.Debug("this is a log.", "v1.0.0", true, 2, glog.Config{
		ServerName:       "glog",
		Version:          "1.0.1",
		ConsoleLog:       true,
		IsUpload:         true,
		LogUrl:           "http://your.domain.com/api/default/%s/_json",
		OpenobserveToken: "token",
	})
}
