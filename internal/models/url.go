package models

import "time"

type Link struct {
	Url        string    `json:"url,omitempty"`
	Alias      string    `json:"alias,omitempty"`
	ExpireTime time.Time `json:"expire_time"`
}

type UrlStats struct {
	Clicks         int       `json:"clicks"`
	LastAccessTime time.Time `json:"last_access_time"`
}
