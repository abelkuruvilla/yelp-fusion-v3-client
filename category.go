package yelp

import "fmt"

type Category struct {
	Alias            string   `json:"alias,omitempty"`
	Title            string   `json:"title,omitempty"`
	ParentAliases    []string `json:"parent_aliases,omitempty"`
	CountryWhitelist []string `json:"country_whitelist,omitempty"`
	CountryBlacklist []string `json:"country_blacklist,omitempty"`
}

type CategoriesSearchResponse struct {
	Categories []Category `json:"categories,omitempty"`
}
type CategoryDetailResponse struct {
	Category Category `json:"category,omitempty"`
}
type CategoriesSearchParam struct {
	Locale string `json:"locale,omitempty"`
}

// CategoriesList is used to fetch the list of categories from Yelp.
// See https://www.yelp.com/developers/documentation/v3/all_categories
func (client *Client) CategoriesSearch(request CategoriesSearchParam) (response CategoriesSearchResponse, err error) {
	url := fmt.Sprintf("/categories")

	if err := client.query(url, &request, &response); err != nil {
		return response, err
	}
	return response, nil

}

// GetCategory is used to get detailed information about a category from YELP API
// See https://www.yelp.com/developers/documentation/v3/category
func (client *Client) GetCategory(alias string, request CategoriesSearchParam) (response CategoryDetailResponse, err error) {
	url := fmt.Sprintf("/categories/%s", alias)

	if err := client.query(url, &request, &response); err != nil {
		return response, err
	}
	return response, nil

}
