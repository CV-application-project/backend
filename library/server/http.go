package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type httpConfig struct {
	Addr              Listen
	MuxOptions        []runtime.ServeMuxOption
	ServerMiddlewares []HTTPServerMiddleware
	ServerHandlers    []HTTPServerHandler
}

func createDefaultHttpConfig() *httpConfig {
	return &httpConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 10080,
		},
		MuxOptions: []runtime.ServeMuxOption{
			ProtoJSONMarshaller(),
			runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
		},
	}
}

func ProtoJSONMarshaller() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true})
}

type httpServer struct {
	server *http.Server
	config *httpConfig
}

func newGatewayServer(c *httpConfig, conn *grpc.ClientConn, servers []ServiceServer) (*httpServer, error) {
	// init mux
	mux := runtime.NewServeMux(c.MuxOptions...)
	var handler http.Handler = mux

	handler = otelhttp.NewHandler(handler, "")

	for i := len(c.ServerMiddlewares) - 1; i >= 0; i-- {
		handler = c.ServerMiddlewares[i](handler)
	}

	httpMux := http.NewServeMux()

	for _, h := range c.ServerHandlers {
		h(httpMux)
	}

	httpMux.Handle("/", handler)

	svr := &http.Server{
		Addr:    c.Addr.String(),
		Handler: httpMux,
	}

	for _, svr := range servers {
		err := svr.RegisterWithHandler(context.Background(), mux, conn)
		if err != nil {
			return nil, fmt.Errorf("failed to register handler. %w", err)
		}
	}

	return &httpServer{
		server: svr,
		config: c,
	}, nil
}

func (s *httpServer) Serve() error {
	log.Println("http server starting at", s.config.Addr.String())
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("Error starting http server, ", err)
		return err
	}

	return nil
}

func (s *httpServer) Shutdown(ctx context.Context) {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	err := s.server.Shutdown(ctx)
	log.Println("All http(s) requests finished")
	if err != nil {
		log.Println("failed to shutdown http-gateway server: ", err)
	}
}
