package logic

import (
	"os"
	"strings"
	"testing"
)

func useTempDir(t *testing.T) {
	t.Helper()

	original, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	tmp := t.TempDir()
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.Chdir(original)
	})
}

func TestAddCobra_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	useTempDir(t)

	err := AddCobra("my-test-cli")
	if err != nil {
		t.Fatalf("AddCobra failed: %v", err)
	}

	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		t.Error("go.mod was not created")
	}

	if _, err := os.Stat(".gitignore"); os.IsNotExist(err) {
		t.Error(".gitignore was not created")
	}

}

func TestAddFastAPI_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	useTempDir(t)

	err := AddFastAPI("0.0.0.0", "8000")
	if err != nil {
		t.Fatalf("AddFastAPI failed: %v", err)
	}

	if _, err := os.Stat("app/app.py"); os.IsNotExist(err) {
		t.Error("app/app.py was not created")
	}

	if _, err := os.Stat("main.py"); os.IsNotExist(err) {
		t.Error("main.py was not created")
	}

	mainContent, err := os.ReadFile("main.py")
	if err != nil {
		t.Fatalf("could not read main.py: %v", err)
	}
	if !strings.Contains(string(mainContent), "0.0.0.0") {
		t.Error("main.py missing host")
	}
	if !strings.Contains(string(mainContent), "8000") {
		t.Error("main.py missing port")
	}

	appContent, err := os.ReadFile("app/app.py")
	if err != nil {
		t.Fatalf("could not read app/app.py: %v", err)
	}
	if !strings.Contains(string(appContent), "from fastapi import FastAPI") {
		t.Error("app/app.py missing FastAPI import")
	}
	
	if _, err := os.Stat("pyproject.toml"); os.IsNotExist(err) {
		t.Error("pyproject.toml was not created")
	}
}