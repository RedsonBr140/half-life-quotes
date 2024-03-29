<img src="docs/demo.gif" alt="Demo GIF." align="center"/>
<h1 align=center>Half-Life quotes API</h1>
<p align=center>A (very) simple API to retrieve some quotes of Half Life.</p>

## Production host
http://half-life-api.herokuapp.com/

## API

`GET /`

Get a random quote in this format:
> https://half-life-api.herokuapp.com/
```json
[
  {
    "quote":"Welcome to the HEV Mark 4 Protective System, for use in hazardous environment conditions.",
    "author":"HEV Suit"
  }
]
```


`GET /{number}`

return an array with `{number}` quotes.
> https://half-life-api.herokuapp.com/5
```json
[
  {
    "quote": "",
    "author": "Gordon Freeman"
  },
  {
    "quote": "Communications interface online.",
    "author": "HEV Suit"
  },
  {
    "quote": "Defensive weapon selection system activated.",
    "author": "HEV Suit"
  },
  {
    "quote": "Communications interface online.",
    "author": "HEV Suit"
  },
  {
    "quote": "Minor fracture detected.",
    "author": "HEV Suit"
  }
]
```

## Docker
Deploy the [Docker Image](https://hub.docker.com/r/redsonbr/half-life-quotes/) of Half-Life Quotes:
```bash
docker pull redsonbr/half-life-quotes:latest
docker run -d --name half-life-quotes -p 80:8080 redsonbr/half-life-quotes:latest
```
If you want to build your own image, go ahead:
```bash
git clone https://github.com/RedsonBr140/half-life-quotes.git
cd half-life-quotes
docker build -t redsonbr/half-life-quotes .
```
You can also change the `redsonbr/half-life-quotes` tag to anything you want.

## Contributing
If you want to add some quotes, just add them in [`quotes.go`](https://github.com/RedsonBr140/half-life-quotes/blob/main/quotes.go) file and do a pull request!

## Credits
inspired by [Breaking Bad quotes](https://github.com/shevabam/breaking-bad-quotes)

## License
MIT License

---
❤ `Keep It Simple, Stupid.`
