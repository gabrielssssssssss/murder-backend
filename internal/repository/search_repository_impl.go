package repository

import (
	"fmt"
	"io"

	"github.com/gabrielssssssssss/murder-backend.git/config"
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

func (client *queryImplementation) AddIndex(query *entity.IndexEntity) (*model.AddIndexResponse, error) {
	_, cancel := config.NewMeilisearchContext()
	defer cancel()

	csvFile, _ := query.Document.Open()
	defer csvFile.Close()

	body, err := io.ReadAll(csvFile)
	if err != nil {
		return nil, fmt.Errorf("An error occured pending file reading (csv).")
	}

	client.db.CreateIndex(&meilisearch.IndexConfig{
		Uid: query.Name,
	})

	response, err := client.db.Index(query.Name).AddDocumentsCsv(body, &meilisearch.CsvDocumentsQuery{
		PrimaryKey:   "Uuid",
		CsvDelimiter: ";",
	})
	if err != nil {
		return nil, fmt.Errorf("An error occured pending file upload.")
	}

	data := model.AddIndexResponse{
		TaskUid:    int(response.TaskUID),
		IndexUid:   response.IndexUID,
		Status:     string(response.Status),
		Type:       string(response.Type),
		EnqueuedAt: response.EnqueuedAt.String(),
	}

	return &data, nil
}

func (client *queryImplementation) GetIndex(query *entity.IndexEntity) (*meilisearch.Stats, error) {
	_, cancel := config.NewMeilisearchContext()
	defer cancel()

	response, err := client.db.GetStats()
	if err != nil {
		return nil, err
	}

	if query.Name != "" {
		for key, val := range response.Indexes {
			if key == query.Name {
				return &meilisearch.Stats{
					DatabaseSize:     response.DatabaseSize,
					UsedDatabaseSize: response.UsedDatabaseSize,
					LastUpdate:       response.LastUpdate,
					Indexes: map[string]meilisearch.StatsIndex{
						query.Name: val,
					},
				}, nil
			}
		}
		return nil, fmt.Errorf("Index name not found.")
	}

	return response, nil
}

func (client *queryImplementation) DeleteIndex(query *entity.IndexEntity) (*model.IndexResult, error) {
	_, cancel := config.NewMeilisearchContext()
	defer cancel()

	_, err := client.db.GetIndex(query.Name)
	if err != nil {
		return nil, fmt.Errorf("Index name doesn't exist.")
	}

	response, err := client.db.DeleteIndex(query.Name)
	if err != nil {
		return nil, fmt.Errorf("An error occured pending index deletion.")
	}

	data := &model.IndexResult{
		TaskUID:    response.TaskUID,
		IndexUID:   response.IndexUID,
		EnqueuedAt: response.EnqueuedAt,
	}

	return data, nil
}
