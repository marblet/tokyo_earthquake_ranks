<template>
  <div>
    <v-text-field v-model="innerVal"></v-text-field>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  computed: {
    innerVal: {
      get () {
        return this.value
      },
      set (value) {
        this.$emit('update:value', value)
      }
    }
  }
}
</script>

<style>

</style>
