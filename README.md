terraform-provider-inwx
==========================

Terraform provider for INWX (InterNetworX)

## Description

With this custom terraform provider plugin you can manage your INWX domains.

## Usage

This plugin is currently not yet part of the official list of [terraform-providers](https://github.com/terraform-providers),
therefore you have to install it manually:

1. [Download](https://github.com/andrexus/terraform-provider-inwx/releases) the
   compiled plugin and make the file executable (`chmod +x terraform-provider-inwx`).  
   Or if you have Go installed:
    ```
    go get github.com/andrexus/terraform-provider-inwx
    go install github.com/andrexus/terraform-provider-inwx
    ```
2. Add plugin binary to your ~/.terraformrc file
    ```
    providers {
       inwx = "/path/to/your/bin/terraform-provider-inwx"
    }
    ```

### Provider Configuration

```
provider "inwx" {
  username = "${var.inwx_username}"
  password = "${var.inwx_password}"
  sandbox  = "${var.inwx_sandbox}" // default is false
  tan      = "${var.inwx_tan}"     // if 2-Factor authentication is enabled for your INWX account
}

// Example record
resource "inwx_record" "example" {
  domain   = "example.com"
  name     = ""
  type     = "MX"
  value    = "mx.example.com"
  ttl      = 3600
  priority = 10
}

```

#### Example (variables.tf)
```
variable "inwx_username" {}
variable "inwx_password" {}
variable "inwx_tan" {}
variable "inwx_sandbox" {
  default = false
}
```
#### Example (terraform.tfvars)
```
inwx_username = "username"
inwx_password = "password"
```

If you don't specify the value for _inwx_tan_ variable (which you normally should not do
because TAN is only valid for 30 seconds) terraform will prompt you to input the TAN in command line.
Alternatively you can pass this variable as a command line parameter.
```
terraform plan -var 'inwx_tan=545817'
```
