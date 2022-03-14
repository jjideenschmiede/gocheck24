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
	"encoding/xml"
)

// OrdersReturn is to decode the xml return
type OrdersReturn struct {
	XmlName       xml.Name                  `xml:"ORDER"`
	OrderHeader   OrdersReturnOrderHeader   `xml:"ORDER_HEADER"`
	OrderItemList OrdersReturnOrderItemList `xml:"ORDER_ITEM_LIST"`
	OrderSummary  OrdersReturnOrderSummary  `xml:"ORDER_SUMMARY"`
}

type OrdersReturnOrderHeader struct {
	XmlName     xml.Name                `xml:"ORDER_HEADER"`
	ControlInfo OrdersReturnControlInfo `xml:"CONTROL_INFO"`
	OrderInfo   OrdersReturnOrderInfo   `xml:"ORDER_INFO"`
}

type OrdersReturnControlInfo struct {
	XmlName        xml.Name `xml:"CONTROL_INFO"`
	GeneratorInfo  string   `xml:"GENERATOR_INFO"`
	GenerationDate string   `xml:"GENERATION_DATE"`
}

type OrdersReturnOrderInfo struct {
	XmlName               xml.Name                          `xml:"ORDER_INFO"`
	OrderId               string                            `xml:"ORDER_ID"`
	OrderDate             string                            `xml:"ORDER_DATE"`
	Parties               OrdersReturnParties               `xml:"PARTIES"`
	OrderPartiesReference OrdersReturnOrderPartiesReference `xml:"ORDER_PARTIES_REFERENCE"`
	Currency              string                            `xml:"CURRENCY"`
	Payment               OrdersReturnPayment               `xml:"PAYMENT"`
	Remarks               []OrdersReturnRemarks             `xml:"REMARKS"`
}

type OrdersReturnParties struct {
	XmlName xml.Name            `xml:"PARTIES"`
	Party   []OrdersReturnParty `xml:"PARTY"`
}

type OrdersReturnParty struct {
	XmlName   xml.Name            `xml:"PARTY"`
	PartyId   string              `xml:"PARTY_ID"`
	PartyRole string              `xml:"PARTY_ROLE"`
	Address   OrdersReturnAddress `xml:"ADDRESS"`
}

type OrdersReturnAddress struct {
	XmlName      xml.Name `xml:"ADDRESS"`
	Name         string   `xml:"NAME"`
	Name2        string   `xml:"NAME2"`
	Name3        string   `xml:"NAME3"`
	Street       string   `xml:"STREET"`
	Zip          string   `xml:"ZIP"`
	City         string   `xml:"CITY"`
	Country      string   `xml:"COUNTRY"`
	CountryCoded string   `xml:"COUNTRY_CODED"`
	Phone        string   `xml:"PHONE"`
	Email        string   `xml:"EMAIL"`
}

type OrdersReturnOrderPartiesReference struct {
	XmlName       xml.Name `xml:"ORDER_PARTIES_REFERENCE"`
	BuyerIdref    string   `xml:"BUYER_IDREF"`
	SupplierIdref string   `xml:"SUPPLIER_IDREF"`
}

type OrdersReturnPayment struct {
	XmlName      xml.Name                 `xml:"PAYMENT"`
	Debit        bool                     `xml:"DEBIT"`
	PaymentTerms OrdersReturnPaymentTerms `xml:"PAYMENT_TERMS"`
}

type OrdersReturnPaymentTerms struct {
	XmlName        xml.Name                   `xml:"PAYMENT_TERMS"`
	PaymentTerm    int                        `xml:"PAYMENT_TERM"`
	TimeForPayment OrdersReturnTimeForPayment `xml:"TIME_FOR_PAYMENT"`
}

type OrdersReturnTimeForPayment struct {
	XmlName     xml.Name `xml:"TIME_FOR_PAYMENT"`
	PaymentDate string   `xml:"PAYMENT_DATE"`
}

type OrdersReturnRemarks struct {
	XmlName xml.Name `xml:"REMARKS"`
	Key     string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type OrdersReturnOrderItemList struct {
	XmlName   xml.Name                `xml:"ORDER_ITEM_LIST"`
	OrderItem []OrdersReturnOrderItem `xml:"ORDER_ITEM"`
}

type OrdersReturnOrderItem struct {
	XmlName         xml.Name                    `xml:"ORDER_ITEM"`
	LineItemId      string                      `xml:"LINE_ITEM_ID"`
	ProductId       OrdersReturnProductId       `xml:"PRODUCT_ID"`
	Quantity        int                         `xml:"QUANTITY"`
	OrderUnit       string                      `xml:"ORDER_UNIT"`
	ProductPriceFix OrdersReturnProductPriceFix `xml:"PRODUCT_PRICE_FIX"`
	PriceLineAmount float64                     `xml:"PRICE_LINE_AMOUNT"`
}

type OrdersReturnProductId struct {
	XmlName          xml.Name `xml:"PRODUCT_ID"`
	SupplierPid      int      `xml:"SUPPLIER_PID"`
	DescriptionShort string   `xml:"DESCRIPTION_SHORT"`
}

type OrdersReturnProductPriceFix struct {
	XmlName       xml.Name                  `xml:"PRODUCT_PRICE_FIX"`
	PriceAmount   float64                   `xml:"PRICE_AMOUNT"`
	TaxDetailsFix OrdersReturnTaxDetailsFix `xml:"TAX_DETAILS_FIX"`
}

type OrdersReturnTaxDetailsFix struct {
	XmlName     xml.Name `xml:"TAX_DETAILS_FIX"`
	TaxCategory string   `xml:"TAX_CATEGORY"`
	TaxType     string   `xml:"TAX_TYPE"`
	Tax         float64  `xml:"TAX"`
}

type OrdersReturnOrderSummary struct {
	XmlName      xml.Name `xml:"ORDER_SUMMARY"`
	TotalItemNum int      `xml:"TOTAL_ITEM_NUM"`
	TotalAmount  float64  `xml:"TOTAL_AMOUNT"`
}

// Orders is to get an offer by id
func Orders(r Request) (OrdersReturn, error) {

	// Config new request
	c := Config{
		Path:   "/shop/document",
		Method: "GET",
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Decode data
	var decode OrdersReturn
	xml.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}
