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

// DISPATCHNOTIFICATION is to structure the body data
type DISPATCHNOTIFICATION struct {
	Xmlns                        string                                               `xml:"xmlns,attr"`
	Xsi                          string                                               `xml:"xmlns:xsi,attr"`
	Bmecat                       string                                               `xml:"xmlns:bmecat,attr"`
	SchemaLocation               string                                               `xml:"xsi:schemaLocation,attr"`
	Version                      string                                               `xml:"version,attr"`
	DispatchNotificationHeader   DispatchNotificationBodyHeader                       `xml:"DISPATCHNOTIFICATION_HEADER"`
	DispatchNotificationItemList DispatchNotificationBodyDispatchNotificationItemList `xml:"DISPATCHNOTIFICATION_ITEM_LIST"`
	DispatchNotificationSummary  DispatchNotificationBodySummary                      `xml:"DISPATCHNOTIFICATION_SUMMARY"`
}

type DispatchNotificationBodyHeader struct {
	DispatchNotificationInfo DispatchNotificationBodyInfo `xml:"DISPATCHNOTIFICATION_INFO"`
}

type DispatchNotificationBodyInfo struct {
	DispatchNotificationId   string                                           `xml:"DISPATCHNOTIFICATION_ID"`
	Parties                  DispatchNotificationBodyParties                  `xml:"PARTIES"`
	SupplierIdref            DispatchNotificationBodySupplierIdref            `xml:"bmecat:SUPPLIER_IDREF"`
	ShipmentPartiesReference DispatchNotificationBodyShipmentPartiesReference `xml:"SHIPMENT_PARTIES_REFERENCE"`
	ShipmentId               string                                           `xml:"SHIPMENT_ID"`
	TrackingTracingUrl       *string                                          `xml:"TRACKING_TRACING_URL"`
}

type DispatchNotificationBodyParties struct {
	Party []DispatchNotificationBodyParty `xml:"PARTY"`
}

type DispatchNotificationBodyParty struct {
	PartyId   DispatchNotificationBodyPartyId `xml:"bmecat:PARTY_ID"`
	PartyRole string                          `xml:"PARTY_ROLE"`
}

type DispatchNotificationBodyPartyId struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type DispatchNotificationBodySupplierIdref struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type DispatchNotificationBodyShipmentPartiesReference struct {
	DeliveryIdref  DispatchNotificationBodyDeliveryIdref  `xml:"DELIVERY_IDREF"`
	DelivererIdref DispatchNotificationBodyDelivererIdref `xml:"DELIVERER_IDREF"`
}

type DispatchNotificationBodyDeliveryIdref struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type DispatchNotificationBodyDelivererIdref struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type DispatchNotificationBodyDispatchNotificationItemList struct {
	DispatchNotificationItem []DispatchNotificationBodyItem `xml:"DISPATCHNOTIFICATION_ITEM"`
}

type DispatchNotificationBodyItem struct {
	LineItemId               string                                            `xml:"LINE_ITEM_ID"`
	ProductId                DispatchNotificationBodyProductId                 `xml:"PRODUCT_ID"`
	Quantity                 int                                               `xml:"QUANTITY"`
	OrderUnit                string                                            `xml:"bmecat:ORDER_UNIT"`
	OrderReference           DispatchNotificationBodyOrderReference            `xml:"ORDER_REFERENCE"`
	ShipmentPartiesReference DispatchNotificationBodyShipmentPartiesReference2 `xml:"SHIPMENT_PARTIES_REFERENCE"`
}

type DispatchNotificationBodyProductId struct {
	SupplierPid int `xml:"bmecat:SUPPLIER_PID"`
}

type DispatchNotificationBodyOrderReference struct {
	OrderId    string `xml:"ORDER_ID"`
	LineItemId string `xml:"LINE_ITEM_ID"`
}

type DispatchNotificationBodySummary struct {
	TotalItemNum int `xml:"TOTAL_ITEM_NUM"`
}

type DispatchNotificationBodyShipmentPartiesReference2 struct {
	DeliveryIdref DispatchNotificationBodyDeliveryIdref `xml:"DELIVERY_IDREF"`
}

// DispatchNotificationReturn is to decode the xml data
type DispatchNotificationReturn struct {
	XmlName xml.Name                         `xml:"ERRORS"`
	Item    []DispatchNotificationReturnItem `xml:"item"`
}

type DispatchNotificationReturnItem struct {
	XmlName xml.Name `xml:"item"`
	Key     string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

// DispatchNotification is to send the shipping information to check24
func DispatchNotification(body DISPATCHNOTIFICATION, r Request) (DispatchNotificationReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return DispatchNotificationReturn{}, err
	}

	// Config new request
	c := Config{
		Path:   "/shop/document",
		Method: "POST",
		Body:   convert,
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return DispatchNotificationReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return DispatchNotificationReturn{}, err
	}

	// Decode data
	var decode DispatchNotificationReturn
	xml.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}
