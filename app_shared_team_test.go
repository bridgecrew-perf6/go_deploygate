package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListAppSharedTeams(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/app/list_app_shared_teams")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListAppSharedTeams(&ListAppSharedTeamsRequest{
		Organization: "test_organization",
		Platform:     "android",
		AppId:        "com.deploygate.sample",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}

func Test_AddAppSharedTeam(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/app/add_app_shared_team")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	// TODO: Success response is empty and couldn't decode.
	_, err = c.AddAppSharedTeam(&AddAppSharedTeamRequest{
		Organization: "test_organization",
		Platform:     "android",
		AppId:        "com.deploygate.sample",
		Team:         "test_team",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RemoveAppSharedTeam(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/app/remove_app_shared_team")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveAppSharedTeam(&RemoveAppSharedTeamRequest{
		Organization: "test_organization",
		Platform:     "android",
		AppId:        "com.deploygate.sample",
		Team:         "test_team",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}
