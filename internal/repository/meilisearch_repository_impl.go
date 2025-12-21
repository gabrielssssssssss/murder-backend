package repository

import (
	"fmt"
	"io"

	"github.com/gabrielssssssssss/murder-backend.git/config"
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

func (client *queryImplementation) GetData(element string) {
	_, cancel := config.NewMeilisearchContext()
	defer cancel()
}

func (client *queryImplementation) AddDocument(query *entity.DocumentEntity) (*model.AddDocumentQueryResponse, error) {
	_, cancel := config.NewMeilisearchContext()
	defer cancel()

	csvFile, _ := query.Document.Open()
	defer csvFile.Close()

	body, _ := io.ReadAll(csvFile)

	client.db.CreateIndex(&meilisearch.IndexConfig{
		Uid: query.Index,
	})

	response, err := client.db.Index(query.Index).AddDocumentsCsv(body, &meilisearch.CsvDocumentsQuery{
		PrimaryKey:   "Uuid",
		CsvDelimiter: ";",
	})
	if err != nil {
		return nil, fmt.Errorf("An error occured pending file upload.")
	}

	data := model.AddDocumentQueryResponse{
		TaskUid:    int(response.TaskUID),
		IndexUid:   response.IndexUID,
		Status:     string(response.Status),
		Type:       string(response.Type),
		EnqueuedAt: response.EnqueuedAt.String(),
	}

	return &data, nil
}

func (client *queryImplementation) DelDocument(query *entity.DocumentEntity) (bool, error) {
	_, cancel := config.NewMeilisearchContext()
	defer cancel()

	response, err := client.db.DeleteIndex(query.Index)
	if err != nil {
		return false, fmt.Errorf("An error occured pending index deletion.")
	}

	fmt.Println(response.Status)

	return true, nil
}
