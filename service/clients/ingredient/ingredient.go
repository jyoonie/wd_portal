package ingredient

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

func (c *Client) GetIngredient(ctx context.Context, id uuid.UUID) (*models.Ingredient, error) {
	targetURL := fmt.Sprintf("%s%s%s", c.baseURL, "/ingredients/", id.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to get ingredient: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to get ingredient: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting ingredient, received status %d from service", resp.StatusCode)
	}

	var ingredient *models.Ingredient
	err = json.NewDecoder(resp.Body).Decode(ingredient)
	if err != nil {
		return nil, fmt.Errorf("error getting ingredient: %w", err)
	}

	return ingredient, nil
}

func (c *Client) CreateIngredient(ctx context.Context, u models.Ingredient) (*models.Ingredient, error) {
	targetURL := fmt.Sprintf("%s%s", c.baseURL, "/ingredients")

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(u)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request to create ingredient: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to create ingredient: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating ingredient, received status %d from service", resp.StatusCode)
	}

	var createdIngredient *models.Ingredient
	err = json.NewDecoder(resp.Body).Decode(createdIngredient)
	if err != nil {
		return nil, fmt.Errorf("error creating ingredient: %w", err)
	}

	return createdIngredient, nil
}

func (c *Client) UpdateIngredient(ctx context.Context, u models.Ingredient) (*models.Ingredient, error) {
	targetURL := fmt.Sprintf("%s%s%s", c.baseURL, "/ingredients/", u.IngredientUUID.String())

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(u)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request to update ingredient: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to update ingredient: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error updating ingredient, received status %d from service", resp.StatusCode)
	}

	var updatedIngredient *models.Ingredient
	err = json.NewDecoder(resp.Body).Decode(updatedIngredient)
	if err != nil {
		return nil, fmt.Errorf("error updating ingredient: %w", err)
	}

	return updatedIngredient, nil
}

//wrapped error?
//you have to read the body after you do the Do function
//written to and read from, so that satisfies the io.reader interface and io.writer interface
