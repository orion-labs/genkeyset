/*
Copyright 2020 The genkeyset Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package genkeyset

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/pkg/errors"
	"time"
)

func toChar(i int) string {
	return string(rune('a' + i))
}

// GenerateKeyID creates a key identifier from the time object given.  The identifier is of the form last 2 digits of the date + 2 digits of the month + 2 digits of the day.
func GenerateKeyID(date time.Time) (keyId string) {
	month := fmt.Sprintf("%02d", int(date.Month()))
	day := fmt.Sprintf("%02d", date.Day())
	keyId = fmt.Sprintf("%d%s%s", date.Year()-2000, month, day)

	return keyId
}

// CreatePrivateKey creates an ECDSA p256 key
func CreatePrivateKey() (privateKey *ecdsa.PrivateKey, err error) {
	pubkeyCurve := elliptic.P256()

	privateKey, err = ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		err = errors.Wrap(err, "failed to generate private key")
		return privateKey, err
	}

	return privateKey, err
}

// GenerateKeySet Creates a JWK KeySet of the length given and returns it.  Keys in the KeySet have the additional attribute 'live' set to true as required by Orionlabs PTT.
func GenerateKeySet(length int) (keyset jwk.Set, err error) {
	keyset.Keys = make([]jwk.Key, 0)

	for i := 0; i < length; i++ {
		privateKey, err := CreatePrivateKey()
		if err != nil {
			return keyset, err
		}

		key, err := jwk.New(privateKey)
		if err != nil {
			err = errors.Wrap(err, "failed to create jwk")
			return keyset, err
		}

		key.Set("live", true)
		key.Set("kid", fmt.Sprintf("%s%s", GenerateKeyID(time.Now()), toChar(i)))
		key.Set("alg", "ES256")

		keyset.Keys = append(keyset.Keys, key)
	}

	return keyset, err
}
