package transfers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"user_service/internal/repository"
)

type Repository struct {
	log    *logrus.Logger
	client *http.Client
}

func NewRepository(client *http.Client, log *logrus.Logger) Repository {
	return Repository{
		log:    log,
		client: client,
	}
}

func (r Repository) Transfer(IDFrom, IDTo, Amount int) error {
	url := "http://localhost:3001/v1/tx-service/transfer"
	body := strings.NewReader(fmt.Sprintf("{\n    \"user_id_from\": %d,\n    \"user_id_to\": %d,\n    \"amount\": %d\n}", IDFrom, IDTo, Amount))

	resp, err := r.client.Post(url, "application/json", body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 300 {
		return repository.InternalServerError{}
	}

	return nil
}
