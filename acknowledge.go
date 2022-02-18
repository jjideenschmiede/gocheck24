package gocheck24

import (
	"fmt"
)

// Acknowledge is to set the order acknowledge to successful
func Acknowledge(documentNumber int, r Request) error {

	// Config new request
	c := Config{
		Path:   fmt.Sprintf("/shop/document/%d/acknowledge", documentNumber),
		Method: "PUT",
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return err
	}

	// Return data
	return nil

}
