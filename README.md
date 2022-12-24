# Go-Token-Checker
Discord Token Checker Proxyless & Fast

Star ⭐ for support!

# Usage
```
1. Download GoLang and the source code or download the exe from releases 
2. Put your tokens into input.txt
3. Run the exe or run the code
4. Results will be printed to the console and saved in the files
```

![NVIDIA_Share_2l1iAdgZbB](https://user-images.githubusercontent.com/110062350/203656244-88bb929b-5081-4489-8740-55b50eb723e7.gif)

# Config Information
```
- The first field is check_flagged, if you would like the checker to check for flagged tokens then set it to true, if not then false
- The second field is offset_ms, this is how long it will wait between each token in milliseconds
```

# Ratelimit Information
```
- If you enable checking for flags and check a lot of tokens, there is the possibility of getting ratelimited
- To avoid ratelimits either don't check for flagged or increase the offset
```

# Other Information
```
- All token formats (token, email:pass:token, etc) are supported as long as the token is the longest part of it
- If you find any bugs with it make a pull request or an issue
- This tool is an educational POC tool to demonstrate how goroutines can be used to send requests 
```

# Benchmarks
```
With checking for flagged enabled and a 1 ms offset, 1,989 tokens got checked in 4.3 seconds 
With checking for flagged enabled and a 1 ms offset, 2,605 tokens got checked in 5.6 seconds
With checking for flagged enabled and a 1 ms offset, 73,311 tokens got checked in 116.5 seconds
With checking for flagged disabled and 1 ms offset, 95,289 tokens got checked in 239.7 seconds
```