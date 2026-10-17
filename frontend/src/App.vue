<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { motion, useScroll, useSpring } from 'motion-v'
import Navbar from './components/Navbar.vue'
import Hero from './components/Hero.vue'
import Services from './components/Services.vue'
import Projects from './components/Projects.vue'
import Clients from './components/Clients.vue'
import Contact from './components/Contact.vue'
import Dashboard from './components/Dashboard.vue'

const currentPath = ref(window.location.pathname)

const isDashboard = computed(() => {
  return currentPath.value === '/dashboard' || currentPath.value.startsWith('/dashboard/')
})

const onPopState = () => {
  currentPath.value = window.location.pathname
}

onMounted(() => {
  window.addEventListener('popstate', onPopState)
})

onUnmounted(() => {
  window.removeEventListener('popstate', onPopState)
})

const { scrollYProgress } = useScroll()
const progress = useSpring(scrollYProgress, {
  stiffness: 120,
  damping: 28,
  mass: 0.4,
})
</script>

<template>
  <!-- Route /dashboard -->
  <Dashboard v-if="isDashboard" />

  <!-- Landing page (/) persis seperti aslinya -->
  <template v-else>
    <motion.div
      class="scroll-progress"
      :style="{ scaleX: progress }"
      aria-hidden="true"
    />
    <div class="layout">
      <Navbar />
      <main>
        <Hero />
        <Services />
        <Projects />
        <Clients />
        <Contact />
      </main>
    </div>
  </template>
</template>

<style scoped>
.scroll-progress {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  transform-origin: left;
  background: linear-gradient(90deg, var(--primary-light), var(--accent));
  z-index: 2000;
}

.layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

main {
  flex: 1;
}
</style>