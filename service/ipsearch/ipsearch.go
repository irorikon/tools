/*
 * @Author: iRorikon
 * @Date: 2023-04-18 11:14:47
 * @FilePath: \api-service\service\ipsearch\ipsearch.go
 */
package ipsearch

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ip2location/ip2location-go/v9"
	"github.com/ipipdotnet/ipdb-go"
	"github.com/irorikon/api-service/command/flags"
	"github.com/irorikon/api-service/model"
	"github.com/oschwald/geoip2-golang"
)

type IpSearch struct {
	DatabaseType string   `json:"type"`
	Database     string   `json:"database"`
	Online       bool     `json:"online"`
	IPs          []string `json:"ips"`
}

func NewIPSearch(databaseType, database string, online bool, ips []string) *IpSearch {
	return &IpSearch{
		DatabaseType: databaseType,
		Database:     database,
		Online:       online,
		IPs:          ips,
	}
}

func (i *IpSearch) SearchByDatabase() ([]model.IP, error) {
	if i.Online {
		// Todo: Online search
	} else if i.DatabaseType != "" {
		// Todo: Database search
		switch strings.ToLower(i.DatabaseType) {
		case "ipip":
			ipinfo, err := i.SearchByIPIP()
			if err != nil {
				return ipinfo, err
			}
			return ipinfo, err
		case "maxmind":
			ipinfo, err := i.SearchByMaxmind()
			if err != nil {
				return ipinfo, err
			}
			return ipinfo, err
		case "ip2location":
			ipinfo, err := i.SearchByIP2Location()
			if err != nil {
				return ipinfo, err
			}
			return ipinfo, err
		}
	} else {
		return nil, errors.New("no database type or online query specified")
	}
	return nil, nil
}

func (i *IpSearch) SearchByOnline() (ipinfos []model.IP, err error) {
	for _, ip := range i.IPs {
		// Todo: check if ip is valid
		ipaddr := net.ParseIP(ip)
		var res *http.Request
		searchURL := fmt.Sprintf("http://ip-api.com/json/%v?lang=zh-CN", ipaddr)
		res, err = http.NewRequest(http.MethodGet, searchURL, nil)
		if err != nil {
			return nil, err
		}
		client := &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 3,
				DisableKeepAlives:   false,
			},
			Timeout: time.Duration(30) * time.Second,
		}

		var resp *http.Response
		resp, err = client.Do(res)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := getResponseBody(resp)
		if err != nil {
			return nil, err
		}
		var ipApiResp model.IPAPTResp
		err = json.Unmarshal(body, &ipApiResp)
		if err != nil {
			return nil, err
		}
		ipinfos = append(ipinfos, model.IP{
			IP:          ipApiResp.Query,
			CountryCode: ipApiResp.CountryCode,
			CountryName: ipApiResp.Country,
			RegionName:  ipApiResp.RegionName,
			CityName:    ipApiResp.City,
			Latitude:    strconv.FormatFloat(ipApiResp.Lat, 'f', 6, 64),
			Longitude:   strconv.FormatFloat(ipApiResp.Lon, 'f', 6, 64),
			AS:          ipApiResp.AS,
			ISP:         ipApiResp.ISP,
		})
	}
	return
}

func getResponseBody(resp *http.Response) (body []byte, err error) {
	var output io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		output, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
		if err != nil {
			return
		}
	default:
		output = resp.Body
		if err != nil {
			return
		}
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(output)
	if err != nil {
		return
	}
	body = buf.Bytes()
	return
}

// Search By IPIP
func (i *IpSearch) SearchByIPIP() (ipinfo []model.IP, err error) {
	db, err := ipdb.NewCity(flags.Database)
	if err != nil {
		return nil, err
	}
	for _, ip := range i.IPs {
		ip = strings.TrimSpace(ip)
		iplist := strings.Split(ip, "/")
		if iplist[0] == "" {
			continue
		}
		ipAddress := net.ParseIP(iplist[0])
		if ipAddress == nil {
			ipinfo = append(ipinfo, model.IP{
				Error: "IP format is incorrect",
			})
			continue
		} else {
			ipdata, err := db.FindInfo(iplist[0], "CN")
			if err != nil {
				return ipinfo, err
			}
			ipinfo = append(ipinfo, model.IP{
				IP:          ip,
				CountryCode: ipdata.CountryCode,
				CountryName: ipdata.CountryName,
				RegionName:  ipdata.RegionName,
				CityName:    ipdata.CityName,
				Latitude:    ipdata.Latitude,
				Longitude:   ipdata.Longitude,
				ASN:         ipdata.ASN,
				IDC:         ipdata.IDC,
				ISP:         ipdata.IspDomain,
			})
		}

	}
	return
}

// Search By Maxmind
func (i *IpSearch) SearchByMaxmind() (ipinfo []model.IP, err error) {
	db, err := geoip2.Open(flags.Database)
	if err != nil {
		return ipinfo, err
	}
	defer db.Close()
	for _, ip := range i.IPs {
		ip = strings.TrimSpace(ip)
		iplist := strings.Split(ip, "/")
		if iplist[0] == "" {
			continue
		}
		ipAddress := net.ParseIP(iplist[0])
		if ipAddress == nil {
			ipinfo = append(ipinfo, model.IP{
				Error: "IP format is incorrect",
			})
			continue
		} else {
			ipAddr := net.ParseIP(iplist[0])
			record, err := db.City(ipAddr)
			if err != nil {
				return ipinfo, err
			}
			var countryName string
			var cityName string
			if record.Country.Names["zh-CN"] != "" {
				countryName = record.Country.Names["zh-CN"]
			} else {
				countryName = record.Country.Names["en"]
			}
			if record.City.Names["zh-CN"] != "" {
				cityName = record.City.Names["zh-CN"]
			} else {
				cityName = record.City.Names["en"]
			}
			ipinfo = append(ipinfo, model.IP{
				IP:          ip,
				CountryCode: record.Continent.Code,
				CountryName: countryName,
				CityName:    cityName,
				Latitude:    strconv.FormatFloat(record.Location.Latitude, 'f', 6, 64),
				Longitude:   strconv.FormatFloat(record.Location.Longitude, 'f', 6, 64),
			})
		}
	}
	return
}

// Search By IP2Location
func (i *IpSearch) SearchByIP2Location() (ipinfo []model.IP, err error) {
	db, err := ip2location.OpenDB(flags.Database)
	if err != nil {
		return
	}
	defer db.Close()
	for _, ip := range i.IPs {
		ip = strings.TrimSpace(ip)
		iplist := strings.Split(ip, "/")
		if iplist[0] == "" {
			continue
		}
		ipAddress := net.ParseIP(iplist[0])
		if ipAddress == nil {
			ipinfo = append(ipinfo, model.IP{
				Error: "IP format is incorrect",
			})
			continue
		} else {
			results, err := db.Get_all(iplist[0])
			if err != nil {
				return ipinfo, err
			}
			ipinfo = append(ipinfo, model.IP{
				IP:          ip,
				CountryCode: results.Country_short,
				CountryName: results.Country_long,
				RegionName:  results.Region,
				CityName:    results.City,
				Latitude:    strconv.FormatFloat(float64(results.Latitude), 'f', 6, 64),
				Longitude:   strconv.FormatFloat(float64(results.Longitude), 'f', 6, 64),
				ISP:         results.Isp,
			})
		}
	}
	return
}
