package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"text/template"
)

type Param struct {
	IpaYear  string
	IpbYear  string
	InumPrgm int
}

func Fetchres() string {

	setYear := flag.String("year", "2026", "Specify the Year")
	setTop := flag.Int("top", 5, "Set Number of programs to list")
	setMonth := flag.String("month", "01", "Specify the Month in numeric value")

	flag.Parse()

	paYear := Param{
		IpaYear:  *setYear + "-" + *setMonth + "-01T00:00:00Z",
		IpbYear:  *setYear + "-" + *setMonth + "-30T23:30:00Z",
		InumPrgm: *setTop,
	}

	tmpl, err := template.New("query").Parse(`
	    {
				posts(postedAfter: "{{ .IpaYear}}" , postedBefore: "{{ .IpbYear}}", first: {{ .InumPrgm}}, featured: true, order: VOTES) {
	 		 		totalCount,
					nodes {
						name,
						website
					}
		}
	     }
	`)

	if err != nil {
		panic(err)
	}

	var s strings.Builder
	err = tmpl.Execute(&s, paYear)
	if err != nil {
		panic(err)
	}

	query := s.String()

	jasonData := map[string]string{
		"query": query,
	}

	jsonValue, err := json.Marshal(jasonData)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "https://api.producthunt.com/v2/api/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	request.Header.Add("content-type", "application/json")
	request.Header.Add("Authorization", "{Your-BearerToken-Here}")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	rData := string(data)
	return rData
}

func main() {

	type Returndata struct {
		Data struct {
			Posts struct {
				TotalCount int `json:"totalCount"`
				Nodes      []struct {
					Name    string `json:"name"`
					Website string `json:"website"`
				} `json:"nodes"`
			} `json:"posts"`
		} `json:"data"`
	}

	getOutput := Fetchres()
	var returndata Returndata

	err := json.Unmarshal([]byte(getOutput), &returndata)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`\?.*$`)

	for _, data := range returndata.Data.Posts.Nodes {
		fmt.Printf("Name: %s\n", data.Name)
		str := data.Website
		webS := re.ReplaceAllString(str, "/")
		fmt.Printf("Website: %s\n", webS)
		fmt.Print("\n")
	}

}
