# fetchph
Go CLI program that fetches Product Hunt's top product's.

## Usage
Clone the repo:
```
git clone https://github.com/alwin-7/fetchph.git
```

Obtain your developer token from producthunt's API dashboard.

Insert your api token in line 72 at main.go file:
```js
70 }
71	request.Header.Add("content-type", "application/json")
72	request.Header.Add("Authorization", "{Your-BearerToken-Here}") <- insert the api token here
```

Compile the code by running:
```
go build
```

Run:
```
./fetchph
```
Default: To fetch top 5 products of the year 2026 in month January.
