# crypto_com
```bash
cp .env.example .env
```

fill the values according the API authentication created in [crypto.com exchange](https://crypto.com/exchange/personal/api-management)


## TL;DR;
```bash
go mod vendor && \
    go build -o bid src/* && \
    ./bid 0.001
```

## Build
```bash
go mod download # fetching dependencies
go build -o bid src/* # build the src content into an executable called 'bid'
./bid 0.001 # runs the executable with the desired BTC amount to buy at average price
```

## Dev
```bash
air . # golang live reload
```


## Description

this project is meant to cancel old orders and place new ones to updated prices;

ideally it runs every 2 hours


## Buy me some weed

- BTC: 18GLPePTDMxPxwuFaFHhVNGZXihsNuYhK5
- ETH: 0x8df8b0c99184d9305018f5d45c13437922d4f9d3
- USDT: 0x8df8b0c99184d9305018f5d45c13437922d4f9d3
