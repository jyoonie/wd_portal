package recipe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	models "github.com/jyoonie/wd_models"
)

type Client struct {
	hc      http.Client
	baseURL string
}

func New() *Client {
	return &Client{
		hc:      *http.DefaultClient,
		baseURL: "", //set this later once the user service URL is known
	}
}

func (c *Client) GetRecipe(ctx context.Context, id uuid.UUID) (*models.Recipe, error) {
	targetURL := fmt.Sprintf("%s%s%s", c.baseURL, "/recipes/", id.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to get recipe: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to get recipe: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting recipe, received status %d from service", resp.StatusCode)
	}

	var recipe *models.Recipe
	err = json.NewDecoder(resp.Body).Decode(recipe)
	if err != nil {
		return nil, fmt.Errorf("error getting recipe: %w", err)
	}

	return recipe, nil
}

func (c *Client) CreateRecipe(ctx context.Context, u models.Recipe) (*models.Recipe, error) {
	targetURL := fmt.Sprintf("%s%s", c.baseURL, "/recipes")

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(u)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request to create recipe: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to create recipe: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating recipe, received status %d from service", resp.StatusCode)
	}

	var createdRecipe *models.Recipe
	err = json.NewDecoder(resp.Body).Decode(createdRecipe)
	if err != nil {
		return nil, fmt.Errorf("error creating recipe: %w", err)
	}

	return createdRecipe, nil
}

func (c *Client) UpdateRecipe(ctx context.Context, u models.Recipe) (*models.Recipe, error) {
	targetURL := fmt.Sprintf("%s%s%s", c.baseURL, "/recipes/", u.RecipeUUID.String())

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(u)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request to update recipe: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to update recipe: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error updating recipe, received status %d from service", resp.StatusCode)
	}

	var updatedRecipe *models.Recipe
	err = json.NewDecoder(resp.Body).Decode(updatedRecipe)
	if err != nil {
		return nil, fmt.Errorf("error updating recipe: %w", err)
	}

	return updatedRecipe, nil
}

//wrapped error?
//you have to read the body after you do the Do function
//written to and read from, so that satisfies the io.reader interface and io.writer interface
