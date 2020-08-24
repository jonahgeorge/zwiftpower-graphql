package zwiftpower

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var zwiftPowerURL = "https://zwiftpower.com"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type LoggingHTTPClient struct {
	HTTPClient
}

func (c *LoggingHTTPClient) Do(req *http.Request) (*http.Response, error) {
	log.Printf("--> %s %s", req.Method, req.URL.EscapedPath()+"?"+req.URL.RawQuery)
	return c.HTTPClient.Do(req)
}

type Client struct {
	HTTPClient
}

func (c Client) ListLeagues(
	ctx context.Context,
	req *ListLeaguesRequest,
) (*ListLeaguesResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodGet,
		zwiftPowerURL+"/api3.php?do=league_list",
		bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := new(ListLeaguesResponse)
	err = json.Unmarshal(buf, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c Client) ListTeams(
	ctx context.Context,
	req *ListTeamsRequest,
) (*ListTeamsResponse, error) {
	r, err := http.NewRequest(http.MethodGet,
		zwiftPowerURL+"/api3.php?do=team_list",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := new(ListTeamsResponse)
	err = json.Unmarshal(buf, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c Client) ListActivities(
	ctx context.Context,
	req *ListActivitiesRequest,
) (*ListActivitiesResponse, error) {
	r, err := http.NewRequest(http.MethodGet,
		zwiftPowerURL+"/api3.php?do=activities",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := new(ListActivitiesResponse)
	err = json.Unmarshal(buf, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c Client) GetEventSprints(
	ctx context.Context,
	req *GetEventSprintsRequest,
) (*GetEventSprintsResponse, error) {
	r, err := http.NewRequest(http.MethodGet,
		zwiftPowerURL+"/api3.php?do=event_sprints?zid="+string(req.Zid),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := new(GetEventSprintsResponse)
	err = json.Unmarshal(buf, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c Client) ListEvents(
	ctx context.Context,
	req *ListEventsRequest,
) (*ListEventsResponse, error) {
	r, err := http.NewRequest(http.MethodGet,
		zwiftPowerURL+"/cache3/lists/0_zwift_event_list_3.json",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := new(ListEventsResponse)
	err = json.Unmarshal(buf, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c Client) GetTeamRiders(
	ctx context.Context,
	req *GetTeamRidersRequest,
) (*GetTeamRidersResponse, error) {
	r, err := http.NewRequest(http.MethodGet,
		zwiftPowerURL+"/cache3/teams/"+req.Id+"_riders.json",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := new(GetTeamRidersResponse)
	err = json.Unmarshal(buf, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// /cache3/global/rider_records_" +ZP_VARS.rider_records_option + ".json" } ;
// /cache3/profile/197816_all.json?_=1591419016055
// /cache3/profile/197816_rider_compare_victims.json
// /cache3/profile/197816_analysis_list.json
// /cache3/profile/197816_signups.json
// /api3.php?do=event_comments

