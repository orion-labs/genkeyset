# genkeyset

[![CircleCI](https://circleci.com/gh/OnBeep/genkeyset.svg?style=svg)](https://circleci.com/gh/OnBeep/genkeyset)

Generate a JWK KeySet for use with Orionlabs PTT.

This tool creates a JWK KeySet perhttps://tools.ietf.org/html/rfc7517 with additional members as provided for in “Additional members” - https://tools.ietf.org/html/rfc7517#section-4.  

At the time of this writing, the additional members used by Orionlabs are:

* *live* (boolean)  Used to flag keys available for use in the Orionlabs PTT system.

## Requirements

To run this tool, you need a Golang environment of at least v1.12.x

https://golang.org/dl/

## Installation

Provided your golang environment is successfully installed, clone this repo and from the root of the repo clone run:

    go build
    
You will find a binary in the current working directory named `genkeyset`

## Usage

To access the internal help menue of `genkeyset`, run:

    ./genkeyset --help
    
To generate a default KeySet with 3 active keys, run:

    ./genkeyset
  
The default output is escaped, and a single line, which isn't particularly readable.  To generated a human readable version, run:

    ./genkeyset -u
    
## Next Steps

Take the escaped output generated from this tool and upload it to SSM Parameter store under the name `/orion-services/observer/session_token_keystore`

Orionlabs OBserver will find it there and the magic will be at your service.

