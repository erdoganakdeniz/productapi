package models

type Product struct {
	Barcode           string  `json:"barcode" bson:"barcode"`
	Title             string  `json:"title" bson:"title"`
	ProductMainId     string  `json:"productMainId" bson:"product_main_id"`
	BrandId           int     `json:"brandId" bson:"brand_id"`
	CategoryId        int     `json:"categoryId" bson:"category_id"`
	Quantity          int     `json:"quantity" bson:"quantity"`
	StockCode         string  `json:"stockCode" bson:"stock_code"`
	DimensionalWeight int     `json:"dimensionalWeight" bson:"dimensional_weight"`
	Description       string  `json:"description" bson:"description"`
	CurrencyType      string  `json:"currencyType" bson:"currency_type"`
	ListPrice         float64 `json:"listPrice" bson:"list_price"`
	SalePrice         float64 `json:"salePrice" bson:"sale_price"`
	VatRate           int     `json:"vatRate" bson:"vat_rate"`
	CargoCompanyId    int     `json:"cargoCompanyId" bson:"cargo_company_id"`
	Images            []struct {
		Url string `json:"url" bson:"url"`
	} `json:"images" bson:"images"`
	Attributes []struct {
		AttributeId          int    `json:"attributeId" bson:"attribute_id"`
		AttributeValueId     int    `json:"attributeValueId,omitempty" bson:"attribute_value_id"`
		CustomAttributeValue string `json:"customAttributeValue,omitempty" bson:"custom_attribute_value"`
	} `json:"attributes" bson:"attributes"`
}
