package prestoClient

import "testing"

func TestPrestoClient(t *testing.T) {
	client, err := CreateClient(Config{URL: "127.0.0.1:8080",
		User:    "root",
		Catalog: "mysql",
		Schema:  "swiggy",
	})
	if err != nil {
		t.Fatalf("Error in creating client %v", err)
	}
	_, err = client.ExecuteSync("select id from mysql.swiggy.swiggy_orders limit 100")
	if err != nil {
		t.Fatalf("Error in executing query %v", err)
	}
}
