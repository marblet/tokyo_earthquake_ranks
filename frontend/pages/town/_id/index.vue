<template>
  <div>
    <div>
      <h2>{{ town.municipality + town.town_name }}</h2>
      <p>地盤分類：{{ town.base_class }}</p>
      <v-simple-table>
        <thead>
          <tr>
            <th></th>
            <th>危険量</th>
            <th>順位</th>
            <th>ランク</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>建物倒壊</td>
            <td>{{ town.collapse_ha }} 棟/ha</td>
            <td>{{ town.collapse_order }} / 5177</td>
            <td>{{ town.collapse_rank }}</td>
          </tr>
          <tr>
            <td>火災</td>
            <td>{{ town.fire_ha }} 棟/ha</td>
            <td>{{ town.fire_order }} / 5177</td>
            <td>{{ town.fire_rank }}</td>
          </tr>
          <tr>
            <td>活動困難</td>
            <td>{{ town.difficulty }}</td>
            <td>{{ town.difficulty_order }} / 5177</td>
            <td>{{ town.difficulty_rank }}</td>
          </tr>
          <tr>
            <td>総合</td>
            <td>{{ town.total_ha }} 棟/ha</td>
            <td>{{ town.total_order }} / 5177</td>
            <td>{{ town.total_rank }}</td>
          </tr>
        </tbody>
      </v-simple-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'TownInfo',
  mixins: [Meta],
  async asyncData ({ params }) {
    if (process.client) {
      return axios
        .get(`/api/towninfo?id=${params.id}`)
        .then(response => {
          return {
            town: response.data
          }
      })
    } else {
      return axios
        .get(`https://tokyo-earthquake-ranks.df.r.appspot.com/api/towninfo?id=${params.id}`)
        .then(response => {
          return {
            town: response.data
          }
      })
    }
    // return axios
    //   .get(`/api/towninfo?id=${params.id}`)
    //   .then(response => {
    //     return {
    //       town: response.data
    //     }
    //   }).catch(err => {
    //     console.log(err)
    //     console.log(axios.defaults.baseURL)
    //   })
  },
  head () {
    return {
      title: this.town.municipality + this.town.town_name
    }
  },
  data () {
    return {
      town: {},
      meta: {
        // title: town.municipality + town.town_name,
        description: 'Towninfo page',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
    }
  },
  methods: {
    update (response) {
      this.town = response.data
    }
  },
}
</script>

<style>
@import '@/assets/styles/loader.css'
</style>