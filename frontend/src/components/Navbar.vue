<script setup>
import { ref } from 'vue'

const isMobileMenuOpen = ref(false)

const navLinks = [
  { name: 'About me', href: '#about' },
  { name: 'Services', href: '#services' },
  { name: 'Projects', href: '#projects' },
  { name: 'Contact', href: '#contact' },
]

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}
</script>

<template>
  <header class="navbar">
    <div class="navbar-container">
      <a href="#about" class="logo">Tylenol</a>

      <nav class="nav-menu" aria-label="Main Navigation">
        <ul class="nav-list">
          <li v-for="link in navLinks" :key="link.name" class="nav-item">
            <a :href="link.href" class="nav-link">{{ link.name }}</a>
          </li>
        </ul>
      </nav>

      <button
        class="mobile-toggle"
        :class="{ active: isMobileMenuOpen }"
        @click="toggleMobileMenu"
        aria-label="Toggle navigation menu"
        :aria-expanded="isMobileMenuOpen"
      >
        <span class="hamburger-bar"></span>
        <span class="hamburger-bar"></span>
        <span class="hamburger-bar"></span>
      </button>
    </div>

    <div class="mobile-menu" :class="{ open: isMobileMenuOpen }">
      <ul class="mobile-nav-list">
        <li v-for="link in navLinks" :key="link.name">
          <a :href="link.href" class="mobile-nav-link" @click="isMobileMenuOpen = false">
            {{ link.name }}
          </a>
        </li>
      </ul>
    </div>
  </header>
</template>

<style scoped>
.navbar {
  position: sticky;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
  background: rgba(11, 26, 51, 0.85);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid var(--border-soft);
  padding: 0;
  display: block;
}

.navbar-container {
  max-width: 1120px;
  margin: 0 auto;
  padding: 0 24px;
  height: 64px;
  display: flex !important;
  flex-direction: row !important;
  flex-wrap: nowrap !important;
  justify-content: space-between !important;
  align-items: center !important;
}

.logo {
  font-weight: 700;
  font-size: 1.35rem;
  letter-spacing: -0.02em;
  color: var(--text);
  transition: opacity 0.2s ease;
  flex-shrink: 0;
}

.logo:hover {
  opacity: 0.85;
}

.nav-menu {
  display: flex;
  align-items: center;
}

.nav-list {
  display: flex;
  align-items: center;
  gap: 1.75rem;
  list-style: none;
  margin: 0;
  padding: 0;
}

.nav-link {
  color: var(--text-muted);
  font-size: 0.875rem;
  font-weight: 500;
  position: relative;
  padding: 0.4rem 0;
  transition: color 0.25s ease;
  white-space: nowrap;
}

.nav-link::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 2px;
  background: var(--primary-light);
  border-radius: 2px;
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.nav-link:hover {
  color: var(--text);
}

.nav-link:hover::after {
  width: 100%;
}

.mobile-toggle {
  display: none;
  flex-direction: column;
  justify-content: center;
  gap: 5px;
  width: 40px;
  height: 40px;
  background: transparent;
  border: none;
  padding: 0;
}

.hamburger-bar {
  display: block;
  width: 24px;
  height: 2px;
  background: var(--text);
  border-radius: 2px;
  margin: 0 auto;
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.mobile-toggle.active .hamburger-bar:nth-child(1) {
  transform: translateY(7px) rotate(45deg);
}

.mobile-toggle.active .hamburger-bar:nth-child(2) {
  opacity: 0;
}

.mobile-toggle.active .hamburger-bar:nth-child(3) {
  transform: translateY(-7px) rotate(-45deg);
}

.mobile-menu {
  display: none;
  background: #0b1a33;
  border-bottom: 1px solid var(--border-soft);
  overflow: hidden;
  max-height: 0;
  transition: max-height 0.3s ease;
}

.mobile-menu.open {
  max-height: 260px;
}

.mobile-nav-list {
  list-style: none;
  margin: 0;
  padding: 1rem 1.5rem 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.mobile-nav-link {
  color: var(--text);
  font-size: 1rem;
  font-weight: 500;
  display: block;
  padding: 0.5rem 0;
  transition: color 0.2s ease, transform 0.2s ease;
}

.mobile-nav-link:hover {
  color: var(--primary-light);
  transform: translateX(4px);
}

@media (max-width: 860px) {
  .nav-menu {
    display: none;
  }

  .mobile-toggle {
    display: flex;
  }

  .mobile-menu {
    display: block;
  }
}
</style>
