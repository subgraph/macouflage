package main

import (
	"strings"
	"fmt"
	lmf "github.com/subgraph/libmacouflage"
)

func getCurrentMacInfo(name string) (result string, err error) {
	currentMacInfo, err := getMacInfo(name, "CurrentMAC")
	if err != nil {
		return
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
	result = fmt.Sprintf("%sPermanent MAC:\t%s (%s)",
		currentMacInfo,
		permanentMac, permanentMacVendor.Vendor)
	return
}

func getMacInfo(name string, macType string) (result string, err error) {
	newMac, err := lmf.GetCurrentMac(name)
	if err != nil {
		return
	}
	newMacVendor, err := lmf.FindVendorByMac(newMac.String())
	if err != nil {
		if err == err.(lmf.NoVendorError) {
			newMacVendor.Vendor = "Unknown"
			err = nil
		} else {
			return
		}
	}
	result = fmt.Sprintf("%s:\t%s (%s)\n",
		macType, newMac, newMacVendor.Vendor)
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

func spoofMacEnding(name string) (err error) {
	currentMacInfo, err := getCurrentMacInfo(name)
	if err != nil {
		return
	}
	fmt.Println(currentMacInfo)
	changed, err := lmf.SpoofMacSameVendor(name, true)
	if err != nil {
		return
	}
	if changed {
		newMac, err2 := getMacInfo(name, "New MAC")
		if err2 != nil {
			err = err2
			return
		}
		fmt.Printf(newMac)
	}
	return
}

func spoofMacAnother(name string) (err error) {
	currentMacInfo, err := getCurrentMacInfo(name)
	if err != nil {
		return
	}
	fmt.Println(currentMacInfo)
	changed, err := lmf.SpoofMacSameDeviceType(name)
	if err != nil {
		return
	}
	if changed {
		newMac, err2 := getMacInfo(name, "New MAC")
		if err2 != nil {
			err = err2
			return
		}
		fmt.Printf(newMac)
	}
	return
}

func spoofMacAny(name string) (err error) {
	currentMacInfo, err := getCurrentMacInfo(name)
	if err != nil {
		return
	}
	fmt.Println(currentMacInfo)
	changed, err := lmf.SpoofMacAnyDeviceType(name)
	if err != nil {
		return
	}
	if changed {
		newMac, err2 := getMacInfo(name, "New MAC")
		if err2 != nil {
			err = err2
			return
		}
		fmt.Printf(newMac)
	}
	return
}

func revertMac(name string) (err error) {
	currentMacInfo, err := getCurrentMacInfo(name)
	if err != nil {
		return
	}
	fmt.Println(currentMacInfo)
	permMac, err := lmf.GetPermanentMac(name)
	if err != nil {
		return
	}
	err = lmf.RevertMac(name)
	if err != nil {
		return
	}
	newMac, err := lmf.GetCurrentMac(name)
	if err != nil {
		return
	}
	if lmf.CompareMacs(permMac, newMac) {
		newMac, err2 := getMacInfo(name, "New MAC")
		if err2 != nil {
			err = err2
			return
		}
		fmt.Printf(newMac)
	}
	return
}

func spoofMacRandom(name string, bia bool) (err error) {
	currentMacInfo, err := getCurrentMacInfo(name)
	if err != nil {
		return
	}
	fmt.Println(currentMacInfo)
	changed, err := lmf.SpoofMacRandom(name, bia)
	if err != nil {
		return
	}
	if changed {
		newMac, err2 := getMacInfo(name, "New MAC")
		if err2 != nil {
			err = err2
			return
		}
		fmt.Printf(newMac)
	}
	return
}

