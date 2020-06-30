<template>
  <div class="top-wrapper">
    <div class="top-content-wrapper">
      <Title />
      <AddressForm />
      <Municipality />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import AddressForm from '@/components/AddressForm.vue'
import Municipality from '@/components/Municipality.vue'
import Title from '@/components/Title.vue'

import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'Home',
  components: {
    AddressForm,
    Municipality,
    Title,
  },
  mixins: [Meta],
  data () {
    return {
      meta: {
        title: 'Tokyo Earthquake Ranks',
        description: 'ページ個別のディスクリプション',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
    }
  }
}
</script>

<style lang="scss">
.top-wrapper {
  background: {
    image: url("~@/assets/3432480_m.jpg");
    attachment: fixed;
  }
}
.top-content-wrapper {
  margin: auto;
  width: 100%;
  max-width: 1200px;
}
</style>