<template>
  <tr>
    <td><nuxt-link :to="`/town/${town.id}`">{{ town.municipality + town.town_name }}</nuxt-link></td>
    <td>{{ town.collapse_rank}}</td>
    <td>{{ town.fire_rank}}</td>
    <td>{{ town.difficulty_rank}}</td>
    <td>{{ town.total_rank }}</td>
  </tr>
</template>

<script>
export default {
  props: {
    town: {
      type: Object
    }
  },
}
</script>

<style>

</style>
