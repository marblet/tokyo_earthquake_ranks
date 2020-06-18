<template>
  <div>
    <v-text-field v-model="address" v-on:keydown.13="clicked"></v-text-field>
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
      console.log(this.address)
    }
  }
}
</script>

<style>

</style>
