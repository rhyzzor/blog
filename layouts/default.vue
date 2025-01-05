<template>
  <div class="container">
    <slot />
  </div>
</template>

<style>
:root {
  --spacing: 16px;
  --spacing-2x: calc(var(--spacing) * 2);
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html,
body {
  font-weight: 300;
  min-height: 100vh;
  line-height: 1.5;
  list-style-type: none;
  font-size: 16px;
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  width: 80vw;
  padding-top: var(--spacing);
  padding-bottom: var(--spacing);
}


.container {
  display: flex;
  flex-direction: column;
  gap: var(--spacing);
  position: relative
}

@media screen and (min-width: 768px) {
  body {
    width: 40vw;
  }
}
</style>