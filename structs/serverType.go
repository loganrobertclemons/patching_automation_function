package structs

import (
	"github.com/ChrisHirsch/puppetdb-client-go"
	util "github.com/loganrobertclemons/patching_automation_function/functions"
)

type ServerType struct {
	IPAddress       string   `json:"ip"`
	Notes           string   `json:"notes"`
	OperatingSystem string   `json:"operating_system"`
	OSVersion       string   `json:"os_version"`
	PackageUpdates  int      `json:"package_updates"`
	PatchWindow     string   `json:"patch_window"`
	PinnedPackages  []string `json:"pinned_packages"`
	SecurityUpdates int      `json:"security_updates"`
}

func (s *ServerType) Init() {
	// var pkgs []string
	// s.PinnedPackages = pkgs
}

func (s *ServerType) parseInventoryResult(Server *puppetdb.Inventory) {
	s.IPAddress = util.GetFactString(Server.Facts["ipaddress"])
	s.OperatingSystem = util.GetFactString(Server.Facts["os"].(map[string]interface{})["name"])
	s.OSVersion = util.GetFactString(Server.Facts["os"].(map[string]interface{})["release"].(map[string]interface{})["full"])
	s.PackageUpdates = util.GetFactInt(Server.Facts["os_patching"].(map[string]interface{})["package_update_count"])
	s.PatchWindow = util.GetFactString(Server.Facts["os_patching"].(map[string]interface{})["patch_window"])
	s.PinnedPackages = util.GetFactArrayOfStrings(Server.Facts["os_patching"].(map[string]interface{})["pinned_packages"].([]interface{}))
	s.SecurityUpdates = util.GetFactInt(Server.Facts["os_patching"].(map[string]interface{})["security_package_update_count"])
}
