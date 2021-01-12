package main

import (
	"fmt"

	"github.com/grokify/oauth2more/ringcentral"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"go.uber.org/zap"
)

type Options struct {
	CredsPath  string `short:"c" long:"credspath" description:"RingCentral Credentials File" required:"true"`
	AccountKey string `short:"a" long:"account" description:"RingCentral Credentials File Entry" required:"true"`
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	credsSet, err := ringcentral.ReadFileCredentialsSet(opts.CredsPath)
	if err != nil {
		logger.Fatal("failed to read creds file",
			zap.String("filepath", opts.CredsPath),
			zap.String("err", err.Error()))
	}
	creds, err := credsSet.Get(opts.AccountKey)
	if err != nil {
		log.Fatal(err)
	}

	token, err := ringcentral.NewTokenCli(creds, "mystate")
	if err != nil {
		logger.Fatal("failed to get new token",
			zap.String("filepath", opts.CredsPath),
			zap.String("account", opts.AccountKey),
			zap.String("err", err.Error()))
	}

	fmtutil.PrintJSON(token)

	fmt.Println("DONE")
}
