# location-taxrate-api

Modified from a [demo application](https://github.com/moficodes/bookdata-api) by [Mofizur Rahman](https://github.com/moficodes) showing how to build rest api using golang. See git history for changes.

Tax rates from
[1](https://assets.ey.com/content/dam/ey-sites/ey-com/en_gl/topics/tax/guides/ey-2021-worldwide-vat-gst-sales-tax-guide-web-v3.pdf?download)

Territories labelled using ISO_3166-2
[2](https://en.wikipedia.org/wiki/ISO_3166-2)
[3](https://www.iso.org/iso-3166-country-codes.html)

## Run Locally

To run, from the root of the repo

```
go run .
```

## Access the app 

The App can be extended to have more than one endpoints

All api endpoints are prefixed with `/api/v1`

To reach any endpoint use `baseurl:8080/api/v1/{endpoint}`

```text
Get taxrate by location
"/taxrate/location/{location}" 
```
