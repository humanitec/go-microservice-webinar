package campaign

import (
	"io/ioutil"
	"log"
)

type ClickService struct {
	Config *Settings
}

func (c *ClickService) AddClick(click Click) error {
	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	return nil
}

func (c *ClickService) Init() error {
	// read if mappings already exists

	// if not create them
	buf, err := ioutil.ReadFile(c.Config.MappingPath)
	if err != nil {
		log.Fatalf("could not initialize the proper")
	}
	log.Printf("%s", buf)
	//mapping := string(buf)

	return nil
}

func NewClickService(config *Settings) *ClickService {
	return &ClickService{Config: config}
}
