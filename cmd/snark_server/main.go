package main

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"micro-snark-server/build"
	"micro-snark-server/snark-ffi"
)

var (
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	node := build.GetAppName()
	version := build.GetVersion()
	return kratos.New(
		kratos.ID(node.ID),
		kratos.Name(node.Name),
		kratos.Version(fmt.Sprintf("%s+%s", version.BaseVersion, version.Commit)),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	app := newApp(nil, nil, nil)
	fmt.Printf("%+v", app)
	snark_ffi.SnarkPost()
}
