<template>
  <div class="about">
    <h1>This is an about page</h1>
    <h2>サイトの説明</h2>
    <h3>建物倒壊</h3>
    <h3>火災</h3>
    <h3>災害時</h3>
  </div>
</template>

<script>
import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'About',
  mixins: [Meta],
  data () {
    return {
      meta: {
        title: 'About',
        description: 'About page',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
    }
  }
}
</script>
