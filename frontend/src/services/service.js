const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export async function fetchServices() {
  const response = await fetch(`${API_BASE_URL}/api/services`)
  if (!response.ok) {
    throw new Error(`Gagal memuat data layanan (Status: ${response.status})`)
  }
  return await response.json()
}

export async function createService(payload) {
  let options = {}
  if (payload instanceof FormData) {
    options = {
      method: 'POST',
      body: payload,
    }
  } else {
    options = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    }
  }

  const response = await fetch(`${API_BASE_URL}/api/services`, options)
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.error || `Gagal menambah layanan (Status: ${response.status})`)
  }
  return await response.json()
}

export async function updateService(id, payload) {
  let options = {}
  if (payload instanceof FormData) {
    options = {
      method: 'PUT',
      body: payload,
    }
  } else {
    options = {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    }
  }

  const response = await fetch(`${API_BASE_URL}/api/services/${id}`, options)
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.error || `Gagal mengubah layanan (Status: ${response.status})`)
  }
  return await response.json()
}

export async function deleteService(id) {
  const response = await fetch(`${API_BASE_URL}/api/services/${id}`, {
    method: 'DELETE',
  })
  if (!response.ok && response.status !== 204) {
    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.error || `Gagal menghapus layanan (Status: ${response.status})`)
  }
  return true
}
