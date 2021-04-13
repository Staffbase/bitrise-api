This repository holds golang client code to consume the bitrse api.
The golang client code is generated from the official open-api specification.

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
   `diff -u swagger.json swagger_patched.json > staffbase.patch`
3. Regenerate golang client code