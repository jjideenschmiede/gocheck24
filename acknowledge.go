package gocheck24

// Acknowledge is to set the order acknowledge to successful
func Acknowledge(documentNumber string, r Request) error {

	// Config new request
	c := Config{
		Path:   "/shop/document/" + documentNumber + "/acknowledge",
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
