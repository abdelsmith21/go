# Hızlı Başlangıç - TV'ye Kurulum (Quick Start)

## Senaryo 1: Windows Kullanıyorsanız

### Adım 1: Tizen Studio İndirin
1. https://developer.samsung.com/tizen/tizen-studio/download adresine gidin
2. "Download" butonuna tıklayın
3. Windows için indirin ve kurun
4. Kurulum sırasında şunları seçin:
   - ✅ TV Extensions
   - ✅ TV Extensions-6.0
   - ✅ Certificate Manager

### Adım 2: TV'nizi Hazırlayın
1. TV'nizde Smart Hub'ı açın
2. "Apps" bölümüne gidin  
3. Uzaktan kumandada **hızlıca** şu tuşlara basın: **1 2 3 4 5**
4. Developer Mode ekranı açılacak → "ON" yapın
5. Bilgisayarınızın IP adresini girin (cmd'de `ipconfig` yazarak bulabilirsiniz)
6. TV'yi yeniden başlatın

### Adım 3: Sertifika Oluşturun
1. Tizen Studio'yu açın
2. Tools → Certificate Manager
3. "+" butonuna tıklayın
4. Samsung → TV seçin
5. Samsung hesabınızla giriş yapın (yoksa oluşturun)
6. TV'nizin DUID'ini ekleyin (TV'de Developer Mode ekranında gösterilir)
7. Sertifikayı oluşturun

### Adım 4: Build Edin
1. Bu projeyi indirin:
   ```
   git clone https://github.com/abdelsmith21/go.git
   cd go
   ```

2. `build-tizen.bat` dosyasını çift tıklayarak çalıştırın

3. Sertifika profili adını girin (Certificate Manager'da oluşturduğunuz)

4. Script `tizen-app.wgt` dosyasını oluşturacak

### Adım 5: TV'ye Yükleyin
1. Tizen Studio'da: Tools → Device Manager
2. Remote Device Manager → "+"
3. TV'nizin IP adresini girin
4. "Connect" butonuna tıklayın
5. TV'ye sağ tıklayın → "Install Application"
6. `tizen-app.wgt` dosyasını seçin
7. Uygulama TV'nize yüklenecek!

---

## Senaryo 2: Linux/macOS Kullanıyorsanız

Aynı adımlar, sadece:
- Adım 1'de Linux/macOS sürümünü indirin
- Adım 4'te `./build-tizen.sh` komutunu çalıştırın
- Terminal'de şu komutları kullanın:
  ```bash
  # TV'ye bağlan
  ~/tizen-studio/tools/sdb connect [TV_IP]:26101
  
  # Yükle
  ~/tizen-studio/tools/sdb install tizen-app.wgt
  ```

---

## Backend Sunucu (ZORUNLU!)

Uygulamanın çalışması için Go backend'i de çalıştırmalısınız:

### 1. Go'yu İndirin
https://go.dev/dl/ - En son sürümü indirip kurun

### 2. Backend'i Çalıştırın
```bash
cd go  # Proje klasörüne gidin

# Xtream Codes bilgilerinizi main.go'da düzenleyin
# Satır 345-347'yi kendi bilgilerinizle değiştirin:
# baseURL := "http://sizin-sunucu.com:8080"
# username := "sizin_kullanici_adi"
# password := "sizin_sifre"

# Sunucuyu başlatın
go run main.go
```

Sunucu `http://localhost:8080` adresinde çalışacak.

### 3. TV Uygulamasını Yapılandırın
TV'de uygulama açıldığında:
- Server URL: `http://[BILGISAYAR_IP]:8080` (örn: `http://192.168.1.100:8080`)
- Username: Xtream Codes kullanıcı adınız
- Password: Xtream Codes şifreniz

**Not:** Bilgisayar IP'nizi bulmak için:
- Windows: `ipconfig` komutunu çalıştırın
- Linux/Mac: `ifconfig` veya `ip addr` komutunu çalıştırın

---

## Sorun mu Yaşıyorsunuz?

### "Tizen Studio bulunamadı" hatası
- Tizen Studio'yu doğru yere kurdunuz mu?
- Windows: `C:\tizen-studio`
- Linux/Mac: `~/tizen-studio`

### "Sertifika profili yok" hatası
- Certificate Manager'da sertifika oluşturdunuz mu?
- TV'nizin DUID'ini sertifikaya eklediniz mi?

### TV'ye bağlanamıyorum
- TV ve bilgisayar aynı ağda mı?
- TV'de Developer Mode açık mı?
- TV'yi yeniden başlattınız mı?

### Daha fazla yardım
Detaylı Türkçe rehber: [DEPLOYMENT_GUIDE_TR.md](DEPLOYMENT_GUIDE_TR.md)

---

## Özet: En Hızlı Yol

```
1. Tizen Studio kur → Sertifika oluştur
2. TV'de Developer Mode aç (1-2-3-4-5 tuşları)
3. build-tizen.bat (Windows) veya ./build-tizen.sh (Linux/Mac) çalıştır
4. Tizen Studio Device Manager ile TV'ye yükle
5. main.go'yu düzenle ve go run main.go ile backend'i başlat
6. TV'de uygulamayı aç ve backend IP'sini gir
```

İyi seyirler! 🎬📺
