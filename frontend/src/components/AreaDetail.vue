<template>
  <v-card>
    <v-card-title>{{ area.areaname }}</v-card-title>
    <v-card-text>Dialog is called</v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    area: {
      type: Object
    }
  }
}
</script>

<style>

</style>