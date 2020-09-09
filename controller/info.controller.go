package controller

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// InfoController ...
type InfoController struct {
	InstanceID string      `json:"instance_id"`
	Host       string      `json:"host"`
	Path       string      `json:"path"`
	Method     string      `json:"method"`
	RealIP     string      `json:"real_ip"`
	Status     int         `json:"response_http_status_code"`
	Envs       interface{} `json:"envs"`
	K8s        interface{} `json:"k8s"`
}

func includes(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

// GetInfo ...
func (controller InfoController) GetInfo(c echo.Context) error {
	allEnvs := os.Environ()
	envs := map[string]string{}
	k8s := map[string]string{}

	blackList := []string{
		"GOPATH",
		"PWD",
		"MAKEFLAGS",
		"GOLANG_VERSION",
		"MFLAGS",
		"SHLVL",
		"PATH",
		"CGO_ENABLED",
		"MAKE_TERMOUT",
		"MAKE_TERMERR",
		"TERM",
		"MAKELEVEL",
		"HOME",
		"HOSTNAME",
	}

	statusCodes := []string{
		"100", "101", "102",
		"200", "201", "202", "203", "204", "205", "206", "207", "208", "226",
		"300", "301", "302", "303", "304", "305", "306", "307", "308",
		"400", "401", "402", "403", "404", "405", "406", "407", "408", "409", "410",
		"410", "411", "412", "413", "414", "415", "416", "417", "418", "420", "422",
		"423", "424", "425", "426", "428", "429", "431", "444", "449", "450", "451", "499",
		"500", "501", "502", "503", "504", "505", "506", "507", "508", "509", "510", "511", "598", "599",
	}

	for _, rawEnv := range allEnvs {
		splittedEnv := strings.Split(rawEnv, "=")
		n := splittedEnv[0]
		v := splittedEnv[1]

		if strings.Contains(n, "_PORT") && strings.Contains(v, "tcp://") {
			k8s[n] = v
			continue
		}

		if strings.Contains(n, "_PORT_") && strings.Contains(n, "_TCP") {
			k8s[n] = v
			continue
		}

		if strings.Contains(n, "_SERVICE_HOST") || strings.Contains(n, "_SERVICE_PORT") {
			k8s[n] = v
			continue
		}

		if !includes(blackList, n) {
			envs[n] = v
		}
	}

	envJSON, err := json.Marshal(envs)
	if err != nil {
		return err
	}
	k8sJSON, err := json.Marshal(k8s)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(string(envJSON)), &controller.Envs)
	json.Unmarshal([]byte(string(k8sJSON)), &controller.K8s)

	controller.InstanceID = os.Getenv("HOSTNAME")
	controller.Path = c.Request().URL.Path
	controller.Host = c.Request().Host
	controller.Method = c.Request().Method
	controller.RealIP = c.RealIP()

	responseStatusSTR := "200"
	if includes(statusCodes, c.Request().URL.Path[1:]) {
		responseStatusSTR = c.Request().URL.Path[1:]
	}

	responseStatus, err := strconv.ParseInt(responseStatusSTR, 10, 64)
	if err != nil {
		return err
	}

	controller.Status = int(responseStatus)
	return c.JSON(int(responseStatus), controller)
}
