package services

import (
	"errors"
	"net/http"

	rthttpclient "github.com/cobalt77/jfrog-client-go/artifactory/httpclient"
	"github.com/cobalt77/jfrog-client-go/auth"
	clientutils "github.com/cobalt77/jfrog-client-go/utils"
	"github.com/cobalt77/jfrog-client-go/utils/errorutils"
	"github.com/cobalt77/jfrog-client-go/utils/log"
)

type DeleteRepositoryService struct {
	client     *rthttpclient.ArtifactoryHttpClient
	ArtDetails auth.ServiceDetails
}

func NewDeleteRepositoryService(client *rthttpclient.ArtifactoryHttpClient) *DeleteRepositoryService {
	return &DeleteRepositoryService{client: client}
}

func (drs *DeleteRepositoryService) GetJfrogHttpClient() *rthttpclient.ArtifactoryHttpClient {
	return drs.client
}

func (drs *DeleteRepositoryService) Delete(repoKey string) error {
	httpClientsDetails := drs.ArtDetails.CreateHttpClientDetails()
	log.Info("Deleting repository...")
	resp, body, err := drs.client.SendDelete(drs.ArtDetails.GetUrl()+"api/repositories/"+repoKey, nil, &httpClientsDetails)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errorutils.CheckError(errors.New("Artifactory response: " + resp.Status + "\n" + clientutils.IndentJson(body)))
	}

	log.Debug("Artifactory response:", resp.Status)
	log.Info("Done deleting repository.")
	return nil
}
