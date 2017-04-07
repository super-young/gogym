Implement a command to generate a random secret.

Similar to openssl's `rand` command. To generate n bytes of random data:

```
$ openssl rand -base64 64
rIvsPesK6UatalH+I0bqJzdryh5u/D0rEKNcarjNdFVoo4x1TG3Tddq2jlNZ7isB
2tYvjAvt0Nv5PL/ylTvTqQ==

$ openssl rand -base64 128
H8BtEQMMkgDOIgxSFATeOEXY7oc2sIeJCF935hkKfeaAzBPmJ3zfSaLhUqvvIwiT
FAHWvY9QIMJqPrYA4oxXyReZXA2Y5DL4WzOMz7CDuk+GcG50l+LrlqSxwJIhLqmD
JGT0uKIxBx6KZeDa1jH1KHHxt669v69TUPiSyxj1ojQ=
```

## Instruction

A typical random function may have enough patterns for attackers to exploit. A cryptographically secured random function makes it harder for attackers to guess the next random value.

* [crypto/rand](https://golang.org/pkg/crypto/rand)
  * Generate cryptographically secured random bytes.
* [encoding/base64](https://golang.org/pkg/crypto/rand/)
  * use `base64.StdEncoding`.

## Examples

```
$ go run gensecret.go 64
AY0oDKCSLMDE9m5SAxqNJuI3TCMbITW2JNDZN0lJQ3aYxGiRSymRXfRpntl5QyKygTj5uyhOl0+KyE6RUQXjDg==

$ go run gensecret.go 128
pCiVozBYn7lTqOL8mCdnYQUhmmvWTZGjAeKS3XzniMY8D/NkDhKuiJBv89s/WDg/e1Y/x96YWv2827mfLoRi11Xl0A0PiAxXWIGFG5sI90ppRjfPCnKw4iwbP2e/HbF6GGScDeZQxRmK/9mTNBEgZrzFU3CGLC9E1/wKEo9BNQI=
```