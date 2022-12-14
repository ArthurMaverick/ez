# EZ
## Install ez binary with curl on Linux 
1. Download ez
```
curl -L https://github.com/ArthurMaverick/ez/releases/download/v0.0.2/ez_0.0.2_Linux_x86_64.tar.gz >> ez.tar.gz
```
2. Install ez

   ```bash
   sudo install -o root -g root -m 0755 ez /usr/local/bin/ez
   ```

   
   If you do not have root access on the target system, you can still install kubectl to the `~/.local/bin` directory:

   ```bash
   chmod +x ez
   mkdir -p ~/.local/bin
   mv ./ez ~/.local/bin/ez
   # and then append (or prepend) ~/.local/bin to $PATH
   ```

   
## ABOUT
### **This command line interface was created with the aim of helping developers by offering a toolkit of tools that are widely used in everyday life**


## Net Command
  The Net tool aims to offer a set of APIs that help with network debugging. Currently Net toolkit only offers the command that identifies useful information about IPs.
        For example using the command **ez net --ip 1.1.1.1** you will get this information:
    
```bash
Continent:
  Code: NA
  Name: North America
Country: US
Location: 34.0522,-118.2437
Anycast: true
City: Los Angeles
Org: AS13335 Cloudflare, Inc.
Timezone: America/Los_Angeles
IP: 1.1.1.1
CountryCurrency:
  Symbol: $
  Code: USD
Hostname: one.one.one.one
CountryName: United States
Postal: 90076
Region: California
CountryFlag:
  Emoji: 🇺🇸
  Unicode: U+1F1FA U+1F1F8

```
## IaC command
  The IaC tool aims to generate cloudformation and terraform templates, currently the iac commands only work to generate cfn templates. I'm working for ez cmd to be able to generate terraform templates too.
    
  ez iac can generate ecs, elb, lambda, vpc, vpc2 and vpn-site-to-site templates.
    To generate cloudformation templates just use the following command:

```bash
    ez iac --provider cfn --create vpc
```

## Cloud command
  The ez integration with the aws cloud is still in development mode but the goal is to bring the most used scripts in the pipelines. centralize These tools can make day-to-day development more flexible.