# bitrise-api

This repository holds golang client code to consume the bitrse api.
The golang client code is generated from the [official open-api specification](https://api-docs.bitrise.io/docs/swagger.json).

In addition, some patches to the specification are applied as the specification does not match the implementation sometimes.

## Usage

The main purpose is to provide a package of the generated golang client code, which you can find in the `go` directory.

### Use in your golang application

`go get github.com/Staffbase/bitrise-api`

## Development

Please find all operation in the `Makefile`.

`make help`

### Create a new patch

To create a new patch you should do the following:

1. Do your manual changes to the swagger_patched.json
2. Create the new patch file:
   `make patch-create`
3. Regenerate golang client code
   `make generate`


<table>
  <tr>
    <td>
      <img src="docs/assets/images/staffbase.png" alt="Staffbase GmbH" width="96" />
    </td>
    <td>
      <b>Staffbase GmbH</b>
      <br />Staffbase is an internal communications platform built to revolutionize the way you work and unite your company. Staffbase is hiring: <a href="https://jobs.staffbase.com" target="_blank" rel="noreferrer">jobs.staffbase.com</a>
      <br /><a href="https://github.com/Staffbase" target="_blank" rel="noreferrer">GitHub</a> | <a href="https://staffbase.com/" target="_blank" rel="noreferrer">Website</a> | <a href="https://jobs.staffbase.com" target="_blank" rel="noreferrer">Jobs</a>
    </td>
  </tr>
</table>
