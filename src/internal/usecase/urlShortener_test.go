package usecase_test

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vindosVP/url-shortener/src/internal/cerrors"
	"github.com/vindosVP/url-shortener/src/internal/pkg/logger/discardLogger"
	"github.com/vindosVP/url-shortener/src/internal/usecase"
	"github.com/vindosVP/url-shortener/src/internal/usecase/mocks"
	"testing"
)

const (
	domain         = "testdomain.com"
	domainProtocol = "https"
)

func TestGetOriginal(t *testing.T) {

	cases := []struct {
		name            string
		alias           string
		originalURL     string
		error           error
		mockOriginalUrl string
		mockAliasExists bool
	}{
		{
			name:            "Original url exists",
			alias:           "AHJDVU89_0",
			originalURL:     "https://github.com/vindosVP",
			mockAliasExists: true,
			mockOriginalUrl: "https://github.com/vindosVP",
		},
		{
			name:            "Original url does not exist",
			alias:           "AHJDVU89_0",
			mockAliasExists: false,
			mockOriginalUrl: "https://github.com/vindosVP",
			error:           cerrors.ErrAliasForURLDoesNotExist,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			l := discardLogger.NewDiscardLogger()
			mockRepo := mocks.NewUrlRepo(t)
			shortener := usecase.NewShortenerUseCase(mockRepo, domain, domainProtocol, l)
			mockRepo.On("AliasExists", tc.alias).Return(tc.mockAliasExists, nil)

			if tc.mockAliasExists {
				mockRepo.On("GetOriginal", tc.alias).Return(tc.mockOriginalUrl, nil)
			}

			gotOriginal, err := shortener.GetOriginal(tc.alias)
			if tc.mockAliasExists {
				require.Equal(t, tc.originalURL, gotOriginal)
			}
			require.Equal(t, tc.error, err)

		})
	}

}

func TestSave(t *testing.T) {
	cases := []struct {
		name              string
		originalUrl       string
		resultAlias       string
		resultError       error
		mockExistingAlias string
		mockAliasExists   bool
		mockSavedAlias    string
		mockError         string
	}{
		{
			name:              "Alias already exists",
			originalUrl:       "https://github.com/vindosVP",
			resultAlias:       "AHJDVU89_0",
			mockExistingAlias: "AHJDVU89_0",
			mockAliasExists:   true,
		},
		{
			name:            "New alias",
			originalUrl:     "https://github.com/vindosVP",
			mockAliasExists: false,
			resultAlias:     "AHJDVU89_0",
			mockSavedAlias:  "AHJDVU89_0",
			resultError:     nil,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := discardLogger.NewDiscardLogger()
			mockRepo := mocks.NewUrlRepo(t)
			shortener := usecase.NewShortenerUseCase(mockRepo, domain, domainProtocol, l)
			mockRepo.On("AliasForURLExists", tc.originalUrl).Return(tc.mockAliasExists, nil)
			if !tc.mockAliasExists {
				mockRepo.On("AliasExists", mock.AnythingOfType("string")).Return(false, nil)
			}
			if tc.mockSavedAlias != "" {
				mockRepo.On("Save", tc.originalUrl, mock.AnythingOfType("string")).Return(tc.mockSavedAlias, nil)
			}

			if tc.mockAliasExists && tc.mockExistingAlias != "" {
				mockRepo.On("GetAlias", tc.originalUrl).Return(tc.mockExistingAlias, nil)
			}

			gotAlias, err := shortener.Save(tc.originalUrl)
			require.Equal(t, fmt.Sprintf("%s://%s/%s", domainProtocol, domain, tc.resultAlias), gotAlias)
			require.Equal(t, tc.resultError, err)

		})
	}

}
