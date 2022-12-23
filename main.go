package main

import (
	"encoding/json"
	"fmt"
	"gochecker/src"
	"io/ioutil"
	"strconv"
	"sync"
	"time"
)

// Main function
func main() {

	// Checker and logo
	var checker = source.Checker{}
	checker.Cls()
	checker.Logo()

	// Reading tokens
	token, err := checker.Read_tokens("tokens.txt")
	checker.Errs(err)

	// Reading config
	configFile, err := ioutil.ReadFile("config.json")
	checker.Errs(err)
	var config source.Config
	err = json.Unmarshal(configFile, &config)
	checker.Errs(err)

	// Colors
	c, r, rb, u := "\u001b[30;1m", "\033[39m", "\u001b[0m", "\u001b[4m"

	// Console title function
	go func(flagged bool) {
		for {
			if flagged {
				source.SetTitle(fmt.Sprintf("GO Token Checker - github.com/Tainted06/Go-Token-Checker - Flagged: %d - Locked: %d - Invalid: %d - Valid: %d - Total: %d", checker.Flagged, checker.Locked, checker.Invalid, checker.Valid, checker.All))
			} else {
				source.SetTitle(fmt.Sprintf("GO Token Checker - github.com/Tainted06/Go-Token-Checker - Locked: %d - Invalid: %d - Valid: %d - Total: %d", checker.Locked, checker.Invalid, checker.Valid, checker.All))
			}
		}
	}(config.CheckFlagged)

	// Starting
	start := time.Now()
	if len(token) != 0 {
		var wg sync.WaitGroup
		wg.Add(len(token))
		for i := 0; i < len(token); i++ {

			// Waiting
			time.Sleep(time.Duration(config.Offset) * time.Millisecond)

			// Checking function
			go func(i int) {
				defer wg.Done()
				defer func() { checker.All++ }()

				// Checking Flagged
				if config.CheckFlagged {
					resp := checker.CheckFlagged(token[i])
					if resp == "flagged" {
						checker.Flagged++
						checker.OpnFile("flagged.txt", token[i])
						fmt.Println("[\033[31mX"+r+"] ("+u+""+c+"FLAGGED"+u+""+rb+"):", checker.FormatToken(token[i]))
						return
					} else if resp == "ratelimit" {
						fmt.Println("[\033[33m/"+r+"] ("+u+""+c+"RATELIMIT"+u+""+rb+"):", "Ratelimited when checking for flagged")
					}
				}

				// Check token
				checker.Check(token[i])

				// Save token
				if checker.Resp == true {
					checker.OpnFile("valid.txt", token[i])
				} else if checker.Resp == false {
					checker.OpnFile("locked.txt", token[i])
				} else {
					checker.OpnFile("invalid.txt", token[i])
				}

			}(i)
		}
		wg.Wait()
	}

	// Finished report
	elapsed := time.Since(start).Seconds()
	elapsedStr := strconv.FormatFloat(elapsed, 'f', 5, 64)
	if config.CheckFlagged {
		fmt.Println("[\033[32m✓\033[39m] (TIME\033[39m):", elapsedStr+" s", "\033[39m(\033[33mLOCKED\033[39m):", checker.Locked, "\033[39m(\033[33mFLAGGED\033[39m):", checker.Flagged, "(\033[31mINVALID\033[39m):", checker.Invalid, "(\033[32mVALID\033[39m):", checker.Valid, "(\u001b[34;1mTOTAL\033[39m):", checker.All)
	} else {
		fmt.Println("[\033[32m✓\033[39m] (TIME\033[39m):", elapsedStr+" s", "\033[39m(\033[33mLOCKED\033[39m):", checker.Locked, "(\033[31mINVALID\033[39m):", checker.Invalid, "(\033[32mVALID\033[39m):", checker.Valid, "(\u001b[34;1mTOTAL\033[39m):", checker.All)
	}
}
