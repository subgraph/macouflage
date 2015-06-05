package main

import (
	"strings"
	"fmt"
	lmf "github.com/subgraph/libmacouflage"
)
func getMacInfo(name string) (result string, err error) {
	currentMac, err := lmf.GetCurrentMac(name)
	if err != nil {
		return
	}
	currentMacVendor, err := lmf.FindVendorByMac(currentMac.String())
	if err != nil {
		if strings.HasPrefix(err.Error(),
			"No vendor found in OuiDb for vendor prefix") {
			currentMacVendor.Vendor = "Unknown"
		} else {
			return
		}
	}
	permanentMac, err := lmf.GetPermanentMac(name)
	if err != nil {
		fmt.Println(err)
	}
	permanentMacVendor, err := lmf.FindVendorByMac(permanentMac.String())
	if err != nil {
		if strings.HasPrefix(err.Error(),
			"No vendor found in OuiDb for vendor prefix") {
			permanentMacVendor.Vendor = "Unknown"
		} else {
			return
		}
	}

	result = fmt.Sprintf("Current MAC: %s (%s)\nPermanent MAC: %s (%s)",
		currentMac, currentMacVendor.Vendor,
		permanentMac, permanentMacVendor.Vendor)
	return
}

func listVendors(keyword string, isPopular bool) (results string, err error) {
	var ouis []lmf.Oui
	var vendors []string

	if isPopular {
		ouis, err = lmf.FindAllPopularOuis()
		if err != nil {
			return
		}
	} else {
		ouis, err = lmf.FindVendorsByKeyword(keyword)
		if err != nil {
			return
		}
	}
	if len(ouis) == 0 {
		results = fmt.Sprintf("No vendors found in search.")
		return
	} else {
		vendors = append(vendors, fmt.Sprintf("#\tVendorPrefix\tVendor"))
		for i, result := range ouis {
			vendors = append(vendors, fmt.Sprintf("%d\t%s\t%s", i+1,
				result.VendorPrefix, result.Vendor))
		}
		results = strings.Join(vendors, "\n")
	}
	return
}
