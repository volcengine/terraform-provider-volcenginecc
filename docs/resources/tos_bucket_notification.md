---
page_title: "{.Name} {.Type} - {.ProviderName}"
subcategory: "TOS"
description: |-
{ .Description | plainmarkdown | trimspace | prefixlines "  " }
---

# {.Name} ({.Type})

{ .Description | trimspace }

## Example Usage

{ tffile (printf "examples/resources/%s/tos_bucket_notification.tf" .Name)}

{ .SchemaMarkdown | trimspace }
{- if .HasImport }

## Import

Import is supported using the following syntax:

{ codefile "shell" .ImportFile }
{- end }
