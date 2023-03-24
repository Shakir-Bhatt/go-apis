package models

import "time"

type Product struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	BodyHTML          string    `json:"body_html"`
	Vendor            string    `json:"vendor"`
	ProductType       string    `json:"product_type"`
	CreatedAt         time.Time `json:"created_at"`
	Handle            string    `json:"handle"`
	UpdatedAt         time.Time `json:"updated_at"`
	PublishedAt       time.Time `json:"published_at"`
	TemplateSuffix    string    `json:"template_suffix"`
	Status            string    `json:"status"`
	PublishedScope    string    `json:"published_scope"`
	Tags              string    `json:"tags"`
	AdminGraphqlAPIID string    `json:"admin_graphql_api_id"`
	Variants          []struct {
		ID                   int64       `json:"id"`
		ProductID            int64       `json:"product_id"`
		Title                string      `json:"title"`
		Price                string      `json:"price"`
		Sku                  string      `json:"sku"`
		Position             int         `json:"position"`
		InventoryPolicy      string      `json:"inventory_policy"`
		CompareAtPrice       string      `json:"compare_at_price"`
		FulfillmentService   string      `json:"fulfillment_service"`
		InventoryManagement  interface{} `json:"inventory_management"`
		Option1              string      `json:"option1"`
		Option2              interface{} `json:"option2"`
		Option3              interface{} `json:"option3"`
		CreatedAt            time.Time   `json:"created_at"`
		UpdatedAt            time.Time   `json:"updated_at"`
		Taxable              bool        `json:"taxable"`
		Barcode              string      `json:"barcode"`
		Grams                int         `json:"grams"`
		ImageID              interface{} `json:"image_id"`
		Weight               int         `json:"weight"`
		WeightUnit           string      `json:"weight_unit"`
		InventoryItemID      int64       `json:"inventory_item_id"`
		InventoryQuantity    int         `json:"inventory_quantity"`
		OldInventoryQuantity int         `json:"old_inventory_quantity"`
		RequiresShipping     bool        `json:"requires_shipping"`
		AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
	} `json:"variants"`
	Options []struct {
		ID        int64    `json:"id"`
		ProductID int64    `json:"product_id"`
		Name      string   `json:"name"`
		Position  int      `json:"position"`
		Values    []string `json:"values"`
	} `json:"options"`
	Images []struct {
		ID                int64         `json:"id"`
		ProductID         int64         `json:"product_id"`
		Position          int           `json:"position"`
		CreatedAt         time.Time     `json:"created_at"`
		UpdatedAt         time.Time     `json:"updated_at"`
		Alt               interface{}   `json:"alt"`
		Width             int           `json:"width"`
		Height            int           `json:"height"`
		Src               string        `json:"src"`
		VariantIds        []interface{} `json:"variant_ids"`
		AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
	} `json:"images"`
	Image struct {
		ID                int64         `json:"id"`
		ProductID         int64         `json:"product_id"`
		Position          int           `json:"position"`
		CreatedAt         time.Time     `json:"created_at"`
		UpdatedAt         time.Time     `json:"updated_at"`
		Alt               interface{}   `json:"alt"`
		Width             int           `json:"width"`
		Height            int           `json:"height"`
		Src               string        `json:"src"`
		VariantIds        []interface{} `json:"variant_ids"`
		AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
	} `json:"image"`
}

type ProductWithImgae struct {
	ID       int64
	Title    string
	Vendor   string
	ImageUrl string
}
