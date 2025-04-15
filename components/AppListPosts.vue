<template>
  <section>
    <h2>Posts</h2>

    <div class="posts">
      <NuxtLink v-for="post in posts" :key="post.id" :to="post.path" class="post">
        <strong>
          <p>{{ `${dateFormat(post.date)} - ${post.title}` }}</p>
        </strong>

        <p>
          <strong>Description:</strong>
          {{ post.description }}
        </p>

      </NuxtLink>
    </div>
  </section>
</template>

<script lang="ts" setup>
  const { data: posts } = await useAsyncData('blog', () => queryCollection('blog').order('date', 'DESC').all())


</script>

<style>
.posts {
  display: flex;
  flex-direction: column;
  gap: var(--spacing);
}

.post {
  padding: var(--spacing-2);
  border-radius: var(--border-radius);
  background-color: var(--color-background);
  box-shadow: var(--shadow-1);
  transition: all 0.3s ease;

  strong {
    font-weight: bold;
  }
}
</style>