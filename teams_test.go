package clickup

import (
	"testing"
)

func TestGetTeams(t *testing.T) {
	resp, _, err := testClient.Teams.Get()
	if err != nil {
		t.Error(err)
	}

	var testTeam *Team
	for _, team := range resp.Teams {
		if team.GetName() == "ClickUp API Test" {
			testTeam = &team
		}
	}

	if testTeam == nil {
		t.Fatal("could not find test team (workspace)")
	}

	if testTeam.GetID() == "" {
		t.Error("could not get team id")
	}

	if testTeam.GetColor() == "" {
		t.Error("could not get team color")
	}
}
