// Copyright 2015 The tgbot Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"github.com/ChimeraCoder/anaconda"
)

type cmdTweet struct {
	description string
	syntax      string
	re          *regexp.Regexp
	w           io.Writer
	config      TweetConfig
}

type TweetConfig struct {
	Enabled bool
	Consumer_Key string
	Consumer_Secret string
	Access_Token string
	Access_Token_Secret string
}

func NewCmdTweet(w io.Writer, config TweetConfig) Command {
	return &cmdTweet{
		syntax:      "!tw tweet",
		description: "Tweet a message",
		re:          regexp.MustCompile(`^!tw .+`),
		w:           w,
		config:      config,
	}
}

func (cmd *cmdTweet) Enabled() bool {
	return cmd.config.Enabled
}

func (cmd *cmdTweet) Syntax() string {
	return cmd.syntax
}

func (cmd *cmdTweet) Description() string {
	return cmd.description
}

func (cmd *cmdTweet) Match(text string) bool {
	return cmd.re.MatchString(text)
}

func (cmd *cmdTweet) Run(title, from, text string) error {
	tweetText := strings.TrimSpace(strings.TrimPrefix(text, "!tw"))
	tweetLen:=len(tweetText)

	anaconda.SetConsumerKey(cmd.config.Consumer_Key)
	anaconda.SetConsumerSecret(cmd.config.Consumer_Secret)
	api := anaconda.NewTwitterApi(cmd.config.Access_Token, cmd.config.Access_Token_Secret)

	if tweetLen > 140{
		fmt.Fprintf(cmd.w, "msg %v %v chars? Mmm to much for me, size actually matters\n", title, tweetLen)
		return nil
	} else {
		_ , err := api.PostTweet(tweetText, nil)
		if err != nil {
			fmt.Fprintf(cmd.w, "msg %v Useless humans...something went wrong\n", title)
		} else {
			fmt.Fprintf(cmd.w, "msg %v Congrats you did it, new boring tweet posted\n", title)
		}
		return nil
	}
}

func (cmd *cmdTweet) Shutdown() error {
	return nil
}
