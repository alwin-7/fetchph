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

Default: To fetch top 5 products of the year 2026 in month January.
```

Flags:

```
./fetchph --help

--month 07 -> To fetch the top products for a specific month.
In numeric value 07 = Month July

--top 10 -> To fetch the top 10 products of a specific month.
In numeric value 10

--year 2026 -> To fetch the top products from the year 2026.
In numeric value 2026
```

Example Usage:
```cmd
~$ ./fetchph --year 2024 --top 3 --month 04
Name: Supabase
Website: https://www.producthunt.com/r/NEVU7ICMHAJAYL/

Name: Cal.com Platform
Website: https://www.producthunt.com/r/6TKDB3VFB3IDQT/

Name: PaddleBoat
Website: https://www.producthunt.com/r/JDLZRYWPFITVYZ/
```
