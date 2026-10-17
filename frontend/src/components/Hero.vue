<script setup>
import { Motion } from 'motion-v'
import heroImg from '../assets/verstappen.jpg'
import { resolveIcon } from '../assets/icons.js'

const techIcons = [
  { name: 'Tailwind', slug: 'tailwindcss', color: '38BDF8' },
  { name: 'Next.js', slug: 'nextdotjs', color: 'FFFFFF' },
  { name: 'Redux', slug: 'redux', color: '764ABC' },
  { name: 'Figma', slug: 'figma', color: 'F24E1E' },
  { name: 'HTML5', slug: 'html5', color: 'E34F26' },
  { name: 'CSS3', slug: 'css3', color: '1572B6' },
  { name: 'JavaScript', slug: 'javascript', color: 'F7DF1E' },
  { name: 'React', slug: 'react', color: '61DAFB' },
]

const iconUrl = (icon) => resolveIcon(icon.slug, icon.color)
const ease = [0.16, 1, 0.3, 1]

const strip = {
  hidden: {},
  show: { transition: { staggerChildren: 0.05, delayChildren: 0.3 } },
}
const iconItem = {
  hidden: { opacity: 0, y: 15, scale: 0.7 },
  show: { opacity: 1, y: 0, scale: 1, transition: { type: 'spring', stiffness: 280, damping: 20 } },
}
</script>

<template>
  <section id="about" class="hero section">
    <div class="container hero-grid">
      <Motion
        as="div"
        class="hero-photo"
        :initial="{ opacity: 0, scale: 0.8, rotate: -6 }"
        :animate="{ opacity: 1, scale: 1, rotate: 0 }"
        :transition="{ duration: 0.9, ease }"
      >
        <div class="photo-circle-wrapper">
          <div class="glow-ring"></div>
          <div class="photo-frame">
            <img :src="heroImg" alt="Max Verstappen" />
          </div>
        </div>
      </Motion>

      <div class="hero-copy">
        <Motion
          as="h1"
          class="hero-title"
          :initial="{ opacity: 0, y: 30 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.9, delay: 0.15, ease }"
        >
          Web <span class="highlight">DEVELOPER</span><br />
          And UI UX <span class="highlight">DESIGNER</span>
        </Motion>

        <Motion
          as="div"
          class="tech-strip"
          :variants="strip"
          initial="hidden"
          animate="show"
        >
          <Motion
            v-for="icon in techIcons"
            :key="icon.name"
            :variants="iconItem"
            as="img"
            :src="iconUrl(icon)"
            :alt="icon.name"
            class="tech-icon"
            loading="lazy"
            :while-hover="{ scale: 1.25, y: -4, rotate: 6 }"
            :while-tap="{ scale: 0.95 }"
          />
        </Motion>

        <Motion
          as="p"
          class="hero-desc"
          :initial="{ opacity: 0, y: 20 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.7, delay: 0.45, ease }"
        >
          I build fast, accessible and pixel-perfect web experiences. From
          responsive interfaces to scalable single page applications, I turn
          ideas into clean, production-ready products.
        </Motion>

        <Motion
          as="a"
          href="#contact"
          class="btn-accent"
          :initial="{ opacity: 0, y: 20 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.7, delay: 0.6, ease }"
          :while-hover="{ scale: 1.06, y: -3 }"
          :while-tap="{ scale: 0.95 }"
        >
          Contact me
        </Motion>
      </div>
    </div>
  </section>
</template>

<style scoped>
.hero {
  padding-top: 56px;
  padding-bottom: 72px;
}

.hero-grid {
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 56px;
  align-items: center;
}

.hero-photo {
  display: flex;
  justify-content: center;
}

.photo-circle-wrapper {
  position: relative;
  width: 270px;
  height: 270px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.glow-ring {
  position: absolute;
  inset: -6px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.5), rgba(147, 51, 234, 0.5));
  filter: blur(14px);
  opacity: 0.75;
}

.photo-frame {
  position: relative;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  padding: 6px;
  border: 2px solid rgba(59, 130, 246, 0.5);
  background: radial-gradient(circle at 30% 30%, rgba(30, 58, 138, 0.6), rgba(15, 23, 42, 0.9));
  box-shadow: 0 20px 45px rgba(0, 0, 0, 0.6), inset 0 0 25px rgba(59, 130, 246, 0.3);
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
}

.photo-frame img {
  display: block;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.hero-title {
  font-size: clamp(2rem, 4vw, 2.75rem);
  font-weight: 800;
  letter-spacing: -0.02em;
  margin-bottom: 20px;
}

.tech-strip {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin-bottom: 24px;
}

.tech-icon {
  width: 30px;
  height: 30px;
  object-fit: contain;
  cursor: pointer;
}

.hero-desc {
  color: var(--text-muted);
  font-size: 0.95rem;
  max-width: 520px;
  margin-bottom: 28px;
}

.btn-accent {
  display: inline-block;
  background: var(--accent);
  color: #fff;
  font-size: 0.85rem;
  font-weight: 600;
  padding: 11px 26px;
  border-radius: 9999px;
  box-shadow: 0 8px 24px rgba(147, 51, 234, 0.35);
}

@media (max-width: 768px) {
  .hero-grid {
    grid-template-columns: 1fr;
    gap: 32px;
    text-align: center;
  }

  .photo-circle-wrapper {
    width: 220px;
    height: 220px;
    margin: 0 auto;
  }

  .hero-desc {
    margin-left: auto;
    margin-right: auto;
  }

  .tech-strip {
    justify-content: center;
  }
}
</style>
