<template>
<div>
  <AddressForm v-bind:value.sync="inputAddress"/>
  <v-btn v-on:click=this.searchAreas>Search</v-btn>
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
      .get('/api/areas?address=&page=1&displaynum=20')
      .then(responce => this.update(responce))
  },
  components: {
    AddressForm,
    AreaListItem
  },
  methods: {
    searchAreas: function () {
      axios
        .get(`/api/areas?address=${this.inputAddress}&page=1&displaynum=20`)
        .then(responce => this.update(responce))
    },
    update: function (responce) {
      this.matchedNum = responce.data.matched_num
      this.length = Math.ceil(this.matchedNum / this.pageSize)
      this.matchedAreas = responce.data.matched_areas
    }
  },
  watch: {
    page: function (newPage) {
      axios
        .get(`/api/areas?address=${this.inputAddress}&page=${newPage}&displaynum=20`)
        .then(responce => this.update(responce))
    }
  }
}
</script>

<style>

</style>
