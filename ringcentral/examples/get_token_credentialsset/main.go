package main

import (
	"fmt"
	"os"

	"github.com/grokify/oauth2more/ringcentral"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Options struct {
	CredsPath  string `short:"c" long:"credspath" description:"RingCentral Credentials File" required:"true"`
	AccountKey string `short:"a" long:"account" description:"RingCentral Credentials File Entry" required:"true"`
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to parse CLI options")
	}

	credsSet, err := ringcentral.ReadFileCredentialsSet(opts.CredsPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("filepath", opts.CredsPath).
			Msg("failed to read creds file")
	}
	creds, err := credsSet.Get(opts.AccountKey)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to get credentials from credentialset")
	}

	token, err := ringcentral.NewTokenCli(creds, "mystate")
	if err != nil {
		log.Fatal().
			Err(err).
			Str("filepath", opts.CredsPath).
			Str("account", opts.AccountKey).
			Msg("failed to get new token")
	}

	fmtutil.PrintJSON(token)

	fmt.Println("DONE")
}
