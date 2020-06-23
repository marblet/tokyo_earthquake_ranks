<template>
  <div class="home">
    <Title />
    <AddressForm />
  </div>
</template>

<script>
// @ is an alias to /src
import AddressForm from '@/components/AddressForm.vue'
import Title from '@/components/Title.vue'
import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'Home',
  components: {
    AddressForm,
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
