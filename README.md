# genkeyset

[![Current Release](https://img.shields.io/github/release/orion-labs/genkeyset.svg)](https://img.shields.io/github/release/orion-labs/genkeyset.svg)

[![CircleCI](https://circleci.com/gh/orion-labs/genkeyset.svg?style=svg)](https://circleci.com/gh/orion-labs/genkeyset)

[![Coverage Status](https://codecov.io/gh/orion-labs/genkeyset/branch/master/graph/badge.svg)](https://codecov.io/gh/orion-labs/genkeyset)

[![Go Report Card](https://goreportcard.com/badge/github.com/orion-labs/genkeyset)](https://goreportcard.com/report/github.com/orion-labs/genkeyset)

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/orion-labs/genkeyset/pkg/genkeyset)

Generate a JWK KeySet for use with Orionlabs PTT.

This tool creates a JWK KeySet per https://tools.ietf.org/html/rfc7517 with additional members as provided for in “Additional members” - https://tools.ietf.org/html/rfc7517#section-4.  

At the time of this writing, the additional members used by Orionlabs are:

* *live* (boolean)  Used to flag keys available for use in the Orionlabs PTT system.

## Installation

1. Navigate to the releases page: https://github.com/orion-labs/genkeyset/releases

2. Download the appropriate version for your Operating System 

    (hint: The MacOs version is `genkeyset_darwin_amd64`)

3. If you're running this on a Mac or Linux system, you need to make it executable.  Open a terminal of some sort, navigate to where you downloaded the binary, and use the `chmod` command:

        cd ~/Downloads <enter>
        
        chmod 755 genkeyset_darwin_amd64
        
        
## Usage

Note, depending on which version you downloaded, your tool program may have another name.

e.g. `genkeyset_darwin_amd64` for Mac, or `genkeyset_windows_amd64.exe` for Windows.

### Normal Use
To generate a default KeySet with 3 active keys, run:

    ./genkeyset
  
The output will be an escaped JSON blob suitable for inclusion in your Orion PTT System instance.

### Accessing the Help Menu

To access the internal help menu of `genkeyset`, run:

    ./genkeyset -h
    
### Legacy CloudFormation Use

The default output is unescaped, which doesn't work for version 1.0.0 of the Orion PTT System in AWS.  To generate an escaped version, run:

    ./genkeyset -e

## Building From Source

### Requirements

To run this tool, you need a Golang environment of at least v1.12.x

https://golang.org/dl/

### Installation

Provided your golang environment is successfully installed, clone this repo and from the root of the repo clone run:

    go build
    
You will find a binary in the current working directory named `genkeyset`.  Run it with the instructions above.

