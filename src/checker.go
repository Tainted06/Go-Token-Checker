package source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Normal checker
func (Xc *Checker) Check(token string) string {

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	Xc.Errs(err)

	// Add headers
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("authorization", Xc.FormatToken(token))
	req.Header.Set("referer", "https://discord.com/app")
	req.Header.Set("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Brave\";v=\"108\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("x-discord-locale", "en-US")
	req.Header.Set("x-super-properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEwOC4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTA4LjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjE2NTQ4NSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=")

	// Process request
	resp, err := Xc.Client.Do(req)
	Xc.Errs(err)
	var typ = Xc.Resp

	// Process response
	if resp.StatusCode == 200 {
		typ = true
		Xc.InvalidResp = false
		fmt.Println("["+g+"✓"+r+"] ("+u+""+g+"VALID"+u+""+rb+"):", Xc.FormatToken(token))
		Xc.Valid++
	} else if resp.StatusCode == 403 {
		typ = false
		Xc.InvalidResp = false
		fmt.Println("[\033[33m/"+r+"] ("+u+""+c+"LOCKED"+u+""+rb+"):", Xc.FormatToken(token))
		Xc.Locked++
	} else {
		typ = false
		Xc.InvalidResp = true
		fmt.Println("[\033[31mX"+r+"] ("+u+""+c+"INVALID"+u+""+rb+"):", Xc.FormatToken(token))
		Xc.Invalid++
	}
	Xc.Resp = typ
	Xc.Token = token
	return Xc.Token
}

// Flagged checker
func (Xc *Checker) CheckFlagged(token string) string {

	// Create request
	req, err := http.NewRequest("GET", flagged_url, nil)
	Xc.Errs(err)

	// Add headers
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("authorization", Xc.FormatToken(token))
	req.Header.Set("referer", "https://discord.com/app")
	req.Header.Set("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Brave\";v=\"108\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("x-discord-locale", "en-US")
	req.Header.Set("x-super-properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEwOC4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTA4LjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjE2NTQ4NSwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=")

	// Process request
	resp, err := Xc.Client.Do(req)
	Xc.Errs(err)

	// Process response
	if resp.StatusCode == 200 {
		content, err := ioutil.ReadAll(resp.Body)
		Xc.Errs(err)
		var jsonContent map[string]interface{}
		json.Unmarshal(content, &jsonContent)
		flagsInt, _ := jsonContent["flags"].(float64)
		publicFlagsInt, _ := jsonContent["public_flags"].(float64)
		if flagsInt >= 1048576 || publicFlagsInt >= 1048576 {
			return "flagged"
		} else {
			return "not flagged"
		}
	} else if resp.StatusCode == 429 {
		return "ratelimit"
	} else {
		return "-"
	}
}
