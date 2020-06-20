<template>
<div class="arealist">
  <AddressForm v-bind:value.sync="inputAddress"/>
  <p>全{{this.matchedNum}}件中{{ numTopArea }}件目から{{ numBottomArea }}件目を表示</p>
  <v-simple-table>
    <thead>
      <td>地名</td>
      <td>建物倒壊危険度</td>
      <td>火災危険度</td>
      <td>災害時活動困難度</td>
      <td>総合危険度</td>
    </thead>
    <tbody>
      <AreaListItem
        v-for="area in matchedAreas"
        v-bind:key="area.id"
        v-bind:area="area"
      ></AreaListItem>
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
import AreaListItem from './AreaListItem'
import axios from 'axios'
export default {
  name: 'AreaList',
  components: {
    AddressForm,
    AreaListItem
  },
  data () {
    return {
      inputAddress: '',
      matchedAreas: [],
      page: 1,
      displayNum: 20,
      matchedNum: 0,
      length: 0
    }
  },
  computed: {
    numTopArea () {
      return Math.min((this.page - 1) * this.displayNum + 1, this.matchedNum)
    },
    numBottomArea () {
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
      this.searchAreas()
    }
  },
  created () {
    // itinialize matchedAreas using API
    axios
      .get(`/api/areas?address=&page=1&displaynum=${this.displayNum}`)
      .then(responce => this.update(responce))
  },
  methods: {
    searchAreas: function () {
      this.page = 1
      axios
        .get(`/api/areas?address=${this.inputAddress}&page=1&displaynum=${this.displayNum}`)
        .then(responce => this.update(responce))
    },
    update: function (responce) {
      this.matchedNum = responce.data.matched_num
      this.length = Math.ceil(this.matchedNum / this.displayNum)
      this.matchedAreas = responce.data.matched_areas
    }
  }
}
</script>

<style lang="scss">
.arealist {
  width: 1200px;
  margin: auto;
  padding: {
    bottom: 10px;
    left: 5px;
    right: 5px;
    top: 10px;
  }
}
</style>
