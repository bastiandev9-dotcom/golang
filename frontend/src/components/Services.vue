<script setup>
import { ref, onMounted } from 'vue'
import { Motion } from 'motion-v'
import { fetchServices } from '../services/service.js'

const ease = [0.16, 1, 0.3, 1]

const services = ref([])
const loading = ref(true)
const error = ref(null)

const loadServices = async () => {
  loading.value = true
  error.value = null
  try {
    const data = await fetchServices()
    services.value = data
  } catch (err) {
    error.value = err.message || 'Terjadi kesalahan saat memuat data layanan.'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadServices()
})
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

      <!-- Keadaan Memuat (Loading) -->
      <div v-if="loading" class="state-container">
        <div class="spinner"></div>
        <p class="state-text">Memuat layanan...</p>
      </div>

      <!-- Pesan Gagal (Error) -->
      <div v-else-if="error" class="state-container error-state">
        <i class="bi bi-exclamation-triangle-fill error-icon"></i>
        <p class="error-text">{{ error }}</p>
        <button class="retry-btn" @click="loadServices">Coba Lagi</button>
      </div>

      <!-- Grid Layanan -->
      <div v-else-if="services.length > 0" class="services-grid">
        <Motion
          v-for="(service, i) in services"
          :key="service.id || service.name"
          as="article"
          class="service-card"
          :initial="{ opacity: 0, y: 40 }"
          :while-in-view="{ opacity: 1, y: 0 }"
          :viewport="{ once: true, amount: 0.15 }"
          :transition="{ duration: 0.55, delay: i * 0.1, ease }"
          :while-hover="{ y: -6, scale: 1.02 }"
        >
          <span class="service-badge">{{ service.name }}</span>

          <!-- Gambar Layanan jika ada -->
          <div v-if="service.image_url" class="service-image-box">
            <img :src="service.image_url" :alt="service.name" class="service-image" loading="lazy" />
          </div>

          <div class="service-icons">
            <i
              v-if="service.icon"
              :class="['bi', `bi-${service.icon}`, 'service-icon-bi']"
            ></i>
          </div>

          <p class="service-desc">{{ service.description }}</p>
        </Motion>
      </div>

      <!-- Jika Layanan Kosong -->
      <div v-else class="state-container">
        <p class="state-text">Belum ada layanan yang tersedia.</p>
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

.state-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 48px 24px;
  background: linear-gradient(180deg, var(--bg-card-2), var(--bg-card));
  border: 1px dashed var(--border-card);
  border-radius: var(--radius);
  text-align: center;
}

.spinner {
  width: 36px;
  height: 36px;
  border: 3px solid rgba(59, 130, 246, 0.2);
  border-top-color: var(--primary-light);
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.state-text {
  color: var(--text-muted);
  font-size: 0.95rem;
}

.error-state {
  border-color: rgba(239, 68, 68, 0.3);
  background: rgba(239, 68, 68, 0.05);
}

.error-icon {
  font-size: 2rem;
  color: #ef4444;
  margin-bottom: 12px;
}

.error-text {
  color: #fca5a5;
  font-size: 0.95rem;
  margin-bottom: 16px;
}

.retry-btn {
  background: var(--primary);
  color: #ffffff;
  border: none;
  border-radius: 6px;
  padding: 8px 18px;
  font-size: 0.85rem;
  font-weight: 600;
  transition: background 0.2s ease;
}

.retry-btn:hover {
  background: var(--primary-light);
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

.service-image-box {
  width: 100%;
  height: 120px;
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 14px;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.service-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.service-card:hover .service-image {
  transform: scale(1.05);
}

.service-icons {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 16px;
  min-height: 32px;
}

.service-icon-bi {
  font-size: 1.75rem;
  color: var(--primary-light);
  transition: transform 0.2s ease;
}

.service-icon-bi:hover {
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