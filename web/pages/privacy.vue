<template>
  <div class="min-h-screen bg-neutral-950">
    <!-- Hero -->
    <div class="relative overflow-hidden">
      <div class="absolute inset-0 bg-gradient-to-b from-neutral-950 via-neutral-950 to-neutral-900"></div>
      <div class="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,.015)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,.015)_1px,transparent_1px)] bg-[size:64px_64px]"></div>
      
      <div class="relative max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 lg:pt-20 pb-12 text-center">
        <div class="inline-flex items-center gap-2 px-3 py-1.5 rounded-full bg-white/5 border border-white/10 backdrop-blur-sm mb-6">
          <svg class="w-4 h-4 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
          </svg>
          <span class="text-xs text-neutral-300">Your data, protected</span>
        </div>
        
        <h1 class="font-display text-3xl sm:text-4xl lg:text-5xl font-bold text-white mb-4">
          Privacy Policy
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
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">1</span>
              Introduction
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              Unicorn Sport ("we," "our," or "us") is committed to protecting your privacy. This Privacy Policy explains how we collect, use, disclose, and safeguard your information when you use our platform to discover and connect with African football talent.
            </p>
            <p class="text-neutral-400 leading-relaxed mt-4">
              We operate as a professional talent discovery platform, connecting verified youth football players with scouts and clubs worldwide. As we handle information about young athletes, we take extra care to ensure the highest standards of data protection.
            </p>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">2</span>
              Information We Collect
            </h2>
            
            <h3 class="text-lg font-medium text-white mt-6 mb-3">For Players (Under 18)</h3>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li>Full name, date of birth, and nationality</li>
              <li>Physical attributes (height, weight, position)</li>
              <li>Football performance data and statistics</li>
              <li>Video footage from matches and training sessions</li>
              <li>Academy or school affiliation</li>
              <li>Parental/guardian consent documentation</li>
              <li>Government ID for age verification (NIN or International Passport)</li>
            </ul>

            <h3 class="text-lg font-medium text-white mt-6 mb-3">For Scouts & Clubs</h3>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li>Contact information (name, email, phone)</li>
              <li>Organization details and credentials</li>
              <li>Payment and billing information</li>
              <li>Usage data and interaction history</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">3</span>
              How We Use Your Information
            </h2>
            <ul class="list-disc list-inside text-neutral-400 space-y-2">
              <li>Create and manage player profiles on our platform</li>
              <li>Verify player identity and age through government documents</li>
              <li>Enable scouts to discover and connect with talent</li>
              <li>Process subscription payments securely via Stripe</li>
              <li>Send important platform updates and notifications</li>
              <li>Improve our services and user experience</li>
              <li>Comply with legal obligations</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">4</span>
              Children's Privacy (COPPA & GDPR-K)
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              We are acutely aware of our responsibilities regarding children's data. Our platform primarily features youth athletes under 18. We implement the following safeguards:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li>Parental or guardian consent is required before any player profile is created</li>
              <li>We collect only information necessary for the platform's purpose</li>
              <li>Parents/guardians can request access to, correction of, or deletion of their child's data</li>
              <li>Player contact information is never publicly displayed</li>
              <li>All scout contact requests are moderated and logged</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">5</span>
              Data Security
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              We implement industry-standard security measures to protect your data:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li>SSL/TLS encryption for all data transmission</li>
              <li>Secure cloud storage with AWS</li>
              <li>Regular security audits and penetration testing</li>
              <li>Role-based access control for platform administrators</li>
              <li>Comprehensive audit logging for compliance</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">6</span>
              Data Sharing
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              We do not sell your personal information. We may share data with:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li><strong class="text-white">Subscribed Scouts:</strong> Player profiles, highlights, and approved contact information</li>
              <li><strong class="text-white">Payment Processors:</strong> Stripe for secure payment handling</li>
              <li><strong class="text-white">Cloud Services:</strong> AWS for hosting and content delivery</li>
              <li><strong class="text-white">Legal Authorities:</strong> When required by law or to protect rights</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">7</span>
              Your Rights
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              You have the right to:
            </p>
            <ul class="list-disc list-inside text-neutral-400 space-y-2 mt-4">
              <li>Access the personal data we hold about you</li>
              <li>Request correction of inaccurate information</li>
              <li>Request deletion of your data (subject to legal requirements)</li>
              <li>Withdraw consent for data processing</li>
              <li>Export your data in a portable format</li>
              <li>Lodge a complaint with a data protection authority</li>
            </ul>
          </section>

          <section class="mb-12">
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">8</span>
              Cookies & Tracking
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              We use essential cookies for platform functionality and analytics cookies to improve our services. You can manage cookie preferences through your browser settings.
            </p>
          </section>

          <section>
            <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-3">
              <span class="w-8 h-8 rounded-lg bg-primary-500/20 flex items-center justify-center text-primary-400 text-sm font-bold">9</span>
              Contact Us
            </h2>
            <p class="text-neutral-400 leading-relaxed">
              For privacy-related inquiries or to exercise your rights, please contact us:
            </p>
            <div class="mt-4 p-4 bg-neutral-800/50 rounded-xl">
              <p class="text-white font-medium">Unicorn Sport - Data Protection</p>
              <p class="text-neutral-400 mt-1">Email: privacy@unicornsport.africa</p>
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
  title: 'Privacy Policy | Unicorn Sport',
  meta: [
    { name: 'description', content: 'Learn how Unicorn Sport protects your privacy and handles your data.' }
  ]
})

const lastUpdated = 'January 27, 2026'
</script>
