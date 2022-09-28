# des-ecb-
Команди для Windows:
Побудувати проект: go build -o encrypt.exe
Команди для DES
Зашифрувати: ./encrypt.exe -key akasimov -i sample.txt -o sample_enc.txt
Розшифрувати: ./encrypt.exe -key akasimov -i sample_enc.txt -o sample_dec.txt -decrypt
Команди для Triple DES:
Зашифрувати: ./encrypt.exe -triple -i sample.txt -o sample_enc_triple.txt
Розшифрувати: ./encrypt.exe -triple -o sample_dec_triple.txt -i sample_enc_triple.txt -decrypt
