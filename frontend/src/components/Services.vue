<script setup>
import { Motion } from 'motion-v'
import { resolveIcon } from '../assets/icons.js'

const ease = [0.16, 1, 0.3, 1]

const services = [
  {
    title: 'Front end web development',
    icons: [
      { name: 'HTML5', slug: 'html5', color: 'E34F26' },
      { name: 'CSS3', slug: 'css3', color: '1572B6' },
      { name: 'JavaScript', slug: 'javascript', color: 'F7DF1E' },
    ],
    desc: 'Modern, responsive interfaces built with semantic HTML, clean CSS and vanilla JavaScript that feel great on every device.',
  },
  {
    title: 'UI UX Design',
    icons: [
      { name: 'Photoshop', slug: 'adobephotoshop', color: '31A8FF' },
      { name: 'Figma', slug: 'figma', color: 'F24E1E' },
      { name: 'Illustrator', slug: 'adobeillustrator', color: 'FF9A00' },
    ],
    desc: 'User-centered design from wireframe to polished prototype, backed by consistent design systems and usability testing.',
  },
  {
    title: 'CMS website',
    icons: [
      { name: 'WordPress', slug: 'wordpress', color: '21759B' },
      { name: 'Shopify', slug: 'shopify', color: '7AB55C' },
      { name: 'Wix', slug: 'wix', color: 'FFFFFF' },
    ],
    desc: 'Content-driven websites and storefronts on popular CMS platforms, easy for your team to manage and scale.',
  },
  {
    title: 'Single page application',
    icons: [
      { name: 'Tailwind', slug: 'tailwindcss', color: '38BDF8' },
      { name: 'Next.js', slug: 'nextdotjs', color: 'FFFFFF' },
      { name: 'Redux', slug: 'redux', color: '764ABC' },
      { name: 'React', slug: 'react', color: '61DAFB' },
    ],
    desc: 'Fast, interactive single page apps with component-driven architecture, state management and API integration.',
  },
]

const iconUrl = (icon) => resolveIcon(icon.slug, icon.color)
</script>

<template>
  <section id="services" class="section">
    <div class="container">
      <Motion
        as="h2"
        class="section-title"
        :initial="{ opacity: 0, y: 24 }"
        :while-in-view="{ opacity: 1, y: 0 }"
        :viewport="{ once: true, amount: 0.4 }"
        :transition="{ duration: 0.6, ease }"
      >
        <span class="highlight">My</span> services
      </Motion>

      <div class="services-grid">
        <Motion
          v-for="(service, i) in services"
          :key="service.title"
          as="article"
          class="service-card"
          :initial="{ opacity: 0, y: 40 }"
          :while-in-view="{ opacity: 1, y: 0 }"
          :viewport="{ once: true, amount: 0.15 }"
          :transition="{ duration: 0.55, delay: i * 0.1, ease }"
          :while-hover="{ y: -6, scale: 1.02 }"
        >
          <span class="service-badge">{{ service.title }}</span>

          <div class="service-icons">
            <img
              v-for="icon in service.icons"
              :key="icon.name"
              :src="iconUrl(icon)"
              :alt="icon.name"
              class="service-icon"
              loading="lazy"
            />
          </div>

          <p class="service-desc">{{ service.desc }}</p>
        </Motion>
      </div>
    </div>
  </section>
</template>

<style scoped>
.section-title {
  text-align: center;
  font-size: 1.75rem;
  font-weight: 800;
  letter-spacing: -0.02em;
  margin-bottom: 40px;
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.service-card {
  background: linear-gradient(180deg, var(--bg-card-2), var(--bg-card));
  border: 1px solid var(--border-card);
  border-radius: var(--radius);
  padding: 16px;
  text-align: center;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.25);
}

.service-badge {
  display: block;
  background: var(--bg-badge);
  color: var(--text);
  font-size: 0.7rem;
  font-weight: 600;
  padding: 6px 10px;
  border-radius: 6px;
  margin-bottom: 16px;
}

.service-icons {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 16px;
}

.service-icon {
  width: 24px;
  height: 24px;
  object-fit: contain;
  transition: transform 0.2s ease;
}

.service-icon:hover {
  transform: scale(1.2);
}

.service-desc {
  color: var(--text-muted);
  font-size: 0.78rem;
  line-height: 1.55;
}

@media (max-width: 1024px) {
  .services-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 560px) {
  .services-grid {
    grid-template-columns: 1fr;
  }
}
</style>