<template>
  <div class="min-h-screen bg-neutral-950">
    <!-- Hero -->
    <div class="relative overflow-hidden">
      <div class="absolute inset-0 bg-gradient-to-b from-neutral-950 via-neutral-950 to-neutral-900"></div>
      <div class="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,.015)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,.015)_1px,transparent_1px)] bg-[size:64px_64px]"></div>
      
      <!-- Glowing orbs -->
      <div class="absolute top-0 left-1/4 w-[400px] h-[300px] bg-primary-500/10 rounded-full blur-[100px]"></div>
      <div class="absolute top-0 right-1/4 w-[300px] h-[200px] bg-emerald-500/10 rounded-full blur-[80px]"></div>
      
      <div class="relative max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 lg:pt-20 pb-12 text-center">
        <div class="inline-flex items-center gap-2 px-3 py-1.5 rounded-full bg-white/5 border border-white/10 backdrop-blur-sm mb-6">
          <span class="relative flex h-2 w-2">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
            <span class="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
          </span>
          <span class="text-xs text-neutral-300">We're here to help</span>
        </div>
        
        <h1 class="font-display text-3xl sm:text-4xl lg:text-5xl font-bold text-white mb-4">
          Get in <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-400 to-emerald-400">Touch</span>
        </h1>
        <p class="text-neutral-400 max-w-xl mx-auto">
          Have questions about our platform? Want to partner with us? We'd love to hear from you.
        </p>
      </div>
    </div>

    <!-- Content -->
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 pb-24">
      <div class="grid lg:grid-cols-5 gap-8 lg:gap-12">
        
        <!-- Contact Info - Order changes on mobile: form first, then info -->
        <div class="lg:col-span-2 space-y-6 order-2 lg:order-1">
          <!-- Quick Contact Cards -->
          <div class="bg-neutral-900/50 border border-neutral-800 rounded-2xl p-6">
            <h3 class="font-semibold text-white mb-4 flex items-center gap-2">
              <svg class="w-5 h-5 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
              Email Us
            </h3>
            <div class="space-y-3">
              <a href="mailto:hello@unicornsport.africa" class="block text-neutral-400 hover:text-white transition-colors">
                <span class="text-xs text-neutral-500">General Inquiries</span><br>
                hello@unicornsport.africa
              </a>
              <a href="mailto:scouts@unicornsport.africa" class="block text-neutral-400 hover:text-white transition-colors">
                <span class="text-xs text-neutral-500">For Scouts & Clubs</span><br>
                scouts@unicornsport.africa
              </a>
              <a href="mailto:academies@unicornsport.africa" class="block text-neutral-400 hover:text-white transition-colors">
                <span class="text-xs text-neutral-500">Academy Partnerships</span><br>
                academies@unicornsport.africa
              </a>
            </div>
          </div>

          <div class="bg-neutral-900/50 border border-neutral-800 rounded-2xl p-6">
            <h3 class="font-semibold text-white mb-4 flex items-center gap-2">
              <svg class="w-5 h-5 text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              Location
            </h3>
            <p class="text-neutral-400">
              Unicorn Sport Academy<br>
              Lagos, Nigeria<br>
              <span class="text-xs text-neutral-500">West Africa</span>
            </p>
          </div>

          <div class="bg-neutral-900/50 border border-neutral-800 rounded-2xl p-6">
            <h3 class="font-semibold text-white mb-4 flex items-center gap-2">
              <svg class="w-5 h-5 text-secondary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              Response Time
            </h3>
            <p class="text-neutral-400 text-sm">
              We typically respond within <span class="text-white font-medium">24 hours</span> on business days.
            </p>
            <div class="mt-3 flex items-center gap-2 text-xs text-neutral-500">
              <span class="w-2 h-2 rounded-full bg-emerald-500"></span>
              Currently online (WAT timezone)
            </div>
          </div>

          <!-- Social Links -->
          <div class="bg-neutral-900/50 border border-neutral-800 rounded-2xl p-6">
            <h3 class="font-semibold text-white mb-4">Follow Us</h3>
            <div class="flex gap-3">
              <a href="#" class="w-10 h-10 rounded-xl bg-white/5 border border-white/10 flex items-center justify-center text-neutral-400 hover:bg-primary-500 hover:border-primary-500 hover:text-white transition-all">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/>
                </svg>
              </a>
              <a href="#" class="w-10 h-10 rounded-xl bg-white/5 border border-white/10 flex items-center justify-center text-neutral-400 hover:bg-gradient-to-br hover:from-purple-500 hover:to-pink-500 hover:border-transparent hover:text-white transition-all">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/>
                </svg>
              </a>
              <a href="#" class="w-10 h-10 rounded-xl bg-white/5 border border-white/10 flex items-center justify-center text-neutral-400 hover:bg-red-600 hover:border-red-600 hover:text-white transition-all">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/>
                </svg>
              </a>
              <a href="#" class="w-10 h-10 rounded-xl bg-white/5 border border-white/10 flex items-center justify-center text-neutral-400 hover:bg-blue-600 hover:border-blue-600 hover:text-white transition-all">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
                </svg>
              </a>
            </div>
          </div>
        </div>

        <!-- Contact Form -->
        <div class="lg:col-span-3 order-1 lg:order-2">
          <div class="bg-neutral-900/50 border border-neutral-800 rounded-2xl p-6 lg:p-8">
            <h2 class="text-xl font-semibold text-white mb-6">Send us a message</h2>
            
            <form @submit.prevent="submitForm" class="space-y-5">
              <div class="grid sm:grid-cols-2 gap-5">
                <div>
                  <label for="name" class="block text-sm font-medium text-neutral-400 mb-2">Your Name</label>
                  <input 
                    id="name"
                    v-model="form.name"
                    type="text" 
                    required
                    placeholder="John Doe"
                    class="w-full px-4 py-3 bg-neutral-800/50 border border-neutral-700 rounded-xl text-white placeholder-neutral-500 focus:outline-none focus:border-primary-500 transition-all"
                  />
                </div>
                <div>
                  <label for="email" class="block text-sm font-medium text-neutral-400 mb-2">Email Address</label>
                  <input 
                    id="email"
                    v-model="form.email"
                    type="email" 
                    required
                    placeholder="john@example.com"
                    class="w-full px-4 py-3 bg-neutral-800/50 border border-neutral-700 rounded-xl text-white placeholder-neutral-500 focus:outline-none focus:border-primary-500 transition-all"
                  />
                </div>
              </div>

              <div>
                <label for="subject" class="block text-sm font-medium text-neutral-400 mb-2">Subject</label>
                <select 
                  id="subject"
                  v-model="form.subject"
                  required
                  class="w-full px-4 py-3 bg-neutral-800/50 border border-neutral-700 rounded-xl text-white focus:outline-none focus:border-primary-500 transition-all appearance-none"
                >
                  <option value="" disabled>Select a topic</option>
                  <option value="general">General Inquiry</option>
                  <option value="scout">Scout/Club Account</option>
                  <option value="academy">Academy Partnership</option>
                  <option value="player">Player Inquiry</option>
                  <option value="billing">Billing & Subscriptions</option>
                  <option value="technical">Technical Support</option>
                  <option value="press">Press & Media</option>
                  <option value="other">Other</option>
                </select>
              </div>

              <div>
                <label for="organization" class="block text-sm font-medium text-neutral-400 mb-2">Organization (optional)</label>
                <input 
                  id="organization"
                  v-model="form.organization"
                  type="text" 
                  placeholder="Your club or company name"
                  class="w-full px-4 py-3 bg-neutral-800/50 border border-neutral-700 rounded-xl text-white placeholder-neutral-500 focus:outline-none focus:border-primary-500 transition-all"
                />
              </div>

              <div>
                <label for="message" class="block text-sm font-medium text-neutral-400 mb-2">Message</label>
                <textarea 
                  id="message"
                  v-model="form.message"
                  required
                  rows="5"
                  placeholder="Tell us how we can help..."
                  class="w-full px-4 py-3 bg-neutral-800/50 border border-neutral-700 rounded-xl text-white placeholder-neutral-500 focus:outline-none focus:border-primary-500 transition-all resize-none"
                ></textarea>
              </div>

              <button 
                type="submit"
                :disabled="isSubmitting"
                class="w-full py-3 px-6 bg-gradient-to-r from-primary-500 to-emerald-500 text-white font-semibold rounded-xl hover:shadow-lg hover:shadow-primary-500/25 transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
              >
                <svg v-if="isSubmitting" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                </svg>
                <span>{{ isSubmitting ? 'Sending...' : 'Send Message' }}</span>
                <svg v-if="!isSubmitting" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                </svg>
              </button>

              <!-- Success Message -->
              <div v-if="submitted" class="p-4 bg-emerald-500/10 border border-emerald-500/30 rounded-xl">
                <div class="flex items-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-emerald-500/20 flex items-center justify-center flex-shrink-0">
                    <svg class="w-4 h-4 text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                  </div>
                  <div>
                    <p class="text-emerald-400 font-medium">Message sent successfully!</p>
                    <p class="text-neutral-400 text-sm">We'll get back to you within 24 hours.</p>
                  </div>
                </div>
              </div>
            </form>
          </div>

          <!-- FAQ Teaser -->
          <div class="mt-6 p-6 bg-neutral-900/30 border border-neutral-800/50 rounded-2xl">
            <div class="flex items-start gap-4">
              <div class="w-10 h-10 rounded-xl bg-primary-500/10 flex items-center justify-center flex-shrink-0">
                <svg class="w-5 h-5 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <h3 class="font-medium text-white mb-1">Looking for quick answers?</h3>
                <p class="text-neutral-400 text-sm">Check out our <NuxtLink to="/pricing#faq" class="text-primary-400 hover:underline">FAQ section</NuxtLink> on the pricing page for common questions about subscriptions, features, and more.</p>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
useHead({
  title: 'Contact Us | Unicorn Sport',
  meta: [
    { name: 'description', content: 'Get in touch with Unicorn Sport. Contact us for scout accounts, academy partnerships, or general inquiries.' }
  ]
})

const form = ref({
  name: '',
  email: '',
  subject: '',
  organization: '',
  message: ''
})

const isSubmitting = ref(false)
const submitted = ref(false)

async function submitForm() {
  isSubmitting.value = true
  
  // Simulate API call
  await new Promise(resolve => setTimeout(resolve, 1500))
  
  console.log('Form submitted:', form.value)
  
  // Reset form
  form.value = {
    name: '',
    email: '',
    subject: '',
    organization: '',
    message: ''
  }
  
  isSubmitting.value = false
  submitted.value = true
  
  // Hide success message after 5 seconds
  setTimeout(() => {
    submitted.value = false
  }, 5000)
}
</script>
