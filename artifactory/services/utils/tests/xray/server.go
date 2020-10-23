package xray

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
	"github.com/cobalt77/jfrog-client-go/utils/log"
	clienttests "github.com/cobalt77/jfrog-client-go/utils/tests"
)

const (
	CleanScanBuildName  = "cleanBuildName"
	FatalScanBuildName  = "fatalBuildName"
	VulnerableBuildName = "vulnerableBuildName"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buildName, err := jsonparser.GetString(body, "buildName")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch buildName {
	case CleanScanBuildName:
		fmt.Fprint(w, CleanXrayScanResponse)
		return
	case FatalScanBuildName:
		fmt.Fprint(w, FatalErrorXrayScanResponse)
		return
	case VulnerableBuildName:
		fmt.Fprint(w, VulnerableXrayScanResponse)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func StartXrayMockServer() int {
	handlers := clienttests.HttpServerHandlers{}
	handlers["/api/xray/scanBuild"] = handler
	handlers["/"] = http.NotFound

	port, err := clienttests.StartHttpServer(handlers)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	return port
}
