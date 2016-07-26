package ex08

/*
In this problem we're given a big list of base64 encoded strings, and we're
given the task of finding out which one has been encoded with AES-ECB.

This should be pretty straightforward. A major failing of ECB mode (and what
allows for chosen-plaintext type attack on it and so on) is that a given 16-byte
block in the input (assuming AES-128) will always produce the same output in
the ciphertext.

So we'll just look for that pattern!
*/

func detectECBmode(ciphertext []byte) bool {
}
