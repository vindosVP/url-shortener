package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/stretchr/testify/require"
	"github.com/vindosVP/url-shortener/src/internal/cerrors"
	"github.com/vindosVP/url-shortener/src/internal/controller/http"
	"github.com/vindosVP/url-shortener/src/internal/pkg/logger/discardLogger"
	"github.com/vindosVP/url-shortener/src/internal/usecase/mocks"
	"io"
	"net/http/httptest"
	"testing"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type response struct {
	Status string `json:"status"`
	URL    string `json:"url"`
	Error  string `json:"errors"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
	Pattern     = "https:\\/\\/testdomain.com\\/\\b([-a-zA-Z0-9_]{10})$"
)

func setupMiddlewares(handler *fiber.App) {
	handler.Use(requestid.New(requestid.Config{
		Header:     "X-Request-ID",
		ContextKey: "request-id",
	}))
}

func TestGetOriginalHandler(t *testing.T) {
	cases := []struct {
		name         string
		alias        string
		respError    string
		respCode     int
		respStatus   string
		mockOriginal string
		mockError    error
	}{
		{
			name:         "Success",
			alias:        "https://testdomain.com/AHJDVU89_0",
			respCode:     200,
			respStatus:   StatusOK,
			mockOriginal: "https://github.com/vindosVP",
		},
		{
			name:       "Invalid alias",
			alias:      "https://testdomain.com/AHJDVU89_?",
			respCode:   400,
			respStatus: StatusError,
			respError:  "url is not valid",
		},
		{
			name:       "Internal error",
			alias:      "https://testdomain.com/AHJDVU89_0",
			respCode:   500,
			respStatus: StatusError,
			respError:  "failed to get original url",
			mockError:  errors.New("unexpected error"),
		},
		{
			name:       "Original not found",
			alias:      "https://testdomain.com/AHJDVU89_0",
			respCode:   404,
			respStatus: StatusError,
			respError:  "ulr with this alias does not exist",
			mockError:  cerrors.ErrAliasForURLDoesNotExist,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			app := fiber.New()
			setupMiddlewares(app)
			l := discardLogger.NewDiscardLogger()
			shortenerMock := mocks.NewShortener(t)
			http.SetupRoutes(app, shortenerMock, l)
			shortenerMock.On("GetLinkPattern").Return(Pattern).Once()

			if tc.mockError != nil || tc.mockOriginal != "" {
				shortenerMock.On("GetOriginal", tc.alias[len(tc.alias)-10:]).Return(tc.mockOriginal, tc.mockError).Once()
			}

			body := shortenRequest{URL: tc.alias}
			var b bytes.Buffer
			err := json.NewEncoder(&b).Encode(body)
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest("GET", "/", &b)
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, 1)
			if err != nil {
				t.Fatal(err)
			}

			require.Equal(t, tc.respCode, resp.StatusCode)
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			var tcResp = response{}
			err = json.Unmarshal(respBody, &tcResp)
			if err != nil {
				t.Fatal(err)
			}

			require.Equal(t, tc.respStatus, tcResp.Status)
			require.Equal(t, tc.respError, tcResp.Error)

			if tc.respStatus == StatusOK {
				require.Equal(t, tc.mockOriginal, tcResp.URL)
			}

		})

	}
}

func TestSaveHandler(t *testing.T) {
	cases := []struct {
		name       string
		url        string
		respError  string
		respCode   int
		respStatus string
		mockAlias  string
		mockError  error
	}{
		{
			name:       "Success",
			url:        "https://github.com/vindosVP",
			respCode:   200,
			respStatus: StatusOK,
			mockAlias:  "https://testdomain.com/AHJDVU89_0",
		},
		{
			name:       "Invalid URL",
			url:        "invalid url",
			respCode:   400,
			respStatus: StatusError,
			respError:  "field URL is not a valid URL",
		},
		{
			name:       "Internal error",
			url:        "https://github.com/vindosVP",
			respCode:   500,
			respStatus: StatusError,
			respError:  "failed to save alias",
			mockError:  errors.New("unexpected error"),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			app := fiber.New()
			setupMiddlewares(app)
			l := discardLogger.NewDiscardLogger()
			shortenerMock := mocks.NewShortener(t)
			http.SetupRoutes(app, shortenerMock, l)

			if tc.mockError != nil || tc.mockAlias != "" {
				shortenerMock.On("Save", tc.url).Return(tc.mockAlias, tc.mockError).Once()
			}

			body := shortenRequest{URL: tc.url}
			var b bytes.Buffer
			err := json.NewEncoder(&b).Encode(body)
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest("POST", "/", &b)
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, 1)
			if err != nil {
				t.Fatal(err)
			}

			require.Equal(t, tc.respCode, resp.StatusCode)
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			var tcResp = response{}
			err = json.Unmarshal(respBody, &tcResp)
			if err != nil {
				t.Fatal(err)
			}

			require.Equal(t, tc.respStatus, tcResp.Status)
			require.Equal(t, tc.respError, tcResp.Error)

			if tc.respStatus == StatusOK {
				require.Equal(t, tc.mockAlias, tcResp.URL)
			}

		})

	}
}
