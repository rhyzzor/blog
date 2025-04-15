<template>
  <div class="pre">
    <pre class="pre-body" :class="$props.class"><slot/></pre>
  </div>
</template>


<script setup lang="ts">
const props = defineProps({
  code: {
    type: String,
    default: ''
  },
  language: {
    type: String,
    default: null
  },
  filename: {
    type: String,
    default: null
  },
  highlights: {
    type: Array as () => number[],
    default: () => []
  },
  meta: {
    type: String,
    default: null
  },
  class: {
    type: String,
    default: null
  }
})

</script>

<style lang="scss">
.pre {
  overflow-x: hidden;
  border-radius: 6px;
  background-color: #282c34;
  border: 1px solid  #282c34;
  color: #abb2bf;

 

  &-body {
    margin: 0;
    padding: .75rem 0 .75rem 0;
    border-bottom-left-radius: 6px;
    border-bottom-right-radius: 6px;
    overflow-x: auto;

    code {
      display: inline-block;
      width: 100%;
    }

    .line {
      padding: 0 .75rem;
      line-height: 1.6;
      

      span {
        background-color: transparent !important;
      }

      &::before {
        content: attr(line);
        padding-right: 1.25rem;
        display: inline-block;
        opacity: 0.8;
      }

    }
  }
}

pre code .line {
  display: block;
}
</style>
