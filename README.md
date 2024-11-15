# bubble-pocs #

A project to try to study bubble tea TUI library.


## reference links ##

- [Intro to Bubble Tea in Go - Andy Haskell](https://dev.to/andyhaskell/intro-to-bubble-tea-in-go-21lg)
- [Building UI of Golang CLI app with Bubble Tea - Vladimir Dulenov](https://medium.com/@originalrad50/building-ui-of-golang-cli-app-with-bubble-tea-68b61e25445e)
- [I don't get Bubbletea - Reddit](https://www.reddit.com/r/golang/comments/xvrhow/i_dont_get_bubbletea/)
- [Views - Bubble Tea examples - GitHub](https://github.com/charmbracelet/bubbletea/blob/main/examples/views/main.go)
- [bubbletea/examples/textinputs/main.go - GitHub](https://github.com/charmbracelet/bubbletea/blob/v1.2.2/examples/textinputs/main.go)
- [Aikar's Start Script Generator tsx code - Github](https://github.com/PaperMC/docs/blob/6d98e168827d8129bc57fc2f687c559f561a560b/src/components/StartScriptGenerator.tsx)
- [Aikar's Flags - PaperMC](https://docs.papermc.io/paper/aikars-flags)
- [Aikar's Flags' based Startup Script Generator](https://flags.sh/)


## some useful APIs to interact ##

- [Mojang API](https://wiki.vg/Mojang_API)
- [Powerful Minecraft APIs](https://api.minetools.eu/)
- [CoinDesk - Public APIs](https://publicapis.io/coin-desk-api)
- [Purpur Minecraft Server Versions](https://purpurmc.org/docs/purpur/)

```shell
curl https://api.coindesk.com/v1/bpi/currentprice/BRL.json | jq .

{
  "time": {
    "updated": "Oct 18, 2024 00:10:45 UTC",
    "updatedISO": "2024-10-18T00:10:45+00:00",
    "updateduk": "Oct 18, 2024 at 01:10 BST"
  },
  "disclaimer": "This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org",
  "bpi": {
    "USD": {
      "code": "USD",
      "rate": "67,319.253",
      "description": "United States Dollar",
      "rate_float": 67319.2527
    },
    "BRL": {
      "code": "BRL",
      "rate": "380,575.931",
      "description": "Brazil Real",
      "rate_float": 380575.9312
    }
  }
}
```

```shell
curl https://api.coindesk.com/v1/bpi/currentprice.json | jq .

{
  "time": {
    "updated": "Oct 18, 2024 00:14:07 UTC",
    "updatedISO": "2024-10-18T00:14:07+00:00",
    "updateduk": "Oct 18, 2024 at 01:14 BST"
  },
  "disclaimer": "This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org",
  "chartName": "Bitcoin",
  "bpi": {
    "USD": {
      "code": "USD",
      "symbol": "&#36;",
      "rate": "67,286.319",
      "description": "United States Dollar",
      "rate_float": 67286.3186
    },
    "GBP": {
      "code": "GBP",
      "symbol": "&pound;",
      "rate": "51,699.645",
      "description": "British Pound Sterling",
      "rate_float": 51699.6448
    },
    "EUR": {
      "code": "EUR",
      "symbol": "&euro;",
      "rate": "62,134.205",
      "description": "Euro",
      "rate_float": 62134.2052
    }
  }
}
```


```shell
curl -L \
  -H "Accept: application/vnd.github+json" \
  https://api.github.com/repos/PurpurMC/PurpurExtras/releases
```