package translator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Translate translates text from one language to another using Google Translate
func Translate(text, from, to string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("empty text provided for translation")
	}

	// Build the URL for Google Translate
	baseURL := "https://translate.googleapis.com/translate_a/single"

	// Prepare query parameters
	params := url.Values{}
	params.Add("client", "gtx")
	params.Add("sl", from) // Source language
	params.Add("tl", to)   // Target language
	params.Add("dt", "t")  // Return translated text
	params.Add("q", text)  // Text to translate

	// Construct the full URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Add headers to make request more reliable
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "*/*")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("API returned non-OK status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Read response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Parse the response
	var result []interface{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %v, body: %s", err, string(bodyBytes))
	}

	// Extract translation from the response
	// Google Translate returns a complex nested structure
	translation := extractTranslation(result)
	if translation == "" {
		return "", fmt.Errorf("could not extract translation from response: %s", string(bodyBytes))
	}

	return translation, nil
}

// extractTranslation extracts the translated text from Google Translate's response
func extractTranslation(data []interface{}) string {
	if len(data) == 0 {
		return ""
	}

	// The first element contains the translations
	translations, ok := data[0].([]interface{})
	if !ok {
		return ""
	}

	var translatedParts []string

	// Each element in translations is a pair [translated text, original text]
	for _, part := range translations {
		pair, ok := part.([]interface{})
		if !ok {
			continue
		}

		if len(pair) == 0 {
			continue
		}

		// The first element of the pair is the translated text
		translatedText, ok := pair[0].(string)
		if !ok {
			continue
		}

		translatedParts = append(translatedParts, translatedText)
	}

	return strings.Join(translatedParts, "")
}

// GetSupportedLanguages returns language codes supported by Google Translate
func GetSupportedLanguages() []string {
	return []string{
		"af", "sq", "am", "ar", "hy", "az", "eu", "be", "bn", "bs", "bg", "ca", "ceb", "zh-CN", "zh-TW",
		"co", "hr", "cs", "da", "nl", "en", "eo", "et", "fi", "fr", "fy", "gl", "ka", "de", "el", "gu",
		"ht", "ha", "haw", "he", "hi", "hmn", "hu", "is", "ig", "id", "ga", "it", "ja", "jv", "kn", "kk",
		"km", "rw", "ko", "ku", "ky", "lo", "la", "lv", "lt", "lb", "mk", "mg", "ms", "ml", "mt", "mi",
		"mr", "mn", "my", "ne", "no", "ny", "or", "ps", "fa", "pl", "pt", "pa", "ro", "ru", "sm", "gd",
		"sr", "st", "sn", "sd", "si", "sk", "sl", "so", "es", "su", "sw", "sv", "tl", "tg", "ta", "tt",
		"te", "th", "tr", "tk", "uk", "ur", "ug", "uz", "vi", "cy", "xh", "yi", "yo", "zu",
	}
}
