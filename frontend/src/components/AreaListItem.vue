<template>
  <tr>
    <td>{{ area.areaname }}</td>
    <td>{{ area.collapse_rank}}</td>
    <td>{{ area.fire_rank}}</td>
    <td>{{ area.difficulty_rank}}</td>
    <td>{{ area.total_rank }}</td>
  </tr>
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
