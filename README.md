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
  sandbox = true // default is false
}
```
