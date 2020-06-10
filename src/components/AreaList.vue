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
      v-for="area in getAreas"
      v-bind:key="area.id"
      v-bind:area="area"
    ></AreaListItem>
  </table>
  <v-pagination
    v-model="page"
    :length="length"
  ></v-pagination>

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
      inputAddress: '',
      areas: areas,
      match: [],
      page: 1,
      pageSize: 20
    }
  },
  created () {
    this.match = this.areas
    this.length = Math.ceil(this.match.length / this.pageSize)
  },
  computed: {
    getAreas: function () {
      return this.match.slice((this.page - 1) * this.pageSize, this.page * this.pageSize)
    }
  },
  components: {
    AddressForm,
    AreaListItem
  },
  watch: {
    inputAddress: function (newAddress) {
      this.match = this.areas.filter(area => area.name.indexOf(newAddress) > -1)
      this.length = Math.ceil(this.match.length / this.pageSize)
    }
  }
}
</script>

<style>

</style>
