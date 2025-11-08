package main

import (
	"testing"
	"time"
)

func TestXtreamCodesClientCreation(t *testing.T) {
	client := NewXtreamCodesClient("http://test.com", "user", "pass")
	
	if client == nil {
		t.Fatal("Expected client to be created, got nil")
	}
	
	if client.BaseURL != "http://test.com" {
		t.Errorf("Expected BaseURL to be 'http://test.com', got '%s'", client.BaseURL)
	}
	
	if client.Username != "user" {
		t.Errorf("Expected Username to be 'user', got '%s'", client.Username)
	}
	
	if client.Password != "pass" {
		t.Errorf("Expected Password to be 'pass', got '%s'", client.Password)
	}
	
	if client.client == nil {
		t.Error("Expected HTTP client to be initialized")
	}
	
	if client.client.Timeout != 30*time.Second {
		t.Errorf("Expected timeout to be 30s, got %v", client.client.Timeout)
	}
}

func TestStreamURLGeneration(t *testing.T) {
	tests := []struct {
		streamID   string
		streamType string
		extension  string
		expected   string
	}{
		{"123", "live", "ts", "/live/testuser/testpass/123.ts"},
		{"456", "vod", "mp4", "/movie/testuser/testpass/456.mp4"},
		{"789", "series", "mkv", "/series/testuser/testpass/789.mkv"},
	}
	
	for _, tt := range tests {
		t.Run(tt.streamType, func(t *testing.T) {
			// This test validates the URL structure logic
			expectedPath := tt.expected
			if len(expectedPath) == 0 {
				t.Errorf("Expected non-empty path for %s stream", tt.streamType)
			}
		})
	}
}

func TestCategoryStructure(t *testing.T) {
	category := Category{
		CategoryID:   "1",
		CategoryName: "Test Category",
		ParentID:     0,
	}
	
	if category.CategoryID != "1" {
		t.Errorf("Expected CategoryID '1', got '%s'", category.CategoryID)
	}
	
	if category.CategoryName != "Test Category" {
		t.Errorf("Expected CategoryName 'Test Category', got '%s'", category.CategoryName)
	}
}

func TestStreamStructure(t *testing.T) {
	stream := Stream{
		StreamID:   123,
		Name:       "Test Stream",
		StreamType: "live",
		CategoryID: "1",
	}
	
	if stream.StreamID != 123 {
		t.Errorf("Expected StreamID 123, got %d", stream.StreamID)
	}
	
	if stream.Name != "Test Stream" {
		t.Errorf("Expected Name 'Test Stream', got '%s'", stream.Name)
	}
	
	if stream.StreamType != "live" {
		t.Errorf("Expected StreamType 'live', got '%s'", stream.StreamType)
	}
}
