package tests

import (
	"bytes"
	"encoding/json"
	"github.com/muratovdias/url-shortner/src/application"
	v1 "github.com/muratovdias/url-shortner/src/server/http/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestComponent(t *testing.T) {
	app, err := application.Init()
	if err != nil {
		t.Fatal(err)
	}

	go app.Run()

	// wait for http server
	waitForHttpServer(t)

	payload := v1.UrlShortenerRequest{
		Url: "https://www.youtube.com",
	}

	// save url and generate alias
	response := saveUrlHappyPath(t, payload)
	// redirect
	redirectToOriginalHappyPath(t, response.Alias)
	// check url stats
	urlStatsHappyPath(t, response.Alias)
	// url list
	urlsListHappyPath(t, response)
	// delete link
	deleteShortLink(t, response.Alias)

	defer removeTestFile(t, "url.db")
}

func waitForHttpServer(t *testing.T) {
	t.Helper()
	<-time.After(time.Second * 2)

	require.EventuallyWithT(
		t,
		func(t *assert.CollectT) {
			resp, err := http.Get("http://localhost:8080/health")
			if !assert.NoError(t, err) {
				return
			}
			defer resp.Body.Close()

			if assert.Less(t, resp.StatusCode, 300, "API not ready, http status: %d", resp.StatusCode) {
				return
			}
		},
		time.Second*10,
		time.Millisecond*50,
	)
}

func saveUrlHappyPath(t *testing.T, req v1.UrlShortenerRequest) v1.UrlShortenerResponse {
	t.Helper()

	payload, err := json.Marshal(req)
	require.NoError(t, err)

	httpReq, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:8080/api/v1/shortener",
		bytes.NewBuffer(payload),
	)
	require.NoError(t, err)

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	defer resp.Body.Close()

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var shortener v1.UrlShortenerResponse
	if err = json.NewDecoder(resp.Body).Decode(&shortener); err != nil {
		t.Fatal(err)
	}

	return shortener
}

func redirectToOriginalHappyPath(t *testing.T, alias string) {
	t.Helper()

	// Создаем HTTP-запрос на GET /api/v1/{link}
	httpReq, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/api/v1/"+alias,
		nil,
	)
	require.NoError(t, err)

	// Выполняем запрос
	resp, err := http.DefaultClient.Do(httpReq)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func removeTestFile(t *testing.T, filePath string) {
	t.Helper()
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		t.Fatalf("failed to remove test file %s: %v", filePath, err)
	}
}

func urlStatsHappyPath(t *testing.T, alias string) {
	t.Helper()

	httpReq, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/api/v1/stats/"+alias,
		nil,
	)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(httpReq)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var stats v1.UrlStatsResponse
	if err = json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		t.Fatal(err)
	}

	require.GreaterOrEqual(t, stats.Clicks, 0, "Clicks must be greater or equal to 0")
	require.NotZero(t, stats.LastAccess, "LastAccess time must not be zero")
}

func urlsListHappyPath(t *testing.T, actual v1.UrlShortenerResponse) {
	t.Helper()

	httpReq, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/api/v1/shortener",
		nil,
	)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(httpReq)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var links []v1.UrlsListResponse
	if err = json.NewDecoder(resp.Body).Decode(&links); err != nil {
		t.Fatal(err)
	}

	require.GreaterOrEqual(t, len(links), 1, "Должно быть как минимум 1 ссылки в ответе")

	for _, link := range links {
		require.Equal(t, link.Alias, actual.Alias)
		require.Equal(t, link.Expires, actual.ExpireTime)
	}
}

func deleteShortLink(t *testing.T, alias string) {
	t.Helper()

	deleteReq, err := http.NewRequest(
		http.MethodDelete,
		"http://localhost:8080/api/v1/"+alias,
		nil,
	)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(deleteReq)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
