package main

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
    DB_USER     = "pubguser"
    DB_PASSWORD = "Pubg$1234"
    DB_NAME     = "dev"
    DB_HOSTNAME = "sm-achawan-l01-1.tailc73dd.ts.net"
)

type wrapped_data struct {
    TotalAssists                 int64   `gorm:"Total_assists"`
    TotalChickens                int64   `gorm:"Total_chickens"`
    TotalCnt                     int64   `gorm:"Total_cnt"`
    TotalKills                   int64   `gorm:"Total_kills"`
    TotalTs                      float64 `gorm:"Total_ts"`
    AccountID                    string  `gorm:"accountID"`
    AvgKills                     int64   `gorm:"avg_kills"`
    AvgMeleeKills                int64   `gorm:"avg_melee_kills"`
    AvgMostWeaponKills           float64 `gorm:"avg_most_weapon_kills"`
    AvgPickupCnt                 int64   `gorm:"avg_pickup_cnt"`
    AvgThrowableKills            int64   `gorm:"avg_throwable_kills"`
    Country                      string  `gorm:"country"`
    Device                       string  `gorm:"device"`
    Firstlogindate               string  `gorm:"firstlogindate"`
    HeadshotCnt                  int64   `gorm:"headshot_cnt"`
    HeadshotRate                 float64 `gorm:"headshot_rate"`
    MostUsedWeapon               string  `gorm:"most_used_weapon"`
    OfficialAssists              int64   `gorm:"official_assists"`
    OfficialChickens             int64   `gorm:"official_chickens"`
    OfficialCnt                  int64   `gorm:"official_cnt"`
    OfficialKills                int64   `gorm:"official_kills"`
    OfficialTs                   float64 `gorm:"official_ts"`
    PercentileAvgKills           float64 `gorm:"percentile_avg_kills"`
    PercentileAvgMeleeKills      float64 `gorm:"percentile_avg_melee_kills"`
    PercentileAvgMostWeaponKills float64 `gorm:"percentile_avg_most_weapon_kills"`
    PercentileAvgPickupCnt       float64 `gorm:"percentile_avg_pickup_cnt"`
    PercentileAvgThrowableKills  float64 `gorm:"percentile_avg_throwable_kills"`
    Platform                     string  `gorm:"platform"`
    PubgRegion                   string  `gorm:"pubg_region"`
    RankAssists                  int64   `gorm:"rank_assists"`
    RankChickens                 int64   `gorm:"rank_chickens"`
    RankCnt                      int64   `gorm:"rank_cnt"`
    RankKills                    int64   `gorm:"rank_kills"`
    RankTs                       int64   `gorm:"rank_ts"`
}

var db *gorm.DB

func main() {
    // dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
    //     DB_USER, DB_PASSWORD, DB_NAME)
    // db, err := sql.Open("postgres", dbinfo)
    // checkErr(err)
    // defer db.Close()

    // conn := fmt.Sprintf("host = %s  port = 5432 user = %s password = %s dbname = %s sslmode=disable", DB_HOSTNAME, DB_USER, DB_PASSWORD, DB_NAME)
    // conn := "postgres://root:password@localhost:9494/postgres?sslmode=disable"
    //conn := "postgres://localhost:9494/postgres?sslmode=disable"
    conn := fmt.Sprintf("host = %s  port = 9494 user = %s password = %s dbname = %s sslmode=disable", "localhost", "", "", "postgres")

    var err error
    db, err = gorm.Open("postgres", conn)
    if err != nil {
        fmt.Println(err)
        panic(err)

    }

    defer db.Close()
    fmt.Printf("%s\n", conn)
    // db.Create(&wrapped_data{Device: "PC", AccountID: "12345"})

    // var data []wrapped_data
    // db.Find(&data)
    // // db.Where("accountid = ?", accid).Find(&wrapped_data)

    // // fmt.Printf("%+v", info)
    // for _, player := range data {
    //     fmt.Printf("%+v\n", player)

    // }

    accid := "account.0e09f9cb97d94d84812b8f1ac1f78728"
    var data1 wrapped_data
    // db.Raw("SELECT * FROM wrapped_data WHERE accountid = ?", accid).Scan(&data1)
    db.Raw("SELECT * FROM wrapped_data WHERE accountid = ?", accid).Scan(&data1)
    fmt.Printf("%+v", data1)

    // fmt.Println("# Querying")
    // rows, _ := db.Query("select * from wrapped_data")
    // checkErr(err)
    // // fmt.Printf("rows changed %+v", rows)

    // var resp PlayerData
    // // err = rows.Scan(resp)
    // for rows.Next() {
    //     err = rows.StructScan(&resp)
    //     checkErr(err)
    //     fmt.Printf("\n total kills %v", resp)
    // }

    fmt.Println("rows changed")
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
