<script setup>
const props = defineProps({
	title: { type: String, required: false, default: "title" },
	description: { type: String, required: false, default: "description" },
});
</script>

<template>
  <div class="h-full w-full flex items-start justify-start border-solid border-blue-500 border-[12px] bg-gray-50">
    <div class="flex items-start justify-start h-full">
      <div class="flex flex-col justify-between w-full h-full">
        <div class="content-top">
          <h1 class="text-[60px] p-20 pb-0 font-black text-left">
            {{ title }}
          </h1>
          <p class="text-2xl px-20 mt-4">
            {{ description }}
          </p>
        </div>
        <p class="text-2xl pb-10 px-20 font-bold mb-0">
          blog.rhyzzor.com
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.content-top {
  display: flex;
  flex-direction: column;
}
</style>
