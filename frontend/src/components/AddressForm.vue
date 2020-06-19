<template>
  <div>
    <v-text-field
      v-model="address"
      v-on:keydown.13="clicked"
      outlined
    ></v-text-field>
  </div>
</template>

<script>
export default {
  data () {
    return {
      address: ""
    }
  },
  methods: {
    clicked () {
      this.$emit('update:value', this.address)
    }
  }
}
</script>

<style>

</style>
