package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/satioO/togo/app/domain"
)

// ResourceUsecase ...
type ResourceUsecase struct{}

// NewResourceUsecase ...
func NewResourceUsecase() *ResourceUsecase {
	return &ResourceUsecase{}
}

// ListResources ...
func (r *ResourceUsecase) ListResources(context.Context) ([]domain.Resource, error) {
	var files []domain.Resource
	var resource domain.Resource
	path := "./static/resources/"

	fileInfo, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, v := range fileInfo {
		file, _ := ioutil.ReadFile(path + v.Name())
		json.Unmarshal(file, &resource)
		files = append(files, resource)
	}

	return files, nil
}

// CreateResource ...
func (r *ResourceUsecase) CreateResource(ctx context.Context, resource domain.Resource) error {
	file, err := json.Marshal(&resource)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./static/resources/"+resource.Name+".json", file, 0644)
	return err
}
