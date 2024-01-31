// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddItemToReceiptInput struct {
	ReceiptID string   `json:"receiptId"`
	Name      string   `json:"name"`
	Price     *float64 `json:"price,omitempty"`
}

type AssignUserToItemInput struct {
	ItemID string `json:"itemId"`
	UserID string `json:"userId"`
}

type Item struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Price    *float64 `json:"price,omitempty"`
	SharedBy []*User  `json:"sharedBy"`
}

type Mutation struct {
}

type Query struct {
}

type Receipt struct {
	ID          string   `json:"id"`
	OwnedBy     *User    `json:"ownedBy,omitempty"`
	Description string   `json:"description"`
	Total       *float64 `json:"total,omitempty"`
	Items       []*Item  `json:"items"`
}

type ReceiptInput struct {
	Description string   `json:"description"`
	Price       *float64 `json:"price,omitempty"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserInput struct {
	Username string `json:"username"`
}

type UserWithJwt struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}
