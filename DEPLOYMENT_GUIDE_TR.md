# Tizen TV Deployment Guide (Türkçe / Turkish)

Bu rehber, IPTV Player uygulamasını Samsung Tizen TV'nize kurmanız için detaylı adımları içerir.

This guide contains detailed steps to install the IPTV Player application on your Samsung Tizen TV.

---

## Hızlı Kurulum (Quick Installation)

### Seçenek 1: Tizen Studio ile (Önerilen / Recommended)

#### Adım 1: Gerekli Yazılımları İndirin
1. **Tizen Studio** indirin: https://developer.samsung.com/tizen/tizen-studio/download
2. Windows, macOS veya Linux sürümünü seçin

#### Adım 2: Tizen Studio'yu Kurun
1. İndirdiğiniz dosyayı çalıştırın
2. Kurulum sırasında şunları seçin:
   - "TV Extensions" (TV Eklentileri)
   - "TV Extensions-6.0" veya daha yüksek
   - "Certificate Manager" (Sertifika Yöneticisi)
   - "Samsung Certificate Extension"

#### Adım 3: TV'nizi Geliştirici Moduna Alın
1. TV'nizde Smart Hub'ı açın
2. "Apps" (Uygulamalar) bölümüne gidin
3. Uzaktan kumandanızda şu sırayı tuşlayın:
   ```
   1 2 3 4 5 (hızlı bir şekilde)
   ```
4. "Developer Mode" açılacak
5. "ON" seçin
6. Bilgisayarınızın IP adresini girin (Host PC IP)
7. TV'yi yeniden başlatın

#### Adım 4: Sertifika Oluşturun
1. Tizen Studio'yu açın
2. Tools → Certificate Manager
3. "+" butonuna tıklayın
4. "Samsung" seçin
5. "TV" seçin
6. Hesap bilgilerinizi girin (Samsung Developer hesabı gerekli)
7. Sertifikayı oluşturun

#### Adım 5: TV'ye Bağlanın
1. TV'nizin IP adresini öğrenin (Ayarlar → Genel → Ağ → Ağ Durumu → IP Ayarları)
2. Tizen Studio'da: Tools → Device Manager
3. "Remote Device Manager" → "+" 
4. TV'nizin IP adresini girin
5. "Add" → "Connect"

#### Adım 6: Projeyi İçe Aktarın
1. File → Import → Tizen → Tizen Project
2. Bu repository'deki `tizen-app` klasörünü seçin
3. "Finish"

#### Adım 7: Uygulamayı Yükleyin
1. Proje üzerine sağ tıklayın
2. "Run As" → "Tizen Web Application"
3. TV'nizde uygulama otomatik olarak açılacak

---

### Seçenek 2: WGT Dosyası ile Manuel Kurulum

Bu yöntem daha basittir ama Tizen Studio'nun yine de kurulu olması gerekir.

#### Adım 1: WGT Dosyası Oluşturun

##### Windows'ta:
```batch
cd tizen-app
"C:\tizen-studio\tools\ide\bin\tizen.bat" package -t wgt -s <certificate-profile-name> -- .
```

##### macOS/Linux'ta:
```bash
cd tizen-app
~/tizen-studio/tools/ide/bin/tizen package -t wgt -s <certificate-profile-name> -- .
```

Bu komut `tizen-app.wgt` dosyası oluşturacak.

**Not:** `<certificate-profile-name>` yerine Tizen Studio'da oluşturduğunuz sertifika profili adını yazın.

#### Adım 2: WGT Dosyasını TV'ye Yükleyin

##### Yöntem A: Tizen Studio Device Manager ile
1. Tizen Studio → Tools → Device Manager
2. TV'nize bağlanın (yukardaki adımlar gibi)
3. Sağ tıklayın → "Install Application"
4. `tizen-app.wgt` dosyasını seçin

##### Yöntem B: SDB (Smart Development Bridge) ile
```bash
# TV'ye bağlan
~/tizen-studio/tools/sdb connect <TV_IP>:26101

# Bağlantıyı kontrol et
~/tizen-studio/tools/sdb devices

# WGT dosyasını yükle
~/tizen-studio/tools/sdb install tizen-app.wgt
```

---

### Seçenek 3: USB ile Kurulum (2021+ TV'ler için)

#### Adım 1: WGT Dosyası Hazırlayın
1. Yukarıdaki yöntemle `.wgt` dosyası oluşturun
2. Dosyayı bir USB belleğe kopyalayın
3. USB belleği TV'ye takın

#### Adım 2: TV'de Kurulum
1. TV'de "Apps" bölümüne gidin
2. "Settings" (Ayarlar) → "About" → "Developer Mode" açın
3. "Install App" seçeneğini arayın
4. USB'den `.wgt` dosyasını seçin

**Not:** Bu özellik bazı 2021 modellerinde olmayabilir.

---

## Sorun Giderme (Troubleshooting)

### "Author signature is invalid" hatası
- Sertifikanızı doğru oluşturduğunuzdan emin olun
- TV'nin DUID'ini sertifikaya eklediğinizden emin olun

### TV'ye bağlanamıyorum
- TV ve bilgisayarınızın aynı ağda olduğundan emin olun
- TV'de Developer Mode'un açık olduğunu kontrol edin
- Güvenlik duvarınızı kontrol edin (26101 portu açık olmalı)

### Uygulama yüklenmiyor
- TV modelinizin Tizen 6.0 veya üstü olduğundan emin olun
- Sertifikanın TV'niz için oluşturulduğunu kontrol edin

### Video oynatmıyor
- Xtream Codes sunucu bilgilerinizi kontrol edin
- İnternet bağlantınızı kontrol edin
- Sunucunun stream formatının TV'nizle uyumlu olduğundan emin olun

---

## Backend Sunucuyu Başlatma

Uygulamanın tam çalışması için Go backend sunucusunu da çalıştırmanız gerekir:

```bash
# Repoyu klonlayın
git clone https://github.com/abdelsmith21/go.git
cd go

# Bağımlılıkları indirin
go mod download

# main.go'da Xtream Codes bilgilerinizi düzenleyin
# baseURL, username, password değerlerini değiştirin

# Sunucuyu başlatın
go run main.go
```

Sunucu `http://localhost:8080` adresinde çalışacak.

TV uygulamasında, sunucu adresi olarak bilgisayarınızın IP adresini kullanın:
- Örnek: `http://192.168.1.100:8080`

---

## İpuçları

1. **İlk Kurulum:** Uygulamayı ilk açtığınızda:
   - Server URL: Backend sunucunuzun adresi (örn: `http://192.168.1.100:8080`)
   - Username: Xtream Codes kullanıcı adınız
   - Password: Xtream Codes şifreniz

2. **Uzaktan Kumanda Kısayolları:**
   - Kanal Değiştirme: CH+ / CH- tuşları
   - Altyazı: Kırmızı tuş
   - Ses Seçimi: Yeşil tuş
   - Oynat/Durdur: Play/Pause tuşu
   - Geri Dön: Back tuşu

3. **Performans:** İlk açılışta kategoriler yüklenirken biraz bekleyebilirsiniz.

---

## Ek Kaynaklar

- Samsung Developers: https://developer.samsung.com/smarttv
- Tizen Documentation: https://docs.tizen.org/application/web/
- Tizen Studio Guide: https://developer.tizen.org/development/tizen-studio

---

## Destek

Sorun yaşarsanız GitHub'da issue açabilirsiniz:
https://github.com/abdelsmith21/go/issues
