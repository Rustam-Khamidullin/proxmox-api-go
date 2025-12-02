package test

import (
	"os"

	testConstant "github.com/Rustam-Khamidullin/proxmox-api-go/test"
)

func SetEnvironmentVariables() {
	os.Setenv("PM_API_URL", testConstant.ApiURL)
	os.Setenv("PM_USER", testConstant.UserID)
	os.Setenv("PM_PASS", testConstant.Password)
}
