/*
 * @Author: iRorikon
 * @Date: 2023-04-18 11:19:58
 * @FilePath: \api-service\model\ip.go
 */
package model

type IP struct {
	IP          string `json:"ip"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	RegionName  string `json:"region_name"`
	CityName    string `json:"city_name"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	AS          string `json:"as"`
	ASN         string `json:"asn"`
	IDC         string `json:"idc"`
	ISP         string `json:"isp"`
	Error       string `json:"error"`
}

type IPAPTResp struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
}
