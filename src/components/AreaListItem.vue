<template>
  <tr>
    <td>{{ area.name }}</td>
    <td>{{ area.order }}</td>
    <td>{{ area.rank }}</td>
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
