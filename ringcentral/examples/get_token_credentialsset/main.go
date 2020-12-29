package main

import (
	"fmt"

	"github.com/grokify/oauth2more/ringcentral"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

type Options struct {
	CredsPath string `short:"c" long:"credspath" description:"RingCentral Credentials File" required:"true"`
	CredsVar  string `short:"v" long:"credsvar" description:"RingCentral Credentials File Entry" required:"true"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	credsSet, err := ringcentral.ReadFileCredentialsSet(opts.CredsPath)
	if err != nil {
		log.Fatal(err)
	}
	creds, err := credsSet.Get(opts.CredsVar)
	if err != nil {
		log.Fatal(err)
	}
	token, err := creds.NewToken()
	if err != nil {
		log.Fatal(err)
	}

	token.Expiry = token.Expiry.UTC()

	fmtutil.PrintJSON(token)
	/*
		files, err := config.LoadDotEnv(opts.EnvPath, os.Getenv("ENV_PATH"), "./.env")
		if err != nil {
			log.Fatal(errors.Wrap(err, "E_LOAD_DOT_ENV"))
		}
		fmtutil.PrintJSON(files)

		if len(opts.EnvVar) > 0 {
			if len(os.Getenv(opts.EnvVar)) == 0 {
				log.Fatal("E_NO_VAR")
			}

			credentials, err := ringcentral.NewCredentialsJSON([]byte(os.Getenv(opts.EnvVar)))
			if err != nil {
				log.Fatal(
					errors.Wrap(
						err, fmt.Sprintf("E_JSON_UNMARSHAL [%v]", os.Getenv(opts.EnvVar))))
			}
			token, err := credentials.NewToken()
			if err != nil {
				log.Fatal(err)
			}

			token.Expiry = token.Expiry.UTC()

			fmtutil.PrintJSON(token)
		} else {
			fmt.Printf("No EnvVar [-v]\n")
		}
	*/
	fmt.Println("DONE")
}
