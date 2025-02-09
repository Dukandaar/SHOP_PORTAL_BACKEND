package helper

import (
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// Base64Encode encodes data to Base64.
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode decodes Base64-encoded data.
func Base64Decode(base64String string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64String)
}

// Base64EncodeToString encodes a string to Base64.
func Base64EncodeToString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64DecodeToString decodes a Base64-encoded string to a string.
func Base64DecodeToString(base64String string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// Key Parsing Functions
func ParsePublicKey(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid public key PEM")
	}

	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing public key failed: %w", err)
	}
	pubKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return pubKey, nil
}

// ParsePrivateKey parses a PEM-encoded private key.
func ParsePrivateKey(privateKeyPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("invalid private key PEM")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing private key failed: %w", err)
	}

	return privKey, nil
}

// RSA Encryption/Decryption Functions
// Encrypt encrypts data using RSA public key.
func Encrypt(plaintext string, pubKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(plaintext))
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %w", err)
	}
	return ciphertext, nil
}

// Decrypt decrypts data using RSA private key.
func Decrypt(ciphertext []byte, privKey *rsa.PrivateKey) (string, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}
	return string(plaintext), nil
}

// Generate a JWT containing the Base64-encoded, RSA-encrypted ID
func GenerateJWT(encryptedID string) (string, error) {
	claims := jwt.MapClaims{
		"reg_id": encryptedID,                           // Store the encrypted ID in the JWT
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":    time.Now().Unix(),                     // Issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(utils.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("signing token failed: %w", err)
	}

	return signedToken, nil
}

// Parse a JWT and decrypt the ID
func ParseAndDecryptJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return []byte(utils.JwtSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("parsing token failed: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		encryptedID, ok := claims["reg_id"].(string)
		if !ok {
			return "", fmt.Errorf("invalid reg_id in token")
		}

		// Base64 decode
		decodedEncryptedID, err := base64.StdEncoding.DecodeString(encryptedID)
		if err != nil {
			return "", fmt.Errorf("base64 decode failed: %w", err)
		}

		// RSA decrypt (you'll need the private key)
		privKey, err := ParsePrivateKey(utils.PrivateKeyPEM)
		if err != nil {
			return "", fmt.Errorf("parsing private key failed: %w", err)
		}

		decryptedID, err := Decrypt(decodedEncryptedID, privKey)
		if err != nil {
			return "", fmt.Errorf("rsa decryption failed: %w", err)
		}

		return decryptedID, nil
	}

	return "", fmt.Errorf("invalid token claims")
}
