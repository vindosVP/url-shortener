package inmemory_repo_test

import (
	"github.com/stretchr/testify/require"
	"github.com/vindosVP/url-shortener/src/internal/cerrors"
	"github.com/vindosVP/url-shortener/src/internal/pkg/logger/discardLogger"
	"github.com/vindosVP/url-shortener/src/internal/usecase/inmemory_repo"
	"testing"
)

func NewTestRepo() *inmemory_repo.InmemoryRepo {
	l := discardLogger.NewDiscardLogger()
	return inmemory_repo.New(l)
}

func TestSave(t *testing.T) {

	const originalUrl = "https://github.com/vindosVP"
	const alias = "AHJDVU89_0"
	repo := NewTestRepo()

	t.Run("Success", func(t *testing.T) {

		res, err := repo.Save(originalUrl, alias)
		require.Equal(t, alias, res)
		require.Equal(t, nil, err)

	})

	t.Run("Already saved", func(t *testing.T) {

		res, err := repo.Save(originalUrl, alias)
		require.Equal(t, "", res)
		require.Equal(t, cerrors.ErrAliasAlreadySaved, err)

	})

}

func TestGetOriginal(t *testing.T) {

	const originalUrl = "https://github.com/vindosVP"
	const alias = "AHJDVU89_0"
	repo := NewTestRepo()

	t.Run("Success", func(t *testing.T) {

		_, err := repo.Save(originalUrl, alias)
		if err != nil {
			t.Fatal(err)
		}
		res, err := repo.GetOriginal(alias)
		require.Equal(t, originalUrl, res)
		require.Equal(t, nil, err)

	})

	repo = NewTestRepo()

	t.Run("Not exist", func(t *testing.T) {

		res, err := repo.GetOriginal(alias)
		require.Equal(t, "", res)
		require.Equal(t, cerrors.ErrAliasDoesNotExist, err)

	})

}

func TestAliasForURLExists(t *testing.T) {

	const originalUrl = "https://github.com/vindosVP"
	const alias = "AHJDVU89_0"
	repo := NewTestRepo()

	t.Run("Exists", func(t *testing.T) {

		_, err := repo.Save(originalUrl, alias)
		if err != nil {
			t.Fatal(err)
		}
		res, err := repo.AliasForURLExists(originalUrl)
		require.Equal(t, true, res)
		require.Equal(t, nil, err)

	})

	repo = NewTestRepo()

	t.Run("Not exist", func(t *testing.T) {

		res, err := repo.AliasForURLExists(originalUrl)
		require.Equal(t, false, res)
		require.Equal(t, nil, err)

	})

}

func TestAliasExists(t *testing.T) {

	const originalUrl = "https://github.com/vindosVP"
	const alias = "AHJDVU89_0"
	repo := NewTestRepo()

	t.Run("Exists", func(t *testing.T) {

		_, err := repo.Save(originalUrl, alias)
		if err != nil {
			t.Fatal(err)
		}
		res, err := repo.AliasExists(alias)
		require.Equal(t, true, res)
		require.Equal(t, nil, err)

	})

	repo = NewTestRepo()

	t.Run("Not exist", func(t *testing.T) {

		res, err := repo.AliasExists(alias)
		require.Equal(t, false, res)
		require.Equal(t, nil, err)

	})

}

func TestGetAlias(t *testing.T) {

	const originalUrl = "https://github.com/vindosVP"
	const alias = "AHJDVU89_0"
	repo := NewTestRepo()

	t.Run("Success", func(t *testing.T) {

		_, err := repo.Save(originalUrl, alias)
		if err != nil {
			t.Fatal(err)
		}
		res, err := repo.GetAlias(originalUrl)
		require.Equal(t, alias, res)
		require.Equal(t, nil, err)

	})

	repo = NewTestRepo()

	t.Run("Not exist", func(t *testing.T) {

		res, err := repo.GetAlias(originalUrl)
		require.Equal(t, "", res)
		require.Equal(t, cerrors.ErrAliasForURLDoesNotExist, err)

	})

}
