<template>
<div class="townlist">
  <p>全{{this.matchedNum}}件中{{ numTopTown }}件目から{{ numBottomTown }}件目を表示</p>
  <v-simple-table>
    <thead>
      <tr>
        <th>地名</th>
        <th>建物倒壊危険度</th>
        <th>火災危険度</th>
        <th>災害時活動困難度</th>
        <th>総合危険度</th>
      </tr>
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
    v-model="pageNum"
    :length="length"
  ></v-pagination>
</div>
</template>

<script>
import axios from 'axios'
import TownListItem from './TownListItem'

export default {
  name: 'TownList',
  components: {
    TownListItem
  },
  data () {
    return {
      pageNum: 1,
      matchedTowns: [],
      displayNum: 20,
      matchedNum: 0,
      length: 0,
    }
  },
  computed: {
    numTopTown () {
      return Math.min((this.pageNum - 1) * this.displayNum + 1, this.matchedNum)
    },
    numBottomTown () {
      return Math.min(this.pageNum * this.displayNum, this.matchedNum)
    }
  },
  watch: {
    pageNum (newPage) {
      // need to fix 'avoided redundant navigation' error
      this.$router.push({path: 'search', query: {address: this.address, page: newPage}})//.catch(()=>{})
    },
    $route () {
      this.callAPI()
    },
  },
  created () {
    this.callAPI()
  },
  methods: {
    InfiniteLoad () {
      
    },
    callAPI () {
      this.address = this.$route.query.address
      this.pageNum = parseInt(this.$route.query.page)
      axios
        .get(`/api/towns?address=${this.address}&page=${this.pageNum}&displaynum=${this.displayNum}`)
        .then(responce => this.update(responce))
    },
    update: function (responce) {
      this.matchedNum = responce.data.matched_num
      this.length = Math.ceil(this.matchedNum / this.displayNum)
      this.matchedTowns = responce.data.matched_towns
    }
  }
}
</script>

<style lang="scss">
</style>
