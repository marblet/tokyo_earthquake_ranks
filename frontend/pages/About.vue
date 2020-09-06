<template>
  <div class="about-wrapper">
    <h1>This is an about page</h1>
    <h2>サイトの説明</h2>
    <h3>建物倒壊</h3>
    <CollapseDescription />
    <h3>火災</h3>
    <FireDescription />
    <h3>災害時</h3>
    <DifficultyDescription />
    <h3>総合</h3>
    <p>建物倒壊危険度・火災危険度に災害時活動困難度を加味し、地震の揺れに起因する町丁目ごとの危険性を5段階にランク付けした。</p>
  </div>
</template>

<script>
import Meta from '@/assets/mixins/Meta.js'
import FireDescription from '@/components/FireDescription.vue'
import CollapseDescription from '@/components/CollapseDescription.vue'
import DifficultyDescription from '@/components/DifficultyDescription.vue'

export default {
  name: 'About',
  components: {
    FireDescription,
    CollapseDescription,
    DifficultyDescription,
  },
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
