# CartShipment


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Address`                                                                  | [*shared.AddressReference](../../../pkg/models/shared/addressreference.md) | :heavy_minus_sign:                                                         | N/A                                                                        |                                                                            |
| `Carrier`                                                                  | **string*                                                                  | :heavy_minus_sign:                                                         | The name of the carrier selected.                                          | FedEx                                                                      |
| `Cost`                                                                     | [*shared.Amount](../../../pkg/models/shared/amount.md)                     | :heavy_minus_sign:                                                         | A monetary amount, i.e. a base unit amount and a supported currency.       |                                                                            |