terraform-provider-inwx
==========================

Terraform provider for INWX (InterNetworX)

## Description

With this custom terraform provider plugin you can manage your INWX domains.

## Usage

Add plugin binary to your ~/.terraformrc file
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
  sandbox = "${var.inwx_sandbox}" // default is false
  TAN = "${var.inwx_tan}" // if 2-Factor authentication is enabled for your INWX account
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