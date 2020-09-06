<template>
  <div class="address-field">
    <v-text-field
      v-model="address"
      v-on:keydown.13="entered"
      solo
      append-icon="mdi-magnify"
      placeholder="千代田区飯田橋"
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
    entered () {
      // avoid redundant navigation
      if (this.address === this.$route.query.address) return
      this.$router.push({path: 'search', query: {address: this.address}})
    }
  }
}
</script>

<style>
.address-field {
  margin: auto;
}
</style>
