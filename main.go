package main

import (
	"fmt"
	"log"
)

func main() {
	printBanner()
	takeInput()
	if len(targets) > 0 {
		log.Println("Genzai is starting...")
		log.Println("Loading Genzai DB...")
		loadDB()
		log.Println("Loading Vendor DB...")
		loadVendorDB()
		fmt.Println("\n ")
	}

	genzaiOutput.Targets = targets

	for _, target := range targets {
		fmt.Println()
		log.Println("Starting the scan for", target)
		product := targetDetection(target)
		if product != "" {
			var targetResult genzaiResult
			targetResult.Target = target
			targetResult.IoTidentified = product
			log.Println("IoT Dashboard Discovered:", product)
			log.Println("Trying for default vendor-specific [", product, "] passwords...")
			passIssue := vendorpassScan(target, product)
			if passIssue.URL != "" {
				targetResult.Issues = append(targetResult.Issues, passIssue)
			}

			genzaiOutput.Results = append(genzaiOutput.Results, targetResult)
		}
	}

	/*
		for key, entry := range genzaiDB {
			fmt.Printf("Entry %s:\n", key)
			fmt.Println("Matches:", entry.Matchers.Strings)
			fmt.Println("Response Code:", entry.Matchers.ResponseCode)

			fmt.Println("Headers:")
			for headerKey, headerValue := range entry.Matchers.Headers {
				fmt.Printf("  %s: %v\n", headerKey, headerValue)
			}
			fmt.Println()
		}
	*/
	generateOutput()
}