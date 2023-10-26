# CreditCardOutput


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Expiration`                                                  | *string*                                                      | :heavy_check_mark:                                            | The expiration date, in YYYY-MM format.                       | 2029-03                                                       |
| `Last4`                                                       | *string*                                                      | :heavy_check_mark:                                            | The account number's last four digits.                        | 1004                                                          |
| `Network`                                                     | [CreditCardNetwork](../../models/shared/creditcardnetwork.md) | :heavy_check_mark:                                            | The credit card's network.                                    | visa                                                          |