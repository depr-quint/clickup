package clickup

import (
	"fmt"
	"testing"
	"time"
)

func TestFolder(t *testing.T) {
	team := getTestTeam()
	if team == nil {
		t.Fail()
		return
	}

	space := createTestSpace(team, "Folder")
	if space == nil {
		t.Fail()
		return
	}
	defer deleteTestSpace(space)

	var folder *Folder
	t.Run("create", func(t *testing.T) {
		folder = testCreateFolder(space, t)
	})
	t.Run("get all", func(t *testing.T) {
		testGetAllFolders(space, t)
	})
	t.Run("update", func(t *testing.T) {
		folder = testUpdateFolder(folder, t)
	})
	t.Run("get", func(t *testing.T) {
		folder = testGetFolder(folder, t)
	})
	t.Run("delete", func(t *testing.T) {
		testDeleteFolder(folder, t)
	})
}

func testCreateFolder(space *Space, t *testing.T) *Folder {
	name := fmt.Sprintf("Folder Test (%s)", time.Now().Format("2006-02-01"))
	folder, _, err := testClient.Folders.Create(space.GetID(), &Folder{
		Name: &name,
	})
	if err != nil {
		t.Error(err)
		return nil
	}

	if folder.GetName() != name {
		t.Errorf("names do not match, got %v, expected %v", folder.GetName(), name)
	}
	return folder
}

func testGetAllFolders(space *Space, t *testing.T) {
	resp, _, err := testClient.Folders.GetAll(space.GetID(), false)
	if err != nil {
		t.Error(err)
		return
	}

	if len(resp.Folders) < 1 {
		t.Errorf("expected at least one space, got %d", len(resp.Folders))
	}
}

func testUpdateFolder(original *Folder, t *testing.T) *Folder {
	name := fmt.Sprintf("Folder Test (%s) Updated", time.Now().Format("2006-02-01"))
	folder, _, err := testClient.Folders.Update(original.GetID(), &Folder{
		Name: &name,
	})
	if err != nil {
		t.Error(err)
		return nil
	}

	if folder.GetID() != original.GetID() {
		t.Errorf("ids do not match, got %v expected %v", folder.GetID(), original.GetID())
	}

	if folder.GetName() != name {
		t.Errorf("names do not match, got %v, expected %v", folder.GetName(), name)
	}
	return folder
}

func testGetFolder(original *Folder, t *testing.T) *Folder {
	folder, _, err := testClient.Folders.Get(original.GetID())
	if err != nil {
		t.Error(err)
		return nil
	}

	if folder.GetID() != original.GetID() {
		t.Errorf("ids do not match, got %v expected %v", folder.GetID(), original.GetID())
	}

	if folder.GetName() != original.GetName() {
		t.Errorf("names do not match, got %v, expected %v", folder.GetName(), original.GetName())
	}
	return folder
}

func testDeleteFolder(folder *Folder, t *testing.T) {
	_, err := testClient.Folders.Delete(folder.GetID())
	if err != nil {
		t.Error(err)
	}
}
