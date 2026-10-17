<script setup>
import { ref, onMounted } from 'vue'
import {
  fetchServices,
  createService,
  updateService,
  deleteService,
} from '../services/service.js'

const services = ref([])
const loading = ref(true)
const error = ref(null)
const submitting = ref(false)
const modalError = ref(null)

// Form state
const isEdit = ref(false)
const form = ref({
  id: null,
  name: '',
  description: '',
  icon: '',
  image_url: '',
})

// Delete confirmation state
const serviceToDelete = ref(null)
const deleting = ref(false)

const loadServices = async () => {
  loading.value = true
  error.value = null
  try {
    const data = await fetchServices()
    services.value = data
  } catch (err) {
    error.value = err.message || 'Gagal memuat daftar layanan.'
  } finally {
    loading.value = false
  }
}

// Bootstrap Modal Helpers
const getModalInstance = (elementId) => {
  const el = document.getElementById(elementId)
  if (!el || !window.bootstrap) return null
  return window.bootstrap.Modal.getOrCreateInstance(el)
}

const openAddModal = () => {
  isEdit.value = false
  modalError.value = null
  form.value = {
    id: null,
    name: '',
    description: '',
    icon: '',
    image_url: '',
  }
  const modal = getModalInstance('serviceModal')
  if (modal) modal.show()
}

const openEditModal = (service) => {
  isEdit.value = true
  modalError.value = null
  form.value = {
    id: service.id,
    name: service.name || '',
    description: service.description || '',
    icon: service.icon || '',
    image_url: service.image_url || '',
  }
  const modal = getModalInstance('serviceModal')
  if (modal) modal.show()
}

const handleSubmit = async () => {
  modalError.value = null
  if (!form.value.name.trim()) {
    modalError.value = "Nama layanan wajib diisi."
    return
  }

  submitting.value = true
  try {
    const payload = {
      name: form.value.name.trim(),
      description: form.value.description.trim(),
      icon: form.value.icon.trim(),
      image_url: form.value.image_url ? form.value.image_url.trim() : null,
    }

    if (isEdit.value) {
      await updateService(form.value.id, payload)
    } else {
      await createService(payload)
    }

    const modal = getModalInstance('serviceModal')
    if (modal) modal.hide()
    await loadServices()
  } catch (err) {
    modalError.value = err.message || 'Terjadi kesalahan saat menyimpan data.'
  } finally {
    submitting.value = false
  }
}

const confirmDelete = (service) => {
  serviceToDelete.value = service
  const modal = getModalInstance('deleteConfirmModal')
  if (modal) modal.show()
}

const handleDelete = async () => {
  if (!serviceToDelete.value) return
  deleting.value = true
  try {
    await deleteService(serviceToDelete.value.id)
    const modal = getModalInstance('deleteConfirmModal')
    if (modal) modal.hide()
    serviceToDelete.value = null
    await loadServices()
  } catch (err) {
    alert(err.message || 'Gagal menghapus layanan.')
  } finally {
    deleting.value = false
  }
}

const navigateToHome = () => {
  window.location.pathname = '/'
}

onMounted(() => {
  loadServices()
})
</script>

<template>
  <div class="dashboard-page">
    <!-- Dashboard Header -->
    <header class="dashboard-navbar border-bottom py-3 px-4">
      <div class="container-fluid d-flex justify-content-between align-items-center">
        <div class="d-flex align-items-center gap-3">
          <span class="brand-title fw-bold fs-4">Dashboard Layanan</span>
          <span class="badge bg-primary-subtle text-primary border border-primary-subtle">Admin</span>
        </div>
        <div>
          <button class="btn btn-outline-light btn-sm" @click="navigateToHome">
            <i class="bi bi-box-arrow-left me-1"></i> Kembali ke Landing Page
          </button>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="container py-4">
      <div class="d-flex justify-content-between align-items-center mb-4">
        <div>
          <h1 class="h3 fw-bold mb-1">Daftar Layanan</h1>
          <p class="text-secondary small mb-0">Kelola dan perbarui data layanan yang tampil di landing page.</p>
        </div>
        <button class="btn btn-primary d-flex align-items-center gap-2" @click="openAddModal">
          <i class="bi bi-plus-lg"></i> Tambah Layanan
        </button>
      </div>

      <!-- Error State -->
      <div v-if="error" class="alert alert-danger d-flex align-items-center justify-content-between" role="alert">
        <div>
          <i class="bi bi-exclamation-triangle-fill me-2"></i>
          <span>{{ error }}</span>
        </div>
        <button class="btn btn-outline-danger btn-sm" @click="loadServices">Coba Lagi</button>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border text-primary" role="status">
          <span class="visually-hidden">Memuat data...</span>
        </div>
        <p class="text-secondary mt-2 small">Memuat daftar layanan...</p>
      </div>

      <!-- Services Table -->
      <div v-else class="card shadow-sm border-0 dashboard-card">
        <div class="table-responsive">
          <table class="table table-hover align-middle mb-0">
            <thead class="table-dark">
              <tr>
                <th style="width: 70px;" class="text-center">ID</th>
                <th style="width: 80px;" class="text-center">Ikon</th>
                <th style="width: 90px;" class="text-center">Gambar</th>
                <th>Nama Layanan</th>
                <th>Deskripsi</th>
                <th style="width: 150px;">Dibuat Pada</th>
                <th style="width: 140px;" class="text-center">Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="service in services" :key="service.id">
                <td class="text-center text-secondary fw-semibold">{{ service.id }}</td>
                <td class="text-center">
                  <div class="icon-preview-box mx-auto">
                    <i v-if="service.icon" :class="['bi', `bi-${service.icon}`, 'fs-5 text-primary']"></i>
                    <span v-else class="text-secondary">-</span>
                  </div>
                </td>
                <td class="text-center">
                  <img
                    v-if="service.image_url"
                    :src="service.image_url"
                    :alt="service.name"
                    class="rounded object-fit-cover border border-secondary"
                    style="width: 48px; height: 36px;"
                  />
                  <span v-else class="text-secondary small">-</span>
                </td>
                <td class="fw-semibold text-light">{{ service.name }}</td>
                <td class="text-secondary small text-truncate" style="max-width: 320px;" :title="service.description">
                  {{ service.description || '-' }}
                </td>
                <td class="text-secondary small">
                  {{ service.created_at ? new Date(service.created_at).toLocaleDateString('id-ID') : '-' }}
                </td>
                <td class="text-center">
                  <div class="btn-group btn-group-sm">
                    <button class="btn btn-outline-primary" title="Ubah Layanan" @click="openEditModal(service)">
                      <i class="bi bi-pencil-square"></i>
                    </button>
                    <button class="btn btn-outline-danger" title="Hapus Layanan" @click="confirmDelete(service)">
                      <i class="bi bi-trash3"></i>
                    </button>
                  </div>
                </td>
              </tr>
              <tr v-if="services.length === 0">
                <td colspan="6" class="text-center py-5 text-secondary">
                  <i class="bi bi-inbox fs-2 d-block mb-2 text-secondary opacity-50"></i>
                  Belum ada data layanan. Klik "Tambah Layanan" untuk menambahkan.
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </main>

    <!-- Modal Form (Tambah / Ubah) -->
    <div
      class="modal fade"
      id="serviceModal"
      tabindex="-1"
      aria-labelledby="serviceModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content text-light border-secondary">
          <div class="modal-header border-secondary">
            <h5 class="modal-title" id="serviceModalLabel">
              {{ isEdit ? 'Ubah Layanan' : 'Tambah Layanan Baru' }}
            </h5>
            <button type="button" class="btn-close btn-close-white" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <form @submit.prevent="handleSubmit">
            <div class="modal-body">
              <div v-if="modalError" class="alert alert-danger py-2 px-3 small">
                <i class="bi bi-exclamation-circle-fill me-1"></i> {{ modalError }}
              </div>

              <!-- Nama Layanan -->
              <div class="mb-3">
                <label for="serviceName" class="form-label small fw-semibold">Nama Layanan <span class="text-danger">*</span></label>
                <input
                  id="serviceName"
                  v-model="form.name"
                  type="text"
                  class="form-control form-control-dark"
                  placeholder="Contoh: Web Development"
                  required
                />
              </div>

              <!-- Ikon Bootstrap -->
              <div class="mb-3">
                <label for="serviceIcon" class="form-label small fw-semibold">
                  Nama Ikon (Bootstrap Icons)
                </label>
                <div class="input-group">
                  <span class="input-group-text bg-dark border-secondary text-secondary">bi-</span>
                  <input
                    id="serviceIcon"
                    v-model="form.icon"
                    type="text"
                    class="form-control form-control-dark"
                    placeholder="Contoh: code-slash, palette, layers"
                  />
                  <span v-if="form.icon" class="input-group-text bg-dark border-secondary text-primary">
                    <i :class="['bi', `bi-${form.icon}`]"></i>
                  </span>
                </div>
                <div class="form-text text-secondary small">
                  Nama ikon dari <a href="https://icons.getbootstrap.com/" target="_blank" class="text-primary">Bootstrap Icons</a> (tanpa awalan <code>bi-</code>).
                </div>
              </div>

              <!-- Image URL -->
              <div class="mb-3">
                <label for="serviceImage" class="form-label small fw-semibold">URL Gambar (Opsional)</label>
                <input
                  id="serviceImage"
                  v-model="form.image_url"
                  type="url"
                  class="form-control form-control-dark"
                  placeholder="https://images.unsplash.com/..."
                />
              </div>

              <!-- Deskripsi -->
              <div class="mb-3">
                <label for="serviceDesc" class="form-label small fw-semibold">Deskripsi</label>
                <textarea
                  id="serviceDesc"
                  v-model="form.description"
                  class="form-control form-control-dark"
                  rows="3"
                  placeholder="Deskripsi singkat mengenai layanan ini..."
                ></textarea>
              </div>
            </div>
            <div class="modal-footer border-secondary">
              <button type="button" class="btn btn-outline-secondary" data-bs-dismiss="modal">Batal</button>
              <button type="submit" class="btn btn-primary" :disabled="submitting">
                <span v-if="submitting" class="spinner-border spinner-border-sm me-1" role="status"></span>
                {{ isEdit ? 'Simpan Perubahan' : 'Tambah Layanan' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Modal Konfirmasi Hapus -->
    <div
      class="modal fade"
      id="deleteConfirmModal"
      tabindex="-1"
      aria-labelledby="deleteConfirmModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered modal-sm">
        <div class="modal-content text-light border-danger">
          <div class="modal-header border-secondary">
            <h5 class="modal-title text-danger fs-6" id="deleteConfirmModalLabel">
              <i class="bi bi-exclamation-triangle-fill me-1"></i> Konfirmasi Hapus
            </h5>
            <button type="button" class="btn-close btn-close-white" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body text-center py-4">
            <p class="mb-2">Apakah Anda yakin ingin menghapus layanan ini?</p>
            <p class="fw-bold text-light mb-0" v-if="serviceToDelete">
              "{{ serviceToDelete.name }}"
            </p>
            <small class="text-secondary d-block mt-2">Tindakan ini tidak dapat dibatalkan.</small>
          </div>
          <div class="modal-footer border-secondary justify-content-center">
            <button type="button" class="btn btn-outline-secondary btn-sm" data-bs-dismiss="modal">Batal</button>
            <button type="button" class="btn btn-danger btn-sm" :disabled="deleting" @click="handleDelete">
              <span v-if="deleting" class="spinner-border spinner-border-sm me-1" role="status"></span>
              Ya, Hapus
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-page {
  min-height: 100vh;
  background-color: #0b1329;
  color: #f1f5f9;
}

.dashboard-navbar {
  background: rgba(15, 23, 42, 0.85);
  backdrop-filter: blur(12px);
  border-color: rgba(255, 255, 255, 0.08) !important;
}

.brand-title {
  background: linear-gradient(90deg, #60a5fa, #c084fc);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.dashboard-card {
  background: #111e38;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.35);
}

.table {
  color: #e2e8f0;
  --bs-table-bg: transparent;
  --bs-table-striped-bg: rgba(255, 255, 255, 0.02);
}

.table-dark {
  --bs-table-bg: #162447;
  border-bottom: 2px solid rgba(255, 255, 255, 0.1);
}

.table > :not(caption) > * > * {
  padding: 14px 16px;
  border-bottom-color: rgba(255, 255, 255, 0.06);
}

.icon-preview-box {
  width: 38px;
  height: 38px;
  border-radius: 8px;
  background: rgba(37, 99, 235, 0.12);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background-color: #132040;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.6);
}

.form-control-dark {
  background-color: #0d172e;
  border-color: rgba(255, 255, 255, 0.15);
  color: #f1f5f9;
}

.form-control-dark:focus {
  background-color: #0f1d3a;
  border-color: #3b82f6;
  color: #ffffff;
  box-shadow: 0 0 0 0.25rem rgba(59, 130, 246, 0.25);
}
</style>
