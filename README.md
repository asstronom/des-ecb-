# des-ecb-
Команди для Windows:\n
Побудувати проект: go build -o encrypt.exe\n
Команди для DES\n
Зашифрувати: ./encrypt.exe -key akasimov -i sample.txt -o sample_enc.txt\n
Розшифрувати: ./encrypt.exe -key akasimov -i sample_enc.txt -o sample_dec.txt -decrypt\n
Команди для Triple DES:\n
Зашифрувати: ./encrypt.exe -triple -i sample.txt -o sample_enc_triple.txt\n
Розшифрувати: ./encrypt.exe -triple -o sample_dec_triple.txt -i sample_enc_triple.txt -decrypt\n
