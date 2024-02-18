package ipfs

import (
	"bytes"
	"context"
	"github.com/genovatix/algoliocdm/log"
	"github.com/genovatix/algoliocdm/pkg/cryptography"
	"github.com/genovatix/algoliocdm/pkg/storage"
	shell "github.com/ipfs/go-ipfs-api"
	"go.dedis.ch/kyber/v3"
	"go.uber.org/zap"
	"io"
)

// IPFSStorageDriver implements the StorageDriver interface for IPFS.
type IPFSStorageDriver struct {
	shell      *shell.Shell
	keyManager cryptography.KeyManager
}

// NewIPFSStorageDriver initializes a new IPFS storage driver.
func NewIPFSDriver(ipfsNodeAddress string, keyManager cryptography.KeyManager) *IPFSStorageDriver {
	return &IPFSStorageDriver{
		shell:      shell.NewShell(ipfsNodeAddress),
		keyManager: keyManager,
	}
}

// Create stores data in IPFS and returns the hash of the stored content.
func (ipfs *IPFSStorageDriver) Create(ctx context.Context, publicKey kyber.Point, data []byte) (string, error) {
	// Encrypt the data before adding it to IPFS.
	encryptedData, err := ipfs.keyManager.Encrypt(publicKey, data)
	if err != nil {
		log.Logger.Error("Failed to encrypt data", zap.Error(err))
		return "", err
	}

	cid, err := ipfs.shell.Add(bytes.NewReader(encryptedData))
	if err != nil {
		log.Logger.Error("Failed to add encrypted data to IPFS", zap.Error(err))
		return "", err
	}

	log.Logger.Info("Encrypted data added to IPFS", zap.String("cid", cid))
	return cid, nil
}

// Read retrieves data from IPFS by the content hash.
func (ipfs *IPFSStorageDriver) Read(ctx context.Context, privateKey kyber.Scalar, cid string) ([]byte, error) {
	readCloser, err := ipfs.shell.Cat(cid)
	if err != nil {
		log.Logger.Error("Failed to retrieve data from IPFS", zap.String("cid", cid), zap.Error(err))
		return nil, err
	}
	defer readCloser.Close()

	encryptedData, err := io.ReadAll(readCloser)
	if err != nil {
		return nil, err
	}

	// Decrypt the data after retrieving it from IPFS.
	data, err := ipfs.keyManager.Decrypt(privateKey, encryptedData)
	if err != nil {
		log.Logger.Error("Failed to decrypt data", zap.Error(err))
		return nil, err
	}

	return data, nil
}

// Update is not directly supported by IPFS as it is immutable; data must be re-added for changes.
func (ipfs *IPFSStorageDriver) Update(key string, value []byte) error {
	log.Logger.Warn("Update operation is not supported by IPFS; use Create to add new data")
	return storage.ErrStorageFail
}

// Delete is not directly supported by IPFS as it is immutable; data will be garbage collected if not pinned.
func (ipfs *IPFSStorageDriver) Delete(key string) error {
	log.Logger.Warn("Delete operation is not supported by IPFS")
	return storage.ErrStorageFail
}
