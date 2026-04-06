package dockerhub

import (
	"fmt"
	"net/http"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/go-resty/resty/v2"
	"github.com/zalando/go-keyring"
)

func (s *DockerHub) ListImages() (*[]types.RegistryImage, error) {
	token, err := keyring.Get("cruise", fmt.Sprintf("%s/%s", s.Provider(), s.Domain()))
	if err != nil {
		return nil, err
	}

	client := resty.New().R().SetAuthToken(token)
	url := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/?page_size=100", s.Username)

	var imgs []types.RegistryImage

	for url != "" {
		var result struct {
			Results []struct {
				Name        string `json:"name"`
				Namespace   string `json:"namespace"`
				Description string `json:"description"`
				IsPrivate   bool   `json:"is_private"`
				PullCount   int64  `json:"pull_count"`
			} `json:"results"`
			Next string `json:"next"`
		}

		resp, err := client.SetResult(&result).Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode() != http.StatusOK {
			return nil, fmt.Errorf("failed to list repositories")
		}

		for _, repo := range result.Results {
			imgs = append(imgs, types.RegistryImage{
				Provider:    s.Provider(),
				Domain:      s.Domain(),
				Project:     repo.Namespace,
				Name:        repo.Name,
				Description: repo.Description,
				IsPrivate:   repo.IsPrivate,
				PullCount:   repo.PullCount,
			})
		}

		url = result.Next
	}

	return &imgs, nil
}
