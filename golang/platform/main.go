package main

import (
	"platform/config"
	"platform/logging"
    "platform/services"
)


func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		message, ok := section.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}

func main() {
	// var cfg config.Configuration
	// var err error
	// cfg, err = config.Load("config/config.json")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// var logger logging.Logger = logging.NewDefaultLogger(cfg)
	// writeMessage(logger, cfg)

    // services.RegisterDefaultServices()
    //
    // var cfg config.Configuration
    // services.Getservice(&cfg)
    //
    // var logger logging.Logger
    // services.Getservice(&logger)
    //
    // writeMessage(logger, cfg)

    services.RegisterDefaultServices()
    services.Call(writeMessage)

    val := struct {
        message string
        logging.Logger
    }{
        message: "Hello from the struct",
    }
    services.Populate(&val)
    val.Logger.Debug(val.message)
}
