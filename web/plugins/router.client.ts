// Plugin to handle client-side navigation
import type { RouteLocationNormalized } from 'vue-router'

export default defineNuxtPlugin(() => {
  const router = useRouter()

  // Force scroll to top on navigation
  router.afterEach((to: RouteLocationNormalized, from: RouteLocationNormalized) => {
    if (to.path !== from.path) {
      nextTick(() => {
        window.scrollTo({ top: 0, behavior: 'instant' })
      })
    }
  })
})
