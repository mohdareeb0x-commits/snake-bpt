package logic

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
)

func RunCmd(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	return cmd.Run()
}

func AddCobra(cliName string) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	deps := []string{"go", "cobra-cli", "git"}
	for _, dep := range deps {
		if err := CheckDependency(dep); err != nil {
			return err
		}
	}

	err = CheckInternet()
	if err != nil {
		return err
	}

	if err := RunCmd(dir, "go", "mod", "init", cliName); err != nil {
		return errors.New("CLI initialization stopped: CLI might be already initialized")
	}
	pterm.Info.Println("Initialized go module")

	if err := RunCmd(dir, "cobra-cli", "init"); err != nil {
		return errors.New("CLI initialization stopped: CLI might be already initialized")
	}
	pterm.Info.Println("Initialized cobra-cli")

	if err := RunCmd(dir, "go", "mod", "tidy"); err != nil {
		return errors.New("CLI initialization stopped: Unknown Error")
	}

	if err := RunCmd(dir, "git", "init"); err != nil {
		return errors.New("CLI initialization stopped: Unable to intialize git")
	}
	pterm.Info.Println("Initialized git")

	if err := RunCmd(dir, "touch", ".gitignore"); err != nil {
		return errors.New("CLI initialization stopped: Unable to intialize git")
	}
	pterm.Info.Println("Added .gitignore")

	return nil
}

func AddFastAPI(host string, port string) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	deps := []string{"uv", "git"}
	for _, dep := range deps {
		if err := CheckDependency(dep); err != nil {
			return err
		}
	}

	err = CheckInternet()
	if err != nil {
		return err
	}

	if err := RunCmd(dir, "uv", "init", "."); err != nil {
		return errors.New("FastAPI project initialization stopped: Project might be already initialized")
	}
	pterm.Info.Println("Initialized uv project")

	if err := RunCmd(dir, "uv", "add", "fastapi[standard]"); err != nil {
		return errors.New("FastAPI project initialization stopped: Dependencies installation problem")
	}
	pterm.Info.Println("Added fastapi[standard]")

	if err := RunCmd(dir, "uv", "add", "uvicorn[standard]"); err != nil {
		return errors.New("FastAPI project initialization stopped: Dependencies installation problem")
	}
	pterm.Info.Println("Added uvicorn[standard]")

	if err := RunCmd(dir, "mkdir", "app"); err != nil {
		return errors.New("CLI initialization stopped: Unable to create directories")
	}

	if err := RunCmd(dir, "touch", "app/app.py"); err != nil {
		return errors.New("CLI initialization stopped: Unable to create files")
	}
	pterm.Info.Println("Initialized 'app' directory")

	file, err := os.OpenFile("app/app.py", os.O_RDWR, 0666)

	if err != nil {
		return errors.New("CLI initialization stopped: Unable to write files")
	}

	defer file.Close()

	data := "from fastapi import FastAPI\n\napp = FastAPI()\n\n@app.get(\"/\")\nasync def default():\n\treturn {\n\t\t\"health\": \"okay\"\n\t}"

	Wdata := []byte(data)

	_, err = file.Write(Wdata)
	if err != nil {
		return errors.New("CLI initialization stopped: Unable to write files")
	}

	file, err = os.OpenFile("main.py", os.O_RDWR, 0666)

	if err != nil {
		return errors.New("CLI initialization stopped: Unable to write files")
	}

	defer file.Close()

	data = fmt.Sprintf("import uvicorn\n\nif __name__ == \"__main__\":\n\tuvicorn.run(\"app.app:app\", host=\"%s\", port=%s, reload=True)", host, port)

	Wdata = []byte(data)

	_, err = file.Write(Wdata)

	if err != nil {
		return errors.New("CLI initialization stopped: Unable to write files")
	}
	pterm.Info.Println("Added required code")

	return nil
}
