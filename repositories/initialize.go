package repositories

import (
	"log"
	"time"

	"github.com/digitalocean/godo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func init() {

	db, err := gorm.Open(sqlite.Open("./database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Droplet{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func CreateDroplets(droplets []godo.Droplet) {
	// Record all droplets
	for _, droplet := range droplets {
		log.Println(droplet.Name)
		d := Droplet{
			ID:     droplet.ID,
			Name:   droplet.Name,
			Memory: droplet.Memory,
			Vcpus:  droplet.Vcpus,
			Disk:   droplet.Disk,
			Region: droplet.Region.Slug,
			//Size:      droplet.Size,
			SizeSlug: droplet.SizeSlug,
			Status:   droplet.Status,
			Created:  droplet.Created,
			//VolumeIDs: droplet.VolumeIDs,
		}
		if len(droplet.Networks.V4) > 0 {
			d.NetworkV4IP = droplet.Networks.V4[0].IPAddress
		}

		DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&d)
	}

}

type Droplets struct {
	droplets []Droplet
}

type Droplet struct {
	gorm.Model
	ID     int    `json:"id,float64,omitempty"`
	Name   string `json:"name,omitempty"`
	Memory int    `json:"memory,omitempty"`
	Vcpus  int    `json:"vcpus,omitempty"`
	Disk   int    `json:"disk,omitempty"`
	Region string `json:"region,omitempty"`
	//Image     *Image `json:"image,omitempty"`
	//Size      *Size  `json:"size,omitempty"`
	SizeSlug string `json:"size_slug,omitempty"`
	//BackupIDs []int  `json:"backup_ids,omitempty"`
	//NextBackupWindow *BackupWindow `json:"next_backup_window,omitempty"`
	//SnapshotIDs []int    `json:"snapshot_ids,omitempty"`
	//Features    []string `json:"features,omitempty"`
	Locked bool   `json:"locked,bool,omitempty"`
	Status string `json:"status,omitempty"`
	//Networks         *Networks     `json:"networks,omitempty"`
	NetworkV4IP string `json:"network_v4_ip"`
	Created     string `json:"created_at,omitempty"`
	//Kernel           *Kernel       `json:"kernel,omitempty"`
	//Tags      []string `json:"tags,omitempty"`
	//VolumeIDs []string `json:"volume_ids"`
	//VPCUUID   string   `json:"vpc_uuid,omitempty"`
}

type Region struct {
	Slug      string   `json:"slug,omitempty"`
	Name      string   `json:"name,omitempty"`
	Sizes     []string `json:"sizes,omitempty"`
	Available bool     `json:"available,omitempty"`
	Features  []string `json:"features,omitempty"`
}
type Image struct {
	ID            int      `json:"id,float64,omitempty"`
	Name          string   `json:"name,omitempty"`
	Type          string   `json:"type,omitempty"`
	Distribution  string   `json:"distribution,omitempty"`
	Slug          string   `json:"slug,omitempty"`
	Public        bool     `json:"public,omitempty"`
	Regions       []string `json:"regions,omitempty"`
	MinDiskSize   int      `json:"min_disk_size,omitempty"`
	SizeGigaBytes float64  `json:"size_gigabytes,omitempty"`
	Created       string   `json:"created_at,omitempty"`
	Description   string   `json:"description,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Status        string   `json:"status,omitempty"`
	ErrorMessage  string   `json:"error_message,omitempty"`
}
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	PriceMonthly float64  `json:"price_monthly,omitempty"`
	PriceHourly  float64  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	Available    bool     `json:"available,omitempty"`
	Transfer     float64  `json:"transfer,omitempty"`
}
type BackupWindow struct {
	Start *Timestamp `json:"start,omitempty"`
	End   *Timestamp `json:"end,omitempty"`
}
type Networks struct {
	V4 []NetworkV4 `json:"v4,omitempty"`
	V6 []NetworkV6 `json:"v6,omitempty"`
}
type Kernel struct {
	ID      int    `json:"id,float64,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
type Timestamp struct {
	time.Time
}
type NetworkV4 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}
type NetworkV6 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   int    `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}
