package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/tensuqiuwulu/bupda-reset-otp/config"
	"github.com/tensuqiuwulu/bupda-reset-otp/delivery/controller"
	"github.com/tensuqiuwulu/bupda-reset-otp/helper"
	"github.com/tensuqiuwulu/bupda-reset-otp/repository"
	"github.com/tensuqiuwulu/bupda-reset-otp/service"
)

func main() {
	// init config
	appConfig := config.GetConfig()

	// Set Database Configuration
	dbConnect := repository.NewDatabaseConnection(&appConfig.Database)

	// Set timezone
	location, err := time.LoadLocation(appConfig.Timezone.Timezone)
	time.Local = location
	log.Println("Location:", location, err)

	// Info app server
	log.Println("Server App : ", string(appConfig.Application.Server))

	// Set Logger
	logrusLogger := helper.NewLogger(appConfig.Log)
	log.Println("Logger : ", logrusLogger)

	scheduler := cron.New(cron.WithLocation(location), cron.WithLogger(cron.DefaultLogger))

	otpManagerRepository := repository.NewOtpManagerRepositoryImpl(&appConfig.Database)

	otpManagerService := service.NewOtpManagerServiceImpl(dbConnect, logrusLogger, otpManagerRepository)

	otpManagerController := controller.NewOtpManagerControllerImpl(otpManagerService)

	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	scheduler.AddFunc("* */23 * * *", func() { otpManagerController.UpdatePhoneLimit() })

	// start scheduler
	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	repository.Close(dbConnect)
}
