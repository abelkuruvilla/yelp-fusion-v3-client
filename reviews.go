package yelp

import "fmt"

type Review struct {
	ID          string `json:"id,omitempty"`
	Text        string `json:"text,omitempty"`
	URL         string `json:"url,omitempty"`
	Rating      int    `json:"rating,omitempty"`
	TimeCreated string `json:"time_created,omitempty"`
	User        User   `json:"user,omitempty"`
}

type User struct {
	ID         string `json:"id,omitempty"`
	ProfileURL string `json:"profile_url,omitempty"`
	Name       string `json:"name,omitempty"`
	ImageURL   string `json:"image_url,omitempty"`
}

type ReviewResponse struct {
	Total             uint     `json:"total,omitempty"`
	PossibleLanguages []string `json:"possible_languages,omitempty"`
	Reviews           []Review `json:"reviews,omitempty"`
}

// ReviewParams represents the query parameters that can be used while fetching a review
type ReviewParams struct {
	Locale string `json:"locale,omitempty"`
}

// addReviews is used to add the reviews of a business to the struct
func (b *Business) addReviews(client *Client, query ReviewParams) error {

	resp, err := client.FetchReviews(b.ID, query)
	if err != nil {
		return err
	}
	b.Reviews = resp
	return nil

}

// FetchReviews is used to fetch the reviews of a business from Yelp. Use the ReviewParam if you want to specify the query params.
// More Documentation at https://www.yelp.com/developers/documentation/v3/business_reviews
func (client *Client) FetchReviews(id string, query ReviewParams) (response ReviewResponse, err error) {

	url := fmt.Sprintf("/businesses/%s/reviews", id)

	if err := client.query(url, &query, &response); err != nil {
		return response, err
	}
	return response, nil
}
