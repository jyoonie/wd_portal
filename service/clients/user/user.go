package user

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

func (c *Client) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	targetURL := fmt.Sprintf("%s%s%s", c.baseURL, "/users/", id.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to get user: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to get user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting user, received status %d from service", resp.StatusCode)
	}

	var user *models.User
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (c *Client) CreateUser(ctx context.Context, u models.User) (*models.User, error) {
	targetURL := fmt.Sprintf("%s%s", c.baseURL, "/users")

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(u)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request to create user: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to create user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating user, received status %d from service", resp.StatusCode)
	}

	var createdUser *models.User
	err = json.NewDecoder(resp.Body).Decode(createdUser)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return createdUser, nil
}

func (c *Client) UpdateUser(ctx context.Context, u models.User) (*models.User, error) {
	targetURL := fmt.Sprintf("%s%s%s", c.baseURL, "/users/", u.UserUUID.String())

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(u)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request to update user: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to update user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error updating user, received status %d from service", resp.StatusCode)
	}

	var updatedUser *models.User
	err = json.NewDecoder(resp.Body).Decode(updatedUser)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return updatedUser, nil
}

//wrapped error?
//you have to read the body after you do the Do function
//written to and read from, so that satisfies the io.reader interface and io.writer interface
