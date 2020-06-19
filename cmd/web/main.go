package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/DiaElectronics/online_kasse/cmd/web/api"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/DiaElectronics/online_kasse/cmd/web/dal"
	"github.com/DiaElectronics/online_kasse/cmd/web/device"
	"github.com/DiaElectronics/online_kasse/cmd/web/memdb"
	"github.com/powerman/structlog"
)

var (
	log         = structlog.New()
	cfg         = getConfig()
	errNoDevice = errors.New("no device found")
)

func run(errc chan<- error) {
	time.Sleep(time.Second * 10)

	var mutex sync.Mutex
	var db app.DataAccessLayer
	db, err := dal.NewPostgresDAL(cfg)
	if err != nil {
		if cfg.Host == "" && cfg.User == "" && cfg.Password == "" {
			log.PrintErr(err)
			log.Info("USING MEM DB")
			db = memdb.New()
		} else {
			errc <- err
			return
		}
	}

	dev, err := device.NewKaznacheyFA(&mutex)
	if err != nil {
		errc <- err
		return
	}

	if dev == nil {
		log.PrintErr("can't find a printer, just demo mode", err, errNoDevice)
	}

	err = dev.PingDevice()
	if err != nil {
		log.PrintErr("can't find a printer, running a demo mode", err, errNoDevice)
		dev = nil
	}

	application := app.NewApplication(db, dev, errc)
	server := api.NewWebServer(application)

	server.Start(errc)
}

func main() {
	log.Info("Server is preparing to start ...", "+")

	errc := make(chan error)
	go run(errc)
	for {
		err := <-errc
		if err != nil {
			log.PrintErr("ERR_URG", "err", err)
		}
	}
}

func getConfig() dal.Config {
	cfg := dal.Config{}
	flag.StringVar(&cfg.User, "dbuser", "", "db user")
	flag.StringVar(&cfg.Password, "dbpass", "", "db pass")
	flag.StringVar(&cfg.Host, "dbhost", "", "db host [ADDR]:PORT")
	flag.Usage = flag.PrintDefaults
	flag.Parse()
	return cfg
}
