package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// XtreamCodesClient handles communication with Xtream Codes API
type XtreamCodesClient struct {
	BaseURL  string
	Username string
	Password string
	client   *http.Client
}

// AuthResponse represents authentication response
type AuthResponse struct {
	UserInfo struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		Message        string `json:"message"`
		Auth           int    `json:"auth"`
		Status         string `json:"status"`
		ExpDate        string `json:"exp_date"`
		IsTrial        string `json:"is_trial"`
		ActiveCons     string `json:"active_cons"`
		CreatedAt      string `json:"created_at"`
		MaxConnections string `json:"max_connections"`
	} `json:"user_info"`
	ServerInfo struct {
		URL         string `json:"url"`
		Port        string `json:"port"`
		HTTPSPort   string `json:"https_port"`
		ServerProto string `json:"server_protocol"`
		RTMPPort    string `json:"rtmp_port"`
		Timezone    string `json:"timezone"`
	} `json:"server_info"`
}

// Category represents a content category
type Category struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentID     int    `json:"parent_id"`
}

// Stream represents a live stream
type Stream struct {
	Num              int    `json:"num"`
	Name             string `json:"name"`
	StreamType       string `json:"stream_type"`
	StreamID         int    `json:"stream_id"`
	StreamIcon       string `json:"stream_icon"`
	EPGChannelID     string `json:"epg_channel_id"`
	Added            string `json:"added"`
	CategoryID       string `json:"category_id"`
	CustomSid        string `json:"custom_sid"`
	TVArchive        int    `json:"tv_archive"`
	DirectSource     string `json:"direct_source"`
	TVArchiveDuration int   `json:"tv_archive_duration"`
}

// VODInfo represents video on demand info
type VODInfo struct {
	StreamID       int    `json:"stream_id"`
	Name           string `json:"name"`
	Added          string `json:"added"`
	CategoryID     string `json:"category_id"`
	ContainerExtension string `json:"container_extension"`
	CustomSid      string `json:"custom_sid"`
	DirectSource   string `json:"direct_source"`
}

// SeriesInfo represents TV series info
type SeriesInfo struct {
	SeriesID   int    `json:"series_id"`
	Name       string `json:"name"`
	Cover      string `json:"cover"`
	Plot       string `json:"plot"`
	Cast       string `json:"cast"`
	Director   string `json:"director"`
	Genre      string `json:"genre"`
	ReleaseDate string `json:"releaseDate"`
	Rating     string `json:"rating"`
	CategoryID string `json:"category_id"`
}

// NewXtreamCodesClient creates a new Xtream Codes API client
func NewXtreamCodesClient(baseURL, username, password string) *XtreamCodesClient {
	return &XtreamCodesClient{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Authenticate authenticates with the Xtream Codes server
func (x *XtreamCodesClient) Authenticate() (*AuthResponse, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s", x.BaseURL, x.Username, x.Password)
	
	resp, err := x.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, err
	}

	return &authResp, nil
}

// GetLiveStreams gets live streams
func (x *XtreamCodesClient) GetLiveStreams(categoryID string) ([]Stream, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_live_streams", 
		x.BaseURL, x.Username, x.Password)
	
	if categoryID != "" {
		url += "&category_id=" + categoryID
	}

	resp, err := x.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var streams []Stream
	if err := json.NewDecoder(resp.Body).Decode(&streams); err != nil {
		return nil, err
	}

	return streams, nil
}

// GetLiveCategories gets live stream categories
func (x *XtreamCodesClient) GetLiveCategories() ([]Category, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_live_categories",
		x.BaseURL, x.Username, x.Password)

	resp, err := x.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var categories []Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetVODStreams gets VOD streams
func (x *XtreamCodesClient) GetVODStreams(categoryID string) ([]VODInfo, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_vod_streams",
		x.BaseURL, x.Username, x.Password)

	if categoryID != "" {
		url += "&category_id=" + categoryID
	}

	resp, err := x.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var vods []VODInfo
	if err := json.NewDecoder(resp.Body).Decode(&vods); err != nil {
		return nil, err
	}

	return vods, nil
}

// GetVODCategories gets VOD categories
func (x *XtreamCodesClient) GetVODCategories() ([]Category, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_vod_categories",
		x.BaseURL, x.Username, x.Password)

	resp, err := x.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var categories []Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetSeries gets TV series
func (x *XtreamCodesClient) GetSeries(categoryID string) ([]SeriesInfo, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_series",
		x.BaseURL, x.Username, x.Password)

	if categoryID != "" {
		url += "&category_id=" + categoryID
	}

	resp, err := x.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var series []SeriesInfo
	if err := json.NewDecoder(resp.Body).Decode(&series); err != nil {
		return nil, err
	}

	return series, nil
}

// GetSeriesCategories gets series categories
func (x *XtreamCodesClient) GetSeriesCategories() ([]Category, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_series_categories",
		x.BaseURL, x.Username, x.Password)

	resp, err := x.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var categories []Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// ProxyStream proxies a stream from Xtream Codes
func (x *XtreamCodesClient) ProxyStream(streamID, streamType, extension string) (io.ReadCloser, string, error) {
	var url string
	switch streamType {
	case "live":
		url = fmt.Sprintf("%s/live/%s/%s/%s.%s", x.BaseURL, x.Username, x.Password, streamID, extension)
	case "vod":
		url = fmt.Sprintf("%s/movie/%s/%s/%s.%s", x.BaseURL, x.Username, x.Password, streamID, extension)
	case "series":
		url = fmt.Sprintf("%s/series/%s/%s/%s.%s", x.BaseURL, x.Username, x.Password, streamID, extension)
	default:
		return nil, "", fmt.Errorf("invalid stream type: %s", streamType)
	}

	resp, err := x.client.Get(url)
	if err != nil {
		return nil, "", err
	}

	contentType := resp.Header.Get("Content-Type")
	return resp.Body, contentType, nil
}

var client *XtreamCodesClient

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	auth, err := client.Authenticate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth)
}

func liveCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	categories, err := client.GetLiveCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func liveStreamsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	categoryID := r.URL.Query().Get("category_id")
	streams, err := client.GetLiveStreams(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(streams)
}

func vodCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	categories, err := client.GetVODCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func vodStreamsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	categoryID := r.URL.Query().Get("category_id")
	vods, err := client.GetVODStreams(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vods)
}

func seriesCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	categories, err := client.GetSeriesCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func seriesStreamsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	categoryID := r.URL.Query().Get("category_id")
	series, err := client.GetSeries(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}

func streamProxyHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	streamID := vars["id"]
	streamType := r.URL.Query().Get("type")
	extension := r.URL.Query().Get("ext")

	if extension == "" {
		extension = "ts"
	}

	body, contentType, err := client.ProxyStream(streamID, streamType, extension)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer body.Close()

	w.Header().Set("Content-Type", contentType)
	io.Copy(w, body)
}

func main() {
	// Initialize client with environment variables or default values
	baseURL := "http://example.com:8080"
	username := "demo"
	password := "demo"

	client = NewXtreamCodesClient(baseURL, username, password)

	r := mux.NewRouter()

	// API endpoints
	r.HandleFunc("/api/auth", authHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/live/categories", liveCategoriesHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/live/streams", liveStreamsHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/vod/categories", vodCategoriesHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/vod/streams", vodStreamsHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/series/categories", seriesCategoriesHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/series/streams", seriesStreamsHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stream/{id}", streamProxyHandler).Methods("GET", "OPTIONS")

	// Serve static files for Tizen app
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./tizen-app")))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
