package search

import (
	"github.com/meilisearch/meilisearch-go"
)

type SearchClient struct {
	Host   string
	Client *meilisearch.Client
}

// NewSearchClient creates a new search client
func NewSearchClient(host string, apiKey string) *SearchClient {
	return &SearchClient{
		Host: host,
		Client: meilisearch.NewClient(meilisearch.ClientConfig{
			Host:   host,
			APIKey: apiKey,
		}),
	}
}

// AddDocuments adds single or multiple documents to an index
func (s *SearchClient) AddDocument(indexName string, document map[string]interface{}) error {
	_, err := s.Client.Index(indexName).AddDocuments(document)
	if err != nil {
		return err
	}
	return nil
}

// Search searches an index
func (s *SearchClient) Search(indexName string, query string, offset int64, limit int64, filter string) (*meilisearch.SearchResponse, error) {
	res, err := s.Client.Index(indexName).Search(query, &meilisearch.SearchRequest{
		Offset: offset,
		Limit:  limit,
		Filter: filter,
	})
	if err != nil {
		return &meilisearch.SearchResponse{}, err
	}
	return res, nil
}

// DeleteIndexDocuments deletes all documents of an index
func (s *SearchClient) DeleteIndexDocuments(indexName string) error {
	_, err := s.Client.Index(indexName).DeleteAllDocuments()
	if err != nil {
		return err
	}
	return nil
}

// GetDocument returns a document
func (s *SearchClient) GetDocument(indexName string, documentID string) (interface{}, error) {
	var res interface{}
	err := s.Client.Index(indexName).GetDocument(documentID, &meilisearch.DocumentQuery{}, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetDocuments returns multiple documents
func (s *SearchClient) GetDocuments(indexName string) ([]map[string]interface{}, error) {
	var res meilisearch.DocumentsResult
	err := s.Client.Index(indexName).GetDocuments(&meilisearch.DocumentsQuery{}, &res)
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}

// UpdateDocument updates a document
func (s *SearchClient) UpdateDocument(indexName string, documentID string, document map[string]interface{}) error {
	_, err := s.Client.Index(indexName).UpdateDocuments(document, documentID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDocument deletes a document
func (s *SearchClient) DeleteDocument(indexName string, documentID string) error {
	_, err := s.Client.Index(indexName).DeleteDocument(documentID)
	if err != nil {
		return err
	}
	return nil
}
