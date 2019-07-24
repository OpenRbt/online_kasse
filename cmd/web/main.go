package main

import (
	"fmt"
	"flag"
	"strconv"
	"reflect"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/powerman/structlog"
)

func main() {
	flag.Parse()

	log = structlog.New()
		
}
