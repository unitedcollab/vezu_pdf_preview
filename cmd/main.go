package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/unitedcollab/vezu_pdf_preview/internal/platform/config"
	"github.com/unitedcollab/vezu_pdf_preview/internal/platform/server"
	"github.com/unitedcollab/vezu_pdf_preview/internal/transport/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pdfConverterServiceServer := grpc.NewPDFConverterServiceServer()
	grpcServer := server.NewServer()
	pdfConverterServiceServer.Register(grpcServer)

	uri := net.JoinHostPort(cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", uri)
	if err != nil {
		log.Fatal(err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		<-shutdown
		grpcServer.GracefulStop()
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
