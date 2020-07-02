<template>
  <div class="municipalitylist-wrapper">
    <div class="municipalitylist">
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
import TownListItem from '@/components/TownListItem.vue'
import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'Search',
  components: {
    InfiniteLoading,
    TownListItem,
  },
  mixins: [Meta],
  async asyncData ({ params, error }) {
    // should be fixed in the future
    const q = `/api/municipality/${params.areacode}?page=1&displaynum=20`
    const url = process.client ? q : 'https://tokyo-earthquake-ranks.df.r.appspot.com' + q
    return axios
      .get(url)
      .then(response => {
        if (response.data.Error || response.data.matched_num == 0) {
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
  head () {
    return {
      title: this.towns[0].municipality
    }
  },
  data () {
    return {
      length: 0,
      pageNum: 1,
      displayNum: 20,
      meta: {
        description: 'About page',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
    }
  },
  methods: {
    infiniteHandler ($state) {
      this.pageNum += 1;
      axios.get(`/api/municipality/${this.$route.params['areacode']}?page=${this.pageNum}&displaynum=${this.displayNum}`
      ).then( response  => {
        this.matchedNum = response.data.matched_num
        if (response.data.matched_towns.length > 0) {
          this.towns.push(...response.data.matched_towns);
          $state.loaded();
        } else {
          $state.complete();
        }
      });
    },
  },
}
</script>

<style lang="scss">
.municipalitylist-wrapper {
  margin: auto;
  width: 100%;
  max-width: 1200px;
}
</style>