package des

type TripleDES struct {
	blocks [3]*DES
}

func NewTripleDES(k1, k2, k3 []byte) (*TripleDES, error) {
	triple := TripleDES{}
	var err error
	triple.blocks[0], err = NewDES(k1)
	if err != nil {
		return nil, err
	}
	triple.blocks[1], err = NewDES(k2)
	if err != nil {
		return nil, err
	}
	triple.blocks[2], err = NewDES(k3)
	if err != nil {
		return nil, err
	}
	return &triple, nil
}

func (d *TripleDES) Encrypt(block []byte) ([]byte, error) {
	encrypted := make([]byte, len(block))
	copy(encrypted, block)
	for i := range d.blocks[:] {
		var err error
		encrypted, err = d.blocks[i].Encrypt(encrypted)
		if err != nil {
			return nil, err
		}
	}
	return encrypted, nil
}

func (d *TripleDES) Decrypt(block []byte) ([]byte, error) {
	decrypted := make([]byte, len(block))
	copy(decrypted, block)
	for i := 2; i >= 0; i-- {
		var err error
		decrypted, err = d.blocks[i].Decrypt(decrypted)
		if err != nil {
			return nil, err
		}
	}
	return decrypted, nil
}
