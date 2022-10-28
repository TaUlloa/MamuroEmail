package main

import(
    "bufio"
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "os"

    client "github.com/zinclabs/sdk-go-zincsearch"


	func main() {
	query := "query_example" // string | Query

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.Document.Bulk(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Document.Bulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Bulk`: MetaHTTPResponseRecordCount
	fmt.Fprintf(os.Stdout, "Response from `Document.Bulk`: %v\n", resp)
}

)   
func Search(reader *bufio.Scanner, querytype string) {
	key := ReadText(reader, "Enter key")
	value := ReadText(reader, "Enter value")
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			querytype: map[string]interface{}{
				key: value,
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	response, _ := es.Search(es.Search.WithIndex("stsc"), es.Search.WithBody(&buffer))
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		craft := hit.(map[string]interface{})["_source"].(map[string]interface{})
		Print(craft)
	}
}