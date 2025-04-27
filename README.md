# m365recon

m365recon tool performs information gathering from microsoft endpoints for a given domain.
The tool was inspired from AADInternals module which performs osint and providers. This is a personal hobby project I wrote to learn go and build a tool that would be useful to me. All credits goes to the original researchers.

Currently the tool is very basic, I plan to write additional modules.

## Usage
The `-d` or `--domain` flag is required
```shell
                                
       ___ ___ ___                     
 _____|_  |  _|  _|___ ___ ___ ___ ___ 
|     |_  | . |_  |  _| -_|  _| . |   |
|_|_|_|___|___|___|_| |___|___|___|_|_|
                                       

Microsoft Recon Tool by xcr-19

Usage:
  m365recon [flags]

Flags:
  -d, --domain string   Domain name
  -h, --help            help for m365recon
```

## To-Do
- [ ] Add Search by Tenant
- [ ] Domain and Email search
- [ ] Add Brute module
- [ ] Verify Email Address Module  

## References
- https://aadinternals.com/post/just-looking/ - Dr Nestori Syynimaa
- https://www.msxfaq.de/cloud/authentifizierung/getcredentialtype.htm

## Disclosure
This tool is provided for educational purposes only. The author assumes no responsibility for:
- Any misuse or illegal activities conducted with this tool
- Any damage or disruption caused by the use of this tool
- Any legal consequences resulting from the use of this tool

Users are solely responsible for:
- Ensuring they have proper authorization before using this tool
- Complying with all applicable laws and regulations
- Using the tool in an ethical and responsible manner

This tool is intended for educational and authorized security research purposes only. Unauthorized scanning of domains without permission is strictly prohibited. 