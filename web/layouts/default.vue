<template>
  <div class="min-h-screen flex flex-col">
    <LayoutAppHeader />
    <main class="flex-1 pt-16 lg:pt-20">
      <slot />
    </main>
    <LayoutAppFooter />
  </div>
</template>
