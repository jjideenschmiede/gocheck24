# Library for Check24

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gocheck24.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gocheck24/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gocheck24/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gocheck24)](https://goreportcard.com/report/github.com/jjideenschmiede/gocheck24) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gocheck24?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gocheck24) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gocheck24) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

Here you can find our library for shopware 6. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.
## Install

```console
go get github.com/jjideenschmiede/gocheck24
```

## How to use?

Currently we have the following functions covered:

- [Orders & Acknowledge](https://github.com/jjideenschmiede/gocheck24#orders-acknowledge)
- [Dispatch notification](https://github.com/jjideenschmiede/gocheck24#dispatch-notification)

## Orders & Acknowledge

### Order

If you want to read out an order, you can do this with the following function. The user data is required for this. Only one order can be read out at a time. This must be confirmed afterwards with the acknowledge.

```go
// Define request
r := gocheck24.Request{
    Username: "partner187",
    Password: "your_password",
}

// Get orders
order, err := gocheck24.Orders(r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(order)
}
```

### Acknowledge

With this function you confirm the order. The document number is required for this.

```go
// Define request
r := gocheck24.Request{
    Username: "partner187",
    Password: "your_password",
}

// Set acknowledge
err := gocheck24.Acknowledge("3", r)
if err != nil {
    log.Fatalln(err)
}
```

### Dispatch notification

With this function you can transfer the shipping confirmation back to Check24.

```go
// Define request
r := gocheck24.Request{
    Username: "partner187",
    Password: "your_password",
}

// Define body for request
body := gocheck24.DISPATCHNOTIFICATION{
    Xmlns:          "http://www.opentrans.org/XMLSchema/2.1",
    Xsi:            "http://www.w3.org/2001/XMLSchema-instance",
    Bmecat:         "https://www.bme.de/initiativen/bmecat/bmecat-2005",
    SchemaLocation: "http://www.opentrans.org/XMLSchema/2.1 https://merchantcenter.check24.de/sdk/opentrans/schema-definitions/opentrans_2_1.xsd https://www.bme.de/initiativen/bmecat/bmecat-2005 https://merchantcenter.check24.de/sdk/opentrans/schema-definitions/bmecat_2005.xsd http://www.w3.org/2005/05/xmlmime https://merchantcenter.check24.de/sdk/opentrans/schema-definitions/xmlmime.xsd http://www.w3.org/2000/09/xmldsig# https://merchantcenter.check24.de/sdk/opentrans/schema-definitions/xmldsig-core-schema.xsd",
    Version:        "2.1",
    DispatchNotificationHeader: gocheck24.DispatchNotificationBodyHeader{
        DispatchNotificationInfo: gocheck24.DispatchNotificationBodyInfo{
            DispatchNotificationId: "4Y0KGY",
            Parties: gocheck24.DispatchNotificationBodyParties{
                Party: []gocheck24.DispatchNotificationBodyParty{},
            },
            SupplierIdref: gocheck24.DispatchNotificationBodySupplierIdref{
                Type:  "check24",
                Value: "26110",
            },
            ShipmentPartiesReference: gocheck24.DispatchNotificationBodyShipmentPartiesReference{
                DeliveryIdref: gocheck24.DispatchNotificationBodyDeliveryIdref{
                    Type:  "check24",
                    Value: "C24-DA-2085219",
                },
                DelivererIdref: gocheck24.DispatchNotificationBodyDelivererIdref{
                    Type:  "supplier_specific",
                    Value: "dhl",
                },
            },
            ShipmentId:         "123456789",
            TrackingTracingUrl: nil,
		},
    },
    DispatchNotificationItemList: gocheck24.DispatchNotificationBodyDispatchNotificationItemList{
        DispatchNotificationItem: []gocheck24.DispatchNotificationBodyItem{},
	},
    DispatchNotificationSummary: gocheck24.DispatchNotificationBodySummary{
        TotalItemNum: 1,
    },
}

//  Add dispatch notification header
body.DispatchNotificationHeader.DispatchNotificationInfo.Parties.Party = append(body.DispatchNotificationHeader.DispatchNotificationInfo.Parties.Party, gocheck24.DispatchNotificationBodyParty{
    PartyId: gocheck24.DispatchNotificationBodyPartyId{
        Type:  "supplier_specific",
        Value: "dhl",
    },
    PartyRole: "deliverer",
})

// Add dispatch notification item list
body.DispatchNotificationItemList.DispatchNotificationItem = append(body.DispatchNotificationItemList.DispatchNotificationItem, gocheck24.DispatchNotificationBodyItem{
    LineItemId: "am1WTzBYQ0NCaUFHcFdWbHhqRlIrZz09",
    ProductId: gocheck24.DispatchNotificationBodyProductId{
		SupplierPid: 5431,
    },
    Quantity:  1,
    OrderUnit: "C62",
    OrderReference: gocheck24.DispatchNotificationBodyOrderReference{
        OrderId:    "4Y0KGY",
        LineItemId: "am1WTzBYQ0NCaUFHcFdWbHhqRlIrZz09",
    },
    ShipmentPartiesReference: gocheck24.DispatchNotificationBodyShipmentPartiesReference2{
        DeliveryIdref: gocheck24.DispatchNotificationBodyDeliveryIdref{
            Type:  "check24",
            Value: "C24-SA-31122",
        },
    },
})

// Set dispatch notification
err := gocheck24.DispatchNotification(body, r)
if err != nil {
    log.Fatalln(err)
}
```
