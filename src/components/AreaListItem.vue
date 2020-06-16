<template>
  <tr>
    <td>{{ area.areaname }}</td>
    <td>{{ area.total_order }}</td>
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
