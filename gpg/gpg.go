package gpg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

const OpenPgpURL = `https://keys.openpgp.org/vks/v1/by-email/`

type Encryptor struct {
	PlainMessage *crypto.PlainMessage
	PubKey       *crypto.Key
}

func NewEncryptor() *Encryptor {
	e := &Encryptor{}
	e.PubKey = nil
	e.PlainMessage = nil

	return e
}

func (e *Encryptor) setMessage(message []byte) {
	e.PlainMessage = crypto.NewPlainMessage(message)
}

func (e *Encryptor) AddRecipentFromURL(url *url.URL) {
	pKey, err := e.getGpgPublicKeyFromURL(url)
	if err != nil {
		fmt.Println(err)
		e.PubKey = nil
	}
	e.PubKey = pKey
}
func (e *Encryptor) AddRecipentFromOpenPGPDotOrg(email string) error {
	pKey, err := e.getGpgPublicKeyFromOpenPGPDotOrg(email)
	if err != nil {
		pKey = nil
		return fmt.Errorf("cannot add recipent from opengpg: %w", err)
	}
	e.PubKey = pKey
	return nil
}

func (e *Encryptor) Encrypt(message []byte) (string, error) {
	e.setMessage(message)
	_, err := e.isReadyToEncrypt()
	if err != nil {
		return "", err
	}

	k, _ := e.PubKey.Armor()
	encryptedMessage, err := helper.EncryptBinaryMessageArmored(k, message)
	if err != nil {
		return "", err
	}
	return encryptedMessage, nil
}

func (e Encryptor) isReadyToEncrypt() (bool, error) {
	if e.PubKey != nil {

		if !e.PubKey.CanEncrypt() {
			return false, errors.New("the pubkey cannot encrypt")
		}
		if e.PlainMessage == nil {
			return false, errors.New("there is no message to encrypt, the message is nil")
		}
		if string(e.PlainMessage.Data) == "" {
			return false, errors.New("the message to encrypt is an empty string")
		}

	}
	if e.PubKey == nil {
		return false, errors.New("the pubkey is nil")
	}
	return true, nil
}

func (e Encryptor) getGpgPublicKeyFromURL(u *url.URL) (*crypto.Key, error) {

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	publicKey, err := crypto.NewKeyFromArmoredReader(resp.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	return publicKey, nil
}

func (e Encryptor) getGpgPublicKeyFromOpenPGPDotOrg(email string) (*crypto.Key, error) {

	u, _ := url.Parse(OpenPgpURL + email)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	publicKey, err := crypto.NewKeyFromArmored(string(body))

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	return publicKey, nil
}
