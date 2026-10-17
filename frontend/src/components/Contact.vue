<script setup>
import { ref } from 'vue'
import { Motion } from 'motion-v'
import { resolveIcon } from '../assets/icons.js'
import contactArt from '../assets/contact-art.svg'

const ease = [0.16, 1, 0.3, 1]

const form = ref({ name: '', email: '', message: '' })
const sent = ref(false)

const socials = [
  { name: 'Facebook', slug: 'facebook' },
  { name: 'Discord', slug: 'discord' },
  { name: 'WhatsApp', slug: 'whatsapp' },
  { name: 'LinkedIn', slug: 'linkedin' },
  { name: 'Upwork', slug: 'upwork' },
  { name: 'Twitter', slug: 'twitter' },
  { name: 'GitHub', slug: 'github' },
]

const logoUrl = (slug) => resolveIcon(slug, '94a3b8')

const handleSubmit = () => {
  sent.value = true
  form.value = { name: '', email: '', message: '' }
  setTimeout(() => (sent.value = false), 3000)
}
</script>

<template>
  <section id="contact" class="section contact">
    <div class="container">
      <Motion
        as="h2"
        class="contact-title"
        :initial="{ opacity: 0, y: 20 }"
        :while-in-view="{ opacity: 1, y: 0 }"
        :viewport="{ once: true, amount: 0.4 }"
        :transition="{ duration: 0.6, ease }"
      >
        Contact
      </Motion>

      <div class="contact-grid">
        <Motion
          as="div"
          class="contact-illustration"
          aria-hidden="true"
          :initial="{ opacity: 0, x: -40 }"
          :while-in-view="{ opacity: 1, x: 0 }"
          :viewport="{ once: true, amount: 0.2 }"
          :transition="{ duration: 0.7, ease }"
        >
          <img :src="contactArt" class="illustration-art" alt="Contact Illustration" />
        </Motion>

        <Motion
          as="div"
          class="contact-form-card"
          :initial="{ opacity: 0, x: 40 }"
          :while-in-view="{ opacity: 1, x: 0 }"
          :viewport="{ once: true, amount: 0.2 }"
          :transition="{ duration: 0.7, ease }"
        >
          <h3 class="form-heading">
            Contact with me for any kind of<br />
            <span class="highlight">Queries</span>
          </h3>

          <form @submit.prevent="handleSubmit">
            <label class="field">
              <span class="field-label">Name</span>
              <input v-model="form.name" type="text" required placeholder="Your name" />
            </label>

            <label class="field">
              <span class="field-label">Email</span>
              <input v-model="form.email" type="email" required placeholder="you@email.com" />
            </label>

            <label class="field">
              <span class="field-label">Message</span>
              <textarea v-model="form.message" required placeholder="Tell me about your project"></textarea>
            </label>

            <Motion
              as="button"
              type="submit"
              class="btn-submit"
              :while-hover="{ scale: 1.02, filter: 'brightness(1.1)' }"
              :while-tap="{ scale: 0.98 }"
            >
              Submit
            </Motion>

            <p v-if="sent" class="form-note">Thanks! Your message has been sent.</p>
          </form>
        </Motion>
      </div>

      <div class="socials">
        <Motion
          v-for="(social, index) in socials"
          :key="social.name"
          as="a"
          href="#contact"
          class="social-link"
          :aria-label="social.name"
          :initial="{ opacity: 0, y: 15 }"
          :while-in-view="{ opacity: 1, y: 0 }"
          :viewport="{ once: true }"
          :transition="{ duration: 0.4, delay: index * 0.05 }"
          :while-hover="{ scale: 1.2, backgroundColor: 'rgba(59, 130, 246, 0.25)', y: -3 }"
          :while-tap="{ scale: 0.92 }"
        >
          <img :src="logoUrl(social.slug)" :alt="social.name" loading="lazy" />
        </Motion>
      </div>

      <div class="copyright">
        <span class="line"></span>
        <span>Copyright @2026 all right reserved</span>
        <span class="line"></span>
      </div>
    </div>
  </section>
</template>

<style scoped>
.contact-title {
  text-align: center;
  font-size: 2rem;
  font-weight: 800;
  margin-bottom: 40px;
}

.contact-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40px;
  align-items: center;
  margin-bottom: 56px;
}

.contact-illustration {
  position: relative;
  display: flex;
  justify-content: center;
}

.illustration-art {
  width: 100%;
  max-width: 440px;
  height: auto;
}

.contact-form-card {
  background: var(--bg-2);
  border: 2px solid var(--primary);
  border-radius: 6px;
  padding: 26px;
  box-shadow: var(--shadow);
}

.form-heading {
  font-size: 1.1rem;
  font-weight: 700;
  line-height: 1.4;
  margin-bottom: 22px;
}

.field {
  display: block;
  margin-bottom: 18px;
}

.field-label {
  display: block;
  font-size: 0.7rem;
  color: var(--text-muted);
  margin-bottom: 2px;
}

.field input,
.field textarea {
  width: 100%;
  background: transparent;
  border: none;
  border-bottom: 1px solid var(--border-soft);
  color: var(--text);
  font-size: 0.8rem;
  padding: 6px 0;
  outline: none;
  transition: border-color 0.2s ease;
}

.field input::placeholder,
.field textarea::placeholder {
  color: var(--text-dim);
}

.field input:focus,
.field textarea:focus {
  border-bottom-color: var(--primary-light);
}

.field textarea {
  resize: vertical;
  min-height: 76px;
  font-family: inherit;
}

.btn-submit {
  width: 100%;
  height: 44px;
  background: var(--primary);
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 0.9rem;
  font-weight: 600;
  margin-top: 6px;
  cursor: pointer;
}

.form-note {
  margin-top: 12px;
  color: #4ade80;
  font-size: 0.78rem;
  text-align: center;
}

.socials {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 14px;
  margin-bottom: 26px;
}

.social-link {
  width: 36px;
  height: 36px;
  display: grid;
  place-items: center;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.05);
}

.social-link img {
  width: 16px;
  height: 16px;
  object-fit: contain;
}

.copyright {
  display: flex;
  align-items: center;
  gap: 14px;
  justify-content: center;
  color: var(--text-dim);
  font-size: 0.72rem;
}

.copyright .line {
  flex: 1;
  height: 1px;
  background: var(--border-soft);
  max-width: 180px;
}

@media (max-width: 768px) {
  .contact-grid {
    grid-template-columns: 1fr;
    gap: 28px;
  }
}
</style>