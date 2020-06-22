<template>
  <div class="address-field">
    <v-text-field
      v-model="address"
      v-on:keydown.13="clicked"
      solo
      append-icon="mdi-magnify"
    ></v-text-field>
  </div>
</template>

<script>
export default {
  props: {
    inputAddress: String
  },
  data () {
    return {
      address: this.$route.query.address
    }
  },
  methods: {
    clicked () {
      // avoid redundant navigation
      if (this.address === this.$route.query.address && parseInt(this.$route.query.page) === 1) return
      this.$router.push({path: 'search', query: {address: this.address, page: 1}})
    }
  }
}
</script>

<style>
.address-field {
  width: 90%;
  margin: auto;
}
</style>
