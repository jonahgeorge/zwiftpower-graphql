package zwiftpower

import (
	"net/http"
	"os"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c MockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

func TestListLeagues(t *testing.T) {
	zp := Client{&MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			f, err := os.Open("testdata/league_list.json")
			if err != nil {
				return nil, err
			}
			return &http.Response{Body: f}, nil
		},
	}}

	resp, err := zp.ListLeagues(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", resp)
}

func TestListTeams(t *testing.T) {
	zp := Client{&MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			f, err := os.Open("testdata/team_list.json")
			if err != nil {
				return nil, err
			}
			return &http.Response{Body: f}, nil
		},
	}}

	resp, err := zp.ListTeams(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", resp)
}
