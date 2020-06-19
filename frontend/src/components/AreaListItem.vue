<template>
  <tr v-on:click="dialog=true">
    <td>{{ area.areaname }}</td>
    <td>{{ area.collapse_rank}}</td>
    <td>{{ area.fire_rank}}</td>
    <td>{{ area.difficulty_rank}}</td>
    <td>{{ area.total_rank }}</td>
    <v-dialog v-model="dialog">
      <AreaDetail :area="area"></AreaDetail>
    </v-dialog>
  </tr>
</template>

<script>
import AreaDetail from './AreaDetail'
export default {
  components: {
    AreaDetail
  },
  props: {
    area: {
      type: Object
    }
  },
  data () {
    return {
      dialog: false
    }
  }
}
</script>

<style>

</style>
