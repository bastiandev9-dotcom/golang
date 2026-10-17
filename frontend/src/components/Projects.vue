<script setup>
import { Motion } from 'motion-v'
import { resolveIcon } from '../assets/icons.js'

const ease = [0.16, 1, 0.3, 1]

const projects = [
  {
    title: 'Mega charts',
    stack: ['javascript', 'react', 'nextdotjs', 'tailwindcss'],
    watermark: 'react',
    brand: 'Tylenol',
    tint: 'linear-gradient(135deg, #1e3a8a, #0f2a4a)',
    accent: '#38bdf8',
  },
  {
    title: 'Naked Beauty New York',
    stack: ['html5', 'css3', 'javascript', 'react'],
    watermark: 'javascript',
    brand: 'Tylenol',
    tint: 'linear-gradient(135deg, #7e22ce, #3b0764)',
    accent: '#e9d5ff',
  },
  {
    title: 'Cesars Tire',
    stack: ['wordpress', 'php', 'css3', 'javascript'],
    watermark: 'wordpress',
    brand: 'Tylenol',
    tint: 'linear-gradient(135deg, #065f46, #0b3b2e)',
    accent: '#6ee7b7',
  },
  {
    title: 'RadhaMadav Dham',
    stack: ['figma', 'html5', 'css3', 'javascript'],
    watermark: 'figma',
    brand: 'Tylenol',
    tint: 'linear-gradient(135deg, #b45309, #7c2d12)',
    accent: '#fcd34d',
  },
]

const iconUrl = (slug, color = 'FFFFFF') => resolveIcon(slug, color)
</script>

<template>
  <section id="projects" class="section">
    <div class="container">
      <Motion
        as="h2"
        class="section-title"
        :initial="{ opacity: 0, y: 24 }"
        :while-in-view="{ opacity: 1, y: 0 }"
        :viewport="{ once: true, amount: 0.4 }"
        :transition="{ duration: 0.6, ease }"
      >
        Some of my recent<br /><span class="highlight">projects</span>
      </Motion>

      <div class="projects-grid">
        <Motion
          v-for="(project, i) in projects"
          :key="project.title"
          as="article"
          class="project-card"
          :style="{ '--tint': project.tint, '--accent': project.accent }"
          :initial="{ opacity: 0, y: 40, scale: 0.95 }"
          :while-in-view="{ opacity: 1, y: 0, scale: 1 }"
          :viewport="{ once: true, amount: 0.15 }"
          :transition="{ duration: 0.6, delay: i * 0.1, ease }"
          :while-hover="{ y: -6, scale: 1.01 }"
          :while-tap="{ scale: 0.99 }"
        >
          <img
            class="project-watermark"
            :src="iconUrl(project.watermark)"
            alt=""
            aria-hidden="true"
            loading="lazy"
          />

          <div class="project-info">
            <h3 class="project-title">{{ project.title }}</h3>

            <div class="project-stack">
              <img
                v-for="slug in project.stack"
                :key="slug"
                :src="iconUrl(slug)"
                :alt="slug"
                class="stack-dot"
                loading="lazy"
              />
            </div>

            <div class="project-brand">
              <span class="brand-mark">{{ project.brand.charAt(0) }}</span>
              <span class="brand-name">{{ project.brand }}</span>
            </div>

            <Motion
              as="a"
              href="#projects"
              class="btn-site"
              :while-hover="{ scale: 1.08, filter: 'brightness(1.15)' }"
              :while-tap="{ scale: 0.95 }"
            >
              Visit site
            </Motion>
          </div>

          <div class="project-preview">
            <div class="mock">
              <div class="mock-bar">
                <span></span><span></span><span></span>
              </div>
              <div class="mock-body">
                <div class="mock-line w70"></div>
                <div class="mock-line w45"></div>
                <div class="mock-blocks">
                  <div class="mock-block"></div>
                  <div class="mock-block"></div>
                </div>
              </div>
            </div>
          </div>
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

.projects-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
}

.project-card {
  position: relative;
  overflow: hidden;
  border-radius: var(--radius-lg);
  padding: 22px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 18px;
  align-items: center;
  background: var(--tint);
  border: 1px solid var(--border-card);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.project-card:hover {
  border-color: rgba(59, 130, 246, 0.45);
  box-shadow: var(--shadow);
}

.project-watermark {
  position: absolute;
  right: -18px;
  bottom: -22px;
  width: 130px;
  height: 130px;
  opacity: 0.08;
  pointer-events: none;
}

.project-info {
  position: relative;
  z-index: 1;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.project-title {
  font-size: 1rem;
  font-weight: 700;
  color: #fff;
}

.project-stack {
  display: flex;
  gap: 6px;
}

.stack-dot {
  width: 16px;
  height: 16px;
  object-fit: contain;
}

.project-brand {
  display: flex;
  align-items: center;
  gap: 8px;
}

.brand-mark {
  width: 22px;
  height: 22px;
  display: grid;
  place-items: center;
  border-radius: 50%;
  background: var(--primary-light);
  color: #fff;
  font-size: 0.7rem;
  font-weight: 700;
}

.brand-name {
  color: rgba(255, 255, 255, 0.85);
  font-size: 0.8rem;
  font-weight: 600;
}

.btn-site {
  align-self: flex-start;
  background: var(--accent);
  color: #fff;
  font-size: 0.72rem;
  font-weight: 600;
  padding: 8px 18px;
  border-radius: 9999px;
  display: inline-block;
}

.project-preview {
  position: relative;
  z-index: 1;
  border: 2px solid rgba(124, 58, 237, 0.85);
  border-radius: 8px;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.06);
}

.mock {
  background: rgba(10, 20, 40, 0.7);
}

.mock-bar {
  display: flex;
  gap: 4px;
  padding: 6px 8px;
  background: rgba(255, 255, 255, 0.08);
}

.mock-bar span {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.35);
}

.mock-body {
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.mock-line {
  height: 6px;
  border-radius: 3px;
  background: var(--accent);
  opacity: 0.7;
}

.mock-line.w70 {
  width: 70%;
}

.mock-line.w45 {
  width: 45%;
}

.mock-blocks {
  display: flex;
  gap: 8px;
  margin-top: 4px;
}

.mock-block {
  flex: 1;
  height: 26px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.14);
}

@media (max-width: 900px) {
  .projects-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 460px) {
  .project-card {
    grid-template-columns: 1fr;
  }
}
</style>