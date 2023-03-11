package tool

import (
	"fmt"
	"testing"
)

func TestRSASign(t *testing.T) {
	got, err := RSASign([]byte("string"), "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQC3SYirUDkCOXfymgWF6jXzBzjFbGFAE3uINs2guvWUS7tYD/8h\n2WeUq6ejDQ/xcTYi3Y4tNatmfdVThkA6rZ+G6sXS97fyllBZIEGC2K9na6Zl+4yO\n7j7QrP6W32xZr+Z1frCKp4csEHFEc58oBPG/Dsds4ApR7kF38DeG9H3uXwIDAQAB\nAoGAExLP6iP7CsQ4O0LT+E+bNaM2wTS6GhTs8gvh8iwCimAnKs7fWgJpyQrj8w/U\n7Oc/Hvm3ZfUD1TKbFIoP/Qs8Yny9MCV9s1NpFjhLVQfQCjMzCGfq1xbEJf5xNCPw\nHJsWd2UyC+rp3bxobVgHNjznifTusOOPpMq8wMcRcw+3gsECQQDk04l7p/4S4M7+\no0xG7Fmf86FxIV0B0RXKKVF17HQKYmloOPfIjmLBAnb227/ljOfJ1kih5IRsTT+x\nDJzC/5gJAkEAzQ2Um4YXAWSq9irZz5UIkbIIJWX6KKqCyLzZttDOQPPQMSOZl29g\nuQ4K4Vc3UhOngUbOkxxUjNtQRuREFl7dJwJAa1vEOTwMPJc1Bste0je5pQ4NRKK+\nnEeYzYytJ4KUvvqFMdzohDQpqRya7B8V3YFKjqv2z94DMzzbERo1wldg4QJAcm+0\n0wMHjch/vPh3LGlRKfaAo1aBQPbAHIWAv41Svl8TfokOq7wF7+ENY2tIPW8omXZJ\nzUIPUbkH2TiFYzK4GwJBAIqSz77Jh8pdIH2vfjNBgYMFM6h7HSgHUOqKnxuvWZGm\nCxwWKkcUU6vTu6jwwYRXutKHrDqNZs0hMseOWTDd8qs=\n-----END RSA PRIVATE KEY-----")
	fmt.Println(got, err)
}

func TestRSAVerify(t *testing.T) {
	g := RSAVerify([]byte("string"), "SlNYP9F4vJxjk/lEsiwoIC1GxSFtVXCCI8yizwa3KzfCzLbSmUSHEaEvD6LTNHastD8pg5J3PXgtGwbjulnUjZzqNz0Gc13CRJr+slmg3PgMBSO65nRKnAq5GQMlOf+9YU4zBU0Xp0Bs+jYno/fEtjhHmWVQ/jShHuz68UlVCWE=", "/Users/yidejia/www/upi-zero/data/cert/rsa_public_key.pem")
	fmt.Println(g)
}

