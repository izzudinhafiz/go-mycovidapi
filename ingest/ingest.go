package ingest

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/izzudinhafiz/go-mycovidapi/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Files struct {
	fileName string
	fileURL string
}

var files = []Files{
	{"cases_malaysia.csv","https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/cases_malaysia.csv"},
	{"cases_state.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/cases_state.csv"},
	{"clusters.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/clusters.csv"},
	{"deaths_age.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/deaths_age.csv"},
	{"deaths_malaysia.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/deaths_malaysia.csv"},
	{"deaths_state.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/deaths_state.csv"},
	{"hospital.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/hospital.csv"},
	{"icu.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/icu.csv"},
	{"pkrc.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/pkrc.csv"},
	{"tests_malaysia.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/tests_malaysia.csv"},
	{"tests_state.csv", "https://raw.githubusercontent.com/MoH-Malaysia/covid19-public/main/epidemic/tests_state.csv"},
}

func Ingest()  {
	for _, file := range files {
		downloadFile(file.fileName, file.fileURL)
	}
	db := startDB()
	writeCases(db)
	writeDeaths(db)
	writeClusters(db)
	writeHealthcares(db)


	for _, file := range files {
		err := os.Remove(file.fileName)
		log.Infof("Removing %v\n", file.fileName)
		if err != nil {
			log.Errorf("Could not remove file: %v\n", file.fileName)
			log.Error(err)
		}
	}

	log.Info("DONE!")
}

func openFileWithErrorHandling(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Failed to open %v\n\t%v\n", filePath, err)
	}
	return f
}

func downloadFile(filepath string, url string) {
	log.Infof("Downloading from %v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Failed to download %v\n", url)
		log.Error(err)
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		log.Errorf("Failed to create %v\n", filepath)
		log.Error(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Errorf("Failed to write to %v\n", filepath)
		log.Error(err)
	}
}

func startDB() *gorm.DB {
	var dsn = fmt.Sprintf(
		"host=%v user=%v password=%v port=%v database=%v sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Cases{})
	db.AutoMigrate(&models.CasesAgeCategory{})
	db.AutoMigrate(&models.CasesCluster{})
	db.AutoMigrate(&models.Clusters{})
	db.AutoMigrate(&models.Deaths{})
	db.AutoMigrate(&models.Hospital{})
	db.AutoMigrate(&models.ICU{})
	db.AutoMigrate(&models.QuarantineCentre{})

	return db
}
