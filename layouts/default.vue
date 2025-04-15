<template>
  <div class="container">
    <AppHeader />
    <slot />
  </div>
</template>

<script lang="ts" setup>


</script>

<style>
.container {
  display: flex;
  flex-direction: column;
  gap: var(--spacing);
  position: relative;
}
</style>