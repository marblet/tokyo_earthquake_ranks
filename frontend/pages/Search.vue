<template>
  <div class="search-wrapper">
    <AddressForm />
    <div class="townlist">
      <p>全{{ this.matchedNum }}件</p>
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
            v-for="town in towns"
            v-bind:key="town.id"
            v-bind:town="town"
          ></TownListItem>
        </tbody>
      </v-simple-table>
    </div>
    <client-only>
      <infinite-loading @infinite="infiniteHandler" spinner="bubbles">
      </infinite-loading>
    </client-only>
  </div>
</template>

<script>
import axios from 'axios'
import InfiniteLoading from 'vue-infinite-loading'
import AddressForm from '@/components/AddressForm.vue'
import TownListItem from '@/components/TownListItem.vue'
import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'Search',
  components: {
    AddressForm,
    InfiniteLoading,
    TownListItem,
  },
  mixins: [Meta],
  async asyncData ({ query, error }) {
    // should be fixed in the future
    const encodedAddress = encodeURI(query.address)
    const q = `/api/towns?address=${encodedAddress}&page=1&displaynum=20`
    const url = process.client ? q : 'https://tokyo-earthquake-ranks.df.r.appspot.com' + q
    return axios
    .get(url)
    .then(response => {
      if (response.data.Error) {
        return error({
          statusCode: 404
        })
      }
      return {
        matchedNum: response.data.matched_num,
        towns: response.data.matched_towns
      }
    })
  },
  data () {
    return {
      length: 0,
      pageNum: 1,
      displayNum: 20,
      meta: {
        title: this.$route.query.address,
        description: `${this.$route.query.address}の検索結果`,
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
    }
  },
  methods: {
    infiniteHandler ($state) {
      const address = this.$route.query.address
      this.pageNum += 1
      axios.get(`/api/towns?address=${address}&page=${this.pageNum}&displaynum=${this.displayNum}`
      ).then( response  => {
        if (response.data.matched_towns.length > 0) {
          this.towns.push(...response.data.matched_towns)
          $state.loaded()
        } else {
          $state.complete()
        }
      });
    },
    reset () {
      this.matchedNum = 0
      this.length = 0
      this.pageNum = 1
      this.towns = []
    }
  },
  watch: {
    $route () {
      this.address = this.$route.query.address
      this.meta.title = this.address
      this.reset()
    }
  },
  watchQuery: ['address']
}
</script>

<style lang="scss">
.search-wrapper {
  margin: auto;
  width: 100%;
  max-width: 1200px;
}
</style>