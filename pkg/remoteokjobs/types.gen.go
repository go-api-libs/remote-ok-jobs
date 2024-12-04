// This file provides types for the API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package remoteokjobs

import (
	"net/url"
	"time"
)

// GetAPIOkJSONResponse defines a model
type GetAPIOkJSONResponse []GetAPIOkJSONResponseItems

// GetAPIOkJSONResponseItems defines a model
type GetAPIOkJSONResponseItems struct {
	LastUpdated int        `json:"last_updated"`
	Legal       string     `json:"legal"`
	Slug        *string    `json:"slug"`
	ID          *string    `json:"id"`
	Epoch       *int       `json:"epoch"`
	Date        *time.Time `json:"date"`
	Company     *string    `json:"company"`
	CompanyLogo *url.URL   `json:"company_logo"`
	Position    *string    `json:"position"`
	Tags        *[]string  `json:"tags"`
	Description *string    `json:"description"`
	Location    *string    `json:"location"`
	SalaryMin   *int       `json:"salary_min"`
	SalaryMax   *int       `json:"salary_max"`
	ApplyURL    *url.URL   `json:"apply_url"`
	Logo        *url.URL   `json:"logo"`
	URL         *url.URL   `json:"url"`
	Original    *bool      `json:"original"`
	Verified    *bool      `json:"verified"`
}