package campaign

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ClickService struct {
	Config *Settings
	client *http.Client
}

func (c *ClickService) AddClick(click Click) error {
	// convert to json
	body, err := json.Marshal(click)
	if err != nil {
		return fmt.Errorf("error marshaling data - AddClick - %w", err)
	}

	// create request to save the data to elasticsearch
	url := fmt.Sprintf("%v://%v:%v/%v/_doc/%v", c.Config.DatabaseProto, c.Config.DatabaseHost, c.Config.DatabasePort, c.Config.DatabaseIndex, Hash(click))
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error creating request for saving click - AddClick - %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if res, err := c.client.Do(req); err != nil {
		return fmt.Errorf("error saving click - %w", err)
	} else if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		b, err := GetStringFromBuffer(res.Body)
		if err != nil {
			return fmt.Errorf("error handling the response body: %w", err)
		}
		return fmt.Errorf("error creating the index %s, status code %v, body %v", c.Config.DatabaseIndex, res.StatusCode, b)
	}
	return nil
}

func (c *ClickService) Init() error {
	// check if mappings already exists
	url := fmt.Sprintf("%v://%v:%v/%v/_mapping", c.Config.DatabaseProto, c.Config.DatabaseHost, c.Config.DatabasePort, c.Config.DatabaseIndex)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating new GET request - %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if res, err := c.client.Do(req); err != nil {
		return fmt.Errorf("error sending mapping to elasticsearch - %w", err)
	} else if res.StatusCode == http.StatusOK {
		log.Printf("index %s exists, ignoring index creation", c.Config.DatabaseIndex)
		return nil
	} else if res.StatusCode == http.StatusNotFound {
		log.Printf("index %s does not exists", c.Config.DatabaseIndex)
	} else {
		b, err := GetStringFromBuffer(res.Body)
		if err != nil {
			return fmt.Errorf("error handling the response body: %w", err)
		}
		return errors.New(fmt.Sprintf("something weird happened: %v status code, %v response.", res.StatusCode, b))
	}

	// if not, create the index
	buf, err := os.Open(c.Config.MappingPath)
	if err != nil {
		return fmt.Errorf("possible problem with file path - %w", err)
	}
	url = fmt.Sprintf("%v://%v:%v/%v", c.Config.DatabaseProto, c.Config.DatabaseHost, c.Config.DatabasePort, c.Config.DatabaseIndex)
	req, err = http.NewRequest(http.MethodPut, url, buf)
	if err != nil {
		return fmt.Errorf("error creating new PUT request - %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if res, err := c.client.Do(req); err != nil {
		return fmt.Errorf("error creating mapping in elasticsearch - %w", err)
	} else if res.StatusCode != http.StatusOK {
		b, err := GetStringFromBuffer(res.Body)
		if err != nil {
			return fmt.Errorf("error handling the response body: %w", err)
		}
		return fmt.Errorf("error creating the index %s, status code %v, body %v", c.Config.DatabaseIndex, res.StatusCode, b)
	}
	log.Print(fmt.Sprintf("index %s was created successfully", c.Config.DatabaseIndex))
	return nil
}

func NewClickService(config *Settings) *ClickService {
	return &ClickService{Config: config, client: &http.Client{}}
}
