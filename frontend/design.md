# Design Spec — Personal Portfolio (Tylenol)

> Sumber: Figma proto `qcf35nGjGTTViQnIQjkQPV` node `15:640` — 4 frame screenshot (Hero, Services, Projects+Clients, Contact). Figma link diblokir 403 tanpa auth, jadi spec di-derive dari visual screenshot.

## 1. Overview

Landing page one-page portfolio dark theme. Scroll vertikal: `Navbar → Hero → My Services → Recent Projects → Our Clients → Contact → Footer`. Aesthetic: dark navy gradient dengan grid dotted background, aksen biru elektrik + ungu, kartu semi-transparan dengan border tipis.

Stack target: `Vue 3 + Vite` (sudah ada `frontend/src/App.vue`, `Navbar.vue`).

## 2. Theme & Tokens

| Token | Value | Penggunaan |
|-------|-------|------------|
| `bg` | `#0A1931` → `#0F172A` gradient, + grid pattern `rgba(255,255,255,0.04)` | body background |
| `bg-card` | `#132040` / `#162447` `rgba(19,32,64,0.8)` | card services & projects |
| `bg-card-header` | `#1E3458` / `#243A63` | header kecil di atas card Services |
| `border-card` | `rgba(59,130,246,0.15)` / `rgba(255,255,255,0.08)` | border kartu |
| `primary` | `#2563EB` / `#3B82F6` | kata kunci heading, border form, submit btn |
| `accent` | `#7C3AED` → `#9333EA` / `#A855F7` | CTA `Contact me`, `Visit site` pill |
| `text` | `#F1F5F9` / `#E2E8F0` | heading putih |
| `text-muted` | `#94A3B8` / `#9CA3AF` | paragraf deskripsi |
| `pill-bg` | `#E5E7EB` / `#D1D5DB` | bar logo clients |
| `font` | `Inter` / `Poppins`, sans | seluruh UI |
| `radius-card` | `12-16px` | services & project cards |
| `radius-pill` | `9999px` | button & clients bar |
| `shadow` | `0 8px 32px rgba(0,0,0,0.4)` | card |

Efek global: `background: radial-gradient` biru di sudut kiri-bawah & kanan-atas, plus grid lines halus di area hero.

## 3. Layout & Sections

### 3.1 Navbar
- Sticky top, `height ~64px`, bg `#0B1A33` semi-transparan, border bottom halus.
- Kiri: logo text `Tylenol` putih bold.
- Kanan: link `About me | Services | Projects | Contact` (text 14px, `text-muted`, hover putih). Anchor ke section id.
- Mobile: hamburger → dropdown.

### 3.2 Hero
- Dua kolom (kiri foto, kanan copy) di desktop; stack vertikal di mobile.
- Heading besar: `Web DEVELOPER / And UI UX DESIGNER` — `DEVELOPER` & `DESIGNER` warna `primary` biru. Font 42-48px bold, line-height ketat.
- Tech strip di bawah heading: icon row `Tailwind, Next.js, Redux/Zustand, Figma, HTML5, CSS3, JS, React` (ukuran 28-32px, gap 12px).
- Foto: card 4:5, border `rgba(59,130,246,0.3)` + padding 6px, `border-radius 12px`, object-cover.
- Deskripsi: paragraf 14-15px `text-muted`, max-width ~520px.
- CTA: `Contact me` pill kecil ungu `bg #9333EA`, text putih 13px, `padding 10px 22px`.

### 3.3 My Services
- Heading center: `My services` — `My` biru, `services` putih, 28px bold.
- Grid: 4 kolom desktop, 2 kolom tablet, 1 kolom mobile. Gap 20px.
- Card: bg `bg-card`, border tipis, `radius 12px`, `padding 16px` + header badge di atas.
  - Badge header: pill persegi kecil `bg #1E3458`, text 11px putih, center.
  - Icon row 24px per tech (tengah).
  - Deskripsi: text 12-13px `text-muted` center, 4 baris.
- 4 kartu:
  1. `Front end web development` — HTML/CSS/JS
  2. `UI UX Design` — Ps/Figma/Ai
  3. `CMS website` — WordPress/Shopify/Wix
  4. `Single page application` — Tailwind/Next/Redux/React

### 3.4 Some of my recent projects
- Heading center: `Some of my recent / projects` — `projects` biru.
- Grid 2×2, gap 20px.
- Project card: `bg #0F2A2E` / `#0F2530` gelap kehijauan, `radius 16px`, `padding 20px`, `display flex` (info kiri, preview kanan). Overflow hidden, ada watermark icon besar transparan di background (JS/React/WordPress sesuai project).
  - Kiri: Title 16px putih bold, tech dots 4× `w16 h16 circle` kecil, logo `S / Sandeep Kumar` (placeholder brand, biru), button `Visit site` pill ungu full-width kecil.
  - Kanan: preview image dengan border ungu `3px solid #7C3AED`, `radius 8px`, aspect 16:9.
- Daftar project:
  - `Mega charts` — chart dashboard
  - `Naked Beauty New York` — beauty brand hero
  - `Cesars Tire` — 3-card blog
  - `RadhaMadav Dham` — temple/NGO landing

### 3.5 Our Clients
- Heading center: `OUR / CLIENTS` putih, `OUR` 18px, `CLIENTS` 36px extra-bold.
- Pill bar: `bg #E5E7EB`, `radius 9999px`, `padding 16px 32px`, flex row logo hitam monokrom, `justify space-between`, gap 24px, overflow-x auto di mobile.
- Logo: Apple, MI, Nike, Adidas, PUMA, HP, ASUS, AMD (versi hitam silhouette).

### 3.6 Contact
- Heading center: `Contact` putih 32px bold.
- Dua kolom: kiri ilustrasi, kanan form. Stack di mobile.
- Ilustrasi: karakter ungu-putih melompat menempel laptop/dokumen, shape blob `#8B7FE8` di belakang, plus dekor `X` dan `▲` kecil.
- Form card: `bg #0F172A`, `border 2px solid #2563EB`, `radius 4px`, `padding 0`.
  - Header: `Contact with me for any kind of / Queries` — `Queries` biru 20px bold.
  - Field: `Name`, `Email` (input text), `Message` (textarea 100px), style underline only — `bg transparent`, `border-bottom 1px solid rgba(255,255,255,0.1)`, placeholder `text-muted` 12px.
  - Submit: `bg #2563EB`, text putih, full-width, `height 44px`.
- Footer sosmed row center: icon bulat kecil Facebook/Discord/WhatsApp/LinkedIn/Upwork/Twitter/WhatsApp (16px, `text-muted` hover putih).
- Copyright: `Copyright @2026 all right reserved` 12px `text-muted`, dengan garis horizontal tipis di kiri-kanan (flex divider).

## 4. Komponen Reusable

| Komponen | Props | Catatan |
|----------|-------|---------|
| `Navbar` | `links[]` | sudah ada — sesuaikan warna ke `#0B1A33`, link `About me/Services/Projects/Contact` |
| `Hero` | `title, techIcons[], photo, cta` | grid + watermark |
| `ServiceCard` | `title, icons[], desc` | badge header |
| `ProjectCard` | `title, stack[], image, href` | watermark bg icon |
| `ClientsBar` | `logos[]` | pill horizontal scroll |
| `ContactForm` | `onSubmit` | validasi required, underline input |
| `SocialRow` | `icons[]` | footer |

## 5. Responsive

- Breakpoint: `768px` (tablet), `1024px` (desktop).
- Hero & Contact: `flex-col` di `<768px`.
- Services: `grid-cols-1 → 2 → 4`.
- Projects: `grid-cols-1 → 2`.
- Clients pill: `overflow-x auto`, `scroll-snap` di mobile.
- Navbar: hamburger `<768px`.

## 6. Aset yang Dibutuhkan

- Foto profil (portrait) + semua tech logo SVG (Tailwind, Next, Redux, Figma, HTML, CSS, JS, React, Ps, Ai, WP, Shopify, Wix).
- Screenshot preview 4 project (sudah ada di Figma — export 800×500).
- Logo clients SVG monokrom.
- Ilustrasi contact (export SVG/PNG dari Figma).

## 7. Catatan Implementasi

- Anchor smooth scroll: `html { scroll-behavior: smooth }`, offset untuk sticky navbar.
- Grid background hero: `background-image: linear-gradient(rgba(255,255,255,0.03) 1px, transparent 1px), linear-gradient(90deg, ...)` atau SVG pattern.
- Semua card hover: `transform translateY(-4px) + shadow` transisi 200ms.
- Form: controlled input Vue `v-model`, submit prevent + toast.
- Branding `Sandeep Kumar` di project card — ganti jadi `Tylenol` saat implement.
- Aksesibilitas: `alt` untuk image, `aria-label` nav, contrast `text-muted` minimal 4.5:1 di bg gelap.
