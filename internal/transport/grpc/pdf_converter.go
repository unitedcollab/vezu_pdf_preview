package grpc

import (
	"context"

	"github.com/h2non/bimg"
	v1 "github.com/unitedcollab/vezu_pdf_preview/build/grpc/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PDFConverterServiceServer struct{}

var _ v1.PDFConverterServiceServer = (*PDFConverterServiceServer)(nil)

func NewPDFConverterServiceServer() *PDFConverterServiceServer {
	return &PDFConverterServiceServer{}
}

func (p *PDFConverterServiceServer) Register(grpcServer *grpc.Server) {
	v1.RegisterPDFConverterServiceServer(grpcServer, p)
}

func (p *PDFConverterServiceServer) ConvertPDFToJPG(ctx context.Context, request *v1.ConvertPDFToJPGRequest) (*v1.ConvertPDFToJPGResponse, error) {
	jpegImage, err := bimg.
		NewImage(request.GetPdfFileContent()).
		Convert(bimg.JPEG)
	if err != nil {
		return &v1.ConvertPDFToJPGResponse{}, status.Error(codes.Internal, err.Error())
	}

	width := int(request.GetWidth())
	height := int(request.GetHeight())

	resizedJpegImage, err := bimg.
		NewImage(jpegImage).
		Resize(width, height)
	if err != nil {
		return &v1.ConvertPDFToJPGResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &v1.ConvertPDFToJPGResponse{
		JpgFileContent: resizedJpegImage,
	}, nil
}
