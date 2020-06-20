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
.address-field {
  width: 90%;
  margin: auto;
}
</style>
