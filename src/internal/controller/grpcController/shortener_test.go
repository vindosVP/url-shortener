package grpcController_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/vindosVP/url-shortener/src/internal/cerrors"
	"github.com/vindosVP/url-shortener/src/internal/controller/grpcController"
	"github.com/vindosVP/url-shortener/src/internal/pkg/logger/discardLogger"
	"github.com/vindosVP/url-shortener/src/internal/usecase/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const Pattern = "https:\\/\\/testdomain.com\\/\\b([-a-zA-Z0-9_]{10})$"

func TestGet(t *testing.T) {

	cases := []struct {
		name                 string
		req                  *grpcController.GetRequest
		resp                 *grpcController.GetResponse
		checkResp            bool
		mockPattern          bool
		expErr               error
		mockGetOriginalResp  string
		mockGetOriginalError error
	}{
		{
			name:                 "Success",
			req:                  &grpcController.GetRequest{ShortenedUrl: "https://testdomain.com/AHJDVU89_0"},
			resp:                 &grpcController.GetResponse{Url: "https://github.com/vindosVP"},
			checkResp:            true,
			mockPattern:          true,
			expErr:               nil,
			mockGetOriginalResp:  "https://github.com/vindosVP",
			mockGetOriginalError: nil,
		},
		{
			name:      "Invalid url 1",
			req:       &grpcController.GetRequest{ShortenedUrl: "invalid url"},
			checkResp: false,
			expErr:    status.Error(codes.InvalidArgument, cerrors.ErrInvalidUrl.Error()),
		},
		{
			name:        "Invalid url 2",
			mockPattern: true,
			req:         &grpcController.GetRequest{ShortenedUrl: "https://testdomain.com/AHJDVU89_?"},
			checkResp:   false,
			expErr:      status.Error(codes.InvalidArgument, cerrors.ErrInvalidUrl.Error()),
		},
		{
			name:                 "Alias does not exist",
			mockPattern:          true,
			req:                  &grpcController.GetRequest{ShortenedUrl: "https://testdomain.com/AHJDVU89_0"},
			checkResp:            false,
			mockGetOriginalError: cerrors.ErrAliasForURLDoesNotExist,
			expErr:               status.Error(codes.NotFound, cerrors.ErrAliasForURLDoesNotExist.Error()),
		},
		{
			name:                 "Unexpected error",
			mockPattern:          true,
			req:                  &grpcController.GetRequest{ShortenedUrl: "https://testdomain.com/AHJDVU89_0"},
			checkResp:            false,
			mockGetOriginalError: errors.New("unexpected error"),
			expErr:               status.Error(codes.Internal, "failed to get original url"),
		},
	}

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		_ = lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	l := discardLogger.NewDiscardLogger()
	shortenerMock := mocks.NewShortener(t)
	GrpcShortener := grpcController.NewShortener(shortenerMock, l)
	grpcController.RegisterUrlShortenerServer(srv, GrpcShortener)

	go func() {
		err := srv.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	dialer := func(ctx context.Context, string2 string) (net.Conn, error) {
		return lis.Dial()
	}
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	t.Cleanup(func() {
		_ = conn.Close()
	})

	if err != nil {
		t.Fatal(err)
	}

	client := grpcController.NewUrlShortenerClient(conn)

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.mockPattern {
				shortenerMock.On("GetLinkPattern").Return(Pattern).Once()
			}

			if tc.mockGetOriginalResp != "" || tc.mockGetOriginalError != nil {
				shortenerMock.On("GetOriginal", tc.req.ShortenedUrl[len(tc.req.ShortenedUrl)-10:]).Return(tc.mockGetOriginalResp, tc.mockGetOriginalError).Once()
			}

			resp, err := client.Get(context.Background(), tc.req)
			if tc.checkResp {
				require.Equal(t, tc.resp.Url, resp.Url)
			}
			if tc.expErr != nil {
				require.ErrorIs(t, err, tc.expErr)
			}
		})
	}

}

func TestSave(t *testing.T) {

	cases := []struct {
		name          string
		req           *grpcController.SaveRequest
		resp          *grpcController.SaveResponse
		checkResp     bool
		expErr        error
		mockSaveResp  string
		mockSaveError error
	}{
		{
			name:          "Success",
			req:           &grpcController.SaveRequest{Url: "https://github.com/vindosVP"},
			resp:          &grpcController.SaveResponse{ShortenedUrl: "https://testdomain.com/AHJDVU89_0"},
			checkResp:     true,
			expErr:        nil,
			mockSaveResp:  "https://testdomain.com/AHJDVU89_0",
			mockSaveError: nil,
		},
		{
			name:      "Invalid url",
			req:       &grpcController.SaveRequest{Url: "invalid url"},
			checkResp: false,
			expErr:    status.Error(codes.InvalidArgument, cerrors.ErrInvalidUrl.Error()),
		},
		{
			name:          "Unexpected error",
			req:           &grpcController.SaveRequest{Url: "https://github.com/vindosVP"},
			checkResp:     false,
			expErr:        status.Error(codes.Internal, "failed to save alias"),
			mockSaveError: errors.New("unexpected error"),
		},
	}

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		_ = lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	l := discardLogger.NewDiscardLogger()
	shortenerMock := mocks.NewShortener(t)
	GrpcShortener := grpcController.NewShortener(shortenerMock, l)
	grpcController.RegisterUrlShortenerServer(srv, GrpcShortener)

	go func() {
		err := srv.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	dialer := func(ctx context.Context, string2 string) (net.Conn, error) {
		return lis.Dial()
	}
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	t.Cleanup(func() {
		_ = conn.Close()
	})

	if err != nil {
		t.Fatal(err)
	}

	client := grpcController.NewUrlShortenerClient(conn)

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.mockSaveResp != "" || tc.mockSaveError != nil {
				shortenerMock.On("Save", tc.req.Url).Return(tc.mockSaveResp, tc.mockSaveError).Once()
			}

			resp, err := client.Save(context.Background(), tc.req)
			if tc.checkResp {
				require.Equal(t, tc.resp.ShortenedUrl, resp.ShortenedUrl)
			}
			if tc.expErr != nil {
				require.ErrorIs(t, err, tc.expErr)
			}

		})
	}

}
