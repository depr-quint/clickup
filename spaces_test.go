package clickup

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestSpace(t *testing.T) {
	team := getTestTeam()
	if team == nil {
		t.Fail()
		return
	}

	var space *Space
	t.Run("create", func(t *testing.T) {
		space = testCreateSpace(team, t)
	})
	t.Run("get all", func(t *testing.T) {
		testGetAllSpaces(team, t)
	})
	t.Run("update", func(t *testing.T) {
		space = testUpdateSpace(space, t)
	})
	t.Run("get", func(t *testing.T) {
		space = testGetSpace(space, t)
	})
	t.Run("delete", func(t *testing.T) {
		testDeleteSpace(space, t)
	})
	t.Run("get deleted", func(t *testing.T) {
		testGetDeletedSpace(space, t)
	})
}

func testCreateSpace(team *Team, t *testing.T) *Space {
	name := fmt.Sprintf("Space Test (%s)", time.Now().Format("2006-02-01"))
	space, _, err := testClient.Spaces.Create(team.GetID(), &Space{
		Name: &name,
	})
	if err != nil {
		t.Error(err)
		return nil
	}

	if space.GetName() != name {
		fmt.Errorf("names do not match, got %v, expected %v", space.GetName(), name)
	}
	return space
}

func testGetAllSpaces(team *Team, t *testing.T) {
	resp, _, err := testClient.Spaces.GetAll(team.GetID(), false)
	if err != nil {
		t.Error(err)
		return
	}

	if len(resp.Spaces) < 1 {
		t.Errorf("expected at least one space, got %d", len(resp.Spaces))
	}
}

func testUpdateSpace(original *Space, t *testing.T) *Space {
	name := fmt.Sprintf("Space Test (%s) Updated", time.Now().Format("2006-02-01"))
	space, _, err := testClient.Spaces.Update(original.GetID(), &Space{
		Name: &name,
	})
	if err != nil {
		t.Error(err)
		return nil
	}

	if space.GetID() != original.GetID() {
		t.Errorf("ids do not match, got %v expected %v", space.GetID(), original.GetID())
	}

	if space.GetName() != name {
		fmt.Errorf("names do not match, got %v, expected %v", space.GetName(), name)
	}
	return space
}

func testGetSpace(original *Space, t *testing.T) *Space {
	space, _, err := testClient.Spaces.Get(original.GetID())
	if err != nil {
		t.Error(err)
		return nil
	}

	if space.GetID() != original.GetID() {
		t.Errorf("ids do not match, got %v expected %v", space.GetID(), original.GetID())
	}

	if space.GetName() != original.GetName() {
		fmt.Errorf("names do not match, got %v, expected %v", space.GetName(), original.GetName())
	}
	return space
}

func testDeleteSpace(space *Space, t *testing.T) {
	_, err := testClient.Spaces.Delete(space.GetID())
	if err != nil {
		t.Error(err)
	}
}

func testGetDeletedSpace(space *Space, t *testing.T) {
	space, resp, err := testClient.Spaces.Get(space.GetID())
	if err == nil {
		t.Error("expected error")
	}

	if resp == nil {
		t.Error("expected response")
		return
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("exprected 404 error code, got %d", resp.StatusCode)
	}
}
