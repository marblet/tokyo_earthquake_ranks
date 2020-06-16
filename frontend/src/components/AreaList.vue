<template>
<div>
  <AddressForm v-bind:value.sync="inputAddress"/>
  <p>Your input is {{ inputAddress }}</p>
  <p>Hit: {{this.matchedNum}}</p>
  <table>
    <tr>
      <td>地名</td>
      <td>総合順位</td>
      <td>総合ランク</td>
    </tr>
    <AreaListItem
      v-for="area in matchedAreas"
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
import axios from 'axios'
export default {
  name: 'AreaList',
  data () {
    return {
      inputAddress: '',
      matchedAreas: [],
      page: 1,
      pageSize: 20,
      matchedNum: 0,
      length: 0
    }
  },
  created () {
    // itinialize matchedAreas using API
    axios
      .get('http://127.0.0.1:8080/api/areas?address=&page=1&displaynum=20')
      .then(responce => this.update(responce))
  },
  components: {
    AddressForm,
    AreaListItem
  },
  methods: {
    update: function (responce) {
      this.matchedNum = responce.data.matched_num
      this.length = Math.ceil(this.matchedNum / this.pageSize)
      this.matchedAreas = responce.data.matched_areas
    }
  },
  watch: {
    inputAddress: function (newAddress) {
      axios
        .get(`http://localhost:8080/api/areas?address=${newAddress}&page=1&displaynum=20`)
        .then(responce => this.update(responce))
    },
    page: function (newPage) {
      axios
        .get(`http://localhost:8080/api/areas?address=${this.inputAddress}&page=${newPage}&displaynum=20`)
        .then(responce => this.update(responce))
    }
  }
}
</script>

<style>

</style>
