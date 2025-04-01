package VectorDB

import (
	"bytes"
	Models "colabfiltering/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const IndexName = "user-preferences"

type VectorClient interface {
	SaveUserEmbeddings(user Models.User) error
	BulkSaveEmbeddings(users []Models.User) error
}

type UserPreferences struct {
	UserID             string   `json:"_id"` // Document ID
	FavoriteMusicGenre string   `json:"favorite_music_genre"`
	FavoriteVacation   string   `json:"favorite_vacation_activity"`
	FavoriteCar        string   `json:"favorite_car"`
	TopPlacesVisited   []string `json:"top_places_visited"`
}

type vClient struct {
	baseUrl string
}

func NewClient() (*vClient, error) {
	baseUrl := "http://localhost:8882"

	client := &vClient{
		baseUrl: baseUrl,
	}

	url := baseUrl + "/indexes/" + IndexName

	jsonBody := []byte(`{ "model": "hf/e5-small-v2", "type": "unstructured" }`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{}
	res, err := c.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusConflict {
		fmt.Println("Another index creation/deletion operation is in progress.")
		return client, nil
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Response status: %d\n", res.StatusCode)
		resBody, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Printf("Could not read response body: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Response body: %s\n", resBody)
	}

	return client, nil
}

func (v *vClient) SaveEmbeddings(user Models.User) error {
	url := v.baseUrl + "/indexes/" + IndexName + "/documents"

	data := map[string]interface{}{
		"documents": []Models.Preferences{user.Preferences},
		"tensorFields": []string{
			"ethnicities",
			"hobbies",
			"favFoods",
			"musicGenres",
			"personalityTraits",
		},
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("failed to marshal json: %s", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to store embedding: %s", resp.Status)
	}

	fmt.Println("Embedding stored successfully")

	return nil
}

func (v *vClient) BulkSaveEmbeddings(users []interface{}) error {
	url := v.baseUrl + "/indexes/" + IndexName + "/documents"

	documents := make([]UserPreferences, len(users))

	for i, user := range users {
		documents[i] = user.(UserPreferences)
	}

	data := map[string]interface{}{
		"documents": documents,
		"tensorFields": []string{
			"ethnicities",
			"hobbies",
			"favFoods",
			"musicGenres",
			"personalityTraits",
		},
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("failed to marshal json: %s", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to store embedding: %s", resp.Status)
	}

	fmt.Println("Embedding stored successfully")

	return nil
}

func (v *vClient) QuerySimilarUsers() error {
	url := v.baseUrl + "/indexes/" + IndexName + "/search"

	queryData := map[string]interface{}{
		"q":     "I like Classic music",
		"limit": 10, // Return top 10 similar results
	}

	jsonData, err := json.Marshal(queryData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to query similar users: %s", resp.Status)
	}

	return nil
}
