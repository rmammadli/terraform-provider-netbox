# Terraform Provider Netbox
This repo aims to be a usefull collection for implementing netbox dcim resources.

## Build provider

Run the following command to build the netbox provider

```shell
$ go build -o terraform-provider-netbox
```

## Test sample configuration

Build and install the provider ( currently for windows only, please make appropriate changes in Makefile for other os )

```shell
$ make install
```

Then, navigate to the `terraform` directory.
Please define the netbox url, netbox token and connection type, which might be http or https depending on your netbox configuration.
These inputs can be provided as either environment variables or input variables in terraform, for example:

environment variables:
```shell
export NETBOX_HOST=mynetbox.local
export NETBOX_TOKEN=1234567abc1234567abc
export CONNECTION_TYPE=https
```


terraform input variables:
```shell
variable "netbox_host" {
    type = string
    description = "netbox host url"
    default = "mynetbox.local"
}

variable "netbox_token" {
    type = string
    description = "netbox token"
    default = "1234567abc1234567abc"
}


variable "connection_type"  {
    type = string
    description = "http / https"
    default = "https"
}
```

```shell
$ cd terraform
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
