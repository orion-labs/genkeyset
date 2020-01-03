# genkeyset

Generate a JWK KeySet for use with Orionlabs PTT.

This tool creates a JWK KeySet per https://tools.ietf.org/html/rfc7517 with additional members as provided for in “Additional members” - https://tools.ietf.org/html/rfc7517#section-4.  

At the time of this writing, the additional members used by Orionlabs are:

* *live* (boolean)  Used to flag keys available for use in the Orionlabs PTT system.

## Installation

1. Navigate to the releases page: https://github.com/orion-labs/genkeyset/releases

2. Download the appropriate version for your Operating System (*hint:* The MacOs version is `genkeyset_darwin_amd64`)

3. Open a terminal of some sort, navigate to where you downloaded the binary, and run it.  e.g.:

        cd Downloads <enter>
        ./genkeyset_darwin_amd64 <enter>
   
The output will be an escaped JSON blob suitable for inclusion in your Orion Cloudformation Template.


## Usage

To access the internal help menue of `genkeyset`, run:

    ./genkeyset -h
    
To generate a default KeySet with 3 active keys, run:

    ./genkeyset
  
The default output is escaped, and a single line, which isn't particularly readable.  To generated a human readable version, run:

    ./genkeyset -u
    
## Next Steps

Take the escaped output generated from this tool and upload it to SSM Parameter store under the name `/orion-services/observer/session_token_keystore`

Orionlabs OBserver will find it there and the magic will be at your service.

## Building From Source

### Requirements

To run this tool, you need a Golang environment of at least v1.12.x

https://golang.org/dl/

### Installation

Provided your golang environment is successfully installed, clone this repo and from the root of the repo clone run:

    go build
    
You will find a binary in the current working directory named `genkeyset`

