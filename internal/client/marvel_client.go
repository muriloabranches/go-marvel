package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MarvelClient struct {
	BaseURL    string
	PublicKey  string
	PrivateKey string
	Timestamp  int64
	Limit      int
	Hash       string
}

func NewMarvelClient(baseURL, publicKey, privateKey string) *MarvelClient {
	ts := time.Now().UnixMilli()
	hash := generateHash(ts, privateKey, publicKey)

	return &MarvelClient{
		BaseURL:    baseURL,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Timestamp:  ts,
		Limit:      100,
		Hash:       hash,
	}
}

type Comics struct {
	Available int `json:"available"`
}

type Thumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type Result struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Comics      Comics    `json:"comics"`
}

type Data struct {
	Offset  int      `json:"offset"`
	Limit   int      `json:"limit"`
	Total   int      `json:"total"`
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

type Response struct {
	Code            int    `json:"code"`
	AttributionText string `json:"attributionText"`
	Data            Data   `json:"data"`
}

func (mc *MarvelClient) GetAllCharacters() ([]*CharacterClientResponse, error) {
	var responses []*CharacterClientResponse

	offset := 0
	url := mc.GenerateURL(offset)
	resp, err := request(url)
	if err != nil {
		return nil, err
	}

	total := resp.Data.Total
	responses = append(responses, convertToClientResponse(resp.Data, resp.AttributionText)...)

	offset += mc.Limit
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := offset; i <= total; i += mc.Limit {
		wg.Add(1)
		url := mc.GenerateURL(i)

		go func() {
			defer wg.Done()

			resp, err := request(url)
			if err != nil {
				fmt.Println(err)
			}

			mu.Lock()
			responses = append(responses, convertToClientResponse(resp.Data, resp.AttributionText)...)
			mu.Unlock()
		}()
	}
	wg.Wait()

	return responses, nil
}

func (mc *MarvelClient) GenerateURL(offset int) string {
	url := fmt.Sprintf("%s/characters?ts=%d&apikey=%s&hash=%s&limit=%d&offset=%d", mc.BaseURL, mc.Timestamp, mc.PublicKey, mc.Hash, mc.Limit, offset)
	return url
}

func generateHash(ts int64, privateKey, publicKey string) string {
	strSlice := []string{strconv.FormatInt(ts, 10), privateKey, publicKey}
	str := strings.Join(strSlice, "")
	hash := md5.Sum([]byte(str))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}

func request(url string) (*Response, error) {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func convertToClientResponse(data Data, attributionText string) []*CharacterClientResponse {
	var responses []*CharacterClientResponse
	for _, r := range data.Results {
		if r.Comics.Available == 0 {
			continue
		}

		if strings.Contains(r.Thumbnail.Path, "image_not_available") {
			continue
		}

		if r.Thumbnail.Extension != "jpg" {
			continue
		}

		responses = append(responses, NewCharacterClientResponse(r.Name, r.Description, strconv.Itoa(r.ID), attributionText, fmt.Sprintf("%s.%s", r.Thumbnail.Path, r.Thumbnail.Extension)))
	}

	return responses
}
