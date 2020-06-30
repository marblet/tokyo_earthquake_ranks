<template>
  <div class="about-wrapper">
    <h1>This is an about page</h1>
    <h2>サイトの説明</h2>
    <h3>建物倒壊</h3>
    <p>{{ description.collapse.overview }}</p>
    <h3>火災</h3>
    <p>{{ description.fire.overview }}</p>
    <h3>災害時</h3>
    <p>{{ description.difficulty.overview }}</p>
    <h3>総合</h3>
    <p>{{ description.total.overview }}</p>
  </div>
</template>

<script>
import Meta from '@/assets/mixins/Meta.js'
import desc from '@/assets/jsons/description.json'

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
      },
      description: desc
    }
  }
}
</script>

<style lang="scss">
.about-wrapper {
  margin: auto;
  width: 100%;
  max-width: 1200px;
}
</style>
