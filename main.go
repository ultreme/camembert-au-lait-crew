package main // import "ultre.me/calcbiz"

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gogo/gateway"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"ultre.me/calcbiz/api"
	"ultre.me/calcbiz/svc"
)

// VERSION represents the version of the Camembert au lait crew's website
const VERSION = "2.1.0"

func main() {
	rand.Seed(time.Now().UnixNano())

	app := cli.NewApp()
	app.Name = "calcbiz"
	app.Usage = "Camembert au lait crew's web server"
	app.Version = VERSION
	app.Flags = []cli.Flag{}

	app.Before = func(c *cli.Context) error {
		config := zap.NewDevelopmentConfig()
		config.Level.SetLevel(zap.DebugLevel)
		config.DisableStacktrace = true
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		l, err := config.Build()
		if err != nil {
			return errors.Wrap(err, "failed to configure logger")
		}
		zap.ReplaceGlobals(l)
		zap.L().Debug("logger initialized")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:   "server",
			Usage:  "Start a calc-www server",
			Action: server,
			Flags: []cli.Flag{
				// server options
				cli.StringFlag{
					Name:  "http-bind",
					Usage: "TCP port address for HTTP server",
					Value: ":9000",
				},
				cli.StringFlag{
					Name:  "grpc-bind",
					Usage: "TCP port address for gRPC server",
					Value: ":9001",
				},
				// service options
				cli.StringFlag{
					Name:   "soundcloud-client-id",
					Value:  "<configure-me>",
					Usage:  "SoundCloud CLIENT_ID",
					EnvVar: "SOUNDCLOUD_CLIENT_ID",
				},
				cli.IntFlag{
					Name:   "soundcloud-user-id",
					Value:  96137699,
					Usage:  "SoundCloud USER_ID",
					EnvVar: "SOUNDCLOUD_USER_ID",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type serverOptions struct {
	// SQLpath        string
	// JWTKey         string
	// WithReflection bool
	GRPCBind       string
	HTTPBind       string
	ServiceOptions svc.Options
}

func serverOptionsFromCliContext(c *cli.Context) serverOptions {
	return serverOptions{
		GRPCBind: c.String("grpc-bind"),
		HTTPBind: c.String("http-bind"),
		ServiceOptions: svc.Options{
			SoundcloudUserID:   c.Int("soundcloud-client-id"),
			SoundcloudClientID: c.String("soundcloud-user-id"),
		},
	}
}

func server(c *cli.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opts := serverOptionsFromCliContext(c)

	errs := make(chan error)
	go func() { errs <- errors.Wrap(startGRPCServer(ctx, &opts), "gRPC server error") }()
	go func() { errs <- errors.Wrap(startHTTPServer(ctx, &opts), "HTTP server error") }()
	return <-errs
}

func startHTTPServer(ctx context.Context, opts *serverOptions) error {
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &gateway.JSONPb{
			EmitDefaults: false,
			Indent:       "  ",
			OrigName:     true,
		}),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	grpcOpts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterServerHandlerFromEndpoint(ctx, gwmux, opts.GRPCBind, grpcOpts); err != nil {
		return err
	}
	zap.L().Info("starting HTTP server", zap.String("bind", opts.HTTPBind))
	mux := http.NewServeMux()
	mux.Handle("/api/", gwmux)
	return http.ListenAndServe(opts.HTTPBind, mux)
}

func startGRPCServer(ctx context.Context, opts *serverOptions) error {
	listener, err := net.Listen("tcp", opts.GRPCBind)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	defer func() {
		if err := listener.Close(); err != nil {
			zap.L().Error(
				"failed to close listener",
				zap.String("address", opts.GRPCBind),
				zap.Error(err),
			)
		}
	}()

	grpcLogger := zap.L().Named("grpc")
	serverStreamOpts := []grpc.StreamServerInterceptor{
		grpc_recovery.StreamServerInterceptor(),
		//grpc_auth.StreamServerInterceptor(authFunc),
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_zap.StreamServerInterceptor(grpcLogger),
		grpc_recovery.StreamServerInterceptor(),
	}
	serverUnaryOpts := []grpc.UnaryServerInterceptor{
		grpc_recovery.UnaryServerInterceptor(),
		//grpc_auth.UnaryServerInterceptor(authFunc),
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(grpcLogger),
		grpc_recovery.UnaryServerInterceptor(),
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(serverStreamOpts...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(serverUnaryOpts...)),
	)

	svc, err := svc.New(opts.ServiceOptions)
	if err != nil {
		return errors.Wrap(err, "failed to initialize service")
	}
	api.RegisterServerServer(grpcServer, svc)
	//if opts.WithReflection {
	reflection.Register(grpcServer)
	//}

	go func() {
		defer grpcServer.GracefulStop()
		<-ctx.Done()
	}()

	zap.L().Info("starting gRPC server", zap.String("bind", opts.GRPCBind))
	return grpcServer.Serve(listener)
}

/*

	r := chi.NewRouter()
	//r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(5 * time.Second))


	})

	// static files
	box := packr.NewBox("./static")
	r.Handle("/", http.FileServer(box))

	// FIXME: handle socket.io
	http.Handle("/", r)
	log.Infof("Listening and serving HTTP on %s", c.String("bind-address"))
	return http.ListenAndServe(c.String("bind-address"), nil)
}
*/
