package model

type Promotion struct {
	Id                      string         `json:"Id"`
	Product                 interface{}    `json:"Product"`
	Price                   float64        `json:"Price"`
	OldPrice                float64        `json:"OldPrice"`
	PromotionGainText       string         `json:"PromotionGainText"`
	PromotionGainTextLines  []string       `json:"PromotionGainTextLines"`
	PromotionGainValue      float64        `json:"PromotionGainValue"`
	PromotionTotalCount     int            `json:"PromotionTotalCount"`
	SaleCount               int            `json:"SaleCount"`
	MaxBuyCount             int            `json:"MaxBuyCount"`
	PromotionName           string         `json:"PromotionName"`
	PromotionEnd            string         `json:"PromotionEnd"`
	HtmlContent             interface{}    `json:"HtmlContent"`
	PromotionPhoto          PromotionPhoto `json:"PromotionPhoto"`
	IsActive                bool           `json:"IsActive"`
	IsSuspended             bool           `json:"IsSuspended"`
	MinimumInstallmentValue float64        `json:"MinimumInstallmentValue"`
}

type PromotionPhoto struct {
	Url          string      `json:"Url"`
	ThumbnailUrl string      `json:"ThumbnailUrl"`
	UrlTemplate  interface{} `json:"UrlTemplate"`
}
