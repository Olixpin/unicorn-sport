<template>
  <div class="min-h-screen bg-neutral-950">
    <!-- Hero -->
    <div class="relative overflow-hidden">
      <div class="absolute inset-0 bg-gradient-to-b from-neutral-950 via-neutral-950 to-neutral-900"></div>
      <div class="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,.015)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,.015)_1px,transparent_1px)] bg-[size:64px_64px]"></div>
      
      <div class="relative max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 lg:pt-20 pb-12 text-center">
        <div class="inline-flex items-center gap-2 px-3 py-1.5 rounded-full bg-white/5 border border-white/10 backdrop-blur-sm mb-6">
          <svg class="w-4 h-4 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <span class="text-xs text-neutral-300">Legal agreement</span>
        </div>
        
        <h1 class="font-display text-3xl sm:text-4xl lg:text-5xl font-bold text-white mb-4">
          Terms of Service
        </h1>
        <p class="text-neutral-400 max-w-xl mx-auto">
          Last updated: {{ lastUpdated }}
        </p>
      </div>
    </div>

    <!-- Content -->
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 pb-24">
      <div class="bg-neutral-900/50 border border-neutral-800 rounded-2xl p-8 lg:p-12">
        <div class="prose prose-invert prose-neutral max-w-none">
          
          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">1</span>
              Agreement to Terms
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              By accessing or using Unicorn Sport ("Platform"), you agree to be bound by these Terms of Service ("Terms"). If you disagree with any part of these terms, you may not access the Platform.
            </p>
            <p class="text-neutral-400 leading-relaxed mt-4">
              Unicorn Sport is a professional talent discovery platform operated by a scouting academy based in Nigeria. We connect verified African youth football talent with scouts and clubs worldwide.
            </p>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">2</span>
              Platform Description
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              Unicorn Sport provides:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li>A platform for scouts and clubs to discover verified African football talent</li>
              <li>Player profiles with verified age and identity information</li>
              <li>Professional video content including highlights and full match recordings</li>
              <li>Secure communication channels between scouts and players/academies</li>
              <li>Subscription-based access to premium content and features</li>
            </ul>
            <p class="text-neutral-400 leading-relaxed mt-4">
              <strong class="text-white">Important:</strong> We are NOT a user-generated content platform. All player profiles and videos are created by our academy staff and verified before publication.
            </p>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">3</span>
              User Accounts
            </h2>
            
            <h3 class="text-lg font-medium text-white mt-6 mb-3">Account Creation</h3>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li>You must provide accurate and complete information when creating an account</li>
              <li>You are responsible for maintaining the security of your account credentials</li>
              <li>You must notify us immediately of any unauthorized access to your account</li>
              <li>You must be at least 18 years old to create a scout or club account</li>
            </ul>

            <h3 class="text-lg font-medium text-white mt-6 mb-3">Account Types</h3>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li><strong class="text-white">Scout/Club Accounts:</strong> For professional scouts and football clubs seeking talent</li>
              <li><strong class="text-white">Academy Accounts:</strong> For partner academies managing player profiles</li>
              <li><strong class="text-white">Admin Accounts:</strong> For Unicorn Sport staff only</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">4</span>
              Subscriptions & Payments
            </h2>
            
            <h3 class="text-lg font-medium text-white mt-6 mb-3">Subscription Tiers</h3>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li><strong class="text-white">Explorer (Free):</strong> Browse profiles and watch highlight clips</li>
              <li><strong class="text-white">Scout ($79-99/mo):</strong> Full match access, shortlists, advanced filters</li>
              <li><strong class="text-white">Pro ($239-299/mo):</strong> Contact players, export reports, priority support</li>
              <li><strong class="text-white">Club ($799-999/mo):</strong> Unlimited contacts, API access, dedicated account manager</li>
            </ul>

            <h3 class="text-lg font-medium text-white mt-6 mb-3">Billing</h3>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li>Payments are processed securely through Stripe</li>
              <li>Subscriptions renew automatically unless cancelled</li>
              <li>Annual subscriptions receive a 20% discount</li>
              <li>Refunds are available within 14 days of initial purchase</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">5</span>
              Acceptable Use
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              You agree NOT to:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li>Use the Platform for any illegal purpose</li>
              <li>Download, copy, or redistribute player videos without authorization</li>
              <li>Contact players or their families outside of the Platform's official channels</li>
              <li>Misrepresent your identity, affiliation, or intentions</li>
              <li>Attempt to circumvent subscription paywalls or access controls</li>
              <li>Scrape, crawl, or harvest data from the Platform</li>
              <li>Interfere with the Platform's security or integrity</li>
              <li>Use the Platform to exploit, harm, or manipulate minors</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">6</span>
              Player Protection
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              Given that our platform features youth athletes, we enforce strict protection measures:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li>All contact with players must go through official Platform channels</li>
              <li>Contact requests are logged and moderated by our team</li>
              <li>Any suspicious or inappropriate behavior will result in immediate account termination</li>
              <li>We cooperate with law enforcement on any safeguarding concerns</li>
              <li>Scouts must verify their professional credentials before contacting players</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">7</span>
              Intellectual Property
            </h2>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li>All content on the Platform (videos, profiles, design) is owned by Unicorn Sport</li>
              <li>You are granted a limited license to view content for personal, professional evaluation purposes</li>
              <li>You may not reproduce, distribute, or create derivative works without written permission</li>
              <li>Our trademarks, logos, and branding may not be used without authorization</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">8</span>
              Disclaimers
            </h2>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li>The Platform is provided "as is" without warranties of any kind</li>
              <li>We do not guarantee player performance or suitability for any club</li>
              <li>We are not responsible for outcomes of scout-player relationships initiated through the Platform</li>
              <li>We reserve the right to modify or discontinue features at any time</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">9</span>
              Termination
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              We may terminate or suspend your account immediately, without prior notice, for:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li>Violation of these Terms</li>
              <li>Fraudulent or illegal activity</li>
              <li>Behavior that endangers or exploits minors</li>
              <li>Non-payment of subscription fees</li>
              <li>Any reason at our sole discretion</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">10</span>
              Governing Law
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              These Terms shall be governed by and construed in accordance with the laws of the Federal Republic of Nigeria, without regard to its conflict of law provisions. Any disputes arising under these Terms shall be subject to the exclusive jurisdiction of the courts in Lagos, Nigeria.
            </p>
          </section>

          <section>
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-emerald-500/20 flex items-center justify-center text-emerald-400 text-sm font-bold">11</span>
              Contact
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              For questions about these Terms, please contact us:
            </p>
            <div class="mt-4 p-4 bg-neutral-800/50 rounded-xl">
              <p class="text-white font-medium">Unicorn Sport - Legal</p>
              <p class="text-neutral-400 mt-1">Email: legal@unicornsport.africa</p>
              <p class="text-neutral-400">Address: Lagos, Nigeria</p>
            </div>
          </section>

        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
useHead({
  title: 'Terms of Service | Unicorn Sport',
  meta: [
    { name: 'description', content: 'Read the Terms of Service for using the Unicorn Sport platform.' }
  ]
})

const lastUpdated = 'January 27, 2026'
</script>
