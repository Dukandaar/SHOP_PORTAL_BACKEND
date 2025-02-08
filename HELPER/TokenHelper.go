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

// RSA Encryption/Decryption Functions
func Encrypt(plaintext string, pubKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(plaintext))
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %w", err)
	}
	return ciphertext, nil
}

// base 64 enryption
func Base64Encode(encryptedID []byte) string {
	base64EncryptedID := base64.StdEncoding.EncodeToString(encryptedID)
	fmt.Println("Base64 Encrypted ID:", base64EncryptedID)
	return base64EncryptedID
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
