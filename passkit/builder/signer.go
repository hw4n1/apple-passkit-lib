package builder

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log/slog"
	"os"

	"go.mozilla.org/pkcs7"
)

func getKeyLocationsFromEnv() (map[string]string, bool) {
	keysLocation := make(map[string]string)

	value, ok := os.LookupEnv("SIGNER_KEY_PATH")

	if !ok {
		slog.Warn("signer key is missing. unable to sign pass")
		return nil, false
	}
	keysLocation["signerKey"] = value

	value, ok = os.LookupEnv("SIGNER_CERT_PATH")
	if !ok {
		slog.Warn("signer certificate is missing. unable to sign pass")
		return nil, false
	}
	keysLocation["signerCert"] = value

	value, ok = os.LookupEnv("WWDR_CERT_PATH")
	if !ok {
		slog.Warn("Apple WWDR certificate is missing. unable to sign pass")
		return nil, false
	}
	keysLocation["wwdrCert"] = value

	return keysLocation, true
}

func SignPass(passName string) bool {
	keysLocation, found := getKeyLocationsFromEnv()

	if !found {
		slog.Warn("unable to load signing keys")
		return false
	}

	passPath := passName + ".pass"
	manifestPath := passPath + "/manifest.json"
	signaturePath := passPath + "/signature"

	// Load signer certificate and private key
	tlsCert, err := tls.LoadX509KeyPair(keysLocation["signerCert"], keysLocation["signerKey"])

	if err != nil {
		slog.Error("error loading certificates", "err", err)
		return false
	}

	signerCert, _ := x509.ParseCertificate(tlsCert.Certificate[0])
	privateKey := tlsCert.PrivateKey.(*rsa.PrivateKey)

	// Load WWDR cert
	wwdrPem, _ := os.ReadFile(keysLocation["wwdrCert"])
	block, _ := pem.Decode(wwdrPem)
	wwdrCert, _ := x509.ParseCertificate(block.Bytes)

	// Read manifest
	manifest, _ := os.ReadFile(manifestPath)

	// Perform sign - Use SHA256 for content digest
	signedData, err := pkcs7.NewSignedData(manifest)

	if err != nil {
		slog.Error("error creating pkcs7 data for manifest.", "err", err)
		return false
	}

	signedData.SetDigestAlgorithm(pkcs7.OIDDigestAlgorithmSHA256)
	err = signedData.AddSigner(signerCert, privateKey, pkcs7.SignerInfoConfig{})

	if err != nil {
		slog.Error("error signing pass manifest.", "err", err)
		return false
	}

	signedData.AddCertificate(wwdrCert)
	signedData.Detach()

	der, _ := signedData.Finish()
	err = os.WriteFile(signaturePath, der, 0644)

	if err != nil {
		slog.Error("error writing pass signature.", "err", err)
		return false
	}

	slog.Info("pass successfully signed")

	return true
}
