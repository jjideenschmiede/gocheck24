//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2022 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gocheck24.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gocheck24

import (
	"errors"
	"strings"
)

// statusCodes is to get the status code & return an error
func statusCodes(status string) error {

	// Check each status codes
	switch {
	case strings.Contains(status, "401"):
		return errors.New("access is not authorized")
	case strings.Contains(status, "404"):
		return errors.New("not found")
	default:
		return nil
	}

}
