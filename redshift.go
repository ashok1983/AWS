package paintjobs

import (
	"context"
	"log"
	"strconv"
	"time"

	g "git.projectbro.com/isd/opendata-api"
	"git.projectbro.com/isd/opendata-api/config"
	"git.projectbro.com/isd/opendata-api/models"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata"
	"github.com/aws/aws-sdk-go/aws"
	rsvc "github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
)

type PlayerData struct {
	Device                       string `json:"device"`
	Platform                     string `json:"platform"`
	Country                      string `json:"country"`
	PubgRegion                   string `json:"pubg_region"`
	AccountID                    string `json:"accountId"`
	Firstlogindate               string `json:"firstlogindate"`
	TotalTs                      string `json:"Total_ts"`
	OfficialTs                   string `json:"official_ts"`
	RankTs                       string `json:"rank_ts"`
	TotalCnt                     string `json:"Total_cnt"`
	OfficialCnt                  string `json:"official_cnt"`
	RankCnt                      string `json:"rank_cnt"`
	TotalKills                   string `json:"Total_kills"`
	OfficialKills                string `json:"official_kills"`
	RankKills                    string `json:"rank_kills"`
	TotalAssists                 string `json:"Total_assists"`
	OfficialAssists              string `json:"official_assists"`
	RankAssists                  string `json:"rank_assists"`
	TotalChickens                string `json:"Total_chickens"`
	OfficialChickens             string `json:"official_chickens"`
	RankChickens                 string `json:"rank_chickens"`
	HeadshotCnt                  string `json:"headshot_cnt"`
	HeadshotRate                 string `json:"headshot_rate"`
	MostUsedWeapon               string `json:"most_used_weapon"`
	AvgMostWeaponKills           string `json:"avg_most_weapon_kills"`
	AvgKills                     string `json:"avg_kills"`
	AvgMeleeKills                string `json:"avg_melee_kills"`
	AvgThrowableKills            string `json:"avg_throwable_kills"`
	AvgPickupCnt                 string `json:"avg_pickup_cnt"`
	PercentileAvgMostWeaponKills string `json:"percentile_avg_most_weapon_kills"`
	PercentileAvgKills           string `json:"percentile_avg_kills"`
	PercentileAvgMeleeKills      string `json:"percentile_avg_melee_kills"`
	PercentileAvgThrowableKills  string `json:"percentile_avg_throwable_kills"`
	PercentileAvgPickupCnt       string `json:"percentile_avg_pickup_cnt"`
}

const (
	createTable = `CREATE TABLE wrapped_data(
	device                           VARCHAR(20) NOT NULL,
	platform                         VARCHAR(20) NOT NULL,
	country                          VARCHAR(20) NOT NULL,
	pubg_region                      VARCHAR(20) NOT NULL,
	accountid                        VARCHAR(40) NOT NULL,
	firstlogindate                   DATE NOT NULL,
	total_ts                         FLOAT NOT NULL,
	official_ts                      FLOAT NOT NULL,
	rank_ts                          FLOAT NOT NULL,
	total_cnt                        BIGINT NOT NULL,
	official_cnt                     BIGINT NOT NULL,
	rank_cnt                         BIGINT NOT NULL,
	total_kills                      BIGINT NOT NULL,
	official_kills                   BIGINT NOT NULL,
	rank_kills                       BIGINT NOT NULL,
	total_assists                    BIGINT NOT NULL,
	official_assists                 BIGINT NOT NULL,
	rank_assists                     BIGINT NOT NULL,
	total_chickens                   BIGINT NOT NULL,
	official_chickens                BIGINT NOT NULL,
	rank_chickens                    BIGINT NOT NULL,
	headshot_cnt                     BIGINT NOT NULL,
	headshot_rate                    FLOAT NOT NULL,
	most_used_weapon                 VARCHAR(20) NOT NULL,
	avg_most_weapon_kills            FLOAT NOT NULL,
	avg_kills                        FLOAT NOT NULL,
	avg_melee_kills                  FLOAT NOT NULL,
	avg_throwable_kills              FLOAT NOT NULL,
	avg_pickup_cnt                   FLOAT NOT NULL,
	percentile_avg_most_weapon_kills FLOAT NOT NULL,
	percentile_avg_kills             FLOAT NOT NULL,
	percentile_avg_melee_kills       FLOAT NOT NULL,
	percentile_avg_throwable_kills   FLOAT NOT NULL,
	percentile_avg_pickup_cnt        FLOAT NOT NULL
 );`
)

// Redshift describes a redshift service
type Redshift interface {
	Get(string, string) (*models.TotalKillInfo, error)
	LoadData(string, string) error
}

// RedshiftClient  struct
type RedshiftService struct {
	config   *config.Config
	rsClient *rsvc.RedshiftDataAPIService
}

// NewRedshiftClient creates a new RedshiftClient
func NewRedshiftClient(sc *g.Gamelocker) RedshiftService {
	rs := RedshiftService{
		config:   sc.Config,
		rsClient: sc.RedshiftClient,
	}

	return rs
}

func (rs *RedshiftService) Get(accID string, fileName string) (*models.TotalKillInfo, error) {
	response := &models.TotalKillInfo{}
	var playerInfo PlayerData
	response.ID = accID
	response.FileName = fileName
	response.TotalKill, _ = strconv.Atoi(playerInfo.TotalKills)
	return response, nil
}

func (rs *RedshiftService) LoadData(fileName string) (string, error) {

	log.Println("Using the RedshiftData SDK to query Redshift")

	// Connect to RedshiftData API
	cfg, err := awscfg.LoadDefaultConfig(context.TODO(), awscfg.WithRegion(rs.config.AwsRedshiftRegion))
	if err != nil {
		log.Printf("Unable to load SDK config, %v", err)
		return "", err
	}
	svc := redshiftdata.NewFromConfig(cfg)

	// Fire off a query
	statement, statementErr := svc.ExecuteStatement(context.TODO(), &redshiftdata.ExecuteStatementInput{
		ClusterIdentifier: aws.String(rs.config.AwsRedshiftClusterID),
		Database:          aws.String(rs.config.AwsRedshiftDBName),
		Sql:               aws.String("select * from wrapped_data limit 1;"),
		SecretArn:         aws.String(rs.config.AwsRedshiftIAMRole),
	})

	if statementErr != nil {
		log.Printf("%#v\n", statementErr.Error())
		return "", statementErr
	}
	log.Println("Statement ID for retrieving results later", statement.Id)

	// Wait for a response. You can add an automatic retry here
	time.Sleep(1 * time.Second)

	// Retrieve results
	result, resultErr := svc.GetStatementResult(context.TODO(), &redshiftdata.GetStatementResultInput{
		Id: statement.Id,
	})
	if resultErr != nil {
		log.Println(resultErr)
		return "", resultErr
	}

	// Do something with results
	log.Printf("result is %+v\n", result)
	return "Load Success 200 OK", nil
}
