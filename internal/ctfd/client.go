package ctfd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type CTFdClient interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
	CreateTeam(ctx context.Context, req CreateTeamRequest) (CreateTeamResponse, error)
	AddUserToTeam(ctx context.Context, teamID, userID int) error
	UserExists(ctx context.Context, user_id uint) (bool, error)
	TeamExists(ctx context.Context, team_id uint) (bool, error)
}

type Client struct {
	baseURL    string
	adminToken string
	client     *http.Client
}

func NewCTFdClient(baseURL, adminToken string) CTFdClient {

	client := &Client{
		baseURL:    strings.TrimSuffix(baseURL, "/"),
		adminToken: adminToken,
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}

	return client
}

func (c *Client) CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	var response CreateUserResponse

	url := fmt.Sprintf("%s/users", c.baseURL)

	err := c.makeRequest(ctx, "POST", url, req, &response)
	if err != nil {
		return response, fmt.Errorf("failed to create user: %w", err)
	}

	if !response.Success {
		return response, fmt.Errorf("CTFd API returned failure for user creation")
	}
	response.Data.CTFDPass = req.Password
	return response, nil
}

func (c *Client) CreateTeam(ctx context.Context, req CreateTeamRequest) (CreateTeamResponse, error) {
	var response CreateTeamResponse

	url := fmt.Sprintf("%s/teams", c.baseURL)

	err := c.makeRequest(ctx, "POST", url, req, &response)
	if err != nil {
		return response, err
	}

	if !response.Success {
		return response, err
	}

	return response, nil
}

func (c *Client) AddUserToTeam(ctx context.Context, teamID, userID int) error {

	return nil
}

func (c *Client) UserExists(ctx context.Context, user_id uint) (bool, error) {
	var response ExistenceResponse
	url := fmt.Sprintf("%s/users/%d", c.baseURL, user_id)

	err := c.makeRequest(ctx, "GET", url, nil, &response)
	if err != nil {
		return response.Success, fmt.Errorf("failed to create team: %w", err)
	}

	if !response.Success {
		return response.Success, fmt.Errorf("CTFd API returned failure for team creation")
	}

	return response.Success, nil
}

func (c *Client) TeamExists(ctx context.Context, team_id uint) (bool, error) {
	var response ExistenceResponse
	url := fmt.Sprintf("%s/teams/%d", c.baseURL, team_id)

	err := c.makeRequest(ctx, "GET", url, nil, &response)
	if err != nil {
		return response.Success, fmt.Errorf("failed to create team: %w", err)
	}

	if !response.Success {
		return response.Success, fmt.Errorf("CTFd API returned failure for team creation")
	}

	return response.Success, nil
}

func (c *Client) makeRequest(ctx context.Context, method, url string, requestData interface{}, responseData interface{}) error {
	var body io.Reader

	if requestData != nil {
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			return fmt.Errorf("failed to marshal request data: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.adminToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {

		var errorResp struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		if err := json.Unmarshal(respBody, &errorResp); err == nil && errorResp.Message != "" {
			return fmt.Errorf("CTFd API error (%d): %s", resp.StatusCode, errorResp.Message)
		}

		return fmt.Errorf("CTFd API returned status %d: %s", resp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, responseData); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w. Response: %s", err, string(respBody))
	}

	return nil
}
