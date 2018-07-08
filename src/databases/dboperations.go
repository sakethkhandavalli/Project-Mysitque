package data

import (
	"time"
)

const BLOCKLIMIT = 8

type BlockInfo struct {
	BlockId          string       `json:"blockid"`
	PackageId        string       `json:"packageid , omitempty"`
	Status           string       `json:"packagestatus , omitempty"`
	ConnectedSensors int          `json:"connectedsensors ,omitempty"`
	SensorInfo       []SensorInfo `json:"sensorinfo , omitempty"`
	GeoLocation      string       `json:"geohash , omitempty"`
	Data             string       `json: "blockdata , omitempty"`
	LastUpdated      time.Time    `json:"lastupdated"`
}

type SensorInfo struct {
	SensorId     string  `json:"sensorid"`
	Status       string  `json:"sensorstatus"`
	CurrentValue float64 `json:"currentvalue"`
	ValueUnit    string  `json:"valuetype"`
}

type PackageInfo struct {
	PackageId        string       `json:"packageid"`
	Status           string       `json:"packagestatus"`
	ConnectedSensors int          `json:"connectedsensors ,omitempty"`
	SensorInfo       []SensorInfo `json:"sensorinfo , omitempty"`
	GeoLocation      string       `json:"geohash , omitempty"`
	BlocksNumber     int          `json:"numberofblocks"`
	BlockId          []string     `json:"blockid"`
	UpdateFrequency  int          `json:"numberoftimes "`
	LastUpdated      []time.Time  `json:"lastupdated"`
}

var packagelist []PackageInfo
var blocklist []BlockInfo

func InitBlockList() {
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j49902934nesf", PackageId: "2987K", Status: "ALERT", ConnectedSensors: 0, GeoLocation: "Hyderabad", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j49909sdmmgke", PackageId: "2987J", Status: "ALERT", ConnectedSensors: 0, GeoLocation: "Hyderabad", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j4999dd923042", PackageId: "2987I", Status: "ACTIVE", ConnectedSensors: 0, GeoLocation: "Hyderabad", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j499234msd984", PackageId: "2987H", Status: "ACTIVE", ConnectedSensors: 0, GeoLocation: "Hyderabad", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j4993254mdfls", PackageId: "2987G", Status: "INACTIVE", ConnectedSensors: 0, GeoLocation: "Delhi", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j499hyfisndkf", PackageId: "2987F", Status: "INACTIVE", ConnectedSensors: 0, GeoLocation: "Delhi", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j499jhnkslwif", PackageId: "2987E", Status: "ALERT", ConnectedSensors: 0, GeoLocation: "Delhi", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j499lkjuhnkjs", PackageId: "2987D", Status: "INACTIVE", ConnectedSensors: 0, GeoLocation: "Delhi", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n1234njuihgbasdfnasds", PackageId: "2987C", Status: "ACTIVE", ConnectedSensors: 0, GeoLocation: "Lahore", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "asb30345n12341b3j499asdfnasd", PackageId: "2987B", Status: "INACTIVE", ConnectedSensors: 0, GeoLocation: "Lahore", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "jijadfn9n12341b3j499asdfnasd", PackageId: "2987B", Status: "INACTIVE", GeoLocation: "Lahore", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "3534frdgn12341b3j499asdfnasd", Status: "ALERT", GeoLocation: "Bangalore", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "sdgfsd43n12341b3j499asdfnasd", Status: "ALERT", GeoLocation: "Bangalore", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "htrr4we4n12341b3j499asdfnasd",Status: "ACTIVE",  GeoLocation: "Bangalore", LastUpdated: time.Now()})
	blocklist = append(blocklist, BlockInfo{BlockId: "876834fsn12341b3j499asdfnasd",Status: "ACTIVE",  GeoLocation: "Bangalore", LastUpdated: time.Now()})

}

func InitPackageList() {
	packagelist = append(packagelist, PackageInfo{PackageId: "2987B", Status: "INACTIVE", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j499asdfnasd"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987C", Status: "ACTIVE", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n1234njuihgbasdfnasds"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987D", Status: "INACTIVE", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j499lkjuhnkjs"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987E", Status: "ALERT", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j499jhnkslwif"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987F", Status: "INACTIVE", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j499hyfisndkf"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987G", Status: "INACTIVE", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j4993254mdfls"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987H", Status: "ACTIVE", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j499234msd984"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987I", Status: "ACTIVE", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j4999dd923042"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987J", Status: "ALERT", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j49909sdmmgke"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
	packagelist = append(packagelist, PackageInfo{PackageId: "2987K", Status: "ALERT", ConnectedSensors: 0, BlocksNumber: 1, BlockId: []string{"asb30345n12341b3j49902934nesf"}, UpdateFrequency: 1, LastUpdated: []time.Time{time.Now()}})
}
func GetPackage() *PackageInfo {
	packageinfo := &PackageInfo{}
	packageinfo.PackageId = "2987A"
	packageinfo.Status = "ACTIVE"
	packageinfo.ConnectedSensors = 0
	packageinfo.BlockId = []string{"asb30345n12341b3j499asdfnasd"}
	packageinfo.LastUpdated = []time.Time{time.Now()}

	return packageinfo
}

func GetAllPackage() []PackageInfo {
	return packagelist
}

func GetAllBlock() []BlockInfo {
	return blocklist
}
func GetBlock() BlockInfo {
	var blockinfo BlockInfo
	return blockinfo
}
func GetBlockInfoStructById(blockid string) BlockInfo {
	for _, item := range blocklist {
		if item.BlockId == blockid {
			return item
		}
	}
	blockinfo := BlockInfo{}
	return blockinfo

}

func GetAllBlocksForAPackage(packageId string) []BlockInfo {
	blocksListForPackage := []BlockInfo{}
	for _, item := range blocklist {
		if item.PackageId == packageId {
			blocksListForPackage = append(blocksListForPackage, item)
		}
	}
	return blocksListForPackage
}

func GetRecentBlocksForThePackage(packageId string) []BlockInfo {
	blocksListForPackage := []BlockInfo{}
	var count = 0
	for _, item := range blocklist {
		if item.PackageId == packageId {
			blocksListForPackage = append(blocksListForPackage, item)
			count++
		}
		if count == BLOCKLIMIT {
			return blocksListForPackage
		}
	}
	return blocksListForPackage
}
