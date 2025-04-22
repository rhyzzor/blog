<template>
  <template v-if="post">
    <div>
      <section class="information">
        <h2>{{ post.title }}</h2>
        <p>
          <strong> Created at:</strong> {{ dateFormat(post.date, true) }}
        </p>
        <p>
          <strong>Reading time:</strong> {{ calculateReadingTime(post.rawbody) }}

        </p>
      </section>

      <ContentRenderer :value="post" class="content" tag="section" />
    </div>
  </template>

  <template v-else>
    <AppNotFound />
  </template>
</template>

<script lang="ts" setup>
const slug = useRoute().params.slug as string;

const { data: post } = await useAsyncData(`blog-${slug}`, () => {
	return queryCollection("blog").path(`/blog/${slug}`).first();
});

if (post.value?.ogImage) {
	defineOgImage(post.value.ogImage);
}

useHead(post.value?.head || {});
useSeoMeta(post.value?.seo || {});
</script>

<style>
section strong {
  font-weight: bold;
}

.content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing);

  img {
    width: 100%;
  }
}
</style>