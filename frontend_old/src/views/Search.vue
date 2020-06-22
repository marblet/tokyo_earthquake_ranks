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

export default {
  name: 'Search',
  components: {
    AddressForm,
    TownList,
  },
  data () {
    return {
      address: this.$route.query.address,
    }
  },
  watch: {
    $route () {
      this.address = this.$route.query.address
      this.$emit('updateHead')
    }
  },
  head () {
    return {
      title: {
        inner: this.address
      }
    }
  },
}
</script>