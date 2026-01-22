package ctfd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	url := fmt.Sprintf("%s/api/v1/users", c.baseURL)

	err, respCode := c.makeRequest(ctx, "POST", url, req, &response)
	response.Data.CTFDPass = req.Password
	response.StatusCode = respCode
	if err != nil {
		return response, fmt.Errorf("failed to create user: %w", err)
	}

	if !response.Success {
		return response, fmt.Errorf("CTFd API returned failure for user creation")
	}

	return response, nil
}

func (c *Client) CreateTeam(ctx context.Context, req CreateTeamRequest) (CreateTeamResponse, error) {
	var response CreateTeamResponse

	url := fmt.Sprintf("%s/api/v1/teams", c.baseURL)

	err, respCode := c.makeRequest(ctx, "POST", url, req, &response)
	response.StatusCode = respCode
	if err != nil {
		return response, err
	}

	if !response.Success {
		return response, fmt.Errorf("CTFd API returned failure for team creation")
	}
	return response, nil
}

func (c *Client) AddUserToTeam(ctx context.Context, teamID, userID int) error {
	var response ExistenceResponse

	req := AddToTeamRequest{UserID: userID}
	url := fmt.Sprintf("%s/api/v1/teams/%d/members", c.baseURL, teamID)

	err, respCode := c.makeRequest(ctx, "POST", url, req, &response)
	response.StatusCode = respCode
	if err != nil {
		return err
	}

	if !response.Success {
		return fmt.Errorf("CTFd API returned failure for adding user to team")
	}
	return nil
}

func (c *Client) UserExists(ctx context.Context, user_id uint) (bool, error) {
	var response ExistenceResponse
	url := fmt.Sprintf("%s/api/v1/users/%d", c.baseURL, user_id)

	err, respCode := c.makeRequest(ctx, "GET", url, nil, &response)
	response.StatusCode = respCode
	if err != nil {
		return response.Success, fmt.Errorf("failed to check user: %w", err)
	}
	return response.Success, nil
}

func (c *Client) TeamExists(ctx context.Context, team_id uint) (bool, error) {
	var response ExistenceResponse
	url := fmt.Sprintf("%s/api/v1/teams/%d", c.baseURL, team_id)

	err, respCode := c.makeRequest(ctx, "GET", url, nil, &response)
	response.StatusCode = respCode
	if err != nil {
		return response.Success, fmt.Errorf("failed to check team: %w", err)
	}
	return response.Success, nil
}
func (c *Client) retryRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	if resp.StatusCode >= 500 {
		resp.Body.Close()
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}
	return resp, err
}
func (c *Client) makeRequest(ctx context.Context, method, url string, requestData interface{}, responseData interface{}) (error, int) {
	var body io.Reader
	var jsonData []byte
	if requestData != nil {
		data, err := json.Marshal(requestData)
		if err != nil {
			log.Printf("failed to marshal request data")
			return err, 0
		}
		body = bytes.NewBuffer(data)
		jsonData = data
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		log.Println("error occurred while making request")
		return err, 0
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.adminToken))

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println("failed to send request")
		return err, 0
	}

	if resp.StatusCode >= 500 {
		log.Println("re-attempting request")
		resp.Body.Close()
		for i := range 5 {
			if jsonData != nil {
				body = bytes.NewBuffer(jsonData)
			}

			req, err := http.NewRequestWithContext(ctx, method, url, body)
			if err != nil {
				log.Printf("attempt %d failed: %s", i, err.Error())
				continue
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.adminToken))

			resp, err = c.retryRequest(req)
			if err != nil {
				log.Printf("attempt %d failed: %s", i, err.Error())
			} else {
				break
			}
			time.Sleep(time.Second * 1)
		}
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err), 0
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {

		var errorResp struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		if err := json.Unmarshal(respBody, &errorResp); err == nil && errorResp.Message != "" {
			return fmt.Errorf("CTFd API error (%d): %s", resp.StatusCode, errorResp.Message), resp.StatusCode
		}

		return fmt.Errorf("CTFd API returned status %d: %s", resp.StatusCode, string(respBody)), resp.StatusCode
	}

	if err := json.Unmarshal(respBody, responseData); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w. Response: %s", err, string(respBody)), 0
	}
	return nil, resp.StatusCode
}
