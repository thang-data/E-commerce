package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

var env *ENV

type FrontendUrls struct {
	FrontendLoginUrl          string
	FrontendForgotPasswordUrl string
}
type Urls struct { // other urls
	PublicLogoUrl string
}

type ENV struct {
	// Mode
	ApiMode                     bool
	BatchMode                   bool
	MaxNumOfRecordsProcessAsync int // Maximum number of records that batch will process async at same time

	// API
	APIPort       string
	BackendApiUrl string
	FrontendUrl   string
	FrontendUrls
	Urls

	// DB
	DBDriver           string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBHostReadonly     string
	DBPortReadonly     string
	DBUserReadonly     string
	DBPasswordReadonly string
	DBNameReadonly     string

	// Log
	LogLevel string

	// Email
	SMTPHost          string
	SMTPPort          string
	SMTPUsername      string
	SMTPPassword      string
	SMTPDisplayName   string
	SMTPSenderAddress string

	SendGridSenderName    string
	SendGridSenderAddress string
	SendGridApiKey        string

	// Admin
	InitialAdminEmail    string
	InitialAdminPassword string
}

const (
	ApiMode                     = "API_MODE"
	BatchMode                   = "BATCH_MODE"
	MaxNumOfRecordsProcessAsync = "MAX_NUM_OF_RECORDS_PROCESS_ASYNC"

	//Backend URL
	BackendApiUrl = "BACKEND_API_URL"

	//Frontend URL
	FrontendUrl = "FRONTEND_URL"

	//DB
	DBHost             = "DB_HOST"
	DBPort             = "DB_PORT"
	DBUser             = "DB_USER"
	DBPassword         = "DB_PASSWORD"
	DBName             = "DB_NAME"
	DBHostReadonly     = "DB_HOST_READONLY"
	DBPortReadonly     = "DB_PORT_READONLY"
	DBUserReadonly     = "DB_USER_READONLY"
	DBPasswordReadonly = "DB_PASSWORD_READONLY"
	DBNameReadonly     = "DB_NAME_READONLY"

	LogLevel = "LOG_LEVEL"

	SMTPHost          = "SMTP_HOST"
	SMTPPort          = "SMTP_PORT"
	SMTPUsername      = "SMTP_USERNAME"
	SMTPPassword      = "SMTP_PASSWORD"
	SMTPDisplayName   = "SMTP_DISPLAY_NAME"
	SMTPSenderAddress = "SMTP_SENDER_ADDRESS"

	SendGridSenderName    = "SENDGRID_SENDER_NAME"
	SendGridSenderAddress = "SENDGRID_SENDER_ADDRESS"
	SendGridApiKey        = "SENDGRID_API_KEY"
	InitialAdminEmail     = "INITIAL_ADMIN_EMAIL"
	InitialAdminPassword  = "INITIAL_ADMIN_PASSWORD"
)

func FetchEnvironmentVarialble() {
	env = &ENV{

		APIPort:       "1324",
		BackendApiUrl: os.Getenv(BackendApiUrl),
		FrontendUrl:   os.Getenv(FrontendUrl),

		DBDriver:              "postgres",
		DBHost:                os.Getenv(DBHost),
		DBPort:                os.Getenv(DBPort),
		DBUser:                os.Getenv(DBUser),
		DBPassword:            os.Getenv(DBPassword),
		DBName:                os.Getenv(DBName),
		DBHostReadonly:        os.Getenv(DBHostReadonly),
		DBPortReadonly:        os.Getenv(DBPortReadonly),
		DBUserReadonly:        os.Getenv(DBUserReadonly),
		DBPasswordReadonly:    os.Getenv(DBPasswordReadonly),
		DBNameReadonly:        os.Getenv(DBNameReadonly),
		SMTPHost:              os.Getenv(SMTPHost),
		SMTPPort:              os.Getenv(SMTPPort),
		SMTPUsername:          os.Getenv(SMTPUsername),
		SMTPPassword:          os.Getenv(SMTPPassword),
		SMTPDisplayName:       os.Getenv(SMTPDisplayName),
		SMTPSenderAddress:     os.Getenv(SMTPSenderAddress),
		SendGridSenderName:    os.Getenv(SendGridSenderName),
		SendGridSenderAddress: os.Getenv(SendGridSenderAddress),
		SendGridApiKey:        os.Getenv(SendGridApiKey),
		InitialAdminEmail:     os.Getenv(InitialAdminEmail),
		InitialAdminPassword:  os.Getenv(InitialAdminPassword),

		LogLevel: os.Getenv(LogLevel),
	}

	apiModeEnv := os.Getenv(ApiMode)
	batchModeEnv := os.Getenv(BatchMode)
	maxNumOfRecordsEnv := os.Getenv(MaxNumOfRecordsProcessAsync)
	apiMode, _ := strconv.ParseBool(apiModeEnv)
	batchMode, _ := strconv.ParseBool(batchModeEnv)
	maxNumOfRecords, _ := strconv.Atoi(maxNumOfRecordsEnv)

	// Format URLs
	env.BackendApiUrl = formatUrl(env.BackendApiUrl)
	env.FrontendUrl = formatUrl(env.FrontendUrl)

	env.FrontendLoginUrl = env.FrontendUrl + "/login"
	env.FrontendForgotPasswordUrl = env.FrontendUrl + "/signup"
	env.PublicLogoUrl = env.BackendApiUrl + "/files/public/logo.png"

	if apiMode == false && batchMode == false { // default without ENV
		logrus.Info("apiMode = false & batchMode = false FROM ENV => default will be all mode (API and Batch)")
		apiMode = true
		batchMode = true
	}
	env.ApiMode = apiMode
	env.BatchMode = batchMode
	logrus.Info("apiMode = ", apiMode)
	logrus.Info("batchMode = ", batchMode)

	if maxNumOfRecords < 1 {
		logrus.Info("maxNumOfRecords < 1 => default maxNumOfRecords will be 3")
		maxNumOfRecords = 3
	}
	env.MaxNumOfRecordsProcessAsync = maxNumOfRecords
	logrus.Info("maxNumOfRecordsProcessAsync = ", maxNumOfRecords)

	// DB
	if len(env.DBHostReadonly) == 0 {
		logrus.Info("DBHostReadonly is empty => use the same info with DBHost")
		env.DBHostReadonly = env.DBHost
		env.DBPortReadonly = env.DBPort
		env.DBUserReadonly = env.DBUser
		env.DBPasswordReadonly = env.DBPassword
		env.DBNameReadonly = env.DBName
	}
}

func formatUrl(url string) string {
	if url == "" {
		return ""
	}

	if url[len(url)-1:] == "/" {
		return url[:len(url)-1]
	} else {
		return url
	}
}

func GetConfig() ENV {
	if env == nil {
		FetchEnvironmentVarialble()
	}
	return *env
}
