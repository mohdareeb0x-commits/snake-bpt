package logic

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
)

func AddCobra(cliName string) error {
	_, err := exec.Command("go", "mod", "init", cliName).Output()
	if err != nil {
		return errors.New("CLI initialization stopped: CLI might be already initialized")
	}

	_, err = exec.Command("cobra-cli", "init").Output()
	if err != nil {
		return errors.New("CLI initialization stopped: CLI might be already initialized")
	}

	_, err = exec.Command("go", "mod", "tidy").Output()
	if err != nil {
		return errors.New("CLI initialization stopped: Unknown Error")
	}

	_, err = exec.Command("git", "init").Output()
	if err != nil {
		return errors.New("CLI initialization stopped: Unable to intialize git")
	}

	_, err = exec.Command("touch", ".gitignore").Output()
	if err != nil {
		return errors.New("CLI initialization stopped: Unable to intialize git")
	}

	return nil
}

func AddFastAPI(host string, port string) error {
	_, err := exec.Command("uv", "init", ".").Output()
	if err != nil {
		return errors.New("FastAPI project initialization stopped: Project might be already initialized")
	}
	pterm.Info.Println("Initialized uv project")

	_, err = exec.Command("uv", "add", "fastapi[standard]").Output()
	if err != nil {
		return errors.New("FastAPI project initialization stopped: Dependencies installation problem")
	}
	pterm.Info.Println("Added fastapi[standard]")

	_, err = exec.Command("uv", "add", "uvicorn[standard]").Output()
	if err != nil {
		return errors.New("FastAPI project initialization stopped: Dependencies installation problem")
	}
	pterm.Info.Println("Added uvicorn[standard]")

	_, err = exec.Command("mkdir", "app").Output()
	if err != nil {
		return errors.New("CLI initialization stopped: Unable to create directories")
	}
	
	_, err = exec.Command("touch", "app/app.py").Output()
	if err != nil {
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

