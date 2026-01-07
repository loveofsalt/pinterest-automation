package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Simplified structures for Base64 only
type MediaSource struct {
	SourceType  string `json:"source_type"`
	Data        string `json:"data"`
	ContentType string `json:"content_type"`
}

type PinRequest struct {
	BoardID        string      `json:"board_id"`
	BoardSectionID string      `json:"board_section_id,omitempty"`
	Title          string      `json:"title,omitempty"`
	Description    string      `json:"description,omitempty"`
	Link           string      `json:"link,omitempty"`
	AltText        string      `json:"alt_text,omitempty"`
	Note           string      `json:"note,omitempty"`
	MediaSource    MediaSource `json:"media_source"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func main() {
	// 1. Gather Inputs
	appID := os.Getenv("PINTEREST_APP_ID")
	appSecret := os.Getenv("PINTEREST_APP_SECRET")
	refreshToken := os.Getenv("PINTEREST_REFRESH_TOKEN")
	boardID := os.Getenv("PINTEREST_BOARD_ID")
	
	// Workflow Inputs
	filePath := os.Getenv("INPUT_FILE_PATH") // Renamed for clarity
	title := os.Getenv("INPUT_TITLE")
	desc := os.Getenv("INPUT_DESCRIPTION")
	link := os.Getenv("INPUT_LINK")
	altText := os.Getenv("INPUT_ALT_TEXT")
	sectionID := os.Getenv("INPUT_SECTION_ID")
	note := os.Getenv("INPUT_NOTE")

	if appID == "" || refreshToken == "" {
		log.Fatal("❌ Missing critical secrets (App ID or Refresh Token)")
	}

	// 2. Process the Image File
	log.Printf("📂 Reading file: %s", filePath)
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("❌ Error reading file: %v. \n(Did you commit the image to the repo and check the path spelling?)", err)
	}

	// Detect Content-Type
	mimeType := http.DetectContentType(fileBytes)
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		log.Fatalf("❌ Invalid file type: %s. Pinterest only accepts jpg or png.", mimeType)
	}

	// Encode to Base64
	base64Str := base64.StdEncoding.EncodeToString(fileBytes)

	// 3. Authenticate and Pin
	log.Println("🔄 Authenticating...")
	token := getAccessToken(appID, appSecret, refreshToken)

	log.Println("📌 Uploading and Pinning...")
	createPin(token, boardID, sectionID, title, desc, link, altText, note, base64Str, mimeType)
}

func getAccessToken(clientID, clientSecret, refreshToken string) string {
	auth := clientID + ":" + clientSecret
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, _ := http.NewRequest("POST", "https://api.pinterest.com/v5/oauth/token", strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "Basic "+encodedAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("❌ API Auth Error: %s", string(body))
	}

	var tokenResp TokenResponse
	json.NewDecoder(resp.Body).Decode(&tokenResp)
	return tokenResp.AccessToken
}

func createPin(token, boardID, sectionID, title, desc, link, altText, note, base64Data, contentType string) {
	payload := PinRequest{
		BoardID:        boardID,
		BoardSectionID: sectionID,
		Title:          title,
		Description:    desc,
		Link:           link,
		AltText:        altText,
		Note:           note,
		MediaSource: MediaSource{
			SourceType:  "image_base64",
			Data:        base64Data,
			ContentType: contentType,
		},
	}

	jsonData, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api-sandbox.pinterest.com/v5/pins", bytes.NewBuffer(jsonData))
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("❌ Error creating pin (Status %d): %s", resp.StatusCode, string(body))
	}

	log.Println("✅ Pin successfully created!")
}
