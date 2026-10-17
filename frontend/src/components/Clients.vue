<script setup>
import { Motion } from 'motion-v'

const clients = [
  { name: 'Apple', slug: 'apple' },
  { name: 'Xiaomi', slug: 'xiaomi' },
  { name: 'Nike', slug: 'nike' },
  { name: 'Adidas', slug: 'adidas' },
  { name: 'Puma', slug: 'puma' },
  { name: 'HP', slug: 'hp' },
  { name: 'ASUS', slug: 'asus' },
  { name: 'AMD', slug: 'amd' },
]

const logoUrl = (slug) => `https://cdn.simpleicons.org/${slug}/111827`
</script>

<template>
  <section id="clients" class="section">
    <div class="container">
      <Motion
        as="h2"
        class="clients-title"
        :initial="{ opacity: 0, y: 24 }"
        :while-in-view="{ opacity: 1, y: 0 }"
        :viewport="{ once: true, amount: 0.4 }"
        :transition="{ duration: 0.6, ease: [0.16, 1, 0.3, 1] }"
      >
        <span class="clients-small">OUR</span>
        <span class="clients-big">CLIENTS</span>
      </Motion>

      <Motion
        class="clients-wrapper"
        :initial="{ opacity: 0, y: 30 }"
        :while-in-view="{ opacity: 1, y: 0 }"
        :viewport="{ once: true, amount: 0.2 }"
        :transition="{ duration: 0.6, ease: [0.16, 1, 0.3, 1] }"
      >
        <div class="clients-bar">
          <div class="clients-track">
            <div v-for="group in 2" :key="group" class="clients-group">
              <img
                v-for="client in clients"
                :key="client.name + group"
                :src="logoUrl(client.slug)"
                :alt="client.name"
                class="client-logo"
                loading="lazy"
              />
            </div>
          </div>
        </div>
      </Motion>
    </div>
  </section>
</template>

<style scoped>
.clients-title {
  text-align: center;
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin-bottom: 32px;
}

.clients-small {
  font-size: 1.1rem;
  font-weight: 800;
  letter-spacing: 0.18em;
  color: var(--text);
}

.clients-big {
  font-size: 2.25rem;
  font-weight: 900;
  letter-spacing: 0.06em;
  color: #fff;
}

.clients-wrapper {
  background: #e5e7eb;
  border-radius: 9999px;
  padding: 18px 0;
  overflow: hidden;
}

.clients-bar {
  display: flex;
  width: 100%;
  overflow: hidden;
  position: relative;
}

.clients-track {
  display: flex;
  gap: 48px;
  width: max-content;
  animation: marquee 25s linear infinite;
}

.clients-bar:hover .clients-track {
  animation-play-state: paused;
}

.clients-group {
  display: flex;
  gap: 48px;
  align-items: center;
}

.client-logo {
  height: 26px;
  width: auto;
  object-fit: contain;
  flex-shrink: 0;
  opacity: 0.85;
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.client-logo:hover {
  opacity: 1;
  transform: scale(1.08);
}

@keyframes marquee {
  to {
    transform: translateX(-50%);
  }
}

@media (max-width: 768px) {
  .clients-wrapper {
    border-radius: 24px;
    padding: 16px 0;
  }

  .clients-track,
  .clients-group {
    gap: 32px;
  }

  .client-logo {
    height: 22px;
  }
}
</style>