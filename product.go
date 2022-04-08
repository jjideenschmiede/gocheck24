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

// ProductsBody is to structure the body data for the csv file
type ProductsBody struct {
	ProductId     string
	Manufacturer  string
	Mpnr          string
	Ean           string
	ProductName   string
	Description   string
	CategoryPath  string
	Price         string
	MinPrice      string
	ImageUrl      string
	ImageUrl1     string
	ImageUrl2     string
	ImageUrl3     string
	ImageUrl4     string
	ImageUrl5     string
	ImageUrl6     string
	ImageUrl7     string
	ImageUrl8     string
	ImageUrl9     string
	ImageUrl10    string
	ImageUrl11    string
	DeliveryTime  string
	DeliveryCosts string
	Weight        string
	Stock         string
	ShippingType  string
}
