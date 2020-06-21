<template>
<div class="townlist">
  <AddressForm v-bind:value.sync="inputAddress"/>
  <p>全{{this.matchedNum}}件中{{ numTopTown }}件目から{{ numBottomTown }}件目を表示</p>
  <v-simple-table>
    <thead>
      <td>地名</td>
      <td>建物倒壊危険度</td>
      <td>火災危険度</td>
      <td>災害時活動困難度</td>
      <td>総合危険度</td>
    </thead>
    <tbody>
      <TownListItem
        v-for="town in matchedTowns"
        v-bind:key="town.id"
        v-bind:town="town"
      ></TownListItem>
    </tbody>
  </v-simple-table>
  <v-pagination
    v-model="page"
    :length="length"
  ></v-pagination>
</div>
</template>

<script>
import AddressForm from './AddressForm'
import axios from 'axios'
import TownListItem from './TownListItem'
export default {
  name: 'TownList',
  components: {
    AddressForm,
    TownListItem
  },
  data () {
    return {
      inputAddress: '',
      matchedTowns: [],
      page: 1,
      displayNum: 20,
      matchedNum: 0,
      length: 0
    }
  },
  computed: {
    numTopTown () {
      return Math.min((this.page - 1) * this.displayNum + 1, this.matchedNum)
    },
    numBottomTown () {
      return Math.min(this.page * this.displayNum, this.matchedNum)
    }
  },
  watch: {
    page: function (newPage) {
      axios
        .get(`/api/areas?address=${this.inputAddress}&page=${newPage}&displaynum=${this.displayNum}`)
        .then(responce => this.update(responce))
    },
    inputAddress: function () {
      this.searchTowns()
    }
  },
  created () {
    // itinialize matchedTowns using API
    axios
      .get(`/api/areas?address=&page=1&displaynum=${this.displayNum}`)
      .then(responce => this.update(responce))
  },
  methods: {
    searchTowns: function () {
      this.page = 1
      axios
        .get(`/api/areas?address=${this.inputAddress}&page=1&displaynum=${this.displayNum}`)
        .then(responce => this.update(responce))
    },
    update: function (responce) {
      this.matchedNum = responce.data.matched_num
      this.length = Math.ceil(this.matchedNum / this.displayNum)
      this.matchedTowns = responce.data.matched_areas
    }
  }
}
</script>

<style lang="scss">
</style>
