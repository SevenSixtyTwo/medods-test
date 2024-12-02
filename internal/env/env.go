package env

import utilenv "medods-test/util/env"

var (
	POSTGRES_CONN string
	LOG_LEVEL     string

	SMTP_FROM     string
	SMTP_SERVER   string
	SMTP_PASSWORD string
)

func init() {
	utilenv.LoadFileEnv("./config/config.env")
	utilenv.LoadFileEnv("./secrets/.smtp.env")

	utilenv.LoadStrVar(&POSTGRES_CONN, "POSTGRES_CONN")
	utilenv.LoadStrVar(&LOG_LEVEL, "LOG_LEVEL")

	utilenv.LoadStrVar(&SMTP_FROM, "FROM")
	utilenv.LoadStrVar(&SMTP_PASSWORD, "PASSWORD")
	utilenv.LoadStrVar(&SMTP_SERVER, "SERVER")
}
