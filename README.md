# WORQ Todo App

Bu proje TDD metodolojisi kullanılarak geliştirilmiş full-stack bir Todo uygulamasıdır.

## 🏗️ Proje Yapısı
worq-todo-app/
├── frontend/ # React + TypeScript + Tailwind
│ ├── src/
│ │ ├── components/ # React bileşenleri
│ │ ├── hooks/ # Custom hooks
│ │ ├── services/ # API çağrıları
│ │ ├── types/ # TypeScript tipleri
│ │ ├── utils/ # Yardımcı fonksiyonlar
│ │ └── tests/ # Test dosyaları
│ ├── cypress/ # E2E testler
│ └── package.json
├── backend/ # Golang API
│ ├── cmd/ # Ana uygulama giriş noktası
│ ├── internal/ # İç paketler
│ │ ├── handler/ # HTTP handlers
│ │ ├── service/ # Business logic
│ │ ├── repository/ # Database katmanı
│ │ └── models/ # Veri modelleri
│ ├── pkg/ # Dışa açık paketler
│ └── go.mod
├── deployment/ # Deployment dosyaları
│ ├── docker/ # Dockerfile'lar
│ └── kubernetes/ # K8s YAML'ları
├── docs/ # Dokümantasyon
└── .github/workflows/ # CI/CD pipeline

## 🎯 TDD Süreci

Bu projede **Test Driven Development (TDD)** metodolojisi kullanılmaktadır:

### TDD Döngüsü: Red → Green → Refactor

1. **🔴 RED**: Önce failing test yaz
2. **🟢 GREEN**: Test'i geçecek minimum kod yaz  
3. **🔄 REFACTOR**: Kodu iyileştir, testler hala geçsin

## 🚀 Kurulum ve Çalıştırma

*(Bu bölüm geliştirme sürecinde güncellenecek)*

---

**Geliştirme süreci:** TDD ile adım adım ilerliyoruz!