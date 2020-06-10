<template>
<div>
  <AddressForm v-bind:value.sync="inputAddress"/>
  <p>Your input is {{ inputAddress }}</p>
  <p>Hit: {{match.length}}</p>
  <table>
    <tr>
      <td>地名</td>
      <td>総合順位</td>
      <td>総合ランク</td>
    </tr>
    <AreaListItem
      v-for="area in match"
      v-bind:key="area.id"
      v-bind:area="area"
    ></AreaListItem>
  </table>

</div>
</template>

<script>
import AddressForm from './AddressForm'
import AreaListItem from './AreaListItem'
import areas from '../assets/all2.json'
export default {
  name: 'AreaList',
  data () {
    return {
      inputAddress: '千',
      areas: areas
    }
  },
  computed: {
    match: function () {
      return this.areas.filter(area => area.name.indexOf(this.inputAddress) > -1)
    }
  },
  components: {
    AddressForm,
    AreaListItem
  }
}
</script>

<style>

</style>
