<template>
  <div class="search">
    <AddressForm />
    <TownList />
  </div>
</template>

<script>
// @ is an alias to /src
import AddressForm from '@/components/AddressForm.vue'
import TownList from '@/components/TownList.vue'
import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'Search',
  components: {
    AddressForm,
    TownList,
  },
  mixins: [Meta],
  data () {
    return {
      address: this.$route.query.address,
      meta: {
        title: this.$route.query.address,
        description: 'About page',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
    }
  },
  watch: {
    $route () {
      this.address = this.$route.query.address
      this.meta.title = this.address
    }
  },
}
</script>